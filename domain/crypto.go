package domain

import (
	"encoding/json"
	"fmt"
	"github.com/move-ton/ever-client-go/util"
)

const (
	// DefaultWordCount Word count in mnemonic phrase on default
	DefaultWordCount = 12

	// DefaultDictionary dictionary mnemonic phrase on default
	DefaultDictionary = 1

	// TonMnemonicDictionary ...
	TonMnemonicDictionary MnemonicDictionary = 0

	// EnglishMnemonicDictionary ...
	EnglishMnemonicDictionary MnemonicDictionary = 1

	// ChineseSimplifiedMnemonicDictionary ...
	ChineseSimplifiedMnemonicDictionary MnemonicDictionary = 2

	// ChineseTraditionalMnemonicDictionary ...
	ChineseTraditionalMnemonicDictionary MnemonicDictionary = 3

	// FrenchMnemonicDictionary ...
	FrenchMnemonicDictionary MnemonicDictionary = 4

	// ItalianMnemonicDictionary ...
	ItalianMnemonicDictionary MnemonicDictionary = 5

	// JapaneseMnemonicDictionary ...
	JapaneseMnemonicDictionary MnemonicDictionary = 6

	// KoreanMnemonicDictionary ...
	KoreanMnemonicDictionary MnemonicDictionary = 7

	// SpanishMnemonicDictionary ...
	SpanishMnemonicDictionary MnemonicDictionary = 8

	// CipherModeCBC ...
	CipherModeCBC CipherMode = "CBC"

	// CipherModeCFB ...
	CipherModeCFB CipherMode = "CFB"

	// CipherModeCTR ...
	CipherModeCTR CipherMode = "CTR"

	// CipherModeECB ...
	CipherModeECB CipherMode = "ECB"

	// CipherModeOFB ...
	CipherModeOFB CipherMode = "OFB"
)

var CryptoErrorCode map[string]int

type (
	CryptoBoxHandle     int
	SigningBoxHandle    int
	EncryptionBoxHandle int
	MnemonicDictionary  int

	// EncryptionBoxInfo - Encryption box information.
	EncryptionBoxInfo struct {
		HDPath    string          `json:"hdpath,omitempty"`
		Algorithm string          `json:"algorithm,omitempty"`
		Options   json.RawMessage `json:"options,omitempty"`
		Public    json.RawMessage `json:"public,omitempty"`
	}

	EncryptionAlgorithmAESVariant struct {
		Value AesParamsEB `json:"value"`
	}

	EncryptionAlgorithmChaCha20Variant struct {
		Value ChaCha20ParamsEB `json:"value"`
	}

	EncryptionAlgorithmNaclBoxVariant struct {
		Value NaclBoxParamsEB `json:"value"`
	}

	EncryptionAlgorithmNaclSecretBoxVariant struct {
		Value NaclSecretBoxParamsEB `json:"value"`
	}

	EncryptionAlgorithm struct {
		ValueEnumType interface{}
	}

	BoxEncryptionAlgorithmChaCha20Variant struct {
		Value ChaCha20ParamsCB `json:"value"`
	}

	BoxEncryptionAlgorithmNaclBoxVariant struct {
		Value NaclBoxParamsCB `json:"value"`
	}

	BoxEncryptionAlgorithmNaclSecretBoxVariant struct {
		Value NaclSecretBoxParamsCB `json:"value"`
	}

	BoxEncryptionAlgorithm struct {
		ValueEnumType interface{}
	}

	CipherMode string

	AesParamsEB struct {
		Mode CipherMode `json:"mode"`
		Key  string     `json:"key"`
		IV   string     `json:"iv,omitempty"`
	}

	AesInfo struct {
		Mode CipherMode `json:"mode"`
		IV   string     `json:"iv,omitempty"`
	}

	ChaCha20ParamsEB struct {
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	NaclBoxParamsEB struct {
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
		Nonce       string `json:"nonce"`
	}

	NaclSecretBoxParamsEB struct {
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	ChaCha20ParamsCB struct {
		Nonce string `json:"nonce"`
	}

	NaclBoxParamsCB struct {
		TheirPublic string `json:"their_public"`
		Nonce       string `json:"nonce"`
	}

	NaclSecretBoxParamsCB struct {
		Nonce string `json:"nonce"`
	}

	CryptoBoxSecret struct {
		ValueEnumType interface{}
	}

	// CryptoBoxSecretRandomSeedPhraseVariant - Creates Crypto Box from a random seed phrase. This option can be
	// used if a developer doesn't want the seed phrase to leave the core library's memory, where it is stored encrypted.
	// This type should be used upon the first wallet initialization, all further initializations should use
	// EncryptedSecret type instead.
	// Get encrypted_secret with get_crypto_box_info function and store it on your side.
	CryptoBoxSecretRandomSeedPhraseVariant struct {
		Dictionary MnemonicDictionary `json:"dictionary"`
		WordCount  int                `json:"wordcount"`
	}

	// CryptoBoxSecretPredefinedSeedPhraseVariant - Restores crypto box instance from an existing seed phrase.
	// This type should be used when Crypto Box is initialized from a seed phrase, entered by a user.
	// This type should be used only upon the first wallet initialization, all further initializations should
	// use EncryptedSecret type instead.
	// Get encrypted_secret with get_crypto_box_info function and store it on your side.
	CryptoBoxSecretPredefinedSeedPhraseVariant struct {
		Phrase     string             `json:"phrase"`
		Dictionary MnemonicDictionary `json:"dictionary"`
		WordCount  int                `json:"wordcount"`
	}

	// CryptoBoxSecretEncryptedSecretVariant - Use this type for wallet reinitializations, when you already
	// have encrypted_secret on hands. To get encrypted_secret, use get_crypto_box_info function after
	// you initialized your crypto box for the first time.
	// It is an object, containing seed phrase or private key, encrypted with secret_encryption_salt
	// and password from password_provider.
	// Note that if you want to change salt or password provider, then you need to reinitialize the wallet
	// with PredefinedSeedPhrase, then get EncryptedSecret via get_crypto_box_info, store it somewhere, and only after that initialize the wallet with EncryptedSecret type
	CryptoBoxSecretEncryptedSecretVariant struct {
		EncryptedSecret string `json:"encrypted_secret"`
	}

	ParamsOfFactorize struct {
		Composite string `json:"composite"`
	}

	ResultOfFactorize struct {
		Factors []string `json:"factors"`
	}

	ParamsOfModularPower struct {
		Base     string `json:"base"`
		Exponent string `json:"exponent"`
		Modulus  string `json:"modulus"`
	}

	ResultOfModularPower struct {
		ModularPower string `json:"modular_power"`
	}

	ParamsOfTonCrc16 struct {
		Data string `json:"data"`
	}

	ResultOfTonCrc16 struct {
		Crc int `json:"crc"`
	}

	ParamsOfGenerateRandomBytes struct {
		Length int `json:"length"`
	}

	ResultOfGenerateRandomBytes struct {
		Bytes string `json:"bytes"`
	}

	ParamsOfConvertPublicKeyToTonSafeFormat struct {
		PublicKey string `json:"public_key"`
	}

	ResultOfConvertPublicKeyToTonSafeFormat struct {
		TonPublicKey string `json:"ton_public_key"`
	}

	KeyPair struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	}

	ParamsOfSign struct {
		Unsigned string   `json:"unsigned"`
		Keys     *KeyPair `json:"keys"`
	}

	ResultOfSign struct {
		Signed    string `json:"signed"`
		Signature string `json:"signature"`
	}

	ParamsOfVerifySignature struct {
		Signed string `json:"signed"`
		Public string `json:"public"`
	}

	ResultOfVerifySignature struct {
		Unsigned string `json:"unsigned"`
	}

	ParamsOfHash struct {
		Data string `json:"data"`
	}

	ResultOfHash struct {
		Hash string `json:"hash"`
	}

	ParamsOfScrypt struct {
		Password string `json:"password"`
		Salt     string `json:"salt"`
		LogN     int    `json:"log_n"`
		R        int    `json:"r"`
		P        int    `json:"p"`
		DkLen    int    `json:"dk_len"`
	}

	ResultOfScrypt struct {
		Key string `json:"key"`
	}

	ParamsOfNaclSignKeyPairFromSecret struct {
		Secret string `json:"secret"`
	}

	ParamsOfNaclSign struct {
		Unsigned string `json:"unsigned"`
		Secret   string `json:"secret"`
	}

	ResultOfNaclSign struct {
		Signed string `json:"signed"`
	}

	ParamsOfNaclSignOpen struct {
		Signed string `json:"signed"`
		Public string `json:"public"`
	}

	ResultOfNaclSignOpen struct {
		Unsigned string `json:"unsigned"`
	}

	ResultOfNaclSignDetached struct {
		Signature string `json:"signature"`
	}

	ParamsOfNaclSignDetachedVerify struct {
		Unsigned  string `json:"unsigned"`
		Signature string `json:"signature"`
		Public    string `json:"public"`
	}

	ResultOfNaclSignDetachedVerify struct {
		Succeeded bool `json:"succeeded"`
	}

	ParamsOfNaclBoxKeyPairFromSecret struct {
		Secret string `json:"secret"`
	}

	ParamsOfNaclBox struct {
		Decrypted   string `json:"decrypted"`
		Nonce       string `json:"nonce"`
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
	}

	ResultOfNaclBox struct {
		Encrypted string `json:"encrypted"`
	}

	ParamsOfNaclBoxOpen struct {
		Encrypted   string `json:"encrypted"`
		Nonce       string `json:"nonce"`
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
	}

	ResultOfNaclBoxOpen struct {
		Decrypted string `json:"decrypted"`
	}

	ParamsOfNaclSecretBox struct {
		Decrypted string `json:"decrypted"`
		Nonce     string `json:"nonce"`
		Key       string `json:"key"`
	}

	ParamsOfNaclSecretBoxOpen struct {
		Encrypted string `json:"encrypted"`
		Nonce     string `json:"nonce"`
		Key       string `json:"key"`
	}

	ParamsOfMnemonicWords struct {
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
	}

	ResultOfMnemonicWords struct {
		Words string `json:"words"`
	}

	ParamsOfMnemonicFromRandom struct {
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
		WordCount  *int                `json:"word_count,omitempty"`
	}

	ResultOfMnemonicFromRandom struct {
		Phrase string `json:"phrase"`
	}

	ParamsOfMnemonicFromEntropy struct {
		Entropy    string              `json:"entropy"`
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
		WordCount  *int                `json:"word_count,omitempty"`
	}

	ResultOfMnemonicFromEntropy struct {
		Phrase string `json:"phrase"`
	}

	ParamsOfMnemonicVerify struct {
		Phrase     string              `json:"phrase"`
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
		WordCount  *int                `json:"word_count,omitempty"`
	}

	ResultOfMnemonicVerify struct {
		Valid bool `json:"valid"`
	}

	ParamsOfMnemonicDeriveSignKeys struct {
		Phrase     string              `json:"phrase"`
		Path       string              `json:"path,omitempty"`
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
		WordCount  *int                `json:"word_count,omitempty"`
	}

	ParamsOfHDKeyXPrvFromMnemonic struct {
		Phrase     string              `json:"phrase"`
		Dictionary *MnemonicDictionary `json:"dictionary,omitempty"`
		WordCount  *int                `json:"word_count,omitempty"`
	}

	ResultOfHDKeyXPrvFromMnemonic struct {
		Xprv string `json:"xprv"`
	}

	ParamsOfHDKeyDeriveFromXPrv struct {
		Xprv       string `json:"xprv"`
		ChildIndex int    `json:"child_index"`
		Hardened   bool   `json:"hardened"`
	}

	ResultOfHDKeyDeriveFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	ParamsOfHDKeyDeriveFromXPrvPath struct {
		Xprv string `json:"xprv"`
		Path string `json:"path"`
	}

	ResultOfHDKeyDeriveFromXPrvPath struct {
		Xprv string `json:"xprv"`
	}

	ParamsOfHDKeySecretFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	ResultOfHDKeySecretFromXPrv struct {
		Secret string `json:"secret"`
	}

	ParamsOfHDKeyPublicFromXPrv struct {
		Xprv string `json:"xprv"`
	}

	ResultOfHDKeyPublicFromXPrv struct {
		Public string `json:"public"`
	}

	ParamsOfChaCha20 struct {
		Data  string `json:"data"`
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	ResultOfChaCha20 struct {
		Data string `json:"data"`
	}

	ParamsOfCreateCryptoBox struct {
		SecretEncryptionSalt string          `json:"secret_encryption_salt"`
		Secret               CryptoBoxSecret `json:"secret"`
	}

	RegisteredCryptoBox struct {
		Handle CryptoBoxHandle `json:"handle"`
	}

	ResultOfGetCryptoBoxInfo struct {
		EncryptedSecret string `json:"encrypted_secret"`
	}

	ResultOfGetCryptoBoxSeedPhrase struct {
		Phrase     string             `json:"phrase"`
		Dictionary MnemonicDictionary `json:"dictionary"`
		Wordcount  int                `json:"wordcount"`
	}

	ParamsOfGetSigningBoxFromCryptoBox struct {
		Handle         int    `json:"handle"`
		HdPath         string `json:"hdpath,omitempty"`
		SecretLifetime int    `json:"secret_lifetime,omitempty"`
	}

	ParamsOfGetEncryptionBoxFromCryptoBox struct {
		Handle         int                     `json:"handle"`
		HdPath         string                  `json:"hdpath,omitempty"`
		Algorithm      *BoxEncryptionAlgorithm `json:"algorithm"`
		SecretLifetime int                     `json:"secret_lifetime,omitempty"`
	}

	RegisteredSigningBox struct {
		Handle SigningBoxHandle `json:"handle"`
	}

	ParamsOfAppSigningBox struct {
		ValueEnumType interface{}
	}

	ParamsOfAppSigningBoxGetPublicKey struct{}

	ParamsOfAppSigningBoxSign struct {
		Unsigned string `json:"unsigned"`
	}

	ParamsOfAppPasswordProvider struct {
		ValueEnumType interface{}
	}

	ParamsOfAppPasswordProviderGetPassword struct {
		EncryptionPublicKey string `json:"encryption_public_key"`
	}

	ResultOfAppPasswordProvider struct {
		ValueEnumType interface{}
	}

	ResultOfAppPasswordProviderGetPassword struct {
		EncryptedPassword   string `json:"encrypted_password"`
		AppEncryptionPubkey string `json:"app_encryption_pubkey"`
	}

	ResultOfAppSigningBox struct {
		ValueEnumType interface{}
	}

	ResultOfAppSigningBoxGetPublicKey struct {
		PublicKey string `json:"public_key"`
	}

	ResultOfAppSigningBoxSign struct {
		Signature string `json:"signature"`
	}

	ResultOfSigningBoxGetPublicKey struct {
		PubKey string `json:"PubKey"`
	}

	ParamsOfSigningBoxSign struct {
		SigningBox SigningBoxHandle `json:"signing_box"`
		Unsigned   string           `json:"unsigned"`
	}

	ResultOfSigningBoxSign struct {
		Signature string `json:"signature"`
	}

	RegisteredEncryptionBox struct {
		Handle EncryptionBoxHandle `json:"handle"`
	}

	// ParamsOfAppEncryptionBox - Encryption box callbacks.
	ParamsOfAppEncryptionBox struct {
		ValueEnumType interface{}
	}

	ParamsOfAppEncryptionBoxGetInfo struct{}

	ParamsOfAppEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	ParamsOfAppEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	// ResultOfAppEncryptionBox - Returning values from signing box callbacks.
	ResultOfAppEncryptionBox struct {
		ValueEnumType interface{}
	}

	ResultOfAppEncryptionBoxGetInfo struct {
		Info EncryptionBoxInfo `json:"info"`
	}

	ResultOfAppEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	ResultOfAppEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	ParamsOfEncryptionBoxGetInfo struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
	}

	ResultOfEncryptionBoxGetInfo struct {
		Info EncryptionBoxInfo `json:"info"`
	}

	ParamsOfEncryptionBoxEncrypt struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
		Data          string              `json:"data"`
	}

	ResultOfEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	ParamsOfEncryptionBoxDecrypt struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
		Data          string              `json:"data"`
	}

	ResultOfEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	ParamsOfCreateEncryptionBox struct {
		Algorithm EncryptionAlgorithm `json:"algorithm"`
	}

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
		HDKeyXprvFromMnemonic(*ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error)
		HDKeyDeriveFromXprv(*ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error)
		HDKeyDeriveFromXprvPath(*ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error)
		HDKeySecretFromXprv(*ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error)
		HDKeyPublicFromXprv(*ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error)
		Chacha20(*ParamsOfChaCha20) (*ResultOfChaCha20, error)
		CreateCryptoBox(*ParamsOfCreateCryptoBox, AppPasswordProvider) (*RegisteredCryptoBox, error)
		RemoveCryptoBox(*RegisteredCryptoBox) error
		GetCryptoBoxInfo(*RegisteredCryptoBox) (*ResultOfGetCryptoBoxInfo, error)
		GetCryptoBoxSeedPhrase(*RegisteredCryptoBox) (*ResultOfGetCryptoBoxSeedPhrase, error)
		GetSigningBoxFromCryptoBox(*ParamsOfGetSigningBoxFromCryptoBox) (*RegisteredSigningBox, error)
		GetEncryptionBoxFromCryptoBox(box *ParamsOfGetEncryptionBoxFromCryptoBox) (*RegisteredEncryptionBox, error)
		ClearCryptoBoxSecretCache(*RegisteredCryptoBox) error
		RegisterSigningBox(AppSigningBox) (*RegisteredSigningBox, error)
		GetSigningBox(*KeyPair) (*RegisteredSigningBox, error)
		SigningBoxGetPublicKey(*RegisteredSigningBox) (*ResultOfSigningBoxGetPublicKey, error)
		SigningBoxSign(*ParamsOfSigningBoxSign) (*ResultOfSigningBoxSign, error)
		RemoveSigningBox(*RegisteredSigningBox) error
		RegisterEncryptionBox(AppEncryptionBox) (*RegisteredEncryptionBox, error)
		RemoveEncryptionBox(*RegisteredEncryptionBox) error
		EncryptionBoxGetInfo(*ParamsOfEncryptionBoxGetInfo) (*ResultOfEncryptionBoxGetInfo, error)
		EncryptionBoxEncrypt(*ParamsOfEncryptionBoxEncrypt) (*ResultOfEncryptionBoxEncrypt, error)
		EncryptionBoxDecrypt(*ParamsOfEncryptionBoxDecrypt) (*ResultOfEncryptionBoxDecrypt, error)
		CreateEncryptionBox(*ParamsOfCreateEncryptionBox) (*RegisteredEncryptionBox, error)
	}
)

func init() {
	// List errors crypto module
	CryptoErrorCode = map[string]int{
		"InvalidPublicKey":                    100,
		"InvalidSecretKey":                    101,
		"InvalidKey":                          102,
		"InvalidFactorizeChallenge":           106,
		"InvalidBigInt":                       107,
		"ScryptFailed":                        108,
		"InvalidKeySize":                      109,
		"NaclSecretBoxFailed":                 110,
		"NaclBoxFailed":                       111,
		"NaclSignFailed":                      112,
		"Bip39InvalidEntropy":                 113,
		"Bip39InvalidPhrase":                  114,
		"Bip32InvalidKey":                     115,
		"Bip32InvalidDerivePath":              116,
		"Bip39InvalidDictionary":              117,
		"Bip39InvalidWordCount":               118,
		"MnemonicGenerationFailed":            119,
		"MnemonicFromEntropyFailed":           120,
		"SigningBoxNotRegistered":             121,
		"InvalidSignature":                    122,
		"EncryptionBoxNotRegistered":          123,
		"InvalidIvSize":                       124,
		"UnsupportedCipherMode":               125,
		"CannotCreateCipher":                  126,
		"EncryptDataError":                    127,
		"DecryptDataError":                    128,
		"IvRequired":                          129,
		"CryptoBoxNotRegistered":              130,
		"InvalidCryptoBoxType":                131,
		"CryptoBoxSecretSerializationError":   132,
		"CryptoBoxSecretDeserializationError": 133,
		"InvalidNonceSize":                    134,
	}
}

func NewDefaultParamsOfMnemonicWords() *ParamsOfMnemonicWords {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfMnemonicWords{Dictionary: &dict}
}

func NewDefaultParamsOfMnemonicFromRandom() *ParamsOfMnemonicFromRandom {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfMnemonicFromRandom{Dictionary: &dict, WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func NewDefaultParamsOfMnemonicFromEntropy() *ParamsOfMnemonicFromEntropy {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfMnemonicFromEntropy{Entropy: "", Dictionary: &dict, WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func NewDefaultParamsOfMnemonicVerify() *ParamsOfMnemonicVerify {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfMnemonicVerify{Phrase: "", Dictionary: &dict, WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func NewDefaultParamsOfMnemonicDeriveSignKeys() *ParamsOfMnemonicDeriveSignKeys {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfMnemonicDeriveSignKeys{Phrase: "", Dictionary: &dict, WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func NewDefaultParamsOfHDKeyXPrvFromMnemonic() *ParamsOfHDKeyXPrvFromMnemonic {
	dict := MnemonicDictionary(DefaultDictionary)
	return &ParamsOfHDKeyXPrvFromMnemonic{Phrase: "", Dictionary: &dict, WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func (cbs *CryptoBoxSecret) MarshalJSON() ([]byte, error) {
	switch value := (cbs.ValueEnumType).(type) {
	case CryptoBoxSecretRandomSeedPhraseVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretRandomSeedPhraseVariant
		}{"RandomSeedPhrase", value})
	case CryptoBoxSecretPredefinedSeedPhraseVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretPredefinedSeedPhraseVariant
		}{"PredefinedSeedPhrase", value})
	case CryptoBoxSecretEncryptedSecretVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretEncryptedSecretVariant
		}{"EncryptedSecret", value})
	default:
		return nil, fmt.Errorf("unsupported type for CryptoBoxSecret %v", cbs.ValueEnumType)
	}
}

func (cbs *CryptoBoxSecret) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "RandomSeedPhrase":
		var valueEnum CryptoBoxSecretRandomSeedPhraseVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		cbs.ValueEnumType = valueEnum
	case "PredefinedSeedPhrase":
		var valueEnum CryptoBoxSecretPredefinedSeedPhraseVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		cbs.ValueEnumType = valueEnum
	case "EncryptedSecret":
		var valueEnum CryptoBoxSecretEncryptedSecretVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		cbs.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for CryptoBoxSecret %v", typeD.Type)
	}
	return nil
}

func (pOAPP *ParamsOfAppPasswordProvider) MarshalJSON() ([]byte, error) {
	switch value := (pOAPP.ValueEnumType).(type) {
	case ParamsOfAppPasswordProviderGetPassword:
		return json.Marshal(struct {
			ParamsOfAppPasswordProviderGetPassword
			Type string `json:"type"`
		}{
			value,
			"GetPassword",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppPasswordProvider %v", pOAPP.ValueEnumType)
	}
}

func (pOAPP *ParamsOfAppPasswordProvider) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "GetPassword":
		var valueEnum ParamsOfAppPasswordProviderGetPassword
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOAPP.ValueEnumType = valueEnum

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppPasswordProvider %v", typeD.Type)
	}

	return nil
}

func (rOAPP *ResultOfAppPasswordProvider) MarshalJSON() ([]byte, error) {
	switch value := (rOAPP.ValueEnumType).(type) {
	case ResultOfAppPasswordProviderGetPassword:
		return json.Marshal(struct {
			ResultOfAppPasswordProviderGetPassword
			Type string `json:"type"`
		}{
			value,
			"GetPassword",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppPasswordProvider %v", rOAPP.ValueEnumType)
	}
}

func (rOAPP *ResultOfAppPasswordProvider) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "GetPassword":
		var valueEnum ResultOfAppPasswordProviderGetPassword
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		rOAPP.ValueEnumType = valueEnum

	default:
		return fmt.Errorf("unsupported type for ResultOfAppPasswordProvider %v", typeD.Type)
	}

	return nil
}

func (ea *EncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	switch value := (ea.ValueEnumType).(type) {
	case EncryptionAlgorithmAESVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			EncryptionAlgorithmAESVariant
		}{"AES", value})
	case EncryptionAlgorithmChaCha20Variant:
		return json.Marshal(struct {
			Type string `json:"type"`
			EncryptionAlgorithmChaCha20Variant
		}{"ChaCha20", value})
	case EncryptionAlgorithmNaclBoxVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			EncryptionAlgorithmNaclBoxVariant
		}{"NaclBox", value})
	case EncryptionAlgorithmNaclSecretBoxVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			EncryptionAlgorithmNaclSecretBoxVariant
		}{"NaclSecretBox", value})
	default:
		return nil, fmt.Errorf("unsupported type for EncryptionAlgorithm %v", ea.ValueEnumType)
	}
}

func (ea *EncryptionAlgorithm) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "AES":
		var valueEnum EncryptionAlgorithmAESVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "ChaCha20":
		var valueEnum EncryptionAlgorithmChaCha20Variant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "NaclBox":
		var valueEnum EncryptionAlgorithmNaclBoxVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "NaclSecretBox":
		var valueEnum EncryptionAlgorithmNaclSecretBoxVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for EncryptionAlgorithm %v", typeD.Type)
	}
	return nil
}

func (bea *BoxEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	switch value := (bea.ValueEnumType).(type) {
	case BoxEncryptionAlgorithmChaCha20Variant:
		return json.Marshal(struct {
			Type string `json:"type"`
			BoxEncryptionAlgorithmChaCha20Variant
		}{"ChaCha20", value})
	case BoxEncryptionAlgorithmNaclBoxVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			BoxEncryptionAlgorithmNaclBoxVariant
		}{"NaclBox", value})
	case BoxEncryptionAlgorithmNaclSecretBoxVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			BoxEncryptionAlgorithmNaclSecretBoxVariant
		}{"NaclSecretBox", value})
	default:
		return nil, fmt.Errorf("unsupported type for BoxEncryptionAlgorithm %v", bea.ValueEnumType)
	}
}

func (bea *BoxEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "ChaCha20":
		var valueEnum BoxEncryptionAlgorithmChaCha20Variant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bea.ValueEnumType = valueEnum
	case "NaclBox":
		var valueEnum BoxEncryptionAlgorithmNaclBoxVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bea.ValueEnumType = valueEnum
	case "NaclSecretBox":
		var valueEnum BoxEncryptionAlgorithmNaclSecretBoxVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bea.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for BoxEncryptionAlgorithm %v", typeD.Type)
	}
	return nil
}

func (pOASB *ParamsOfAppSigningBox) MarshalJSON() ([]byte, error) {
	switch value := (pOASB.ValueEnumType).(type) {
	case ParamsOfAppSigningBoxGetPublicKey:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAppSigningBoxGetPublicKey
		}{"GetPublicKey", value})
	case ParamsOfAppSigningBoxSign:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAppSigningBoxSign
		}{"Sign", value})
	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppSigningBox %v", pOASB.ValueEnumType)
	}
}

func (pOASB *ParamsOfAppSigningBox) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "GetPublicKey":
		var valueEnum ParamsOfAppSigningBoxGetPublicKey
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	case "Sign":
		var valueEnum ParamsOfAppSigningBoxSign
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ParamsOfAppSigningBox %v", typeD.Type)
	}
	return nil
}

func NewParamsOfAppSigningBox(value interface{}) *ParamsOfAppSigningBox {
	return &ParamsOfAppSigningBox{ValueEnumType: value}
}

func (rOASB *ResultOfAppSigningBox) MarshalJSON() ([]byte, error) {
	switch value := (rOASB.ValueEnumType).(type) {
	case ResultOfAppSigningBoxGetPublicKey:
		return json.Marshal(struct {
			Type string `json:"type"`
			ResultOfAppSigningBoxGetPublicKey
		}{"GetPublicKey", value})
	case ResultOfAppSigningBoxSign:
		return json.Marshal(struct {
			Type string `json:"type"`
			ResultOfAppSigningBoxSign
		}{"Sign", value})
	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppSigningBox %v", rOASB.ValueEnumType)
	}
}

func (rOASB *ResultOfAppSigningBox) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "GetPublicKey":
		var valueEnum ResultOfAppSigningBoxGetPublicKey
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		rOASB.ValueEnumType = valueEnum
	case "Sign":
		var valueEnum ResultOfAppSigningBoxSign
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		rOASB.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ResultOfAppSigningBox %v", typeD.Type)
	}
	return nil
}

func NewResultOfAppSigningBox(value interface{}) *ResultOfAppSigningBox {
	return &ResultOfAppSigningBox{ValueEnumType: value}
}

func (pOASB *ParamsOfAppEncryptionBox) MarshalJSON() ([]byte, error) {
	switch value := (pOASB.ValueEnumType).(type) {
	case ParamsOfAppEncryptionBoxGetInfo:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAppEncryptionBoxGetInfo
		}{"GetInfo", value})
	case ParamsOfAppEncryptionBoxEncrypt:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAppEncryptionBoxEncrypt
		}{"Encrypt", value})
	case ParamsOfAppEncryptionBoxDecrypt:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAppEncryptionBoxDecrypt
		}{"Decrypt", value})
	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppEncryptionBox %v", pOASB.ValueEnumType)
	}
}

func (pOASB *ParamsOfAppEncryptionBox) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "GetInfo":
		var valueEnum ParamsOfAppEncryptionBoxGetInfo
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	case "Encrypt":
		var valueEnum ParamsOfAppEncryptionBoxEncrypt
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	case "Decrypt":
		var valueEnum ParamsOfAppEncryptionBoxDecrypt
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ParamsOfAppEncryptionBox %v", typeD.Type)
	}
	return nil
}

// NewParamsOfAppEncryptionBox ...
func NewParamsOfAppEncryptionBox(value interface{}) *ParamsOfAppEncryptionBox {
	return &ParamsOfAppEncryptionBox{ValueEnumType: value}
}

func (pOASB *ResultOfAppEncryptionBox) MarshalJSON() ([]byte, error) {
	switch value := (pOASB.ValueEnumType).(type) {
	case ResultOfAppEncryptionBoxGetInfo:
		return json.Marshal(struct {
			Type string `json:"type"`
			ResultOfAppEncryptionBoxGetInfo
		}{"GetInfo", value})
	case ResultOfAppEncryptionBoxEncrypt:
		return json.Marshal(struct {
			Type string `json:"type"`
			ResultOfAppEncryptionBoxEncrypt
		}{"Encrypt", value})
	case ResultOfAppEncryptionBoxDecrypt:
		return json.Marshal(struct {
			Type string `json:"type"`
			ResultOfAppEncryptionBoxDecrypt
		}{"Decrypt", value})
	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppEncryptionBox %v", pOASB.ValueEnumType)
	}
}

func (pOASB *ResultOfAppEncryptionBox) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}
	switch typeD.Type {
	case "GetInfo":
		var valueEnum ResultOfAppEncryptionBoxGetInfo
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	case "Encrypt":
		var valueEnum ResultOfAppEncryptionBoxEncrypt
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	case "Decrypt":
		var valueEnum ResultOfAppEncryptionBoxDecrypt
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOASB.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ResultOfAppEncryptionBox %v", typeD.Type)
	}
	return nil
}

func NewResultOfAppEncryptionBox(value interface{}) *ResultOfAppEncryptionBox {
	return &ResultOfAppEncryptionBox{ValueEnumType: value}
}
