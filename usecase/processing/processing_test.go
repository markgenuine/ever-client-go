package processing

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

	eventsAbi := &domain.AbiContract{}
	err = json.Unmarshal(byteAbi, &eventsAbi)
	assert.Equal(t, nil, err)

	abiValue := domain.NewAbiContract(eventsAbi)

	fileTvc, err := os.Open("../samples/Events.tvc")
	assert.Equal(t, nil, err)
	byteTvc, err := ioutil.ReadAll(fileTvc)
	assert.Equal(t, nil, err)
	deploySet := domain.DeploySet{Tvc: base64.StdEncoding.EncodeToString(byteTvc)}

	t.Run("TestProcessMessage", func(t *testing.T) {
		// # Prepare data for deployment message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSignerKeys(keypair)
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keypair.Public}}

		// # Encode deployment message
		encoded, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: abiValue, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
		assert.Equal(t, nil, err)

		// # Send grams
		fileAbiG, err := os.Open("../samples/Giver.abi.json")
		assert.Equal(t, nil, err)
		byteAbiG, err := ioutil.ReadAll(fileAbiG)
		assert.Equal(t, nil, err)

		eventsAbiG := &domain.AbiContract{}
		err = json.Unmarshal(byteAbiG, &eventsAbiG)
		assert.Equal(t, nil, err)

		giverAbi := domain.NewAbiContract(eventsAbiG)
		callSetN := domain.CallSet{}
		callSetN.FunctionName = "grant"
		callSetN.Input = json.RawMessage(`{"dest":"` + encoded.Address + `"}`)
		assert.Equal(t, nil, err)

		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     giverAbi,
				Signer:  domain.NewSignerNone(),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		// # Deploy account
		result, err := procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:       abiValue,
				Signer:    signer,
				DeploySet: &deploySet,
				CallSet:   &callSet}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)
		var object map[string]json.RawMessage
		err = json.Unmarshal(result.Transaction, &object)
		assert.Equal(t, nil, err)
		assert.Equal(t, `"`+encoded.Address+`"`, string(object["account_addr"]))
		assert.Equal(t, `"finalized"`, string(object["status_name"]))
		assert.Equal(t, 0, len(result.OutMessages))

		// # Contract execution error
		callSetErr := domain.CallSet{FunctionName: "returnValue", Input: json.RawMessage(`{"id": -1}`)}
		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     abiValue,
				Signer:  signer,
				Address: encoded.Address,
				CallSet: &callSetErr,
			}, SendEvents: false}, nil)
		assert.NotEqual(t, nil, err)
	})

	t.Run("TestProcessMessageWithEvents", func(t *testing.T) {
		// # Prepare data for deployment message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSignerKeys(keypair)
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keypair.Public}}

		// # Encode deployment message
		encoded, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: abiValue, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
		assert.Equal(t, nil, err)

		// # Send grams
		fileAbiG, err := os.Open("../samples/Giver.abi.json")
		assert.Equal(t, nil, err)
		byteAbiG, err := ioutil.ReadAll(fileAbiG)
		assert.Equal(t, nil, err)

		eventsAbiG := &domain.AbiContract{}
		err = json.Unmarshal(byteAbiG, &eventsAbiG)
		assert.Equal(t, nil, err)

		giverAbi := domain.NewAbiContract(eventsAbiG)
		callSetN := domain.CallSet{}
		callSetN.FunctionName = "grant"
		callSetN.Input = json.RawMessage(`{"dest":"` + encoded.Address + `"}`)
		assert.Equal(t, nil, err)

		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     giverAbi,
				Signer:  domain.NewSignerNone(),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		events := make(chan *domain.ProcessingEvent, 10)

		// # Deploy account
		generator, err := procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:       abiValue,
				Signer:    signer,
				DeploySet: &deploySet,
				CallSet:   &callSet}, SendEvents: true}, func(event *domain.ProcessingEvent) { events <- event })
		assert.Equal(t, nil, err)
		close(events)

		if len(events) > 0 {
			for event := range events {
				fmt.Println("Type: " + string(event.Type) + "; Shard block id: " + event.ShardBlockID)
			}
		}

		var (
			objmap                  map[string]json.RawMessage
			accountAddr, statusName string
		)
		err = json.Unmarshal(generator.Transaction, &objmap)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(objmap["account_addr"], &accountAddr)
		assert.Equal(t, nil, err)

		err = json.Unmarshal(objmap["status_name"], &statusName)
		assert.Equal(t, nil, err)

		assert.Equal(t, encoded.Address, accountAddr)
		assert.Equal(t, "finalized", statusName)
		assert.Equal(t, 0, len(generator.OutMessages))
	})

	t.Run("TestWaitForTransaction", func(t *testing.T) {
		// # Create deploy message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSignerKeys(keypair)
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keypair.Public}}

		// # Encode deployment message
		encoded, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: abiValue, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
		assert.Equal(t, nil, err)

		// # Send grams
		fileAbiG, err := os.Open("../samples/Giver.abi.json")
		assert.Equal(t, nil, err)
		byteAbiG, err := ioutil.ReadAll(fileAbiG)
		assert.Equal(t, nil, err)

		eventsAbiG := &domain.AbiContract{}
		err = json.Unmarshal(byteAbiG, &eventsAbiG)
		assert.Equal(t, nil, err)

		giverAbi := domain.NewAbiContract(eventsAbiG)
		callSetN := domain.CallSet{}
		callSetN.FunctionName = "grant"
		callSetN.Input = json.RawMessage(`{"dest":"` + encoded.Address + `"}`)
		assert.Equal(t, nil, err)
		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     giverAbi,
				Signer:  domain.NewSignerNone(),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		// # Send message
		shardBlockID, err := procUC.SendMessage(&domain.ParamsOfSendMessage{Message: encoded.Message, SendEvents: false, Abi: abiValue}, nil)
		assert.Equal(t, nil, err)

		//  # Wait for transaction
		result, err := procUC.WaitForTransaction(&domain.ParamsOfWaitForTransaction{Message: encoded.Message, ShardBlockID: shardBlockID.ShardBlockID, SendEvents: false, Abi: abiValue}, nil)
		assert.Equal(t, nil, err)
		assert.Equal(t, 0, len(result.OutMessages))
		assert.Equal(t, 0, len(result.Decoded.OutMessages))
		assert.Equal(t, json.RawMessage("null"), result.Decoded.Output)
	})

	t.Run("TestWaitForTransactionWithEvents", func(t *testing.T) {
		// # Create deploy message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSignerKeys(keypair)
		callSet := domain.CallSet{FunctionName: "constructor", Header: &domain.FunctionHeader{PubKey: keypair.Public}}

		// # Encode deployment message
		encoded, err := abiUC.EncodeMessage(&domain.ParamsOfEncodeMessage{Abi: abiValue, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
		assert.Equal(t, nil, err)

		// # Send grams
		fileAbiG, err := os.Open("../samples/Giver.abi.json")
		assert.Equal(t, nil, err)
		byteAbiG, err := ioutil.ReadAll(fileAbiG)
		assert.Equal(t, nil, err)

		eventsAbiG := &domain.AbiContract{}
		err = json.Unmarshal(byteAbiG, &eventsAbiG)
		assert.Equal(t, nil, err)

		giverAbi := domain.NewAbiContract(eventsAbiG)
		callSetN := domain.CallSet{}
		callSetN.FunctionName = "grant"
		callSetN.Input = json.RawMessage(`{"dest":"` + encoded.Address + `"}`)
		assert.Equal(t, nil, err)
		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     giverAbi,
				Signer:  domain.NewSignerNone(),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		events := make(chan *domain.ProcessingEvent, 10)

		// # Send message
		shardBlockID, err := procUC.SendMessage(&domain.ParamsOfSendMessage{Message: encoded.Message, SendEvents: true, Abi: abiValue}, func(event *domain.ProcessingEvent) { events <- event })
		assert.Equal(t, nil, err)
		close(events)

		if len(events) > 0 {
			for event := range events {
				fmt.Println("Type: " + string(event.Type) + "; Shard block id: " + event.ShardBlockID)
			}
		}

		events = make(chan *domain.ProcessingEvent, 10)
		//  # Wait for transaction
		result, err := procUC.WaitForTransaction(&domain.ParamsOfWaitForTransaction{Message: encoded.Message, ShardBlockID: shardBlockID.ShardBlockID, SendEvents: true, Abi: abiValue}, func(event *domain.ProcessingEvent) { events <- event })
		assert.Equal(t, nil, err)
		close(events)

		if len(events) > 0 {
			for event := range events {
				fmt.Println("Type: " + string(event.Type) + "; Shard block id: " + event.ShardBlockID)
			}
		}

		assert.Equal(t, 0, len(result.OutMessages))
		assert.Equal(t, 0, len(result.Decoded.OutMessages))
		assert.Equal(t, json.RawMessage("null"), result.Decoded.Output)
	})
}
