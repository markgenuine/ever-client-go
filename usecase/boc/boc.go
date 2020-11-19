package boc

import "github.com/move-ton/ton-client-go/domain"

// import (
// 	"github.com/move-ton/ton-client-go/domain"
// )

type boc struct {
	config domain.Config
	client domain.ClientGateway
}

// NewBoc ...
func NewBoc(
	config domain.Config,
	client domain.ClientGateway,
) domain.BocUseCase {
	return &boc{
		config: config,
		client: client,
	}
}

// ParseMessage method boc.parse_message
func (b *boc) ParseMessage(pOP domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_message", pOP, result)
	return result, err
}

// ParseTransaction method boc.parse_transaction
func (b *boc) ParseTransaction(pOP domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_transaction", pOP, result)
	return result, err
}

// ParseAccount method boc.parse_account
func (b *boc) ParseAccount(pOP domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_account", pOP, result)
	return result, err
}

// ParseBlock method boc.parse_block
func (b *boc) ParseBlock(pOP domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_block", pOP, result)
	return result, err
}

// Parse method boc.parse_shardstate
func (b *boc) ParseShardstate(pOPS domain.ParamsOfParseShardstate) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_shardstate", pOPS, result)
	return result, err
}

// GetBlockhainConfig method boc.get_blockchain_config
func (b *boc) GetBlockhainConfig(pOGBC domain.ParamsOfGetBlockchainConfig) (*domain.ResultOfGetBlockchainConfig, error) {
	result := new(domain.ResultOfGetBlockchainConfig)
	err := b.client.GetResult("boc.get_blockchain_config", pOGBC, result)
	return result, err
}

// Parse method boc.get_boc_hash
func (b *boc) GetBocHash(pOGBH domain.ParamsOfGetBocHash) (*domain.ResultOfGetBocHash, error) {
	result := new(domain.ResultOfGetBocHash)
	err := b.client.GetResult("boc.get_boc_hash", pOGBH, result)
	return result, err
}
