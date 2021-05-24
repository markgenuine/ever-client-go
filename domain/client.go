package domain

import (
	"encoding/json"
	"fmt"
)

const (
	// AppRequestResultTypeError ...
	AppRequestResultTypeError AppRequestResultType = "Error"

	// AppRequestResultTypeOk ...
	AppRequestResultTypeOk AppRequestResultType = "Ok"
)

//ClientErrorCode ...
var ClientErrorCode map[string]int

type (

	// ClientError ...
	ClientError struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}

	// ClientResponse ...
	ClientResponse struct {
		Code  int
		Data  []byte
		Error error
	}

	// AppRequestResultType ...
	AppRequestResultType string

	// AppRequestResult ...
	AppRequestResult struct {
		Type   AppRequestResultType `json:"type"`
		Text   string               `json:"text,omitempty"`
		Result interface{}          `json:"result,omitempty"`
	}

	// ResultOfVersion ...
	ResultOfVersion struct {
		Version string `json:"version"`
	}

	// ResultOfGetAPIReference ...
	ResultOfGetAPIReference struct {
		API *API `json:"api"`
	}

	// API ...
	API struct {
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

	// ResultOfBuildInfo ...
	ResultOfBuildInfo struct {
		BuildNumber  int                    `json:"build_number"`
		Dependencies []*BuildInfoDependency `json:"dependencies"`
	}

	// BuildInfoDependency ...
	BuildInfoDependency struct {
		Name      string `json:"name"`
		GitCommit string `json:"git_commit"`
	}

	// ParamsOfResolveAppRequest ...
	ParamsOfResolveAppRequest struct {
		AppRequestID int               `json:"app_request_id"`
		Result       *AppRequestResult `json:"result"`
	}

	//ParamsOfAppRequest ...
	ParamsOfAppRequest struct {
		AppRequestID int             `json:"app_request_id"`
		RequestData  json.RawMessage `json:"request_data"`
	}

	// ClientGateway ...
	ClientGateway interface {
		Destroy()
		GetResult(string, interface{}, interface{}) error
		Request(string, interface{}) (<-chan *ClientResponse, error)
		GetResponse(string, interface{}) ([]byte, error)
		GetAPIReference() (*ResultOfGetAPIReference, error)
		Version() (*ResultOfVersion, error)
		GetBuildInfo() (*ResultOfBuildInfo, error)
		ResolveAppRequest(*ParamsOfResolveAppRequest) error
	}

	//AppSigningBox ...
	AppSigningBox interface {
		Request(ParamsOfAppSigningBox) (ResultOfAppSigningBox, error)
		Notify(ParamsOfAppSigningBox)
	}

	//AppDebotBrowser ...
	AppDebotBrowser interface {
		Request(ParamsOfAppDebotBrowser) (ResultOfAppDebotBrowser, error)
		Notify(ParamsOfAppDebotBrowser)
	}
)

func init() {
	// ClientErrorCode - list with error client.
	ClientErrorCode = map[string]int{
		"NotImplemented":                      1,
		"InvalidHex":                          2,
		"InvalidBase64":                       3,
		"InvalidAddress":                      4,
		"CallbackParamsCantBeConvertedToJson": 5,
		"WebsocketConnectError":               6,
		"WebsocketReceiveError":               7,
		"WebsocketSendError":                  8,
		"HttpClientCreateError":               9,
		"HttpRequestCreateError":              10,
		"HttpRequestSendError":                11,
		"HttpRequestParseError":               12,
		"CallbackNotRegistered":               13,
		"NetModuleNotInit":                    14,
		"InvalidConfig":                       15,
		"CannotCreateRuntime":                 16,
		"InvalidContextHandle":                17,
		"CannotSerializeResult":               18,
		"CannotSerializeError":                19,
		"CannotConvertJsValueToJson":          20,
		"CannotReceiveSpawnedResult":          21,
		"SetTimerError":                       22,
		"InvalidParams":                       23,
		"ContractsAddressConversionFailed":    24,
		"UnknownFunction":                     25,
		"AppRequestError":                     26,
		"NoSuchRequest":                       27,
		"CanNotSendRequestResult":             28,
		"CanNotReceiveRequestResult":          29,
		"CanNotParseRequestResult":            30,
		"UnexpectedCallbackResponse":          31,
		"CanNotParseNumber":                   32,
		"InternalError":                       33,
	}
}

// DynBufferForResponses ...
func DynBufferForResponses(in <-chan *ClientResponse) <-chan *ClientResponse {
	out := make(chan *ClientResponse, 1)
	var storage []*ClientResponse
	go func() {
		defer close(out)
		for {
			if len(storage) == 0 {
				item, ok := <-in
				if !ok {
					return
				}
				storage = append(storage, item)

				continue
			}

			select {
			case item, ok := <-in:
				if ok {
					storage = append(storage, item)
				} else {
					for _, item := range storage {
						out <- item
					}

					return
				}
			case out <- storage[0]:
				if len(storage) == 1 {
					storage = nil
				} else {
					storage = storage[1:]
				}
			}
		}
	}()

	return out
}

// HandleEvents ...
func HandleEvents(responses <-chan *ClientResponse, callback EventCallback, result interface{}) error {
	for r := range responses {
		switch r.Code {
		case 100:
			event := &ProcessingEvent{}
			if err := json.Unmarshal(r.Data, event); err != nil {
				panic(err)
			}
			callback(event)
		case 1:
			return r.Error
		case 0:
			if err := json.Unmarshal(r.Data, result); err != nil {
				panic(err)
			}

			return nil
		default:
			panic(fmt.Errorf("unknown response type code %v", r.Code))
		}
	}

	return nil
}

//AppRequestResultError - Variant constructor Application request result error.
func AppRequestResultError(text string) *AppRequestResult {
	return &AppRequestResult{Type: "Error", Text: text}
}

//AppRequestResultOk - Variant constructor Application request result Ok.
func AppRequestResultOk(result json.RawMessage) *AppRequestResult {
	return &AppRequestResult{Type: "Ok", Result: result}
}
