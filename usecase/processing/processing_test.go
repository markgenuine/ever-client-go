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
	configConn := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
	clientConn, err := client.NewClientGateway(configConn)
	assert.Equal(t, nil, err)

	procUC := processing{
		config: configConn,
		client: clientConn,
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

	type resultData struct {
		AccountAddr string `json:"account_addr,omitempty"`
		StatusName  string `json:"status_name,omitempty"`
	}

	t.Run("TestProcessMessage", func(t *testing.T) {
		// # Prepare data for deployment message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSigner(domain.SignerKeys{keypair})
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

		_, err = procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:     domain.NewAbiContract(eventsAbiG),
				Signer:  domain.NewSigner(domain.SignerNone{}),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &domain.CallSet{
					FunctionName: "grant",
					Input:        json.RawMessage(`{"dest":"` + encoded.Address + `"}`),
				}}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		// # Deploy account
		result, err := procUC.ProcessMessage(&domain.ParamsOfProcessMessage{
			MessageEncodeParams: &domain.ParamsOfEncodeMessage{
				Abi:       abiValue,
				Signer:    signer,
				DeploySet: &deploySet,
				CallSet:   &callSet}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		resultSt := &resultData{}
		err = json.Unmarshal(result.Transaction, resultSt)
		assert.Equal(t, nil, err)
		assert.Equal(t, encoded.Address, resultSt.AccountAddr)
		assert.Equal(t, `finalized`, resultSt.StatusName)
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
		signer := domain.NewSigner(domain.SignerKeys{keypair})
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
				Signer:  domain.NewSigner(domain.SignerNone{}),
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

		printMessage(events)

		resSt := &resultData{}
		err = json.Unmarshal(generator.Transaction, &resSt)
		assert.Equal(t, nil, err)

		assert.Equal(t, encoded.Address, resSt.AccountAddr)
		assert.Equal(t, "finalized", resSt.StatusName)
		assert.Equal(t, 0, len(generator.OutMessages))
	})

	t.Run("TestWaitForTransaction", func(t *testing.T) {
		// # Create deploy message
		keypair, err := cryptoUC.GenerateRandomSignKeys()
		assert.Equal(t, nil, err)
		signer := domain.NewSigner(domain.SignerKeys{keypair})
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
				Signer:  domain.NewSigner(domain.SignerNone{}),
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
		signer := domain.NewSigner(domain.SignerKeys{keypair})
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
				Signer:  domain.NewSigner(domain.SignerNone{}),
				Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
				CallSet: &callSetN}, SendEvents: false}, nil)
		assert.Equal(t, nil, err)

		events := make(chan *domain.ProcessingEvent, 10)

		// # Send message
		shardBlockID, err := procUC.SendMessage(&domain.ParamsOfSendMessage{Message: encoded.Message, SendEvents: true, Abi: abiValue}, func(event *domain.ProcessingEvent) { events <- event })
		assert.Equal(t, nil, err)
		close(events)

		printMessage(events)

		events = make(chan *domain.ProcessingEvent, 10)
		//  # Wait for transaction
		result, err := procUC.WaitForTransaction(&domain.ParamsOfWaitForTransaction{Message: encoded.Message, ShardBlockID: shardBlockID.ShardBlockID, SendEvents: true, Abi: abiValue}, func(event *domain.ProcessingEvent) { events <- event })
		assert.Equal(t, nil, err)
		close(events)

		printMessage(events)

		assert.Equal(t, 0, len(result.OutMessages))
		assert.Equal(t, 0, len(result.Decoded.OutMessages))
		assert.Equal(t, json.RawMessage("null"), result.Decoded.Output)
	})
}

func printMessage(events chan *domain.ProcessingEvent) {
	if len(events) == 0 {
		return
	}

	for event := range events {
		switch valueType := event.ValueEnumType.(type) {
		case domain.ProcessingEventWillSend:
			fmt.Println("Type: ProcessingEventWillSend; Shard block id: " + valueType.ShardBlockID)
		case domain.ProcessingEventDidSend:
			fmt.Println("Type: ProcessingEventDidSend; Shard block id: " + valueType.ShardBlockID)
		case domain.ProcessingEventSendFailed:
			fmt.Println("Type: ProcessingEventSendFailed; Shard block id: " + valueType.ShardBlockID)
		case domain.ProcessingEventWillFetchNextBlock:
			fmt.Println("Type: ProcessingEventWillFetchNextBlock; Shard block id: " + valueType.ShardBlockID)
		case domain.ProcessingEventFetchNextBlockFailed:
			fmt.Println("Type: ProcessingEventFetchNextBlockFailed; Shard block id: " + valueType.ShardBlockID)
		default:
			continue
		}
	}
}
