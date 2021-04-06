package domain

const (

	// AddressStringFormatTypeID ...
	AddressStringFormatTypeID AddressStringFormatType = "AccountId"

	// AddressStringFormatTypeHex ...
	AddressStringFormatTypeHex AddressStringFormatType = "Hex"

	// AddressStringFormatTypeBase64 ...
	AddressStringFormatTypeBase64 AddressStringFormatType = "Base64"
)

type (
	// AddressStringFormatType ...
	AddressStringFormatType string

	// AddressStringFormat
	AddressStringFormat struct {
		Type   AddressStringFormatType `json:"type"`
		URL    bool                    `json:"url,omitempty"`
		Test   bool                    `json:"test,omitempty"`
		Bounce bool                    `json:"bounce,omitempty"`
	}

	// ParamsOfConvertAddress ...
	ParamsOfConvertAddress struct {
		Address      string               `json:"address"`
		OutputFormat *AddressStringFormat `json:"output_format"`
	}

	// ResultOfConvertAddress ...
	ResultOfConvertAddress struct {
		Address string `json:"address"`
	}

	// ParamsOfCalcStorageFee ...
	ParamsOfCalcStorageFee struct {
		Account string `json:"account"`
		Period  int    `json:"period"`
	}

	// ResultOfCalcStorageFee ...
	ResultOfCalcStorageFee struct {
		Fee string `json:"fee"`
	}

	//UtilsUseCase ...
	UtilsUseCase interface {
		ConvertAddress(*ParamsOfConvertAddress) (*ResultOfConvertAddress, error)
		CalcStorageFee(pOCA *ParamsOfCalcStorageFee) (*ResultOfCalcStorageFee, error)
	}
)

// AddressStringFormatAccountId ...
func AddressStringFormatAccountId() *AddressStringFormat {
	return &AddressStringFormat{Type: AddressStringFormatTypeID}
}

// AddressStringFormatHex ...
func AddressStringFormatHex() *AddressStringFormat {
	return &AddressStringFormat{Type: AddressStringFormatTypeHex}
}

// AddressStringFormatBase64 ...
func AddressStringFormatBase64(url, test, bounce bool) *AddressStringFormat {
	return &AddressStringFormat{Type: AddressStringFormatTypeBase64, URL: url, Test: test, Bounce: bounce}
}
