package boc

import (
	"github.com/markgenuine/ton-client-go/domain"
)

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
func (b *boc) ParseMessage(pOP domain.ParamsOfParse) (int, error) {
	return b.client.Request("boc.parse_message", pOP)
}

// ParseTransaction method boc.parse_transaction
func (b *boc) ParseTransaction(pOP domain.ParamsOfParse) (int, error) {
	return b.client.Request("boc.parse_transaction", pOP)
}

// ParseAccount method boc.parse_account
func (b *boc) ParseAccount(pOP domain.ParamsOfParse) (int, error) {
	return b.client.Request("boc.parse_account", pOP)
}

// ParseBlock method boc.parse_block
func (b *boc) ParseBlock(pOP domain.ParamsOfParse) (int, error) {
	return b.client.Request("boc.parse_block", pOP)
}

// Parse method boc.parse_shardstate
func (b *boc) ParseShardstate(pOPS domain.ParamsOfParseShardstate) (int, error) {
	return b.client.Request("boc.parse_shardstate", pOPS)
}

// GetBlockhainConfig method boc.get_blockchain_config
func (b *boc) GetBlockhainConfig(pOGBC domain.ParamsOfGetBlockchainConfig) (int, error) {
	return b.client.Request("boc.get_blockchain_config", pOGBC)
}

// Parse method boc.get_boc_hash
func (b *boc) GetBocHash(pOGBH domain.ParamsOfGetBocHash) (int, error) {
	return b.client.Request("boc.get_boc_hash", pOGBH)
}
