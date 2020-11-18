package crypto

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/util"
	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {

	client, err := client.NewClientGateway(domain.NewDefaultConfig(2))
	assert.Equal(t, nil, err)
	defer client.Destroy()

	cryptoUC := crypto{
		config: domain.NewDefaultConfig(2),
		client: client,
	}
	defer cryptoUC.client.Destroy()

	t.Run("TestFactorize", func(t *testing.T) {
		idFactorize, err := cryptoUC.Factorize(domain.ParamsOfFactorize{Composite: "17ED48941A08F981"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idFactorize)
		valueFactorize, err := cryptoUC.client.GetResp(idFactorize)
		assert.Equal(t, nil, err)
		resultFactorize := valueFactorize.(domain.ResultOfFactorize)
		assert.Equal(t, "494C553B", resultFactorize.Factors[0])
		assert.Equal(t, "53911073", resultFactorize.Factors[1])
	})

	t.Run("TestModularPower", func(t *testing.T) {
		idModularPower, err := cryptoUC.ModularPower(domain.ParamsOfModularPower{Base: "0123456789ABCDEF", Exponent: "0123", Modulus: "01234567"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idModularPower)
		valueMP, err := cryptoUC.client.GetResp(idModularPower)
		assert.Equal(t, nil, err)
		assert.Equal(t, "63bfdf", valueMP.(domain.ResultOfModularPower).ModularPower)
	})

	t.Run("TestTonCrc16", func(t *testing.T) {
		idTonCrc16, err := cryptoUC.TonCrc16(domain.ParamsOfTonCrc16{Data: base64.StdEncoding.EncodeToString(util.FromHex("0123456789abcdef"))})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idTonCrc16)
		valueTonCRC16, err := cryptoUC.client.GetResp(idTonCrc16)
		assert.Equal(t, nil, err)
		assert.Equal(t, 43349, valueTonCRC16.(domain.ResultOfTonCrc16).Crc)
	})

	t.Run("TestGenerateRandomBytes", func(t *testing.T) {
		idRandomGB, err := cryptoUC.GenerateRandomBytes(domain.ParamsOfGenerateRandomBytes{Length: 32})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idRandomGB)
		valueGenRB, err := cryptoUC.client.GetResp(idRandomGB)
		assert.Equal(t, nil, err)
		resultGenRB := valueGenRB.(domain.ResultOfGenerateRandomBytes)
		assert.Equal(t, 44, len(resultGenRB.Bytes))
		resultDecode, err := base64.StdEncoding.DecodeString(resultGenRB.Bytes)
		assert.Equal(t, nil, err)
		assert.Equal(t, 32, len(resultDecode))
	})

	t.Run("TestConvertPublicKeyString", func(t *testing.T) {
		idConvertPubKey, err := cryptoUC.ConvertPublicKeyString(domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: "06117f59ade83e097e0fb33e5d29e8735bda82b3bf78a015542aaa853bb69600"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConvertPubKey)
		valueConvPubKey, err := cryptoUC.client.GetResp(idConvertPubKey)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuYGEX9Zreg-CX4Psz5dKehzW9qCs794oBVUKqqFO7aWAOTD", valueConvPubKey.(domain.ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey)

	})

	t.Run("TestGenerateRandomSignKeys", func(t *testing.T) {
		idGenerateRandomKeys, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idGenerateRandomKeys)
		valueRandomGenerateKeys, err := cryptoUC.client.GetResp(idGenerateRandomKeys)
		assert.Equal(t, nil, err)
		resultRandomGK := valueRandomGenerateKeys.(domain.KeyPair)
		assert.Equal(t, 64, len(resultRandomGK.Public))
		assert.Equal(t, 64, len(resultRandomGK.Secret))
		assert.NotEqual(t, resultRandomGK.Secret, resultRandomGK.Public)
	})

	t.Run("TestSha256", func(t *testing.T) {
		idSha256_1, err := cryptoUC.Sha256(domain.ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idSha256_1)

		idSha256_2, err := cryptoUC.Sha256(domain.ParamsOfHash{Data: base64.StdEncoding.EncodeToString(util.FromHex("544f4e206973206f757220667574757265"))})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idSha256_2)

		idSha256_3, err := cryptoUC.Sha256(domain.ParamsOfHash{Data: "VE9OIGlzIG91ciBmdXR1cmU="})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idSha256_3)

		valueSha256_1, err := cryptoUC.client.GetResp(idSha256_1)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", valueSha256_1.(domain.ResultOfHash).Hash)

		valueSha256_2, err := cryptoUC.client.GetResp(idSha256_2)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", valueSha256_2.(domain.ResultOfHash).Hash)

		valueSha256_3, err := cryptoUC.client.GetResp(idSha256_3)
		assert.Equal(t, nil, err)
		assert.Equal(t, "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951", valueSha256_3.(domain.ResultOfHash).Hash)
	})

	t.Run("TestSha512", func(t *testing.T) {
		idSha512, err := cryptoUC.Sha512(domain.ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idSha512)
		valueSha512, err := cryptoUC.client.GetResp(idSha512)
		assert.Equal(t, nil, err)
		assert.Equal(t, "4c52dd4cefc68319bac5e97c1f0d18ae8194fb0dd8d9e090ba8376834a0756175a9a736d1e69cb1a58d25c3d554b02a2b8ed9c3ae5cbeeccc3277746a363a434", valueSha512.(domain.ResultOfHash).Hash)
	})

	t.Run("TestScrypt", func(t *testing.T) {
		idScrypt, err := cryptoUC.Scrypt(domain.ParamsOfScrypt{Password: base64.StdEncoding.EncodeToString([]byte("Test Password")), Salt: base64.StdEncoding.EncodeToString([]byte("Test Salt")), LogN: 10, R: 8, P: 16, DkLen: 64})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idScrypt)
		valueScrypt, err := cryptoUC.client.GetResp(idScrypt)
		assert.Equal(t, nil, err)
		assert.Equal(t, "52e7fcf91356eca55fc5d52f16f5d777e3521f54e3c570c9bbb7df58fc15add73994e5db42be368de7ebed93c9d4f21f9be7cc453358d734b04a057d0ed3626d", valueScrypt.(domain.ResultOfScrypt).Key)
	})

	t.Run("TestNaclBoxKeypair", func(t *testing.T) {
		idNaclBoxKeypair, err := cryptoUC.NaclBoxKeypair()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclBoxKeypair)
		valueNBK, err := cryptoUC.client.GetResp(idNaclBoxKeypair)
		assert.Equal(t, nil, err)
		resultNBK := valueNBK.(domain.KeyPair)
		assert.Equal(t, 64, len(resultNBK.Public))
		assert.Equal(t, 64, len(resultNBK.Secret))
		assert.NotEqual(t, resultNBK.Secret, resultNBK.Public)
	})

	t.Run("TestNaclSignKeypairFromSecretKey", func(t *testing.T) {
		idNaclBoxKeypairFromSK, err := cryptoUC.NaclSignKeypairFromSecretKey(domain.ParamsOfNaclSignKeyPairFromSecret{Secret: "8fb4f2d256e57138fb310b0a6dac5bbc4bee09eb4821223a720e5b8e1f3dd674"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclBoxKeypairFromSK)
		valueNBKF, err := cryptoUC.client.GetResp(idNaclBoxKeypairFromSK)
		assert.Equal(t, nil, err)
		assert.Equal(t, "aa5533618573860a7e1bf19f34bd292871710ed5b2eafa0dcdbb33405f2231c6", valueNBKF.(domain.KeyPair).Public)
	})

	decrypted := base64.StdEncoding.EncodeToString([]byte("Test Message"))
	nonce := "cd7f99924bf422544046e83595dd5803f17536f5c9a11746"
	theirPublic := "c4e2d9fe6a6baf8d1812b799856ef2a306291be7a7024837ad33a8530db79c6b"
	secret := "d9b9dc5033fb416134e5d2107fdbacab5aadb297cb82dbdcd137d663bac59f7f"

	t.Run("TestNaclBoxAndNaclBoxOpen", func(t *testing.T) {
		idBox, err := cryptoUC.NaclBox(domain.ParamsOfNaclBox{Decrypted: decrypted, Nonce: nonce, TheirPublic: theirPublic, Secret: secret})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idBox)
		valBox, err := cryptoUC.client.GetResp(idBox)
		assert.Equal(t, nil, err)
		box := valBox.(domain.ResultOfNaclBox).Encrypted
		assert.Equal(t, "li4XED4kx/pjQ2qdP0eR2d/K30uN94voNADxwA==", box)

		idOpenBox, err := cryptoUC.NaclBoxOpen(domain.ParamsOfNaclBoxOpen{Encrypted: box, Nonce: nonce, TheirPublic: theirPublic, Secret: secret})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idOpenBox)
		valOB, err := cryptoUC.client.GetResp(idOpenBox)
		assert.Equal(t, nil, err)
		assert.Equal(t, decrypted, valOB.(domain.ResultOfNaclBoxOpen).Decrypted)
	})

	unsigned := base64.StdEncoding.EncodeToString([]byte("Test Message"))
	secret = "56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"
	t.Run("TestNaclSignAndNaclSignDetached", func(t *testing.T) {
		idNaclS, err := cryptoUC.NaclSign(domain.ParamsOfNaclSign{Unsigned: unsigned, Secret: secret})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclS)
		valNaclS, err := cryptoUC.client.GetResp(idNaclS)
		assert.Equal(t, nil, err)
		assert.Equal(t, "+wz+QO6l1slgZS5s65BNqKcu4vz24FCJz4NSAxef9lu0jFfs8x3PzSZRC+pn5k8+aJi3xYMA3BQzglQmjK3hA1Rlc3QgTWVzc2FnZQ==", valNaclS.(domain.ResultOfNaclSign).Signed)

		idNaclSO, err := cryptoUC.NaclSignOpen(domain.ParamsOfNaclSignOpen{Signed: valNaclS.(domain.ResultOfNaclSign).Signed, Public: "1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclSO)
		valNaclSO, err := cryptoUC.client.GetResp(idNaclSO)
		assert.Equal(t, nil, err)
		assert.Equal(t, unsigned, valNaclSO.(domain.ResultOfNaclSignOpen).Unsigned)

		idNaclSignDet, err := cryptoUC.NaclSignDetached(domain.ParamsOfNaclSign{Unsigned: unsigned, Secret: secret})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclSignDet)
		valNaclSDet, err := cryptoUC.client.GetResp(idNaclSignDet)
		assert.Equal(t, nil, err)
		assert.Equal(t, "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade103", valNaclSDet.(domain.ResultOfNaclSignDetached).Signature)
	})

	decrypted = base64.StdEncoding.EncodeToString([]byte(`Text with \' and \" and : {}`))
	nonce = "2a33564717595ebe53d91a785b9e068aba625c8453a76e45"
	key := "8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8"
	t.Run("TestNaclSecretBoxAndNaclSecretBoxOpen", func(t *testing.T) {
		idNaclSecretBox, err := cryptoUC.NaclSecretBox(domain.ParamsOfNaclSecretBox{Decrypted: decrypted, Nonce: nonce, Key: key})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclSecretBox)
		valNaclSB, err := client.GetResp(idNaclSecretBox)
		assert.Equal(t, nil, err)
		box := valNaclSB.(domain.ResultOfNaclBox).Encrypted
		assert.Equal(t, "lrWXE3+j/uXvhcgH0uvJa4QF2d/C7wMWAG3rmnk9vYiDgjL/Sf0qkXzpWK0=", box)

		idNaclSecretBoxOpen, err := cryptoUC.NaclSecretBoxOpen(domain.ParamsOfNaclSecretBoxOpen{Encrypted: box, Nonce: nonce, Key: key})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idNaclSecretBoxOpen)
		valNaclSBO, err := cryptoUC.client.GetResp(idNaclSecretBoxOpen)
		assert.Equal(t, nil, err)
		assert.Equal(t, decrypted, valNaclSBO.(domain.ResultOfNaclBoxOpen).Decrypted)
	})

	t.Run("TestMnemonicWords", func(t *testing.T) {
		idMnemWords, err := cryptoUC.MnemonicWords(domain.NewDefaultParamsOfMnemonicWords())
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idMnemWords)
		valMnemWords, err := cryptoUC.client.GetResp(idMnemWords)
		assert.Equal(t, nil, err)
		assert.Equal(t, 2048, len(strings.Fields(valMnemWords.(domain.ResultOfMnemonicWords).Words)))
	})

	lenMnem := domain.WordCounList()
	dictMnem := domain.DictionaryList()
	t.Run("TestMnemonicFromRandom", func(t *testing.T) {
		for vall := range lenMnem {
			for _, vald := range dictMnem {
				idMnemRandom, err := cryptoUC.MnemonicFromRandom(domain.ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: vall})
				assert.Equal(t, nil, err)
				assert.NotEqual(t, 0, idMnemRandom)
				valMnemRand, err := cryptoUC.client.GetResp(idMnemRandom)
				assert.Equal(t, nil, err)
				assert.Equal(t, vall, len(strings.Fields(valMnemRand.(domain.ResultOfMnemonicFromRandom).Phrase)))
			}
		}
	})

	t.Run("TestMnemonicFromEntropy", func(t *testing.T) {
		defaultEntropy := domain.NewDefaultParamsOfMnemonicFromEntropy()
		defaultEntropy.Entropy = "00112233445566778899AABBCCDDEEFF"
		idMnemFromEntropy, err := cryptoUC.MnemonicFromEntropy(defaultEntropy)
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idMnemFromEntropy)
		valMnemEntr, err := cryptoUC.client.GetResp(idMnemFromEntropy)
		assert.Equal(t, nil, err)
		assert.Equal(t, "abandon math mimic master filter design carbon crystal rookie group knife young", valMnemEntr.(domain.ResultOfMnemonicFromEntropy).Phrase)
	})

	t.Run("TestMnemonicVerify", func(t *testing.T) {
		for vall := range lenMnem {
			for _, vald := range dictMnem {
				idMnemRandom2, err := cryptoUC.MnemonicFromRandom(domain.ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: vall})
				assert.Equal(t, nil, err)
				assert.NotEqual(t, 0, idMnemRandom2)
				valMnemRand2, err := cryptoUC.client.GetResp(idMnemRandom2)
				assert.Equal(t, nil, err)

				idMnemVerify, err := cryptoUC.MnemonicVerify(domain.ParamsOfMnemonicVerify{Phrase: valMnemRand2.(domain.ResultOfMnemonicFromRandom).Phrase, Dictionary: vald, WordCount: vall})
				assert.Equal(t, nil, err)
				assert.NotEqual(t, 0, idMnemVerify)
				valMnemVerify, err := cryptoUC.client.GetResp(idMnemVerify)
				assert.Equal(t, nil, err)
				assert.True(t, valMnemVerify.(domain.ResultOfMnemonicVerify).Valid)
			}
		}
	})

	phrase := "unit follow zone decline glare flower crisp vocal adapt magic much mesh cherry teach mechanic rain float vicious solution assume hedgehog rail sort chuckle"
	t.Run("TestMnemonicDeriveSignKeys", func(t *testing.T) {
		idMnemRandom, err := cryptoUC.MnemonicFromRandom(domain.NewDefaultParamsOfMnemonicFromRandom())
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idMnemRandom)
		valueMnemR, err := cryptoUC.client.GetResp(idMnemRandom)
		assert.Equal(t, nil, err)

		MnemDSK := domain.NewDefaultParamsOfMnemonicDeriveSignKeys()
		MnemDSK.Phrase = valueMnemR.(domain.ResultOfMnemonicFromRandom).Phrase
		idMnemDeriveSK, err := cryptoUC.MnemonicDeriveSignKeys(MnemDSK)
		assert.NotEqual(t, 0, idMnemDeriveSK)
		valueDerivSK, err := cryptoUC.client.GetResp(idMnemDeriveSK)
		assert.Equal(t, nil, err)
		assert.IsType(t, domain.KeyPair{}, valueDerivSK.(domain.KeyPair))

		idMnemDeriveSK1, err := cryptoUC.MnemonicDeriveSignKeys(domain.ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Dictionary: domain.DictionaryList()["TON"], WordCount: domain.WordCounList()[24]})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idMnemDeriveSK1)
		keyPair1, err := cryptoUC.client.GetResp(idMnemDeriveSK1)
		assert.Equal(t, nil, err)

		idPublicSafe1, err := cryptoUC.ConvertPublicKeyString(domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair1.(domain.KeyPair).Public})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idPublicSafe1)
		publicSafe1, err := cryptoUC.client.GetResp(idPublicSafe1)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PuYTvCuf__YXhp-4jv3TXTHL0iK65ImwxG0RGrYc1sP3H4KS", publicSafe1.(domain.ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey)

		idMnemDerSK2, err := cryptoUC.MnemonicDeriveSignKeys(domain.ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Path: "m", Dictionary: dictMnem["TON"], WordCount: lenMnem[24]})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idMnemDerSK2)
		keyPair2, err := cryptoUC.client.GetResp(idMnemDerSK2)
		assert.Equal(t, nil, err)

		idPublicSafe2, err := cryptoUC.ConvertPublicKeyString(domain.ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair2.(domain.KeyPair).Public})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idPublicSafe2)
		publicSafe2, err := cryptoUC.client.GetResp(idPublicSafe2)
		assert.Equal(t, nil, err)
		assert.Equal(t, "PubDdJkMyss2qHywFuVP1vzww0TpsLxnRNnbifTCcu-XEgW0", publicSafe2.(domain.ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey)
	})

	masterXPrv := "xprv9s21ZrQH143K25JhKqEwvJW7QAiVvkmi4WRenBZanA6kxHKtKAQQKwZG65kCyW5jWJ8NY9e3GkRoistUjjcpHNsGBUv94istDPXvqGNuWpC"
	t.Run("TestHdkeyXprvFromMnemonic", func(t *testing.T) {
		idXprv, err := cryptoUC.HdkeyXprvFromMnemonic(domain.ParamsOfHDKeyXPrvFromMnemonic{Phrase: "abuse boss fly battle rubber wasp afraid hamster guide essence vibrant tattoo", Dictionary: dictMnem["ENGLISH"], WordCount: 12})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idXprv)
		xprv, err := cryptoUC.client.GetResp(idXprv)
		assert.Equal(t, nil, err)
		assert.Equal(t, masterXPrv, xprv.(domain.ResultOfHDKeyXPrvFromMnemonic).Xprv)
	})

	t.Run("TestHdkeyXprvDerive", func(t *testing.T) {
		idDeriveXPrv, err := cryptoUC.HdkeyXprvDerive(domain.ParamsOfHDKeyDeriveFromXPrv{Xprv: masterXPrv, ChildIndex: 0, Hardened: false})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idDeriveXPrv)
		deriveXPrv, err := cryptoUC.client.GetResp(idDeriveXPrv)
		assert.Equal(t, nil, err)
		assert.Equal(t, "xprv9uZwtSeoKf1swgAkVVCEUmC2at6t7MCJoHnBbn1MWJZyxQ4cySkVXPyNh7zjf9VjsP4vEHDDD2a6R35cHubg4WpzXRzniYiy8aJh1gNnBKv", deriveXPrv.(domain.ResultOfHDKeyDeriveFromXPrv).Xprv)
	})

	t.Run("TestHdkeyXprvDerivePath", func(t *testing.T) {
		idDeriveXPrvPath, err := cryptoUC.HdkeyXprvDerivePath(domain.ParamsOfHDKeyDeriveFromXPrvPath{Xprv: masterXPrv, Path: "m/44'/60'/0'/0'"})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idDeriveXPrvPath)
		deriveXPrvP, err := cryptoUC.client.GetResp(idDeriveXPrvPath)
		assert.Equal(t, nil, err)
		assert.Equal(t, "xprvA1KNMo63UcGjmDF1bX39Cw2BXGUwrwMjeD5qvQ3tA3qS3mZQkGtpf4DHq8FDLKAvAjXsYGLHDP2dVzLu9ycta8PXLuSYib2T3vzLf3brVgZ", deriveXPrvP.(domain.ResultOfHDKeyDeriveFromXPrvPath).Xprv)
	})

	t.Run("TestHdkeyXprvSecrets", func(t *testing.T) {
		idSecretXPrv, err := cryptoUC.HdkeyXprvSecret(domain.ParamsOfHDKeySecretFromXPrv{Xprv: masterXPrv})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idSecretXPrv)
		secretXPrv, err := cryptoUC.client.GetResp(idSecretXPrv)
		assert.Equal(t, "0c91e53128fa4d67589d63a6c44049c1068ec28a63069a55ca3de30c57f8b365", secretXPrv.(domain.ResultOfHDKeySecretFromXPrv).Secret)
	})

	t.Run("TestHdkeyXprvPublic", func(t *testing.T) {
		idPublicXPrv, err := cryptoUC.HdkeyXprvPublic(domain.ParamsOfHDKeyPublicFromXPrv{Xprv: masterXPrv})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idPublicXPrv)
		publicXPrv, err := cryptoUC.client.GetResp(idPublicXPrv)
		assert.Equal(t, nil, err)
		assert.Equal(t, "02a8eb63085f73c33fa31b4d1134259406347284f8dab6fc68f4bf8c96f6c39b75", publicXPrv.(domain.ResultOfHDKeyPublicFromXPrv).Public)
	})
}
