package contracts

import (
	"encoding/json"
	goton "github.com/move-ton/go-ton-sdk"
)

//ABI ...
type ABI struct {
	ABIVersion int      `json:"ABI version"`
	Header     []string `json:"header"`
	Functions  []struct {
		Name    string        `json:"name"`
		Inputs  []interface{} `json:"inputs"`
		Outputs []interface{} `json:"outputs"`
	} `json:"functions"`
	Data   []interface{} `json:"data"`
	Events []interface{} `json:"events"`
}

type RunFunctionCallSet struct {
	Abi          ABI    `json:"abi"`
	FunctionName string `json:"functionName"`
	Header       string `json:"header"`
	Input        string `json:"input"`
}

type ResultOfLocalRun struct {
	Output *json.RawMessage `json:"output"`
	Fees   RunFees          `json:"fees,omitempty"`
	//Account Contract         `json:"account,omitempty"`
	Account interface{} `json:"account,omitempty"`
}

type RunFees struct {
	InMsgFwdFee      string `json:"inMsgFwdFee"`
	StorageFee       string `json:"storageFee"`
	GasFee           string `json:"gasFee"`
	OutMsgsFwdFee    string `json:"outMsgsFwdFee"`
	TotalAccountFees string `json:"totalAccountFees"`
	TotalOutput      string `json:"totalOutput"`
}

type Contract struct {
	IDMsg        string      `json:"MsgAddressInt"`           //id: MsgAddressInt,
	AccType      interface{} `json:"acc_type"`                // acc_type: AccountStatus,
	Balance      string      `json:"balance"`                 // balance: u64,
	BalanceOther interface{} `json:"balance_other,omitempty"` // balance_other: Option<Vec<OtherCurrencyValue>>,
	Code         interface{} `json:"code,omitempty"`          // code: Option<Cell>,
	CodeHash     string      `json:"code_hash,omitempty"`
	DataHash     string      `json:"data_hash,omitempty"`
	Data         interface{} `json:"data,omitempty"` // data: Option<Cell>,
	Boc          interface{} `json:"boc,omitempty"`  // boc: Option<Cell>,
	LastPaid     int         `json:"last_paid"`
}

type ParamsOfLocalRunWithMsg struct {
	Address       string      `json:"address"`
	Account       interface{} `json:"account,omitempty"`
	Abi           ABI         `json:"abi,omitempty"`
	FunctionName  string      `json:"functionName,omitempty"`
	MessageBase64 string      `json:"messageBase64"`
	FullRun       bool        `json:"fullRun"`
	//Time          int         `json:"time"`
}

type ParamsOfLocalRun struct {
	Address string `json:"address"`
	//Account string `json:"account"`
	//CallSet RunFunctionCallSet `json:"call_set"`
	Abi          ABI           `json:"abi"`
	FunctionName string        `json:"functionName"`
	Input        interface{}   `json:"input"`
	KeyPair      *goton.TONKey `json:"keyPair,omitempty"`
	FullRun      bool          `json:"fullRun"`
}

type ParamsOfRun struct {
	Address string `json:"address"`
	//CallSet RunFunctionCallSet
	Abi          ABI           `json:"abi"`
	FunctionName string        `json:"functionName"`
	KeyPair      *goton.TONKey `json:"keypair,omitempty"`
	//TryIndex     int          `json:"tryIndex,omitempty"`
	Input interface{} `json:"input"`
}

type EncodedMessage struct {
	Address           string `json:"address,omitempty"`
	MessageID         string `json:"messageId"`
	MessageBodyBase64 string `json:"messageBodyBase64"`
	Expire            int    `json:"expire,omitempty"`
}

type ParamsOfLocalRunGet struct {
	Address      string          `json:"address,omitempty"`
	FunctionName string          `json:"functionName"`
	BocBase64    string          `json:"BocBase64,omitempty"`
	CodeBase64   string          `json:"CodeBase64,omitempty"`
	DataBase64   string          `json:"DataBase64,omitempty"`
	Input        json.RawMessage `json:"input,omitempty"`
	Balance      string          `json:"balance,omitempty"`
	LastPaid     int             `json:"lastPaid,omitempty"`
}

type LoadParams struct {
	Address string `json:"address"`
}

type LoadResult struct {
	ID           string `json:"id"`
	BalanceGrams string `json:"balanceGrams"`
}

type ParamsOfGetDeployAddress struct {
	Abi         ABI             `json:"abi"`
	InitParams  json.RawMessage `json:"initParams,omitempty"`
	ImageBase64 string          `json:"imageBase64"`
	KeyPair     *goton.TONKey   `json:"keyPair"`
	WorkChainID int             `json:"workChainID,omitempty"`
}

type ParamsOfConvertAddress struct {
	Address      string              `json:"address"`
	ConvertTo    string              `json:"convertTo"`
	Base64Params Base64AddressParams `json:"base64Params,omitempty"`
}

type Base64AddressParams struct {
	Url    bool `json:"url"`
	Test   bool `json:"test"`
	Bounce bool `json:"bounce"`
}

type ResultOfConvertAddress struct {
	Address string
}

type ParamsOfDeploy struct {
	//CallSet     DeployFunctionCallSet `json:"callSet"`              //call_set: DeployFunctionCallSet,
	Abi               ABI             `json:"abi"`
	ConstructorParams json.RawMessage `json:"constructorParams"`
	InitParams        json.RawMessage `json:"initParams,omitempty"` //init_params: Option<serde_json::Value>,
	ImageBase64       string          `json:"imageBase64"`          //image_base64: String,
	KeyPair           *goton.TONKey   `json:"keyPair"`
	WorkchainID       int             `json:"workchainID,omitempty"` // workchain_id: Option<i32>,
	TryIndex          int             `json:"tryIndex,omitempty"`
}

//type DeployFunctionCallSet struct {
//	Abi               ABI             `json:"abi"`
//	ConstructorHeader json.RawMessage `json:"constructorHeader,omitempty"`
//	ConstructorParams json.RawMessage `json:"constructorParams"`
//}

type ResultOfDeploy struct {
	Address         string          `json:"address"`
	AlreadyDeployed bool            `json:"alreadyDeployed"`
	Fees            RunFees         `json:"fees,omitempty"`
	Transaction     json.RawMessage `json:"transaction"`
}

type ParamsOfEncodeUnsignedDeployMessage struct {
	Abi               ABI             `json:"abi"`
	ConstructorParams json.RawMessage `json:"constructorParams"`
	//CallSet      DeployFunctionCallSet `json:"callSet"`
	InitParams   json.RawMessage `json:"initParams,omitempty"`
	ImageBase64  string          `json:"imageBase64"`
	PublicKeyHex string          `json:"publicKeyHex"`
	WorkchainID  int             `json:"workchainID,omitempty"`
	TryIndex     int             `json:"tryIndex,omitempty"`
}

type ResultOfEncodeUnsignedDeployMessage struct {
	Encoded    EncodedUnsignedMessage
	AddressHex string
}

type EncodedUnsignedMessage struct {
	UnsignedBytesBase64 string `json:"unsignedBytesBase64"`
	BytesToSignBase64   string `json:"bytesToSignBase64"`
	Expire              int    `json:"expire,omitempty"`
}

type ParamsOfGetDeployData struct {
	Abi          ABI             `json:"abi,omitempty"`
	InitParams   json.RawMessage `json:"initParams,omitempty"`
	ImageBase64  string          `json:"imageBase64,omitempty"`
	PublicKeyHex string          `json:"publicKeyHex"`
	WorkchainID  int             `json:"workchainID,omitempty"`
}

type ResultOfGetDeployData struct {
	ImageBase64 string `json:"imageBase64,omitempty"`
	AccountID   string `json:"accountID,omitempty"`
	Address     string `json:"address,omitempty"`
	DataBase64  string `json:"dataBase64"`
}

type ParamsOfResolveError struct {
	Address       string          `json:"address"`
	Account       json.RawMessage `json:"Account"`
	MessageBase64 string          `json:"messageBase64"`
	Time          int             `json:"time"`
	MainError     json.RawMessage `json:"mainError"`
}

type ParamsOfEncodeUnsignedRunMessage struct {
	Address      string `json:"address"`
	Abi          ABI    `json:"abi"`
	FunctionName string `json:"functionName"`
	//CallSet      RunFunctionCallSet `json:"callSet"`
	TryIndex int         `json:"tryIndex,omitempty"`
	Input    interface{} `json:"input"`
}

type ParamsOfGetRunBody struct {
	Abi          ABI           `json:"abi"`
	FunctionName string        `json:"function"`
	Header       interface{}   `json:"header,omitempty"`
	Params       interface{}   `json:"params"`
	Internal     bool          `json:"internal"`
	KeyPair      *goton.TONKey `json:"keyPair,omitempty"`
}

type ResultOfGetRunBody struct {
	BodyBase64 string `json:"bodyBase64"`
}

type ResultOfRun struct {
	Output      *json.RawMessage
	Fees        RunFees
	Transaction *json.RawMessage
}

type ParamsOfDecodeRunOutput struct {
	Abi          ABI    `json:"abi"`
	FunctionName string `json:"functionName"`
	BodyBase64   string `json:"bodyBase64"`
	Internal     bool   `json:"internal"`
}

type ResultOfDecode struct {
	Output *json.RawMessage `json:"output"`
}

type ParamsOfDecodeUnknownRun struct {
	Abi        ABI    `json:"abi"`
	BodyBase64 string `json:"bodyBase64"`
	Internal   bool   `json:"internal"`
}

type ResultOfDecodeUnknownRun struct {
	Function string           `json:"function"`
	Output   *json.RawMessage `json:"output"`
}

type ParamsOfEncodeMessageWithSign struct {
	Abi                 ABI    `json:"abi"`
	UnsignedBytesBase64 string `json:"unsignedBytesBase64"`
	SignBytesBase64     string `json:"signBytesBase64"`
	PublicKeyHex        string `json:"publicKeyHex,omitempty"`
	Expire              int    `json:"expire,omitempty"`
}

type InputBoc struct {
	BocBase64 string `json:"bocBase64"`
}

type ParamsOfGetFunctionId struct {
	Abi      ABI    `json:"abi"`
	Function string `json:"function"`
	Input    bool   `json:"input"`
}

type ResultOfGetFunctionId struct {
	ID int `json:"id"`
}

type ParamsOfFindShard struct {
	Shards  *json.RawMessage `json:"shards"`
	Address string           `json:"address"`
}

type ParamsOfProcessMessage struct {
	Abi          ABI            `json:"abi,omitempty"`
	FunctionName string         `json:"functionName,omitempty"`
	Message      EncodedMessage `json:"message"`
	InfiniteWait bool           `json:"infiniteWait"`
}

type ParamsOfProcessTransaction struct {
	Transaction  interface{} `json:"transaction"`
	Abi          ABI         `json:"abi,omitempty"`
	FunctionName string      `json:"functionName,omitempty"`
	Address      string      `json:"address"`
}

type ParamsOfWaitTransaction struct {
	Abi                    ABI                    `json:"abi,omitempty"`
	FunctionName           string                 `json:"functionName,omitempty"`
	Message                EncodedMessage         `json:"message"`
	MessageProcessingState MessageProcessingState `json:"messageProcessingState"`
	InfiniteWait           bool                   `json:"infiniteWait"`
}

type MessageProcessingState struct {
	LastBlockID string `json:"lastBlockId"`
	SendingTime int    `json:"sendingTime"`
}
