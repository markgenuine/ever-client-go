package domain

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const (
	// ContractAbiType ...
	ContractAbiType AbiType = "Contract"

	// JSONAbiType ...
	JSONAbiType AbiType = "Json"

	// HandleAbiType ...
	HandleAbiType AbiType = "Handle"

	// SerializedAbiType ...
	SerializedAbiType AbiType = "Serialized"

	// MessageBodyTypeInput - Message contains the input of the ABI function.
	MessageBodyTypeInput MessageBodyType = "Input"

	// MessageBodyTypeOutput - Message contains the output of the ABI function.
	MessageBodyTypeOutput MessageBodyType = "Output"

	// MessageBodyTypeInternalOutput - Message contains the input of the imported ABI function.
	//Occurs when contract sends an internal message to other contract.
	MessageBodyTypeInternalOutput MessageBodyType = "InternalOutput"

	// MessageBodyTypeEvent - Message contains the input of the ABI event.
	MessageBodyTypeEvent MessageBodyType = "Event"
)

// AbiErrorCode ...
var AbiErrorCode map[string]int

type (

	// AbiType ...
	AbiType string

	// Abi ...
	Abi struct {
		Type  AbiType     `json:"type"`
		Value interface{} `json:"value"`
	}

	// AbiHandle ...
	AbiHandle int

	// FunctionHeader ...
	FunctionHeader struct {
		Expire *int     `json:"expire,omitempty"`
		Time   *big.Int `json:"time,omitempty"`
		PubKey string   `json:"pubkey,omitempty"`
	}

	// CallSet ...
	CallSet struct {
		FunctionName string          `json:"function_name"`
		Header       *FunctionHeader `json:"header,omitempty"`
		Input        interface{}     `json:"input,omitempty"`
	}

	// DeploySet ...
	DeploySet struct {
		Tvc           string      `json:"tvc"`
		WorkchainID   *int        `json:"workchain_id,omitempty"`
		InitialData   interface{} `json:"initial_data,omitempty"`
		InitialPubKey string      `json:"initial_pubkey,omitempty"`
	}

	// SignerNone No keys are provided. Creates an unsigned message.
	SignerNone struct{}

	// SignerExternal Only public key is provided to generate unsigned message and data_to_sign which can be signed later.
	SignerExternal struct {
		PublicKey string `json:"public_key"`
	}

	// SignerKeys Key pair is provided for signing
	SignerKeys struct {
		Keys *KeyPair `json:"keys"`
	}

	// SignerSigningBox Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
	SignerSigningBox struct {
		Handle SigningBoxHandle `json:"handle"`
	}

	// Signer ...
	Signer struct {
		ValueEnumType interface{}
	}

	// MessageBodyType ...
	MessageBodyType string

	// StateInitSourceMessage Deploy message.
	StateInitSourceMessage struct {
		Source *MessageSource `json:"source"`
	}

	// StateInitSourceStateInit State init data
	StateInitSourceStateInit struct {
		Code    string `json:"code"`
		Data    string `json:"data"`
		Library string `json:"library,omitempty"`
	}

	// StateInitSourceTvc Content of the TVC file. Encoded in base64.
	StateInitSourceTvc struct {
		Tvc        string           `json:"tvc"`
		PublicKey  string           `json:"public_key,omitempty"`
		InitParams *StateInitParams `json:"init_params,omitempty"`
	}

	StateInitSource struct {
		ValueEnumType interface{}
	}

	// StateInitParams ...
	StateInitParams struct {
		Abi   Abi         `json:"abi"`
		Value interface{} `json:"value"`
	}

	// MessageSourceEncoded ...
	MessageSourceEncoded struct {
		Message string `json:"message"`
		Abi     *Abi   `json:"abi,omitempty"`
	}

	// MessageSourceEncodingParams ...
	MessageSourceEncodingParams struct {
		*ParamsOfEncodeMessage
	}

	// MessageSource ...
	MessageSource struct {
		ValueEnumType interface{}
	}

	// AbiParam ...
	AbiParam struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiEvent ...
	AbiEvent struct {
		Name   string      `json:"name"`
		Inputs []*AbiParam `json:"inputs"`
		ID     string      `json:"id,omitempty"`
	}

	// AbiData ...
	AbiData struct {
		Key        int         `json:"key"`
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiFunctions ...
	AbiFunctions struct {
		Name    string      `json:"name"`
		Inputs  []*AbiParam `json:"inputs"`
		Outputs []*AbiParam `json:"outputs"`
		ID      string      `json:"id,omitempty"`
	}

	// AbiContract ...
	AbiContract struct {
		AbiVersion *int            `json:"ABI version,omitempty"`
		Version    string          `json:"version,omitempty"`
		Header     []string        `json:"header,omitempty"`
		Functions  []*AbiFunctions `json:"functions,omitempty"`
		Events     []*AbiEvent     `json:"events,omitempty"`
		Data       []*AbiData      `json:"data,omitempty"`
		Fields     []*AbiParam     `json:"fields,omitempty"`
	}

	// ParamsOfEncodeMessageBody ...
	ParamsOfEncodeMessageBody struct {
		Abi                *Abi     `json:"abi"`
		CallSet            *CallSet `json:"call_set"`
		IsInternal         bool     `json:"is_internal"`
		Signer             *Signer  `json:"signer"`
		ProcessingTryIndex *int     `json:"processing_try_index,omitempty"`
	}

	// ResultOfEncodeMessageBody ...
	ResultOfEncodeMessageBody struct {
		Body       string `json:"body"`
		DataToSign string `json:"data_to_sign,omitempty"`
	}

	// ParamsOfAttachSignatureToMessageBody ...
	ParamsOfAttachSignatureToMessageBody struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	// ResultOfAttachSignatureToMessageBody ...
	ResultOfAttachSignatureToMessageBody struct {
		Body string `json:"body"`
	}

	// ParamsOfEncodeMessage ...
	ParamsOfEncodeMessage struct {
		Abi                *Abi       `json:"abi"`
		Address            string     `json:"address,omitempty"`
		DeploySet          *DeploySet `json:"deploy_set,omitempty"`
		CallSet            *CallSet   `json:"call_set,omitempty"`
		Signer             *Signer    `json:"signer"`
		ProcessingTryIndex *int       `json:"processing_try_index"`
	}

	// ResultOfEncodeMessage ...
	ResultOfEncodeMessage struct {
		Message    string `json:"message"`
		DataToSign string `json:"data_to_sign,omitempty"`
		Address    string `json:"address"`
		MessageID  string `json:"message_id"`
	}

	// ParamsOfEncodeInternalMessage ...
	ParamsOfEncodeInternalMessage struct {
		Abi        *Abi       `json:"abi,omitempty"`
		Address    string     `json:"address,omitempty"`
		SrcAddress string     `json:"src_address,omitempty"`
		DeploySet  *DeploySet `json:"deploy_set,omitempty"`
		CallSet    *CallSet   `json:"call_set,omitempty"`
		Value      string     `json:"value"`
		Bounce     *bool      `json:"bounce"`
		EnableIhr  *bool      `json:"enable_ihr"`
	}

	// ResultOfEncodeInternalMessage ...
	ResultOfEncodeInternalMessage struct {
		Message   string `json:"message"`
		Address   string `json:"address"`
		MessageID string `json:"message_id"`
	}

	// ParamsOfAttachSignature ...
	ParamsOfAttachSignature struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	// ResultOfAttachSignature ...
	ResultOfAttachSignature struct {
		Message   string `json:"message"`
		MessageID string `json:"message_id"`
	}

	// ParamsOfDecodeMessage ...
	ParamsOfDecodeMessage struct {
		Abi     *Abi   `json:"abi"`
		Message string `json:"message"`
	}

	// DecodedMessageBody ...
	DecodedMessageBody struct {
		BodyType MessageBodyType `json:"body_type"`
		Name     string          `json:"name"`
		Value    json.RawMessage `json:"value,omitempty"`
		Header   *FunctionHeader `json:"header,omitempty"`
	}

	// ParamsOfDecodeMessageBody ...
	ParamsOfDecodeMessageBody struct {
		Abi        *Abi   `json:"abi"`
		Body       string `json:"body"`
		IsInternal bool   `json:"is_internal"`
	}

	// ParamsOfEncodeAccount ...
	ParamsOfEncodeAccount struct {
		StateInit   *StateInitSource `json:"state_init"`
		Balance     *big.Int         `json:"balance,omitempty"`
		LastTransLt *big.Int         `json:"last_trans_lt,omitempty"`
		LastPaid    *int             `json:"last_paid,omitempty"`
		BocCache    *BocCacheType    `json:"boc_cache,omitempty"`
	}

	// ResultOfEncodeAccount ...
	ResultOfEncodeAccount struct {
		Account string `json:"account"`
		ID      string `json:"id"`
	}

	// ParamsOfDecodeAccountData ...
	ParamsOfDecodeAccountData struct {
		Abi  *Abi   `json:"abi"`
		Data string `json:"data"`
	}

	// ResultOfDecodeData ...
	ResultOfDecodeData struct {
		Data json.RawMessage `json:"data"`
	}

	// ParamsOfUpdateInitialData ...
	ParamsOfUpdateInitialData struct {
		Abi           *Abi          `json:"abi,omitempty"`
		Data          string        `json:"data"`
		InitialData   interface{}   `json:"initial_data,omitempty"`
		InitialPubKey string        `json:"initial_pubkey,omitempty"`
		BocCache      *BocCacheType `json:"boc_cache,omitempty"`
	}

	// ResultOfUpdateInitialData ...
	ResultOfUpdateInitialData struct {
		Data string `json:"data"`
	}

	// ParamsOfEncodeInitialData ...
	ParamsOfEncodeInitialData struct {
		Abi           *Abi          `json:"abi,omitempty"`
		InitialData   interface{}   `json:"initial_data,omitempty"`
		InitialPubKey string        `json:"initial_pubkey,omitempty"`
		BocCache      *BocCacheType `json:"boc_cache,omitempty"`
	}

	// ResultOfEncodeInitialData ...
	ResultOfEncodeInitialData struct {
		Data string `json:"data"`
	}

	// ParamsOfDecodeInitialData ...
	ParamsOfDecodeInitialData struct {
		Abi  *Abi   `json:"abi,omitempty"`
		Data string `json:"data"`
	}

	// ResultOfDecodeInitialData ...
	ResultOfDecodeInitialData struct {
		InitialData   interface{} `json:"initial_data,omitempty"`
		InitialPubKey string      `json:"initial_pubkey,omitempty"`
	}

	// ParamsOfDecodeBoc ...
	ParamsOfDecodeBoc struct {
		Params       []*AbiParam `json:"params"`
		Boc          string      `json:"boc"`
		AllowPartial bool        `json:"allow_partial"`
	}

	// ResultOfDecodeBoc ...
	ResultOfDecodeBoc struct {
		Data json.RawMessage `json:"data"`
	}

	// ParamsOfAbiEncodeBoc ...
	ParamsOfAbiEncodeBoc struct {
		Params []*AbiParam `json:"params"`
		Data json.RawMessage `json:"data"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	// ResultOfAbiEncodeBoc ...
	ResultOfAbiEncodeBoc struct {
		Boc string `json:"boc"`
	}

	//AbiUseCase ...
	AbiUseCase interface {
		EncodeMessageBody(*ParamsOfEncodeMessageBody) (*ResultOfEncodeMessageBody, error)
		AttachSignatureToMessageBody(*ParamsOfAttachSignatureToMessageBody) (*ResultOfAttachSignatureToMessageBody, error)
		EncodeMessage(*ParamsOfEncodeMessage) (*ResultOfEncodeMessage, error)
		EncodeInternalMessage(*ParamsOfEncodeInternalMessage) (*ResultOfEncodeInternalMessage, error)
		AttachSignature(*ParamsOfAttachSignature) (*ResultOfAttachSignature, error)
		DecodeMessage(*ParamsOfDecodeMessage) (*DecodedMessageBody, error)
		DecodeMessageBody(*ParamsOfDecodeMessageBody) (*DecodedMessageBody, error)
		EncodeAccount(*ParamsOfEncodeAccount) (*ResultOfEncodeAccount, error)
		DecodeAccountData(*ParamsOfDecodeAccountData) (*ResultOfDecodeData, error)
		UpdateInitialData(*ParamsOfUpdateInitialData) (*ResultOfUpdateInitialData, error)
		EncodeInitialData(*ParamsOfEncodeInitialData) (*ResultOfEncodeInitialData, error)
		DecodeInitialData(*ParamsOfDecodeInitialData) (*ResultOfDecodeInitialData, error)
		DecodeBoc(*ParamsOfDecodeBoc) (*ResultOfDecodeBoc, error)
		EncodeBoc(*ParamsOfAbiEncodeBoc) (*ResultOfAbiEncodeBoc, error)
	}
)

func init() {
	AbiErrorCode = map[string]int{
		"RequiredAddressMissingForEncodeMessage":    301,
		"RequiredCallSetMissingForEncodeMessage":    302,
		"InvalidJson":                               303,
		"InvalidMessage":                            304,
		"EncodeDeployMessageFailed":                 305,
		"EncodeRunMessageFailed":                    306,
		"AttachSignatureFailed":                     307,
		"InvalidTvcImage":                           308,
		"RequiredPublicKeyMissingForFunctionHeader": 309,
		"InvalidSigner":                             310,
		"InvalidAbi":                                311,
		"InvalidFunctionId":                         312,
		"InvalidData":                               313,
		"EncodeInitialDataFailed":                   314,
	}
}

func NewAbiContract(value *AbiContract) *Abi {
	return &Abi{Type: ContractAbiType, Value: value}
}

func NewAbiJSON(value string) *Abi {
	return &Abi{Type: JSONAbiType, Value: value}
}

func NewAbiHandle(value *AbiHandle) *Abi {
	return &Abi{Type: HandleAbiType, Value: value}
}

func NewAbiSerialized(value *AbiContract) *Abi {
	return &Abi{Type: SerializedAbiType, Value: value}
}

func (s *Signer) MarshalJSON() ([]byte, error) {
	switch value := (s.ValueEnumType).(type) {
	case SignerNone:
		return json.Marshal(struct {
			Type string `json:"type"`
			SignerNone
		}{"None", value})
	case SignerExternal:
		return json.Marshal(struct {
			Type string `json:"type"`
			SignerExternal
		}{"External", value})
	case SignerKeys:
		return json.Marshal(struct {
			Type string `json:"type"`
			SignerKeys
		}{"Keys", value})
	case SignerSigningBox:
		return json.Marshal(struct {
			Type string `json:"type"`
			SignerSigningBox
		}{"SigningBox", value})
	default:
		return nil, fmt.Errorf("unsupported type for Signer %v", s.ValueEnumType)
	}
}

func (s *Signer) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "None":
		var valueEnum SignerNone
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	case "External":
		var valueEnum SignerExternal
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	case "Keys":
		var valueEnum SignerKeys
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	case "Serialized":
		var valueEnum SignerSigningBox
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for Signer %v", typeD.Type)
	}
	return nil
}

// NewSigner ...
func NewSigner(value interface{}) *Signer {
	return &Signer{ValueEnumType: value}
}

func (s *StateInitSource) MarshalJSON() ([]byte, error) {
	switch value := (s.ValueEnumType).(type) {
	case StateInitSourceMessage:
		return json.Marshal(struct {
			Type string `json:"type"`
			StateInitSourceMessage
		}{"Message", value})
	case StateInitSourceStateInit:
		return json.Marshal(struct {
			Type string `json:"type"`
			StateInitSourceStateInit
		}{"StateInit", value})
	case StateInitSourceTvc:
		return json.Marshal(struct {
			Type string `json:"type"`
			StateInitSourceTvc
		}{"Tvc", value})
	default:
		return nil, fmt.Errorf("unsupported type for StateInitSource %v", s.ValueEnumType)
	}
}

func (s *StateInitSource) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Message":
		var valueEnum StateInitSourceMessage
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	case "StateInit":
		var valueEnum StateInitSourceStateInit
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	case "Tvc":
		var valueEnum StateInitSourceTvc
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		s.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for StateInitSource %v", typeD.Type)
	}
	return nil
}

// NewStateInitSource ...
func NewStateInitSource(value interface{}) *StateInitSource {
	return &StateInitSource{ValueEnumType: value}
}

func (ms *MessageSource) MarshalJSON() ([]byte, error) {
	switch value := (ms.ValueEnumType).(type) {
	case MessageSourceEncoded:
		return json.Marshal(struct {
			Type string `json:"type"`
			MessageSourceEncoded
		}{"Encoded", value})
	case MessageSourceEncodingParams:
		return json.Marshal(struct {
			Type string `json:"type"`
			MessageSourceEncodingParams
		}{"EncodingParams", value})
	default:
		return nil, fmt.Errorf("unsupported type for MessageSource %v", ms.ValueEnumType)
	}
}

func (ms *MessageSource) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Encoded":
		var valueEnum MessageSourceEncoded
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ms.ValueEnumType = valueEnum
	case "EncodingParams":
		var valueEnum MessageSourceEncodingParams
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ms.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for MessageSource %v", typeD.Type)
	}
	return nil
}

// NewStateInitSource ...
func NewMessageSource(value interface{}) *MessageSource {
	return &MessageSource{ValueEnumType: value}
}
