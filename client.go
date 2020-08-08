package goton

/*
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#include "./lib/client_method.h"
*/
import "C"
import (
	"encoding/json"
	"errors"
	"unsafe"
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

//NewClient create context
func NewClient() (*Client, error) {

	client := Client{
		client: C.tc_create_context(),
	}

	if client.client != C.uint32_t(1) {
		return &client, errors.New("Context don't connect")
	}

	return &client, nil

}

//Destroy disconnect context
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

//Version get version library rust sdk
func (client *Client) Version() (string, error) {
	return client.request("version", "")
}

//Setup set parameters connection
func (client *Client) setup(config *TomlConfig) (string, error) {
	req, err := json.Marshal(&config)
	if err != nil {
		return "", errors.New("Error conver to config in json!")
	}

	result, err := client.request("setup", string(req))

	return result, err
}
