package abi

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/move-ton/ever-client-go/domain"
	"github.com/move-ton/ever-client-go/gateway/client"
	"github.com/move-ton/ever-client-go/usecase/boc"
	"github.com/move-ton/ever-client-go/usecase/crypto"
	"github.com/move-ton/ever-client-go/util"
	"github.com/stretchr/testify/assert"
)

func TestAbi(t *testing.T) {

	configConn := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls(), "")
	clientConn, err := client.NewClientGateway(configConn)
	assert.Equal(t, nil, err)
	defer clientConn.Destroy()

	abiUC := abi{
		config: configConn,
		client: clientConn,
	}
	defer abiUC.client.Destroy()

	var (
		objmap map[string]json.RawMessage
	)

	bocUC := boc.NewBoc(abiUC.config, abiUC.client)

	// Events contract params
	fileAbi, err := os.Open("../samples/Events.abi.json")
	assert.Equal(t, nil, err)
	byteAbi, err := ioutil.ReadAll(fileAbi)
	assert.Equal(t, nil, err)

	nn := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &nn)
	assert.Equal(t, nil, err)
	eventsAbi := domain.NewAbiContract(nn)

	fileTvc, err := os.Open("../samples/Events.tvc")
	assert.Equal(t, nil, err)
	byteTvc, err := ioutil.ReadAll(fileTvc)
	assert.Equal(t, nil, err)
	eventsTvc := base64.StdEncoding.EncodeToString(byteTvc)

	eventsTime := big.NewInt(1599458364291)
	eventsExpire := util.IntToPointerInt(1599458404)
	keyPair := &domain.KeyPair{Public: "4c7c408ff1ddebb8d6405ee979c716a14fdd6cc08124107a61d3c25597099499", Secret: "cc8929d635719612a9478b9cd17675a39cfad52d8959e8a177389b8c0b9122a7"}

	t.Run("TestDecodeMessage", func(t *testing.T) {
		message := "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
		params := &domain.ParamsOfDecodeMessage{Abi: eventsAbi, Message: message}
		decoded, err := abiUC.DecodeMessage(params)
		assert.Equal(t, nil, err)
		assert.Equal(t, domain.MessageBodyTypeInput, decoded.BodyType)

		err = json.Unmarshal(decoded.Value, &objmap)
		assert.Equal(t, nil, err)

		assert.Equal(t, `"0x0000000000000000000000000000000000000000000000000000000000000000"`, string(objmap["id"]))
		assert.Equal(t, eventsExpire, decoded.Header.Expire)
		assert.Equal(t, eventsTime, decoded.Header.Time)
		assert.Equal(t, keyPair.Public, decoded.Header.PubKey)

		message = "te6ccgEBAQEAVQAApeACvg5/pmQpY4m61HmJ0ne+zjHJu3MNG8rJxUDLbHKBu/AAAAAAAAAMJL6z6ro48sYvAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA"
		params = &domain.ParamsOfDecodeMessage{Abi: eventsAbi, Message: message}
		decoded, err = abiUC.DecodeMessage(params)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(decoded.Value, &objmap)
		assert.Equal(t, nil, err)

		assert.Equal(t, domain.MessageBodyTypeEvent, decoded.BodyType)
		assert.Equal(t, `"0x0000000000000000000000000000000000000000000000000000000000000000"`, string(objmap["id"]))
		assert.Nil(t, decoded.Header)

		message = "te6ccgEBAQEAVQAApeACvg5/pmQpY4m61HmJ0ne+zjHJu3MNG8rJxUDLbHKBu/AAAAAAAAAMKr6z6rxK3xYJAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA"
		params = &domain.ParamsOfDecodeMessage{Abi: eventsAbi, Message: message}
		decoded, err = abiUC.DecodeMessage(params)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(decoded.Value, &objmap)
		assert.Equal(t, nil, err)

		assert.Equal(t, domain.MessageBodyTypeOutput, decoded.BodyType)
		assert.Equal(t, `"0x0000000000000000000000000000000000000000000000000000000000000000"`, string(objmap["value0"]))
		assert.Nil(t, decoded.Header)

		params = &domain.ParamsOfDecodeMessage{Abi: eventsAbi, Message: "0x0"}
		decoded, err = abiUC.DecodeMessage(params)
		assert.NotEqual(t, nil, err)
	})

	t.Run("TestDecodeMessageBody", func(t *testing.T) {
		message := "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
		paramsBocObj := &domain.ParamsOfParse{Boc: message}
		bocObj, err := bocUC.ParseMessage(paramsBocObj)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(bocObj.Parsed, &objmap)
		assert.Equal(t, nil, err)

		var body string
		bytesBody, err := objmap["body"].MarshalJSON()
		assert.Equal(t, nil, err)

		err = json.Unmarshal(bytesBody, &body)
		assert.Equal(t, nil, err)

		decodeParams := &domain.ParamsOfDecodeMessageBody{Abi: eventsAbi, Body: body, IsInternal: false}
		decoded, err := abiUC.DecodeMessageBody(decodeParams)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(decoded.Value, &objmap)
		assert.Equal(t, nil, err)

		assert.Equal(t, domain.MessageBodyTypeInput, decoded.BodyType)
		assert.Equal(t, `"0x0000000000000000000000000000000000000000000000000000000000000000"`, string(objmap["id"]))
		assert.Equal(t, eventsExpire, decoded.Header.Expire)
		assert.Equal(t, eventsTime, decoded.Header.Time)
		assert.Equal(t, keyPair.Public, decoded.Header.PubKey)
	})

	t.Run("TestEncodeMessage", func(t *testing.T) {
		deploySet := &domain.DeploySet{Tvc: eventsTvc}
		callSet := &domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keyPair.Public, Time: eventsTime, Expire: eventsExpire}}
		signer := domain.NewSigner(domain.SignerExternal{PublicKey: keyPair.Public})

		// Create unsigned deployment message
		unsigned, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: eventsAbi, Signer: signer, DeploySet: deploySet, CallSet: callSet})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECFwEAA2gAAqeIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEZTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gGAQEBwAICA88gBQMBAd4EAAPQIABB2mPiBH+O713GsgL3S844tQp+62YECSCD0w6eEqy4TKTMAib/APSkICLAAZL0oOGK7VNYMPShCQcBCvSkIPShCAAAAgEgDAoByP9/Ie1E0CDXScIBjhDT/9M/0wDRf/hh+Gb4Y/hijhj0BXABgED0DvK91wv/+GJw+GNw+GZ/+GHi0wABjh2BAgDXGCD5AQHTAAGU0/8DAZMC+ELiIPhl+RDyqJXTAAHyeuLTPwELAGqOHvhDIbkgnzAg+COBA+iogggbd0Cgud6S+GPggDTyNNjTHwH4I7zyudMfAfAB+EdukvI83gIBIBINAgEgDw4AvbqLVfP/hBbo417UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLe+Ebyc3H4ZtH4APhCyMv/+EPPCz/4Rs8LAMntVH/4Z4AgEgERAA5biABrW/CC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb30gyupo6H0gb+j8IpA3SRg4b3whXXlwMnwAZGT9ghBkZ8KEZ0aCBAfQAAAAAAAAAAAAAAAAACBni2TAgEB9gBh8IWRl//wh54Wf/CNnhYBk9qo//DPAAxbmTwqLfCC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb2uG/8rqaOhp/+/o/ABkRe4AAAAAAAAAAAAAAAAIZ4tnwOfI48sYvRDnhf/kuP2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8AIBSBYTAQm4t8WCUBQB/PhBbo4T7UTQ0//TP9MA0X/4Yfhm+GP4Yt7XDf+V1NHQ0//f0fgAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPkceWMXohzwv/yXH7AMiL3AAAAAAAAAAAAAAAABDPFs+Bz5JW+LBKIc8L/8lx+wAw+ELIy//4Q88LP/hGzwsAye1UfxUABPhnAHLccCLQ1gIx0gAw3CHHAJLyO+Ah1w0fkvI84VMRkvI74cEEIoIQ/////byxkvI84AHwAfhHbpLyPN4=", unsigned.Message)
		assert.Equal(t, "KCGM36iTYuCYynk+Jnemis+mcwi3RFCke95i7l96s4Q=", unsigned.DataToSign)

		cryptoUC := crypto.NewCrypto(abiUC.config, abiUC.client)

		// Create detached signature
		signature, err := cryptoUC.Sign(&domain.ParamsOfSign{Unsigned: unsigned.DataToSign, Keys: keyPair})
		assert.Equal(t, nil, err)
		assert.Equal(t, "6272357bccb601db2b821cb0f5f564ab519212d242cf31961fe9a3c50a30b236012618296b4f769355c0e9567cd25b366f3c037435c498c82e5305622adbc70e", signature.Signature)

		// Attach signature to unsigned message
		signed, err := abiUC.AttachSignature(&domain.ParamsOfAttachSignature{Abi: eventsAbi, PublicKey: keyPair.Public, Message: unsigned.Message, Signature: signature.Signature})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECGAEAA6wAA0eIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEbAHAgEA4bE5Gr3mWwDtlcEOWHr6slWoyQlpIWeYyw/00eKFGFkbAJMMFLWnu0mq4HSrPmktmzeeAboa4kxkFymCsRVt44dTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gAQHAAwIDzyAGBAEB3gUAA9AgAEHaY+IEf47vXcayAvdLzji1Cn7rZgQJIIPTDp4SrLhMpMwCJv8A9KQgIsABkvSg4YrtU1gw9KEKCAEK9KQg9KEJAAACASANCwHI/38h7UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLTAAGOHYECANcYIPkBAdMAAZTT/wMBkwL4QuIg+GX5EPKoldMAAfJ64tM/AQwAao4e+EMhuSCfMCD4I4ED6KiCCBt3QKC53pL4Y+CANPI02NMfAfgjvPK50x8B8AH4R26S8jzeAgEgEw4CASAQDwC9uotV8/+EFujjXtRNAg10nCAY4Q0//TP9MA0X/4Yfhm+GP4Yo4Y9AVwAYBA9A7yvdcL//hicPhjcPhmf/hh4t74RvJzcfhm0fgA+ELIy//4Q88LP/hGzwsAye1Uf/hngCASASEQDluIAGtb8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFvfSDK6mjofSBv6PwikDdJGDhvfCFdeXAyfABkZP2CEGRnwoRnRoIEB9AAAAAAAAAAAAAAAAAAIGeLZMCAQH2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8ADFuZPCot8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFva4b/yupo6Gn/7+j8AGRF7gAAAAAAAAAAAAAAAAhni2fA58jjyxi9EOeF/+S4/YAYfCFkZf/8IeeFn/wjZ4WAZPaqP/wzwAgFIFxQBCbi3xYJQFQH8+EFujhPtRNDT/9M/0wDRf/hh+Gb4Y/hi3tcN/5XU0dDT/9/R+ADIi9wAAAAAAAAAAAAAAAAQzxbPgc+Rx5YxeiHPC//JcfsAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPklb4sEohzwv/yXH7ADD4QsjL//hDzws/+EbPCwDJ7VR/FgAE+GcActxwItDWAjHSADDcIccAkvI74CHXDR+S8jzhUxGS8jvhwQQighD////9vLGS8jzgAfAB+EdukvI83g==", signed.Message)

		// Create initially signed message
		signerKeys := domain.NewSigner(domain.SignerKeys{Keys: keyPair})
		signed2, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: eventsAbi, DeploySet: deploySet, CallSet: callSet, Signer: signerKeys})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECGAEAA6wAA0eIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEbAHAgEA4bE5Gr3mWwDtlcEOWHr6slWoyQlpIWeYyw/00eKFGFkbAJMMFLWnu0mq4HSrPmktmzeeAboa4kxkFymCsRVt44dTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gAQHAAwIDzyAGBAEB3gUAA9AgAEHaY+IEf47vXcayAvdLzji1Cn7rZgQJIIPTDp4SrLhMpMwCJv8A9KQgIsABkvSg4YrtU1gw9KEKCAEK9KQg9KEJAAACASANCwHI/38h7UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLTAAGOHYECANcYIPkBAdMAAZTT/wMBkwL4QuIg+GX5EPKoldMAAfJ64tM/AQwAao4e+EMhuSCfMCD4I4ED6KiCCBt3QKC53pL4Y+CANPI02NMfAfgjvPK50x8B8AH4R26S8jzeAgEgEw4CASAQDwC9uotV8/+EFujjXtRNAg10nCAY4Q0//TP9MA0X/4Yfhm+GP4Yo4Y9AVwAYBA9A7yvdcL//hicPhjcPhmf/hh4t74RvJzcfhm0fgA+ELIy//4Q88LP/hGzwsAye1Uf/hngCASASEQDluIAGtb8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFvfSDK6mjofSBv6PwikDdJGDhvfCFdeXAyfABkZP2CEGRnwoRnRoIEB9AAAAAAAAAAAAAAAAAAIGeLZMCAQH2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8ADFuZPCot8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFva4b/yupo6Gn/7+j8AGRF7gAAAAAAAAAAAAAAAAhni2fA58jjyxi9EOeF/+S4/YAYfCFkZf/8IeeFn/wjZ4WAZPaqP/wzwAgFIFxQBCbi3xYJQFQH8+EFujhPtRNDT/9M/0wDRf/hh+Gb4Y/hi3tcN/5XU0dDT/9/R+ADIi9wAAAAAAAAAAAAAAAAQzxbPgc+Rx5YxeiHPC//JcfsAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPklb4sEohzwv/yXH7ADD4QsjL//hDzws/+EbPCwDJ7VR/FgAE+GcActxwItDWAjHSADDcIccAkvI74CHXDR+S8jzhUxGS8jvhwQQighD////9vLGS8jzgAfAB+EdukvI83g==", signed2.Message)

		// Sign with signing box
		sBox, err := cryptoUC.GetSigningBox(keyPair)
		assert.Equal(t, nil, err)

		signedBox, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: eventsAbi, Signer: domain.NewSigner(domain.SignerSigningBox{Handle: sBox.Handle}), DeploySet: deploySet, CallSet: callSet})
		assert.Equal(t, nil, err)

		err = cryptoUC.RemoveSigningBox(sBox)
		assert.Equal(t, nil, err)

		assert.Equal(t, signed2.Message, signedBox.Message)

		// Create run unsigned message
		address := "0:05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309"
		callSet = &domain.CallSet{
			FunctionName: "returnValue",
			Input:        json.RawMessage(`{"id": "0"}`),
			Header: &domain.FunctionHeader{
				PubKey: keyPair.Public,
				Time:   eventsTime,
				Expire: eventsExpire},
		}
		unsigned, err = abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: eventsAbi, CallSet: callSet, Signer: signer, Address: address})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAgEAeAABpYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIFMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKAQBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", unsigned.Message)
		assert.Equal(t, "i4Hs3PB12QA9UBFbOIpkG3JerHHqjm4LgvF4MA7TDsY=", unsigned.DataToSign)

		// Create detached signature
		signature, err = cryptoUC.Sign(&domain.ParamsOfSign{Unsigned: unsigned.DataToSign, Keys: keyPair})
		assert.Equal(t, nil, err)
		assert.Equal(t, "5bbfb7f184f2cb5f019400b9cd497eeaa41f3d5885619e9f7d4fab8dd695f4b3a02159a1422996c1dd7d1be67898bc79c6adba6c65a18101ac5f0a2a2bb8910b", signature.Signature)

		// Attach signature
		signed, err = abiUC.AttachSignature(&domain.ParamsOfAttachSignature{Abi: eventsAbi, PublicKey: keyPair.Public, Message: unsigned.Message, Signature: signature.Signature})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", signed.Message)

		// Create initially signed message
		signed2, err = abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: eventsAbi, CallSet: callSet, Signer: signerKeys, Address: address})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", signed2.Message)
	})

	t.Run("TestEncodeMessageBody", func(t *testing.T) {
		header := domain.FunctionHeader{Expire: eventsExpire, Time: eventsTime, PubKey: keyPair.Public}
		callSet := &domain.CallSet{FunctionName: "returnValue", Header: &header, Input: json.RawMessage(`{"id":"0"}`)}
		signer := domain.NewSigner(domain.SignerKeys{Keys: keyPair})
		encoded, err := abiUC.EncodeMessageBody(&domain.ParamsOfEncodeMessageBody{Abi: eventsAbi, CallSet: callSet, IsInternal: false, Signer: signer})
		assert.Equal(t, nil, err)
		assert.Equal(t, "", encoded.DataToSign)
	})

	t.Run("TestEncodeAccount", func(t *testing.T) {
		// Encode account from encoded deploy message
		encodedDeployMessage := "te6ccgECFwEAA2gAAqeIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEZTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gGAQEBwAICA88gBQMBAd4EAAPQIABB2mPiBH+O713GsgL3S844tQp+62YECSCD0w6eEqy4TKTMAib/APSkICLAAZL0oOGK7VNYMPShCQcBCvSkIPShCAAAAgEgDAoByP9/Ie1E0CDXScIBjhDT/9M/0wDRf/hh+Gb4Y/hijhj0BXABgED0DvK91wv/+GJw+GNw+GZ/+GHi0wABjh2BAgDXGCD5AQHTAAGU0/8DAZMC+ELiIPhl+RDyqJXTAAHyeuLTPwELAGqOHvhDIbkgnzAg+COBA+iogggbd0Cgud6S+GPggDTyNNjTHwH4I7zyudMfAfAB+EdukvI83gIBIBINAgEgDw4AvbqLVfP/hBbo417UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLe+Ebyc3H4ZtH4APhCyMv/+EPPCz/4Rs8LAMntVH/4Z4AgEgERAA5biABrW/CC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb30gyupo6H0gb+j8IpA3SRg4b3whXXlwMnwAZGT9ghBkZ8KEZ0aCBAfQAAAAAAAAAAAAAAAAACBni2TAgEB9gBh8IWRl//wh54Wf/CNnhYBk9qo//DPAAxbmTwqLfCC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb2uG/8rqaOhp/+/o/ABkRe4AAAAAAAAAAAAAAAAIZ4tnwOfI48sYvRDnhf/kuP2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8AIBSBYTAQm4t8WCUBQB/PhBbo4T7UTQ0//TP9MA0X/4Yfhm+GP4Yt7XDf+V1NHQ0//f0fgAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPkceWMXohzwv/yXH7AMiL3AAAAAAAAAAAAAAAABDPFs+Bz5JW+LBKIc8L/8lx+wAw+ELIy//4Q88LP/hGzwsAye1UfxUABPhnAHLccCLQ1gIx0gAw3CHHAJLyO+Ah1w0fkvI84VMRkvI74cEEIoIQ/////byxkvI84AHwAfhHbpLyPN4="
		messageSource := domain.NewMessageSource(domain.MessageSourceEncoded{Message: encodedDeployMessage, Abi: eventsAbi})
		stateInitSourceSS := domain.NewStateInitSource(domain.StateInitSourceMessage{Source: messageSource})
		encoded, err := abiUC.EncodeAccount(&domain.ParamsOfEncodeAccount{StateInit: stateInitSourceSS})
		assert.Equal(t, nil, err)
		assert.Equal(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", encoded.ID)

		// Encode account from encoding params
		deploySet := domain.DeploySet(domain.DeploySet{Tvc: eventsTvc})
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{Expire: eventsExpire, Time: eventsTime, PubKey: keyPair.Public}}
		signerKeys := domain.NewSigner(domain.SignerKeys{Keys: keyPair})

		messageSourceEnconing := domain.NewMessageSource(domain.MessageSourceEncodingParams{&domain.ParamsOfEncodeMessage{Abi: eventsAbi, Signer: signerKeys, DeploySet: &deploySet, CallSet: &callSet}})
		stateInitSourceSS = domain.NewStateInitSource(domain.StateInitSourceMessage{Source: messageSourceEnconing})
		encoded, err = abiUC.EncodeAccount(&domain.ParamsOfEncodeAccount{StateInit: stateInitSourceSS})
		assert.Equal(t, nil, err)
		assert.Equal(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", encoded.ID)

		// Test exception (external signer)
		signerExt := domain.NewSigner(domain.SignerExternal{PublicKey: keyPair.Public})
		msEP := domain.NewMessageSource(domain.MessageSourceEncodingParams{&domain.ParamsOfEncodeMessage{Abi: eventsAbi, Signer: signerExt, DeploySet: &deploySet, CallSet: &callSet}})
		stateInitSourceSS = domain.NewStateInitSource(msEP)
		_, err = abiUC.EncodeAccount(&domain.ParamsOfEncodeAccount{StateInit: stateInitSourceSS})
		assert.NotEqual(t, nil, err)

		// Encode account from TVC
		sIST := domain.StateInitSourceTvc{Tvc: eventsTvc, PublicKey: "", InitParams: nil}
		stateSI := domain.NewStateInitSource(sIST)
		encoded, err = abiUC.EncodeAccount(&domain.ParamsOfEncodeAccount{StateInit: stateSI})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", encoded.ID)

		sIST.PublicKey = keyPair.Public
		encoded, err = abiUC.EncodeAccount(&domain.ParamsOfEncodeAccount{StateInit: domain.NewStateInitSource(sIST)})
		assert.Equal(t, nil, err)
		assert.Equal(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", encoded.ID)
	})

	// Hello params
	fileAbiHello, err := os.Open("../samples/Hello.abi.json")
	assert.Equal(t, nil, err)
	byteAbiHello, err := ioutil.ReadAll(fileAbiHello)
	assert.Equal(t, nil, err)

	nn2 := &domain.AbiContract{}
	err = json.Unmarshal(byteAbiHello, &nn2)
	assert.Equal(t, nil, err)
	helloAbi := domain.NewAbiContract(nn2)

	fileHelloTvc, err := os.Open("../samples/Hello.tvc")
	assert.Equal(t, nil, err)
	byteHelloTvc, err := ioutil.ReadAll(fileHelloTvc)
	assert.Equal(t, nil, err)
	helloTvc := base64.StdEncoding.EncodeToString(byteHelloTvc)

	t.Run("TestEncodeInternalMessageRun", func(t *testing.T) {
		address := "0:1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
		result, err := abiUC.EncodeInternalMessage(&domain.ParamsOfEncodeInternalMessage{Abi: helloAbi, Value: "1000000000", Address: address, CallSet: &domain.CallSet{FunctionName: "sayHello"}, Bounce: util.BoolToPointerBool(true)})
		assert.Equal(t, nil, err)
		assert.Equal(t, address, result.Address)
		assert.Equal(t, "te6ccgEBAQEAOgAAcGIACRorPEhV5veJGis8SFXm94kaKzxIVeb3iRorPEhV5veh3NZQAAAAAAAAAAAAAAAAAABQy+0X", result.Message)

		bocHash, err := bocUC.GetBocHash(&domain.ParamsOfGetBocHash{Boc: result.Message})
		assert.Equal(t, nil, err)
		assert.Equal(t, result.MessageID, bocHash.Hash)

		parsed, err := bocUC.ParseMessage(&domain.ParamsOfParse{Boc: result.Message})
		assert.Equal(t, nil, err)

		err = json.Unmarshal(parsed.Parsed, &objmap)
		assert.Equal(t, nil, err)

		assert.Equal(t, `"internal"`, string(objmap["msg_type_name"]))
		assert.Equal(t, `""`, string(objmap["src"]))
		assert.Equal(t, `"`+address+`"`, string(objmap["dst"]))
		assert.Equal(t, `"0x3b9aca00"`, string(objmap["value"]))
		assert.Equal(t, "true", string(objmap["bounce"]))
		assert.Equal(t, "true", string(objmap["ihr_disabled"]))
	})

	t.Run("TestEncodeInternalMessageDeploy", func(t *testing.T) {
		result, err := abiUC.EncodeInternalMessage(&domain.ParamsOfEncodeInternalMessage{
			Abi:       helloAbi,
			Value:     "0",
			DeploySet: &domain.DeploySet{Tvc: helloTvc},
			CallSet:   &domain.CallSet{FunctionName: "constructor"}})
		assert.Equal(t, nil, err)
		assert.Equal(t, `te6ccgECHAEABG0AAmliADYO5IoxskLmUfURre2fOB04OmP32VjPwA/lDM/Cpvh8AAAAAAAAAAAAAAAAAAIxotV8/gYBAQHAAgIDzyAFAwEB3gQAA9AgAEHYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQCJv8A9KQgIsABkvSg4YrtU1gw9KEJBwEK9KQg9KEIAAACASAMCgHo/38h0wABjiaBAgDXGCD5AQFw7UTQ9AWAQPQO8orXC/8B7Ucib3XtVwMB+RDyqN7tRNAg10nCAY4W9ATTP9MA7UcBb3EBb3YBb3MBb3LtV44Y9AXtRwFvcnBvc3BvdsiAIM9AydBvce1X4tM/Ae1HbxMhuSALAGCfMCD4I4ED6KiCCBt3QKC53pntRyFvUyDtVzCUgDTy8OIw0x8B+CO88rnTHwHxQAECASAYDQIBIBEOAQm6i1Xz+A8B+u1Hb2FujjvtRNAg10nCAY4W9ATTP9MA7UcBb3EBb3YBb3MBb3LtV44Y9AXtRwFvcnBvc3BvdsiAIM9AydBvce1X4t7tR28WkvIzl+1HcW9W7VfiAPgA0fgjtR/tRyBvETAByMsfydBvUe1X7UdvEsj0AO1HbxPPCz/tR28WEAAczwsA7UdvEc8Wye1UcGoCAWoVEgEJtAAa1sATAfztR29hbo477UTQINdJwgGOFvQE0z/TAO1HAW9xAW92AW9zAW9y7VeOGPQF7UcBb3Jwb3Nwb3bIgCDPQMnQb3HtV+Le7UdvZSBukjBw3nDtR28SgED0DvKK1wv/uvLgZPgA+kDRIMjJ+wSBA+hwgQCAyHHPCwEizwoAcc9A+CgUAI7PFiTPFiP6AnHPQHD6AnD6AoBAz0D4I88LH3LPQCDJIvsAXwUw7UdvEsj0AO1HbxPPCz/tR28WzwsA7UdvEc8Wye1UcGrbMAEJtGX2i8AWAfjtR29hbo477UTQINdJwgGOFvQE0z/TAO1HAW9xAW92AW9zAW9y7VeOGPQF7UcBb3Jwb3Nwb3bIgCDPQMnQb3HtV+Le0e1HbxHXCx/IghBQy+0XghCAAAAAsc8LHyHPCx/Ic88LAfgozxZyz0D4Jc8LP4Ahz0AgzzUizzG8FwB4lnHPQCHPF5Vxz0EhzeIgyXH7AFshwP+OHu1HbxLI9ADtR28Tzws/7UdvFs8LAO1HbxHPFsntVN5xatswAgEgGxkBCbtzEuRYGgD47UdvYW6OO+1E0CDXScIBjhb0BNM/0wDtRwFvcQFvdgFvcwFvcu1Xjhj0Be1HAW9ycG9zcG92yIAgz0DJ0G9x7Vfi3vgA0fgjtR/tRyBvETAByMsfydBvUe1X7UdvEsj0AO1HbxPPCz/tR28WzwsA7UdvEc8Wye1UcGrbMADK3XAh10kgwSCOKyDAAI4cI9Bz1yHXCwAgwAGW2zBfB9swltswXwfbMOME2ZbbMF8G2zDjBNngItMfNCB0uyCOFTAgghD/////uiCZMCCCEP////6639+W2zBfB9sw4CMh8UABXwc=`, result.Message)

		bocHash, err := bocUC.GetBocHash(&domain.ParamsOfGetBocHash{Boc: result.Message})
		assert.Equal(t, nil, err)
		assert.Equal(t, result.MessageID, bocHash.Hash)

		parsed, err := bocUC.ParseMessage(&domain.ParamsOfParse{Boc: result.Message})
		assert.Equal(t, nil, err)

		err = json.Unmarshal(parsed.Parsed, &objmap)
		assert.Equal(t, nil, err)

		code, err := bocUC.GetCodeFromTvc(&domain.ParamsOfGetCodeFromTvc{Tvc: helloTvc})
		assert.Equal(t, nil, err)
		assert.Equal(t, string(objmap["code"]), `"`+code.Code+`"`)
	})
}
