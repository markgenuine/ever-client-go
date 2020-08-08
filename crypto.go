package goton

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

//MnemonicDeriveSignKeys method crypto.mnemonic.derive.sign.keys
func (client *Client) MnemonicDeriveSignKeys(mnemonic string) (TONKey, error) {
	result1, err := client.request("crypto.mnemonic.derive.sign.keys", `{"phrase":"`+mnemonic+`"}`)
	var result TONKey
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//HdkeyXprvSecret method crypto.hdkey.xprv.secret
func (client *Client) HdkeyXprvSecret(secret string) (string, error) {
	return client.request("crypto.hdkey.xprv.secret", `{"serialized":"`+secret+`"}`)
}

//HdkeyXprvDerivePath method crypto.hdkey.xprv.derive.path
func (client *Client) HdkeyXprvDerivePath(key, path string) (string, error) {
	return client.request("crypto.hdkey.xprv.derive.path", `{"serialized":"`+key+`", "path":"`+path+`"}`)
}

//HdkeyXprvDerive method crypto.hdkey.xprv.derive
func (client *Client) HdkeyXprvDerive(key, childIndex string) (string, error) {
	return client.request("crypto.hdkey.xprv.derive", `{"serialized":"`+key+`", "index":"`+childIndex+`"}`)
}

//HdkeyXprvFromMnemonic method crypto.hdkey.xprv.from.mnemonic
func (client *Client) HdkeyXprvFromMnemonic(mnemonic string) (string, error) {
	return client.request("crypto.hdkey.xprv.from.mnemonic", `{"phrase":"`+mnemonic+`"}`)
}

//NaclSignOpen method crypto.nacl.sign.open
func (client *Client) NaclSignOpen(key, value, valueType string) (string, error) {
	return client.request("crypto.nacl.sign.open", `{"message":`+fixInputMessage(value, valueType)+`,"key":"`+key+`"}`)
}

//NaclBoxOpen method crypto.nacl.box.open
func (client *Client) NaclBoxOpen(value, valueType, nonce, theirPublic, secret string) (string, error) {
	return client.request("crypto.nacl.sign.open", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"theirPublic":"`+theirPublic+`,"secret":"`+secret+`"}`)
}

//KeystoreRemove method crypto.keystore.remove
func (client *Client) KeystoreRemove(handle string) (string, error) {
	return client.request("crypto.keystore.remove", `"`+handle+`"`)
}

//NaclSecretBox method crypto.nacl.secret.box
func (client *Client) NaclSecretBox(value, valueType, nonce, secret string) (string, error) {
	return client.request("crypto.nacl.secret.box", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"key":"`+secret+`"}`)
}

//NaclSignDetached method crypto.nacl.sign.detached
func (client *Client) NaclSignDetached(value, valueType, secret string) (string, error) {
	return client.request("crypto.nacl.sign.detached", `{"message":`+fixInputMessage(value, valueType)+`,"key":"`+secret+`"}`)
}

//KeystoreClear method crypto.keystore.clear
func (client *Client) KeystoreClear(handle string) (string, error) {
	return client.request("crypto.keystore.clear", "{}")
}

//KeystoreAdd method crypto.keystore.add
func (client *Client) KeystoreAdd(keyStore *TONKey) (string, error) {
	req, err := json.Marshal(&keyStore)
	if err != nil {
		return "", errors.New("Error conver to TONKey")
	}

	return client.request("crypto.keystore.add", string(req))
}

//NaclSignKeypairFromSecretKey method crypto.nacl.sign.keypair.fromSecretKey
func (client *Client) NaclSignKeypairFromSecretKey(secretKey string) (string, error) {
	return client.request("crypto.nacl.sign.keypair.fromSecretKey", `"`+secretKey+`"`)
}

//NaclSecretBoxOpen method crypto.nacl.secret.box.open
func (client *Client) NaclSecretBoxOpen(value, valueType, nonce, theirPublic string) (string, error) {
	return client.request("crypto.nacl.secret.box.open", `{"message":`+fixInputMessage(value, valueType)+`,"nonce":"`+nonce+`,"key":"`+theirPublic+`"}`)
}

//NaclSignKeypair method crypto.nacl.sign.keypair
func (client *Client) NaclSignKeypair(handle string) (string, error) {
	return client.request("rypto.nacl.sign.keypair", "{}")
}

//NaclSign method crypto.nacl.sign
func (client *Client) NaclSign(secretKey string) (string, error) {
	return client.request("crypto.nacl.sign", `"`+secretKey+`"`)
}

//MathModularPower method crypto.math.modularPower
func (client *Client) MathModularPower(base, exponent, modulus string) (string, error) {
	return client.request("crypto.math.modularPower", `{"base":"`+base+`", "exponent":"`+exponent+`", "modulus":"`+modulus+`"}`)
}

//MnemonicFromRandom method crypto.mnemonic.from.random
func (client *Client) MnemonicFromRandom() (string, error) {
	return client.request("crypto.mnemonic.from.random", "{}")
}

//MnemonicVerify method crypto.mnemonic.verify !!!CHECKED!!!
func (client *Client) MnemonicVerify(mnemonic string) (result bool, err error) {
	resultO, err := client.request("crypto.mnemonic.verify", `{"phrase":"`+mnemonic+`"}`)
	if resultO != "false" {
		result = true
	}
	return result, err
}

//RandomGenerateBytes method crypto.random.generateBytes
func (client *Client) RandomGenerateBytes(length int) (string, error) {
	result, err := client.request("crypto.random.generateBytes", `{"length":`+strconv.Itoa(length)+`}`)
	result = strings.ReplaceAll(result,`"`,"")
	return result, err
}

//MathFactorize method crypto.math.factorize
func (client *Client) MathFactorize(hex string) (string, error) {
	return client.request("crypto.math.factorize", `"`+hex+`"`)
}

//NaclBoxKeypairFromSecretKey method crypto.nacl.box.keypair.fromSecretKey
func (client *Client) NaclBoxKeypairFromSecretKey(secretKey string) (string, error) {
	return client.request("crypto.nacl.box.keypair.fromSecretKey", `"`+secretKey+`"`)
}

//NaclBox method crypto.nacl.box
func (client *Client) NaclBox(nonce, theirPublicKey, value, valueType string) (string, error) {
	return client.request("crypto.nacl.box", `{"nonce":"`+nonce+`", "theirPublicKey":"`+theirPublicKey+`", "message":"`+fixInputMessage(value, valueType)+`"}`)
}

//HdkeyXprvPublic method crypto.hdkey.xprv.public
func (client *Client) HdkeyXprvPublic(publicKey string) (string, error) {
	return client.request("crypto.hdkey.xprv.public", `{"serialized":"`+publicKey+`"}`)
}

//Scrypt method crypto.scrypt
func (client *Client) Scrypt(data, password, passwordType, salt, saltType, logN, r, p, dkLen string) (string, error) {
	return client.request("crypto.scrypt", `{"data":"`+data+`", "salt":"`+fixInputMessage(salt, saltType)+`", "password":"`+fixInputMessage(password, passwordType)+`", "logN":"`+logN+`", "r":"`+r+`", "p":"`+p+`", "dkLen":"`+dkLen+`"}`)
}

//Ed25519Keypair method crypto.ed25519.keypair
func (client *Client) Ed25519Keypair() (string, error) {
	return client.request("crypto.ed25519.keypair", "{}")
}

//MnemonicWords method crypto.mnemonic.words return slice byte!!!!!
func (client *Client) MnemonicWords() (string, error) {
	return client.request("crypto.mnemonic.words", "{}")
}

//NaclBoxKeypair method crypto.nacl.box.keypair
func (client *Client) NaclBoxKeypair() (string, error) {
	return client.request("crypto.nacl.box.keypair", "{}")
}

//MnemonicFromEntropy method crypto.mnemonic.from.entropy (text, base64) !!!
func (client *Client) MnemonicFromEntropy(hex string) (string, error) {
	return client.request("crypto.mnemonic.from.entropy", `{"entropy":{"text":"`+hex+`"}}`)
}

//TonCrc16 method crypto.ton_crc16(еще добавить параметры text и base64, может быть сделат структурой и в байты???)!!!
func (client *Client) TonCrc16(hex string) (string, error) {
	return client.request("crypto.ton_crc16", `{"hex":"`+hex+`"}`)
}

//TonPublicKeyString method crypto.ton_public_key_string
func (client *Client) TonPublicKeyString(publicKey string) (string, error) {
	return client.request("crypto.ton_public_key_string", `"`+publicKey+`"`)
}

//Sha256 method crypto.sha256
func (client *Client) Sha256(value, valueType string) (string, error) {
	return client.request("crypto.sha256", `{"message":`+fixInputMessage(value, valueType)+`}`)
}

//Sha512 method crypto.sha512
func (client *Client) Sha512(value, valueType string) (string, error) {
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

	return message
}
