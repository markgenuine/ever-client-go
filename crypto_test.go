package goton

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestCrypto(t *testing.T) {
	config := NewConfig(0)
	client, err := InitClient(config)
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

	idFactorize := client.Factorize(ParamsOfFactorize{Composite: "17ED48941A08F981"})
	valueFactorize, err := client.GetResp(idFactorize)
	if len(valueFactorize.(ResultOfFactorize).Factors) == 0 || err != nil {
		t.Errorf("test Failed - Error get Factorize method, err: %s", err)
	}
	if !(valueFactorize.(ResultOfFactorize).Factors[0] == "494C553B" && valueFactorize.(ResultOfFactorize).Factors[1] == "53911073") {
		t.Errorf("test Failed - error value different factorize value")
	}

	idModularPower := client.ModularPower(ParamsOfModularPower{"0123456789ABCDEF", "0123", "01234567"})
	valueModularPower, err := client.GetResp(idModularPower)
	if err != nil {
		t.Errorf("test Failed - Error get Modular Power method, err: %s", err)
	}
	if valueModularPower.(ResultOfModularPower).ModularPower != "63bfdf" {
		t.Errorf("test Failed - error get math modular power don't corrected, different '63bfdf'")
	}

	idTonCrc16 := client.TonCrc16(ParamsOfTonCrc16{Data: base64.StdEncoding.EncodeToString(fromHex("0123456789abcdef"))})
	valueTonCrc16, err := client.GetResp(idTonCrc16)
	if err != nil {
		t.Errorf("test Failed - Error get TonCRC16 method, err: %s", err)
	}
	if valueTonCrc16.(ResultOfTonCrc16).Crc != 43349 {
		t.Errorf("test Failed - error get math modular power don't corrected, different '63bfdf'")
	}

	idRandomGenerateBytes := client.GenerateRandomBytes(ParamsOfGenerateRandomBytes{Length: 32})
	valueRandomGenerateBytes, err := client.GetResp(idRandomGenerateBytes)
	if err != nil {
		t.Errorf("test Failed - Error get Random Generate Bytes method, err: %s", err)
	}
	val := valueRandomGenerateBytes.(ResultOfGenerateRandomBytes).Bytes
	if len(val) != 44 {
		t.Errorf("test Failed - Error get Random Generate Bytes method, length request != 44")
	}
	val2, _ := base64.StdEncoding.DecodeString(val)
	if len(val2) != 32 {
		t.Errorf("test Failed - Error get Random Generate Bytes method, length request != 32")
	}

	idGenerateRandomKeys := client.GenerateRandomSignKeys()
	valueRandomGenerateKeys, err := client.GetResp(idGenerateRandomKeys)
	if err != nil {
		t.Errorf("test Failed - Error get Random Generate Keys method, err: %s", err)
	}
	val3 := valueRandomGenerateKeys.(KeyPair)
	if len(val3.Public) != 64 || len(val3.Secret) != 64 || val3.Public == val3.Secret {
		t.Errorf("test Failed - Error get Random Generate Keys method, bad keys!")
	}

	idConvertPublicKeyString := client.ConvertPublicKeyString(ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: "06117f59ade83e097e0fb33e5d29e8735bda82b3bf78a015542aaa853bb69600"})
	valueConvertPublicKeyString, err := client.GetResp(idConvertPublicKeyString)
	if err != nil {
		t.Errorf("test Failed - Error get Convert Public Key String method, err: %s", err)
	}

	if valueConvertPublicKeyString.(ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey != "PuYGEX9Zreg-CX4Psz5dKehzW9qCs794oBVUKqqFO7aWAOTD" {
		t.Errorf("test Failed - Error get Convert Public Key method, bad key!")
	}

	idSha256_1 := client.Sha256(ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))})
	valueSha256_1, err := client.GetResp(idSha256_1)
	if err != nil {
		t.Errorf("test Failed - Error get Sha256 method, err: %s", err)
	}
	if valueSha256_1.(ResultOfHash).Hash != "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951" {
		t.Errorf("test Failed - Error get Sha256 method result different of key")
	}

	idSha256_2 := client.Sha256(ParamsOfHash{Data: base64.StdEncoding.EncodeToString(fromHex("544f4e206973206f757220667574757265"))})
	valueSha256_2, err := client.GetResp(idSha256_2)
	if err != nil {
		t.Errorf("test Failed - Error get Sha256 method, err: %s", err)
	}
	if valueSha256_2.(ResultOfHash).Hash != "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951" {
		t.Errorf("test Failed - Error get Sha256 method result different of key")
	}

	idSha256_3 := client.Sha256(ParamsOfHash{Data: "VE9OIGlzIG91ciBmdXR1cmU="})
	valueSha256_3, err := client.GetResp(idSha256_3)
	if err != nil {
		t.Errorf("test Failed - Error get Sha256 method, err: %s", err)
	}
	if valueSha256_3.(ResultOfHash).Hash != "1e7fd5ec201652b5375e5edf3e86d0513394d2c2004dd506415abf0578261951" {
		t.Errorf("test Failed - Error get Sha256 method result different of key")
	}

	idSha512 := client.Sha512(ParamsOfHash{Data: base64.StdEncoding.EncodeToString([]byte("TON is our future"))})
	valueSha512, err := client.GetResp(idSha512)
	if err != nil {
		t.Errorf("test Failed - Error get Sha512 method, err: %s", err)
	}
	if valueSha512.(ResultOfHash).Hash != "4c52dd4cefc68319bac5e97c1f0d18ae8194fb0dd8d9e090ba8376834a0756175a9a736d1e69cb1a58d25c3d554b02a2b8ed9c3ae5cbeeccc3277746a363a434" {
		t.Errorf("test Failed - Error get Sha512 method result different of key")
	}

	idScrypt := client.Scrypt(ParamsOfScrypt{Password: base64.StdEncoding.EncodeToString([]byte("Test Password")), Salt: base64.StdEncoding.EncodeToString([]byte("Test Salt")), LogN: 10, R: 8, P: 16, DkLen: 64})
	valueScrypt, err := client.GetResp(idScrypt)
	if err != nil {
		t.Errorf("test Failed - Error get Scrypt method, err: %s", err)
	}
	if valueScrypt.(ResultOfScrypt).Key != "52e7fcf91356eca55fc5d52f16f5d777e3521f54e3c570c9bbb7df58fc15add73994e5db42be368de7ebed93c9d4f21f9be7cc453358d734b04a057d0ed3626d" {
		t.Errorf("test Failed - Error get Scrypt method result different of key")
	}

	idNaclBoxKeypair := client.NaclBoxKeypair()
	valueNBK, err := client.GetResp(idNaclBoxKeypair)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Box Keypair method, err: %s", err)
	}
	public := valueNBK.(KeyPair).Public
	secret := valueNBK.(KeyPair).Secret
	if len(public) != 64 || len(secret) != 64 || public == secret {
		t.Errorf("test Failed - Error get Nacl Box Keypair method, bad keys!")
	}

	idNaclBoxKeypairFromSK := client.NaclSignKeypairFromSecretKey(ParamsOfNaclSignKeyPairFromSecret{Secret: "8fb4f2d256e57138fb310b0a6dac5bbc4bee09eb4821223a720e5b8e1f3dd674"})
	valueNBKF, err := client.GetResp(idNaclBoxKeypairFromSK)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Box Keypair From Secret method, err: %s", err)
	}
	if valueNBKF.(KeyPair).Public != "aa5533618573860a7e1bf19f34bd292871710ed5b2eafa0dcdbb33405f2231c6" {
		t.Errorf("test Failed - Error get Nacl Box Keypair From Secret method, bad keys!")
	}

	decrypted := base64.StdEncoding.EncodeToString([]byte("Test Message"))
	nonce := "cd7f99924bf422544046e83595dd5803f17536f5c9a11746"
	theirPublic := "c4e2d9fe6a6baf8d1812b799856ef2a306291be7a7024837ad33a8530db79c6b"
	secret = "d9b9dc5033fb416134e5d2107fdbacab5aadb297cb82dbdcd137d663bac59f7f"
	idbox := client.NaclBox(ParamsOfNaclBox{Decrypted: decrypted, Nonce: nonce, TheirPublic: theirPublic, Secret: secret})
	valBox, err := client.GetResp(idbox)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Box in Nacl Box Open, err: %s", err)
	}

	box := valBox.(ResultOfNaclBox).Encrypted
	if box != "li4XED4kx/pjQ2qdP0eR2d/K30uN94voNADxwA==" {
		t.Errorf("test Failed - Error value Nacl Box method")
	}

	idOpenBox := client.NaclBoxOpen(ParamsOfNaclBoxOpen{Encrypted: box, Nonce: nonce, TheirPublic: theirPublic, Secret: secret})
	valOB, err := client.GetResp(idOpenBox)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Box Open, err: %s", err)
	}

	if valOB.(ResultOfNaclBoxOpen).Decrypted != decrypted {
		t.Errorf("test Failed - Error get Nacl Box Open, different result value")
	}

	unsigned := base64.StdEncoding.EncodeToString([]byte("Test Message"))
	secret = "56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"
	idNaclS := client.NaclSign(ParamsOfNaclSign{Unsigned: unsigned, Secret: secret})
	valNaclS, err := client.GetResp(idNaclS)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Sign method, err: %s", err)
	}
	if valNaclS.(ResultOfNaclSign).Signed != "+wz+QO6l1slgZS5s65BNqKcu4vz24FCJz4NSAxef9lu0jFfs8x3PzSZRC+pn5k8+aJi3xYMA3BQzglQmjK3hA1Rlc3QgTWVzc2FnZQ==" {
		t.Errorf("test Failed - Error get Nacl Sign method, result different of keys")
	}

	idNaclSO := client.NaclSignOpen(ParamsOfNaclSignOpen{Signed: valNaclS.(ResultOfNaclSign).Signed, Public: "1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e"})
	valNaclSO, err := client.GetResp(idNaclSO)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Sign Open method, err: %s", err)
	}
	if unsigned != valNaclSO.(ResultOfNaclSignOpen).Unsigned {
		t.Errorf("test Failed - Error get Nacl Sign Open different key")
	}

	idNaclSignDet := client.NaclSignDetached(ParamsOfNaclSign{Unsigned: unsigned, Secret: secret})
	valNaclSDet, err := client.GetResp(idNaclSignDet)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Sign Detached method, err: %s", err)
	}
	if valNaclSDet.(ResultOfNaclSignDetached).Signature != "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade103" {
		t.Errorf("test Failed - Error get Nacl Sign Detached differnet key")
	}

	decrypted = base64.StdEncoding.EncodeToString([]byte(`Text with \' and \" and : {}`))
	nonce = "2a33564717595ebe53d91a785b9e068aba625c8453a76e45"
	key := "8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8"

	idNaclSecretBox := client.NaclSecretBox(ParamsOfNaclSecretBox{Decrypted: decrypted, Nonce: nonce, Key: key})
	valNaclSB, err := client.GetResp(idNaclSecretBox)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Secret Box method, err: %s", err)
	}
	box = valNaclSB.(ResultOfNaclBox).Encrypted
	if box != "lrWXE3+j/uXvhcgH0uvJa4QF2d/C7wMWAG3rmnk9vYiDgjL/Sf0qkXzpWK0=" {
		t.Errorf("test Failed - Error get Nacl Secret Box different of key")
	}

	idNaclSecretBoxOpen := client.NaclSecretBoxOpen(ParamsOfNaclSecretBoxOpen{Encrypted: box, Nonce: nonce, Key: key})
	valNaclSBO, err := client.GetResp(idNaclSecretBoxOpen)
	if err != nil {
		t.Errorf("test Failed - Error get Nacl Secret Box Open, err: %s", err)
	}
	if valNaclSBO.(ResultOfNaclBoxOpen).Decrypted != decrypted {
		t.Errorf("test Failed - Error get Nacl Secret Box Open different of key")
	}

	idMnemWords := client.MnemonicWords(ParamsOfMnemonicWords{})
	valMnemWords, err := client.GetResp(idMnemWords)
	if err != nil {
		t.Errorf("test Failed - Error get Mnemonic Words method, err: %s", err)
	}
	if len(strings.Fields(valMnemWords.(ResultOfMnemonicWords).Words)) != 2048 {
		t.Errorf("test Failed - error different count words of 2048")
	}

	lenMnem := getlengthCounWordsInMnemonic()
	dictMnem := getTONMnemonicDictionary()
	for vall := range lenMnem {
		for _, vald := range dictMnem {
			idMnemRandom := client.MnemonicFromRandom(ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: vall})
			valMnemRand, err := client.GetResp(idMnemRandom)
			if err != nil {
				t.Errorf("test Failed - Error get Mnemonic Random method, err: %s", err)
			}
			if vall != len(strings.Fields(valMnemRand.(ResultOfMnemonicFromRandom).Phrase)) {
				t.Errorf("test Failed - Error get Mnemonic Random length must be %d words", vall)
			}
		}
	}

	idMnemFromEntropy := client.MnemonicFromEntropy(ParamsOfMnemonicFromEntropy{Entropy: "00112233445566778899AABBCCDDEEFF", Dictionary: dictMnem["ENGLISH"], WordCount: 12})
	valMnemEntr, err := client.GetResp(idMnemFromEntropy)
	if err != nil {
		t.Errorf("test Failed - Error get Mnemonic From Entropy method, err: %s", err)
	}
	if valMnemEntr.(ResultOfMnemonicFromEntropy).Phrase != "abandon math mimic master filter design carbon crystal rookie group knife young" {
		t.Errorf("test Failed - Error get Mnemonic From Entropy different of key")
	}

	for vall := range lenMnem {
		for _, vald := range dictMnem {
			idMnemRandom2 := client.MnemonicFromRandom(ParamsOfMnemonicFromRandom{Dictionary: vald, WordCount: vall})
			valMnemRand2, err := client.GetResp(idMnemRandom2)
			if err != nil {
				t.Errorf("test Failed - Error get Mnemonic Random method in Mnemonic Verify, err: %s", err)
			}
			idMnemVerify := client.MnemonicVerify(ParamsOfMnemonicVerify{Phrase: valMnemRand2.(ResultOfMnemonicFromRandom).Phrase, Dictionary: vald, WordCount: vall})
			valMnemVerify, err := client.GetResp(idMnemVerify)
			if err != nil {
				t.Errorf("test Failed - Error get Mnemonic Verify method, err: %s", err)
			}
			if !valMnemVerify.(ResultOfMnemonicVerify).Valid {
				t.Errorf("test Failed - Error get Mnemonic Verify, mnemonic don't valid")
			}
		}
	}

	phrase := "unit follow zone decline glare flower crisp vocal adapt magic much mesh cherry teach mechanic rain float vicious solution assume hedgehog rail sort chuckle"
	idMnemDerSK1 := client.MnemonicDeriveSignKeys(ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Dictionary: dictMnem["TON"], WordCount: lenMnem[24]})
	keyPair1, err := client.GetResp(idMnemDerSK1)
	if err != nil {
		t.Errorf("test Failed - Error get Mnemonic Derive Sign Keys method, err: %s", err)
	}

	idPublicSafe1 := client.ConvertPublicKeyString(ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair1.(KeyPair).Public})
	publicSafe1, err := client.GetResp(idPublicSafe1)
	if err != nil {
		t.Errorf("test Failed - Error get Convert Public Key String in Mnemonic Derive Sign Keys method, err: %s", err)
	}
	if publicSafe1.(ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey != "PuYTvCuf__YXhp-4jv3TXTHL0iK65ImwxG0RGrYc1sP3H4KS" {
		t.Errorf("test Failed - Error get Mnemonic Derive Sign Keys different key!")
	}

	idMnemDerSK2 := client.MnemonicDeriveSignKeys(ParamsOfMnemonicDeriveSignKeys{Phrase: phrase, Path: "m", Dictionary: dictMnem["TON"], WordCount: lenMnem[24]})
	keyPair2, err := client.GetResp(idMnemDerSK2)
	if err != nil {
		t.Errorf("test Failed - Error get Mnemonic Derive Sign Keys method, err: %s", err)
	}

	idPublicSafe2 := client.ConvertPublicKeyString(ParamsOfConvertPublicKeyToTonSafeFormat{PublicKey: keyPair2.(KeyPair).Public})
	publicSafe2, err := client.GetResp(idPublicSafe2)
	if err != nil {
		t.Errorf("test Failed - Error get Convert Public Key String in Mnemonic Derive Sign Keys method, err: %s", err)
	}
	if publicSafe2.(ResultOfConvertPublicKeyToTonSafeFormat).TonPublicKey != "PubDdJkMyss2qHywFuVP1vzww0TpsLxnRNnbifTCcu-XEgW0" {
		t.Errorf("test Failed - Error get Mnemonic Derive Sign Keys different key!")
	}

	// HDKeys
	masterXPrv := "xprv9s21ZrQH143K25JhKqEwvJW7QAiVvkmi4WRenBZanA6kxHKtKAQQKwZG65kCyW5jWJ8NY9e3GkRoistUjjcpHNsGBUv94istDPXvqGNuWpC"
	idXprv := client.HdkeyXprvFromMnemonic(ParamsOfHDKeyXPrvFromMnemonic{Phrase: "abuse boss fly battle rubber wasp afraid hamster guide essence vibrant tattoo", Dictionary: dictMnem["ENGLISH"], WordCount: 12})
	xprv, err := client.GetResp(idXprv)
	if err != nil {
		t.Errorf("test Failed - Error get HDKey XPrv From Mnemonic method, err: %s", err)
	}
	if xprv.(ResultOfHDKeyXPrvFromMnemonic).Xprv != masterXPrv {
		t.Errorf("test Failed - Error get HDKey XPrv From Mnemonic different of key")
	}

	idDeriveXPrv := client.HdkeyXprvDerive(ParamsOfHDKeyDeriveFromXPrv{Xprv: masterXPrv, ChildIndex: 0, Hardened: false})
	deriveXPrv, err := client.GetResp(idDeriveXPrv)
	if err != nil {
		t.Errorf("test Failed - Error get HDKey Derive XPrv method, err: %s", err)
	}
	if deriveXPrv.(ResultOfHDKeyDeriveFromXPrv).Xprv != "xprv9uZwtSeoKf1swgAkVVCEUmC2at6t7MCJoHnBbn1MWJZyxQ4cySkVXPyNh7zjf9VjsP4vEHDDD2a6R35cHubg4WpzXRzniYiy8aJh1gNnBKv" {
		t.Errorf("test Failed - Error get HDKey Derive From XPrv different of key")
	}

	idDeriveXPrvPath := client.HdkeyXprvDerivePath(ParamsOfHDKeyDeriveFromXPrvPath{Xprv: masterXPrv, Path: "m/44'/60'/0'/0'"})
	deriveXPrvP, err := client.GetResp(idDeriveXPrvPath)
	if err != nil {
		t.Errorf("test Failed - Error get HDKey Derive XPrv Path method, err: %s", err)
	}
	if deriveXPrvP.(ResultOfHDKeyDeriveFromXPrvPath).Xprv != "xprvA1KNMo63UcGjmDF1bX39Cw2BXGUwrwMjeD5qvQ3tA3qS3mZQkGtpf4DHq8FDLKAvAjXsYGLHDP2dVzLu9ycta8PXLuSYib2T3vzLf3brVgZ" {
		t.Errorf("test Failed - Error get HDKey Derive From XPrv Path different of key")
	}

	idSecretXPrv := client.HdkeyXprvSecret(ParamsOfHDKeySecretFromXPrv{Xprv: masterXPrv})
	secretXPrv, err := client.GetResp(idSecretXPrv)
	if err != nil {
		t.Errorf("test Failed - Error get HDKey XPrv Secret method, err: %s", err)
	}
	if secretXPrv.(ResultOfHDKeySecretFromXPrv).Secret != "0c91e53128fa4d67589d63a6c44049c1068ec28a63069a55ca3de30c57f8b365" {
		t.Errorf("test Failed - Error get HDKey Secret different of key")
	}

	idPublicXPrv := client.HdkeyXprvPublic(ParamsOfHDKeyPublicFromXPrv{Xprv: masterXPrv})
	publicXPrv, err := client.GetResp(idPublicXPrv)
	if err != nil {
		t.Errorf("test Failed -Error get HDKey XPrv Public method, err: %s", err)
	}
	if publicXPrv.(ResultOfHDKeyPublicFromXPrv).Public != "02a8eb63085f73c33fa31b4d1134259406347284f8dab6fc68f4bf8c96f6c39b75" {
		t.Errorf("test Failed - Error get HDKey Public different of key")
	}
}
