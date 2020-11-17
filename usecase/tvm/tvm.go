package tvm

import (
	"github.com/markgenuine/ton-client-go/domain"
)

type tvm struct {
	config domain.Config
	client domain.ClientGateway
}

// NewTvm ...
func NewTvm(
	config domain.Config,
	client domain.ClientGateway,
) domain.TvmUseCase {
	return &tvm{
		config: config,
		client: client,
	}
}

// RunExecutor method tvm.run_executor
func (t *tvm) RunExecutor(pORE domain.ParamsOfRunExecutor) (int, error) {
	return t.client.Request("tvm.run_executor", pORE)
}

// RunTvm method tvm.run_tvm
func (t *tvm) RunTvm(pORT domain.ParamsOfRunTvm) (int, error) {
	return t.client.Request("tvm.run_tvm", pORT)
}

// RunGet method tvm.run_get
func (t *tvm) RunGet(pORG domain.ParamsOfRunGet) (int, error) {
	return t.client.Request("tvm.run_get", pORG)
}
