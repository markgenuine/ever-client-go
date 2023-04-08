package net

import (
	"encoding/json"
	"github.com/move-ton/ever-client-go/domain"
)

type net struct {
	config domain.ClientConfig
	client domain.ClientGateway
}

func NewNet(
	config domain.ClientConfig,
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

// QueryCollection - Queries collection data.
func (n *net) QueryCollection(pOQC *domain.ParamsOfQueryCollection) (*domain.ResultOfQueryCollection, error) {
	result := new(domain.ResultOfQueryCollection)
	err := n.client.GetResult("net.query_collection", pOQC, result)
	return result, err
}

// AggregateCollection - Aggregates collection data.
func (n *net) AggregateCollection(pOAC *domain.ParamsOfAggregateCollection) (*domain.ResultOfAggregateCollection, error) {
	result := new(domain.ResultOfAggregateCollection)
	err := n.client.GetResult("net.aggregate_collection", pOAC, result)
	return result, err
}

// WaitForCollection - Returns an object that fulfills the conditions or waits for its appearance.
func (n *net) WaitForCollection(pOWFC *domain.ParamsOfWaitForCollection) (*domain.ResultOfWaitForCollection, error) {
	result := new(domain.ResultOfWaitForCollection)
	err := n.client.GetResult("net.wait_for_collection", pOWFC, result)
	return result, err
}

// Unsubscribe - Cancels a subscription.
func (n *net) Unsubscribe(rOSC *domain.ResultOfSubscribeCollection) error {
	_, err := n.client.GetResponse("net.unsubscribe", rOSC)
	return err
}

// SubscribeCollection - Creates a subscription.
// Creates a collection subscription
// Triggers for each insert/update of data that satisfies the filter conditions. The projection fields are limited to
// result fields.
// The subscription is a persistent communication channel between client and Free TON Network. All changes in the blockchain
// will be reflected in realtime. Changes means inserts and updates of the blockchain entities.
func (n *net) SubscribeCollection(pOSC *domain.ParamsOfSubscribeCollection) (<-chan json.RawMessage, *domain.ResultOfSubscribeCollection, error) {
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
	chanResult := make(chan json.RawMessage, 1)
	go func() {
		var body struct {
			Result json.RawMessage `json:"result"`
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

// Subscribe - Creates a subscription.
// The subscription is a persistent communication channel between client and Everscale Network.
func (n *net) Subscribe(pOS *domain.ParamsOfSubscribe) (<-chan json.RawMessage, *domain.ResultOfSubscribeCollection, error) {
	result := new(domain.ResultOfSubscribeCollection)
	responses, err := n.client.Request("net.subscribe", pOS)
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
	chanResult := make(chan json.RawMessage, 1)
	go func() {
		var body struct {
			Result json.RawMessage `json:"result"`
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

// GetEndpoints - Requests the list of alternative endpoints from server.
func (n *net) GetEndpoints() (*domain.ResultOfGetEndpoints, error) {
	result := new(domain.ResultOfGetEndpoints)
	err := n.client.GetResult("net.get_endpoints", nil, result)
	return result, err
}

// QueryCounterparties - Allows to query and paginate through the list of accounts that the specified account
// has interacted with, sorted by the time of the last internal message between accounts
// Attention this query retrieves data from 'Counterparties' service which is not supported
// in the opensource version of DApp Server (and will not be supported) as well as in EVER OS SE
// (will be supported in SE in future), but is always accessible via EVER OS Devnet/Mainnet Clouds
func (n *net) QueryCounterparties(pOQC *domain.ParamsOfQueryCounterparties) (*domain.ResultOfQueryCollection, error) {
	result := new(domain.ResultOfQueryCollection)
	err := n.client.GetResult("net.query_counterparties", pOQC, result)
	return result, err
}

// QueryTransactionTree - Returns transactions tree for specific message.
// Performs recursive retrieval of the transactions tree produced by the specific message: in_msg -> dst_transaction ->
// out_messages -> dst_transaction -> ...
// All retrieved messages and transactions will be included into result.messages and result.transactions respectively.
// The retrieval process will stop when the retrieved transaction count is more than 50.
// It is guaranteed that each message in result.messages has the corresponding transaction in the result.transactions.
// But there are no guaranties that all messages from transactions out_msgs are presented in result.messages. So the
// application have to continue retrieval
// for missing messages if it requires.
func (n *net) QueryTransactionTree(pOQTT *domain.ParamsOfQueryTransactionTree) (*domain.ResultOfQueryTransactionTree, error) {
	result := new(domain.ResultOfQueryTransactionTree)
	err := n.client.GetResult("net.query_transaction_tree", pOQTT, result)
	return result, err
}

// CreateBlockIterator - Creates block iterator.
// Block iterator uses robust iteration methods that guaranties that every block in the specified range isn't missed or
// iterated twice.
func (n *net) CreateBlockIterator(iterator *domain.ParamsOfCreateBlockIterator) (*domain.RegisteredIterator, error) {
	result := new(domain.RegisteredIterator)
	err := n.client.GetResult("net.create_block_iterator", iterator, result)
	return result, err
}

// ResumeBlockIterator - Resumes block iterator.
// The iterator stays exactly at the same position where the resume_state was catched.
// Application should call the remove_iterator when iterator is no longer required.
func (n *net) ResumeBlockIterator(iterator *domain.ParamsOfResumeBlockIterator) (*domain.RegisteredIterator, error) {
	result := new(domain.RegisteredIterator)
	err := n.client.GetResult("net.resume_block_iterator", iterator, result)
	return result, err
}

// CreateTransactionIterator - Creates transaction iterator.
// Transaction iterator uses robust iteration methods that guaranty that every transaction in the specified range isn't
// missed or iterated twice.
func (n *net) CreateTransactionIterator(iterator *domain.ParamsOfCreateTransactionIterator) (*domain.RegisteredIterator, error) {
	result := new(domain.RegisteredIterator)
	err := n.client.GetResult("net.create_transaction_iterator", iterator, result)
	return result, err
}

// ResumeTransactionIterator - Resumes transaction iterator.
// The iterator stays exactly at the same position where the resume_state was caught. Note that resume_state doesn't store
// the account filter. If the application requires to use the same account filter as it was when the iterator was created
// then the application must pass the account filter again in accounts_filter parameter.
// Application should call the remove_iterator when iterator is no longer required.
func (n *net) ResumeTransactionIterator(iterator *domain.ParamsOfResumeTransactionIterator) (*domain.RegisteredIterator, error) {
	result := new(domain.RegisteredIterator)
	err := n.client.GetResult("net.resume_transaction_iterator", iterator, result)
	return result, err
}

// IteratorNext - Returns next available items.
// In addition to available items this function returns the has_more flag indicating that the iterator isn't reach the
// end of the iterated range yet.
// This function can return the empty list of available items but indicates that there are more items is available. This
// situation appears when the iterator doesn't reach iterated range but database doesn't contains available items yet.
// If application requests resume state in return_resume_state parameter then this function returns resume_state that
// can be used later to resume the iteration from the position after returned items.
// The structure of the items returned depends on the iterator used. See the description to the appropriated iterator
// creation function.
func (n *net) IteratorNext(iterator *domain.ParamsOfIteratorNext) (*domain.ResultOfIteratorNext, error) {
	result := new(domain.ResultOfIteratorNext)
	err := n.client.GetResult("net.iterator_next", iterator, result)
	return result, err
}

// RemoveIterator - Removes an iterator
// Frees all resources allocated in library to serve iterator.
// Application always should call the remove_iterator when iterator is no longer required.
func (n *net) RemoveIterator(iterator *domain.RegisteredIterator) error {
	_, err := n.client.GetResponse("net.remove_iterator", iterator)
	return err
}

// GetSignatureID
// Returns signature ID for configured network if it should be used in messages signature
func (n *net) GetSignatureID() (*domain.ResultOfGetSignatureId, error) {
	result := new(domain.ResultOfGetSignatureId)
	err := n.client.GetResult("net.get_signature_id", nil, result)
	return result, err
}
