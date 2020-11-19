package processing

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/usecase/abi"
	"github.com/move-ton/ton-client-go/usecase/crypto"
	"github.com/stretchr/testify/assert"
)

func TestProcessing(t *testing.T) {
	config := domain.NewDefaultConfig(1)
	config.Network.ServerAddress = "https://tonos.freeton.surf"
	clientMain, err := client.NewClientGateway(config)
	assert.Equal(t, nil, err)

	procUC := processing{
		config: config,
		client: clientMain,
	}
	defer procUC.client.Destroy()

	cryptoUC := crypto.NewCrypto(procUC.config, procUC.client)
	abiUC := abi.NewAbi(procUC.config, procUC.client)

	fileAbi, err := os.Open("../samples/Events.abi.json")
	assert.Equal(t, nil, err)
	byteAbi, err := ioutil.ReadAll(fileAbi)
	assert.Equal(t, nil, err)

	eventsAbi := domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &eventsAbi)
	assert.Equal(t, nil, err)

	abiValue := domain.NewAbiContract()
	abiValue.Value = eventsAbi

	fileTvc, err := os.Open("../samples/Events.tvc")
	assert.Equal(t, nil, err)
	byteTvc, err := ioutil.ReadAll(fileTvc)
	assert.Equal(t, nil, err)
	deploySet := domain.DeploySet{Tvc: base64.StdEncoding.EncodeToString(byteTvc)}

	t.Run("TestProcessMessage", func(t *testing.T) {
		// # Prepare data for deployment message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSignerKeys()
		signer.Keys = *keypair
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keypair.Public}}

		// # Encode deployment message
		encoded, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
		assert.Equal(t, nil, err)

		// # Send grams
		fileAbiG, err := os.Open("../samples/Giver.abi.json")
		assert.Equal(t, nil, err)
		byteAbiG, err := ioutil.ReadAll(fileAbiG)
		assert.Equal(t, nil, err)

		eventsAbiG := domain.AbiContract{}
		err = json.Unmarshal(byteAbiG, &eventsAbiG)
		assert.Equal(t, nil, err)

		giverAbi := domain.NewAbiContract()
		giverAbi.Value = eventsAbiG
		callSetN := domain.CallSet{}
		callSetN.FunctionName = "grant"
		callSetN.Input = json.RawMessage(`{"dest":"` + encoded.Address + `"}`)
		assert.Equal(t, nil, err)

		_, err = procUC.ProcessMessage(domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Type:    "EncodingParams",
				Abi:     giverAbi,
				Signer:  domain.NewSignerNone(),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false})
		assert.Equal(t, nil, err)

		// # Deploy account
		_, err = procUC.ProcessMessage(domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Type:      "EncodingParams",
				Abi:       abiValue,
				Signer:    signer,
				DeploySet: &deploySet,
				CallSet:   &callSet}, SendEvents: false})

	})
	assert.Equal(t, nil, err)
}
