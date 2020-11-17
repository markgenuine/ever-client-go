package domain

const (
	// SignNone ...
	SignNone SignerType = "None"

	// SignExternal ...
	SignExternal SignerType = "External"

	// SignKeys ...
	SignKeys SignerType = "Keys"

	// SignSigningBox ...
	SignSigningBox SignerType = "SigningBox"

	// TypeS ...
	TypeS AbiType = "Serialized"

	// TypeH ...
	TypeH AbiType = "Handle"

	// MessageBodyTypeInput ...
	MessageBodyTypeInput MessageBodyType = "Input"

	// MessageBodyTypeOutput ...
	MessageBodyTypeOutput MessageBodyType = "Output"

	// MessageBodyTypeInternalOutput ...
	MessageBodyTypeInternalOutput MessageBodyType = "InternalOutput"

	// MessageBodyTypeEvent ...
	MessageBodyTypeEvent MessageBodyType = "Event"

	// StateInitSourceTypeMessage ...
	StateInitSourceTypeMessage StateInitSourceType = "Message"

	// StateInitSourceTypeStateInit ...
	StateInitSourceTypeStateInit StateInitSourceType = "StateInit"

	// StateInitSourceTypeTvc ...
	StateInitSourceTypeTvc StateInitSourceType = "Tvc"

	// MessageSourceEncoded ...
	MessageSourceEncoded MessageSourceType = "Encoded"

	// MessageSourceEncodingParams ...
	MessageSourceEncodingParams MessageSourceType = "EncodingParams"
)

type (
	// AbiHandler ...
	AbiHandler int

	// SignerType ...
	SignerType string

	// AbiType ...
	AbiType string

	// MessageBodyType ...
	MessageBodyType string

	// StateInitSourceType ...
	StateInitSourceType string

	// MessageSourceType ...
	MessageSourceType string

	// FunctionHeader ...
	FunctionHeader struct {
		Expire int    `json:"expire"`
		Time   string `json:"time"`
		PubKey string `json:"pub_key"`
	}

	// CallSet ...
	CallSet struct {
		FunctionName string          `json:"function_name"`
		Header       *FunctionHeader `json:"header,omitempty"`
		Input        interface{}     `json:"input,omitempty"`
	}

	// DeploySet ...
	DeploySet struct {
		Tvc         string      `json:"tvc"`
		WorkchainID int         `json:"workchain_id"`
		InitialData interface{} `json:"initial_data,omitempty"`
	}

	// SignerNone No keys are provided. Creates an unsigned message.
	SignerNone struct {
		Type SignerType `json:"type"`
	}

	// SignerExternal Only public key is provided to generate unsigned message and data_to_sign which can be signed later.
	SignerExternal struct {
		Type      SignerType `json:"type"`
		PublicKey string     `json:"public_key"`
	}

	// SignerKeys Key pair is provided for signing
	SignerKeys struct {
		Type SignerType `json:"type"`
		Keys KeyPair    `json:"keys"`
	}

	// SignerSigningBox Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
	SignerSigningBox struct {
		Type   SignerType       `json:"type"`
		Handle SigningBoxHandle `json:"handle"`
	}

	// AbiS ...
	AbiS struct {
		Type  AbiType     `json:"type"`
		Value AbiContract `json:"value"`
	}

	// AbiFunctions ...
	AbiFunctions struct {
		Name    string     `json:"name"`
		Inputs  []AbiParam `json:"inputs"`
		Outputs []AbiParam `json:"outputs"`
		ID      int        `json:"id,omitempty"`
	}

	// AbiEvent ...
	AbiEvent struct {
		Name   string     `json:"name"`
		Inputs []AbiParam `json:"inputs"`
		ID     int        `json:"id,omitempty"`
	}

	// AbiData ...
	AbiData struct {
		Key        int         `json:"key"` //bigint????
		Name       string      `json:"name"`
		ParamType  string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiParam ...
	AbiParam struct {
		Name       string      `json:"name"`
		ParamType  string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiContract ...
	AbiContract struct {
		AbiVersion int            `json:"ABI version,omitempty"`
		Header     []string       `json:"header,omitempty"`
		Functions  []AbiFunctions `json:"functions"`
		Events     []*AbiEvent    `json:"events,omitempty"`
		Data       []*AbiData     `json:"data,omitempty"`
	}

	// AbiH ...
	AbiH struct {
		Type  AbiType `json:"type"`
		Value int     `json:"value"`
	}

	// StateInitSourceM Deploy message.
	StateInitSourceM struct {
		Type   StateInitSourceType `json:"type"`
		Source interface{}         `json:"source"`
	}

	// StateInitSourceSI State init data
	StateInitSourceSI struct {
		Type    StateInitSourceType `json:"type"`
		Code    string              `json:"code"`
		Data    string              `json:"data"`
		Library string              `json:"library,omitempty"`
	}

	// StateInitSourceT Content of the TVC file. Encoded in base64.
	StateInitSourceT struct {
		Type       StateInitSourceType `json:"type"`
		Tvc        string              `json:"tvc"`
		PublicKey  string              `json:"public_key,omitempty"`
		InitParams StateInitParams     `json:"init_params,omitempty"`
	}
	// MessageSourceE ...
	MessageSourceE struct {
		Type    MessageSourceType `json:"type"`
		Message string            `json:"message"`
		Abi     interface{}       `json:"abi,omitempty"` //Abi AbiS and AbiH
	}

	// MessageSourceEP ...
	MessageSourceEP struct {
		Type               MessageSourceType `json:"type"`
		Abi                interface{}       `json:"abi,omitempty"` //Abi AbiS and AbiH
		Address            string            `json:"address,omitempty"`
		DeploySet          *DeploySet        `json:"deploy_set,omitempty"`
		CallSet            *CallSet          `json:"call_set,omitempty"`
		Signer             interface{}       `json:"signer"` //SignerNone, SignerExternal .....?????
		ProcessingTryIndex int               `json:"processing_try_index"`
	}

	// StateInitParams ...
	StateInitParams struct {
		Abi   interface{} `json:"abi"` //ABI??? AbiS and AbiH
		Value interface{} `json:"value"`
	}

	// ParamsOfEncodeMessageBody ...
	ParamsOfEncodeMessageBody struct {
		Abi                interface{} `json:"abi"` //ABI??? AbiS and AbiH
		CallSet            *CallSet    `json:"call_set"`
		IsInternal         bool        `json:"is_internal"`
		Signer             interface{} `json:"signer"` //SignerNone, SignerExternal .....?????
		ProcessingTryIndex int         `json:"processing_try_index"`
	}

	// ResultOfEncodeMessageBody ...
	ResultOfEncodeMessageBody struct {
		Body       string `json:"body"`
		DataToSign string `json:"data_to_sign,omitempty"`
	}

	// ParamsOfAttachSignatureToMessageBody ...
	ParamsOfAttachSignatureToMessageBody struct {
		Abi       interface{} `json:"abi"` //ABI??? AbiS and AbiH
		PublicKey string      `json:"public_key"`
		Message   string      `json:"message"`
		Signature string      `json:"signature"`
	}

	// ResultOfAttachSignatureToMessageBody ...
	ResultOfAttachSignatureToMessageBody struct {
		Body string `json:"body"`
	}

	// ParamsOfEncodeMessage ...
	ParamsOfEncodeMessage struct {
		Abi                interface{} `json:"abi"` //ABI??? AbiS and AbiH
		Address            string      `json:"address,omitempty"`
		DeploySet          *DeploySet  `json:"deploy_set,omitempty"`
		CallSet            *CallSet    `json:"call_set,omitempty"`
		Signer             interface{} `json:"signer"` //SignerNone, SignerExternal .....?????
		ProcessingTryIndex int         `json:"processing_try_index"`
	}

	// ResultOfEncodeMessage ...
	ResultOfEncodeMessage struct {
		Message    string `json:"message"`
		DataToSign string `json:"data_to_sign,omitempty"`
		Address    string `json:"address"`
		MessageID  string `json:"message_id"`
	}

	// ParamsOfAttachSignature ...
	ParamsOfAttachSignature struct {
		Abi       interface{} `json:"abi"` //ABI??? AbiS and AbiH
		PublicKey string      `json:"public_key"`
		Message   string      `json:"message"`
		Signature string      `json:"signature"`
	}

	// ResultOfAttachSignature ...
	ResultOfAttachSignature struct {
		Message   string `json:"message"`
		MessageID string `json:"message_id"`
	}

	// ParamsOfDecodeMessage ...
	ParamsOfDecodeMessage struct {
		Abi     interface{} `json:"abi"` //ABI??? AbiS and AbiH
		Message string      `json:"message"`
	}

	// DecodedMessageBody ...
	DecodedMessageBody struct {
		BodyType MessageBodyType `json:"body_type"`
		Name     string          `json:"name"`
		Value    interface{}     `json:"value,omitempty"`
		Header   *FunctionHeader `json:"header,omitempty"`
	}

	// ParamsOfDecodeMessageBody ...
	ParamsOfDecodeMessageBody struct {
		Abi        interface{} `json:"abi"` //ABI??? AbiS and AbiH
		Body       string      `json:"body"`
		IsInternal bool        `json:"is_internal"`
	}

	// ParamsOfEncodeAccount ...
	ParamsOfEncodeAccount struct {
		StateInit   interface{} `json:"state_init"` //StateInitSource
		Balance     string      `json:"balance,omitempty"`
		LastTransLt string      `json:"last_trans_lt,omitempty"`
		LastPaid    int         `json:"last_paid,omitempty"`
	}

	// ResultOfEncodeAccount ...
	ResultOfEncodeAccount struct {
		Account string `json:"account"`
		ID      string `json:"id"`
	}

	//AbiUseCase ...
	AbiUseCase interface {
		EncodeMessageBody(pOEMB ParamsOfEncodeMessageBody) (int, error)
		AttachSignatureToMessageBody(pOASTMB ParamsOfAttachSignatureToMessageBody) (int, error)
		EncodeMessage(pOEM ParamsOfEncodeMessage) (int, error)
		AttachSignature(pOAS ParamsOfAttachSignature) (int, error)
		DecodeMessage(pODM ParamsOfDecodeMessage) (int, error)
		DecodeMessageBody(pODMB ParamsOfDecodeMessageBody) (int, error)
		EncodeAccount(pOEA ParamsOfEncodeAccount) (int, error)
	}
)
