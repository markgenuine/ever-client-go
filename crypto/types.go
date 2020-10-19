package crypto

// ParamsOfFactorize ...
type ParamsOfFactorize struct {
	Composite string `json:"composite"`
}

// ResultOfFactorize ...
type ResultOfFactorize struct {
	Factors []string `json:"factors"`
}

// ParamsOfModularPower ...
type ParamsOfModularPower struct {
	Base     string `json:"base"`
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

// ResultOfModularPower ...
type ResultOfModularPower struct {
	ModularPower string `json:"modular_power"`
}

// ParamsOfTonCrc16 ...
type ParamsOfTonCrc16 struct {
	Data string `json:"data"`
}

// ResultOfTonCrc16 ...
type ResultOfTonCrc16 struct {
	Crc int `json:"crc"`
}

// ParamsOfGenerateRandomBytes ...
type ParamsOfGenerateRandomBytes struct {
	Length int `json:"length"`
}

// ResultOfGenerateRandomBytes ...
type ResultOfGenerateRandomBytes struct {
	Bytes string `json:"bytes"`
}

// ParamsOfConvertPublicKeyToTonSafeFormat ...
type ParamsOfConvertPublicKeyToTonSafeFormat struct {
	PublicKey string `json:"public_key"`
}

// ResultOfConvertPublicKeyToTonSafeFormat ...
type ResultOfConvertPublicKeyToTonSafeFormat struct {
	TonPublicKey string `json:"ton_public_key"`
}

// KeyPair ...
type KeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

// ParamsOfSign ...
type ParamsOfSign struct {
	Unsigned string
	Keys     *KeyPair
}

// ResultOfSign ...
type ResultOfSign struct {
	Signed    string
	Signature string
}

// ParamsOfVerifySignature ...
type ParamsOfVerifySignature struct {
	Signed string
	Public string
}

// ResultOfVerifySignature ...
type ResultOfVerifySignature struct {
	Unsigned string
}

// ParamsOfHash ...
type ParamsOfHash struct {
	Data string
}

// ResultOfHash ...
type ResultOfHash struct {
	Hash string
}

// ParamsOfScrypt ...
type ParamsOfScrypt struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
	LogN     int    `json:"log_n"`
	R        int    `json:"r"`
	P        int    `json:"p"`
	DkLen    int    `json:"dk_len"`
}

// ResultOfScrypt ...
type ResultOfScrypt struct {
	Key string `json:"key"`
}

// ParamsOfNaclSignKeyPairFromSecret ...
type ParamsOfNaclSignKeyPairFromSecret struct {
	Secret string
}

// ParamsOfNaclSign ...
type ParamsOfNaclSign struct {
	Unsigned string
	Secret   string
}

// ResultOfNaclSign ...
type ResultOfNaclSign struct {
	Signed string
}

// ParamsOfNaclSignOpen ...
type ParamsOfNaclSignOpen struct {
	Signed string
	Public string
}

// ResultOfNaclSignOpen ...
type ResultOfNaclSignOpen struct {
	Unsigned string
}

// ResultOfNaclSignDetached ...
type ResultOfNaclSignDetached struct {
	Signature string
}

// ParamsOfNaclBoxKeyPairFromSecret ...
type ParamsOfNaclBoxKeyPairFromSecret struct {
	Secret string
}

// ParamsOfNaclBox ...
type ParamsOfNaclBox struct {
	Decrypted   string
	Nonce       string
	TheirPublic string `json:"their_public"`
	Secret      string
}

// ResultOfNaclBox ...
type ResultOfNaclBox struct {
	Encrypted string
}

// ParamsOfNaclBoxOpen ...
type ParamsOfNaclBoxOpen struct {
	Encrypted   string
	Nonce       string
	TheirPublic string `json:"their_public"`
	Secret      string
}

// ResultOfNaclBoxOpen ...
type ResultOfNaclBoxOpen struct {
	Decrypted string
}

// ParamsOfNaclSecretBox ...
type ParamsOfNaclSecretBox struct {
	Decrypted string
	Nonce     string
	Key       string
}

// ParamsOfNaclSecretBoxOpen ...
type ParamsOfNaclSecretBoxOpen struct {
	Encrypted string
	Nonce     string
	Key       string
}

// ParamsOfMnemonicWords ...
type ParamsOfMnemonicWords struct {
	Dictionary int `json:"dictionary,omitempty"`
}

// ResultOfMnemonicWords ...
type ResultOfMnemonicWords struct {
	Words string `json:"words"`
}

// ParamsOfMnemonicFromRandom ...
type ParamsOfMnemonicFromRandom struct {
	Dictionary int `json:"dictionary,omitempty"`
	WordCount  int `json:"word_count,omitempty"`
}

// ResultOfMnemonicFromRandom ...
type ResultOfMnemonicFromRandom struct {
	Phrase string `json:"phrase"`
}

// ParamsOfMnemonicFromEntropy ...
type ParamsOfMnemonicFromEntropy struct {
	Entropy    string `json:"entropy"`
	Dictionary int    `json:"dictionary,omitempty"`
	WordCount  int    `json:"word_count,omitempty"`
}

// ResultOfMnemonicFromEntropy ...
type ResultOfMnemonicFromEntropy struct {
	Phrase string `json:"phrase"`
}

// ParamsOfMnemonicVerify ...
type ParamsOfMnemonicVerify struct {
	Phrase     string `json:"phrase"`
	Dictionary int    `json:"dictionary,omitempty"`
	WordCount  int    `json:"word_count,omitempty"`
}

// ResultOfMnemonicVerify ...
type ResultOfMnemonicVerify struct {
	Valid bool `json:"valid"`
}

// ParamsOfMnemonicDeriveSignKeys ...
type ParamsOfMnemonicDeriveSignKeys struct {
	Phrase     string `json:"phrase"`
	Path       string `json:"path"`
	Dictionary int    `json:"dictionary,omitempty"`
	WordCount  int    `json:"word_count,omitempty"`
}

// ParamsOfHDKeyXPrvFromMnemonic ...
type ParamsOfHDKeyXPrvFromMnemonic struct {
	Phrase string `json:"phrase"`
}

// ResultOfHDKeyXPrvFromMnemonic ...
type ResultOfHDKeyXPrvFromMnemonic struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeyDeriveFromXPrv ...
type ParamsOfHDKeyDeriveFromXPrv struct {
	Xprv       string `json:"xprv"`
	ChildIndex int    `json:"child_index"`
	Hardened   bool   `json:"hardened"`
}

// ResultOfHDKeyDeriveFromXPrv ...
type ResultOfHDKeyDeriveFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeyDeriveFromXPrvPath ...
type ParamsOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
	Path string `json:"path"`
}

// ResultOfHDKeyDeriveFromXPrvPath ...
type ResultOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
}

// ParamsOfHDKeySecretFromXPrv ...
type ParamsOfHDKeySecretFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ResultOfHDKeySecretFromXPrv ...
type ResultOfHDKeySecretFromXPrv struct {
	Secret string `json:"secret"`
}

// ParamsOfHDKeyPublicFromXPrv ...
type ParamsOfHDKeyPublicFromXPrv struct {
	Xprv string `json:"xprv"`
}

// ResultOfHDKeyPublicFromXPrv ...
type ResultOfHDKeyPublicFromXPrv struct {
	Public string `json:"public"`
}
