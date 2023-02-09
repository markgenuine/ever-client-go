package main

import (
	goever "github.com/move-ton/ever-client-go"
	"github.com/move-ton/ever-client-go/domain"
	"github.com/move-ton/ever-client-go/util"
	"log"
)

package main

import (
"fmt"
"github.com/move-ton/ever-client-go/domain"
"github.com/move-ton/ever-client-go/util"
"log"

goever "github.com/move-ton/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", domain.GetDevNetBaseUrls(), "")
	if err != nil {
		log.Fatal(err)
	}

	defer ever.Client.Destroy()

	HDPATH := "m/44'/396'/0'/0/0"
	params := &domain.ParamsOfMnemonicDeriveSignKeys{
		Phrase:     "action inject penalty envelope rabbit element slim tornado dinner pizza off blood",
		Path:       HDPATH,
		WordCount:  util.IntToPointerInt(12),
		Dictionary: util.IntToPointerInt(1),
	}
	keyPair, err := ever.Crypto.MnemonicDeriveSignKeys(params)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("PublicKey is: ", keyPair.Public)
	log.Print("SecretKey is: ", keyPair.Secret)
}
