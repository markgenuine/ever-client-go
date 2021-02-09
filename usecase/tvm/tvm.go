package tvm

import (
	"github.com/move-ton/ton-client-go/domain"
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
func (t *tvm) RunExecutor(pORE *domain.ParamsOfRunExecutor) (*domain.ResultOfRunExecuteMessage, error) {
	result := new(domain.ResultOfRunExecuteMessage)
	err := t.client.GetResult("tvm.run_executor", pORE, result)
	return result, err
}

// RunTvm method tvm.run_tvm
func (t *tvm) RunTvm(pORT *domain.ParamsOfRunTvm) (*domain.ResultOfRunTvm, error) {
	result := new(domain.ResultOfRunTvm)
	err := t.client.GetResult("tvm.run_tvm", pORT, result)
	return result, err
}

// RunGet method tvm.run_get
func (t *tvm) RunGet(pORG *domain.ParamsOfRunGet) (*domain.ResultOfRunGet, error) {
	result := new(domain.ResultOfRunGet)
	err := t.client.GetResult("tvm.run_get", pORG, result)
	return result, err
}
