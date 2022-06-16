package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/move-ton/ever-client-go/domain"
	"github.com/move-ton/ever-client-go/util"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	goever "github.com/move-ton/ever-client-go"
)

func main() {
	ever, err := goever.NewEver("", domain.GetDevNetBaseUrls())
	if err != nil {
		log.Fatal(err)
	}

	defer ever.Client.Destroy()

	eventsTime := big.NewInt(1599458364291)
	eventsExpire := util.IntToPointerInt(1599458404)
	keyPair, err := ever.Crypto.GenerateRandomSignKeys()
	if err != nil {
		log.Fatalf("Error in generation keys: %s", err)
	}

	fileAbi, err := os.Open("contracts/SafeMultisigWallet.abi.json")
	//fileAbi, err := os.Open("contracts/Events.abi.json")
	if err != nil {
		log.Fatalf("Can't open file %s, error: %s", "..contracts/SafeMultisigWallet.abi.json", err)
	}
	byteAbi, err := ioutil.ReadAll(fileAbi)
	nn := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &nn)
	safeMultisigAbi := domain.NewAbiContract(nn)

	fileTvc, err := os.Open("contracts/SafeMultisigWallet.tvc")
	if err != nil {
		log.Fatalf("Can't open file %s, error: %s", "..contracts/SafeMultisigWallet.tvc", err)
	}

	fileTvcToByte, err := ioutil.ReadAll(fileTvc)
	if err != nil {
		log.Fatalf("Can't converting file to byte: %s", err)
	}

	ownersSlice := []string{"0x1111111111111111111111111111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222222222222222222222222222"}
	ownersJson, _ := json.Marshal(ownersSlice)

	unsignedMessage, err := ever.Abi.EncodeMessage(&domain.ParamsOfEncodeMessage{
		Abi:       safeMultisigAbi,
		Signer:    domain.NewSigner(domain.SignerExternal{PublicKey: keyPair.Public}),
		DeploySet: &domain.DeploySet{Tvc: base64.StdEncoding.EncodeToString(fileTvcToByte)},
		CallSet: &domain.CallSet{
			FunctionName: "constructor",
			Header: &domain.FunctionHeader{
				PubKey: keyPair.Public,
				Time:   eventsTime,
				Expire: eventsExpire,
			},
			Input: json.RawMessage(fmt.Sprintf(`{"owners": %s,"reqConfirms": 2}`, string(ownersJson))),
		},
	})

	if err != nil {
		log.Fatalf("Error in encode message: %s", err)
	}

	fmt.Println("Address for deploy: ", unsignedMessage.Address)
}
