package goton

/*
#cgo linux LDFLAGS: -L./lib/linux -lton_client
#cgo darwin LDFLAGS: -L./lib/darwin -lton_client
#include "./lib/client_method.h"
*/
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

type Client struct {
	mu     sync.Mutex
	client C.uint32_t
}

type TCStringT struct {
	Field C.tc_string_t
}

func InitClient(config *TomlConfig) (*Client, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	res, err := client.setup(config)
	if err != nil {
		client.Destroy()
		return nil, err
	}

	fmt.Println("Setup settings response: ", res)

	return client, nil

}

//NewClient create connect node
func NewClient() (*Client, error) {

	client := Client{
		mu:     sync.Mutex{},
		client: C.tc_create_context(),
	}

	if client.client != C.uint32_t(1) {
		return &client, errors.New("Context don't connect")
	}

	return &client, nil

}

//Destroy disconnect node
func (client *Client) Destroy() {
	client.mu.Lock()
	defer client.mu.Unlock()
	C.tc_destroy_context(client.client)
}

//version
func (client *Client) Version() (result string, err error) {
	return client.request("version", "")
}

func (client *Client) setup(config *TomlConfig) (result string, err error) {
	req, err := json.Marshal(&config)
	if err != nil {
		err = errors.New("Error conver to config in json!")
		return
	}
	result, err = client.request("setup", string(req))
	return
}

//crypto.mnemonic.derive.sign.keys
func (client *Client) mnemonicDeriveSignKeys(mnemonic string) (result TONKey, err error) {
	result1, err := client.request("crypto.mnemonic.derive.sign.keys", `{"phrase":"`+mnemonic+`"}`)
	err = json.Unmarshal([]byte(result1), &result)
	return
}

//crypto.hdkey.xprv.secret
func (client *Client) hdkeyXprvSecret(secret string) (result string, err error) {
	return client.request("crypto.hdkey.xprv.secret", `{"serialized":"`+secret+`"}`)
}

//crypto.hdkey.xprv.derive.path
func (client *Client) hdkeyXprvDerivePath(key, path string) (result string, err error) {
	return client.request("crypto.hdkey.xprv.derive.path", `{"serialized":"`+key+`", "path":"`+path+`"}`)
}

//crypto.hdkey.xprv.derive
func (client *Client) hdkeyXprvDerive(key, childIndex string) (result string, err error) {
	return client.request("crypto.hdkey.xprv.derive", `{"serialized":"`+key+`", "index":"`+childIndex+`"}`)
}

//crypto.hdkey.xprv.from.mnemonic
func (client *Client) hdkeyXprvFromMnemonic(mnemonic string) (result string, err error) {
	return client.request("crypto.hdkey.xprv.from.mnemonic", `{"phrase":"`+mnemonic+`"}`)
}

//crypto.nacl.sign.open
func (client *Client) naclSignOpen(key, value, valueType string) (result string, err error) {
	return client.request("crypto.nacl.sign.open", `{"message":`+fixInputMessage(value, valueType)+`,"key":"`+key+`"}`)
}

//crypto.nacl.box.open
func (client *Client) naclBoxOpen(value, valueType, nonce, theirPublic, secret string) (result string, err error) {
	return client.request("crypto.nacl.sign.open", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"theirPublic":"`+theirPublic+`,"secret":"`+secret+`"}`)
}

//crypto.keystore.remove
func (client *Client) keystoreRemove(handle string) (result string, err error) {
	return client.request("crypto.keystore.remove", `"`+handle+`"`)
}

//crypto.nacl.secret.box
func (client *Client) naclSecretBox(value, valueType, nonce, secret string) (result string, err error) {
	return client.request("crypto.nacl.secret.box", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"key":"`+secret+`"}`)
}

//crypto.nacl.sign.detached
func (client *Client) naclSignDetached(value, valueType, secret string) (result string, err error) {
	return client.request("crypto.nacl.sign.detached", `{"message":`+fixInputMessage(value, valueType)+`,"key":"`+secret+`"}`)
}

//crypto.keystore.clear
func (client *Client) keystoreClear(handle string) (result string, err error) {
	return client.request("crypto.keystore.clear", "{}")
}

//crypto.keystore.add
func (client *Client) keystoreAdd(keyStore *TONKey) (result string, err error) {
	req, err := json.Marshal(&keyStore)
	if err != nil {
		return "", errors.New("Error conver to TONKey")
	}

	return client.request("crypto.keystore.add", string(req))
}

//crypto.nacl.sign.keypair.fromSecretKey
func (client *Client) naclSignKeypairFromSecretKey(secretKey string) (result string, err error) {
	return client.request("crypto.nacl.sign.keypair.fromSecretKey", `"`+secretKey+`"`)
}

//crypto.nacl.secret.box.open
func (client *Client) naclSecretBoxOpen(value, valueType, nonce, theirPublic string) (result string, err error) {
	return client.request("crypto.nacl.secret.box.open", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"key":"`+theirPublic+`"}`)
}

//crypto.nacl.sign.keypair
func (client *Client) naclSignKeypair(handle string) (result string, err error) {
	return client.request("rypto.nacl.sign.keypair", "{}")
}

//crypto.nacl.sign
func (client *Client) naclSign(secretKey string) (result string, err error) {
	return client.request("crypto.nacl.sign", `"`+secretKey+`"`)
}

//crypto.math.modularPower
func (client *Client) mathModularPower(base, exponent, modulus string) (result string, err error) {
	return client.request("crypto.math.modularPower", `{"base":"`+base+`", "exponent":"`+exponent+`", "modulus":"`+modulus+`"}`)
}

//crypto.mnemonic.from.random
func (client *Client) mnemonicFromRandom() (result string, err error) {
	return client.request("crypto.mnemonic.from.random", "{}")
}

//crypto.mnemonic.verify
func (client *Client) mnemonicVerify(mnemonic string) (result bool, err error) {
	resultO, err := client.request("crypto.mnemonic.verify", `{"phrase":"`+mnemonic+`"}`)
	if resultO != "false" {
		result = true
	}
	return
}

//crypto.random.generateBytes
func (client *Client) randomGenerateBytes(length string) (result string, err error) {
	return client.request("crypto.random.generateBytes", `{"length":"`+length+`"}`)
}

//crypto.math.factorize
func (client *Client) mathFactorize(hex string) (result string, err error) {
	return client.request("crypto.math.factorize", `"`+hex+`"`)
}

//crypto.nacl.box.keypair.fromSecretKey
func (client *Client) naclBoxKeypairFromSecretKey(secretKey string) (result string, err error) {
	return client.request("crypto.nacl.box.keypair.fromSecretKey", `"`+secretKey+`"`)
}

//crypto.nacl.box
func (client *Client) naclBox(nonce, theirPublicKey, value, valueType string) (result string, err error) {
	return client.request("crypto.nacl.box", `{"nonce":"`+nonce+`", "theirPublicKey":"`+theirPublicKey+`", "message":"`+fixInputMessage(value, valueType)+`"}`)
}

//crypto.hdkey.xprv.public
func (client *Client) hdkeyXprvPublic(publicKey string) (result string, err error) {
	return client.request("crypto.hdkey.xprv.public", `{"serialized":"`+publicKey+`"}`)
}

//crypto.scrypt
func (client *Client) scrypt(data, password, passwordType, salt, saltType, logN, r, p, dkLen string) (result string, err error) {
	return client.request("crypto.scrypt", `{"data":"`+data+`", "salt":"`+fixInputMessage(salt, saltType)+`", "password":"`+fixInputMessage(password, passwordType)+`", "logN":"`+logN+`", "r":"`+r+`", "p":"`+p+`", "dkLen":"`+dkLen+`"}`)
}

//crypto.ed25519.keypair
func (client *Client) ed25519Keypair() (result string, err error) {
	return client.request("crypto.ed25519.keypair", "{}")
}

//crypto.mnemonic.words
func (client *Client) mnemonicWords() (result string, err error) {
	return client.request("crypto.mnemonic.words", "{}")
}

//crypto.nacl.box.keypair
func (client *Client) naclBoxKeypair() (result string, err error) {
	return client.request("crypto.nacl.box.keypair", "{}")
}

//!!!crypto.mnemonic.from.entropy (text, base64)
func (client *Client) mnemonicFromEntropy(hex string) (result string, err error) {
	return client.request("crypto.mnemonic.from.entropy", `{"entropy":{"text":"`+hex+`"}}`)
}

//!!!crypto.ton_crc16(еще добавить параметры text и base64, может быть сделат структурой и в байты???)
func (client *Client) tonCrc16(hex string) (result string, err error) {
	return client.request("crypto.ton_crc16", `{"hex":"`+hex+`"}`)
}

//crypto.ton_public_key_string
func (client *Client) tonPublicKeyString(publicKey string) (result string, err error) {
	return client.request("crypto.ton_public_key_string", `"`+publicKey+`"`)
}

//crypto.sha256
func (client *Client) sha256(value, valueType string) (result string, err error) {
	result, err = client.request("crypto.sha256", `{"message":`+fixInputMessage(value, valueType)+`}`)
	return
}

//crypto.sha512
func (client *Client) sha512(value, valueType string) (result string, err error) {
	return client.request("crypto.sha512", `{"message":`+fixInputMessage(value, valueType)+`}`)
}

func fixInputMessage(value string, format string) (message string) {
	switch value {
	case "text":
		message = `{"text":"` + value + `"}`
	case "hex":
		message = `{"hex":"` + value + `"}`
	case "base64":
		message = `{"base64":"` + value + `"}`
	}

	return
}

func (client *Client) request(method, params string) (string, error) {

	methodR := C.CString(method)
	defer C.free(unsafe.Pointer(methodR))
	param1 := TCStringT{Field: C.tc_string_t{content: methodR, len: C.uint32_t(len(method))}}

	paramsR := C.CString(params)
	defer C.free(unsafe.Pointer(paramsR))
	param2 := TCStringT{Field: C.tc_string_t{content: paramsR, len: C.uint32_t(len(params))}}

	tcResponseHandle := C.tc_json_request(client.client, param1.Field, param2.Field)
	defer C.tc_destroy_json_response(tcResponseHandle)

	tcResponse := C.tc_read_json_response(tcResponseHandle)

	resultJSON := tcResponse.result_json
	errorJSON := tcResponse.error_json

	if errorJSON.len > 0 {
		return "", errors.New(C.GoString(errorJSON.content))
	}

	return C.GoString(resultJSON.content), nil
}
