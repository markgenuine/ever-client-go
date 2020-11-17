package crypto

import (
	"github.com/markgenuine/ton-client-go/domain"
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
func (c *crypto) Factorize(poF domain.ParamsOfFactorize) (int, error) {
	return c.client.Request("crypto.factorize", poF)
}

// ModularPower method crypto.modular_power
func (c *crypto) ModularPower(pOMP domain.ParamsOfModularPower) (int, error) {
	return c.client.Request("crypto.modular_power", pOMP)
}

// TonCrc16 method crypto.ton_crc16
func (c *crypto) TonCrc16(pOTC domain.ParamsOfTonCrc16) (int, error) {

	return c.client.Request("crypto.ton_crc16", pOTC)
}

// GenerateRandomBytes method crypto.generate_random_bytes
func (c *crypto) GenerateRandomBytes(pOGRB domain.ParamsOfGenerateRandomBytes) (int, error) {

	return c.client.Request("crypto.generate_random_bytes", pOGRB)
}

// ConvertPublicKeyString method crypto.convert_public_key_to_ton_safe_format
func (c *crypto) ConvertPublicKeyString(pOCPTTSF domain.ParamsOfConvertPublicKeyToTonSafeFormat) (int, error) {

	return c.client.Request("crypto.convert_public_key_to_ton_safe_format", pOCPTTSF)
}

// GenerateRandomSignKeys method crypto.generate_random_sign_keys
func (c *crypto) GenerateRandomSignKeys() (int, error) {
	return c.client.Request("crypto.generate_random_sign_keys", "{}")
}

// Sign method crypto.sign
func (c *crypto) Sign(pOS domain.ParamsOfSign) (int, error) {

	return c.client.Request("crypto.sign", pOS)
}

// VerifySignature method crypto.verify_signature
func (c *crypto) VerifySignature(pOVS domain.ParamsOfVerifySignature) (int, error) {

	return c.client.Request("crypto.verify_signature", pOVS)
}

// Sha256 method crypto.sha256
func (c *crypto) Sha256(pOH domain.ParamsOfHash) (int, error) {

	return c.client.Request("crypto.sha256", pOH)
}

// Sha512 method crypto.sha512
func (c *crypto) Sha512(pOH domain.ParamsOfHash) (int, error) {

	return c.client.Request("crypto.sha512", pOH)
}

// Scrypt method crypto.scrypt
func (c *crypto) Scrypt(sD domain.ParamsOfScrypt) (int, error) {

	return c.client.Request("crypto.scrypt", sD)
}

// NaclSignKeypairFromSecretKey method crypto.nacl_sign_keypair_from_secret_key
func (c *crypto) NaclSignKeypairFromSecretKey(pONSKPFC domain.ParamsOfNaclSignKeyPairFromSecret) (int, error) {

	return c.client.Request("crypto.nacl_sign_keypair_from_secret_key", pONSKPFC)
}

// NaclSign method crypto.nacl_sign
func (c *crypto) NaclSign(pONS domain.ParamsOfNaclSign) (int, error) {

	return c.client.Request("crypto.nacl_sign", pONS)
}

// NaclSignOpen method crypto.nacl_sign_open
func (c *crypto) NaclSignOpen(pONSO domain.ParamsOfNaclSignOpen) (int, error) {

	return c.client.Request("crypto.nacl_sign_open", pONSO)
}

// NaclSignDetached method crypto.nacl_sign_detached
func (c *crypto) NaclSignDetached(pONS domain.ParamsOfNaclSign) (int, error) {

	return c.client.Request("crypto.nacl_sign_detached", pONS)
}

// NaclBoxKeypair method crypto.nacl_box_keypair
func (c *crypto) NaclBoxKeypair() (int, error) {

	return c.client.Request("crypto.nacl_box_keypair", "{}")
}

// NaclBoxKeypairFromSecretKey method crypto.nacl_box_keypair_from_secret_key
func (c *crypto) NaclBoxKeypairFromSecretKey(pONKPFS domain.ParamsOfNaclBoxKeyPairFromSecret) (int, error) {

	return c.client.Request("crypto.nacl_box_keypair_from_secret_key", pONKPFS)
}

// NaclBox method crypto.nacl_box
func (c *crypto) NaclBox(pONB domain.ParamsOfNaclBox) (int, error) {

	return c.client.Request("crypto.nacl_box", pONB)
}

// NaclBoxOpen method crypto.nacl_box_open
func (c *crypto) NaclBoxOpen(pONBO domain.ParamsOfNaclBoxOpen) (int, error) {

	return c.client.Request("crypto.nacl_box_open", pONBO)
}

// NaclSecretBox method crypto.nacl_secret_box
func (c *crypto) NaclSecretBox(pONSB domain.ParamsOfNaclSecretBox) (int, error) {

	return c.client.Request("crypto.nacl_secret_box", pONSB)
}

// NaclSecretBoxOpen method crypto.nacl_secret_box_open
func (c *crypto) NaclSecretBoxOpen(pONSBO domain.ParamsOfNaclSecretBoxOpen) (int, error) {
	return c.client.Request("crypto.nacl_secret_box_open", pONSBO)
}

// Mnemonic

// MnemonicWords method crypto.mnemonic_words
func (c *crypto) MnemonicWords(pOMW domain.ParamsOfMnemonicWords) (int, error) {
	return c.client.Request("crypto.mnemonic_words", pOMW)
}

// MnemonicFromRandom method crypto.mnemonic_from_random
func (c *crypto) MnemonicFromRandom(pOMFR domain.ParamsOfMnemonicFromRandom) (int, error) {
	return c.client.Request("crypto.mnemonic_from_random", pOMFR)
}

// MnemonicFromEntropy method crypto.mnemonic_from_entropy
func (c *crypto) MnemonicFromEntropy(pOMFE domain.ParamsOfMnemonicFromEntropy) (int, error) {
	return c.client.Request("crypto.mnemonic_from_entropy", pOMFE)
}

// MnemonicVerify method crypto.mnemonic_verify
func (c *crypto) MnemonicVerify(pOMV domain.ParamsOfMnemonicVerify) (int, error) {
	return c.client.Request("crypto.mnemonic_verify", pOMV)
}

// MnemonicDeriveSignKeys method crypto.mnemonic_derive_sign_keys
func (c *crypto) MnemonicDeriveSignKeys(pOMDSK domain.ParamsOfMnemonicDeriveSignKeys) (int, error) {
	return c.client.Request("crypto.mnemonic_derive_sign_keys", pOMDSK)
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey_xprv_from_mnemonic
func (c *crypto) HdkeyXprvFromMnemonic(pOHKXFM domain.ParamsOfHDKeyXPrvFromMnemonic) (int, error) {
	return c.client.Request("crypto.hdkey_xprv_from_mnemonic", pOHKXFM)
}

// HdkeyXprvDerive method crypto.hdkey_derive_from_xprv
func (c *crypto) HdkeyXprvDerive(hdP domain.ParamsOfHDKeyDeriveFromXPrv) (int, error) {
	return c.client.Request("crypto.hdkey_derive_from_xprv", hdP)
}

// HdkeyXprvDerivePath method crypto.hdkey_derive_from_xprv_path
func (c *crypto) HdkeyXprvDerivePath(hdPD domain.ParamsOfHDKeyDeriveFromXPrvPath) (int, error) {
	return c.client.Request("crypto.hdkey_derive_from_xprv_path", hdPD)
}

// HdkeyXprvSecret method crypto.hdkey_secret_from_xprv
func (c *crypto) HdkeyXprvSecret(pOHKSFXP domain.ParamsOfHDKeySecretFromXPrv) (int, error) {
	return c.client.Request("crypto.hdkey_secret_from_xprv", pOHKSFXP)
}

// HdkeyXprvPublic method crypto.hdkey_public_from_xprv
func (c *crypto) HdkeyXprvPublic(pOHKPFXP domain.ParamsOfHDKeyPublicFromXPrv) (int, error) {
	return c.client.Request("crypto.hdkey_public_from_xprv", pOHKPFXP)
}
