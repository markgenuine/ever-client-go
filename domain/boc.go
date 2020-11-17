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
		ParseMessage(pOP ParamsOfParse) (int, error)
		ParseTransaction(pOP ParamsOfParse) (int, error)
		ParseAccount(pOP ParamsOfParse) (int, error)
		ParseBlock(pOP ParamsOfParse) (int, error)
		ParseShardstate(pOPS ParamsOfParseShardstate) (int, error)
		GetBlockhainConfig(pOGBC ParamsOfGetBlockchainConfig) (int, error)
		GetBocHash(pOGBH ParamsOfGetBocHash) (int, error)
	}
)
