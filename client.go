package goton

/*
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo windows LDFLAGS: -L./lib/windows -lton_client
#include "./lib/client_method.h"

void callB(int request_id, tc_string_data_t paramsJson, int response_type, bool finished);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"sync"
	"unsafe"
)

const (
	// VersionLibSDK ...
	VersionLibSDK = "1.0.0"
)

// Client struct with client date, connect and etc.
type Client struct {
	mutx           sync.Mutex
	client         C.uint32_t
	config         *TomlConfig
	AsyncRequestID int
}

// ClientGateWay ...
type ClientGateWay interface {
	// abi
	EncodeMessageBody(ParamsOfEncodeMessageBody) int
	AttachSignatureToMessageBody(ParamsOfAttachSignatureToMessageBody) int
	EncodeMessage(ParamsOfEncodeMessage) int
	AttachSignature(ParamsOfAttachSignature) int
	DecodeMessage(ParamsOfDecodeMessage) int
	DecodeMessageBody(ParamsOfDecodeMessageBody) int
	EncodeAccount(ParamsOfEncodeAccount) int
	// boc
	ParseMessage(ParamsOfParse) int
	ParseTransaction(ParamsOfParse) int
	ParseAccount(ParamsOfParse) int
	ParseBlock(ParamsOfParse) int
	GetBlockhainConfig(ParamsOfGetBlockchainConfig) int
	// client
	Request(method, params string) int
	Version() int
	GetAPIReference() int
	GetBuildInfo() int
	// crypto
	Factorize(poF ParamsOfFactorize) int
	ModularPower(pOMP ParamsOfModularPower) int
	TonCrc16(pOTC ParamsOfTonCrc16) int
	GenerateRandomBytes(pOGRB ParamsOfGenerateRandomBytes) int
	TonPublicKeyString(pOCPTTSF ParamsOfConvertPublicKeyToTonSafeFormat) int
	GenerateRandomSignKeys() int
	Sign(pOS ParamsOfSign) int
	VerifySignature(pOVS ParamsOfVerifySignature) int
	Sha256(pOH ParamsOfHash) int
	Sha512(pOH ParamsOfHash) int
	Scrypt(sD ParamsOfScrypt) int
	NaclSignKeypairFromSecretKey(pONSKPFC ParamsOfNaclSignKeyPairFromSecret) int
	NaclSign(pONS ParamsOfNaclSign) int
	NaclSignOpen(pONSO ParamsOfNaclSignOpen) int
	NaclSignDetached(pONS ParamsOfNaclSign) int
	NaclBoxKeypair() int
	NaclBoxKeypairFromSecretKey(pONKPFS ParamsOfNaclBoxKeyPairFromSecret) int
	NaclBox(pONB ParamsOfNaclBox) int
	NaclBoxOpen(pONBO ParamsOfNaclBoxOpen) int
	NaclSecretBox(pONSB ParamsOfNaclSecretBox) int
	NaclSecretBoxOpen(pONSBO ParamsOfNaclSecretBoxOpen) int
	MnemonicWords(pOMW ParamsOfMnemonicWords) int
	MnemonicFromRandom(pOMFR ParamsOfMnemonicFromRandom) int
	MnemonicFromEntropy(pOMFE ParamsOfMnemonicFromEntropy) int
	MnemonicVerify(pOMV ParamsOfMnemonicVerify) int
	MnemonicDeriveSignKeys(pOMDSK ParamsOfMnemonicDeriveSignKeys) int
	HdkeyXprvFromMnemonic(pOHKXFM ParamsOfHDKeyXPrvFromMnemonic) int
	HdkeyXprvDerive(hdP ParamsOfHDKeyDeriveFromXPrv) int
	HdkeyXprvDerivePath(hdPD ParamsOfHDKeyDeriveFromXPrvPath) int
	HdkeyXprvSecret(pOHKSFXP ParamsOfHDKeySecretFromXPrv) int
	HdkeyXprvPublic(pOHKPFXP ParamsOfHDKeyPublicFromXPrv) int
	// net
	QueryCollection(ParamsOfQueryCollection) int
	WaitForCollection(ParamsOfWaitForCollection) int
	Unsubscribe(ResultOfSubscribeCollection) int
	SubscribeCollection(ParamsOfSubscribeCollection) int
	// processing
	SendMessage(ParamsOfSendMessage, int) int
	WaitForTransaction(ParamsOfWaitForTransaction, int) int
	ProcessMessage(ParamsOfProcessMessage, int) int
	// tvm
	RunExecutor(ParamsOfRunExecutor) int
	RunTvm(ParamsOfRunTvm) int
	RunGet(ParamsOfRunGet) int
	// utils
}

// InitClient create context and setup settings from file or default settings
func InitClient(config *TomlConfig) (*Client, error) {
	client := Client{
		config:         config,
		AsyncRequestID: 0,
	}
	configTrf, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	contM := C.CString(string(configTrf))
	defer C.free(unsafe.Pointer(contM))
	param1 := C.tc_string_data_t{content: contM, len: C.uint32_t(len(configTrf))}

	response := C.tc_create_context(param1)
	responseStr := C.tc_read_string(response)
	defer C.tc_destroy_string(response)

	stringGo := converToStringGo(responseStr.content, C.int(responseStr.len))

	var resultResp map[string]interface{}
	json.Unmarshal([]byte(stringGo), &resultResp)
	if _, ok := resultResp["error"]; ok {
		return &client, errors.New(stringGo)
	} else if elem, ok := resultResp["result"]; ok {
		client.client = C.uint32_t(elem.(float64))
		if client.client == C.uint32_t(0) {
			return &client, errors.New("Context don't connect")
		}
	}
	return &client, nil
}

// GetResp ...
func (client *Client) GetResp(resp int) (interface{}, error) {
	var (
		mapReq  *AsyncResponse
		typeRes int
	)
	for {
		State.Lock()
		nowInd := State.stores[resp]
		State.Unlock()
		typeRes = nowInd.ResponseType
		if !((typeRes == 0 || typeRes == 1) && nowInd.Finished) {
			continue
		} else {
			mapReq = nowInd
			break
		}
	}

	defer delete(State.stores, resp)
	if typeRes == 1 {
		return nil, errors.New(mapReq.Params)
	}

	return getResponce(mapReq.MethodName, mapReq.Params), nil
}

func getResponce(methodName, resp string) interface{} {
	switch methodName {
	//client
	case "client.version":
		resultStruct := ResultOfVersion{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "client.get_api_reference":
		resultStruct := ResultOfGetAPIReference{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "client.build_info":
		resultStruct := ResultOfBuildInfo{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// abi
	case "abi.encode_message_body":
		resultStruct := ResultOfEncodeMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.attach_signature_to_message_body":
		resultStruct := ResultOfAttachSignatureToMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.encode_message":
		resultStruct := ResultOfEncodeMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.attach_signature":
		resultStruct := ResultOfAttachSignature{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.decode_message", "abi.decode_message_body":
		resultStruct := DecodedMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.encode_account":
		resultStruct := ResultOfEncodeAccount{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// boc
	case "boc.parse_message", "boc.parse_transaction", "boc.parse_account", "boc.parse_block":
		resultStruct := ResultOfParse{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "get_blockchain_config":
		resultStruct := ResultOfGetBlockchainConfig{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// crypto
	case "crypto.factorize":
		resultStruct := ResultOfFactorize{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.modular_power":
		resultStruct := ResultOfModularPower{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.ton_crc16":
		resultStruct := ResultOfTonCrc16{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.generate_random_bytes":
		resultStruct := ResultOfGenerateRandomBytes{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.convert_public_key_to_ton_safe_format":
		resultStruct := ResultOfConvertPublicKeyToTonSafeFormat{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.generate_random_sign_keys", "crypto.nacl_sign_keypair_from_secret_key", "crypto.nacl_box_keypair", "crypto.nacl_box_keypair_from_secret_key", "crypto.mnemonic_derive_sign_keys":
		resultStruct := KeyPair{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.sign":
		resultStruct := ResultOfSign{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.verify_signature":
		resultStruct := ResultOfVerifySignature{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.sha256", "crypto.sha512":
		resultStruct := ResultOfHash{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.scrypt":
		resultStruct := ResultOfScrypt{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign":
		resultStruct := ResultOfNaclSign{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign_open":
		resultStruct := ResultOfNaclSignOpen{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign_detached":
		resultStruct := ResultOfNaclSignDetached{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_box", "crypto.nacl_secret_box":
		resultStruct := ResultOfNaclBox{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_box_open", "crypto.nacl_secret_box_open":
		resultStruct := ResultOfNaclBoxOpen{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_words":
		resultStruct := ResultOfMnemonicWords{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_from_random":
		resultStruct := ResultOfMnemonicFromRandom{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_from_entropy":
		resultStruct := ResultOfMnemonicFromEntropy{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_verify":
		resultStruct := ResultOfMnemonicVerify{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_xprv_from_mnemonic":
		resultStruct := ResultOfHDKeyXPrvFromMnemonic{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_derive_from_xprv":
		resultStruct := ResultOfHDKeyDeriveFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_derive_from_xprv_path":
		resultStruct := ResultOfHDKeyDeriveFromXPrvPath{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_secret_from_xprv":
		resultStruct := ResultOfHDKeySecretFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_public_from_xprv":
		resultStruct := ResultOfHDKeyPublicFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// net
	case "net.query_collection":
		resultStruct := ResultOfQueryCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "net.wait_for_collection":
		resultStruct := ResultOfWaitForCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "net.unsubscribe": //??????
		// resultStruct := ResultOfWaitForCollection{}
		// json.Unmarshal([]byte(resp), &resultStruct)
		// return resultStruct
	case "net.subscribe_collection":
		resultStruct := ResultOfSubscribeCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// processing
	case "processing.send_message":
		resultStruct := ResultOfSendMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "processing.wait_for_transaction", "processing.process_message":
		resultStruct := ResultOfProcessMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// tvm
	case "tvm.run_executor":
		resultStruct := ResultOfRunExecuteMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "tvm.run_tvm":
		resultStruct := ResultOfRunTvm{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "tvm.run_get":
		resultStruct := ResultOfRunGet{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// utils
	case "utils.convert_address":
		resultStruct := ResultOfRunGet{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	}

	return nil
}

// Destroy disconnect node
func (client *Client) Destroy() {
	client.mutx.Lock()
	defer client.mutx.Unlock()
	C.tc_destroy_context(client.client)
}

// Request ...
func (client *Client) Request(method string, params string) int {
	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_data_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_data_t{content: paramsCS, len: C.uint32_t(len(params))}

	client.mutx.Lock()
	client.AsyncRequestID++
	res := &AsyncResponse{
		ReqID:      client.AsyncRequestID,
		MethodName: method,
	}
	client.mutx.Unlock()
	State.Lock()
	State.stores[res.ReqID] = res
	State.Unlock()
	C.tc_request(client.client, param1, param2, C.uint32_t(res.ReqID), C.tc_response_handler_t(C.callB))
	return res.ReqID
}

//export callB
func callB(requestID C.int, paramsJSON C.tc_string_data_t, responseType C.int, finished C.bool) {
	State.Lock()
	defer State.Unlock()
	reg := State.stores[int(requestID)]
	reg.Params = converToStringGo(paramsJSON.content, C.int(paramsJSON.len))
	reg.ResponseType = int(responseType)
	reg.Finished = bool(finished)
}

func converToStringGo(valueString *C.char, valueLen C.int) string {
	return C.GoStringN(valueString, valueLen)
}

// Version ...
func (client *Client) Version() int {
	return client.Request("client.version", "")
}

// GetAPIReference ...
func (client *Client) GetAPIReference() int {
	return client.Request("client.get_api_reference", "")
}

// GetBuildInfo ...
func (client *Client) GetBuildInfo() int {
	return client.Request("client.build_info", "")
}
