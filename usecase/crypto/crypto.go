package crypto

import (
	"encoding/json"
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

// Factorize - Performs prime factorization – decomposition of a composite number into a product
// of smaller prime integers (factors).
func (c *crypto) Factorize(poF *domain.ParamsOfFactorize) (*domain.ResultOfFactorize, error) {
	result := new(domain.ResultOfFactorize)
	err := c.client.GetResult("crypto.factorize", poF, result)
	return result, err
}

// ModularPower - Performs modular exponentiation for big integers (base^exponent mod modulus).
func (c *crypto) ModularPower(pOMP *domain.ParamsOfModularPower) (*domain.ResultOfModularPower, error) {
	result := new(domain.ResultOfModularPower)
	err := c.client.GetResult("crypto.modular_power", pOMP, result)
	return result, err
}

// TonCrc16 - Calculates CRC16 using TON algorithm.
func (c *crypto) TonCrc16(pOTC *domain.ParamsOfTonCrc16) (*domain.ResultOfTonCrc16, error) {
	result := new(domain.ResultOfTonCrc16)
	err := c.client.GetResult("crypto.ton_crc16", pOTC, result)
	return result, err
}

// GenerateRandomBytes Generates random byte array of the specified length and returns it in base64 format.
func (c *crypto) GenerateRandomBytes(pOGRB *domain.ParamsOfGenerateRandomBytes) (*domain.ResultOfGenerateRandomBytes, error) {
	result := new(domain.ResultOfGenerateRandomBytes)
	err := c.client.GetResult("crypto.generate_random_bytes", pOGRB, result)
	return result, err
}

// ConvertPublicKeyString - Converts public key to ton safe_format.
func (c *crypto) ConvertPublicKeyString(pOCPTTSF *domain.ParamsOfConvertPublicKeyToTonSafeFormat) (*domain.ResultOfConvertPublicKeyToTonSafeFormat, error) {
	result := new(domain.ResultOfConvertPublicKeyToTonSafeFormat)
	err := c.client.GetResult("crypto.convert_public_key_to_ton_safe_format", pOCPTTSF, result)
	return result, err
}

// GenerateRandomSignKeys - Generates random ed25519 key pair.
func (c *crypto) GenerateRandomSignKeys() (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.generate_random_sign_keys", "{}", result)
	return result, err
}

// Sign - Signs a data using the provided keys.
func (c *crypto) Sign(pOS *domain.ParamsOfSign) (*domain.ResultOfSign, error) {
	result := new(domain.ResultOfSign)
	err := c.client.GetResult("crypto.sign", pOS, result)
	return result, err
}

// VerifySignature - Verifies signed data using the provided public key. Raises error if verification is failed.
func (c *crypto) VerifySignature(pOVS *domain.ParamsOfVerifySignature) (*domain.ResultOfVerifySignature, error) {
	result := new(domain.ResultOfVerifySignature)
	err := c.client.GetResult("crypto.verify_signature", pOVS, result)
	return result, err
}

// Sha256 - Calculates SHA256 hash of the specified data.
func (c *crypto) Sha256(pOH *domain.ParamsOfHash) (*domain.ResultOfHash, error) {
	result := new(domain.ResultOfHash)
	err := c.client.GetResult("crypto.sha256", pOH, result)
	return result, err
}

// Sha512 - Calculates SHA512 hash of the specified data.
func (c *crypto) Sha512(pOH *domain.ParamsOfHash) (*domain.ResultOfHash, error) {
	result := new(domain.ResultOfHash)
	err := c.client.GetResult("crypto.sha512", pOH, result)
	return result, err
}

// Scrypt - Derives key from password and key using scrypt algorithm.
func (c *crypto) Scrypt(sD *domain.ParamsOfScrypt) (*domain.ResultOfScrypt, error) {
	result := new(domain.ResultOfScrypt)
	err := c.client.GetResult("crypto.scrypt", sD, result)
	return result, err
}

// NaclSignKeypairFromSecretKey - Generates a key pair for signing from the secret key.
func (c *crypto) NaclSignKeypairFromSecretKey(pONSKPFC *domain.ParamsOfNaclSignKeyPairFromSecret) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_sign_keypair_from_secret_key", pONSKPFC, result)
	return result, err
}

// NaclSign - Signs data using the signer's secret key.
func (c *crypto) NaclSign(pONS *domain.ParamsOfNaclSign) (*domain.ResultOfNaclSign, error) {
	result := new(domain.ResultOfNaclSign)
	err := c.client.GetResult("crypto.nacl_sign", pONS, result)
	return result, err
}

// NaclSignOpen - Verifies the signature and returns the unsigned message.
//Verifies the signature in signed using the signer's public key public and returns the message unsigned.
//If the signature fails verification, crypto_sign_open raises an exception.
func (c *crypto) NaclSignOpen(pONSO *domain.ParamsOfNaclSignOpen) (*domain.ResultOfNaclSignOpen, error) {
	result := new(domain.ResultOfNaclSignOpen)
	err := c.client.GetResult("crypto.nacl_sign_open", pONSO, result)
	return result, err
}

// NaclSignDetached - Signs the message using the secret key and returns a signature.
//Signs the message unsigned using the secret key secret and returns a signature signature.
func (c *crypto) NaclSignDetached(pONS *domain.ParamsOfNaclSign) (*domain.ResultOfNaclSignDetached, error) {
	result := new(domain.ResultOfNaclSignDetached)
	err := c.client.GetResult("crypto.nacl_sign_detached", pONS, result)
	return result, err
}

// NaclSignDetachedVerify - Verifies the signature with public key and unsigned data.
func (c *crypto) NaclSignDetachedVerify(pONSDV *domain.ParamsOfNaclSignDetachedVerify) (*domain.ResultOfNaclSignDetachedVerify, error) {
	result := new(domain.ResultOfNaclSignDetachedVerify)
	err := c.client.GetResult("crypto.nacl_sign_detached_verify", pONSDV, result)
	return result, err
}

// NaclBoxKeypair - Generates a random NaCl key pair.
func (c *crypto) NaclBoxKeypair() (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_box_keypair", "{}", result)
	return result, err
}

// NaclBoxKeypairFromSecretKey - Generates key pair from a secret key.
func (c *crypto) NaclBoxKeypairFromSecretKey(pONKPFS *domain.ParamsOfNaclBoxKeyPairFromSecret) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.nacl_box_keypair_from_secret_key", pONKPFS, result)
	return result, err
}

// NaclBox - Public key authenticated encryption. Encrypt and authenticate a message using
// the senders secret key, the receivers public key, and a nonce.
func (c *crypto) NaclBox(pONB *domain.ParamsOfNaclBox) (*domain.ResultOfNaclBox, error) {
	result := new(domain.ResultOfNaclBox)
	err := c.client.GetResult("crypto.nacl_box", pONB, result)
	return result, err
}

// NaclBoxOpen - Decrypt and verify the cipher text using the recievers secret key, the senders public
// key, and the nonce.
func (c *crypto) NaclBoxOpen(pONBO *domain.ParamsOfNaclBoxOpen) (*domain.ResultOfNaclBoxOpen, error) {
	result := new(domain.ResultOfNaclBoxOpen)
	err := c.client.GetResult("crypto.nacl_box_open", pONBO, result)
	return result, err
}

// NaclSecretBox - Encrypt and authenticate message using nonce and secret key.
func (c *crypto) NaclSecretBox(pONSB *domain.ParamsOfNaclSecretBox) (*domain.ResultOfNaclBox, error) {
	result := new(domain.ResultOfNaclBox)
	err := c.client.GetResult("crypto.nacl_secret_box", pONSB, result)
	return result, err
}

// NaclSecretBoxOpen - Decrypts and verifies cipher text using nonce and secret key.
func (c *crypto) NaclSecretBoxOpen(pONSBO *domain.ParamsOfNaclSecretBoxOpen) (*domain.ResultOfNaclBoxOpen, error) {
	result := new(domain.ResultOfNaclBoxOpen)
	err := c.client.GetResult("crypto.nacl_secret_box_open", pONSBO, result)
	return result, err
}

// MnemonicWords - Prints the list of words from the specified dictionary.
func (c *crypto) MnemonicWords(pOMW *domain.ParamsOfMnemonicWords) (*domain.ResultOfMnemonicWords, error) {
	result := new(domain.ResultOfMnemonicWords)
	err := c.client.GetResult("crypto.mnemonic_words", pOMW, result)
	return result, err
}

// MnemonicFromRandom - Generates a random mnemonic from the specified dictionary and word count.
func (c *crypto) MnemonicFromRandom(pOMFR *domain.ParamsOfMnemonicFromRandom) (*domain.ResultOfMnemonicFromRandom, error) {
	result := new(domain.ResultOfMnemonicFromRandom)
	err := c.client.GetResult("crypto.mnemonic_from_random", pOMFR, result)
	return result, err
}

// MnemonicFromEntropy - Generates mnemonic from pre-generated entropy.
func (c *crypto) MnemonicFromEntropy(pOMFE *domain.ParamsOfMnemonicFromEntropy) (*domain.ResultOfMnemonicFromEntropy, error) {
	result := new(domain.ResultOfMnemonicFromEntropy)
	err := c.client.GetResult("crypto.mnemonic_from_entropy", pOMFE, result)
	return result, err
}

// MnemonicVerify - The phrase supplied will be checked for word length and validated according to the
// checksum specified in BIP0039.
func (c *crypto) MnemonicVerify(pOMV *domain.ParamsOfMnemonicVerify) (*domain.ResultOfMnemonicVerify, error) {
	result := new(domain.ResultOfMnemonicVerify)
	err := c.client.GetResult("crypto.mnemonic_verify", pOMV, result)
	return result, err
}

// MnemonicDeriveSignKeys - Validates the seed phrase, generates master key and then derives the key pair from
// the master key and the specified path.
func (c *crypto) MnemonicDeriveSignKeys(pOMDSK *domain.ParamsOfMnemonicDeriveSignKeys) (*domain.KeyPair, error) {
	result := new(domain.KeyPair)
	err := c.client.GetResult("crypto.mnemonic_derive_sign_keys", pOMDSK, result)
	return result, err
}

// HdkeyXprvFromMnemonic - Generates an extended master private key that will be the root for all the derived keys.
func (c *crypto) HdkeyXprvFromMnemonic(pOHKXFM *domain.ParamsOfHDKeyXPrvFromMnemonic) (*domain.ResultOfHDKeyXPrvFromMnemonic, error) {
	result := new(domain.ResultOfHDKeyXPrvFromMnemonic)
	err := c.client.GetResult("crypto.hdkey_xprv_from_mnemonic", pOHKXFM, result)
	return result, err
}

// HdkeyDeriveFromXprv - Returns extended private key derived from the specified extended private key and child index.
func (c *crypto) HdkeyDeriveFromXprv(hdP *domain.ParamsOfHDKeyDeriveFromXPrv) (*domain.ResultOfHDKeyDeriveFromXPrv, error) {
	result := new(domain.ResultOfHDKeyDeriveFromXPrv)
	err := c.client.GetResult("crypto.hdkey_derive_from_xprv", hdP, result)
	return result, err
}

// HdkeyDeriveFromXprvPath - Derives the extended private key from the specified key and path.
func (c *crypto) HdkeyDeriveFromXprvPath(hdPD *domain.ParamsOfHDKeyDeriveFromXPrvPath) (*domain.ResultOfHDKeyDeriveFromXPrvPath, error) {
	result := new(domain.ResultOfHDKeyDeriveFromXPrvPath)
	err := c.client.GetResult("crypto.hdkey_derive_from_xprv_path", hdPD, result)
	return result, err
}

// HdkeySecretFromXprv - Extracts the private key from the serialized extended private key.
func (c *crypto) HdkeySecretFromXprv(pOHKSFXP *domain.ParamsOfHDKeySecretFromXPrv) (*domain.ResultOfHDKeySecretFromXPrv, error) {
	result := new(domain.ResultOfHDKeySecretFromXPrv)
	err := c.client.GetResult("crypto.hdkey_secret_from_xprv", pOHKSFXP, result)
	return result, err
}

// HdkeyPublicFromXprv - Extracts the public key from the serialized extended private key.
func (c *crypto) HdkeyPublicFromXprv(pOHKPFXP *domain.ParamsOfHDKeyPublicFromXPrv) (*domain.ResultOfHDKeyPublicFromXPrv, error) {
	result := new(domain.ResultOfHDKeyPublicFromXPrv)
	err := c.client.GetResult("crypto.hdkey_public_from_xprv", pOHKPFXP, result)
	return result, err
}

// Chacha20 - Performs symmetric chacha20 encryption.
func (c *crypto) Chacha20(pOFCC *domain.ParamsOfChaCha20) (*domain.ResultOfChaCha20, error) {
	result := new(domain.ResultOfChaCha20)
	err := c.client.GetResult("crypto.hdkey_secret_from_xprv", pOFCC, result)
	return result, err
}

// RegisterSigningBox - Register an application implemented signing box.
func (c *crypto) RegisterSigningBox(app domain.AppSigningBox) (*domain.RegisteredSigningBox, error) {
	result := new(domain.RegisteredSigningBox)
	responses, err := c.client.Request("crypto.register_signing_box", nil)
	if err != nil {
		return nil, err
	}
	response := <-responses
	if response.Code == 1 {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == 3 {
				c.appRequestCryptoRegisterSigningBox(r.Data, app)
			}
			if r.Code == 4 {
				c.appNotifyCryptoRegisterSigningBox(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *crypto) appRequestCryptoRegisterSigningBox(payload []byte, app domain.AppSigningBox) {
	var appRequest domain.ParamsOfAppRequest
	var appParams domain.ParamsOfAppSigningBox
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	appResponse, err := app.Request(appParams)
	appRequestResult := &domain.AppRequestResult{}
	if err != nil {
		appRequestResult.Type = domain.AppRequestResultTypeError
		appRequestResult.Text = err.Error()
	} else {
		appRequestResult.Type = domain.AppRequestResultTypeOk
		appRequestResult.Result, _ = json.Marshal(appResponse)
	}
	err = c.client.ResolveAppRequest(&domain.ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (c *crypto) appNotifyCryptoRegisterSigningBox(payload []byte, app domain.AppSigningBox) {
	var appParams domain.ParamsOfAppSigningBox
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// GetSigningBox - Creates a default signing box implementation.
func (c *crypto) GetSigningBox(keypair *domain.KeyPair) (*domain.RegisteredSigningBox, error) {
	result := new(domain.RegisteredSigningBox)
	err := c.client.GetResult("crypto.get_signing_box", keypair, result)
	return result, err
}

// SigningBoxGetPublicKey - Returns public key of signing key pair.
func (c *crypto) SigningBoxGetPublicKey(keypair *domain.RegisteredSigningBox) (*domain.ResultOfSigningBoxGetPublicKey, error) {
	result := new(domain.ResultOfSigningBoxGetPublicKey)
	err := c.client.GetResult("crypto.signing_box_get_public_key", keypair, result)
	return result, err
}

// SigningBoxSign - Returns signed user data.
func (c *crypto) SigningBoxSign(pOSBS *domain.ParamsOfSigningBoxSign) (*domain.ResultOfSigningBoxSign, error) {
	result := new(domain.ResultOfSigningBoxSign)
	err := c.client.GetResult("crypto.signing_box_sign", pOSBS, result)
	return result, err
}

// SigningBoxSign - Removes signing box from SDK.
func (c *crypto) RemoveSigningBox(rSB *domain.RegisteredSigningBox) error {
	_, err := c.client.GetResponse("client.resolve_app_request", rSB)
	return err
}
