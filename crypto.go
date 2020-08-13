package goton

import (
	"encoding/json"
	"errors"
	"strings"
)

//FixInputMessage make InputMessage struct
func FixInputMessage(value string, format string) *InputMessage {
	inpMess := &InputMessage{}
	switch format {
	case "text":
		inpMess.Text = value
	case "hex":
		inpMess.Hex = value
	case "base64":
		inpMess.Base64 = value
	}

	return inpMess
}

//MathFactorize method crypto.math.factorize
func (client *Client) MathFactorize(value string) (*MaxFactorizeResult, error) {
	mathFact := &MaxFactorizeResult{}
	value, err := client.request("crypto.math.factorize", stringWithQuer(value))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), mathFact)
	if err != nil {
		return nil, err
	}

	return mathFact, nil
}

//MathModularPower method crypto.math.modularPower
func (client *Client) MathModularPower(modularPower *ModularPowerRequest) (string, error) {
	request, err := json.Marshal(modularPower)
	if err != nil {
		return "", err
	}
	return client.request("crypto.math.modularPower", string(request))
}

//TonCrc16 method crypto.ton_crc16
func (client *Client) TonCrc16(inpMess *InputMessage) (string, error) {
	request, err := json.Marshal(inpMess)
	if err != nil {
		return "", err
	}
	return client.request("crypto.ton_crc16", string(request))
}

//RandomGenerateBytes method crypto.random.generateBytes
//Generate string of the specified lenght length int, outputEncoding string
func (client *Client) RandomGenerateBytes(ranBytes *RandomGenerateBytesRequest) (string, error) {
	request, err := json.Marshal(ranBytes)
	if err != nil {
		return "", err
	}
	return client.request("crypto.random.generateBytes", string(request))
}

//Ed25519Keypair method crypto.ed25519.keypair
func (client *Client) Ed25519Keypair() (*TONKey, error) {
	result := &TONKey{}
	result1, err := client.request("crypto.ed25519.keypair", "{}")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//TonPublicKeyString method crypto.ton_public_key_string
func (client *Client) TonPublicKeyString(publicKey string) (string, error) {
	return client.request("crypto.ton_public_key_string", stringWithQuer(publicKey))
}

// Sha512 method crypto.sha512
// Get string sha2512 format
func (client *Client) Sha512(minpMess *MessageInputMessage) (string, error) {
	request, err := json.Marshal(minpMess)
	if err != nil {
		return "", err
	}
	return client.request("crypto.sha512", string(request))
}

//Sha256 method crypto.sha256
//Get string sha256 format
func (client *Client) Sha256(minpMess *MessageInputMessage) (string, error) {
	request, err := json.Marshal(minpMess)
	if err != nil {
		return "", err
	}
	return client.request("crypto.sha256", string(request))
}

//Scrypt method crypto.scrypt
func (client *Client) Scrypt(sD *ScryptDate) (string, error) {
	request, err := json.Marshal(sD)
	if err != nil {
		return "", err
	}

	return client.request("crypto.scrypt", string(request))
}

//NaclSignKeypair method crypto.nacl.sign.keypair
func (client *Client) NaclSignKeypair() (*TONKey, error) {
	result := &TONKey{}
	result1, err := client.request("crypto.nacl.sign.keypair", "{}")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//NaclSignKeypairFromSecretKey method crypto.nacl.sign.keypair.fromSecretKey
func (client *Client) NaclSignKeypairFromSecretKey(secretKey string) (*TONKey, error) {
	result := &TONKey{}
	result1, err := client.request("crypto.nacl.sign.keypair.fromSecretKey", stringWithQuer(secretKey))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//NaclBoxKeypair method crypto.nacl.box.keypair
func (client *Client) NaclBoxKeypair() (*TONKey, error) {
	result := &TONKey{}
	result1, err := client.request("crypto.nacl.box.keypair", "{}")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//NaclBoxKeypairFromSecretKey method crypto.nacl.box.keypair.fromSecretKey
func (client *Client) NaclBoxKeypairFromSecretKey(secretKey string) (*TONKey, error) {
	result := &TONKey{}
	result1, err := client.request("crypto.nacl.box.keypair.fromSecretKey", stringWithQuer(secretKey))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//NaclBox method crypto.nacl.box
func (client *Client) NaclBox(nBS *NaclBoxIn) (string, error) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.box", string(request))
}

//NaclBoxOpen method crypto.nacl.box.open
func (client *Client) NaclBoxOpen(nBS *NaclBoxIn) (string, error) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.box.open", string(request))
}

//NaclSecretBox method crypto.nacl.secret.box
func (client *Client) NaclSecretBox(nSB *NaclSecretBox) (string, error) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.secret.box", string(request))
}

//NaclSecretBoxOpen method crypto.nacl.secret.box.open
func (client *Client) NaclSecretBoxOpen(nSB *NaclSecretBox) (string, error) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.secret.box.open", string(request))
}

//NaclSign method crypto.nacl.sign
func (client *Client) NaclSign(nS *NaclSign) (string, error) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.sign", string(request))
}

//NaclSignOpen method crypto.nacl.sign.open
func (client *Client) NaclSignOpen(nS *NaclSign) (string, error) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.sign.open", string(request))
}

//NaclSignDetached method crypto.nacl.sign.detached
func (client *Client) NaclSignDetached(nS *NaclSign) (string, error) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.nacl.sign.detached", string(request))
}

//Mnemonic

//MnemonicWords method crypto.mnemonic.words
//Get all words from mnemonic phrases
func (client *Client) MnemonicWords() ([]string, error) {
	result, err := client.request("crypto.mnemonic.words", "{}")
	return strings.Fields(result), err
}

//MnemonicFromRandom method crypto.mnemonic.from.random
//Generate mnemonic phrase of the specified length
func (client *Client) MnemonicFromRandom(mFRR *MnemonicStructRequest) (string, error) {
	if _, ok := lensMnemonic[mFRR.WordCount]; !ok {
		mFRR.WordCount = 24
	}
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", err
	}
	return client.request("crypto.mnemonic.from.random", string(request))
}

//MnemonicFromEntropy method crypto.mnemonic.from.entropy
func (client *Client) MnemonicFromEntropy(mFRR *MnemonicStructRequest) (string, error) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", err
	}
	return client.request("crypto.mnemonic.from.entropy", string(request))
}

//MnemonicVerify method crypto.mnemonic.verify
//Checked mnemonic phrase to correct
func (client *Client) MnemonicVerify(mFRR *MnemonicStructRequest) (result bool, err error) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return false, err
	}
	resultO, err := client.request("crypto.mnemonic.verify", string(request))
	if resultO != "false" {
		result = true
	}
	return result, err
}

//MnemonicDeriveSignKeys method crypto.mnemonic.derive.sign.keys
//Convert mnemonic phrase to public and secret key
func (client *Client) MnemonicDeriveSignKeys(mnemonic string) (*TONKey, error) {
	result := &TONKey{}
	lenMnemonic := len(strings.Fields(mnemonic))
	if _, ok := lensMnemonic[lenMnemonic]; !ok {
		return nil, errors.New("Length mnemonic phrase not allowed")
	}

	type reqStr struct {
		Phrase string `json:"phrase"`
	}
	strtt := &reqStr{mnemonic}
	request, err := json.Marshal(strtt)
	if err != nil {
		return nil, err
	}

	result1, err := client.request("crypto.mnemonic.derive.sign.keys", string(request))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(result1), &result)
	return result, err
}

//HDKeys

//HdkeyXprvFromMnemonic method crypto.hdkey.xprv.from.mnemonic
//Get bip32 key from mnemonic phrase
func (client *Client) HdkeyXprvFromMnemonic(mSR *MnemonicStructRequest) (string, error) {
	lenMnemonic := len(strings.Fields(mSR.Phrase))
	if _, ok := lensMnemonic[lenMnemonic]; !ok {
		return "", errors.New("Length mnemonic phrase not allowed")
	}

	request, err := json.Marshal(mSR)
	if err != nil {
		return "", err
	}
	return client.request("crypto.hdkey.xprv.from.mnemonic", string(request))
}

//HdkeyXprvSecret method crypto.hdkey.xprv.secret
//Get secret key bip32 from bip32 key
func (client *Client) HdkeyXprvSecret(hdS *HDSerialized) (string, error) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.hdkey.xprv.secret", string(request))
}

//HdkeyXprvPublic method crypto.hdkey.xprv.public
//Get private key bip32 from bip32 key
func (client *Client) HdkeyXprvPublic(hdS *HDSerialized) (string, error) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", err
	}
	return client.request("crypto.hdkey.xprv.public", string(request))
}

// HdkeyXprvDerive method crypto.hdkey.xprv.derive
func (client *Client) HdkeyXprvDerive(hdP *HDDerivery) (string, error) {
	request, err := json.Marshal(hdP)
	if err != nil {
		return "", err
	}
	return client.request("crypto.hdkey.xprv.derive", string(request))
}

//HdkeyXprvDerivePath method crypto.hdkey.xprv.derive.path
func (client *Client) HdkeyXprvDerivePath(hdPD *HDPathDerivery) (string, error) {
	request, err := json.Marshal(hdPD)
	if err != nil {
		return "", err
	}
	return client.request("crypto.hdkey.xprv.derive.path", string(request))
}
