package goton

// ParseMessage method boc.parse_message
func (client *Client) ParseMessage(pOP ParamsOfParse) int {
	return client.Request("boc.parse_message", structToJSON(pOP))
}

// ParseTransaction method boc.parse_transaction
func (client *Client) ParseTransaction(pOP ParamsOfParse) int {
	return client.Request("boc.parse_transaction", structToJSON(pOP))
}

// ParseAccount method boc.parse_account
func (client *Client) ParseAccount(pOP ParamsOfParse) int {
	return client.Request("boc.parse_account", structToJSON(pOP))
}

// ParseBlock method boc.parse_block
func (client *Client) ParseBlock(pOP ParamsOfParse) int {
	return client.Request("boc.parse_block", structToJSON(pOP))
}

// GetBlockhainConfig method boc.get_blockchain_config
func (client *Client) GetBlockhainConfig(pOGBC ParamsOfGetBlockchainConfig) int {
	return client.Request("boc.get_blockchain_config", structToJSON(pOGBC))
}
