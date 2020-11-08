package goton

import (
	"encoding/json"
	"errors"
	"sync"
)

var (
	Chains        = getChains()
	LensMnemonic  = getlengthCounWordsInMnemonic()
	SortDirection = getSortDirection()
	// TONMnemonicDictionary map with dictionary
	TONMnemonicDictionary = getTONMnemonicDictionary()
	State                 = &state{&sync.Mutex{}, map[int]*AsyncResponse{}}
)

// ClientError ...
type ClientError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// TomlConfig struct with config data
type TomlConfig struct {
	Network Network `toml:"network" json:"network,omitempty"`
	Crypto  `toml:"crypto" json:"crypto,omitempty"`
	Abi     `toml:"abi" json:"abi,omitempty"`
}

// Network ...
type Network struct {
	ServerAddress            string `toml:"server_address" json:"server_address,omitempty"`
	NetworkRetriesCount      int    `toml:"network_retries_count" json:"network_retries_count,omitempty"`
	MessageRetriesCount      int    `toml:"message_retries_count" json:"message_retries_count,omitempty"`
	MessageProcessingTimeout int    `toml:"message_processing_timeout" json:"message_processing_timeout,omitempty"`
	WaitForTimeout           int    `toml:"wait_for_timeout" json:"wait_for_timeout,omitempty"`
	OutOfSyncThreshold       int    `toml:"out_of_sync_threshold" json:"out_of_sync_threshold,omitempty"`
	AccessKey                string `toml:"access_key" json:"access_key,omitempty"`
}

// Crypto ...
type Crypto struct {
	MnemonicDictionary  int    `toml:"mnemonic_dictionary" json:"mnemonic_dictionary,omitempty"`
	MnemonicWordCount   int    `toml:"mnemonic_word_count" json:"mnemonic_word_count,omitempty"`
	HdkeyDerivationPath string `toml:"hdkey_derivation_path" json:"hdkey_derivation_path,omitempty"`
	HdkeyCompliant      bool   `toml:"hdkey_compliant" json:"hdkey_compliant,omitempty"`
}

// Abi ...
type Abi struct {
	Workchain                          int     `toml:"workchain" json:"workchain,omitempty"`
	MessageExpirationTimeout           int     `toml:"message_expiration_timeout" json:"message_expiration_timeout"`
	MessageExpirationTimeoutGrowFactor float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor"`
}

// NewConfig create new config for connect client
// chanID 0-devnet, 1-mainnet,
func NewConfig(chanID int) *TomlConfig {
	config := TomlConfig{
		Network: Network{
			ServerAddress:            Chains[chanID],
			NetworkRetriesCount:      0,
			MessageRetriesCount:      10,
			MessageProcessingTimeout: 40000, //ms
			WaitForTimeout:           40000, //ms
			OutOfSyncThreshold:       15000, //ms
			AccessKey:                "",
		},
		Crypto: Crypto{
			MnemonicDictionary:  1,
			MnemonicWordCount:   12,
			HdkeyDerivationPath: "",
			HdkeyCompliant:      false,
		},
		Abi: Abi{
			MessageExpirationTimeout:           40000, //ms
			MessageExpirationTimeoutGrowFactor: 1.5},
	}

	return &config
}

func getChains() map[int]string {
	return map[int]string{0: "net.ton.dev", 1: "main.ton.dev"}
}

func getlengthCounWordsInMnemonic() map[int]int {
	return map[int]int{12: 12, 15: 15, 18: 18, 21: 21, 24: 24}
}

func getTONMnemonicDictionary() map[string]int {
	return map[string]int{
		"TON":                 0,
		"ENGLISH":             1,
		"CHINESE_SIMPLIFIED":  2,
		"CHINESE_TRADITIONAL": 3,
		"FRENCH":              4,
		"ITALIAN":             5,
		"JAPANESE":            6,
		"KOREAN":              7,
		"SPANISH":             8,
	}
}

func getSortDirection() map[int]string {
	return map[int]string{
		0: "ASC",
		1: "DESC",
	}
}

// AsyncResponse ...
type AsyncResponse struct {
	ReqID        int
	MethodName   string
	Params       string
	ResponseType int
	Finished     bool
}

type state struct {
	*sync.Mutex
	stores map[int]*AsyncResponse
}

//Client

// ResultOfVersion ...
type ResultOfVersion struct {
	Version string `json:"version"`
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

// ResultOfBuildInfo ...
type ResultOfBuildInfo struct {
	BuildInfo interface{} `json:"build_info"`
}

//  Abi

// AbiHandler ...
type AbiHandler int

// SignerType ...
type SignerType string

// AbiType ...
type AbiType string

// MessageBodyType ...
type MessageBodyType string

// StateInitSourceType ...
type StateInitSourceType string

// MessageSourceType ...
type MessageSourceType string

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

	MessageBodyTypeInput          MessageBodyType = "Input"
	MessageBodyTypeOutput         MessageBodyType = "Output"
	MessageBodyTypeInternalOutput MessageBodyType = "InternalOutput"
	MessageBodyTypeEvent          MessageBodyType = "Event"

	StateInitSourceTypeMessage   StateInitSourceType = "Message"
	StateInitSourceTypeStateInit StateInitSourceType = "StateInit"
	StateInitSourceTypeTvc       StateInitSourceType = "Tvc"

	MessageSourceEncoded        MessageSourceType = "Encoded"
	MessageSourceEncodingParams MessageSourceType = "EncodingParams"
)

// FunctionHeader ...
type FunctionHeader struct {
	Expire int    `json:"expire"`
	Time   string `json:"time"`    //bigint???
	PubKey string `json:"pub_key"` //hex
}

// CallSet ...
type CallSet struct {
	FunctionName string          `json:"function_name"`
	Header       *FunctionHeader `json:"header,omitempty"`
	Input        interface{}     `json:"input,omitempty"`
}

// DeploySet ...
type DeploySet struct {
	Tvc         string      `json:"tvc"`
	WorkchainID int         `json:"workchain_id,omitempty"` //default 0
	InitialData interface{} `json:"initial_data,omitempty"`
}

// SignerNone No keys are provided. Creates an unsigned message.
type SignerNone struct {
	Type SignerType `json:"type"`
}

// SignerExternal Only public key is provided to generate unsigned message and data_to_sign which can be signed later.
type SignerExternal struct {
	Type      SignerType `json:"type"`
	PublicKey string     `json:"public_key"`
}

// SignerKeys Key pair is provided for signing
type SignerKeys struct {
	Type SignerType `json:"type"`
	Keys KeyPair    `json:"keys"`
}

// SignerSigningBox Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
type SignerSigningBox struct {
	Type   SignerType       `json:"type"`
	Handle SigningBoxHandle `json:"handle"`
}

// AbiS ...
type AbiS struct {
	Type  AbiType     `json:"type"`
	Value AbiContract `json:"value"`
}

// AbiFunctions ...
type AbiFunctions struct {
	Name    string     `json:"name"`
	Inputs  []AbiParam `json:"inputs"`
	Outputs []AbiParam `json:"outputs"`
	ID      int        `json:"id,omitempty"`
}

// AbiEvent ...
type AbiEvent struct {
	Name   string     `json:"name"`
	Inputs []AbiParam `json:"inputs"`
	ID     int        `json:"id,omitempty"`
}

// AbiData ...
type AbiData struct {
	Key        int        `json:"key"` //bigint????
	Name       string     `json:"name"`
	ParamType  string     `json:"type"`
	Components []AbiParam `json:"components,omitempty"`
}

// AbiParam ...
type AbiParam struct {
	Name       string     `json:"name"`
	ParamType  string     `json:"type"`
	Components []AbiParam `json:"components,omitempty"`
}

// AbiContract ...
type AbiContract struct {
	AbiVersion int            `json:"ABI version,omitempty"`
	Header     []string       `json:"header,omitempty"`
	Functions  []AbiFunctions `json:"functions"` //?: AbiFunction[],
	Events     []AbiEvent     `json:"events,omitempty"`
	Data       []AbiData      `json:"data,omitempty"`
}

// AbiH ...
type AbiH struct {
	Type  AbiType `json:"type"`
	Value int     `json:"value"`
}

// IsValid ...
func (at AbiType) IsValid() error {
	switch at {
	case TypeS, TypeH:
		return nil
	}
	return errors.New("Invalid abi value")
}

// IsValid ...
func (mbt MessageBodyType) IsValid() error {
	switch mbt {
	case MessageBodyTypeInput, MessageBodyTypeOutput, MessageBodyTypeInternalOutput, MessageBodyTypeEvent:
		return nil
	}
	return errors.New("Invalid message body value")
}

// IsValid ...
func (sT SignerType) IsValid() error {
	switch sT {
	case SignNone, SignExternal, SignKeys, SignSigningBox:
		return nil
	}
	return errors.New("Invalid Signer body value")
}

// StateInitSourceM Deploy message.
type StateInitSourceM struct {
	Type   StateInitSourceType `json:"type"`
	Source interface{}         `json:"source"`
}

// StateInitSourceSI State init data
type StateInitSourceSI struct {
	Type    StateInitSourceType `json:"type"`
	Code    string              `json:"code"`
	Data    string              `json:"data"`
	Library string              `json:"library,omitempty"`
}

// StateInitSourceT Content of the TVC file. Encoded in base64.
type StateInitSourceT struct {
	Type       StateInitSourceType `json:"type"`
	Tvc        string              `json:"tvc"`
	PublicKey  string              `json:"public_key,omitempty"`
	InitParams StateInitParams     `json:"init_params,omitempty"`
}

// IsValid ...
func (sist StateInitSourceType) IsValid() error {
	switch sist {
	case StateInitSourceTypeMessage, StateInitSourceTypeStateInit, StateInitSourceTypeTvc:
		return nil
	}
	return errors.New("Invalid message body value")
}

// MessageSourceE ...
type MessageSourceE struct {
	Type    MessageSourceType `json:"type"`
	Message string            `json:"message"`
	Abi     interface{}       `json:"abi,omitempty"` //Abi AbiS and AbiH
}

// MessageSourceEP ...
type MessageSourceEP struct {
	Type               MessageSourceType `json:"type"`
	Abi                interface{}       `json:"abi,omitempty"` //Abi AbiS and AbiH
	Address            string            `json:"address,omitempty"`
	DeploySet          *DeploySet        `json:"deploy_set,omitempty"`
	CallSet            *CallSet          `json:"call_set,omitempty"`
	Signer             interface{}       `json:"signer"`                         //SignerNone, SignerExternal .....?????
	ProcessingTryIndex int               `json:"processing_try_index,omitempty"` //Default value is 0.
}

// StateInitParams ...
type StateInitParams struct {
	Abi   interface{} `json:"abi"` //ABI??? AbiS and AbiH
	Value interface{} `json:"value"`
}

// ParamsOfEncodeMessageBody ...
type ParamsOfEncodeMessageBody struct {
	Abi                interface{} `json:"abi"` //ABI??? AbiS and AbiH
	CallSet            *CallSet    `json:"call_set"`
	IsInternal         bool        `json:"is_internal"`
	Signer             interface{} `json:"signer"`                         //SignerNone, SignerExternal .....?????
	ProcessingTryIndex int         `json:"processing_try_index,omitempty"` //Default value is 0.
}

// ResultOfEncodeMessageBody ...
type ResultOfEncodeMessageBody struct {
	Body       string `json:"body"`
	DataToSign string `json:"data_to_sign,omitempty"`
}

// ParamsOfAttachSignatureToMessageBody ...
type ParamsOfAttachSignatureToMessageBody struct {
	Abi       interface{} `json:"abi"` //ABI??? AbiS and AbiH
	PublicKey string      `json:"public_key"`
	Message   string      `json:"message"`
	Signature string      `json:"signature"`
}

// ResultOfAttachSignatureToMessageBody ...
type ResultOfAttachSignatureToMessageBody struct {
	Body string `json:"body"`
}

// ParamsOfEncodeMessage ...
type ParamsOfEncodeMessage struct {
	Abi                interface{} `json:"abi"` //ABI??? AbiS and AbiH
	Address            string      `json:"address,omitempty"`
	DeploySet          *DeploySet  `json:"deploy_set,omitempty"`
	CallSet            *CallSet    `json:"call_set,omitempty"`
	Signer             interface{} `json:"signer"`               //SignerNone, SignerExternal .....?????
	ProcessingTryIndex int         `json:"processing_try_index"` //Default value is 0.
}

// ResultOfEncodeMessage ...
type ResultOfEncodeMessage struct {
	Message    string `json:"message"`
	DataToSign string `json:"data_to_sign, omitempty"`
	Address    string `json:"address"`
	MessageID  string `json:"message_id"`
}

// ParamsOfAttachSignature ...
type ParamsOfAttachSignature struct {
	Abi       interface{} `json:"abi"` //ABI??? AbiS and AbiH
	PublicKey string      `json:"public_key"`
	Message   string      `json:"message"`
	Signature string      `json:"signature"`
}

// ResultOfAttachSignature ...
type ResultOfAttachSignature struct {
	Message   string `json:"message"`
	MessageID string `json:"message_id"`
}

// ParamsOfDecodeMessage ...
type ParamsOfDecodeMessage struct {
	Abi     interface{} `json:"abi"` //ABI??? AbiS and AbiH
	Message string      `json:"message"`
}

// DecodedMessageBody ...
type DecodedMessageBody struct {
	BodyType MessageBodyType `json:"body_type"`
	Name     string          `json:"name"`
	Value    interface{}     `json:"value,omitempty"`
	Header   *FunctionHeader `json:"header,omitempty"`
}

// ParamsOfDecodeMessageBody ...
type ParamsOfDecodeMessageBody struct {
	Abi        interface{} `json:"abi"` //ABI??? AbiS and AbiH
	Body       string      `json:"body"`
	IsInternal bool        `json:"is_internal"`
}

// ParamsOfEncodeAccount ...
type ParamsOfEncodeAccount struct {
	StateInit   interface{} `json:"state_init"`              //StateInitSource
	Balance     string      `json:"balance,omitempty"`       //bigint?
	LastTransLt string      `json:"last_trans_lt,omitempty"` //bigint?
	LastPaid    int         `json:"last_paid,omitempty"`
}

// ResultOfEncodeAccount ...
type ResultOfEncodeAccount struct {
	Account string `json:"account"`
	ID      string `json:"id"`
}

// Boc

// ParamsOfParse ...
type ParamsOfParse struct {
	Boc string `json:"boc"`
}

// ResultOfParse ...
type ResultOfParse struct {
	Parsed json.RawMessage `json:"parsed"`
}

// ParamsOfGetBlockchainConfig ...
type ParamsOfGetBlockchainConfig struct {
	BlockBoc string `json:"block_boc"`
}

// ResultOfGetBlockchainConfig ...
type ResultOfGetBlockchainConfig struct {
	ConfigBoc string `json:"config_boc"`
}

//Crypto

// SigningBoxHandle ...
type SigningBoxHandle int

// ParamsOfFactorize ...
type ParamsOfFactorize struct {
	Composite string `json:"composite"`
}

// ResultOfFactorize ...
type ResultOfFactorize struct {
	Factors []string `json:"factors"`
}

// ParamsOfModularPower ...
type ParamsOfModularPower struct {
	Base     string `json:"base"`
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

// ResultOfModularPower ...
type ResultOfModularPower struct {
	ModularPower string `json:"modular_power"`
}

// ParamsOfTonCrc16 ...
type ParamsOfTonCrc16 struct {
	Data string `json:"data"`
}

// ResultOfTonCrc16 ...
type ResultOfTonCrc16 struct {
	Crc int `json:"crc"`
}

// ParamsOfGenerateRandomBytes ...
type ParamsOfGenerateRandomBytes struct {
	Length int `json:"length"`
}

// ResultOfGenerateRandomBytes ...
type ResultOfGenerateRandomBytes struct {
	Bytes string `json:"bytes"`
}

// ParamsOfConvertPublicKeyToTonSafeFormat ...
type ParamsOfConvertPublicKeyToTonSafeFormat struct {
	PublicKey string `json:"public_key"`
}

// ResultOfConvertPublicKeyToTonSafeFormat ...
type ResultOfConvertPublicKeyToTonSafeFormat struct {
	TonPublicKey string `json:"ton_public_key"`
}

// KeyPair ...
type KeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

// ParamsOfSign ...
type ParamsOfSign struct {
	Unsigned string   `json:"unsigned"`
	Keys     *KeyPair `json:"keys"`
}

// ResultOfSign ...
type ResultOfSign struct {
	Signed    string `json:"signed"`
	Signature string `json:"signature"`
}

// ParamsOfVerifySignature ...
type ParamsOfVerifySignature struct {
	Signed string `json:"signed"`
	Public string `json:"public"`
}

// ResultOfVerifySignature ...
type ResultOfVerifySignature struct {
	Unsigned string `json:"unsigned"`
}

// ParamsOfHash ...
type ParamsOfHash struct {
	Data string `json:"data"`
}

// ResultOfHash ...
type ResultOfHash struct {
	Hash string `json:"hash"`
}

// ParamsOfScrypt ...
type ParamsOfScrypt struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
	LogN     int    `json:"log_n"`
	R        int    `json:"r"`
	P        int    `json:"p"`
	DkLen    int    `json:"dk_len"`
}

// ResultOfScrypt ...
type ResultOfScrypt struct {
	Key string `json:"key"`
}

// ParamsOfNaclSignKeyPairFromSecret ...
type ParamsOfNaclSignKeyPairFromSecret struct {
	Secret string `json:"secret"`
}

// ParamsOfNaclSign ...
type ParamsOfNaclSign struct {
	Unsigned string `json:"unsigned"`
	Secret   string `json:"secret"`
}

// ResultOfNaclSign ...
type ResultOfNaclSign struct {
	Signed string `json:"signed"`
}

// ParamsOfNaclSignOpen ...
type ParamsOfNaclSignOpen struct {
	Signed string `json:"signed"`
	Public string `json:"public"`
}

// ResultOfNaclSignOpen ...
type ResultOfNaclSignOpen struct {
	Unsigned string `json:"unsigned"`
}

// ResultOfNaclSignDetached ...
type ResultOfNaclSignDetached struct {
	Signature string `json:"signature"`
}

// ParamsOfNaclBoxKeyPairFromSecret ...
type ParamsOfNaclBoxKeyPairFromSecret struct {
	Secret string `json:"secret"`
}

// ParamsOfNaclBox ...
type ParamsOfNaclBox struct {
	Decrypted   string `json:"decrypted"`
	Nonce       string `json:"nonce"`
	TheirPublic string `json:"their_public"`
	Secret      string `json:"secret"`
}

// ResultOfNaclBox ...
type ResultOfNaclBox struct {
	Encrypted string `json:"encrypted"`
}

// ParamsOfNaclBoxOpen ...
type ParamsOfNaclBoxOpen struct {
	Encrypted   string `json:"encrypted"`
	Nonce       string `json:"nonce"`
	TheirPublic string `json:"their_public"`
	Secret      string `json:"secret"`
}

// ResultOfNaclBoxOpen ...
type ResultOfNaclBoxOpen struct {
	Decrypted string `json:"decrypted"`
}

// ParamsOfNaclSecretBox ...
type ParamsOfNaclSecretBox struct {
	Decrypted string `json:"decrypted"`
	Nonce     string `json:"nonce"`
	Key       string `json:"key"`
}

// ParamsOfNaclSecretBoxOpen ...
type ParamsOfNaclSecretBoxOpen struct {
	Encrypted string `json:"encrypted"`
	Nonce     string `json:"nonce"`
	Key       string `json:"key"`
}

// ParamsOfMnemonicWords ...
type ParamsOfMnemonicWords struct {
	Dictionary int `json:"dictionary"` ///omitempty and default value
}

// ResultOfMnemonicWords ...
type ResultOfMnemonicWords struct {
	Words string `json:"words"`
}

// ParamsOfMnemonicFromRandom ...
type ParamsOfMnemonicFromRandom struct {
	Dictionary int `json:"dictionary"` ///omitempty and default value
	WordCount  int `json:"word_count"` ///omitempty and default value
}

// ResultOfMnemonicFromRandom ...
type ResultOfMnemonicFromRandom struct {
	Phrase string `json:"phrase"`
}

// ParamsOfMnemonicFromEntropy ...
type ParamsOfMnemonicFromEntropy struct {
	Entropy    string `json:"entropy"`
	Dictionary int    `json:"dictionary"` ///omitempty and default value
	WordCount  int    `json:"word_count"` ///omitempty and default value
}

// ResultOfMnemonicFromEntropy ...
type ResultOfMnemonicFromEntropy struct {
	Phrase string `json:"phrase"`
}

// ParamsOfMnemonicVerify ...
type ParamsOfMnemonicVerify struct {
	Phrase     string `json:"phrase"`
	Dictionary int    `json:"dictionary"` ///omitempty and default value
	WordCount  int    `json:"word_count"` ///omitempty and default value
}

// ResultOfMnemonicVerify ...
type ResultOfMnemonicVerify struct {
	Valid bool `json:"valid"`
}

// ParamsOfMnemonicDeriveSignKeys ...
type ParamsOfMnemonicDeriveSignKeys struct {
	Phrase     string `json:"phrase"`
	Path       string `json:"path,omitempty"`
	Dictionary int    `json:"dictionary"` ///omitempty and default value
	WordCount  int    `json:"word_count"` ///omitempty and default value
}

// ParamsOfHDKeyXPrvFromMnemonic ...
type ParamsOfHDKeyXPrvFromMnemonic struct {
	Phrase     string `json:"phrase"`
	Dictionary int    `json:"dictionary"` ///omitempty and default value
	WordCount  int    `json:"word_count"` ///omitempty and default value
}

// ResultOfHDKeyXPrvFromMnemonic ...
type ResultOfHDKeyXPrvFromMnemonic struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeyDeriveFromXPrv ...
type ParamsOfHDKeyDeriveFromXPrv struct {
	Xprv       string `json:"xprv"`
	ChildIndex int    `json:"child_index"`
	Hardened   bool   `json:"hardened"`
}

// ResultOfHDKeyDeriveFromXPrv ...
type ResultOfHDKeyDeriveFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeyDeriveFromXPrvPath ...
type ParamsOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
	Path string `json:"path"`
}

// ResultOfHDKeyDeriveFromXPrvPath ...
type ResultOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeySecretFromXPrv ...
type ParamsOfHDKeySecretFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ResultOfHDKeySecretFromXPrv ...
type ResultOfHDKeySecretFromXPrv struct {
	Secret string `json:"secret"`
}

// ParamsOfHDKeyPublicFromXPrv ...
type ParamsOfHDKeyPublicFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ResultOfHDKeyPublicFromXPrv ...
type ResultOfHDKeyPublicFromXPrv struct {
	Public string `json:"public"`
}

//  Net

const (
	SortDirectionASC  = "ASC"
	SortDirectionDESC = "DESC"
)

// OrderBy ...
type OrderBy struct {
	Path      string `json:"path"`
	Direction string `json:"direction"`
}

// ParamsOfQueryCollection ...
type ParamsOfQueryCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Order      []OrderBy   `json:"order,omitempty"`
	Limit      int         `json:"limit,omitempty"`
}

// ResultOfQueryCollection ...
type ResultOfQueryCollection struct {
	Result []json.RawMessage `json:"result"`
}

// ParamsOfWaitForCollection ...
type ParamsOfWaitForCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Timeout    int         `json:"timeout,omitempty"`
}

// ResultOfWaitForCollection ...
type ResultOfWaitForCollection struct {
	Result json.RawMessage `json:"result"`
}

// ResultOfSubscribeCollection ...
type ResultOfSubscribeCollection struct {
	Handle int `json:"handle"`
}

// ParamsOfSubscribeCollection ...
type ParamsOfSubscribeCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
}

// processing

type ProcessingEventType string

const (
	PEWillFetchFirstBlock       ProcessingEventType = "WillFetchFirstBlock"
	PEWillFetchFirstBlockFailed ProcessingEventType = "FetchFirstBlockFailed"
	PEWillSend                  ProcessingEventType = "WillSend"
	PEDidSend                   ProcessingEventType = "DidSend"
	PESendFailed                ProcessingEventType = "SendFailed"
	PEWillFetchNextBlock        ProcessingEventType = "WillFetchNextBlock"
	PEFetchNextBlockFailed      ProcessingEventType = "FetchNextBlockFailed"
	PEMessageExpired            ProcessingEventType = "MessageExpired"
)

//ProcessingEventWFFB - WillFetchFirstBlock
type ProcessingEventWFFB struct {
	Type ProcessingEventType `json:"type"`
}

//ProcessingEventFFBF - FetchFirstBlockFailed
type ProcessingEventFFBF struct {
	Type  ProcessingEventType `json:"type"`
	Error ClientError         `json:"error"`
}

//ProcessingEventWS - WillSend
type ProcessingEventWS struct {
	Type         ProcessingEventType `json:"type"`
	ShardBlockID string              `json:"shard_block_id"`
	MessageID    string              `json:"message_id"`
	Message      string              `json:"Message"`
}

//ProcessingEventDS - DidSend
type ProcessingEventDS struct {
	Type         ProcessingEventType `json:"type"`
	ShardBlockID string              `json:"shard_block_id"`
	MessageID    string              `json:"message_id"`
	Message      string              `json:"Message"`
}

//ProcessingEventSF - SendFailed
type ProcessingEventSF struct {
	Type         ProcessingEventType `json:"type"`
	ShardBlockID string              `json:"shard_block_id"`
	MessageID    string              `json:"message_id"`
	Message      string              `json:"Message"`
	Error        ClientError         `json:"error"`
}

//ProcessingEventWFNB - WillFetchNextBlock
type ProcessingEventWFNB struct {
	Type         ProcessingEventType `json:"type"`
	ShardBlockID string              `json:"shard_block_id"`
	MessageID    string              `json:"message_id"`
	Message      string              `json:"Message"`
}

//ProcessingEventFNBF - FetchNextBlockFailed
type ProcessingEventFNBF struct {
	Type         ProcessingEventType `json:"type"`
	ShardBlockID string              `json:"shard_block_id"`
	MessageID    string              `json:"message_id"`
	Message      string              `json:"Message"`
	Error        ClientError         `json:"error"`
}

//ProcessingEventME - MessageExpired
type ProcessingEventME struct {
	Type      ProcessingEventType `json:"type"`
	MessageID string              `json:"message_id"`
	Message   string              `json:"Message"`
	Error     ClientError         `json:"error"`
}

// ParamsOfSendMessage ...
type ParamsOfSendMessage struct {
	Message    string      `json:"message"`
	Abi        interface{} `json:"abi, omitempty"` //ABI??? AbiS and AbiH
	SendEvents bool        `json:"send_events"`
}

// ResultOfSendMessage ...
type ResultOfSendMessage struct {
	ShardBlockID string `json:"shard_block_id"`
}

// ParamsOfWaitForTransaction ...
type ParamsOfWaitForTransaction struct {
	Abi          interface{} `json:"abi, omitempty"` //ABI??? AbiS and AbiH
	Message      string      `json:"message"`
	ShardBlockID string      `json:"shard_block_id"`
	SendEvents   bool        `json:"send_events"`
}

// ResultOfProcessMessage ...
type ResultOfProcessMessage struct {
	Transaction json.RawMessage  `json:"transaction"`
	OutMessages []string         `json:"out_messages"`
	Decoded     *DecodedOutput   `json:"decoded,omitempty"`
	Fees        *TransactionFees `json:"fees"`
}

// DecodedOutput ...
type DecodedOutput struct {
	OutMessages []DecodedMessageBody `json:"out_messages,omitempty"`
	Output      json.RawMessage      `json:"output,omitempty"`
}

// ParamsOfProcessMessage ...
type ParamsOfProcessMessage struct {
	MessageEncodeParams *ParamsOfEncodeMessage `json:"message_encode_params"`
	SendEvents          bool                   `json:"send_events"`
}

// tvm

// ExecutionOptions ...
type ExecutionOptions struct {
	BlockchainConfig string `json:"blockchain_config,omitempty"`
	BlockTime        int    `json:"block_time,omitempty"`
	BlockLt          string `json:"block_lt,omitempty"`       //???bigint
	TransactionLt    string `json:"transaction_lt,omitempty"` //??? bigint
}

// AccountForExecutorType ...
type AccountForExecutorType string

const (
	AccountForExecutorTypeNone    = AccountForExecutorType("None")
	AccountForExecutorTypeUninit  = AccountForExecutorType("Uninit")
	AccountForExecutorTypeAccount = AccountForExecutorType("Account")
)

// AccountForExecutorNone ...
type AccountForExecutorNone struct {
	Type AccountForExecutorType `json:"type"`
}

// AccountForExecutorUninit ...
type AccountForExecutorUninit struct {
	Type AccountForExecutorType `json:"type"`
}

// AccountForExecutorAccount ...
type AccountForExecutorAccount struct {
	Type             AccountForExecutorType `json:"type"`
	Boc              string                 `json:"boc"`
	UnlimitedBalance bool                   `json:"unlimited_balance,omitempty"`
}

// ParamsOfRunExecutor ...
type ParamsOfRunExecutor struct {
	Message              string            `json:"message"`
	Account              interface{}       `json:"account"` // AccountForExecutor
	ExecutionOptions     *ExecutionOptions `json:"execution_options,omitempty"`
	Abi                  interface{}       `json:"abi, omitempty"` //ABI??? AbiS and AbiH
	SkipTransactionCheck bool              `json:"skip_transaction_check,omitempty"`
}

// ResultOfRunExecuteMessage ...
type ResultOfRunExecuteMessage struct {
	Transaction json.RawMessage  `json:"transaction,omitempty"`
	OutMessages []string         `json:"out_messages"`
	Decoded     *DecodedOutput   `json:"decoded, omitempty"`
	Account     string           `json:"account"`
	Fees        *TransactionFees `json:"fees"`
}

// ParamsOfRunTvm ...
type ParamsOfRunTvm struct {
	Message          string            `json:"message"`
	Account          string            `json:"account"`
	Abi              interface{}       `json:"abi, omitempty"` //ABI??? AbiS and AbiH
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

// ResultOfRunTvm ...
type ResultOfRunTvm struct {
	OutMessages []string       `json:"out_messages"`
	Decoded     *DecodedOutput `json:"decoded, omitempty"`
	Account     string         `json:"account"`
}

// ParamsOfRunGet ...
type ParamsOfRunGet struct {
	Account          string            `json:"account"`
	FunctionName     string            `json:"function_name,omitempty"`
	Input            interface{}       `json:"input,omitempty"`
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

// ResultOfRunGet ...
type ResultOfRunGet struct {
	Output interface{} `json:"output"`
}

// TransactionFees ...
type TransactionFees struct {
	InMsgFwdFee      string `json:"in_msg_fwd_fee"`     //???bigint
	StorageFee       string `json:"storage_fee"`        //???bigint
	GasFee           string `json:"gas_fee"`            //???bigint
	OutMsgsFwdFee    string `json:"out_msgs_fwd_fee"`   //???bigint
	TotalAccountFees string `json:"total_account_fees"` //???bigint
	TotalOutput      string `json:"total_output"`       //???bigint
}

// utils

// AddressStringFormatType ...
type AddressStringFormatType string

const (
	ASFAccountID AddressStringFormatType = "AccountId"
	ASFHex       AddressStringFormatType = "Hex"
	ASFBase64    AddressStringFormatType = "Base64"
)

// AddressStringFormatAccountID ...
type AddressStringFormatAccountID struct {
	Type AddressStringFormatType `json:"Type"`
}

// AddressStringFormatHex ...
type AddressStringFormatHex struct {
	Type AddressStringFormatType `json:"Type"`
}

// AddressStringFormatBase64 ...
type AddressStringFormatBase64 struct {
	Type   AddressStringFormatType `json:"Type"`
	Url    bool                    `json:"url"`
	Test   bool                    `json:"test"`
	Bounce bool                    `json:"bounce"`
}

// ParamsOfConvertAddress ...
type ParamsOfConvertAddress struct {
	Address    string      `json:"address"`
	OutputName interface{} `json:"output_format"` // AddressStringFormat ///??????
}

// ResultOfConvertAddress ...
type ResultOfConvertAddress struct {
	Address string `json:"address"`
}
