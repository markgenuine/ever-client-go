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

type ResultOfExecuteMessage struct {
	Transaction interface{}   `json:"transaction,omitempty"`
	OutMessages []interface{} `json:"out_messages"`
	Decoded     DecodedOutput `json:"decoded,omitempty"`
	Account     interface{}   `json:"account"`
}

type DecodedOutput struct {
	// out_messages: DecodedMessageBody | null[],
	Output json.RawMessage `json:"output"`
}

// out_messages: DecodedMessageBody?[] – Decoded bodies of the out messages.
// output?: any – Decoded body of the function output message.

// DecodedMessageBody

// type DecodedMessageBody = {
//     body_type: MessageBodyType,
//     name: string,
//     value?: any,
//     header?: FunctionHeader
// };
// body_type: MessageBodyType – Type of the message body content.
// name: string – Function or event name.
// value?: any – Parameters or result value.
// header?: FunctionHeader – Function header.

// MessageBodyType

// type MessageBodyType = 'Input' | 'Output' | 'InternalOutput' | 'Event';
// One of the following value:

// Input – Message contains the input of the ABI function.
// Output – Message contains the output of the ABI function.
// InternalOutput – Message contains the input of the imported ABI function.
// Event – Message contains the input of the ABI event.

// FunctionHeader

// The ABI function header.

// Includes several hidden function parameters that contract uses for security and replay protection reasons.

// The actual set of header fields depends on the contract's ABI.

// type FunctionHeader = {
//     expire?: number,
//     time?: bigint,
//     pubkey?: string
// };
// expire?: number – Message expiration time in seconds.
// time?: bigint – Message creation time in milliseconds.
// pubkey?: string – Public key used to sign message. Encoded with hex.

// ExecuteMessage ...
func ExecuteMessage(pOEM *ParamsOfExecuteMessage) (string, string) {
	request, err := json.Marshal(pOEM)
	if err != nil {
		return "", ""
	}

	return "tvm.execute_message", string(request)
}

// ExecuteMessageResult ...
func ExecuteMessageResult(resp string, err error) (*ResultOfExecuteMessage, error) {
	if err != nil {
		return nil, err
	}

	rOEM := &ResultOfExecuteMessage{}
	err = json.Unmarshal([]byte(resp), rOEM)
	if err != nil {
		return nil, err
	}

	return rOEM, nil
}

// ExecuteGet ...
// func ExecuteGet(pOEM *ParamsOfExecuteMessage) (string, string) {
// 	request, err := json.Marshal(pOEM)
// 	if err != nil {
// 		return "", ""
// 	}
// 	return "tvm.execute_get", string(request)
// }
