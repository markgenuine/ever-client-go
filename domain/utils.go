package domain

import (
	"encoding/json"
	"fmt"
)

const (
	AccountIDAccountAddresType AccountAddressType = "AccountId"
	HexAccountAddresType       AccountAddressType = "Hex"
	Base64AccountAddresType    AccountAddressType = "Base64"
)

type (
	AccountAddressType string

	AddressStringFormatAccountID struct{}

	AddressStringFormatHex struct{}

	AddressStringFormatBase64 struct {
		URL    bool `json:"url"`
		Test   bool `json:"test"`
		Bounce bool `json:"bounce"`
	}

	AddressStringFormat struct {
		ValueEnumType interface{}
	}

	ParamsOfConvertAddress struct {
		Address      string               `json:"address"`
		OutputFormat *AddressStringFormat `json:"output_format"`
	}

	ResultOfConvertAddress struct {
		Address string `json:"address"`
	}

	ParamsOfGetAddressType struct {
		Address string `json:"address"`
	}

	ResultOfGetAddressType struct {
		AddressType AccountAddressType `json:"address_type"`
	}

	ParamsOfCalcStorageFee struct {
		Account string `json:"account"`
		Period  int    `json:"period"`
	}

	ResultOfCalcStorageFee struct {
		Fee string `json:"fee"`
	}

	ParamsOfCompressZstd struct {
		Uncompressed string `json:"uncompressed"`
		Level        *int   `json:"level,omitempty"`
	}

	ResultOfCompressZstd struct {
		Compressed string `json:"compressed"`
	}

	ParamsOfDecompressZstd struct {
		Compressed string `json:"compressed"`
	}

	ResultOfDecompressZstd struct {
		Decompressed string `json:"decompressed"`
	}

	UtilsUseCase interface {
		ConvertAddress(*ParamsOfConvertAddress) (*ResultOfConvertAddress, error)
		GetAddressType(*ParamsOfGetAddressType) (*ResultOfGetAddressType, error)
		CalcStorageFee(pOCA *ParamsOfCalcStorageFee) (*ResultOfCalcStorageFee, error)
		CompressZstd(pOCA *ParamsOfCompressZstd) (*ResultOfCompressZstd, error)
		DecompressZstd(pOCA *ParamsOfDecompressZstd) (*ResultOfDecompressZstd, error)
	}
)

func (a *AddressStringFormat) MarshalJSON() ([]byte, error) {
	switch value := (a.ValueEnumType).(type) {
	case AddressStringFormatAccountID:
		return json.Marshal(struct {
			Type string `json:"type"`
			AddressStringFormatAccountID
		}{"AccountId", value})
	case AddressStringFormatHex:
		return json.Marshal(struct {
			Type string `json:"type"`
			AddressStringFormatHex
		}{"Hex", value})
	case AddressStringFormatBase64:
		return json.Marshal(struct {
			Type string `json:"type"`
			AddressStringFormatBase64
		}{"Base64", value})
	default:
		return nil, fmt.Errorf("unsupported type for AddressStringFormat %v", a.ValueEnumType)
	}
}

func (a *AddressStringFormat) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "AccountId":
		var valueEnum AddressStringFormatAccountID
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		a.ValueEnumType = valueEnum
	case "Hex":
		var valueEnum AddressStringFormatHex
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		a.ValueEnumType = valueEnum
	case "Base64":
		var valueEnum AddressStringFormatBase64
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		a.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for AddressStringFormat %v", typeD.Type)
	}
	return nil
}

func NewAddressStringFormat(value interface{}) *AddressStringFormat {
	return &AddressStringFormat{ValueEnumType: value}
}
