package domain

import (
	"encoding/json"
	"fmt"
	"github.com/move-ton/ever-client-go/util"
)

const (
	BaseCustomUrl = "https://everos.freeton.surf/"

	// NetworkQueriesProtocolHTTP -  Each GraphQL query uses separate HTTP request.
	NetworkQueriesProtocolHTTP NetworkQueriesProtocol = "HTTP"
	// NetworkQueriesProtocolWS - All GraphQL queries will be served using single web socket connection.
	NetworkQueriesProtocolWS NetworkQueriesProtocol = "WS"
)

// ClientErrorCode ...
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
		Data  []byte
		Code  uint32
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
		Type string `json:"type"`
	}

	// ClientGateway ...
	ClientGateway interface {
		Destroy()
		GetResult(string, interface{}, interface{}) error
		Request(string, interface{}) (<-chan *ClientResponse, error)
		GetResponse(string, interface{}) ([]byte, error)
		GetAPIReference() (*ResultOfGetAPIReference, error)
		Version() (*ResultOfVersion, error)
		Config() (*ClientConfig, error)
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

type (
	// NetworkQueriesProtocol - Network protocol used to perform GraphQL queries.
	NetworkQueriesProtocol string

	// ClientConfig ...
	ClientConfig struct {
		Binding          *BindingConfig `json:"binding,omitempty"`
		Network          *NetworkConfig `json:"network,omitempty"`
		Crypto           *CryptoConfig  `json:"crypto,omitempty"`
		Abi              *AbiConfig     `json:"abi,omitempty"`
		Boc              *BocConfig     `json:"boc,omitempty"`
		ProofsConfig     *ProofsConfig  `json:"proofs,omitempty"`
		LocalStoragePath string         `json:"local_storage_path,omitempty"`
	}

	// Binding config for information about Binding
	BindingConfig struct {
		Library string `json:"library,omitempty"`
		Version string `json:"version,omitempty"`
	}

	// Network - Network config.
	NetworkConfig struct {
		ServerAddress            string                 `json:"server_address,omitempty"`
		Endpoints                []string               `json:"endpoints,omitempty"`
		NetworkRetriesCount      *int                   `json:"network_retries_count,omitempty"`
		MaxReconnectTimeOut      *int                   `json:"max_reconnect_timeout,omitempty"`
		ReconnectTimeout         *int                   `json:"reconnect_timeout,omitempty"`
		MessageRetriesCount      *int                   `json:"message_retries_count,omitempty"`
		MessageProcessingTimeout *int                   `json:"message_processing_timeout,omitempty"`
		WaitForTimeout           *int                   `json:"wait_for_timeout,omitempty"`
		OutOfSyncThreshold       *int                   `json:"out_of_sync_threshold,omitempty"`
		SendingEndpointCount     *int                   `json:"sending_endpoint_count,omitempty"`
		LatencyDetectionInterval *int                   `json:"latency_detection_interval,omitempty"`
		MaxLatency               *int                   `json:"max_latency,omitempty"`
		QueryTimeout             *int                   `json:"query_timeout,omitempty"`
		QueriesProtocol          NetworkQueriesProtocol `json:"queries_protocol,omitempty"`
		FirstRempStatusTimeout   *int                   `json:"first_remp_status_timeout,omitempty"`
		NextRempStatusTimeout    *int                   `json:"next_remp_status_timeout,omitempty"`
		SignatureID              *int                   `json:"signature_id,omitempty"`
		AccessKey                string                 `json:"access_key,omitempty"`
	}

	// Crypto ...
	CryptoConfig struct {
		MnemonicDictionary  *MnemonicDictionary `toml:"mnemonic_dictionary" json:"mnemonic_dictionary,omitempty"`
		MnemonicWordCount   *int                `toml:"mnemonic_word_count" json:"mnemonic_word_count,omitempty"`
		HdKeyDerivationPath string              `toml:"hdkey_derivation_path" json:"hdkey_derivation_path,omitempty"`
	}

	// AbiConfig ...
	AbiConfig struct {
		WorkChain                          *int     `toml:"workchain" json:"workchain,omitempty"`
		MessageExpirationTimeout           *int     `toml:"message_expiration_timeout" json:"message_expiration_timeout,omitempty"`
		MessageExpirationTimeoutGrowFactor *float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor,omitempty"`
	}

	// BocConfig ...
	BocConfig struct {
		CacheMaxSize *int `toml:"cache_max_size" json:"cache_max_size,omitempty"`
	}

	// ProofsConfig ...
	ProofsConfig struct {
		CacheInLocalStorage *bool `toml:"cache_in_local_storage"  json:"cache_in_local_storage,omitempty"`
	}
)

// NewDefaultConfig create new config for connect client.
func NewDefaultConfig(address string, endPoints []string, accessKey string) ClientConfig {
	mdict := EnglishMnemonicDictionary
	config := ClientConfig{
		Network: &NetworkConfig{
			ServerAddress:            address,
			Endpoints:                endPoints,
			MessageRetriesCount:      util.IntToPointerInt(5),
			MessageProcessingTimeout: util.IntToPointerInt(40000), //ms
			WaitForTimeout:           util.IntToPointerInt(40000), //ms
			OutOfSyncThreshold:       util.IntToPointerInt(15000), //ms
			SendingEndpointCount:     util.IntToPointerInt(1),
			LatencyDetectionInterval: util.IntToPointerInt(60000), //1 min
			MaxLatency:               util.IntToPointerInt(60000), //1 min
			QueryTimeout:             util.IntToPointerInt(60000), //1 min
			FirstRempStatusTimeout:   util.IntToPointerInt(1000),  //1 sec
			NextRempStatusTimeout:    util.IntToPointerInt(5000),  //5 sec
			AccessKey:                accessKey,
		},
		Crypto: &CryptoConfig{
			MnemonicDictionary:  &mdict,
			MnemonicWordCount:   util.IntToPointerInt(12),
			HdKeyDerivationPath: "",
		},
		Abi: &AbiConfig{
			WorkChain:                          util.IntToPointerInt(0),
			MessageExpirationTimeout:           util.IntToPointerInt(40000), //ms
			MessageExpirationTimeoutGrowFactor: util.Float32ToPointerFloat32(1.5)},
	}
	return config
}

// WordCountList list length mnemonic phrases
func WordCountList() map[int]*int {
	return map[int]*int{12: util.IntToPointerInt(12), 15: util.IntToPointerInt(15), 18: util.IntToPointerInt(18), 21: util.IntToPointerInt(21), 24: util.IntToPointerInt(24)}
}

// GetMainNetBaseUrls return endpoints main net.
func GetMainNetBaseUrls() []string {
	return []string{"https://mainnet.evercloud.dev/662c63951d884a52bbc20f775b5c4baf/graphql/"}
}

// GetDevNetBaseUrls return endpoint dev net.
func GetDevNetBaseUrls() []string {
	return []string{"https://devnet.evercloud.dev/662c63951d884a52bbc20f775b5c4baf/graphql/"}
}

// GetLocalNetBaseUrls return endpoint localhost net.
func GetLocalNetBaseUrls() []string {
	return []string{"http://0.0.0.0/", "http://127.0.0.1/", "http://localhost/"}
}

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
		"InvalidData":                         36,
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
