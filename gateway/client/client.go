package client

/*
//#cgo darwin LDFLAGS: -L../../lib/darwin -lton_client
//#cgo linux LDFLAGS: -L../../lib/linux -lton_client
//#cgo windows LDFLAGS: -L../../lib/windows -lton_client
#cgo darwin LDFLAGS: -L${SRCDIR}/lib/darwin -lton_client
#cgo linux LDFLAGS: -L${SRCDIR}/lib/linux -lton_client

#include "client_method.h"
void callB(int request_id, tc_string_data_t paramsJson, int response_type, bool finished);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/move-ton/ton-client-go/domain"
)

const (
	// VersionLibSDK ...
	VersionLibSDK = "1.14.0"
)

var mainStore = NewStore()

type clientGateway struct {
	client      C.uint32_t
	config      domain.Config
	closeCanals chan struct{}
}

func tcStringData(in []byte) C.tc_string_data_t {
	return C.tc_string_data_t{
		len:     C.uint32_t(len(in)),
		content: C.CString(string(in)),
	}
}

func tcStringToByte(data C.tc_string_data_t) []byte {
	if data.len == 0 {
		return nil
	}

	return C.GoBytes(unsafe.Pointer(data.content), C.int(data.len))
}

// NewClientGateway ...
func NewClientGateway(config domain.Config) (domain.ClientGateway, error) {
	cc := clientGateway{
		config:      config,
		closeCanals: make(chan struct{}),
	}

	configTrf, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	handler := C.tc_create_context(tcStringData(configTrf))
	defer C.tc_destroy_string(handler)
	response := tcStringToByte(C.tc_read_string(handler))

	var skdResponse SDKResponse
	err = json.Unmarshal(response, &skdResponse)
	if err != nil {
		return nil, err
	}
	if skdResponse.Error != nil {
		fmt.Println(skdResponse.Error)
	}
	cc.client = C.uint32_t(skdResponse.Result)

	return &cc, nil
}

func (c *clientGateway) Destroy() {
	C.tc_destroy_context(c.client)
}

//export callB
func callB(requestIDin C.int, paramsJSON C.tc_string_data_t, responseTypein C.int, finishedin C.bool) {
	requestID := int(requestIDin)
	params := C.GoBytes(unsafe.Pointer(paramsJSON.content), C.int(paramsJSON.len))
	responseType := int(responseTypein)
	finished := bool(finishedin)

	responses, closeSignal, isFound := mainStore.GetChannels(requestID, finished)
	if !isFound {
		return
	}

	if responseType == 2 {
		if finished {
			close(responses)
		}

		return
	}

	select {
	case responses <- newResponse(params, responseType):
		if finished {
			close(responses)
		}
	case <-closeSignal:
		close(responses)
		mainStore.DeleteRequestID(requestID)
	}
}

func newResponse(rawBytes []byte, responseType int) *domain.ClientResponse {
	res := &domain.ClientResponse{
		Code: responseType,
	}
	if responseType == 1 {
		res.Error = errors.New(string(rawBytes))
	} else if responseType == 0 || responseType == 100 {
		res.Data = rawBytes
	}

	return res
}

func (c *clientGateway) GetResult(method string, paramIn interface{}, resultStruct interface{}) error {
	rawData, err := c.GetResponse(method, paramIn)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawData, resultStruct)
}

func (c *clientGateway) Request(method string, paramIn interface{}) (<-chan *domain.ClientResponse, error) {
	var (
		rawBody []byte
		err     error
	)

	if paramIn != nil {
		rawBody, err = json.Marshal(paramIn)
		if err != nil {
			return nil, err
		}
	}

	responsChan := make(chan *domain.ClientResponse, 1)
	requestID := mainStore.SetChannels(responsChan, c.closeCanals)
	C.tc_request(c.client, tcStringData([]byte(method)), tcStringData(rawBody), C.uint32_t(requestID), C.tc_response_handler_t(C.callB))

	return responsChan, nil
}

func (c *clientGateway) GetResponse(method string, paramIn interface{}) ([]byte, error) {
	responsChan, err := c.Request(method, paramIn)
	if err != nil {
		return nil, err
	}
	var data []byte

	for {
		select {
		case r, ok := <-responsChan:
			if !ok {
				return data, err
			}
			if r.Error != nil && err == nil {
				err = r.Error
			}
			if r.Data != nil && data == nil {
				data = r.Data
			}
		case <-c.closeCanals:
			return nil, errors.New("channels is closed")
		}
	}
}
