package goever

import (
	"github.com/markgenuine/ever-client-go/domain"
	clientgw "github.com/markgenuine/ever-client-go/gateway/client"
	"github.com/markgenuine/ever-client-go/usecase/abi"
	"github.com/markgenuine/ever-client-go/usecase/boc"
	"github.com/markgenuine/ever-client-go/usecase/crypto"
	"github.com/markgenuine/ever-client-go/usecase/debot"
	"github.com/markgenuine/ever-client-go/usecase/net"
	"github.com/markgenuine/ever-client-go/usecase/processing"
	"github.com/markgenuine/ever-client-go/usecase/proofs"
	"github.com/markgenuine/ever-client-go/usecase/tvm"
	"github.com/markgenuine/ever-client-go/usecase/utils"
)

// Ever Obj ...
type Ever struct {
	Abi        domain.AbiUseCase
	Boc        domain.BocUseCase
	Client     domain.ClientGateway
	Crypto     domain.CryptoUseCase
	Debot      domain.DebotUseCase
	Net        domain.NetUseCase
	Processing domain.ProcessingUseCase
	Proofs     domain.ProofsUseCase
	Tvm        domain.TvmUseCase
	Utils      domain.UtilsUseCase
}

// NewEverWithConfig ...
func NewEverWithConfig(config domain.ClientConfig) (*Ever, error) {
	client, err := clientgw.NewClientGateway(config)
	if err != nil {
		return nil, err
	}
	ever := &Ever{
		Abi:        abi.NewAbi(config, client),
		Boc:        boc.NewBoc(config, client),
		Client:     client,
		Crypto:     crypto.NewCrypto(config, client),
		Debot:      debot.NewDebot(config, client),
		Net:        net.NewNet(config, client),
		Proofs:     proofs.NewProofs(config, client),
		Processing: processing.NewProcessing(config, client),
		Tvm:        tvm.NewTvm(config, client),
		Utils:      utils.NewUtils(config, client),
	}
	return ever, nil
}

// NewEver ...
func NewEver(address string, endPoints []string, accessKey string) (*Ever, error) {
	conf := domain.NewDefaultConfig(address, endPoints, accessKey)
	return NewEverWithConfig(conf)
}
