package crypto

import (
	"strings"
	"testing"

	goton "github.com/move-ton/ton-client-go"
)

func TestCrypto(t *testing.T) {

	client, err := goton.InitClient(goton.NewConfig(0))
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

	// Math

	t.Run("TestMathFactorize", func(t *testing.T) {
		value, _ := MathFactorizeResp(client.Request(MathFactorize("17ED48941A08F981")))
		if value.A != "494C553B" || value.B != "53911073" {
			t.Errorf("test Failed - error value different factorize value")
		}
	})

	t.Run("TestMathModularPower", func(t *testing.T) {
		mP := &goton.ModularPowerRequest{"0123456789ABCDEF", "0123", "01234567"}
		value, _ := client.Request(MathModularPower(mP))
		if value != "63bfdf" {
			t.Errorf("test Failed - error get math modular power don't corrected, different '63bfdf'")
		}
	})

	t.Run("TestTonCrc16", func(t *testing.T) {
		value, _ := client.Request(TonCrc16(goton.FixInputMessage("move-ton", "text")))
		if value != "36005" {
			t.Errorf("test Failed - error value different Ton public key")
		}
	})

	// Random

	t.Run("TestRandomGenerateBytesHex", func(t *testing.T) {
		value, _ := client.Request(RandomGenerateBytes(&goton.RandomGenerateBytesRequest{32, goton.TONOutputEncodingHex}))

		if len(value) != (64) {
			t.Errorf("test Failed - value don't correct")
		}
	})

	t.Run("TestRandomGenerateBytesHexUppercase", func(t *testing.T) {
		value, _ := client.Request(RandomGenerateBytes(&goton.RandomGenerateBytesRequest{32, goton.TONOutputEncodingHexUppercase}))

		if len(value) != (64) {
			t.Errorf("test Failed - value don't correct")
		}
	})

	t.Run("TestRandomGenerateBytesBase64", func(t *testing.T) {
		value, _ := client.Request(RandomGenerateBytes(&goton.RandomGenerateBytesRequest{32, goton.TONOutputEncodingBase64}))

		if len(value) != (44) {
			t.Errorf("test Failed - value don't correct")
		}
	})

	// ed25519

	t.Run("TestEd25519Keypair", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(Ed25519Keypair()))
		if len(value.Public) != 64 {
			t.Errorf("test Failed - value public don't correct")
		}
		if len(value.Secret) != 64 {
			t.Errorf("test Failed - value secret don't correct")
		}
		if value.Public == value.Secret {
			t.Errorf("test Failed - value public or secret don't correct")
		}
	})

	// TON
	t.Run("TestTonPublicKeyString", func(t *testing.T) {
		value, _ := client.Request(TonPublicKeyString("1123fd8e34a7ec75c15121a3ca455e0b3788f28847f1e69eefe1c2c0aa08adbe"))
		if value != "PuYRI_2ONKfsdcFRIaPKRV4LN4jyiEfx5p7v4cLAqgitvmpx" {
			t.Errorf("test Failed - error value different Ton public key")
		}
	})

	// sha

	t.Run("TestSha512", func(t *testing.T) {
		value, _ := client.Request(Sha512(&goton.MessageInputMessage{goton.FixInputMessage("Message to hash with sha 512", goton.TONInputEncodingText)}))
		if value != "2616a44e0da827f0244e93c2b0b914223737a6129bc938b8edf2780ac9482960baa9b7c7cdb11457c1cebd5ae77e295ed94577f32d4c963dc35482991442daa5" {
			t.Errorf("test Failed - value difference to sha512 value")
		}
	})

	t.Run("TestSha256_1", func(t *testing.T) {
		value, _ := client.Request(Sha256(&goton.MessageInputMessage{goton.FixInputMessage("4d65737361676520746f206861736820776974682073686120323536", goton.TONInputEncodingHex)}))
		if value != "16fd057308dd358d5a9b3ba2de766b2dfd5e308478fc1f7ba5988db2493852f5" {
			t.Errorf("test Failed - value difference to sha256 value")
		}
	})

	t.Run("TestSha256_2", func(t *testing.T) {
		value, _ := client.Request(Sha256(&goton.MessageInputMessage{goton.FixInputMessage("TWVzc2FnZSB0byBoYXNoIHdpdGggc2hhIDI1Ng==", goton.TONInputEncodingBase64)}))
		if value != "16fd057308dd358d5a9b3ba2de766b2dfd5e308478fc1f7ba5988db2493852f5" {
			t.Errorf("test Failed - value difference to sha256 value")
		}
	})

	// scrypt

	t.Run("TestScrypt", func(t *testing.T) {
		value, _ := client.Request(Scrypt(&goton.ScryptDate{goton.FixInputMessage("password", goton.TONInputEncodingText), goton.FixInputMessage("salt", goton.TONInputEncodingText), 10, 8, 16, 64, ""}))
		if value != "1effd93afcf2b28964026631bf4362b0e5ed83cbd5f326b72eb687bfbc7ac56756f8d92337963b22c53ecab5e8de24f3b24053bfb5341c28f162aca6b0898a6e" {
			t.Errorf("test Failed - error exc scrypt value different example")
		}
	})

	// nacl keys

	t.Run("TestNaclBoxKeypair", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(NaclBoxKeypair()))
		if len(value.Public) != 64 {
			t.Errorf("test Failed - value public don't correct")
		}
		if len(value.Secret) != 64 {
			t.Errorf("test Failed - value secret don't correct")
		}
		if value.Public == value.Secret {
			t.Errorf("test Failed - value public or secret don't correct")
		}
	})

	t.Run("TestNaclBoxKeypairFromSecretKey", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(NaclBoxKeypairFromSecretKey("e207b5966fb2c5be1b71ed94ea813202706ab84253bdf4dc55232f82a1caf0d4")))
		if value.Public != "a53b003d3ffc1e159355cb37332d67fc235a7feb6381e36c803274074dc3933a" {
			t.Errorf("test Failed - error Nacl Box public key don't correct")
		}
	})

	t.Run("TestNaclSignKeypair", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(NaclSignKeypair()))
		if len(value.Public) != 64 {
			t.Errorf("test Failed - value public don't correct")
		}
		if len(value.Secret) != 128 {
			t.Errorf("test Failed - value secret don't correct")
		}
	})

	t.Run("TestNaclSignKeypairFromSecretKey", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(NaclSignKeypairFromSecretKey("8fb4f2d256e57138fb310b0a6dac5bbc4bee09eb4821223a720e5b8e1f3dd674")))
		if value.Public != "aa5533618573860a7e1bf19f34bd292871710ed5b2eafa0dcdbb33405f2231c6" {
			t.Errorf("test Failed - error Nacl public key don't correct")
		}
	})

	// nacl box

	t.Run("TestNaclBox", func(t *testing.T) {
		value, _ := client.Request(NaclBox(&goton.NaclBoxIn{goton.FixInputMessage("Test Message", goton.TONInputEncodingText),
			"cd7f99924bf422544046e83595dd5803f17536f5c9a11746",
			"c4e2d9fe6a6baf8d1812b799856ef2a306291be7a7024837ad33a8530db79c6b",
			"d9b9dc5033fb416134e5d2107fdbacab5aadb297cb82dbdcd137d663bac59f7f",
			"",
		}))
		if value != "962e17103e24c7fa63436a9d3f4791d9dfcadf4b8df78be83400f1c0" {
			t.Errorf("test Failed - error Nacl Box don't correct")
		}
	})

	// nacl box open
	t.Run("TestNaclBoxOpen", func(t *testing.T) {
		value, _ := client.Request(NaclBoxOpen(&goton.NaclBoxIn{goton.FixInputMessage("962e17103e24c7fa63436a9d3f4791d9dfcadf4b8df78be83400f1c0", goton.TONInputEncodingHex),
			"cd7f99924bf422544046e83595dd5803f17536f5c9a11746",
			"c4e2d9fe6a6baf8d1812b799856ef2a306291be7a7024837ad33a8530db79c6b",
			"d9b9dc5033fb416134e5d2107fdbacab5aadb297cb82dbdcd137d663bac59f7f",
			goton.TONOutputEncodingText,
		}))
		if value != "Test Message" {
			t.Errorf("test Failed - error Nacl Box Open don't correct")
		}
	})

	// nacl Secret box
	t.Run("TestNaclSecretBox", func(t *testing.T) {
		value, _ := client.Request(NaclSecretBox(&goton.NaclSecretBox{goton.FixInputMessage("Test Message", goton.TONInputEncodingText),
			"2a33564717595ebe53d91a785b9e068aba625c8453a76e45",
			"8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8",
			"",
		}))
		if value != "24bede8ca59ed8a5e6aec9ece35c9f5e8405d2dfc2d50f111b2cd0d8" {
			t.Errorf("test Failed - error Nacl Secret Box don't correct")
		}
	})

	// nacl Secret box open
	t.Run("TestNaclSecretBoxOpen", func(t *testing.T) {
		value, _ := client.Request(NaclSecretBoxOpen(&goton.NaclSecretBox{goton.FixInputMessage("24bede8ca59ed8a5e6aec9ece35c9f5e8405d2dfc2d50f111b2cd0d8", goton.TONInputEncodingHex),
			"2a33564717595ebe53d91a785b9e068aba625c8453a76e45",
			"8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8",
			goton.TONOutputEncodingText}))

		if value != "Test Message" {
			t.Errorf("test Failed - error Nacl Secret Box Open don't correct")
		}
	})

	// nacl Secret box + nacl Secret box open
	t.Run("TestNaclSecretBox+NaclSecretBoxOpen", func(t *testing.T) {
		value1, _ := client.Request(NaclSecretBox(&goton.NaclSecretBox{goton.FixInputMessage(`Text with \' and \" and : {}`, goton.TONInputEncodingText),
			"2a33564717595ebe53d91a785b9e068aba625c8453a76e45",
			"8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8",
			goton.TONOutputEncodingBase64}))

		_, err = client.Request(NaclSecretBoxOpen(&goton.NaclSecretBox{goton.FixInputMessage(value1, goton.TONInputEncodingBase64),
			"2a33564717595ebe53d91a785b9e068aba625c8453a76e45",
			"8f68445b4e78c000fe4d6b7fc826879c1e63e3118379219a754ae66327764bd8",
			goton.TONOutputEncodingText}))

		if err != nil {
			t.Errorf("test Failed - error Nacl Secret Box + Nacl Secret Box Open don't correct")
		}
	})

	// nacl sign
	t.Run("TestNaclSign", func(t *testing.T) {
		value, _ := client.Request(NaclSign(&goton.NaclSign{goton.FixInputMessage("Test Message", goton.TONInputEncodingText),
			"56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e",
			""}))
		if value != "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade10354657374204d657373616765" {
			t.Errorf("test Failed - error Nacl Sign don't correct")
		}
	})

	// naclSignOpen
	t.Run("TestNaclSignOpen", func(t *testing.T) {
		value, _ := client.Request(NaclSignOpen(&goton.NaclSign{goton.FixInputMessage("fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade10354657374204d657373616765", goton.TONInputEncodingHex),
			"1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e",
			goton.TONOutputEncodingText}))
		if value != "Test Message" {
			t.Errorf("test Failed - error Nacl Sign Open don't correct")
		}
	})

	// naclSignDetached
	t.Run("TestNaclSignDetached", func(t *testing.T) {
		value, _ := client.Request(NaclSignDetached(&goton.NaclSign{goton.FixInputMessage("Test Message", goton.TONInputEncodingText),
			"56b6a77093d6fdf14e593f36275d872d75de5b341942376b2a08759f3cbae78f1869b7ef29d58026217e9cf163cbfbd0de889bdf1bf4daebf5433a312f5b8d6e",
			"",
		}))
		if value != "fb0cfe40eea5d6c960652e6ceb904da8a72ee2fcf6e05089cf835203179ff65bb48c57ecf31dcfcd26510bea67e64f3e6898b7c58300dc14338254268cade103" {
			t.Errorf("test Failed - error Nacl Sign Open don't correct")
		}
	})

	t.Run("TestMnemonicWords", func(t *testing.T) {
		value, _ := MnemonicWordsResp(client.Request(MnemonicWords()))
		if len(value) != 2048 {
			t.Errorf("test Failed - error different count words of 2048")
		}
	})

	t.Run("TestMnemonicFromRandom", func(t *testing.T) {
		value, _ := client.Request(MnemonicFromRandom(&goton.MnemonicStructRequest{nil, goton.TONMnemonicDictionary["ENGLISH"], ""}))
		if len(strings.Fields(value)) != 24 {
			t.Errorf("test Failed - length must be 24 words")
		}

		value, _ = client.Request(MnemonicFromRandom(&goton.MnemonicStructRequest{nil, goton.TONMnemonicDictionary["ENGLISH"], ""}))
		if len(strings.Fields(value)) != 24 {
			t.Errorf("test Failed - length must be 24 words")
		}
	})

	t.Run("TestMnemonicFromEntropy", func(t *testing.T) {
		value, _ := client.Request(MnemonicFromEntropy(&goton.MnemonicStructRequest{goton.FixInputMessage("00112233445566778899AABBCCDDEEFF", goton.TONInputEncodingHex),
			goton.TONMnemonicDictionary["ENGLISH"], ""}))
		if value != "abandon math mimic master filter design carbon crystal rookie group knife young" {
			t.Errorf("test Failed - error get mnemonic from entropy value different with constants")
		}
	})

	t.Run("TestMnemonicVerify", func(t *testing.T) {
		value, _ := MnemonicVerifyResp(client.Request(MnemonicVerify(&goton.MnemonicStructRequest{nil, goton.TONMnemonicDictionary["ENGLISH"], "abandon math mimic master filter design carbon crystal rookie group knife young"})))
		if !value {
			t.Errorf("test Failed - mnemonic phrase don't check")
		}
	})

	t.Run("TestMnemonicVerify2", func(t *testing.T) {
		value, _ := MnemonicVerifyResp(client.Request(MnemonicVerify(&goton.MnemonicStructRequest{nil, goton.TONMnemonicDictionary["ENGLISH"], "abandon math mimic"})))
		if value {
			t.Errorf("test Failed - mnemonic phrase don't check")
		}
	})

	t.Run("TestMnemonicDeriveSignKeys", func(t *testing.T) {
		value, _ := GenerateRandomSignKeysResp(client.Request(MnemonicDeriveSignKeys("unit follow zone decline glare flower crisp vocal adapt magic much mesh cherry teach mechanic rain float vicious solution assume hedgehog rail sort chuckle")))
		if value.Public != "c374990ccacb36a87cb016e54fd6fcf0c344e9b0bc6744d9db89f4c272ef9712" {
			t.Errorf("test Failed - public key different")
		}
		if value.Secret != "92328f6ff49bb225167ec94f2b146a9560bdc5f3c4ff416624d60ed6e23e0d01" {
			t.Errorf("test Failed - secret key different")
		}
	})

	// HDKeys

	t.Run("TestHDkeyXprv", func(t *testing.T) {
		master, _ := client.Request(HdkeyXprvFromMnemonic(&goton.MnemonicStructRequest{nil, goton.TONMnemonicDictionary["ENGLISH"], "abuse boss fly battle rubber wasp afraid hamster guide essence vibrant tattoo"}))
		if master != "xprv9s21ZrQH143K25JhKqEwvJW7QAiVvkmi4WRenBZanA6kxHKtKAQQKwZG65kCyW5jWJ8NY9e3GkRoistUjjcpHNsGBUv94istDPXvqGNuWpC" {
			t.Errorf("test Failed - HDKey Xprv from mnemonic")
		}

		valueSecret, _ := client.Request(HdkeyXprvSecret(&goton.HDSerialized{master}))
		if valueSecret != "0c91e53128fa4d67589d63a6c44049c1068ec28a63069a55ca3de30c57f8b365" {
			t.Errorf("test Failed - HDKey Xprv Secret")
		}

		valuePublic, _ := client.Request(HdkeyXprvPublic(&goton.HDSerialized{master}))
		if valuePublic != "02a8eb63085f73c33fa31b4d1134259406347284f8dab6fc68f4bf8c96f6c39b75" {
			t.Errorf("test Failed - HDKey Xprv Public")
		}

		valueDerive, _ := client.Request(HdkeyXprvDerive(&goton.HDDerivery{master, goton.TONMnemonicDictionary["ENGLISH"], false}))
		if valueDerive != "xprv9uZwtSeoKf1syu4qHcHUviGu86r8btERk8ZXy8aQNyNXd2REUH266qPqW5k4rAiyU8UTnJKqsfZgT95i2oSpro7hqK5wUem9FXVEZzKvYKD" {
			t.Errorf("test Failed - HDKey Xprv Derive")
		}

		valueSecretDerivery, _ := client.Request(HdkeyXprvSecret(&goton.HDSerialized{valueDerive}))
		if valueSecretDerivery != "ee692074141fca59f11648ea844b5ffc50e857f232601a988947335ceb33d5f6" {
			t.Errorf("test Failed - HDKey Xprv Secret")
		}

		valuePublicDerivery, _ := client.Request(HdkeyXprvPublic(&goton.HDSerialized{valueDerive}))
		if valuePublicDerivery != "029715758a768381e36d419cdecccd67f40d4b855a1fb985cb2f40c00e14403dfb" {
			t.Errorf("test Failed - HDKey Xprv Secret")
		}

		valuePathDerivery, _ := client.Request(HdkeyXprvDerivePath(&goton.HDPathDerivery{master, "m/44'/60'/0'/0'"}))
		if valuePathDerivery != "xprvA1KNMo63UcGjmDF1bX39Cw2BXGUwrwMjeD5qvQ3tA3qS3mZQkGtpf4DHq8FDLKAvAjXsYGLHDP2dVzLu9ycta8PXLuSYib2T3vzLf3brVgZ" {
			t.Errorf("test Failed - HDKey Xprv Path Derive")
		}

		valueSecretPathDerivery, _ := client.Request(HdkeyXprvSecret(&goton.HDSerialized{valuePathDerivery}))
		if valueSecretPathDerivery != "1c566ade41169763b155761406d3cef08b29b31cf8014f51be08c0cb4e67c5e1" {
			t.Errorf("test Failed - HDKey Xprv Secret")
		}

		valuePublicPathDerivery, _ := client.Request(HdkeyXprvPublic(&goton.HDSerialized{valuePathDerivery}))
		if valuePublicPathDerivery != "02a87d9764eedaacee45b0f777b5a242939b05fa06873bf511ca9a59cb46a5f526" {
			t.Errorf("test Failed - HDKey Xprv Public")
		}
	})
}
