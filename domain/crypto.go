package domain

const (
	// DefaultWordCount Word count in mnemonic phrase on default
	DefaultWordCount = 12

	// DefaultDictionary dictionary mnemonic phrase on default
	DefaultDictionary = 1
)

// CryptoErrorCode ...
var CryptoErrorCode map[string]int

type (
	// SigningBoxHandle ...
	SigningBoxHandle int

	// ParamsOfFactorize ...
	ParamsOfFactorize struct {
		Composite string `json:"composite"`
	}

	// ResultOfFactorize ...
	ResultOfFactorize struct {
		Factors []string `json:"factors"`
	}

	// ParamsOfModularPower ...
	ParamsOfModularPower struct {
		Base     string `json:"base"`
		Exponent string `json:"exponent"`
		Modulus  string `json:"modulus"`
	}

	// ResultOfModularPower ...
	ResultOfModularPower struct {
		ModularPower string `json:"modular_power"`
	}

	// ParamsOfTonCrc16 ...
	ParamsOfTonCrc16 struct {
		Data string `json:"data"`
	}

	// ResultOfTonCrc16 ...
	ResultOfTonCrc16 struct {
		Crc int `json:"crc"`
	}

	// ParamsOfGenerateRandomBytes ...
	ParamsOfGenerateRandomBytes struct {
		Length int `json:"length"`
	}

	// ResultOfGenerateRandomBytes ...
	ResultOfGenerateRandomBytes struct {
		Bytes string `json:"bytes"`
	}

	// ParamsOfConvertPublicKeyToTonSafeFormat ...
	ParamsOfConvertPublicKeyToTonSafeFormat struct {
		PublicKey string `json:"public_key"`
	}

	// ResultOfConvertPublicKeyToTonSafeFormat ...
	ResultOfConvertPublicKeyToTonSafeFormat struct {
		TonPublicKey string `json:"ton_public_key"`
	}

	// KeyPair ...
	KeyPair struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	}

	// ParamsOfSign ...
	ParamsOfSign struct {
		Unsigned string  `json:"unsigned"`
		Keys     KeyPair `json:"keys"`
	}

	// ResultOfSign ...
	ResultOfSign struct {
		Signed    string `json:"signed"`
		Signature string `json:"signature"`
	}

	// ParamsOfVerifySignature ...
	ParamsOfVerifySignature struct {
		Signed string `json:"signed"`
		Public string `json:"public"`
	}

	// ResultOfVerifySignature ...
	ResultOfVerifySignature struct {
		Unsigned string `json:"unsigned"`
	}

	// ParamsOfHash ...
	ParamsOfHash struct {
		Data string `json:"data"`
	}

	// ResultOfHash ...
	ResultOfHash struct {
		Hash string `json:"hash"`
	}

	// ParamsOfScrypt ...
	ParamsOfScrypt struct {
		Password string `json:"password"`
		Salt     string `json:"salt"`
		LogN     int    `json:"log_n"`
		R        int    `json:"r"`
		P        int    `json:"p"`
		DkLen    int    `json:"dk_len"`
	}

	// ResultOfScrypt ...
	ResultOfScrypt struct {
		Key string `json:"key"`
	}

	// ParamsOfNaclSignKeyPairFromSecret ...
	ParamsOfNaclSignKeyPairFromSecret struct {
		Secret string `json:"secret"`
	}

	// ParamsOfNaclSign ...
	ParamsOfNaclSign struct {
		Unsigned string `json:"unsigned"`
		Secret   string `json:"secret"`
	}

	// ResultOfNaclSign ...
	ResultOfNaclSign struct {
		Signed string `json:"signed"`
	}

	// ParamsOfNaclSignOpen ...
	ParamsOfNaclSignOpen struct {
		Signed string `json:"signed"`
		Public string `json:"public"`
	}

	// ResultOfNaclSignOpen ...
	ResultOfNaclSignOpen struct {
		Unsigned string `json:"unsigned"`
	}

	// ResultOfNaclSignDetached ...
	ResultOfNaclSignDetached struct {
		Signature string `json:"signature"`
	}

	// ParamsOfNaclSignDetachedVerify ...
	ParamsOfNaclSignDetachedVerify struct {
		Unsigned  string `json:"unsigned"`
		Signature string `json:"signature"`
		Public    string `json:"public"`
	}

	// ResultOfNaclSignDetachedVerify ...
	ResultOfNaclSignDetachedVerify struct {
		Succeeded bool `json:"succeeded"`
	}

	// ParamsOfNaclBoxKeyPairFromSecret ...
	ParamsOfNaclBoxKeyPairFromSecret struct {
		Secret string `json:"secret"`
	}

	// ParamsOfNaclBox ...
	ParamsOfNaclBox struct {
		Decrypted   string `json:"decrypted"`
		Nonce       string `json:"nonce"`
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
	}

	// ResultOfNaclBox ...
	ResultOfNaclBox struct {
		Encrypted string `json:"encrypted"`
	}

	// ParamsOfNaclBoxOpen ...
	ParamsOfNaclBoxOpen struct {
		Encrypted   string `json:"encrypted"`
		Nonce       string `json:"nonce"`
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
	}

	// ResultOfNaclBoxOpen ...
	ResultOfNaclBoxOpen struct {
		Decrypted string `json:"decrypted"`
	}

	// ParamsOfNaclSecretBox ...
	ParamsOfNaclSecretBox struct {
		Decrypted string `json:"decrypted"`
		Nonce     string `json:"nonce"`
		Key       string `json:"key"`
	}

	// ParamsOfNaclSecretBoxOpen ...
	ParamsOfNaclSecretBoxOpen struct {
		Encrypted string `json:"encrypted"`
		Nonce     string `json:"nonce"`
		Key       string `json:"key"`
	}

	// ParamsOfMnemonicWords ...
	ParamsOfMnemonicWords struct {
		Dictionary int `json:"dictionary,omitempty"`
	}

	// ResultOfMnemonicWords ...
	ResultOfMnemonicWords struct {
		Words string `json:"words"`
	}

	// ParamsOfMnemonicFromRandom ...
	ParamsOfMnemonicFromRandom struct {
		Dictionary int `json:"dictionary,omitempty"`
		WordCount  int `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicFromRandom ...
	ResultOfMnemonicFromRandom struct {
		Phrase string `json:"phrase"`
	}

	// ParamsOfMnemonicFromEntropy ...
	ParamsOfMnemonicFromEntropy struct {
		Entropy    string `json:"entropy"`
		Dictionary int    `json:"dictionary,omitempty"`
		WordCount  int    `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicFromEntropy ...
	ResultOfMnemonicFromEntropy struct {
		Phrase string `json:"phrase"`
	}

	// ParamsOfMnemonicVerify ...
	ParamsOfMnemonicVerify struct {
		Phrase     string `json:"phrase"`
		Dictionary int    `json:"dictionary,omitempty"`
		WordCount  int    `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicVerify ...
	ResultOfMnemonicVerify struct {
		Valid bool `json:"valid"`
	}

	// ParamsOfMnemonicDeriveSignKeys ...
	ParamsOfMnemonicDeriveSignKeys struct {
		Phrase     string `json:"phrase"`
		Path       string `json:"path,omitempty"`
		Dictionary int    `json:"dictionary,omitempty"`
		WordCount  int    `json:"word_count,omitempty"`
	}

	// ParamsOfHDKeyXPrvFromMnemonic ...
	ParamsOfHDKeyXPrvFromMnemonic struct {
		Phrase     string `json:"phrase"`
		Dictionary int    `json:"dictionary,omitempty"`
		WordCount  int    `json:"word_count,omitempty"`
	}

	// ResultOfHDKeyXPrvFromMnemonic ...
	ResultOfHDKeyXPrvFromMnemonic struct {
		Xprv string `json:"xprv"`
	}

	// ParamsOfHDKeyDeriveFromXPrv ...
	ParamsOfHDKeyDeriveFromXPrv struct {
		Xprv       string `json:"xprv"`
		ChildIndex int    `json:"child_index"`
		Hardened   bool   `json:"hardened"`
	}

	// ResultOfHDKeyDeriveFromXPrv ...
	ResultOfHDKeyDeriveFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	// ParamsOfHDKeyDeriveFromXPrvPath ...
	ParamsOfHDKeyDeriveFromXPrvPath struct {
		Xprv string `json:"xprv"`
		Path string `json:"path"`
	}

	// ResultOfHDKeyDeriveFromXPrvPath ...
	ResultOfHDKeyDeriveFromXPrvPath struct {
		Xprv string `json:"xprv"`
	}

	// ParamsOfHDKeySecretFromXPrv ...
	ParamsOfHDKeySecretFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	// ResultOfHDKeySecretFromXPrv ...
	ResultOfHDKeySecretFromXPrv struct {
		Secret string `json:"secret"`
	}

	// ParamsOfHDKeyPublicFromXPrv ...
	ParamsOfHDKeyPublicFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	// ResultOfHDKeyPublicFromXPrv ...
	ResultOfHDKeyPublicFromXPrv struct {
		Public string `json:"public"`
	}

	// ParamsOfChaCha20 ...
	ParamsOfChaCha20 struct {
		Data  string `json:"data"`
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	// ResultOfChaCha20 ...
	ResultOfChaCha20 struct {
		Data string `json:"data"`
	}

	// RegisteredSigningBox ...
	RegisteredSigningBox struct {
		Handle SigningBoxHandle
	}

	// ParamsOfAppSigningBox ...
	ParamsOfAppSigningBox struct {
		Type     string `json:"type"`
		Unsigned string `json:"unsigned,omitempty"`
	}

	// ResultOfAppSigningBox ...
	ResultOfAppSigningBox struct {
		Type      string `json:"type"`
		PublicKey string `json:"public_key,omitempty"`
		Signature string `json:"signature,omitempty"`
	}

	// ResultOfSigningBoxGetPublicKey ...
	ResultOfSigningBoxGetPublicKey struct {
		PubKey string `json:"pubkey"`
	}

	// EventCallbackRegisterSigningBox ...
	//EventCallbackRegisterSigningBox func(event *ParamsOfAppSigningBox)

	// ParamsOfSigningBoxSign ...
	ParamsOfSigningBoxSign struct {
		SigningBox SigningBoxHandle `json:"signing_box"`
		Unsigned   string           `json:"unsigned"`
	}

	// ResultOfSigningBoxSign ...
	ResultOfSigningBoxSign struct {
		Signature string `json:"signature"`
	}

	// CryptoUseCase ...
	CryptoUseCase interface {
		Factorize(*ParamsOfFactorize) (*ResultOfFactorize, error)
		ModularPower(*ParamsOfModularPower) (*ResultOfModularPower, error)
		TonCrc16(*ParamsOfTonCrc16) (*ResultOfTonCrc16, error)
		GenerateRandomBytes(*ParamsOfGenerateRandomBytes) (*ResultOfGenerateRandomBytes, error)
		ConvertPublicKeyString(*ParamsOfConvertPublicKeyToTonSafeFormat) (*ResultOfConvertPublicKeyToTonSafeFormat, error)
		GenerateRandomSignKeys() (*KeyPair, error)
		Sign(*ParamsOfSign) (*ResultOfSign, error)
		VerifySignature(*ParamsOfVerifySignature) (*ResultOfVerifySignature, error)
		Sha256(*ParamsOfHash) (*ResultOfHash, error)
		Sha512(*ParamsOfHash) (*ResultOfHash, error)
		Scrypt(*ParamsOfScrypt) (*ResultOfScrypt, error)
		NaclSignKeypairFromSecretKey(*ParamsOfNaclSignKeyPairFromSecret) (*KeyPair, error)
		NaclSign(*ParamsOfNaclSign) (*ResultOfNaclSign, error)
		NaclSignOpen(*ParamsOfNaclSignOpen) (*ResultOfNaclSignOpen, error)
		NaclSignDetached(*ParamsOfNaclSign) (*ResultOfNaclSignDetached, error)
		NaclSignDetachedVerify(*ParamsOfNaclSignDetachedVerify) (*ResultOfNaclSignDetachedVerify, error)
		NaclBoxKeypair() (*KeyPair, error)
		NaclBoxKeypairFromSecretKey(*ParamsOfNaclBoxKeyPairFromSecret) (*KeyPair, error)
		NaclBox(*ParamsOfNaclBox) (*ResultOfNaclBox, error)
		NaclBoxOpen(*ParamsOfNaclBoxOpen) (*ResultOfNaclBoxOpen, error)
		NaclSecretBox(*ParamsOfNaclSecretBox) (*ResultOfNaclBox, error)
		NaclSecretBoxOpen(*ParamsOfNaclSecretBoxOpen) (*ResultOfNaclBoxOpen, error)
		MnemonicWords(*ParamsOfMnemonicWords) (*ResultOfMnemonicWords, error)
		MnemonicFromRandom(*ParamsOfMnemonicFromRandom) (*ResultOfMnemonicFromRandom, error)
		MnemonicFromEntropy(*ParamsOfMnemonicFromEntropy) (*ResultOfMnemonicFromEntropy, error)
		MnemonicVerify(*ParamsOfMnemonicVerify) (*ResultOfMnemonicVerify, error)
		MnemonicDeriveSignKeys(*ParamsOfMnemonicDeriveSignKeys) (*KeyPair, error)
		HdkeyXprvFromMnemonic(*ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error)
		HdkeyDeriveFromXprv(*ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error)
		HdkeyDeriveFromXprvPath(*ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error)
		HdkeySecretFromXprv(*ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error)
		HdkeyPublicFromXprv(*ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error)
		Chacha20(*ParamsOfChaCha20) (*ResultOfChaCha20, error)
		RegisterSigningBox(AppSigningBox) (*RegisteredSigningBox, error)
		GetSigningBox(*KeyPair) (*RegisteredSigningBox, error)
		SigningBoxGetPublicKey(*RegisteredSigningBox) (*ResultOfSigningBoxGetPublicKey, error)
		SigningBoxSign(*ParamsOfSigningBoxSign) (*ResultOfSigningBoxSign, error)
		RemoveSigningBox(rSB *RegisteredSigningBox) error
	}
)

func init() {
	CryptoErrorCode = map[string]int{
		"InvalidPublicKey":          100,
		"InvalidSecretKey":          101,
		"InvalidKey":                102,
		"InvalidFactorizeChallenge": 106,
		"InvalidBigInt":             107,
		"ScryptFailed":              108,
		"InvalidKeySize":            109,
		"NaclSecretBoxFailed":       110,
		"NaclBoxFailed":             111,
		"NaclSignFailed":            112,
		"Bip39InvalidEntropy":       113,
		"Bip39InvalidPhrase":        114,
		"Bip32InvalidKey":           115,
		"Bip32InvalidDerivePath":    116,
		"Bip39InvalidDictionary":    117,
		"Bip39InvalidWordCount":     118,
		"MnemonicGenerationFailed":  119,
		"MnemonicFromEntropyFailed": 120,
		"SigningBoxNotRegistered":   121,
		"InvalidSignature":          122,
	}

}

// NewDefaultParamsOfMnemonicWords ...
func NewDefaultParamsOfMnemonicWords() *ParamsOfMnemonicWords {
	return &ParamsOfMnemonicWords{Dictionary: DefaultDictionary}
}

// NewDefaultParamsOfMnemonicFromRandom ...
func NewDefaultParamsOfMnemonicFromRandom() *ParamsOfMnemonicFromRandom {
	return &ParamsOfMnemonicFromRandom{Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicFromEntropy ...
func NewDefaultParamsOfMnemonicFromEntropy() *ParamsOfMnemonicFromEntropy {
	return &ParamsOfMnemonicFromEntropy{Entropy: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicVerify ...
func NewDefaultParamsOfMnemonicVerify() *ParamsOfMnemonicVerify {
	return &ParamsOfMnemonicVerify{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicDeriveSignKeys ...
func NewDefaultParamsOfMnemonicDeriveSignKeys() *ParamsOfMnemonicDeriveSignKeys {
	return &ParamsOfMnemonicDeriveSignKeys{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfHDKeyXPrvFromMnemonic ...
func NewDefaultParamsOfHDKeyXPrvFromMnemonic() *ParamsOfHDKeyXPrvFromMnemonic {
	return &ParamsOfHDKeyXPrvFromMnemonic{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewParamsOfAppSigningBoxGetPublicKey ...
func NewParamsOfAppSigningBoxGetPublicKey() *ParamsOfAppSigningBox {
	return &ParamsOfAppSigningBox{Type: "GetPublicKey"}
}

// NewParamsOfAppSigningBoxSign ...
func NewParamsOfAppSigningBoxSign() *ParamsOfAppSigningBox {
	return &ParamsOfAppSigningBox{Type: "Sign"}
}

// NewResultOfAppSigningBoxGetPublicKey
func NewResultOfAppSigningBoxGetPublicKey() *ResultOfAppSigningBox {
	return &ResultOfAppSigningBox{Type: "GetPublicKey"}
}

// NewResultOfAppSigningBoxSign
func NewResultOfAppSigningBoxSign() *ResultOfAppSigningBox {
	return &ResultOfAppSigningBox{Type: "Sign"}
}
