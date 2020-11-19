package abi

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/usecase/boc"
	"github.com/move-ton/ton-client-go/usecase/crypto"
	"github.com/stretchr/testify/assert"
)

func TestAbi(t *testing.T) {

	client, err := client.NewClientGateway(domain.NewDefaultConfig(2))
	assert.Equal(t, nil, err)
	defer client.Destroy()

	abiUC := abi{
		config: domain.NewDefaultConfig(2),
		client: client,
	}
	defer abiUC.client.Destroy()

	fileAbi, err := os.Open("../samples/Events.abi.json")
	assert.Equal(t, nil, err)
	byteAbi, err := ioutil.ReadAll(fileAbi)
	assert.Equal(t, nil, err)

	eventsAbi := domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &eventsAbi)
	assert.Equal(t, nil, err)

	fileTvc, err := os.Open("../samples/Events.tvc")
	assert.Equal(t, nil, err)
	byteTvc, err := ioutil.ReadAll(fileTvc)
	assert.Equal(t, nil, err)

	eventsTvc := base64.StdEncoding.EncodeToString(byteTvc)
	eventsTime := 1599458364291
	eventsExpire := 1599458404
	keyPair := domain.KeyPair{Public: "4c7c408ff1ddebb8d6405ee979c716a14fdd6cc08124107a61d3c25597099499", Secret: "cc8929d635719612a9478b9cd17675a39cfad52d8959e8a177389b8c0b9122a7"}

	abiValue := domain.NewAbiContract()
	abiValue.Value = eventsAbi

	t.Run("TestDecodeMessage", func(t *testing.T) {
		var (
			objmap map[string]json.RawMessage
			id     string
		)

		message := "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
		valueDM, err := abiUC.DecodeMessage(domain.ParamsOfDecodeMessage{Abi: abiValue, Message: message})
		assert.Equal(t, nil, err)
		assert.Equal(t, domain.MessageBodyTypeInput, valueDM.BodyType)

		err = json.Unmarshal(valueDM.Value, &objmap)
		assert.Equal(t, nil, err)
		bytesID, err := objmap["id"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesID, &id)
		assert.Equal(t, nil, err)
		assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", id)
		assert.Equal(t, eventsExpire, valueDM.Header.Expire)
		assert.Equal(t, eventsTime, valueDM.Header.Time)
		assert.Equal(t, keyPair.Public, valueDM.Header.PubKey)
		id = ""

		message = "te6ccgEBAQEAVQAApeACvg5/pmQpY4m61HmJ0ne+zjHJu3MNG8rJxUDLbHKBu/AAAAAAAAAMJL6z6ro48sYvAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA"
		valueDM2, err := abiUC.DecodeMessage(domain.ParamsOfDecodeMessage{Abi: abiValue, Message: message})
		assert.Equal(t, nil, err)
		assert.Equal(t, domain.MessageBodyTypeEvent, valueDM2.BodyType)
		err = json.Unmarshal(valueDM2.Value, &objmap)
		assert.Equal(t, nil, err)
		bytesID, err = objmap["id"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesID, &id)
		assert.Equal(t, nil, err)
		assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", id)
		assert.NotEqual(t, nil, valueDM2.Header)

		message = "te6ccgEBAQEAVQAApeACvg5/pmQpY4m61HmJ0ne+zjHJu3MNG8rJxUDLbHKBu/AAAAAAAAAMKr6z6rxK3xYJAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA"
		valueDM3, err := abiUC.DecodeMessage(domain.ParamsOfDecodeMessage{Abi: abiValue, Message: message})
		assert.Equal(t, nil, err)

		assert.Equal(t, domain.MessageBodyTypeOutput, valueDM3.BodyType)
		err = json.Unmarshal(valueDM3.Value, &objmap)
		assert.Equal(t, nil, err)
		bytesID, err = objmap["id"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesID, &id)
		assert.Equal(t, nil, err)
		assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", id)
		assert.NotEqual(t, nil, valueDM2.Header)

		_, err = abiUC.DecodeMessage(domain.ParamsOfDecodeMessage{Abi: abiValue, Message: "0x0"})
		assert.NotEqual(t, nil, err)
	})

	t.Run("TestDecodeMessageBody", func(t *testing.T) {
		message := "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
		bocUC := boc.NewBoc(abiUC.config, abiUC.client)
		valueBoc, err := bocUC.ParseMessage(domain.ParamsOfParse{Boc: message})
		assert.Equal(t, nil, err)
		var (
			objmap   map[string]json.RawMessage
			body, id string
		)
		err = json.Unmarshal(valueBoc.Parsed, &objmap)
		assert.Equal(t, nil, err)

		bytesBody, err := objmap["body"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesBody, &body)
		assert.Equal(t, nil, err)

		valueDMB1, err := abiUC.DecodeMessageBody(domain.ParamsOfDecodeMessageBody{Abi: abiValue, Body: body})
		assert.Equal(t, nil, err)
		assert.Equal(t, domain.MessageBodyTypeInput, valueDMB1.BodyType)

		err = json.Unmarshal(valueDMB1.Value, &objmap)
		assert.Equal(t, nil, err)
		bytesID, err := objmap["id"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesID, &id)
		assert.Equal(t, nil, err)
		assert.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000000", id)
		assert.Equal(t, eventsExpire, valueDMB1.Header.Expire)
		assert.Equal(t, eventsTime, valueDMB1.Header.Time)
		assert.Equal(t, keyPair.Public, valueDMB1.Header.PubKey)
	})

	t.Run("TestEncodeMessage", func(t *testing.T) {
		deploySet := domain.DeploySet(domain.DeploySet{Tvc: eventsTvc})
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{Expire: eventsExpire, Time: eventsTime, PubKey: keyPair.Public}}
		signer := domain.NewSignerExternal()
		signer.PublicKey = keyPair.Public

		valueUnsigned, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue, DeploySet: &deploySet, CallSet: &callSet, Signer: signer})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECFwEAA2gAAqeIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEZTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gGAQEBwAICA88gBQMBAd4EAAPQIABB2mPiBH+O713GsgL3S844tQp+62YECSCD0w6eEqy4TKTMAib/APSkICLAAZL0oOGK7VNYMPShCQcBCvSkIPShCAAAAgEgDAoByP9/Ie1E0CDXScIBjhDT/9M/0wDRf/hh+Gb4Y/hijhj0BXABgED0DvK91wv/+GJw+GNw+GZ/+GHi0wABjh2BAgDXGCD5AQHTAAGU0/8DAZMC+ELiIPhl+RDyqJXTAAHyeuLTPwELAGqOHvhDIbkgnzAg+COBA+iogggbd0Cgud6S+GPggDTyNNjTHwH4I7zyudMfAfAB+EdukvI83gIBIBINAgEgDw4AvbqLVfP/hBbo417UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLe+Ebyc3H4ZtH4APhCyMv/+EPPCz/4Rs8LAMntVH/4Z4AgEgERAA5biABrW/CC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb30gyupo6H0gb+j8IpA3SRg4b3whXXlwMnwAZGT9ghBkZ8KEZ0aCBAfQAAAAAAAAAAAAAAAAACBni2TAgEB9gBh8IWRl//wh54Wf/CNnhYBk9qo//DPAAxbmTwqLfCC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb2uG/8rqaOhp/+/o/ABkRe4AAAAAAAAAAAAAAAAIZ4tnwOfI48sYvRDnhf/kuP2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8AIBSBYTAQm4t8WCUBQB/PhBbo4T7UTQ0//TP9MA0X/4Yfhm+GP4Yt7XDf+V1NHQ0//f0fgAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPkceWMXohzwv/yXH7AMiL3AAAAAAAAAAAAAAAABDPFs+Bz5JW+LBKIc8L/8lx+wAw+ELIy//4Q88LP/hGzwsAye1UfxUABPhnAHLccCLQ1gIx0gAw3CHHAJLyO+Ah1w0fkvI84VMRkvI74cEEIoIQ/////byxkvI84AHwAfhHbpLyPN4=", valueUnsigned.Message)
		assert.Equal(t, "KCGM36iTYuCYynk+Jnemis+mcwi3RFCke95i7l96s4Q=", valueUnsigned.DataToSign)

		cryptoUC := crypto.NewCrypto(abiUC.config, abiUC.client)

		// # Create detached signature
		valueSignature, err := cryptoUC.Sign(domain.ParamsOfSign{Unsigned: valueUnsigned.DataToSign, Keys: keyPair})
		assert.Equal(t, nil, err)
		assert.Equal(t, "6272357bccb601db2b821cb0f5f564ab519212d242cf31961fe9a3c50a30b236012618296b4f769355c0e9567cd25b366f3c037435c498c82e5305622adbc70e", valueSignature.Signature)

		// # Attach signature to unsigned message
		valueSigned, err := abiUC.AttachSignature(domain.ParamsOfAttachSignature{Abi: abiValue, PublicKey: keyPair.Public, Message: valueUnsigned.Message, Signature: valueSignature.Signature})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECGAEAA6wAA0eIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEbAHAgEA4bE5Gr3mWwDtlcEOWHr6slWoyQlpIWeYyw/00eKFGFkbAJMMFLWnu0mq4HSrPmktmzeeAboa4kxkFymCsRVt44dTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gAQHAAwIDzyAGBAEB3gUAA9AgAEHaY+IEf47vXcayAvdLzji1Cn7rZgQJIIPTDp4SrLhMpMwCJv8A9KQgIsABkvSg4YrtU1gw9KEKCAEK9KQg9KEJAAACASANCwHI/38h7UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLTAAGOHYECANcYIPkBAdMAAZTT/wMBkwL4QuIg+GX5EPKoldMAAfJ64tM/AQwAao4e+EMhuSCfMCD4I4ED6KiCCBt3QKC53pL4Y+CANPI02NMfAfgjvPK50x8B8AH4R26S8jzeAgEgEw4CASAQDwC9uotV8/+EFujjXtRNAg10nCAY4Q0//TP9MA0X/4Yfhm+GP4Yo4Y9AVwAYBA9A7yvdcL//hicPhjcPhmf/hh4t74RvJzcfhm0fgA+ELIy//4Q88LP/hGzwsAye1Uf/hngCASASEQDluIAGtb8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFvfSDK6mjofSBv6PwikDdJGDhvfCFdeXAyfABkZP2CEGRnwoRnRoIEB9AAAAAAAAAAAAAAAAAAIGeLZMCAQH2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8ADFuZPCot8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFva4b/yupo6Gn/7+j8AGRF7gAAAAAAAAAAAAAAAAhni2fA58jjyxi9EOeF/+S4/YAYfCFkZf/8IeeFn/wjZ4WAZPaqP/wzwAgFIFxQBCbi3xYJQFQH8+EFujhPtRNDT/9M/0wDRf/hh+Gb4Y/hi3tcN/5XU0dDT/9/R+ADIi9wAAAAAAAAAAAAAAAAQzxbPgc+Rx5YxeiHPC//JcfsAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPklb4sEohzwv/yXH7ADD4QsjL//hDzws/+EbPCwDJ7VR/FgAE+GcActxwItDWAjHSADDcIccAkvI74CHXDR+S8jzhUxGS8jvhwQQighD////9vLGS8jzgAfAB+EdukvI83g==", valueSigned.Message)

		// # Create initially signed message
		signerKeys := domain.NewSignerKeys()
		signerKeys.Keys = keyPair
		valueSigned2, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue, DeploySet: &deploySet, CallSet: &callSet, Signer: signerKeys})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgECGAEAA6wAA0eIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEbAHAgEA4bE5Gr3mWwDtlcEOWHr6slWoyQlpIWeYyw/00eKFGFkbAJMMFLWnu0mq4HSrPmktmzeeAboa4kxkFymCsRVt44dTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gAQHAAwIDzyAGBAEB3gUAA9AgAEHaY+IEf47vXcayAvdLzji1Cn7rZgQJIIPTDp4SrLhMpMwCJv8A9KQgIsABkvSg4YrtU1gw9KEKCAEK9KQg9KEJAAACASANCwHI/38h7UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLTAAGOHYECANcYIPkBAdMAAZTT/wMBkwL4QuIg+GX5EPKoldMAAfJ64tM/AQwAao4e+EMhuSCfMCD4I4ED6KiCCBt3QKC53pL4Y+CANPI02NMfAfgjvPK50x8B8AH4R26S8jzeAgEgEw4CASAQDwC9uotV8/+EFujjXtRNAg10nCAY4Q0//TP9MA0X/4Yfhm+GP4Yo4Y9AVwAYBA9A7yvdcL//hicPhjcPhmf/hh4t74RvJzcfhm0fgA+ELIy//4Q88LP/hGzwsAye1Uf/hngCASASEQDluIAGtb8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFvfSDK6mjofSBv6PwikDdJGDhvfCFdeXAyfABkZP2CEGRnwoRnRoIEB9AAAAAAAAAAAAAAAAAAIGeLZMCAQH2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8ADFuZPCot8ILdHCfaiaGn/6Z/pgGi//DD8M3wx/DFva4b/yupo6Gn/7+j8AGRF7gAAAAAAAAAAAAAAAAhni2fA58jjyxi9EOeF/+S4/YAYfCFkZf/8IeeFn/wjZ4WAZPaqP/wzwAgFIFxQBCbi3xYJQFQH8+EFujhPtRNDT/9M/0wDRf/hh+Gb4Y/hi3tcN/5XU0dDT/9/R+ADIi9wAAAAAAAAAAAAAAAAQzxbPgc+Rx5YxeiHPC//JcfsAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPklb4sEohzwv/yXH7ADD4QsjL//hDzws/+EbPCwDJ7VR/FgAE+GcActxwItDWAjHSADDcIccAkvI74CHXDR+S8jzhUxGS8jvhwQQighD////9vLGS8jzgAfAB+EdukvI83g==", valueSigned2.Message)

		// # Create run unsigned message
		address := "0:05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309"
		callSet = domain.CallSet{FunctionName: "returnValue", Input: json.RawMessage(`{"id": "0"}`), Header: &domain.FunctionHeader{PubKey: keyPair.Public, Time: eventsTime, Expire: eventsExpire}}
		signer.PublicKey = keyPair.Public
		valueUnsigned, err = abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue, CallSet: &callSet, Signer: signer, Address: address})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAgEAeAABpYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIFMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKAQBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", valueUnsigned.Message)
		assert.Equal(t, "i4Hs3PB12QA9UBFbOIpkG3JerHHqjm4LgvF4MA7TDsY=", valueUnsigned.DataToSign)

		// # Create detached signature
		valueSignature, err = cryptoUC.Sign(domain.ParamsOfSign{Unsigned: valueUnsigned.DataToSign, Keys: keyPair})
		assert.Equal(t, nil, err)
		assert.Equal(t, "5bbfb7f184f2cb5f019400b9cd497eeaa41f3d5885619e9f7d4fab8dd695f4b3a02159a1422996c1dd7d1be67898bc79c6adba6c65a18101ac5f0a2a2bb8910b", valueSignature.Signature)

		// # Attach signature
		valueSigned, err = abiUC.AttachSignature(domain.ParamsOfAttachSignature{Abi: abiValue, PublicKey: keyPair.Public, Message: valueUnsigned.Message, Signature: valueSignature.Signature})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", valueSigned.Message)

		// # Create initially signed message
		valueSigned3, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue, CallSet: &callSet, Signer: signerKeys, Address: address})
		assert.Equal(t, nil, err)
		assert.Equal(t, "te6ccgEBAwEAvAABRYgAC31qq9KF9Oifst6LU9U6FQSQQRlCSEMo+A3LN5MvphIMAQHhrd/b+MJ5Za+AygBc5qS/dVIPnqxCsM9PvqfVxutK+lnQEKzQoRTLYO6+jfM8TF4841bdNjLQwIDWL4UVFdxIhdMfECP8d3ruNZAXul5xxahT91swIEkEHph08JVlwmUmQAAAXRnJcuDX1XMZBW+LBKACAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", valueSigned3.Message)
	})

	t.Run("TestEncodeAccount", func(t *testing.T) {
		// # Encode account from encoded deploy message
		encodedDeployMessage := "te6ccgECFwEAA2gAAqeIAAt9aqvShfTon7Lei1PVOhUEkEEZQkhDKPgNyzeTL6YSEZTHxAj/Hd67jWQF7peccWoU/dbMCBJBB6YdPCVZcJlJkAAAF0ZyXLg19VzGRotV8/gGAQEBwAICA88gBQMBAd4EAAPQIABB2mPiBH+O713GsgL3S844tQp+62YECSCD0w6eEqy4TKTMAib/APSkICLAAZL0oOGK7VNYMPShCQcBCvSkIPShCAAAAgEgDAoByP9/Ie1E0CDXScIBjhDT/9M/0wDRf/hh+Gb4Y/hijhj0BXABgED0DvK91wv/+GJw+GNw+GZ/+GHi0wABjh2BAgDXGCD5AQHTAAGU0/8DAZMC+ELiIPhl+RDyqJXTAAHyeuLTPwELAGqOHvhDIbkgnzAg+COBA+iogggbd0Cgud6S+GPggDTyNNjTHwH4I7zyudMfAfAB+EdukvI83gIBIBINAgEgDw4AvbqLVfP/hBbo417UTQINdJwgGOENP/0z/TANF/+GH4Zvhj+GKOGPQFcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLe+Ebyc3H4ZtH4APhCyMv/+EPPCz/4Rs8LAMntVH/4Z4AgEgERAA5biABrW/CC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb30gyupo6H0gb+j8IpA3SRg4b3whXXlwMnwAZGT9ghBkZ8KEZ0aCBAfQAAAAAAAAAAAAAAAAACBni2TAgEB9gBh8IWRl//wh54Wf/CNnhYBk9qo//DPAAxbmTwqLfCC3Rwn2omhp/+mf6YBov/ww/DN8Mfwxb2uG/8rqaOhp/+/o/ABkRe4AAAAAAAAAAAAAAAAIZ4tnwOfI48sYvRDnhf/kuP2AGHwhZGX//CHnhZ/8I2eFgGT2qj/8M8AIBSBYTAQm4t8WCUBQB/PhBbo4T7UTQ0//TP9MA0X/4Yfhm+GP4Yt7XDf+V1NHQ0//f0fgAyIvcAAAAAAAAAAAAAAAAEM8Wz4HPkceWMXohzwv/yXH7AMiL3AAAAAAAAAAAAAAAABDPFs+Bz5JW+LBKIc8L/8lx+wAw+ELIy//4Q88LP/hGzwsAye1UfxUABPhnAHLccCLQ1gIx0gAw3CHHAJLyO+Ah1w0fkvI84VMRkvI74cEEIoIQ/////byxkvI84AHwAfhHbpLyPN4="
		messageSource := domain.NewMessageSourceEncoded()
		messageSource.Abi = &abiValue
		messageSource.Message = encodedDeployMessage
		stateInitSource := domain.NewStateInitSourceMessageEncoded()
		stateInitSource.Source = messageSource

		valueEncoded, err := abiUC.EncodeAccount(domain.ParamsOfEncodeAccount{StateInit: stateInitSource})
		assert.Equal(t, nil, err)
		assert.Equal(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", valueEncoded.ID)

		// # Encode account from encoding params
		deploySet := domain.DeploySet(domain.DeploySet{Tvc: eventsTvc})
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{Expire: eventsExpire, Time: eventsTime, PubKey: keyPair.Public}}
		signerKeys := domain.NewSignerKeys()
		signerKeys.Keys = keyPair
		messageSourceEnconing := domain.NewMessageSourceEncodingParams()
		messageSourceEnconing.Abi = abiValue
		messageSourceEnconing.Signer = signerKeys
		messageSourceEnconing.DeploySet = &deploySet
		messageSourceEnconing.CallSet = &callSet
		stateInitSource.Source = messageSourceEnconing

		valueEncoded, err = abiUC.EncodeAccount(domain.ParamsOfEncodeAccount{StateInit: stateInitSource})
		assert.Equal(t, nil, err)
		assert.Equal(t, "05beb555e942fa744fd96f45a9ea9d0a8248208ca12421947c06e59bc997d309", valueEncoded.ID)
	})
}
