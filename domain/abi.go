package domain

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const (
	ContractAbiType   AbiType = "Contract"
	JSONAbiType       AbiType = "Json"
	HandleAbiType     AbiType = "Handle"
	SerializedAbiType AbiType = "Serialized"

	InputDataLayout  DataLayoutType = "Input"
	OutputDataLayout DataLayoutType = "Output"

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

var AbiErrorCode map[string]int

type (
	AbiHandle      int
	AbiType        string
	DataLayoutType string

	AbiContractVariant struct {
		Value AbiContract `json:"value"`
	}

	AbiJsonVariant struct {
		Value string `json:"value"`
	}

	AbiHandleVariant struct {
		Value AbiHandle `json:"value"`
	}

	AbiSerializedVariant struct {
		Value AbiContract `json:"value"`
	}

	Abi struct {
		Type  AbiType     `json:"type"`
		Value interface{} `json:"value"`
	}

	// FunctionHeader - The ABI function header.
	// Includes several hidden function parameters that contract uses for security, message delivery monitoring and
	// replay protection reasons.
	//The actual set of header fields depends on the contract's ABI. If a contract's ABI does not include some
	// headers, then they are not filled.
	FunctionHeader struct {
		Expire *int     `json:"expire,omitempty"`
		Time   *big.Int `json:"time,omitempty"`
		PubKey string   `json:"pubkey,omitempty"`
	}

	CallSet struct {
		FunctionName string          `json:"function_name"`
		Header       *FunctionHeader `json:"header,omitempty"`
		Input        interface{}     `json:"input,omitempty"`
	}

	DeploySet struct {
		Tvc           string      `json:"tvc,omitempty"`
		Code          string      `json:"code,omitempty"`
		StateInit     string      `json:"state_init,omitempty"`
		WorkchainID   *int        `json:"workchain_id,omitempty"`
		InitialData   interface{} `json:"initial_data,omitempty"`
		InitialPubKey string      `json:"initial_pubkey,omitempty"`
	}

	// SignerNone - No keys are provided. Creates an unsigned message.
	SignerNone struct{}

	// SignerExternal - Only public key is provided to generate unsigned message and data_to_sign which can be signed later.
	SignerExternal struct {
		PublicKey string `json:"public_key"`
	}

	// SignerKeys - Key pair is provided for signing
	SignerKeys struct {
		Keys *KeyPair `json:"keys"`
	}

	// SignerSigningBox - Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
	SignerSigningBox struct {
		Handle SigningBoxHandle `json:"handle"`
	}

	Signer struct {
		ValueEnumType interface{}
	}

	MessageBodyType string

	AbiParam struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
		Init bool `json:"init,omitempty`
	}

	AbiEvent struct {
		Name   string      `json:"name"`
		Inputs []*AbiParam `json:"inputs"`
		ID     string      `json:"id,omitempty"`
	}

	AbiData struct {
		Key        int         `json:"key"`
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	AbiFunctions struct {
		Name    string      `json:"name"`
		Inputs  []*AbiParam `json:"inputs"`
		Outputs []*AbiParam `json:"outputs"`
		ID      string      `json:"id,omitempty"`
	}

	AbiContract struct {
		Abi_Version *int            `json:"ABI version,omitempty"`
		AbiVersion  *int            `json:"abi_version,omitempty"`
		Version     string          `json:"version,omitempty"`
		Header      []string        `json:"header,omitempty"`
		Functions   []*AbiFunctions `json:"functions,omitempty"`
		Events      []*AbiEvent     `json:"events,omitempty"`
		Data        []*AbiData      `json:"data,omitempty"`
		Fields      []*AbiParam     `json:"fields,omitempty"`
	}

	ParamsOfEncodeMessageBody struct {
		Abi                *Abi     `json:"abi"`
		CallSet            *CallSet `json:"call_set"`
		IsInternal         bool     `json:"is_internal"`
		Signer             *Signer  `json:"signer"`
		ProcessingTryIndex *int     `json:"processing_try_index,omitempty"`
		Address            string   `json:"address,omitempty"`
		SignatureID        *int     `json:"signature_id,omitempty"`
	}

	ResultOfEncodeMessageBody struct {
		Body       string `json:"body"`
		DataToSign string `json:"data_to_sign,omitempty"`
	}

	ParamsOfAttachSignatureToMessageBody struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	ResultOfAttachSignatureToMessageBody struct {
		Body string `json:"body"`
	}

	ParamsOfEncodeMessage struct {
		Abi                *Abi       `json:"abi"`
		Address            string     `json:"address,omitempty"`
		DeploySet          *DeploySet `json:"deploy_set,omitempty"`
		CallSet            *CallSet   `json:"call_set,omitempty"`
		Signer             *Signer    `json:"signer"`
		ProcessingTryIndex *int       `json:"processing_try_index"`
		SignatureID        *int       `json:"signature_id,omitempty"`
	}

	ResultOfEncodeMessage struct {
		Message    string `json:"message"`
		DataToSign string `json:"data_to_sign,omitempty"`
		Address    string `json:"address"`
		MessageID  string `json:"message_id"`
	}

	ParamsOfEncodeInternalMessage struct {
		Abi        *Abi       `json:"abi,omitempty"`
		Address    string     `json:"address,omitempty"`
		SrcAddress string     `json:"src_address,omitempty"`
		DeploySet  *DeploySet `json:"deploy_set,omitempty"`
		CallSet    *CallSet   `json:"call_set,omitempty"`
		Value      string     `json:"value"`
		Bounce     *bool      `json:"bounce,omitempty"`
		EnableIhr  *bool      `json:"enable_ihr,omitempty"`
	}

	ResultOfEncodeInternalMessage struct {
		Message   string `json:"message"`
		Address   string `json:"address"`
		MessageID string `json:"message_id"`
	}

	ParamsOfAttachSignature struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	ResultOfAttachSignature struct {
		Message   string `json:"message"`
		MessageID string `json:"message_id"`
	}

	ParamsOfDecodeMessage struct {
		Abi          *Abi            `json:"abi"`
		Message      string          `json:"message"`
		AllowPartial *bool           `json:"allow_partial,omitempty"`
		FunctionName string          `json:"function_name,omitempty"`
		DataLayout   *DataLayoutType `json:"data_layout,omitempty"`
	}

	DecodedMessageBody struct {
		BodyType MessageBodyType `json:"body_type"`
		Name     string          `json:"name"`
		Value    json.RawMessage `json:"value,omitempty"`
		Header   *FunctionHeader `json:"header,omitempty"`
	}

	ParamsOfDecodeMessageBody struct {
		Abi          *Abi            `json:"abi"`
		Body         string          `json:"body"`
		IsInternal   bool            `json:"is_internal"`
		AllowPartial *bool           `json:"allow_partial,omitempty"`
		FunctionName string          `json:"function_name,omitempty"`
		DataLayout   *DataLayoutType `json:"data_layout,omitempty"`
	}

	ParamsOfEncodeAccount struct {
		StateInit   string `json:"state_init"`
		Balance     *big.Int         `json:"balance,omitempty"`
		LastTransLt *big.Int         `json:"last_trans_lt,omitempty"`
		LastPaid    *int             `json:"last_paid,omitempty"`
		BocCache    *BocCacheType    `json:"boc_cache,omitempty"`
	}

	ResultOfEncodeAccount struct {
		Account string `json:"account"`
		ID      string `json:"id"`
	}

	ParamsOfDecodeAccountData struct {
		Abi          *Abi   `json:"abi"`
		Data         string `json:"data"`
		AllowPartial *bool  `json:"allow_partial,omitempty"`
	}

	ResultOfDecodeData struct {
		Data json.RawMessage `json:"data"`
	}

	ParamsOfUpdateInitialData struct {
		Abi           *Abi          `json:"abi"`
		Data          string        `json:"data"`
		InitialData   interface{}   `json:"initial_data,omitempty"`
		InitialPubKey string        `json:"initial_pubkey,omitempty"`
		BocCache      *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfUpdateInitialData struct {
		Data string `json:"data"`
	}

	ParamsOfEncodeInitialData struct {
		Abi           *Abi          `json:"abi"`
		InitialData   interface{}   `json:"initial_data,omitempty"`
		InitialPubKey string        `json:"initial_pubkey,omitempty"`
		BocCache      *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfEncodeInitialData struct {
		Data string `json:"data"`
	}

	ParamsOfDecodeInitialData struct {
		Abi          *Abi   `json:"abi"`
		Data         string `json:"data"`
		AllowPartial *bool  `json:"allow_partial,omitempty"`
	}

	ResultOfDecodeInitialData struct {
		InitialData   interface{} `json:"initial_data"`
		InitialPubKey string      `json:"initial_pubkey,omitempty"`
	}

	ParamsOfDecodeBoc struct {
		Params       []*AbiParam `json:"params"`
		Boc          string      `json:"boc"`
		AllowPartial bool        `json:"allow_partial"`
	}

	ResultOfDecodeBoc struct {
		Data json.RawMessage `json:"data"`
	}

	ParamsOfAbiEncodeBoc struct {
		Params   []*AbiParam     `json:"params"`
		Data     json.RawMessage `json:"data"`
		BocCache *BocCacheType   `json:"boc_cache,omitempty"`
	}

	ResultOfAbiEncodeBoc struct {
		Boc string `json:"boc"`
	}

	ParamsOfCalcFunctionId struct {
		Abi          Abi    `json:"abi"`
		FunctionName string `json:"function_name"`
		Output       *bool  `json:"output,omitempty"`
	}

	ResultOfCalcFunctionId struct {
		FunctionID int `json:"function_id"`
	}

	ParamsOfGetSignatureData struct {
		Abi         Abi    `json:"abi"`
		Message     string `json:"message"`
		SignatureID *int   `json:"signature_id,omitempty"`
	}

	ResultOfGetSignatureData struct {
		Signature string `json:"signature"`
		Unsigned  string `json:"unsigned"`
	}

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
		CalcFunctionID(*ParamsOfCalcFunctionId) (*ResultOfCalcFunctionId, error)
		GetSignatureData(*ParamsOfGetSignatureData) (*ResultOfGetSignatureData, error)
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
		"InvalidFunctionName":                       315,
		"PubKeyNotSupported":                        316,
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

func NewMessageSource(value interface{}) *MessageSource {
	return &MessageSource{ValueEnumType: value}
}
