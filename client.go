package goton

/*
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo windows LDFLAGS: -L./lib/windows -lton_client
#include "./lib/client_method.h"

void callB(int request_id, tc_string_data_t paramsJson, int response_type, bool finished);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"sync"
	"unsafe"
)

const (
	// VersionLibSDK ...
	VersionLibSDK = "1.0.0"
)

// Client struct with client date, connect and etc.
type Client struct {
	mutx           sync.Mutex
	client         C.uint32_t
	config         *TomlConfig
	AsyncRequestID int
}

// ResultOfVersion ...
type ResultOfVersion struct {
	Version string
}

// ResultOfGetAPIReference ...
type ResultOfGetAPIReference struct {
	API API `json:"api"`
}

// API ...
type API struct {
	Modules []struct {
		Description string `json:"description"`
		Functions   []struct {
			Description interface{}   `json:"description"`
			Errors      interface{}   `json:"errors"`
			Name        string        `json:"name"`
			Params      []interface{} `json:"params"`
			Result      struct {
				Ref string `json:"ref"`
			} `json:"result"`
			Summary interface{} `json:"summary"`
		} `json:"functions"`
		Name    string `json:"name"`
		Summary string `json:"summary"`
		Types   []struct {
			Description interface{} `json:"description"`
			Name        string      `json:"name"`
			Struct      []struct {
				Description interface{} `json:"description"`
				Name        string      `json:"name"`
				Ref         string      `json:"ref"`
				Summary     interface{} `json:"summary"`
			} `json:"struct"`
			Summary interface{} `json:"summary"`
		} `json:"types"`
	} `json:"modules"`
	Version string `json:"version"`
}

// AsyncResponse ...
type AsyncResponse struct {
	ReqID        int
	Params       string
	ResponseType int
	Finished     bool
}

var (
	MapStore = make(map[int]*AsyncResponse)
	// MSMu     sync.Mutex
)

// InitClient create context and setup settings from file or default settings
func InitClient(config *TomlConfig) (*Client, error) {
	client := Client{}
	client.config = config
	client.AsyncRequestID = 0

	configTrf, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	contM := C.CString(string(configTrf))
	defer C.free(unsafe.Pointer(contM))
	param1 := C.tc_string_data_t{content: contM, len: C.uint32_t(len(configTrf))}

	response := C.tc_create_context(param1)
	responseStr := C.tc_read_string(response)
	defer C.tc_destroy_string(response)

	stringGo := converToStringGo(responseStr.content, C.int(responseStr.len))

	var resultResp map[string]interface{}
	json.Unmarshal([]byte(stringGo), &resultResp)
	if _, ok := resultResp["error"]; ok {
		return &client, errors.New(stringGo)
	} else if elem, ok := resultResp["result"]; ok {
		client.client = C.uint32_t(elem.(float64))
		if client.client == C.uint32_t(0) {
			return &client, errors.New("Context don't connect")
		}
	}
	return &client, nil
}

// GetResp ...
func (client *Client) GetResp(resp int) (string, error) {
	var mapReq *AsyncResponse
	for {
		nowInd := MapStore[resp]
		if !((nowInd.ResponseType == 0 || nowInd.ResponseType == 1) && nowInd.Finished) {
			continue
		} else {
			mapReq = nowInd
			break
		}
	}
	if nowInd.ResponseType == 1 {
		return "", errors.New(mapReq.Params)
	}
	return mapReq.Params, nil
}

// Destroy disconnect node
func (client *Client) Destroy() {
	client.mutx.Lock()
	defer client.mutx.Unlock()
	C.tc_destroy_context(client.client)
}

func (client *Client) Request(method, params string) (string, error) {
	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_data_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_data_t{content: paramsCS, len: C.uint32_t(len(params))}

	tcResponseHandle := C.tc_request_sync(client.client, param1, param2)
	defer C.tc_destroy_string(tcResponseHandle)
	tcResponse := C.tc_read_string(tcResponseHandle)
	stringGo := converToStringGo(tcResponse.content, C.int(tcResponse.len))
	var resultResp map[string]interface{}
	json.Unmarshal([]byte(stringGo), &resultResp)

	if _, ok := resultResp["error"]; ok {
		return "", errors.New(stringGo)
	} else if elem, ok := resultResp["result"]; ok {
		jsonbody, err := json.Marshal(elem)
		return string(jsonbody), err
	}

	return "", nil
}

// RequestAsync ...
func (client *Client) RequestAsync(method, params string) int {
	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_data_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_data_t{content: paramsCS, len: C.uint32_t(len(params))}

	res := &AsyncResponse{}
	client.mutx.Lock()
	client.AsyncRequestID++
	res.ReqID = client.AsyncRequestID
	client.mutx.Unlock()
	MapStore[res.ReqID] = res
	C.tc_request(client.client, param1, param2, C.uint32_t(res.ReqID), C.tc_response_handler_t(C.callB))
	return res.ReqID
}

//export callB
func callB(requestID C.int, paramsJSON C.tc_string_data_t, responseType C.int, finished C.bool) {
	// respNow := int(responseType)
	// finishedNow := bool(finished)
	// if !(respNow == 0 || respNow == 1) && !finishedNow {
	// 	return
	// }

	// MSMu.Lock()
	reg := MapStore[int(requestID)]
	reg.Params = converToStringGo(paramsJSON.content, C.int(paramsJSON.len))
	reg.ResponseType = int(responseType)
	reg.Finished = bool(finished)
	// MSMu.UnLock()
}

func converToStringGo(valueString *C.char, valueLen C.int) string {
	return C.GoStringN(valueString, valueLen)
}

// Version ...
func Version() (string, string) {
	return "client.version", ""
}

// VersionResult ...
func VersionResult(resp string, err error) (*ResultOfVersion, error) {
	if err != nil {
		return nil, err
	}
	resultStruct := &ResultOfVersion{}
	err = json.Unmarshal([]byte(resp), resultStruct)
	return resultStruct, err
}

// GetAPIReference ...
func GetAPIReference() (string, string) {
	return "client.get_api_reference", ""
}

// GetAPIReferenceResult ...
func GetAPIReferenceResult(resp string, err error) (*ResultOfGetAPIReference, error) {
	if err != nil {
		return nil, err
	}
	resultStruct := &ResultOfGetAPIReference{}
	err = json.Unmarshal([]byte(resp), resultStruct)
	return resultStruct, err
}
