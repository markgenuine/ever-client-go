package tvm

import (
	"encoding/json"

	crypto "github.com/move-ton/ton-client-go/crypto"
)

type SignerNone struct {
	Type string `json:"type"`
}

// type SignerExternal struct {
// 	Type      string
// 	PublicKey string
// }

type SignerKeys struct {
	Type string          `json:"type"`
	Keys *crypto.KeyPair `json:"keys"`
}

// type SignerSigningBox struct {
// 	Type   string
// 	Handle *goton.SigningBoxHandle
// }

// type Signer interface {
// None() *SignerNone
// External(publicKey string) *SignerExternal
// Keys(keys *goton.KeyPair) *SignerKeys
// SigningBox(handle *goton.SigningBoxHandle) *SignerSigningBox
// }

func GetSignerNone() *SignerNone {
	return &SignerNone{Type: "None"}
}

func GetSignerKeys(keys *crypto.KeyPair) *SignerKeys {
	return &SignerKeys{Type: "Keys", Keys: keys}
}

// func (s *SignerExternal) SignerExternal(publicKey string) *SignerExternal {
// 	return &SignerExternal{Type: "External", PublicKey: publicKey}
// }

// func (s *SignerKeys) SignerKeys(keys *goton.KeyPair) *SignerKeys {
// 	return &SignerKeys{Type: "Keys", Keys: keys}
// }

// func (s *SignerSigningBox) SignerSigningBox(handle *goton.SigningBoxHandle) *SignerSigningBox {
// 	return &SignerSigningBox{Type: "SigningBox", Handle: handle}
// }

// type MessageSourceEnc struct {
// 	Type    string     `json:"type"`
// 	Message string     `json:"message"`
// 	Abi     *goton.ABI `json:"abi,omitempty"`
// }

type Abi struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type MessageSourceEncParam struct {
	Type    string `json:"type"`
	Abi     Abi    `json:"abi"`
	Address string `json:"address"`
	// DeploySet *DeploySet `json:"deploy_set,omitempty"`
	CallSet *CallSet    `json:"call_set,omitempty"`
	Signer  *SignerKeys `json:"signer"`
	// ProcessingTryIndex int `json:"processing_try_index,omitempty"`
}

type DeploySet struct {
	Tvc         string      `json:"tvc"`
	WorkchainID int         `json:"workchain_id,omitempty "`
	InitialData interface{} `json:"initial_data, omitempty "` // initial_data?: any
}

type CallSet struct {
	FunctionName string         `json:"function_name"`
	Header       FunctionHeader `json:"header,omitempty"`
	Input        interface{}    `json:"input,omitempty"`
}

type FunctionHeader struct {
	Expire int    `json:"expire,omitempty"`
	Time   string `json:"time,omitempty"` //bigint
	PubKey string `json:"pubkey,omitempty"`
}

// type MessageSource interface {
// 	Encoded(message string, abi *goton.ABI) *MessageSourceEnc
// 	EncodingParams(abi *goton.ABI, address string, tryIndex int) *MessageSourceEncParam
// 	//deploySet DeploySet, callSet CallSet, signer Signer, tryIndex int)
// }

// func (enc *MessageSourceEnc) Encoded(message string, abi *goton.ABI) *MessageSourceEnc {
// 	return &MessageSourceEnc{Type: "Encoded", Message: message, Abi: abi}
// }

func (enc *MessageSourceEncParam) EncodingParams(abi Abi, address string) *MessageSourceEncParam {
	// deploySet DeploySet, callSet CallSet, signer Signer, tryIndex int) *MessageSourceEncParam {
	// return &MessageSourceEncParam{Type: "EncodingParams", Abi: abi, Address: address, DeploySet: &deploySet, CallSet: &callSet, Signer: &signer, ProcessingTryIndex: tryIndex}
	return &MessageSourceEncParam{Type: "EncodingParams", Abi: abi, Address: address}
	// , ProcessingTryIndex: tryIndex}
}

type ExecutionMode interface {
	Full() string
	TvmOnly() string
}

type StringT string

func (st *StringT) Full() string {
	return "Full"
}

func (st *StringT) TvmOnly() string {
	return "TvmOnly"
}

type ExecutionOptions struct {
	BlockchainConfig string `json:"blockchain_config,omitempty"`
	BlockTime        int    `json:"block_time,omitempty"`
	BlockLt          string `json:"block_lt,omitempty"`       //block_lt?: bigint
	TransactionLt    string `json:"transaction_lt,omitempty"` //transaction_lt?: bigint
}

type ParamsOfExecuteMessage struct {
	//Message          *MessageSource    `json:"message"`
	Message *MessageSourceEncParam `json:"message"`
	Account string                 `json:"account"`
	//Mode ExecutionMode `json:"mode"`
	Mode             string            `json:"mode"`
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

// // type ResultOfExecuteMessage struct {
// //     transaction?: any
// //     out_messages: any[]
// //     decoded?: DecodedOutput
// //     account?: any
// // }

// ExecuteMessage ...
func ExecuteMessage(pOEM *ParamsOfExecuteMessage) (string, string) {
	request, err := json.Marshal(pOEM)
	if err != nil {
		return "", ""
	}

	// fmt.Println(string(request))
	return "tvm.execute_message", string(request)
}

// ExecuteGet ...
// func ExecuteGet(pOEM *ParamsOfExecuteMessage) (string, string) {
// 	request, err := json.Marshal(pOEM)
// 	if err != nil {
// 		return "", ""
// 	}
// 	return "tvm.execute_get", string(request)
// }
