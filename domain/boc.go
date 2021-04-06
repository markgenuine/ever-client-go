package domain

import (
	"encoding/json"
)

// BocErrorCode ...
var BocErrorCode map[string]int

const (
	// BocCacheTypePinned ...
	BocCacheTypeTypePinned BocCacheTypeType = "Pinned"

	//BocCacheTypeUnpinned ...
	BocCacheTypeTypeUnpinned BocCacheTypeType = "Unpinned"

	// BuilderOpTypeTypeInteger ...
	BuilderOpTypeTypeInteger BuilderOpTypeType = "Integer"

	// BuilderOpTypeTypeBitString ...
	BuilderOpTypeTypeBitString BuilderOpTypeType = "BitString"

	// BuilderOpTypeTypeCell ...
	BuilderOpTypeTypeCell BuilderOpTypeType = "Cell"

	// BuilderOpTypeTypeCellBoc ...
	BuilderOpTypeTypeCellBoc BuilderOpTypeType = "CellBoc"
)

type (
	// BocCacheTypeType ...
	BocCacheTypeType string

	// BuilderOpTypeType ...
	BuilderOpTypeType string

	// BocCacheType ...
	BocCacheType struct {
		Type BocCacheTypeType `json:"type"`
		Pin  string           `json:"pin,omitempty"`
	}

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

	// ParamsOfBocCacheGet ...
	ParamsOfBocCacheGet struct {
		BocRef string `json:"boc_ref"`
	}

	// ResultOfBocCacheGet ...
	ResultOfBocCacheGet struct {
		Boc string `json:"boc,omitempty"`
	}

	// ParamsOfBocCacheSet ...
	ParamsOfBocCacheSet struct {
		Boc       string        `json:"boc"`
		CacheType *BocCacheType `json:"cache_type"`
	}

	// ResultOfBocCacheSet ...
	ResultOfBocCacheSet struct {
		BocRef string `json:"boc_ref"`
	}

	// ParamsOfBocCacheUnpin ...
	ParamsOfBocCacheUnpin struct {
		Pin    string `json:"pin"`
		BocRef string `json:"boc_ref,omitempty"`
	}

	// BuilderOp - Cell builder operation.
	BuilderOp struct {
		Type           BuilderOpTypeType `json:"type"`
		Size           int               `json:"size,omitempty"`
		Value          json.RawMessage   `json:"value,omitempty"`
		ValueBitString string            `json:"value,omitempty"`
		Builder        []*BuilderOp      `json:"builder,omitempty"`
		Boc            string            `json:"boc,omitempty"`
	}

	// ParamsOfEncodeBoc ...
	ParamsOfEncodeBoc struct {
		Builder  []BuilderOp
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}
	// ResultOfEncodeBoc ...
	ResultOfEncodeBoc struct {
		Boc string
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
		CacheGet(*ParamsOfBocCacheGet) (*ResultOfBocCacheGet, error)
		CacheSet(*ParamsOfBocCacheSet) (*ResultOfBocCacheSet, error)
		CacheUnpin(*ParamsOfBocCacheUnpin) error
		EncodeBoc(*ParamsOfEncodeBoc) (*ResultOfEncodeBoc, error)
	}
)

func init() {
	BocErrorCode = map[string]int{
		"InvalidBoc":            201,
		"SerializationError":    202,
		"InappropriateBlock":    203,
		"MissingSourceBoc":      204,
		"InsufficientCacheSize": 205,
		"BocRefNotFound":        206,
		"InvalidBocRef":         207,
	}
}

// BocCacheTypePinned - First variant constructors boc cache type
func BocCacheTypePinned(pin string) *BocCacheType {
	return &BocCacheType{Type: BocCacheTypeTypePinned, Pin: pin}
}

// BocCacheTypeUnpinned - Second variant constructors boc cache type
func BocCacheTypeUnpinned() *BocCacheType {
	return &BocCacheType{Type: BocCacheTypeTypeUnpinned}
}

// BuilderOpInteger - Variant construction Integer
func BuilderOpInteger(size int, value json.RawMessage) *BuilderOp {
	return &BuilderOp{Type: BuilderOpTypeTypeInteger, Size: size, Value: value}
}

// BuilderOpBitString - Variant construction BitString
func BuilderOpBitString(value string) *BuilderOp {
	return &BuilderOp{Type: BuilderOpTypeTypeBitString, ValueBitString: value}
}

// BuilderOpCell - Variant construction Cell
func BuilderOpCell(builder []*BuilderOp) *BuilderOp {
	return &BuilderOp{Type: BuilderOpTypeTypeCell, Builder: builder}
}

// BuilderOpCellBoc - Variant construction CellBoc
func BuilderOpCellBoc(boc string) *BuilderOp {
	return &BuilderOp{Type: BuilderOpTypeTypeCellBoc, Boc: boc}
}
