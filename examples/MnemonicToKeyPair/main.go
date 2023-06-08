package main

import (
	"github.com/markgenuine/ever-client-go/domain"
	"log"

	goever "github.com/markgenuine/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", domain.GetDevNetBaseUrls(), "")
	if err != nil {
		log.Fatal(err)
	}

	defer ever.Client.Destroy()

	HDPATH := "m/44'/396'/0'/0/0"
	params := &domain.ParamsOfMnemonicDeriveSignKeys{
		Phrase: "action inject penalty envelope rabbit element slim tornado dinner pizza off blood",
		Path:   HDPATH,
	}
	keyPair, err := ever.Crypto.MnemonicDeriveSignKeys(params)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("PublicKey is: ", keyPair.Public)
	log.Print("SecretKey is: ", keyPair.Secret)
}
