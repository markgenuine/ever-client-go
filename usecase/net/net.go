package net

import (
	"encoding/json"

	"github.com/move-ton/ton-client-go/domain"
)

type net struct {
	config domain.Config
	client domain.ClientGateway
}

// NewNet ...
func NewNet(
	config domain.Config,
	client domain.ClientGateway,
) domain.NetUseCase {
	return &net{
		config: config,
		client: client,
	}
}

// QueryCollection method net.query_collection
func (n *net) QueryCollection(pOQC domain.ParamsOfQueryCollection) (*domain.ResultOfQueryCollection, error) {
	result := new(domain.ResultOfQueryCollection)
	err := n.client.GetResult("net.query_collection", pOQC, result)
	return result, err
}

// WaitForCollection net.wait_for_collection
func (n *net) WaitForCollection(pOWFC domain.ParamsOfWaitForCollection) (*domain.ResultOfWaitForCollection, error) {
	result := new(domain.ResultOfWaitForCollection)
	err := n.client.GetResult("net.wait_for_collection", pOWFC, result)
	return result, err
}

// Unsubscribe net.unsubscribe
func (n *net) Unsubscribe(rOSC domain.ResultOfSubscribeCollection) error {
	_, err := n.client.GetResponse("net.unsubscribe", rOSC)
	return err
}

// SubscribeCollection method net.subscribe_collection
func (n *net) SubscribeCollection(pOSC domain.ParamsOfSubscribeCollection) (<-chan interface{}, *domain.ResultOfSubscribeCollection, error) {
	result := new(domain.ResultOfSubscribeCollection)
	responses, err := n.client.Request("net.subscribe_collection", pOSC)
	if err != nil {
		return nil, nil, err
	}

	data := <-responses
	if data.Error != nil {
		return nil, nil, data.Error
	}
	if err := json.Unmarshal(data.Data, result); err != nil {
		return nil, nil, err
	}

	respInBuffer := domain.DynBufferForResponses(responses)
	chanResult := make(chan interface{}, 1)
	go func() {
		var body struct {
			Result interface{} `json:"result"`
		}
		for r := range respInBuffer {
			if err := json.Unmarshal(r.Data, &body); err != nil {
				panic(err)
			}
			chanResult <- body.Result
		}
		close(chanResult)
	}()

	return chanResult, result, nil
}
