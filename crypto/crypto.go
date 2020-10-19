package crypto

import (
	"encoding/json"
)

// Factorize method crypto.factorize
func Factorize(poF *ParamsOfFactorize) (string, string) {
	request, err := json.Marshal(poF)
	if err != nil {
		return "", ""
	}

	return "crypto.factorize", string(request)
}

// FactorizeResult ...
func FactorizeResult(resp string, err error) (*ResultOfFactorize, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfFactorize{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
}

// ModularPower method crypto.modular_power
func ModularPower(pOMP *ParamsOfModularPower) (string, string) {
	request, err := json.Marshal(pOMP)
	if err != nil {
		return "", ""
	}
	return "crypto.modular_power", string(request)
}

// ModularPowerResult ...
func ModularPowerResult(resp string, err error) (*ResultOfModularPower, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfModularPower{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
}

// TonCrc16 method crypto.ton_crc16
func TonCrc16(pOTC *ParamsOfTonCrc16) (string, string) {
	request, err := json.Marshal(pOTC)
	if err != nil {
		return "", ""
	}
	return "crypto.ton_crc16", string(request)
}

// TonCrc16Result ...
func TonCrc16Result(resp string, err error) (*ResultOfTonCrc16, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfTonCrc16{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
}

// GenerateRandomBytes method crypto.generate_random_bytes
func GenerateRandomBytes(pOGRB *ParamsOfGenerateRandomBytes) (string, string) {
	request, err := json.Marshal(pOGRB)
	if err != nil {
		return "", ""
	}
	return "crypto.generate_random_bytes", string(request)
}

// GenerateRandomBytesResult ...
func GenerateRandomBytesResult(resp string, err error) (*ResultOfGenerateRandomBytes, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfGenerateRandomBytes{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
}

//TonPublicKeyString method crypto.convert_public_key_to_ton_safe_format
func TonPublicKeyString(pOCPTTSF *ParamsOfConvertPublicKeyToTonSafeFormat) (string, string) {
	request, err := json.Marshal(pOCPTTSF)
	if err != nil {
		return "", ""
	}
	return "crypto.convert_public_key_to_ton_safe_format", string(request)
}

// TonPublicKeyStringResult ...
func TonPublicKeyStringResult(resp string, err error) (*ResultOfConvertPublicKeyToTonSafeFormat, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfConvertPublicKeyToTonSafeFormat{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
}

// GenerateRandomSignKeys call "crypto.generate_random_sign_keys"
func GenerateRandomSignKeys() (string, string) {
	return "crypto.generate_random_sign_keys", "{}"
}

// GenerateRandomSignKeysResult response method crypto.generate_random_sign_keys
func GenerateRandomSignKeysResult(resp string, err error) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	result := &KeyPair{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Sign call "crypto.sign"
func Sign(pOS *ParamsOfSign) (string, string) {
	request, err := json.Marshal(pOS)
	if err != nil {
		return "", ""
	}
	return "crypto.sign", string(request)
}

// SignResult response method crypto.sign
func SignResult(resp string, err error) (*ResultOfSign, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfSign{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// VerifySignature call "crypto.verify_signature"
func VerifySignature(pOVS *ParamsOfVerifySignature) (string, string) {
	request, err := json.Marshal(pOVS)
	if err != nil {
		return "", ""
	}
	return "crypto.verify_signature", string(request)
}

// VerifySignatureResult response method crypto.verify_signature
func VerifySignatureResult(resp string, err error) (*ResultOfVerifySignature, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfVerifySignature{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Sha256 method crypto.sha256
func Sha256(pOH *ParamsOfHash) (string, string) {
	request, err := json.Marshal(pOH)
	if err != nil {
		return "", ""
	}
	return "crypto.sha256", string(request)
}

// Sha512 method crypto.sha512
func Sha512(pOH *ParamsOfHash) (string, string) {
	request, err := json.Marshal(pOH)
	if err != nil {
		return "", ""
	}
	return "crypto.sha512", string(request)
}

// ShaResult result for sha256 and sha512
func ShaResult(resp string, err error) (*ResultOfHash, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHash{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Scrypt method crypto.scrypt
func Scrypt(sD *ParamsOfScrypt) (string, string) {
	request, err := json.Marshal(sD)
	if err != nil {
		return "", ""
	}

	return "crypto.scrypt", string(request)
}

// ScryptResult result for crypto.scrypt
func ScryptResult(resp string, err error) (*ResultOfScrypt, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfScrypt{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSignKeypairFromSecretKey method crypto.nacl_sign_keypair_from_secret_key
func NaclSignKeypairFromSecretKey(pONSKPFC *ParamsOfNaclSignKeyPairFromSecret) (string, string) {
	request, err := json.Marshal(pONSKPFC)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_sign_keypair_from_secret_key", string(request)
}

// NaclSignKeypairFromSecretKeyResult result for crypto.nacl_sign_keypair_from_secret_key
func NaclSignKeypairFromSecretKeyResult(resp string, err error) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	result := &KeyPair{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSign method crypto.nacl_sign
func NaclSign(pONS *ParamsOfNaclSign) (string, string) {
	request, err := json.Marshal(pONS)
	if err != nil {
		return "", ""
	}

	return "crypto.nacl_sign", string(request)
}

// NaclSignResult result for crypto.nacl_sign
func NaclSignResult(resp string, err error) (*ResultOfNaclSign, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclSign{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSignOpen method crypto.nacl_sign_open
func NaclSignOpen(pONSO *ParamsOfNaclSignOpen) (string, string) {
	request, err := json.Marshal(pONSO)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_sign_open", string(request)
}

// NaclSignOpenResult result for crypto.nacl_sign_open
func NaclSignOpenResult(resp string, err error) (*ResultOfNaclSignOpen, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclSignOpen{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSignDetached method crypto.nacl_sign_detached
func NaclSignDetached(pONS *ParamsOfNaclSign) (string, string) {
	request, err := json.Marshal(pONS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_sign_detached", string(request)
}

// NaclSignDetachedResult result for crypto.nacl_sign_open
func NaclSignDetachedResult(resp string, err error) (*ResultOfNaclSignDetached, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclSignDetached{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBoxKeypair method crypto.nacl_box_keypair
func NaclBoxKeypair() (string, string) {
	return "crypto.nacl_box_keypair", "{}"
}

// NaclBoxKeypairResult result for crypto.nacl_box_keypair
func NaclBoxKeypairResult(resp string, err error) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	result := &KeyPair{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBoxKeypairFromSecretKey method crypto.nacl_box_keypair_from_secret_key
func NaclBoxKeypairFromSecretKey(pONKPFS *ParamsOfNaclBoxKeyPairFromSecret) (string, string) {
	request, err := json.Marshal(pONKPFS)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_box_keypair_from_secret_key", string(request)
}

// NaclBoxKeypairFromSecretKeyResult result for crypto.nacl_box_keypair_from_secret_key
func NaclBoxKeypairFromSecretKeyResult(resp string, err error) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	result := &KeyPair{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBox method crypto.nacl_box
func NaclBox(pONB *ParamsOfNaclBox) (string, string) {
	request, err := json.Marshal(pONB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_box", string(request)
}

// NaclBoxResult result for crypto.nacl_box
func NaclBoxResult(resp string, err error) (*ResultOfNaclBox, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclBox{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclBoxOpen method crypto.nacl_box_open
func NaclBoxOpen(pONBO *ParamsOfNaclBoxOpen) (string, string) {
	request, err := json.Marshal(pONBO)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_box_open", string(request)
}

// NaclBoxOpenResult result for crypto.nacl_box_open
func NaclBoxOpenResult(resp string, err error) (*ResultOfNaclBoxOpen, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclBoxOpen{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSecretBox method crypto.nacl_secret_box
func NaclSecretBox(pONSB ParamsOfNaclSecretBox) (string, string) {
	request, err := json.Marshal(pONSB)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_secret_box", string(request)
}

// NaclSecretBoxResult result for crypto.nacl_secret_box
func NaclSecretBoxResult(resp string, err error) (*ResultOfNaclBox, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclBox{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NaclSecretBoxOpen method crypto.nacl_secret_box_open
func NaclSecretBoxOpen(pONSBO *ParamsOfNaclSecretBoxOpen) (string, string) {
	request, err := json.Marshal(pONSBO)
	if err != nil {
		return "", ""
	}
	return "crypto.nacl_secret_box_open", string(request)
}

// NaclSecretBoxOpenResult result for crypto.nacl_box_open
func NaclSecretBoxOpenResult(resp string, err error) (*ResultOfNaclBoxOpen, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfNaclBoxOpen{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Mnemonic

// MnemonicWords method crypto.mnemonic_words
func MnemonicWords(pOMW *ParamsOfMnemonicWords) (string, string) {
	request, err := json.Marshal(pOMW)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic_words", string(request)
}

// MnemonicWordsResult method crypto.mnemonic_words
func MnemonicWordsResult(resp string, err error) (*ResultOfMnemonicWords, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfMnemonicWords{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MnemonicFromRandom method crypto.mnemonic_from_random
func MnemonicFromRandom(pOMFR *ParamsOfMnemonicFromRandom) (string, string) {
	request, err := json.Marshal(pOMFR)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic_from_random", string(request)
}

// MnemonicFromRandomResult method crypto.mnemonic_words
func MnemonicFromRandomResult(resp string, err error) (*ResultOfMnemonicFromRandom, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfMnemonicFromRandom{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MnemonicFromEntropy method crypto.mnemonic_from_entropy
func MnemonicFromEntropy(pOMFE *ParamsOfMnemonicFromEntropy) (string, string) {
	request, err := json.Marshal(pOMFE)
	if err != nil {
		return "", ""
	}
	return "crypto.mnemonic_from_entropy", string(request)
}

// MnemonicFromEntropyResult method crypto.mnemonic_words
func MnemonicFromEntropyResult(resp string, err error) (*ResultOfMnemonicFromEntropy, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfMnemonicFromEntropy{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MnemonicVerify method crypto.mnemonic_verify
func MnemonicVerify(pOMV *ParamsOfMnemonicVerify) (string, string) {
	request, err := json.Marshal(pOMV)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic_verify", string(request)
}

// MnemonicVerifyResult method crypto.mnemonic_verify
func MnemonicVerifyResult(resp string, err error) (*ResultOfMnemonicVerify, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfMnemonicVerify{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MnemonicDeriveSignKeys method crypto.mnemonic_derive_sign_keys
func MnemonicDeriveSignKeys(pOMDSK *ParamsOfMnemonicDeriveSignKeys) (string, string) {
	request, err := json.Marshal(pOMDSK)
	if err != nil {
		return "", ""
	}

	return "crypto.mnemonic_derive_sign_keys", string(request)
}

// MnemonicDeriveSignKeysResult method crypto.mnemonic_derive_sign_keys
func MnemonicDeriveSignKeysResult(resp string, err error) (*KeyPair, error) {
	if err != nil {
		return nil, err
	}

	result := &KeyPair{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey_xprv_from_mnemonic
func HdkeyXprvFromMnemonic(pOHKXFM *ParamsOfHDKeyXPrvFromMnemonic) (string, string) {
	request, err := json.Marshal(pOHKXFM)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_xprv_from_mnemonic", string(request)
}

// HdkeyXprvFromMnemonicResult method crypto.hdkey_xprv_from_mnemonic
func HdkeyXprvFromMnemonicResult(resp string, err error) (*ResultOfHDKeyXPrvFromMnemonic, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHDKeyXPrvFromMnemonic{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HdkeyXprvDerive method crypto.hdkey_derive_from_xprv
func HdkeyXprvDerive(hdP *ParamsOfHDKeyDeriveFromXPrv) (string, string) {
	request, err := json.Marshal(hdP)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_derive_from_xprv", string(request)
}

// HdkeyXprvDeriveResult method crypto.hdkey_derive_from_xprv
func HdkeyXprvDeriveResult(resp string, err error) (*ResultOfHDKeyDeriveFromXPrv, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHDKeyDeriveFromXPrv{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HdkeyXprvDerivePath method crypto.hdkey_derive_from_xprv_path
func HdkeyXprvDerivePath(hdPD *ParamsOfHDKeyDeriveFromXPrvPath) (string, string) {
	request, err := json.Marshal(hdPD)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_derive_from_xprv_path", string(request)
}

// HdkeyXprvDerivePathResult method crypto.hdkey_derive_from_xprv_path
func HdkeyXprvDerivePathResult(resp string, err error) (*ResultOfHDKeyDeriveFromXPrvPath, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHDKeyDeriveFromXPrvPath{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HdkeyXprvSecret method crypto.hdkey_secret_from_xprv
func HdkeyXprvSecret(pOHKSFXP *ParamsOfHDKeySecretFromXPrv) (string, string) {
	request, err := json.Marshal(pOHKSFXP)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_secret_from_xprv", string(request)
}

// HdkeyXprvSecretResult method crypto.hdkey_secret_from_xprv
func HdkeyXprvSecretResult(resp string, err error) (*ResultOfHDKeySecretFromXPrv, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHDKeySecretFromXPrv{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HdkeyXprvPublic method crypto.hdkey_public_from_xprv
func HdkeyXprvPublic(pOHKPFXP *ParamsOfHDKeyPublicFromXPrv) (string, string) {
	request, err := json.Marshal(pOHKPFXP)
	if err != nil {
		return "", ""
	}
	return "crypto.hdkey_public_from_xprv", string(request)
}

// HdkeyXprvPublicResult method crypto.hdkey_secret_from_xprv
func HdkeyXprvPublicResult(resp string, err error) (*ResultOfHDKeyPublicFromXPrv, error) {
	if err != nil {
		return nil, err
	}

	result := &ResultOfHDKeyPublicFromXPrv{}
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
