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

// CryptoErrorCode ...
var CryptoErrorCode map[string]int

type (
	// CryptoBoxHandle ...
	CryptoBoxHandle int

	// SigningBoxHandle ...
	SigningBoxHandle int

	// EncryptionBoxHandle ...
	EncryptionBoxHandle int

	// EncryptionBoxInfo - Encryption box information.
	EncryptionBoxInfo struct {
		HDPath    string          `json:"hdpath,omitempty"`
		Algorithm string          `json:"algorithm,omitempty"`
		Options   json.RawMessage `json:"options,omitempty"`
		Public    json.RawMessage `json:"public,omitempty"`
	}

	// EncryptionAlgorithm ...
	EncryptionAlgorithm struct {
		ValueEnumType interface{}
	}

	// BoxEncryptionAlgorithm ...
	BoxEncryptionAlgorithm struct {
		ValueEnumType interface{}
	}

	// CipherMode is type string
	CipherMode string

	// AesParams ...
	AesParams struct {
		Mode CipherMode `json:"mode"`
		Key  string     `json:"key"`
		IV   string     `json:"iv,omitempty"`
	}

	// AesInfo ...
	AesInfo struct {
		Mode CipherMode `json:"mode"`
		IV   string     `json:"iv,omitempty"`
	}

	// ChaCha20ParamsEB ...
	ChaCha20ParamsEB struct {
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	// NaclBoxParamsEB ...
	NaclBoxParamsEB struct {
		TheirPublic string `json:"their_public"`
		Secret      string `json:"secret"`
		Nonce       string `json:"nonce"`
	}

	// NaclSecretBoxParamsEB ...
	NaclSecretBoxParamsEB struct {
		Key   string `json:"key"`
		Nonce string `json:"nonce"`
	}

	// CryptoBoxSecret ...
	CryptoBoxSecret struct {
		ValueEnumType interface{}
	}

	// CryptoBoxSecretRandomSeedPhrase ...
	CryptoBoxSecretRandomSeedPhrase struct {
		Dictionary int `json:"dictionary"`
		WordCount  int `json:"wordcount"`
	}

	// CryptoBoxSecretPredefinedSeedPhrase ...
	CryptoBoxSecretPredefinedSeedPhrase struct {
		Phrase     string `json:"phrase"`
		Dictionary int    `json:"dictionary"`
		WordCount  int    `json:"wordcount"`
	}

	// CryptoBoxSecretEncryptedSecret ...
	CryptoBoxSecretEncryptedSecret struct {
		EncryptedSecret string `json:"encrypted_secret"`
	}

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
		Unsigned string   `json:"unsigned"`
		Keys     *KeyPair `json:"keys"`
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
		Dictionary *int `json:"dictionary,omitempty"`
	}

	// ResultOfMnemonicWords ...
	ResultOfMnemonicWords struct {
		Words string `json:"words"`
	}

	// ParamsOfMnemonicFromRandom ...
	ParamsOfMnemonicFromRandom struct {
		Dictionary *int `json:"dictionary,omitempty"`
		WordCount  *int `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicFromRandom ...
	ResultOfMnemonicFromRandom struct {
		Phrase string `json:"phrase"`
	}

	// ParamsOfMnemonicFromEntropy ...
	ParamsOfMnemonicFromEntropy struct {
		Entropy    string `json:"entropy"`
		Dictionary *int   `json:"dictionary,omitempty"`
		WordCount  *int   `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicFromEntropy ...
	ResultOfMnemonicFromEntropy struct {
		Phrase string `json:"phrase"`
	}

	// ParamsOfMnemonicVerify ...
	ParamsOfMnemonicVerify struct {
		Phrase     string `json:"phrase"`
		Dictionary *int   `json:"dictionary,omitempty"`
		WordCount  *int   `json:"word_count,omitempty"`
	}

	// ResultOfMnemonicVerify ...
	ResultOfMnemonicVerify struct {
		Valid bool `json:"valid"`
	}

	// ParamsOfMnemonicDeriveSignKeys ...
	ParamsOfMnemonicDeriveSignKeys struct {
		Phrase     string `json:"phrase"`
		Path       string `json:"path,omitempty"`
		Dictionary *int   `json:"dictionary,omitempty"`
		WordCount  *int   `json:"word_count,omitempty"`
	}

	// ParamsOfHDKeyXPrvFromMnemonic ...
	ParamsOfHDKeyXPrvFromMnemonic struct {
		Phrase     string `json:"phrase"`
		Dictionary *int   `json:"dictionary,omitempty"`
		WordCount  *int   `json:"word_count,omitempty"`
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

	// ParamsOfCreateCryptoBox ...
	ParamsOfCreateCryptoBox struct {
		SecretEncryptionSalt string          `json:"secret_encryption_salt"`
		Secret               CryptoBoxSecret `json:"secret"`
	}

	// RegisteredCryptoBox ...
	RegisteredCryptoBox struct {
		Handle CryptoBoxHandle `json:"handle"`
	}

	// ResultOfGetCryptoBoxInfo ...
	ResultOfGetCryptoBoxInfo struct {
		EncryptedSecret string `json:"encrypted_secret"`
	}

	// ResultOfGetCryptoBoxSeedPhrase ...
	ResultOfGetCryptoBoxSeedPhrase struct {
		Phrase     string `json:"phrase"`
		Dictionary int    `json:"dictionary"`
		Wordcount  int    `json:"wordcount"`
	}

	// ParamsOfGetSigningBoxFromCryptoBox ...
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

	// RegisteredSigningBox ...
	RegisteredSigningBox struct {
		Handle SigningBoxHandle `json:"handle"`
	}

	// ParamsOfAppSigningBox ...
	ParamsOfAppSigningBox struct {
		ValueEnumType interface{}
	}

	// ParamsOfAppSigningBoxGetPublicKey ...
	ParamsOfAppSigningBoxGetPublicKey struct{}

	// ParamsOfAppSigningBoxSign ...
	ParamsOfAppSigningBoxSign struct {
		Unsigned string `json:"unsigned"`
	}

	// ParamsOfAppPasswordProvider ...
	ParamsOfAppPasswordProvider struct {
		ValueEnumType interface{}
	}

	// ParamsOfAppPasswordProviderGetPassword ...
	ParamsOfAppPasswordProviderGetPassword struct {
		EncryptionPublicKey string `json:"encryption_public_key"`
	}

	// ResultOfAppPasswordProvider ...
	ResultOfAppPasswordProvider struct {
		ValueEnumType interface{}
	}

	// ResultOfAppPasswordProviderGetPassword ...
	ResultOfAppPasswordProviderGetPassword struct {
		EncryptedPassword   string `json:"encrypted_password"`
		AppEncryptionPubkey string `json:"app_encryption_pubkey"`
	}

	// ResultOfAppSigningBox ...
	ResultOfAppSigningBox struct {
		ValueEnumType interface{}
	}

	// ResultOfAppSigningBoxGetPublicKey ...
	ResultOfAppSigningBoxGetPublicKey struct {
		PublicKey string `json:"public_key"`
	}

	// ResultOfAppSigningBoxSign ...
	ResultOfAppSigningBoxSign struct {
		Signature string `json:"signature"`
	}

	// ResultOfSigningBoxGetPublicKey ...
	ResultOfSigningBoxGetPublicKey struct {
		PubKey string `json:"PubKey"`
	}

	// ParamsOfSigningBoxSign ...
	ParamsOfSigningBoxSign struct {
		SigningBox SigningBoxHandle `json:"signing_box"`
		Unsigned   string           `json:"unsigned"`
	}

	// ResultOfSigningBoxSign ...
	ResultOfSigningBoxSign struct {
		Signature string `json:"signature"`
	}

	// RegisteredEncryptionBox ...
	RegisteredEncryptionBox struct {
		Handle EncryptionBoxHandle `json:"handle"`
	}

	// ParamsOfAppEncryptionBox - Encryption box callbacks.
	ParamsOfAppEncryptionBox struct {
		ValueEnumType interface{}
	}

	// ParamsOfAppEncryptionBoxGetInfo ...
	ParamsOfAppEncryptionBoxGetInfo struct{}

	// ParamsOfAppEncryptionBoxEncrypt ...
	ParamsOfAppEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	// ParamsOfAppEncryptionBoxDecrypt ...
	ParamsOfAppEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	// ResultOfAppEncryptionBox - Returning values from signing box callbacks.
	ResultOfAppEncryptionBox struct {
		ValueEnumType interface{}
	}

	// ResultOfAppEncryptionBoxGetInfo ...
	ResultOfAppEncryptionBoxGetInfo struct {
		Info EncryptionBoxInfo `json:"info"`
	}

	// ResultOfAppEncryptionBoxEncrypt ...
	ResultOfAppEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	// ResultOfAppEncryptionBoxDecrypt ...
	ResultOfAppEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	// ParamsOfEncryptionBoxGetInfo ...
	ParamsOfEncryptionBoxGetInfo struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
	}

	// ResultOfEncryptionBoxGetInfo ...
	ResultOfEncryptionBoxGetInfo struct {
		Info EncryptionBoxInfo `json:"info"`
	}

	// ParamsOfEncryptionBoxEncrypt ...
	ParamsOfEncryptionBoxEncrypt struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
		Data          string              `json:"data"`
	}

	// ResultOfEncryptionBoxEncrypt ...
	ResultOfEncryptionBoxEncrypt struct {
		Data string `json:"data"`
	}

	// ParamsOfEncryptionBoxDecrypt ...
	ParamsOfEncryptionBoxDecrypt struct {
		EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
		Data          string              `json:"data"`
	}

	// ResultOfEncryptionBoxDecrypt ...
	ResultOfEncryptionBoxDecrypt struct {
		Data string `json:"data"`
	}

	// ParamsOfCreateEncryptionBox ...
	ParamsOfCreateEncryptionBox struct {
		Algorithm EncryptionAlgorithm `json:"algorithm"`
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

// NewDefaultParamsOfMnemonicWords ...
func NewDefaultParamsOfMnemonicWords() *ParamsOfMnemonicWords {
	return &ParamsOfMnemonicWords{Dictionary: util.IntToPointerInt(DefaultDictionary)}
}

// NewDefaultParamsOfMnemonicFromRandom ...
func NewDefaultParamsOfMnemonicFromRandom() *ParamsOfMnemonicFromRandom {
	return &ParamsOfMnemonicFromRandom{Dictionary: util.IntToPointerInt(DefaultDictionary), WordCount: util.IntToPointerInt(DefaultWordCount)}
}

// NewDefaultParamsOfMnemonicFromEntropy ...
func NewDefaultParamsOfMnemonicFromEntropy() *ParamsOfMnemonicFromEntropy {
	return &ParamsOfMnemonicFromEntropy{Entropy: "", Dictionary: util.IntToPointerInt(DefaultDictionary), WordCount: util.IntToPointerInt(DefaultWordCount)}
}

// NewDefaultParamsOfMnemonicVerify ...
func NewDefaultParamsOfMnemonicVerify() *ParamsOfMnemonicVerify {
	return &ParamsOfMnemonicVerify{Phrase: "", Dictionary: util.IntToPointerInt(DefaultDictionary), WordCount: util.IntToPointerInt(DefaultWordCount)}
}

// NewDefaultParamsOfMnemonicDeriveSignKeys ...
func NewDefaultParamsOfMnemonicDeriveSignKeys() *ParamsOfMnemonicDeriveSignKeys {
	return &ParamsOfMnemonicDeriveSignKeys{Phrase: "", Dictionary: util.IntToPointerInt(DefaultDictionary), WordCount: util.IntToPointerInt(DefaultWordCount)}
}

// NewDefaultParamsOfHDKeyXPrvFromMnemonic ...
func NewDefaultParamsOfHDKeyXPrvFromMnemonic() *ParamsOfHDKeyXPrvFromMnemonic {
	return &ParamsOfHDKeyXPrvFromMnemonic{Phrase: "", Dictionary: util.IntToPointerInt(DefaultDictionary), WordCount: util.IntToPointerInt(DefaultWordCount)}
}

func (cbs *CryptoBoxSecret) MarshalJSON() ([]byte, error) {
	switch value := (cbs.ValueEnumType).(type) {
	case CryptoBoxSecretRandomSeedPhrase:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretRandomSeedPhrase
		}{"RandomSeedPhrase", value})
	case CryptoBoxSecretPredefinedSeedPhrase:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretPredefinedSeedPhrase
		}{"PredefinedSeedPhrase", value})
	case CryptoBoxSecretEncryptedSecret:
		return json.Marshal(struct {
			Type string `json:"type"`
			CryptoBoxSecretEncryptedSecret
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
		var valueEnum CryptoBoxSecretRandomSeedPhrase
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		cbs.ValueEnumType = valueEnum
	case "PredefinedSeedPhrase":
		var valueEnum CryptoBoxSecretPredefinedSeedPhrase
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		cbs.ValueEnumType = valueEnum
	case "EncryptedSecret":
		var valueEnum CryptoBoxSecretEncryptedSecret
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
	case AesParams:
		return json.Marshal(struct {
			Type string `json:"type"`
			AesParams
		}{"AES", value})
	case ChaCha20ParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			ChaCha20ParamsEB
		}{"ChaCha20", value})
	case NaclBoxParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			NaclBoxParamsEB
		}{"NaclBox", value})
	case NaclSecretBoxParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			NaclSecretBoxParamsEB
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
		var valueEnum AesParams
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "ChaCha20":
		var valueEnum ChaCha20ParamsEB
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "NaclBox":
		var valueEnum NaclBoxParamsEB
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		ea.ValueEnumType = valueEnum
	case "NaclSecretBox":
		var valueEnum NaclSecretBoxParamsEB
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
	case ChaCha20ParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			ChaCha20ParamsEB
		}{"ChaCha20", value})
	case NaclBoxParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			NaclBoxParamsEB
		}{"NaclBox", value})
	case NaclSecretBoxParamsEB:
		return json.Marshal(struct {
			Type string `json:"type"`
			NaclSecretBoxParamsEB
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
		var valueEnum ChaCha20ParamsEB
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bea.ValueEnumType = valueEnum
	case "NaclBox":
		var valueEnum NaclBoxParamsEB
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bea.ValueEnumType = valueEnum
	case "NaclSecretBox":
		var valueEnum NaclSecretBoxParamsEB
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

// NewParamsOfAppSigningBox ...
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

// NewResultOfAppSigningBox ...
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

// NewResultOfAppEncryptionBox ...
func NewResultOfAppEncryptionBox(value interface{}) *ResultOfAppEncryptionBox {
	return &ResultOfAppEncryptionBox{ValueEnumType: value}
}
