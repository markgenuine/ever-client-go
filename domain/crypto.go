package domain

const (
	// DefaultWordCount Word count in mnemonic phrase on default
	DefaultWordCount = 12

	// DefaultDictionary dictionary mnemonic phrase on default
	DefaultDictionary = 1
)

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
	EventCallbackRegisterSigningBox func(event *ParamsOfAppSigningBox)

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
		Factorize(poF ParamsOfFactorize) (*ResultOfFactorize, error)
		ModularPower(pOMP ParamsOfModularPower) (*ResultOfModularPower, error)
		TonCrc16(pOTC ParamsOfTonCrc16) (*ResultOfTonCrc16, error)
		GenerateRandomBytes(pOGRB ParamsOfGenerateRandomBytes) (*ResultOfGenerateRandomBytes, error)
		ConvertPublicKeyString(pOCPTTSF ParamsOfConvertPublicKeyToTonSafeFormat) (*ResultOfConvertPublicKeyToTonSafeFormat, error)
		GenerateRandomSignKeys() (*KeyPair, error)
		Sign(pOS ParamsOfSign) (*ResultOfSign, error)
		VerifySignature(pOVS ParamsOfVerifySignature) (*ResultOfVerifySignature, error)
		Sha256(pOH ParamsOfHash) (*ResultOfHash, error)
		Sha512(pOH ParamsOfHash) (*ResultOfHash, error)
		Scrypt(sD ParamsOfScrypt) (*ResultOfScrypt, error)
		NaclSignKeypairFromSecretKey(pONSKPFC ParamsOfNaclSignKeyPairFromSecret) (*KeyPair, error)
		NaclSign(pONS ParamsOfNaclSign) (*ResultOfNaclSign, error)
		NaclSignOpen(pONSO ParamsOfNaclSignOpen) (*ResultOfNaclSignOpen, error)
		NaclSignDetached(pONS ParamsOfNaclSign) (*ResultOfNaclSignDetached, error)
		NaclBoxKeypair() (*KeyPair, error)
		NaclBoxKeypairFromSecretKey(pONKPFS ParamsOfNaclBoxKeyPairFromSecret) (*KeyPair, error)
		NaclBox(pONB ParamsOfNaclBox) (*ResultOfNaclBox, error)
		NaclBoxOpen(pONBO ParamsOfNaclBoxOpen) (*ResultOfNaclBoxOpen, error)
		NaclSecretBox(pONSB ParamsOfNaclSecretBox) (*ResultOfNaclBox, error)
		NaclSecretBoxOpen(pONSBO ParamsOfNaclSecretBoxOpen) (*ResultOfNaclBoxOpen, error)
		MnemonicWords(pOMW ParamsOfMnemonicWords) (*ResultOfMnemonicWords, error)
		MnemonicFromRandom(pOMFR ParamsOfMnemonicFromRandom) (*ResultOfMnemonicFromRandom, error)
		MnemonicFromEntropy(pOMFE ParamsOfMnemonicFromEntropy) (*ResultOfMnemonicFromEntropy, error)
		MnemonicVerify(pOMV ParamsOfMnemonicVerify) (*ResultOfMnemonicVerify, error)
		MnemonicDeriveSignKeys(pOMDSK ParamsOfMnemonicDeriveSignKeys) (*KeyPair, error)
		HdkeyXprvFromMnemonic(pOHKXFM ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error)
		HdkeyXprvDerive(hdP ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error)
		HdkeyXprvDerivePath(hdPD ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error)
		HdkeyXprvSecret(pOHKSFXP ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error)
		HdkeyXprvPublic(pOHKPFXP ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error)
		Chacha20(pOFCC ParamsOfChaCha20) (*ResultOfChaCha20, error)
		//RegisterSigningBox(EventCallbackRegisterSigningBox) error
		GetSigningBox(KeyPair) (*RegisteredSigningBox, error)
		SigningBoxGetPublicKey(KeyPair) (*RegisteredSigningBox, error)
		SigningBoxSign(pOSBS ParamsOfSigningBoxSign) (*ResultOfSigningBoxSign, error)
		// remove_signing_box
	}
)

// NewDefaultParamsOfMnemonicWords ...
func NewDefaultParamsOfMnemonicWords() ParamsOfMnemonicWords {
	return ParamsOfMnemonicWords{Dictionary: DefaultDictionary}
}

// NewDefaultParamsOfMnemonicFromRandom ...
func NewDefaultParamsOfMnemonicFromRandom() ParamsOfMnemonicFromRandom {
	return ParamsOfMnemonicFromRandom{Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicFromEntropy ...
func NewDefaultParamsOfMnemonicFromEntropy() ParamsOfMnemonicFromEntropy {
	return ParamsOfMnemonicFromEntropy{Entropy: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicVerify ...
func NewDefaultParamsOfMnemonicVerify() ParamsOfMnemonicVerify {
	return ParamsOfMnemonicVerify{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfMnemonicDeriveSignKeys ...
func NewDefaultParamsOfMnemonicDeriveSignKeys() ParamsOfMnemonicDeriveSignKeys {
	return ParamsOfMnemonicDeriveSignKeys{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewDefaultParamsOfHDKeyXPrvFromMnemonic ...
func NewDefaultParamsOfHDKeyXPrvFromMnemonic() ParamsOfHDKeyXPrvFromMnemonic {
	return ParamsOfHDKeyXPrvFromMnemonic{Phrase: "", Dictionary: DefaultDictionary, WordCount: DefaultWordCount}
}

// NewParamsOfAppSigningBoxGetPublicKey ...
func NewParamsOfAppSigningBoxGetPublicKey() *ParamsOfAppSigningBox {
	return &ParamsOfAppSigningBox{Type: "GetPublicKey"}
}

// NewParamsOfAppSigningBoxSign ...
func NewParamsOfAppSigningBoxSign(unsigned string) *ParamsOfAppSigningBox {
	return &ParamsOfAppSigningBox{Type: "Sign", Unsigned: unsigned}
}
