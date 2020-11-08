package goton

// ConverAddress ...
func (client *Client) ConverAddress(pOCA ParamsOfConvertAddress) int {
	return client.Request("utils.convert_address", structToJSON(pOCA))
}
