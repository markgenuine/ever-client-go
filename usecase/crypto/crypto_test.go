package crypto

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/util"
	"github.com/stretchr/testify/assert"
)

type AppSigningBoxTest struct {
	Public  string
	Private string
}

func (a *AppSigningBoxTest) GetPublicKey() (domain.ResultOfAppSigningBoxGetPublicKey, error) {
	return domain.ResultOfAppSigningBoxGetPublicKey{PublicKey: a.Public}, nil
}

func (a *AppSigningBoxTest) Sign(sign domain.ParamsOfAppSigningBoxSign) (domain.ResultOfAppSigningBoxSign, error) {
	privByte, err := hex.DecodeString(a.Private)
	if err != nil {
		return domain.ResultOfAppSigningBoxSign{}, err
	}
	privKey := ed25519.NewKeyFromSeed(privByte)
	data, err := base64.StdEncoding.DecodeString(sign.Unsigned)
	signature := hex.EncodeToString(ed25519.Sign(privKey, data))
	return domain.ResultOfAppSigningBoxSign{Signature: signature}, nil
}

func TestCrypto(t *testing.T) {
	configConn := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
	clientConn, err := client.NewClientGateway(configConn)
	assert.Equal(t, nil, err)
	defer clientConn.Destroy()

	cryptoUC := crypto{
		config: configConn,
		client: clientConn,
	}
	defer cryptoUC.client.Destroy()

	mnemTest := "abuse boss fly battle rubber wasp afraid hamster guide essence vibrant tattoo"
	masterXPrv := "xprv9s21ZrQH143K25JhKqEwvJW7QAiVvkmi4WRenBZanA6kxHKtKAQQKwZG65kCyW5jWJ8NY9e3GkRoistUjjcpHNsGBUv94istDPXvqGNuWpC"
	lenMnem := domain.WordCountList()
	dictMnem := domain.DictionaryList()
	unsigned := base64.StdEncoding.EncodeToString([]byte("Test Message"))

	t.Run("TestSha256", func(t *testing.T) {
		params := &domain.ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))}
		result, err := cryptoUC.Sha256(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", result.Hash)

		params.Data = base64.StdEncoding.EncodeToString(util.FromHex("544f4e206973206f757220667574757265"))
		result, err = cryptoUC.Sha256(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", result.Hash)

		params.Data = "VE9OIGlzIG91ciBmdXR1cmU="
		result, err = cryptoUC.Sha256(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", result.Hash)
	})

	t.Run("TestSha512", func(t *testing.T) {
		params := &domain.ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))}
		result, err := cryptoUC.Sha512(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "4c52dd4cefc68319bac5e97c1f0d18ae8194fb0dd8d9e090ba8376834a0756175a9a736d1e69cb1a58d25c3d554b02a2b8ed9c3ae5cbeeccc3277746a363a434", result.Hash)
	})

	t.Run("TestHDKeyXprvFromMnemonic", func(t *testing.T) {
		params := &domain.ParamsOfHDKeyXPrvFromMnemonic{Phrase: mnemTest, Dictionary: dictMnem["ENGLISH"], WordCount: util.IntToPointerInt(12)}
		result, err := cryptoUC.HDKeyXprvFromMnemonic(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, masterXPrv, result.Xprv)
	})

	t.Run("TestHDKeyXprvSecrets", func(t *testing.T) {
		params := &domain.ParamsOfHDKeySecretFromXPrv{Xprv: masterXPrv}
		result, err := cryptoUC.HDKeySecretFromXprv(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "0c91e53128fa4d67589d63a6c44049c1068ec28a63069a55ca3de30c57f8b365", result.Secret)

		params.Xprv = ""
		result, err = cryptoUC.HDKeySecretFromXprv(params)
		assert.NotNil(t, err)
	})

	t.Run("TestHDKeyPublicFromXprv", func(t *testing.T) {
		params := &domain.ParamsOfHDKeyPublicFromXPrv{Xprv: masterXPrv}
		result, err := cryptoUC.HDKeyPublicFromXprv(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "7b70008d0c40992283d488b1046739cf827afeabf647a5f07c4ad1e7e45a6f89", result.Public)
	})

	t.Run("TestHDkeyXprvDerive", func(t *testing.T) {
		params := &domain.ParamsOfHDKeyDeriveFromXPrv{Xprv: masterXPrv, ChildIndex: 0, Hardened: false}
		result, err := cryptoUC.HDKeyDeriveFromXprv(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "xprv9uZwtSeoKf1swgAkVVCEUmC2at6t7MCJoHnBbn1MWJZyxQ4cySkVXPyNh7zjf9VjsP4vEHDDD2a6R35cHubg4WpzXRzniYiy8aJh1gNnBKv", result.Xprv)

		params.ChildIndex = -1
		result, err = cryptoUC.HDKeyDeriveFromXprv(params)
		assert.NotNil(t, err)
	})

	t.Run("TestHDKeyXprvDerivePath", func(t *testing.T) {
		params := &domain.ParamsOfHDKeyDeriveFromXPrvPath{Xprv: masterXPrv, Path: "m/44'/60'/0'/0'"}
		result, err := cryptoUC.HDKeyDeriveFromXprvPath(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "xprvA1KNMo63UcGjmDF1bX39Cw2BXGUwrwMjeD5qvQ3tA3qS3mZQkGtpf4DHq8FDLKAvAjXsYGLHDP2dVzLu9ycta8PXLuSYib2T3vzLf3brVgZ", result.Xprv)

		params.Path = "m/"
		result, err = cryptoUC.HDKeyDeriveFromXprvPath(params)
		assert.NotNil(t, err)
	})

	t.Run("TestConvertPublicKeyString", func(t *testing.T) {
		params := &domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: "06117f59ade83e097e0fb33e5d29e8735bda82b3bf78a015542aaa853bb69600"}
		result, err := cryptoUC.ConvertPublicKeyString(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuYGEX9Zreg-CX4Psz5dKehzW9qCs794oBVUKqqFO7aWAOTD", result.TonPublicKey)
	})

	t.Run("TestGenerateRandomSignKeys", func(t *testing.T) {
		result, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		assert.Equal(t, 64, len(result.Public))
		assert.Equal(t, 64, len(result.Secret))
		assert.NotEqual(t, result.Secret, result.Public)
	})

	t.Run("TestSignAndVerify", func(t *testing.T) {
		keys := &domain.KeyPair{Public: "1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e", Secret: "56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f"}

		// Sign message
		signParams := &domain.ParamsOfSign{Unsigned: unsigned, Keys: keys}
		signed, err := cryptoUC.Sign(signParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "+wz+QO6l1slgZS5s65BNqKcu4vz24FCJz4NSAxef9lu0jFfs8x3PzSZRC+pn5k8+aJi3xYMA3BQzglQmjK3hA1Rlc3QgTWVzc2FnZQ==", signed.Signed)
		assert.Equal(t, "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade103", signed.Signature)

		// Verify signature
		verifyParams := &domain.ParamsOfVerifySignature{Signed: signed.Signed, Public: keys.Public}
		verified, err := cryptoUC.VerifySignature(verifyParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, unsigned, verified.Unsigned)
		resultUnsignedDec, err := base64.StdEncoding.DecodeString(unsigned)
		assert.Equal(t, nil, err)
		resultVerifiedUnsignedDec, err := base64.StdEncoding.DecodeString(verified.Unsigned)
		assert.Equal(t, nil, err)
		assert.Equal(t, resultUnsignedDec, resultVerifiedUnsignedDec)

		signParams.Keys = &domain.KeyPair{Public: "1", Secret: "2"}
		signed, err = cryptoUC.Sign(signParams)
		assert.NotNil(t, err)

		verifyParams.Signed = "simple"
		verified, err = cryptoUC.VerifySignature(verifyParams)
		assert.NotNil(t, err)
	})

	t.Run("TestModularPower", func(t *testing.T) {
		params := &domain.ParamsOfModularPower{Base: "0123456789ABCDEF", Exponent: "0123", Modulus: "01234567"}
		result, err := cryptoUC.ModularPower(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "63bfdf", result.ModularPower)

		params.Base = "1"
		params.Modulus = "0.2"
		result, err = cryptoUC.ModularPower(params)
		assert.NotNil(t, err)
	})

	t.Run("TestFactorize", func(t *testing.T) {
		params := &domain.ParamsOfFactorize{Composite: "17ED48941A08F981"}
		result, err := cryptoUC.Factorize(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "494C553B", result.Factors[0])
		assert.Equal(t, "53911073", result.Factors[1])

		params.Composite = "a3"
		result, err = cryptoUC.Factorize(params)
		assert.NotNil(t, err)
	})

	t.Run("TestTonCrc16", func(t *testing.T) {
		params := &domain.ParamsOfTonCrc16{Data: base64.StdEncoding.EncodeToString(util.FromHex("0123456789abcdef"))}
		result, err := cryptoUC.TonCrc16(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, 43349, result.Crc)

		params.Data = "0=="
		result, err = cryptoUC.TonCrc16(params)
		assert.NotNil(t, err)
	})

	t.Run("TestGenerateRandomBytes", func(t *testing.T) {
		params := &domain.ParamsOfGenerateRandomBytes{Length: 32}
		result, err := cryptoUC.GenerateRandomBytes(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, 44, len(result.Bytes))
		resultDecode, err := base64.StdEncoding.DecodeString(result.Bytes)
		assert.Equal(t, nil, err)
		assert.Equal(t, 32, len(resultDecode))
	})

	t.Run("TestMnemonicWords", func(t *testing.T) {
		result, err := cryptoUC.MnemonicWords(domain.NewDefaultParamsOfMnemonicWords())
		assert.Equal(t, nil, err)
		assert.Equal(t, 2048, len(strings.Fields(result.Words)))

		_, err = cryptoUC.MnemonicWords(&domain.ParamsOfMnemonicWords{Dictionary: util.IntToPointerInt(100)})
		assert.NotNil(t, err)
	})

	t.Run("TestMnemonicFromRandom", func(t *testing.T) {
		for vall := range lenMnem {
			for _, vald := range dictMnem {
				params := &domain.ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: util.IntToPointerInt(vall)}
				result, err := cryptoUC.MnemonicFromRandom(params)
				assert.Equal(t, nil, err)
				assert.Equal(t, vall, len(strings.Fields(result.Phrase)))
			}
		}
	})

	t.Run("TestMnemonicFromEntropy", func(t *testing.T) {
		defaultEntropy := domain.NewDefaultParamsOfMnemonicFromEntropy()
		defaultEntropy.Entropy = "00112233445566778899AABBCCDDEEFF"
		result, err := cryptoUC.MnemonicFromEntropy(defaultEntropy)
		assert.Equal(t, nil, err)
		assert.Equal(t, "abandon math mimic master filter design carbon crystal rookie group knife young", result.Phrase)

		defaultEntropy.Entropy = "01"
		_, err = cryptoUC.MnemonicFromEntropy(defaultEntropy)
		assert.NotNil(t, err)
	})

	t.Run("TestMnemonicVerify", func(t *testing.T) {
		for vall := range lenMnem {
			for _, vald := range dictMnem {
				result1, err := cryptoUC.MnemonicFromRandom(&domain.ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: util.IntToPointerInt(vall)})
				assert.Equal(t, nil, err)

				result2, err := cryptoUC.MnemonicVerify(&domain.ParamsOfMnemonicVerify{Phrase: result1.Phrase, Dictionary: vald, WordCount: util.IntToPointerInt(vall)})
				assert.Equal(t, nil, err)
				assert.True(t, result2.Valid)
			}
		}
	})

	t.Run("TestMnemonicDeriveSignKeys", func(t *testing.T) {
		// Derive from random phrase
		result1, err := cryptoUC.MnemonicFromRandom(domain.NewDefaultParamsOfMnemonicFromRandom())
		assert.Equal(t, nil, err)

		MnemDSK := domain.NewDefaultParamsOfMnemonicDeriveSignKeys()
		MnemDSK.Phrase = result1.Phrase

		result2, err := cryptoUC.MnemonicDeriveSignKeys(MnemDSK)
		assert.Equal(t, nil, err)
		assert.IsType(t, &domain.KeyPair{}, result2)

		// Derive from provided phrase amd conver public to ton_safe
		phrase := "unit follow zone decline glare flower crisp vocal adapt magic much mesh cherry teach mechanic rain float vicious solution assume hedgehog rail sort chuckle"
		deriveParams := &domain.ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Dictionary: dictMnem["TON"], WordCount: util.IntToPointerInt(24)}
		keyPair1, err := cryptoUC.MnemonicDeriveSignKeys(deriveParams)
		assert.Equal(t, nil, err)

		convertParams := &domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair1.Public}
		publicSafe1, err := cryptoUC.ConvertPublicKeyString(convertParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuYTvCuf__YXhp-4jv3TXTHL0iK65ImwxG0RGrYc1sP3H4KS", publicSafe1.TonPublicKey)

		// Derive with path
		keyPair2, err := cryptoUC.MnemonicDeriveSignKeys(&domain.ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Path: "m", Dictionary: dictMnem["TON"], WordCount: lenMnem[24]})
		assert.Equal(t, nil, err)

		publicSafe2, err := cryptoUC.ConvertPublicKeyString(&domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair2.Public})
		assert.Equal(t, nil, err)
		assert.Equal(t, "PubDdJkMyss2qHywFuVP1vzww0TpsLxnRNnbifTCcu-XEgW0", publicSafe2.TonPublicKey)

		// Derive from 12-word phrase
		phrase = "abandon math mimic master filter design carbon crystal rookie group knife young"
		deriveParams = &domain.ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Dictionary: util.IntToPointerInt(1)}
		keypair3, err := cryptoUC.MnemonicDeriveSignKeys(deriveParams)
		assert.Equal(t, nil, err)

		convertParams = &domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keypair3.Public}
		result3, err := cryptoUC.ConvertPublicKeyString(convertParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuZhw8W5ejPJwKA68RL7sn4_RNmeH4BIU_mEK7em5d4_-cIx", result3.TonPublicKey)

		// Derive from mnemonic from entropy
		params := &domain.ParamsOfMnemonicFromEntropy{Entropy: "2199ebe996f14d9e4e2595113ad1e627", Dictionary: util.IntToPointerInt(1), WordCount: util.IntToPointerInt(12)}
		mnemonic, err := cryptoUC.MnemonicFromEntropy(params)
		assert.Equal(t, nil, err)

		deriveParams = &domain.ParamsOfMnemonicDeriveSignKeys{Phrase: mnemonic.Phrase, Dictionary: util.IntToPointerInt(1)}
		keypair4, err := cryptoUC.MnemonicDeriveSignKeys(deriveParams)
		assert.Equal(t, nil, err)

		convertParams = &domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keypair4.Public}
		result4, err := cryptoUC.ConvertPublicKeyString(convertParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuZdw_KyXIzo8IksTrERN3_WoAoYTyK7OvM-yaLk711sUIB3", result4.TonPublicKey)
	})

	t.Run("TestNaclSignKeypairFromSecretKey", func(t *testing.T) {
		params := &domain.ParamsOfNaclSignKeyPairFromSecret{Secret: "8fb4f2d256e57138fb310b0a6dac5bbc4bee09eb4821223a720e5b8e1f3dd674"}
		result, err := cryptoUC.NaclSignKeypairFromSecretKey(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "aa5533618573860a7e1bf19f34bd292871710ed5b2eafa0dcdbb33405f2231c6", result.Public)

		params.Secret = "0a"
		_, err = cryptoUC.NaclSignKeypairFromSecretKey(params)
		assert.NotNil(t, err)
	})

	t.Run("TestNaclSign", func(t *testing.T) {
		// Nacl sign data
		secret := "56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"

		params := &domain.ParamsOfNaclSign{Unsigned: unsigned, Secret: secret}
		signed, err := cryptoUC.NaclSign(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "+wz+QO6l1slgZS5s65BNqKcu4vz24FCJz4NSAxef9lu0jFfs8x3PzSZRC+pn5k8+aJi3xYMA3BQzglQmjK3hA1Rlc3QgTWVzc2FnZQ==", signed.Signed)

		// Nacl sign open
		params1 := &domain.ParamsOfNaclSignOpen{Signed: signed.Signed, Public: "1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"}
		result, err := cryptoUC.NaclSignOpen(params1)
		assert.Equal(t, nil, err)
		assert.Equal(t, unsigned, result.Unsigned)

		// Nacl sign detached
		params = &domain.ParamsOfNaclSign{Unsigned: unsigned, Secret: secret}
		result1, err := cryptoUC.NaclSignDetached(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade103", result1.Signature)

		// Nacl sign detached verify signature
		params2 := &domain.ParamsOfNaclSignDetachedVerify{Unsigned: unsigned, Signature: result1.Signature, Public: "1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"}
		result2, err := cryptoUC.NaclSignDetachedVerify(params2)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, result2.Succeeded)

		params.Secret = "0=="
		_, err = cryptoUC.NaclSign(params)
		assert.NotNil(t, err)

		params1 = &domain.ParamsOfNaclSignOpen{Signed: signed.Signed, Public: "0x00"}
		_, err = cryptoUC.NaclSignOpen(params1)
		assert.NotNil(t, err)
	})

	t.Run("TestNaclBoxKeypair", func(t *testing.T) {
		result, err := cryptoUC.NaclBoxKeypair()
		assert.Equal(t, nil, err)
		assert.Equal(t, 64, len(result.Public))
		assert.Equal(t, 64, len(result.Secret))
		assert.NotEqual(t, result.Secret, result.Public)
	})

	t.Run("TestNaclBoxKeypairFromSecretKey", func(t *testing.T) {
		params := &domain.ParamsOfNaclBoxKeyPairFromSecret{Secret: "e207b5966fb2c5be1b71ed94ea813202706ab84253bdf4dc55232f82a1caf0d4"}
		result, err := cryptoUC.NaclBoxKeypairFromSecretKey(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "a53b003d3ffc1e159355cb37332d67fc235a7feb6381e36c803274074dc3933a", result.Public)

		params.Secret = "0x00"
		_, err = cryptoUC.NaclBoxKeypairFromSecretKey(params)
		assert.NotNil(t, err)
	})

	t.Run("TestNaclBoxAndOpen", func(t *testing.T) {
		nonce := "cd7f99924bf422544046e83595dd5803f17536f5c9a11746"
		theirPublic := "c4e2d9fe6a6baf8d1812b799856ef2a306291be7a7024837ad33a8530db79c6b"
		secret := "d9b9dc5033fb416134e5d2107fdbacab5aadb297cb82dbdcd137d663bac59f7f"

		// Create nacl box
		boxParams := &domain.ParamsOfNaclBox{Decrypted: unsigned, Nonce: nonce, TheirPublic: theirPublic, Secret: secret}
		box, err := cryptoUC.NaclBox(boxParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "li4XED4kx/pjQ2qdP0eR2d/K30uN94voNADxwA==", box.Encrypted)

		// Open nacl box
		boxOpenParams := &domain.ParamsOfNaclBoxOpen{Encrypted: box.Encrypted, Nonce: nonce, TheirPublic: theirPublic, Secret: secret}
		opened, err := cryptoUC.NaclBoxOpen(boxOpenParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, unsigned, opened.Decrypted)

		boxParams.Decrypted = "0x00"
		boxParams.TheirPublic = ""
		_, err = cryptoUC.NaclBox(boxParams)
		assert.NotNil(t, err)

		boxOpenParams.Secret = ""
		_, err = cryptoUC.NaclBoxOpen(boxOpenParams)
		assert.NotNil(t, err)
	})

	t.Run("TestNaclSecretBoxAndOpen", func(t *testing.T) {
		nonce := "2a33564717595ebe53d91a785b9e068aba625c8453a76e45"
		key := "8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8"

		// Create nacl secret box
		boxParams := &domain.ParamsOfNaclSecretBox{Decrypted: unsigned, Nonce: nonce, Key: key}
		box, err := cryptoUC.NaclSecretBox(boxParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, "JL7ejKWe2KXmrsns41yfXoQF0t/C1Q8RGyzQ2A==", box.Encrypted)

		//# Open nacl secret box
		boxOpenParams := &domain.ParamsOfNaclSecretBoxOpen{Encrypted: box.Encrypted, Nonce: nonce, Key: key}
		opened, err := cryptoUC.NaclSecretBoxOpen(boxOpenParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, unsigned, opened.Decrypted)

		boxParams.Decrypted = "0x00"
		boxParams.Key = ""
		_, err = cryptoUC.NaclSecretBox(boxParams)
		assert.NotNil(t, err)

		boxOpenParams.Key = ""
		_, err = cryptoUC.NaclSecretBoxOpen(boxOpenParams)
		assert.NotNil(t, err)
	})

	t.Run("TestScrypt", func(t *testing.T) {
		params := &domain.ParamsOfScrypt{
			Password: base64.StdEncoding.EncodeToString([]byte("Test Password")),
			Salt:     base64.StdEncoding.EncodeToString([]byte("Test Salt")),
			LogN:     10,
			R:        8,
			P:        16,
			DkLen:    64}
		result, err := cryptoUC.Scrypt(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "52e7fcf91356eca55fc5d52f16f5d777e3521f54e3c570c9bbb7df58fc15add73994e5db42be368de7ebed93c9d4f21f9be7cc453358d734b04a057d0ed3626d", result.Key)

		params.DkLen = 0
		_, err = cryptoUC.Scrypt(params)
		assert.NotNil(t, err)
	})

	t.Run("TestChacha20", func(t *testing.T) {
		key := strings.Repeat("01", 32)
		nonce := strings.Repeat("ff", 12)
		data := base64.StdEncoding.EncodeToString([]byte("Message"))
		params := &domain.ParamsOfChaCha20{Data: data, Key: key, Nonce: nonce}
		encrypted, err := cryptoUC.Chacha20(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, "w5QOGsJodQ==", encrypted.Data)

		params.Data = encrypted.Data
		decrypted, err := cryptoUC.Chacha20(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, data, decrypted.Data)
	})

	t.Run("TestSigningBox", func(t *testing.T) {
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)

		// Create handle
		signingBox, err := cryptoUC.GetSigningBox(keypair)
		assert.Equal(t, nil, err)
		assert.Equal(t, domain.SigningBoxHandle(1), signingBox.Handle)

		// Get public key from box
		result, err := cryptoUC.SigningBoxGetPublicKey(signingBox)
		assert.Equal(t, nil, err)
		assert.Equal(t, keypair.Public, result.PubKey)

		// Sign with box
		message := base64.StdEncoding.EncodeToString([]byte("Sign with box"))
		params := &domain.ParamsOfSigningBoxSign{SigningBox: signingBox.Handle, Unsigned: message}
		boxResult, err := cryptoUC.SigningBoxSign(params)
		assert.Equal(t, nil, err)

		signParams := &domain.ParamsOfSign{Unsigned: message, Keys: keypair}
		signResult, err := cryptoUC.Sign(signParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, signResult.Signature, boxResult.Signature)

		// Remove signing box
		err = cryptoUC.RemoveSigningBox(signingBox)
		assert.Equal(t, nil, err)
	})

	t.Run("TestRegisterSigningBox", func(t *testing.T) {
		keys, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)

		handle, err := cryptoUC.RegisterSigningBox(&AppSigningBoxTest{Private: keys.Secret, Public: keys.Public})
		assert.Equal(t, nil, err)

		keyResult, err := cryptoUC.SigningBoxGetPublicKey(handle)
		assert.Equal(t, nil, err)
		assert.Equal(t, keyResult.PubKey, keys.Public)

		messageToSign := []byte("Test message")
		signResult, err := cryptoUC.SigningBoxSign(&domain.ParamsOfSigningBoxSign{SigningBox: handle.Handle, Unsigned: base64.StdEncoding.EncodeToString(messageToSign)})
		assert.Equal(t, nil, err)

		pubKeyBytes, err := hex.DecodeString(keyResult.PubKey)
		assert.Equal(t, nil, err)

		signatureBytes, err := hex.DecodeString(signResult.Signature)
		assert.Equal(t, nil, err)

		assert.True(t, ed25519.Verify(pubKeyBytes, messageToSign, signatureBytes))
		assert.Equal(t, nil, cryptoUC.RemoveSigningBox(&domain.RegisteredSigningBox{Handle: handle.Handle}))
	})

}
