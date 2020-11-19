package crypto

import (
	"github.com/move-ton/ton-client-go/domain"
)

type crypto struct {
	config domain.Config
	client domain.ClientGateway
}

// NewCrypto ...
func NewCrypto(config domain.Config, client domain.ClientGateway) domain.CryptoUseCase {
	return &crypto{
		config: config,
		client: client,
	}
}

// Factorize method crypto.factorize
func (c *crypto) Factorize(poF domain.ParamsOfFactorize) (*domain.ResultOfFactorize, error) {
	result := new(domain.ResultOfFactorize)
	err := c.client.GetResult("crypto.factorize", poF, result)
	return result, err
}

// ModularPower method crypto.modular_power
func (c *crypto) ModularPower(pOMP domain.ParamsOfModularPower) (*domain.ResultOfModularPower, error) {
	result := new(domain.ResultOfModularPower)
	err := c.client.GetResult("crypto.modular_power", pOMP, result)
	return result, err
}

// TonCrc16 method crypto.ton_crc16
func (c *crypto) TonCrc16(pOTC domain.ParamsOfTonCrc16) (*domain.ResultOfTonCrc16, error) {
	result := new(domain.ResultOfTonCrc16)
	err := c.client.GetResult("crypto.ton_crc16", pOTC, result)
	return result, err
}

// GenerateRandomBytes method crypto.generate_random_bytes
func (c *crypto) GenerateRandomBytes(pOGRB domain.ParamsOfGenerateRandomBytes) (*domain.ResultOfGenerateRandomBytes, error) {
	result := new(domain.ResultOfGenerateRandomBytes)
	err := c.client.GetResult("crypto.generate_random_bytes", pOGRB, result)
	return result, err
}

// ConvertPublicKeyString method crypto.convert_public_key_to_ton_safe_format
func (c *crypto) ConvertPublicKeyString(pOCPTTSF domain.ParamsOfConvertPublicKeyToTonSafeFormat) (*domain.ResultOfConvertPublicKeyToTonSafeFormat, error) {
	result := new(domain.ResultOfConvertPublicKeyToTonSafeFormat)
	err := c.client.GetResult("crypto.convert_public_key_to_ton_safe_format", pOCPTTSF, result)
	return result, err
}

// GenerateRandomSignKeys method crypto.generate_random_sign_keys
func (c *crypto) GenerateRandomSignKeys() (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.generate_random_sign_keys", "{}", result)
	return result, err
}

// Sign method crypto.sign
func (c *crypto) Sign(pOS domain.ParamsOfSign) (*domain.ResultOfSign, error) {
	result := new(domain.ResultOfSign)
	err := c.client.GetResult("crypto.sign", pOS, result)
	return result, err
}

// VerifySignature method crypto.verify_signature
func (c *crypto) VerifySignature(pOVS domain.ParamsOfVerifySignature) (*domain.ResultOfVerifySignature, error) {
	result := new(domain.ResultOfVerifySignature)
	err := c.client.GetResult("crypto.verify_signature", pOVS, result)
	return result, err
}

// Sha256 method crypto.sha256
func (c *crypto) Sha256(pOH domain.ParamsOfHash) (*domain.ResultOfHash, error) {
	result := new(domain.ResultOfHash)
	err := c.client.GetResult("crypto.sha256", pOH, result)
	return result, err
}

// Sha512 method crypto.sha512
func (c *crypto) Sha512(pOH domain.ParamsOfHash) (*domain.ResultOfHash, error) {
	result := new(domain.ResultOfHash)
	err := c.client.GetResult("crypto.sha512", pOH, result)
	return result, err
}

// Scrypt method crypto.scrypt
func (c *crypto) Scrypt(sD domain.ParamsOfScrypt) (*domain.ResultOfScrypt, error) {
	result := new(domain.ResultOfScrypt)
	err := c.client.GetResult("crypto.scrypt", sD, result)
	return result, err
}

// NaclSignKeypairFromSecretKey method crypto.nacl_sign_keypair_from_secret_key
func (c *crypto) NaclSignKeypairFromSecretKey(pONSKPFC domain.ParamsOfNaclSignKeyPairFromSecret) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_sign_keypair_from_secret_key", pONSKPFC, result)
	return result, err
}

// NaclSign method crypto.nacl_sign
func (c *crypto) NaclSign(pONS domain.ParamsOfNaclSign) (*domain.ResultOfNaclSign, error) {
	result := new(domain.ResultOfNaclSign)
	err := c.client.GetResult("crypto.nacl_sign", pONS, result)
	return result, err
}

// NaclSignOpen method crypto.nacl_sign_open
func (c *crypto) NaclSignOpen(pONSO domain.ParamsOfNaclSignOpen) (*domain.ResultOfNaclSignOpen, error) {
	result := new(domain.ResultOfNaclSignOpen)
	err := c.client.GetResult("crypto.nacl_sign_open", pONSO, result)
	return result, err
}

// NaclSignDetached method crypto.nacl_sign_detached
func (c *crypto) NaclSignDetached(pONS domain.ParamsOfNaclSign) (*domain.ResultOfNaclSignDetached, error) {
	result := new(domain.ResultOfNaclSignDetached)
	err := c.client.GetResult("crypto.nacl_sign_detached", pONS, result)
	return result, err
}

// NaclBoxKeypair method crypto.nacl_box_keypair
func (c *crypto) NaclBoxKeypair() (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_box_keypair", "{}", result)
	return result, err
}

// NaclBoxKeypairFromSecretKey method crypto.nacl_box_keypair_from_secret_key
func (c *crypto) NaclBoxKeypairFromSecretKey(pONKPFS domain.ParamsOfNaclBoxKeyPairFromSecret) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_box_keypair_from_secret_key", pONKPFS, result)
	return result, err
}

// NaclBox method crypto.nacl_box
func (c *crypto) NaclBox(pONB domain.ParamsOfNaclBox) (*domain.ResultOfNaclBox, error) {
	result := new(domain.ResultOfNaclBox)
	err := c.client.GetResult("crypto.nacl_box", pONB, result)
	return result, err
}

// NaclBoxOpen method crypto.nacl_box_open
func (c *crypto) NaclBoxOpen(pONBO domain.ParamsOfNaclBoxOpen) (*domain.ResultOfNaclBoxOpen, error) {
	result := new(domain.ResultOfNaclBoxOpen)
	err := c.client.GetResult("crypto.nacl_box_open", pONBO, result)
	return result, err
}

// NaclSecretBox method crypto.nacl_secret_box
func (c *crypto) NaclSecretBox(pONSB domain.ParamsOfNaclSecretBox) (*domain.ResultOfNaclBox, error) {
	result := new(domain.ResultOfNaclBox)
	err := c.client.GetResult("crypto.nacl_secret_box", pONSB, result)
	return result, err
}

// NaclSecretBoxOpen method crypto.nacl_secret_box_open
func (c *crypto) NaclSecretBoxOpen(pONSBO domain.ParamsOfNaclSecretBoxOpen) (*domain.ResultOfNaclBoxOpen, error) {
	result := new(domain.ResultOfNaclBoxOpen)
	err := c.client.GetResult("crypto.nacl_secret_box_open", pONSBO, result)
	return result, err
}

// Mnemonic

// MnemonicWords method crypto.mnemonic_words
func (c *crypto) MnemonicWords(pOMW domain.ParamsOfMnemonicWords) (*domain.ResultOfMnemonicWords, error) {
	result := new(domain.ResultOfMnemonicWords)
	err := c.client.GetResult("crypto.mnemonic_words", pOMW, result)
	return result, err
}

// MnemonicFromRandom method crypto.mnemonic_from_random
func (c *crypto) MnemonicFromRandom(pOMFR domain.ParamsOfMnemonicFromRandom) (*domain.ResultOfMnemonicFromRandom, error) {
	result := new(domain.ResultOfMnemonicFromRandom)
	err := c.client.GetResult("crypto.mnemonic_from_random", pOMFR, result)
	return result, err
}

// MnemonicFromEntropy method crypto.mnemonic_from_entropy
func (c *crypto) MnemonicFromEntropy(pOMFE domain.ParamsOfMnemonicFromEntropy) (*domain.ResultOfMnemonicFromEntropy, error) {
	result := new(domain.ResultOfMnemonicFromEntropy)
	err := c.client.GetResult("crypto.mnemonic_from_entropy", pOMFE, result)
	return result, err
}

// MnemonicVerify method crypto.mnemonic_verify
func (c *crypto) MnemonicVerify(pOMV domain.ParamsOfMnemonicVerify) (*domain.ResultOfMnemonicVerify, error) {
	result := new(domain.ResultOfMnemonicVerify)
	err := c.client.GetResult("crypto.mnemonic_verify", pOMV, result)
	return result, err
}

// MnemonicDeriveSignKeys method crypto.mnemonic_derive_sign_keys
func (c *crypto) MnemonicDeriveSignKeys(pOMDSK domain.ParamsOfMnemonicDeriveSignKeys) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.mnemonic_derive_sign_keys", pOMDSK, result)
	return result, err
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey_xprv_from_mnemonic
func (c *crypto) HdkeyXprvFromMnemonic(pOHKXFM domain.ParamsOfHDKeyXPrvFromMnemonic) (*domain.ResultOfHDKeyXPrvFromMnemonic, error) {
	result := new(domain.ResultOfHDKeyXPrvFromMnemonic)
	err := c.client.GetResult("crypto.hdkey_xprv_from_mnemonic", pOHKXFM, result)
	return result, err
}

// HdkeyXprvDerive method crypto.hdkey_derive_from_xprv
func (c *crypto) HdkeyXprvDerive(hdP domain.ParamsOfHDKeyDeriveFromXPrv) (*domain.ResultOfHDKeyDeriveFromXPrv, error) {
	result := new(domain.ResultOfHDKeyDeriveFromXPrv)
	err := c.client.GetResult("crypto.hdkey_derive_from_xprv", hdP, result)
	return result, err
}

// HdkeyXprvDerivePath method crypto.hdkey_derive_from_xprv_path
func (c *crypto) HdkeyXprvDerivePath(hdPD domain.ParamsOfHDKeyDeriveFromXPrvPath) (*domain.ResultOfHDKeyDeriveFromXPrvPath, error) {
	result := new(domain.ResultOfHDKeyDeriveFromXPrvPath)
	err := c.client.GetResult("crypto.hdkey_derive_from_xprv_path", hdPD, result)
	return result, err
}

// HdkeyXprvSecret method crypto.hdkey_secret_from_xprv
func (c *crypto) HdkeyXprvSecret(pOHKSFXP domain.ParamsOfHDKeySecretFromXPrv) (*domain.ResultOfHDKeySecretFromXPrv, error) {
	result := new(domain.ResultOfHDKeySecretFromXPrv)
	err := c.client.GetResult("crypto.hdkey_secret_from_xprv", pOHKSFXP, result)
	return result, err
}

// HdkeyXprvPublic method crypto.hdkey_public_from_xprv
func (c *crypto) HdkeyXprvPublic(pOHKPFXP domain.ParamsOfHDKeyPublicFromXPrv) (*domain.ResultOfHDKeyPublicFromXPrv, error) {
	result := new(domain.ResultOfHDKeyPublicFromXPrv)
	err := c.client.GetResult("crypto.hdkey_public_from_xprv", pOHKPFXP, result)
	return result, err
}
