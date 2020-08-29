package goton

/*
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo windows LDFLAGS: -L./lib/windows -lton_client
#include "./lib/client_method.h"

void callB(int request_id, tc_string_t result_json, tc_string_t error_json, int flags);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"unsafe"
)

const (
	//VersionLibSDK ...
	VersionLibSDK = "0.26.0"
)

//Client struct with client date, connect and etc.
type Client struct {
	mutx           sync.Mutex
	client         C.uint32_t
	config         *TomlConfig
	AsyncRequestID int
}

//AsyncResponse ...
type AsyncResponse struct {
	ReqID      int
	ResultJSON string
	ErrorJSON  string
	Flags      int
}

var MapStore = make(map[int]*AsyncResponse)

//InitClient create context and setup settings from file or default settings
func InitClient(config *TomlConfig) (*Client, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	client.config = config
	client.AsyncRequestID = 0

	_, err = client.Request(Setup(config))
	if err != nil {
		client.Destroy()
		return nil, err
	}

	return client, nil
}

func (client *Client) GetResp(resp int) *AsyncResponse {

	var mapReq *AsyncResponse
	for {
		if MapStore[resp].Flags != 1 {
			continue
		} else {
			mapReq = MapStore[resp]
			break
		}
	}
	return mapReq
}

//NewClient create connect node
func NewClient() (*Client, error) {

	client := Client{
		client: C.tc_create_context(),
	}

	if client.client == C.uint32_t(0) {
		return &client, errors.New("Context don't connect")
	}

	return &client, nil
}

//Destroy disconnect node
func (client *Client) Destroy() {
	client.mutx.Lock()
	defer client.mutx.Unlock()
	C.tc_destroy_context(client.client)
}

func (client *Client) Request(method, params string) (string, error) {
	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_t{content: paramsCS, len: C.uint32_t(len(params))}

	tcResponseHandle := C.tc_json_request(client.client, param1, param2)
	defer C.tc_destroy_json_response(tcResponseHandle)

	tcResponse := C.tc_read_json_response(tcResponseHandle)

	resultJSON := tcResponse.result_json
	errorJSON := tcResponse.error_json

	if errorJSON.len > 0 {
		return "", errors.New(converToStringGo(errorJSON.content, C.int(errorJSON.len)))
	}

	return converToStringGo(resultJSON.content, C.int(resultJSON.len)), nil
}

func (client *Client) RequestAsync(method, params string) int {
	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_t{content: paramsCS, len: C.uint32_t(len(params))}

	res := &AsyncResponse{}
	client.mutx.Lock()
	client.AsyncRequestID++
	res.ReqID = client.AsyncRequestID
	client.mutx.Unlock()
	MapStore[res.ReqID] = res
	go C.tc_json_request_async(client.client, param1, param2, C.int(res.ReqID), C.OnResult(C.callB))
	return res.ReqID
}

//export callB
func callB(requestID C.int, resultJSON C.tc_string_t, errorJSON C.tc_string_t, flags C.int) {
	reg := MapStore[int(requestID)]
	reg.ResultJSON = converToStringGo(resultJSON.content, C.int(resultJSON.len))
	reg.ErrorJSON = converToStringGo(errorJSON.content, C.int(errorJSON.len))
	reg.Flags = int(flags)
}

func converToStringGo(valueString *C.char, valueLen C.int) string {
	return deleteQuotesLR(C.GoStringN(valueString, valueLen))
}

func deleteQuotesLR(val string) string {
	return strings.TrimRight(strings.TrimLeft(val, `"`), `"`)
}

//Version ...
func Version() (string, string) {
	return "version", ""
}

//Setup
func Setup(config *TomlConfig) (string, string) {
	req, err := json.Marshal(&config)
	if err != nil {
		err = errors.New("Error conver to config in json")
		return "", ""
	}

	return "setup", string(req)
}
