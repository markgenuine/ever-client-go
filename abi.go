package goton

// EncodeMessageBody method encode_message_body
func (client *Client) EncodeMessageBody(pOEMB ParamsOfEncodeMessageBody) int {
	return client.Request("abi.encode_message_body", structToJSON(pOEMB))
}

// AttachSignatureToMessageBody method attach_signature_to_message_body
func (client *Client) AttachSignatureToMessageBody(pOASTMB ParamsOfAttachSignatureToMessageBody) int {
	return client.Request("abi.attach_signature_to_message_body", structToJSON(pOASTMB))
}

// EncodeMessage method encode_message
func (client *Client) EncodeMessage(pOEM ParamsOfEncodeMessage) int {
	return client.Request("abi.encode_message", structToJSON(pOEM))
}

// AttachSignature method attach_signature
func (client *Client) AttachSignature(pOAS ParamsOfAttachSignature) int {
	return client.Request("abi.attach_signature", structToJSON(pOAS))
}

// DecodeMessage method decode_message
func (client *Client) DecodeMessage(pODM ParamsOfDecodeMessage) int {
	return client.Request("abi.decode_message", structToJSON(pODM))
}

// DecodeMessageBody method decode_message_body
func (client *Client) DecodeMessageBody(pODMB ParamsOfDecodeMessageBody) int {
	return client.Request("abi.decode_message_body", structToJSON(pODMB))
}

// EncodeAccount method encode_account
func (client *Client) EncodeAccount(pOEA ParamsOfEncodeAccount) int {
	return client.Request("abi.encode_account", structToJSON(pOEA))
}
