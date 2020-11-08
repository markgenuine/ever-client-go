package goton

// QueryCollection method net.query_collection
func (client *Client) QueryCollection(pOQC ParamsOfQueryCollection) int {
	return client.Request("net.query_collection", structToJSON(pOQC))
}

// WaitForCollection net.wait_for_collection
func (client *Client) WaitForCollection(pOWFC ParamsOfWaitForCollection) int {
	return client.Request("net.wait_for_collection", structToJSON(pOWFC))
}

// Unsubscribe net.unsubscribe
func (client *Client) Unsubscribe(rOSC ResultOfSubscribeCollection) int {
	return client.Request("net.unsubscribe", structToJSON(rOSC))
}

// SubscribeCollection method net.subscribe_collection
func (client *Client) SubscribeCollection(pOSC ParamsOfSubscribeCollection) int {
	return client.Request("net.subscribe_collection", structToJSON(pOSC))
}
