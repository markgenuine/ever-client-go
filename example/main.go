package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/move-ton/ton-client-go/domain"
	"log"

	goton "github.com/move-ton/ton-client-go"
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

func main() {
	ton, err := goton.NewTon(domain.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer ton.Client.Destroy()

	//value, err := ton.Client.Version()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Version bindings is: ", value.Version)

	//ton,err := NewTonNew()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer ton.Client.Destroy()

	keys, err := ton.Crypto.GenerateRandomSignKeys()
	if err != nil {
		log.Fatal(err)
	}

	appSign := &AppSigningBoxTest{Private: keys.Secret, Public: keys.Public}
	handle, err := ton.Crypto.RegisterSigningBox(appSign)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Handle: ", handle.Handle)
	}

	fmt.Println(appSign.GetPublicKey())

	//keyResult, err := ton.Crypto.SigningBoxGetPublicKey(handle)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	if keyResult.PubKey != keys.Public {
	//		log.Fatal("Don't differen public keys!")
	//	} else {
	//		fmt.Println(keyResult)
	//	}
	//}
	//messageToSign := []byte("Test message")
	//_, err = cryptoUC.SigningBoxSign(&domain.ParamsOfSigningBoxSign{SigningBox: handle.Handle, Unsigned: base64.StdEncoding.EncodeToString(messageToSign)})
	//assert.Equal(t, nil, err)
	//})
}