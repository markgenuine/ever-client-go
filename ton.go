package goton

import (
	"github.com/move-ton/ton-client-go/domain"
	clientgw "github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/usecase/abi"
	"github.com/move-ton/ton-client-go/usecase/boc"
	"github.com/move-ton/ton-client-go/usecase/crypto"
	"github.com/move-ton/ton-client-go/usecase/debot"
	"github.com/move-ton/ton-client-go/usecase/net"
	"github.com/move-ton/ton-client-go/usecase/processing"
	"github.com/move-ton/ton-client-go/usecase/tvm"
	"github.com/move-ton/ton-client-go/usecase/utils"
)

// Ton ...
type Ton struct {
	Abi        domain.AbiUseCase
	Boc        domain.BocUseCase
	Client     domain.ClientGateway
	Crypto     domain.CryptoUseCase
	Debot      domain.DebotUseCase
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
		Debot:      debot.NewDebot(config, client),
		Net:        net.NewNet(config, client),
		Processing: processing.NewProcessing(config, client),
		Tvm:        tvm.NewTvm(config, client),
		Utils:      utils.NewUtils(config, client),
	}
	return ton, nil
}

// NewTon ...
func NewTon(address string) (*Ton, error) {
	conf := domain.NewDefaultConfig(address)
	return NewTonWithConfig(conf)
}
