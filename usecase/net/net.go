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
func (n *net) QueryCollection(pOQC domain.ParamsOfQueryCollection) (int, error) {
	return n.client.Request("net.query_collection", pOQC)
}

// WaitForCollection net.wait_for_collection
func (n *net) WaitForCollection(pOWFC domain.ParamsOfWaitForCollection) (int, error) {
	return n.client.Request("net.wait_for_collection", pOWFC)
}

// Unsubscribe net.unsubscribe
func (n *net) Unsubscribe(rOSC domain.ResultOfSubscribeCollection) (int, error) {
	return n.client.Request("net.unsubscribe", rOSC)
}

// SubscribeCollection method net.subscribe_collection
func (n *net) SubscribeCollection(pOSC domain.ParamsOfSubscribeCollection) (int, error) {
	return n.client.Request("net.subscribe_collection", pOSC)
}
