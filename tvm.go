package goton

// RunExecutor method tvm.run_executor
func (client *Client) RunExecutor(pORE ParamsOfRunExecutor) int {
	return client.Request("tvm.run_executor", structToJSON(pORE))
}

// RunTvm method tvm.run_tvm
func (client *Client) RunTvm(pORT ParamsOfRunTvm) int {
	return client.Request("tvm.run_tvm", structToJSON(pORT))
}

// RunGet method tvm.run_get
func (client *Client) RunGet(pORG ParamsOfRunGet) int {
	return client.Request("tvm.run_get", structToJSON(pORG))
}
