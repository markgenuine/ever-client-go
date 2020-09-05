package crypto

import (
	"encoding/json"
	"strings"

	goton "github.com/move-ton/go-ton-sdk"
)

func stringWithQuer(inStr string) string {
	return `"` + inStr + `"`
}

// MathFactorize method crypto.math.factorize
func MathFactorize(value string) (string, string) {
	return "crypto.math.factorize", stringWithQuer(value)
}

// MathFactorizeResp response method crypto.math.factorize
func MathFactorizeResp(resp string, err error) (*goton.MaxFactorizeResult, error) {
	if err != nil {
		return nil, err
	}

	strT := &goton.MaxFactorizeResult{}
	err = json.Unmarshal([]byte(resp), strT)
	if err != nil {
		return nil, err
	}

	return strT, nil
}

// MathModularPower method crypto.math.modularPower
func MathModularPower(modularPower *goton.ModularPowerRequest) (string, string) {
	request, err := json.Marshal(modularPower)
	if err != nil {
		return "", ""
	}
	return "crypto.math.modularPower", string(request)
}

// TonCrc16 method crypto.ton_crc16
func TonCrc16(inpMess *goton.InputMessage) (string, string) {
	request, err := json.Marshal(inpMess)
	if err != nil {
		return "", ""
	}
	return "crypto.ton_crc16", string(request)
}

// RandomGenerateBytes method crypto.random.generateBytes
func RandomGenerateBytes(ranBytes *goton.RandomGenerateBytesRequest) (string, string) {
	request, err := json.Marshal(ranBytes)
	if err != nil {
		return "", ""
	}
	return "crypto.random.generateBytes", string(request)
}

// Ed25519Keypair method crypto.ed25519.keypair
func Ed25519Keypair() (string, string) {
	return "crypto.ed25519.keypair", "{}"
}

// Ed25519KeypairResp response method crypto.ed25519.keypair
func Ed25519KeypairResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}

	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//TonPublicKeyString method crypto.ton_public_key_string
func TonPublicKeyString(publicKey string) (string, string) {
	return "crypto.ton_public_key_string", stringWithQuer(publicKey)
}

// Sha512 method crypto.sha512
func Sha512(minpMess *goton.MessageInputMessage) (string, string) {
	request, err := json.Marshal(minpMess)
	if err != nil {
		return "", ""
	}
	return "crypto.sha512", string(request)
}

// Sha256 method crypto.sha256
func Sha256(minpMess *goton.MessageInputMessage) (string, string) {
	request, err := json.Marshal(minpMess)
	if err != nil {
		return "", ""
	}
	return "crypto.sha256", string(request)
}

// Scrypt method crypto.scrypt
func Scrypt(sD *goton.ScryptDate) (string, string) {
	request, err := json.Marshal(sD)
	if err != nil {
		return "", ""
	}

	return "crypto.scrypt", string(request)
}

// NaclSignKeypair method crypto.nacl.sign.keypair
func NaclSignKeypair() (string, string) {
	return "crypto.nacl.sign.keypair", "{}"
}

// NaclSignKeypairResp method crypto.nacl.sign.keypair
func NaclSignKeypairResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}

	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSignKeypairFromSecretKey method crypto.nacl.sign.keypair.fromSecretKey
func NaclSignKeypairFromSecretKey(secretKey string) (string, string) {
	return "crypto.nacl.sign.keypair.fromSecretKey", stringWithQuer(secretKey)
}

// NaclSignKeypairFromSecretKeyResp method crypto.nacl.sign.keypair.fromSecretKey
func NaclSignKeypairFromSecretKeyResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}
	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBoxKeypair method crypto.nacl.box.keypair
func NaclBoxKeypair() (string, string) {
	return "crypto.nacl.box.keypair", "{}"
}

// NaclBoxKeypairResp method crypto.nacl.box.keypair
func NaclBoxKeypairResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}

	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBoxKeypairFromSecretKey method crypto.nacl.box.keypair.fromSecretKey
func NaclBoxKeypairFromSecretKey(secretKey string) (string, string) {
	return "crypto.nacl.box.keypair.fromSecretKey", stringWithQuer(secretKey)
}

// NaclBoxKeypairFromSecretKeyResp method crypto.nacl.box.keypair.fromSecretKey
func NaclBoxKeypairFromSecretKeyResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}

	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)

	return result, err
}

// NaclBox method crypto.nacl.box
func NaclBox(nBS *goton.NaclBoxIn) (string, string) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.box", string(request)
}

// NaclBoxOpen method crypto.nacl.box.open
func NaclBoxOpen(nBS *goton.NaclBoxIn) (string, string) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.box.open", string(request)
}

// NaclSecretBox method crypto.nacl.secret.box
func NaclSecretBox(nSB *goton.NaclSecretBox) (string, string) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.secret.box", string(request)
}

// NaclSecretBoxOpen method crypto.nacl.secret.box.open
func NaclSecretBoxOpen(nSB *goton.NaclSecretBox) (string, string) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.secret.box.open", string(request)
}

// NaclSign method crypto.nacl.sign
func NaclSign(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}

	return "crypto.nacl.sign", string(request)
}

// NaclSignOpen method crypto.nacl.sign.open
func NaclSignOpen(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.sign.open", string(request)
}

// NaclSignDetached method crypto.nacl.sign.detached
func NaclSignDetached(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl.sign.detached", string(request)
}

// Mnemonic

// MnemonicWords method crypto.mnemonic.words
func MnemonicWords() (string, string) {
	return "crypto.mnemonic.words", "{}"
}

// MnemonicWordsResp method crypto.mnemonic.words
func MnemonicWordsResp(resp string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	return strings.Fields(resp), nil
}

// MnemonicFromRandom method crypto.mnemonic.from.random
func MnemonicFromRandom(mFRR *goton.MnemonicStructRequest) (string, string) {
	if _, ok := goton.LensMnemonic[mFRR.WordCount]; !ok {
		mFRR.WordCount = 24
	}
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic.from.random", string(request)
}

// MnemonicFromEntropy method crypto.mnemonic.from.entropy
func MnemonicFromEntropy(mFRR *goton.MnemonicStructRequest) (string, string) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic.from.entropy", string(request)
}

// MnemonicVerify method crypto.mnemonic.verify
func MnemonicVerify(mFRR *goton.MnemonicStructRequest) (string, string) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic.verify", string(request)
}

// MnemonicVerifyResp method crypto.mnemonic.verify
func MnemonicVerifyResp(resp string, err error) (bool, error) {
	result := false
	if err != nil {
		return result, err
	}

	if resp != "false" {
		result = true
	}
	return result, nil
}

// MnemonicDeriveSignKeys method crypto.mnemonic.derive.sign.keys
func MnemonicDeriveSignKeys(mnemonic string) (string, string) {
	type reqStr struct {
		Phrase string `json:"phrase"`
	}
	strtt := &reqStr{mnemonic}
	request, err := json.Marshal(strtt)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic.derive.sign.keys", string(request)
}

// MnemonicDeriveSignKeysResp method crypto.mnemonic.derive.sign.keys
func MnemonicDeriveSignKeysResp(resp string, err error) (*goton.TONKey, error) {
	if err != nil {
		return nil, err
	}

	result := &goton.TONKey{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey.xprv.from.mnemonic
func HdkeyXprvFromMnemonic(mSR *goton.MnemonicStructRequest) (string, string) {
	lenMnemonic := len(strings.Fields(mSR.Phrase))
	if _, ok := goton.LensMnemonic[lenMnemonic]; !ok {
		return "", ""
	}

	request, err := json.Marshal(mSR)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey.xprv.from.mnemonic", string(request)
}

// HdkeyXprvSecret method crypto.hdkey.xprv.secret
func HdkeyXprvSecret(hdS *goton.HDSerialized) (string, string) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey.xprv.secret", string(request)
}

// HdkeyXprvPublic method crypto.hdkey.xprv.public
func HdkeyXprvPublic(hdS *goton.HDSerialized) (string, string) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey.xprv.public", string(request)
}

// HdkeyXprvDerive method crypto.hdkey.xprv.derive
func HdkeyXprvDerive(hdP *goton.HDDerivery) (string, string) {
	request, err := json.Marshal(hdP)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey.xprv.derive", string(request)
}

// HdkeyXprvDerivePath method crypto.hdkey.xprv.derive.path
func HdkeyXprvDerivePath(hdPD *goton.HDPathDerivery) (string, string) {
	request, err := json.Marshal(hdPD)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey.xprv.derive.path", string(request)
}
