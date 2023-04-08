package domain

import (
	"encoding/json"
	"fmt"
)

var BocErrorCode map[string]int

type (
	BocCacheType struct {
		ValueEnumType interface{}
	}

	BocCacheTypePinned struct {
		Pin string `json:"pin"`
	}

	// BocCacheTypeUnpinned - BOC is placed into a common BOC pool with
	// limited size regulated by LRU (least recently used) cache lifecycle.
	BocCacheTypeUnpinned struct{}

	ParamsOfParse struct {
		Boc string `json:"boc"`
	}

	ResultOfParse struct {
		Parsed json.RawMessage `json:"parsed"`
	}

	ParamsOfParseShardstate struct {
		Boc         string `json:"boc"`
		ID          string `json:"id"`
		WorkchainID int    `json:"workchain_id"`
	}

	ParamsOfGetBlockchainConfig struct {
		BlockBoc string `json:"block_boc"`
	}

	ResultOfGetBlockchainConfig struct {
		ConfigBoc string `json:"config_boc"`
	}

	ParamsOfGetBocHash struct {
		Boc string `json:"boc"`
	}

	ResultOfGetBocHash struct {
		Hash string `json:"hash"`
	}

	ParamsOfGetBocDepth struct {
		Boc string `json:"boc"`
	}

	ResultOfGetBocDepth struct {
		Depth int `json:"depth"`
	}

	ParamsOfGetCodeFromTvc struct {
		Tvc string `json:"tvc"`
	}

	ResultOfGetCodeFromTvc struct {
		Code string `json:"code"`
	}

	ParamsOfBocCacheGet struct {
		BocRef string `json:"boc_ref"`
	}

	ResultOfBocCacheGet struct {
		Boc string `json:"boc,omitempty"`
	}

	ParamsOfBocCacheSet struct {
		Boc       string        `json:"boc"`
		CacheType *BocCacheType `json:"cache_type"`
	}

	ResultOfBocCacheSet struct {
		BocRef string `json:"boc_ref"`
	}

	ParamsOfBocCacheUnpin struct {
		Pin    string `json:"pin"`
		BocRef string `json:"boc_ref,omitempty"`
	}

	// BuilderOp - Cell builder operation.
	BuilderOp struct {
		ValueEnumType interface{}
	}

	BuilderOpInteger struct {
		Size  int         `json:"size"`
		Value interface{} `json:"value"`
	}

	BuilderOpBitString struct {
		Value string `json:"value"`
	}

	BuilderOpCell struct {
		Builder []*BuilderOp `json:"builder"`
	}

	BuilderOpCellBoc struct {
		Boc string `json:"boc"`
	}

	BuilderOpAddress struct {
		Address string `json:"address"`
	}

	ParamsOfEncodeBoc struct {
		Builder  []*BuilderOp  `json:"builder"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfEncodeBoc struct {
		Boc string
	}

	ParamsOfGetCodeSalt struct {
		Code     string        `json:"code"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfGetCodeSalt struct {
		Salt string `json:"salt,omitempty"`
	}

	ParamsOfSetCodeSalt struct {
		Code     string        `json:"code"`
		Salt     string        `json:"salt"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfSetCodeSalt struct {
		Code string `json:"code"`
	}

	ParamsOfDecodeTvc struct {
		Tvc      string        `json:"tvc"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfDecodeTvc struct {
		Code            string `json:"code,omitempty"`
		CodeHash        string `json:"code_hash,omitempty"`
		CodeDepth       *int   `json:"code_depth,omitempty"`
		Data            string `json:"data,omitempty"`
		DataHash        string `json:"data_hash,omitempty"`
		DataDepth       *int   `json:"data_depth,omitempty"`
		Library         string `json:"library,omitempty"`
		Tick            *bool  `json:"tick,omitempty"`
		Tock            *bool  `json:"tock,omitempty"`
		SplitDepth      *int   `json:"split_depth,omitempty"`
		CompilerVersion string `json:"compiler_version,omitempty"`
	}

	ParamsOfEncodeTvc struct {
		Code       string        `json:"code,omitempty"`
		Data       string        `json:"data,omitempty"`
		Library    string        `json:"library,omitempty"`
		Tick       *bool         `json:"tick,omitempty"`
		Tock       *bool         `json:"tock,omitempty"`
		SplitDepth *int          `json:"split_depth,omitempty"`
		BocCache   *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfEncodeTvc struct {
		Tvc string `json:"tvc"`
	}

	ParamsOfGetCompilerVersion struct {
		Code string `json:"code"`
	}

	ResultOfGetCompilerVersion struct {
		Version string `json:"version,omitempty"`
	}

	ParamsOfEncodeExternalInMessage struct {
		Src      string        `json:"src,omitempty"`
		Dst      string        `json:"dst"`
		Init     string        `json:"init,omitempty"`
		Body     string        `json:"body"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}

	ResultOfEncodeExternalInMessage struct {
		Message   string `json:"message"`
		MessageID string `json:"message_id"`
	}

	BocUseCase interface {
		ParseMessage(*ParamsOfParse) (*ResultOfParse, error)
		ParseTransaction(*ParamsOfParse) (*ResultOfParse, error)
		ParseAccount(*ParamsOfParse) (*ResultOfParse, error)
		ParseBlock(*ParamsOfParse) (*ResultOfParse, error)
		ParseShardstate(*ParamsOfParseShardstate) (*ResultOfParse, error)
		GetBlockhainConfig(*ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error)
		GetBocHash(*ParamsOfGetBocHash) (*ResultOfGetBocHash, error)
		GetBocDepth(*ParamsOfGetBocDepth) (*ResultOfGetBocDepth, error)
		GetCodeFromTvc(*ParamsOfGetCodeFromTvc) (*ResultOfGetCodeFromTvc, error)
		CacheGet(*ParamsOfBocCacheGet) (*ResultOfBocCacheGet, error)
		CacheSet(*ParamsOfBocCacheSet) (*ResultOfBocCacheSet, error)
		CacheUnpin(*ParamsOfBocCacheUnpin) error
		EncodeBoc(*ParamsOfEncodeBoc) (*ResultOfEncodeBoc, error)
		GetCodeSalt(*ParamsOfGetCodeSalt) (*ResultOfGetCodeSalt, error)
		SetCodeSalt(*ParamsOfSetCodeSalt) (*ResultOfSetCodeSalt, error)
		DecodeTvc(*ParamsOfDecodeTvc) (*ResultOfDecodeTvc, error)
		EncodeTvc(*ParamsOfEncodeTvc) (*ResultOfEncodeTvc, error)
		EncodeExternalInMessage(*ParamsOfEncodeExternalInMessage) (*ResultOfEncodeExternalInMessage, error)
		GetCompilerVersion(version *ParamsOfGetCompilerVersion) (*ResultOfGetCompilerVersion, error)
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

func (bCT *BocCacheType) MarshalJSON() ([]byte, error) {
	switch value := (bCT.ValueEnumType).(type) {
	case BocCacheTypePinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			BocCacheTypePinned
		}{"Pinned", value})
	case BocCacheTypeUnpinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			BocCacheTypeUnpinned
		}{"Unpinned", value})
	default:
		return nil, fmt.Errorf("unsupported type for BocCacheType %v", bCT.ValueEnumType)
	}
}

func (bCT *BocCacheType) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Pinned":
		var valueEnum BocCacheTypePinned
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bCT.ValueEnumType = valueEnum
	case "Unpinned":
		var valueEnum BocCacheTypeUnpinned
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bCT.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for BocCacheType %v", typeD.Type)
	}
	return nil
}

// NewBocCacheType ...
func NewBocCacheType(value interface{}) *BocCacheType {
	return &BocCacheType{ValueEnumType: value}
}

func (bO *BuilderOp) MarshalJSON() ([]byte, error) {
	switch value := (bO.ValueEnumType).(type) {
	case BuilderOpInteger:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpInteger
		}{"Integer", value})
	case BuilderOpBitString:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpBitString
		}{"BitString", value})
	case BuilderOpCell:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpCell
		}{"Cell", value})
	case BuilderOpCellBoc:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpCellBoc
		}{"CellBoc", value})
	case BuilderOpAddress:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpAddress
		}{"Address", value})
	default:
		return nil, fmt.Errorf("unsupported type for BuilderOp %v", bO.ValueEnumType)
	}
}

func (bO *BuilderOp) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Integer":
		var valueEnum BuilderOpInteger
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "BitString":
		var valueEnum BuilderOpBitString
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "Cell":
		var valueEnum BuilderOpCell
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "CellBoc":
		var valueEnum BuilderOpCellBoc
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "Address":
		var valueEnum BuilderOpAddress
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for BuilderOp %v", typeD.Type)
	}
	return nil
}

func NewBuilderOp(value interface{}) *BuilderOp {
	return &BuilderOp{ValueEnumType: value}
}
