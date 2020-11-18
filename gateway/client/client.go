package client

/*
#cgo darwin LDFLAGS: -L../../lib/darwin -lton_client
#cgo linux LDFLAGS: -L../../lib/linux -lton_client
#cgo windows LDFLAGS: -L../../lib/windows -lton_client
#include "client_method.h"
void callB(int request_id, tc_string_data_t paramsJson, int response_type, bool finished);
*/
import "C"
import (
	"encoding/json"
	"errors"
	"sync"
	"unsafe"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/util"
)

const (
	// VersionLibSDK ...
	VersionLibSDK = "1.1.1"
)

type clientGateway struct {
	mx             sync.Mutex
	client         C.uint32_t
	config         domain.Config
	AsyncRequestID int
}

func (c *clientGateway) Request(method string, paramsIn interface{}) (int, error) {
	params, err := util.StructToJSON(paramsIn)
	if err != nil {
		return 0, err
	}

	methodsCS := C.CString(method)
	defer C.free(unsafe.Pointer(methodsCS))
	param1 := C.tc_string_data_t{content: methodsCS, len: C.uint32_t(len(method))}

	paramsCS := C.CString(params)
	defer C.free(unsafe.Pointer(paramsCS))
	param2 := C.tc_string_data_t{content: paramsCS, len: C.uint32_t(len(params))}

	//fmt.Println(params)
	c.mx.Lock()
	c.AsyncRequestID++
	res := &AsyncResponse{
		ReqID:      c.AsyncRequestID,
		MethodName: method,
	}
	c.mx.Unlock()
	State.Lock()
	State.stores[res.ReqID] = res
	State.Unlock()
	C.tc_request(c.client, param1, param2, C.uint32_t(res.ReqID), C.tc_response_handler_t(C.callB))
	return res.ReqID, nil
}

func (c *clientGateway) Version() (int, error) {
	return c.Request("client.version", "")
}

func (c *clientGateway) GetAPIReference() (int, error) {
	return c.Request("client.get_api_reference", "")
}

func (c *clientGateway) GetBuildInfo() (int, error) {
	return c.Request("client.build_info", "")
}

// NewClientGateway ...
func NewClientGateway(config domain.Config) (domain.ClientGateway, error) {
	client := clientGateway{
		mx:             sync.Mutex{},
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

	stringGo := util.CToString((*util.ExportedCChar)(responseStr.content), (util.ExportedCInt)(responseStr.len))

	var resultResp map[string]interface{}
	err = json.Unmarshal([]byte(stringGo), &resultResp)
	if err != nil {
		return nil, err
	}
	if _, ok := resultResp["error"]; ok {
		return &client, errors.New(stringGo)
	}
	if elem, ok := resultResp["result"]; ok {
		client.client = C.uint32_t(elem.(float64))
		if client.client == C.uint32_t(0) {
			return &client, errors.New("context don't connect")
		}
	}
	return &client, nil
}

func (c *clientGateway) Destroy() {
	c.mx.Lock()
	defer c.mx.Unlock()
	C.tc_destroy_context(c.client)
}

// GetResp ...
func (c *clientGateway) GetResp(resp int) (interface{}, error) {
	var (
		mapReq  *AsyncResponse
		typeRes int
	)
	for {
		State.Lock()
		nowInd := State.stores[resp]
		State.Unlock()
		typeRes = nowInd.ResponseType
		if !nowInd.Finished {
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

	return getResponse(mapReq.MethodName, mapReq.Params), nil
}

func getResponse(methodName, resp string) interface{} {
	switch methodName {
	//client
	case "client.version":
		resultStruct := domain.ResultOfVersion{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "client.get_api_reference":
		resultStruct := domain.ResultOfGetAPIReference{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "client.build_info":
		resultStruct := domain.ResultOfBuildInfo{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// abi
	case "abi.encode_message_body":
		resultStruct := domain.ResultOfEncodeMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.attach_signature_to_message_body":
		resultStruct := domain.ResultOfAttachSignatureToMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.encode_message":
		resultStruct := domain.ResultOfEncodeMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.attach_signature":
		resultStruct := domain.ResultOfAttachSignature{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.decode_message", "abi.decode_message_body":
		resultStruct := domain.DecodedMessageBody{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "abi.encode_account":
		resultStruct := domain.ResultOfEncodeAccount{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// boc
	case "boc.parse_message", "boc.parse_transaction", "boc.parse_account", "boc.parse_block", "boc.parse_shardstate":
		resultStruct := domain.ResultOfParse{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "boc.get_blockchain_config":
		resultStruct := domain.ResultOfGetBlockchainConfig{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "boc.get_boc_hash":
		resultStruct := domain.ResultOfGetBocHash{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// crypto
	case "crypto.factorize":
		resultStruct := domain.ResultOfFactorize{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.modular_power":
		resultStruct := domain.ResultOfModularPower{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.ton_crc16":
		resultStruct := domain.ResultOfTonCrc16{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.generate_random_bytes":
		resultStruct := domain.ResultOfGenerateRandomBytes{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.convert_public_key_to_ton_safe_format":
		resultStruct := domain.ResultOfConvertPublicKeyToTonSafeFormat{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.generate_random_sign_keys", "crypto.nacl_sign_keypair_from_secret_key", "crypto.nacl_box_keypair", "crypto.nacl_box_keypair_from_secret_key", "crypto.mnemonic_derive_sign_keys":
		resultStruct := domain.KeyPair{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.sign":
		resultStruct := domain.ResultOfSign{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.verify_signature":
		resultStruct := domain.ResultOfVerifySignature{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.sha256", "crypto.sha512":
		resultStruct := domain.ResultOfHash{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.scrypt":
		resultStruct := domain.ResultOfScrypt{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign":
		resultStruct := domain.ResultOfNaclSign{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign_open":
		resultStruct := domain.ResultOfNaclSignOpen{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_sign_detached":
		resultStruct := domain.ResultOfNaclSignDetached{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_box", "crypto.nacl_secret_box":
		resultStruct := domain.ResultOfNaclBox{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.nacl_box_open", "crypto.nacl_secret_box_open":
		resultStruct := domain.ResultOfNaclBoxOpen{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_words":
		resultStruct := domain.ResultOfMnemonicWords{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_from_random":
		resultStruct := domain.ResultOfMnemonicFromRandom{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_from_entropy":
		resultStruct := domain.ResultOfMnemonicFromEntropy{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.mnemonic_verify":
		resultStruct := domain.ResultOfMnemonicVerify{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_xprv_from_mnemonic":
		resultStruct := domain.ResultOfHDKeyXPrvFromMnemonic{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_derive_from_xprv":
		resultStruct := domain.ResultOfHDKeyDeriveFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_derive_from_xprv_path":
		resultStruct := domain.ResultOfHDKeyDeriveFromXPrvPath{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_secret_from_xprv":
		resultStruct := domain.ResultOfHDKeySecretFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "crypto.hdkey_public_from_xprv":
		resultStruct := domain.ResultOfHDKeyPublicFromXPrv{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// net
	case "net.query_collection":
		resultStruct := domain.ResultOfQueryCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "net.wait_for_collection":
		resultStruct := domain.ResultOfWaitForCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "net.unsubscribe":
		return []byte(resp)
	case "net.subscribe_collection":
		resultStruct := domain.ResultOfSubscribeCollection{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// processing
	case "processing.send_message":
		resultStruct := domain.ResultOfSendMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "processing.wait_for_transaction", "processing.process_message":
		resultStruct := domain.ResultOfProcessMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// tvm
	case "tvm.run_executor":
		resultStruct := domain.ResultOfRunExecuteMessage{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "tvm.run_tvm":
		resultStruct := domain.ResultOfRunTvm{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	case "tvm.run_get":
		resultStruct := domain.ResultOfRunGet{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	// utils
	case "utils.convert_address":
		resultStruct := domain.ResultOfConvertAddress{}
		json.Unmarshal([]byte(resp), &resultStruct)
		return resultStruct
	}

	return nil
}

//export callB
func callB(requestID C.int, paramsJSON C.tc_string_data_t, responseType C.int, finished C.bool) {
	State.Lock()
	defer State.Unlock()
	reg := State.stores[int(requestID)]
	reg.Params = util.CToString((*util.ExportedCChar)(paramsJSON.content), (util.ExportedCInt)(C.int(paramsJSON.len)))
	reg.ResponseType = int(responseType)
	reg.Finished = bool(finished)
}
