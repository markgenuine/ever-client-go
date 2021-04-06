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

// Query - Performs DAppServer GraphQL query.
func (n *net) Query(pOQ *domain.ParamsOfQuery) (*domain.ResultOfQuery, error) {
	result := new(domain.ResultOfQuery)
	err := n.client.GetResult("net.query", pOQ, result)
	return result, err
}

// BatchQuery - Performs multiple queries per single fetch.
func (n *net) BatchQuery(pOBQ *domain.ParamsOfBatchQuery) (*domain.ResultOfBatchQuery, error) {
	result := new(domain.ResultOfBatchQuery)
	err := n.client.GetResult("net.batch_query", pOBQ, result)
	return result, err
}

// QueryCollection - Queries collection data
// Queries data that satisfies the filter conditions, limits the number of returned records and orders them.
// The projection fields are limited to result fields.
func (n *net) QueryCollection(pOQC *domain.ParamsOfQueryCollection) (*domain.ResultOfQueryCollection, error) {
	result := new(domain.ResultOfQueryCollection)
	err := n.client.GetResult("net.query_collection", pOQC, result)
	return result, err
}

// AggregateCollection - Aggregates collection data.
// Aggregates values from the specified fields for records that satisfies the filter conditions,
// aggregate_collection.
func (n *net) AggregateCollection(pOAC *domain.ParamsOfAggregateCollection) (*domain.ResultOfAggregateCollection, error) {
	result := new(domain.ResultOfAggregateCollection)
	err := n.client.GetResult("net.aggregate_collection", pOAC, result)
	return result, err
}

// WaitForCollection - Returns an object that fulfills the conditions or waits for its appearance
// Triggers only once. If object that satisfies the filter conditions already exists - returns it immediately.
// If not - waits for insert/update of data within the specified timeout, and returns it.
//The projection fields are limited to result fields.
func (n *net) WaitForCollection(pOWFC *domain.ParamsOfWaitForCollection) (*domain.ResultOfWaitForCollection, error) {
	result := new(domain.ResultOfWaitForCollection)
	err := n.client.GetResult("net.wait_for_collection", pOWFC, result)
	return result, err
}

// Unsubscribe - Cancels a subscription.
// Cancels a subscription specified by its handle.
func (n *net) Unsubscribe(rOSC *domain.ResultOfSubscribeCollection) error {
	_, err := n.client.GetResponse("net.unsubscribe", rOSC)
	return err
}

// SubscribeCollection - Creates a subscription
// Triggers for each insert/update of data that satisfies the filter conditions.
// The projection fields are limited to result fields.
func (n *net) SubscribeCollection(pOSC *domain.ParamsOfSubscribeCollection) (<-chan interface{}, *domain.ResultOfSubscribeCollection, error) {
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

// Suspend - Suspends network module to stop any network activity.
func (n *net) Suspend() error {
	_, err := n.client.GetResponse("net.suspend", nil)
	return err
}

// Resume - Resumes network module to enable network activity.
func (n *net) Resume() error {
	_, err := n.client.GetResponse("net.resume", nil)
	return err
}

// FindLastShardBlock - Returns ID of the last block in a specified account shard.
func (n *net) FindLastShardBlock(pOFLSB *domain.ParamsOfFindLastShardBlock) (*domain.ResultOfFindLastShardBlock, error) {
	result := new(domain.ResultOfFindLastShardBlock)
	err := n.client.GetResult("net.find_last_shard_block", pOFLSB, result)
	return result, err
}

// FetchEndpoints - Requests the list of alternative endpoints from server.
func (n *net) FetchEndpoints() (*domain.EndpointsSet, error) {
	result := new(domain.EndpointsSet)
	err := n.client.GetResult("net.fetch_endpoints", nil, result)
	return result, err
}

// SetEndpoints - Sets the list of endpoints to use on reinit.
func (n *net) SetEndpoints(eS *domain.EndpointsSet) error {
	_, err := n.client.GetResponse("net.set_endpoints", eS)
	return err
}
