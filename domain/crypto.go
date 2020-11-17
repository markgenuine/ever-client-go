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
		Dictionary int    `json:"dictionary"`
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

	// CryptoUseCase ...
	CryptoUseCase interface {
		Factorize(poF ParamsOfFactorize) (int, error)
		ModularPower(pOMP ParamsOfModularPower) (int, error)
		TonCrc16(pOTC ParamsOfTonCrc16) (int, error)
		GenerateRandomBytes(pOGRB ParamsOfGenerateRandomBytes) (int, error)
		ConvertPublicKeyString(pOCPTTSF ParamsOfConvertPublicKeyToTonSafeFormat) (int, error)
		GenerateRandomSignKeys() (int, error)
		Sign(pOS ParamsOfSign) (int, error)
		VerifySignature(pOVS ParamsOfVerifySignature) (int, error)
		Sha256(pOH ParamsOfHash) (int, error)
		Sha512(pOH ParamsOfHash) (int, error)
		Scrypt(sD ParamsOfScrypt) (int, error)
		NaclSignKeypairFromSecretKey(pONSKPFC ParamsOfNaclSignKeyPairFromSecret) (int, error)
		NaclSign(pONS ParamsOfNaclSign) (int, error)
		NaclSignOpen(pONSO ParamsOfNaclSignOpen) (int, error)
		NaclSignDetached(pONS ParamsOfNaclSign) (int, error)
		NaclBoxKeypair() (int, error)
		NaclBoxKeypairFromSecretKey(pONKPFS ParamsOfNaclBoxKeyPairFromSecret) (int, error)
		NaclBox(pONB ParamsOfNaclBox) (int, error)
		NaclBoxOpen(pONBO ParamsOfNaclBoxOpen) (int, error)
		NaclSecretBox(pONSB ParamsOfNaclSecretBox) (int, error)
		NaclSecretBoxOpen(pONSBO ParamsOfNaclSecretBoxOpen) (int, error)
		MnemonicWords(pOMW ParamsOfMnemonicWords) (int, error)
		MnemonicFromRandom(pOMFR ParamsOfMnemonicFromRandom) (int, error)
		MnemonicFromEntropy(pOMFE ParamsOfMnemonicFromEntropy) (int, error)
		MnemonicVerify(pOMV ParamsOfMnemonicVerify) (int, error)
		MnemonicDeriveSignKeys(pOMDSK ParamsOfMnemonicDeriveSignKeys) (int, error)
		HdkeyXprvFromMnemonic(pOHKXFM ParamsOfHDKeyXPrvFromMnemonic) (int, error)
		HdkeyXprvDerive(hdP ParamsOfHDKeyDeriveFromXPrv) (int, error)
		HdkeyXprvDerivePath(hdPD ParamsOfHDKeyDeriveFromXPrvPath) (int, error)
		HdkeyXprvSecret(pOHKSFXP ParamsOfHDKeySecretFromXPrv) (int, error)
		HdkeyXprvPublic(pOHKPFXP ParamsOfHDKeyPublicFromXPrv) (int, error)
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
