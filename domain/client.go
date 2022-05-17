package domain

import (
	"encoding/json"
	"fmt"
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
		Code  uint32
		Data  []byte
		Error error
	}

	// AppRequestResult ...
	AppRequestResult struct {
		ValueEnumType interface{}
	}

	// AppRequestResultError ...
	AppRequestResultError struct {
		Text string `json:"text"`
	}

	// AppRequestResultOk ...
	AppRequestResultOk struct {
		Result json.RawMessage `json:"result"`
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

	// ParamsOfAppRequest ...
	ParamsOfAppRequest struct {
		AppRequestID int             `json:"app_request_id"`
		RequestData  json.RawMessage `json:"request_data"`
	}

	// EnumType ...
	EnumType struct {
		Type string `json"type"`
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

	// AppPasswordProvider ...
	AppPasswordProvider interface {
		GetPassword(ParamsOfAppPasswordProviderGetPassword) (ResultOfAppPasswordProviderGetPassword, error)
	}

	// AppSigningBox ...
	AppSigningBox interface {
		GetPublicKey() (ResultOfAppSigningBoxGetPublicKey, error)
		Sign(ParamsOfAppSigningBoxSign) (ResultOfAppSigningBoxSign, error)
	}

	// AppEncryptionBox ...
	AppEncryptionBox interface {
		GetInfo() (ResultOfAppEncryptionBoxGetInfo, error)
		Encrypt(ParamsOfAppEncryptionBoxEncrypt) (ResultOfAppEncryptionBoxEncrypt, error)
		Decrypt(ParamsOfAppEncryptionBoxDecrypt) (ResultOfAppEncryptionBoxDecrypt, error)
	}

	//AppDebotBrowser ...
	AppDebotBrowser interface {
		Log(ParamsOfAppDebotBrowserLog) error
		Switch(ParamsOfAppDebotBrowserSwitch) error
		SwitchCompleted(ParamsOfAppDebotBrowserSwitchCompleted) error
		ShowAction(ParamsOfAppDebotBrowserShowAction) error
		Input(ParamsOfAppDebotBrowserInput) (ResultOfAppDebotBrowserInput, error)
		GetSigningBox(ParamsOfAppDebotBrowserGetSigningBox) (ResultOfAppDebotBrowserGetSigningBox, error)
		InvokeDebot(ParamsOfAppDebotBrowserInvokeDebot) (ResultOfAppDebotBrowserInvokeDebot, error)
		Send(ParamsOfAppDebotBrowserSend) error
		Approve(ParamsOfAppDebotBrowserApprove) (ResultOfAppDebotBrowserApprove, error)
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
		"InvalidHandle":                       34,
		"LocalStorageError":                   35,
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

func (aRR *AppRequestResult) MarshalJSON() ([]byte, error) {
	switch value := (aRR.ValueEnumType).(type) {
	case AppRequestResultError:
		return json.Marshal(struct {
			Type string `json:"type"`
			AppRequestResultError
		}{"Error", value})
	case AppRequestResultOk:
		return json.Marshal(struct {
			Type string `json:"type"`
			AppRequestResultOk
		}{"Ok", value})
	default:
		return nil, fmt.Errorf("unsupported type for AppRequestResult %v", aRR.ValueEnumType)
	}
}

func (aRR *AppRequestResult) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Error":
		var valueEnum AppRequestResultError
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		aRR.ValueEnumType = valueEnum
	case "Ok":
		var valueEnum AppRequestResultOk
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		aRR.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for AppRequestResult %v", typeD.Type)
	}
	return nil
}

// NewAppRequestResult ...
func NewAppRequestResult(value interface{}) *AppRequestResult {
	return &AppRequestResult{ValueEnumType: value}
}
