package goton

// Factorize method crypto.factorize
func (client *Client) Factorize(poF ParamsOfFactorize) int {
	return client.Request("crypto.factorize", structToJSON(poF))
}

// ModularPower method crypto.modular_power
func (client *Client) ModularPower(pOMP ParamsOfModularPower) int {
	return client.Request("crypto.modular_power", structToJSON(pOMP))
}

// TonCrc16 method crypto.ton_crc16
func (client *Client) TonCrc16(pOTC ParamsOfTonCrc16) int {
	return client.Request("crypto.ton_crc16", structToJSON(pOTC))
}

// GenerateRandomBytes method crypto.generate_random_bytes
func (client *Client) GenerateRandomBytes(pOGRB ParamsOfGenerateRandomBytes) int {
	return client.Request("crypto.generate_random_bytes", structToJSON(pOGRB))
}

// ConvertPublicKeyString method crypto.convert_public_key_to_ton_safe_format
func (client *Client) ConvertPublicKeyString(pOCPTTSF ParamsOfConvertPublicKeyToTonSafeFormat) int {
	return client.Request("crypto.convert_public_key_to_ton_safe_format", structToJSON(pOCPTTSF))
}

// GenerateRandomSignKeys method crypto.generate_random_sign_keys
func (client *Client) GenerateRandomSignKeys() int {
	return client.Request("crypto.generate_random_sign_keys", "{}")
}

// Sign method crypto.sign
func (client *Client) Sign(pOS ParamsOfSign) int {
	return client.Request("crypto.sign", structToJSON(pOS))
}

// VerifySignature method crypto.verify_signature
func (client *Client) VerifySignature(pOVS ParamsOfVerifySignature) int {
	return client.Request("crypto.verify_signature", structToJSON(pOVS))
}

// Sha256 method crypto.sha256
func (client *Client) Sha256(pOH ParamsOfHash) int {
	return client.Request("crypto.sha256", structToJSON(pOH))
}

// Sha512 method crypto.sha512
func (client *Client) Sha512(pOH ParamsOfHash) int {
	return client.Request("crypto.sha512", structToJSON(pOH))
}

// Scrypt method crypto.scrypt
func (client *Client) Scrypt(sD ParamsOfScrypt) int {
	return client.Request("crypto.scrypt", structToJSON(sD))
}

// NaclSignKeypairFromSecretKey method crypto.nacl_sign_keypair_from_secret_key
func (client *Client) NaclSignKeypairFromSecretKey(pONSKPFC ParamsOfNaclSignKeyPairFromSecret) int {
	return client.Request("crypto.nacl_sign_keypair_from_secret_key", structToJSON(pONSKPFC))
}

// NaclSign method crypto.nacl_sign
func (client *Client) NaclSign(pONS ParamsOfNaclSign) int {
	return client.Request("crypto.nacl_sign", structToJSON(pONS))
}

// NaclSignOpen method crypto.nacl_sign_open
func (client *Client) NaclSignOpen(pONSO ParamsOfNaclSignOpen) int {
	return client.Request("crypto.nacl_sign_open", structToJSON(pONSO))
}

// NaclSignDetached method crypto.nacl_sign_detached
func (client *Client) NaclSignDetached(pONS ParamsOfNaclSign) int {
	return client.Request("crypto.nacl_sign_detached", structToJSON(pONS))
}

// NaclBoxKeypair method crypto.nacl_box_keypair
func (client *Client) NaclBoxKeypair() int {
	return client.Request("crypto.nacl_box_keypair", "{}")
}

// NaclBoxKeypairFromSecretKey method crypto.nacl_box_keypair_from_secret_key
func (client *Client) NaclBoxKeypairFromSecretKey(pONKPFS ParamsOfNaclBoxKeyPairFromSecret) int {
	return client.Request("crypto.nacl_box_keypair_from_secret_key", structToJSON(pONKPFS))
}

// NaclBox method crypto.nacl_box
func (client *Client) NaclBox(pONB ParamsOfNaclBox) int {
	return client.Request("crypto.nacl_box", structToJSON(pONB))
}

// NaclBoxOpen method crypto.nacl_box_open
func (client *Client) NaclBoxOpen(pONBO ParamsOfNaclBoxOpen) int {
	return client.Request("crypto.nacl_box_open", structToJSON(pONBO))
}

// NaclSecretBox method crypto.nacl_secret_box
func (client *Client) NaclSecretBox(pONSB ParamsOfNaclSecretBox) int {
	return client.Request("crypto.nacl_secret_box", structToJSON(pONSB))
}

// NaclSecretBoxOpen method crypto.nacl_secret_box_open
func (client *Client) NaclSecretBoxOpen(pONSBO ParamsOfNaclSecretBoxOpen) int {
	return client.Request("crypto.nacl_secret_box_open", structToJSON(pONSBO))
}

// Mnemonic

// MnemonicWords method crypto.mnemonic_words
func (client *Client) MnemonicWords(pOMW ParamsOfMnemonicWords) int {
	return client.Request("crypto.mnemonic_words", structToJSON(pOMW))
}

// MnemonicFromRandom method crypto.mnemonic_from_random
func (client *Client) MnemonicFromRandom(pOMFR ParamsOfMnemonicFromRandom) int {
	return client.Request("crypto.mnemonic_from_random", structToJSON(pOMFR))
}

// MnemonicFromEntropy method crypto.mnemonic_from_entropy
func (client *Client) MnemonicFromEntropy(pOMFE ParamsOfMnemonicFromEntropy) int {
	return client.Request("crypto.mnemonic_from_entropy", structToJSON(pOMFE))
}

// MnemonicVerify method crypto.mnemonic_verify
func (client *Client) MnemonicVerify(pOMV ParamsOfMnemonicVerify) int {
	return client.Request("crypto.mnemonic_verify", structToJSON(pOMV))
}

// MnemonicDeriveSignKeys method crypto.mnemonic_derive_sign_keys
func (client *Client) MnemonicDeriveSignKeys(pOMDSK ParamsOfMnemonicDeriveSignKeys) int {
	return client.Request("crypto.mnemonic_derive_sign_keys", structToJSON(pOMDSK))
}

// HDKeys

// HdkeyXprvFromMnemonic method crypto.hdkey_xprv_from_mnemonic
func (client *Client) HdkeyXprvFromMnemonic(pOHKXFM ParamsOfHDKeyXPrvFromMnemonic) int {
	return client.Request("crypto.hdkey_xprv_from_mnemonic", structToJSON(pOHKXFM))
}

// HdkeyXprvDerive method crypto.hdkey_derive_from_xprv
func (client *Client) HdkeyXprvDerive(hdP ParamsOfHDKeyDeriveFromXPrv) int {
	return client.Request("crypto.hdkey_derive_from_xprv", structToJSON(hdP))
}

// HdkeyXprvDerivePath method crypto.hdkey_derive_from_xprv_path
func (client *Client) HdkeyXprvDerivePath(hdPD ParamsOfHDKeyDeriveFromXPrvPath) int {
	return client.Request("crypto.hdkey_derive_from_xprv_path", structToJSON(hdPD))
}

// HdkeyXprvSecret method crypto.hdkey_secret_from_xprv
func (client *Client) HdkeyXprvSecret(pOHKSFXP ParamsOfHDKeySecretFromXPrv) int {
	return client.Request("crypto.hdkey_secret_from_xprv", structToJSON(pOHKSFXP))
}

// HdkeyXprvPublic method crypto.hdkey_public_from_xprv
func (client *Client) HdkeyXprvPublic(pOHKPFXP ParamsOfHDKeyPublicFromXPrv) int {
	return client.Request("crypto.hdkey_public_from_xprv", structToJSON(pOHKPFXP))
}
