package goton

/*
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo windows LDFLAGS: -L./lib/windows -lton_client
#include "./lib/client_method.h"
*/
import "C"
import (
	"encoding/json"
	"errors"
	"unsafe"
)

const (
	//VersionLibSDK ...
	VersionLibSDK = `"0.25.3"`
)

//Client struct with client date, connect and etc.
type Client struct {
	client C.uint32_t
}

//TCStringT struct for conver struct between C and Go
type TCStringT struct {
	Field C.tc_string_t
}

//InitClient create context and setup settings from file or default settings
func InitClient(config *TomlConfig) (*Client, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_, err = client.setup(config)
	if err != nil {
		client.Destroy()
		return nil, err
	}

	return client, nil

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
	C.tc_destroy_context(client.client)
}

func (client *Client) request(method, params string) (string, error) {

	methodR := C.CString(method)
	defer C.free(unsafe.Pointer(methodR))
	param1 := TCStringT{Field: C.tc_string_t{content: methodR, len: C.uint32_t(len(method))}}

	paramsR := C.CString(params)
	defer C.free(unsafe.Pointer(paramsR))
	param2 := TCStringT{Field: C.tc_string_t{content: paramsR, len: C.uint32_t(len(params))}}

	tcResponseHandle := C.tc_json_request(client.client, param1.Field, param2.Field)
	defer C.tc_destroy_json_response(tcResponseHandle)

	tcResponse := C.tc_read_json_response(tcResponseHandle)

	resultJSON := tcResponse.result_json
	errorJSON := tcResponse.error_json

	if errorJSON.len > 0 {
		return "", errors.New(C.GoString(errorJSON.content))
	}

	return C.GoString(resultJSON.content), nil
}

//Version ...
func (client *Client) Version() (result string, err error) {
	return client.request("version", "")
}

//setup
func (client *Client) setup(config *TomlConfig) (result string, err error) {
	req, err := json.Marshal(&config)
	if err != nil {
		err = errors.New("Error conver to config in json")
		return
	}
	result, err = client.request("setup", string(req))
	return
}
