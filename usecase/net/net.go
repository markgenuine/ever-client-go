package net

import (
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
func (n *net) Unsubscribe(rOSC domain.ResultOfSubscribeCollection) {
	n.client.GetResult("net.unsubscribe", rOSC, nil)
}

// SubscribeCollection method net.subscribe_collection
func (n *net) SubscribeCollection(pOSC domain.ParamsOfSubscribeCollection) (*domain.ResultOfSubscribeCollection, error) {
	result := new(domain.ResultOfSubscribeCollection)
	err := n.client.GetResult("net.subscribe_collection", pOSC, result)
	return result, err
}
