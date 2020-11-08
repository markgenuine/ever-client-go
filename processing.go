package goton

// SendMessage method processing.send_message
func (client *Client) SendMessage(pOSM ParamsOfSendMessage, ResponseHandler int) int { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return client.Request("processing.send_message", structToJSON(pOSM))
}

// WaitForTransaction method processing.wait_for_transaction
func (client *Client) WaitForTransaction(pOWFT ParamsOfWaitForTransaction, ResponseHandler int) int { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return client.Request("processing.wait_for_transaction", structToJSON(pOWFT))
}

// ProcessMessage method processing.process_message
func (client *Client) ProcessMessage(pOPM ParamsOfProcessMessage, ResponseHandler int) int { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return client.Request("processing.process_message", structToJSON(pOPM))
}
