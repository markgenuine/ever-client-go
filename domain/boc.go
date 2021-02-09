package domain

import "encoding/json"

// BocErrorCode ...
var BocErrorCode map[string]int

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

	//ParamsOfGetCodeFromTvc ...
	ParamsOfGetCodeFromTvc struct {
		Tvc string `json:"tvc"`
	}

	//ResultOfGetCodeFromTvc ...
	ResultOfGetCodeFromTvc struct {
		Code string `json:"code"`
	}

	//BocUseCase ...
	BocUseCase interface {
		ParseMessage(*ParamsOfParse) (*ResultOfParse, error)
		ParseTransaction(*ParamsOfParse) (*ResultOfParse, error)
		ParseAccount(*ParamsOfParse) (*ResultOfParse, error)
		ParseBlock(*ParamsOfParse) (*ResultOfParse, error)
		ParseShardstate(*ParamsOfParseShardstate) (*ResultOfParse, error)
		GetBlockhainConfig(*ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error)
		GetBocHash(*ParamsOfGetBocHash) (*ResultOfGetBocHash, error)
		GetCodeFromTvc(*ParamsOfGetCodeFromTvc) (*ResultOfGetCodeFromTvc, error)
	}
)

func init() {
	BocErrorCode = map[string]int{
		"InvalidBoc":         201,
		"SerializationError": 202,
		"InappropriateBlock": 203,
		"MissingSourceBoc":   204,
	}
}
