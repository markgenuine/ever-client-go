package domain

import "encoding/json"

type (
	// ParamsOfParse ...
	ParamsOfParse struct {
		Boc string `json:"boc"`
	}

	// ResultOfParse ...
	ResultOfParse struct {
		Parsed json.RawMessage `json:"parsed"`
	}

	// ParamsOfParseShardstate ...
	ParamsOfParseShardstate struct {
		Boc         string `json:"boc"`
		ID          string `json:"id"`
		WorkchainID int    `json:"workchain_id"`
	}

	// ParamsOfGetBlockchainConfig ...
	ParamsOfGetBlockchainConfig struct {
		BlockBoc string `json:"block_boc"`
	}

	// ResultOfGetBlockchainConfig ...
	ResultOfGetBlockchainConfig struct {
		ConfigBoc string `json:"config_boc"`
	}

	// ParamsOfGetBocHash ...
	ParamsOfGetBocHash struct {
		Boc string `json:"boc"`
	}

	// ResultOfGetBocHash ...
	ResultOfGetBocHash struct {
		Hash string `json:"hash"`
	}

	//BocUseCase ...
	BocUseCase interface {
		ParseMessage(pOP ParamsOfParse) (*ResultOfParse, error)
		ParseTransaction(pOP ParamsOfParse) (*ResultOfParse, error)
		ParseAccount(pOP ParamsOfParse) (*ResultOfParse, error)
		ParseBlock(pOP ParamsOfParse) (*ResultOfParse, error)
		ParseShardstate(pOPS ParamsOfParseShardstate) (*ResultOfParse, error)
		GetBlockhainConfig(pOGBC ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error)
		GetBocHash(pOGBH ParamsOfGetBocHash) (*ResultOfGetBocHash, error)
	}
)
