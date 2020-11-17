package goton

import (
	"github.com/markgenuine/ton-client-go/domain"
	clientgw "github.com/markgenuine/ton-client-go/gateway/client"
	"github.com/markgenuine/ton-client-go/usecase/abi"
	"github.com/markgenuine/ton-client-go/usecase/boc"
	"github.com/markgenuine/ton-client-go/usecase/crypto"
	"github.com/markgenuine/ton-client-go/usecase/net"
	"github.com/markgenuine/ton-client-go/usecase/processing"
	"github.com/markgenuine/ton-client-go/usecase/tvm"
	"github.com/markgenuine/ton-client-go/usecase/utils"
)

// Ton ...
type Ton struct {
	Abi        domain.AbiUseCase
	Boc        domain.BocUseCase
	Client     domain.ClientGateway
	Crypto     domain.CryptoUseCase
	Net        domain.NetUseCase
	Processing domain.ProcessingUseCase
	Tvm        domain.TvmUseCase
	Utils      domain.UtilsUseCase
}

// NewTonWithConfig ...
func NewTonWithConfig(config domain.Config) (*Ton, error) {
	client, err := clientgw.NewClientGateway(config)
	if err != nil {
		return nil, err
	}
	ton := &Ton{
		Abi:        abi.NewAbi(config, client),
		Boc:        boc.NewBoc(config, client),
		Client:     client,
		Crypto:     crypto.NewCrypto(config, client),
		Net:        net.NewNet(config, client),
		Processing: processing.NewProcessing(config, client),
		Tvm:        tvm.NewTvm(config, client),
		Utils:      utils.NewUtils(config, client),
	}
	return ton, nil
}

// NewTon ...
func NewTon(chainID int) (*Ton, error) {
	conf := domain.NewDefaultConfig(chainID)
	return NewTonWithConfig(conf)
}
