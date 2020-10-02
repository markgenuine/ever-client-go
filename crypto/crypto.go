package crypto

import (
	"encoding/json"
	"strings"

	goton "github.com/move-ton/ton-client-go"
)

func stringWithQuer(inStr string) string {
	return `"` + inStr + `"`
}

// MathFactorize method crypto.factorize
func MathFactorize(value string) (string, string) {
	return "crypto.factorize", stringWithQuer(value)
}

// MathFactorizeResp response method crypto.factorize
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

// MathModularPower method crypto.modular_power
func MathModularPower(modularPower *goton.ModularPowerRequest) (string, string) {
	request, err := json.Marshal(modularPower)
	if err != nil {
		return "", ""
	}
	return "crypto.modular_power", string(request)
}

// TonCrc16 method crypto.ton_crc16
func TonCrc16(inpMess *goton.InputMessage) (string, string) {
	request, err := json.Marshal(inpMess)
	if err != nil {
		return "", ""
	}
	return "crypto.ton_crc16", string(request)
}

// RandomGenerateBytes method crypto.generate_random_bytes
func RandomGenerateBytes(ranBytes *goton.RandomGenerateBytesRequest) (string, string) {
	request, err := json.Marshal(ranBytes)
	if err != nil {
		return "", ""
	}
	return "crypto.generate_random_bytes", string(request)
}

// Ed25519Keypair method crypto.ed25519.keypair call GenerateRandomSignKeys
func Ed25519Keypair() (string, string) {
	return GenerateRandomSignKeys()
}

//GenerateRandomSignKeys call "crypto.generate_random_sign_keys"
func GenerateRandomSignKeys() (string, string) {
	return "crypto.generate_random_sign_keys", "{}"
}

// GenerateRandomSignKeysResp response method crypto.generate_random_sign_keys
func GenerateRandomSignKeysResp(resp string, err error) (*goton.TONKey, error) {
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

//TonPublicKeyString method crypto.convert_public_key_to_ton_safe_format
func TonPublicKeyString(publicKey string) (string, string) {
	return "crypto.convert_public_key_to_ton_safe_format", stringWithQuer(publicKey)
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

// NaclSignKeypair method crypto.nacl_sign_keypair call GenerateRandomSignKeys
func NaclSignKeypair() (string, string) {
	return GenerateRandomSignKeys()
}

// NaclSignKeypairFromSecretKey method crypto.nacl_sign_keypair_from_secret
func NaclSignKeypairFromSecretKey(secretKey string) (string, string) {
	return "crypto.nacl_sign_keypair_from_secret", stringWithQuer(secretKey)
}

// NaclBoxKeypair method crypto.nacl_box_keypair
func NaclBoxKeypair() (string, string) {
	return "crypto.nacl_box_keypair", "{}"
}

// NaclBoxKeypairFromSecretKey method crypto.nacl_box_keypair_from_secret
func NaclBoxKeypairFromSecretKey(secretKey string) (string, string) {
	return "crypto.nacl_box_keypair_from_secret", stringWithQuer(secretKey)
}

// NaclBox method crypto.nacl_box
func NaclBox(nBS *goton.NaclBoxIn) (string, string) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_box", string(request)
}

// NaclBoxOpen method crypto.nacl_box_open
func NaclBoxOpen(nBS *goton.NaclBoxIn) (string, string) {
	request, err := json.Marshal(nBS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_box_open", string(request)
}

// NaclSecretBox method crypto.nacl_secret_box
func NaclSecretBox(nSB *goton.NaclSecretBox) (string, string) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_secret_box", string(request)
}

// NaclSecretBoxOpen method crypto.nacl_secret_box_open
func NaclSecretBoxOpen(nSB *goton.NaclSecretBox) (string, string) {
	request, err := json.Marshal(nSB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_secret_box_open", string(request)
}

// NaclSign method crypto.nacl_sign
func NaclSign(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}

	return "crypto.nacl_sign", string(request)
}

// NaclSignOpen method crypto.nacl_sign_open
func NaclSignOpen(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_sign_open", string(request)
}

// NaclSignDetached method crypto.nacl_sign_detached
func NaclSignDetached(nS *goton.NaclSign) (string, string) {
	request, err := json.Marshal(nS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_sign_detached", string(request)
}

// Mnemonic

// MnemonicWords method crypto.mnemonic_words
func MnemonicWords() (string, string) {
	return "crypto.mnemonic_words", "{}"
}

// MnemonicWordsResp method crypto.mnemonic_words
func MnemonicWordsResp(resp string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	return strings.Fields(resp), nil
}

// MnemonicFromRandom method crypto.mnemonic_from_random
func MnemonicFromRandom(mFRR *goton.MnemonicStructRequest) (string, string) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic_from_random", string(request)
}

// MnemonicFromEntropy method crypto.mnemonic_from_entropy
func MnemonicFromEntropy(mFRR *goton.MnemonicStructRequest) (string, string) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic_from_entropy", string(request)
}

// MnemonicVerify method crypto.mnemonic_verify
func MnemonicVerify(mFRR *goton.MnemonicStructRequest) (string, string) {
	request, err := json.Marshal(mFRR)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic_verify", string(request)
}

// MnemonicVerifyResp method crypto.mnemonic_verify
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

// MnemonicDeriveSignKeys method crypto.mnemonic_derive_sign_keys
func MnemonicDeriveSignKeys(mnemonic string) (string, string) {
	type reqStr struct {
		Phrase string `json:"phrase"`
	}
	strtt := &reqStr{mnemonic}
	request, err := json.Marshal(strtt)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic_derive_sign_keys", string(request)
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey_xprv_from_mnemonic
func HdkeyXprvFromMnemonic(mSR *goton.MnemonicStructRequest) (string, string) {
	lenMnemonic := len(strings.Fields(mSR.Phrase))
	if _, ok := goton.LensMnemonic[lenMnemonic]; !ok {
		return "", ""
	}

	request, err := json.Marshal(mSR)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_xprv_from_mnemonic", string(request)
}

// HdkeyXprvSecret method crypto.hdkey_secret_from_xprv
func HdkeyXprvSecret(hdS *goton.HDSerialized) (string, string) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_secret_from_xprv", string(request)
}

// HdkeyXprvPublic method crypto.hdkey_public_from_xprv
func HdkeyXprvPublic(hdS *goton.HDSerialized) (string, string) {
	request, err := json.Marshal(hdS)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_public_from_xprv", string(request)
}

// HdkeyXprvDerive method crypto.hdkey_derive_from_xprv
func HdkeyXprvDerive(hdP *goton.HDDerivery) (string, string) {
	request, err := json.Marshal(hdP)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_derive_from_xprv", string(request)
}

// HdkeyXprvDerivePath method crypto.hdkey_derive_from_xprv_path
func HdkeyXprvDerivePath(hdPD *goton.HDPathDerivery) (string, string) {
	request, err := json.Marshal(hdPD)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_derive_from_xprv_path", string(request)
}
