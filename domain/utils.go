package domain

const (

	// ASFAccountID ...
	ASFAccountID AddressStringFormatType = "AccountId"

	// ASFHex ...
	ASFHex AddressStringFormatType = "Hex"

	// ASFBase64 ...
	ASFBase64 AddressStringFormatType = "Base64"
)

type (

	// AddressStringFormatType ...
	AddressStringFormatType string

	// AddressStringFormatAccountID ...
	AddressStringFormatAccountID struct {
		Type AddressStringFormatType `json:"type"`
	}

	// AddressStringFormatHex ...
	AddressStringFormatHex struct {
		Type AddressStringFormatType `json:"type"`
	}

	// AddressStringFormatBase64 ...
	AddressStringFormatBase64 struct {
		Type   AddressStringFormatType `json:"type"`
		URL    bool                    `json:"url"`
		Test   bool                    `json:"test"`
		Bounce bool                    `json:"bounce"`
	}

	// ParamsOfConvertAddress ...
	ParamsOfConvertAddress struct {
		Address      string      `json:"address"`
		OutputFormat interface{} `json:"output_format"`
	}

	// ResultOfConvertAddress ...
	ResultOfConvertAddress struct {
		Address string `json:"address"`
	}

	//UtilsUseCase ...
	UtilsUseCase interface {
		ConverAddress(pOCA ParamsOfConvertAddress) (int, error)
	}
)

// NewAddressStringFormatAccountID ...
func NewAddressStringFormatAccountID() AddressStringFormatAccountID {
	return AddressStringFormatAccountID{Type: ASFAccountID}
}

// NewAddressStringFormatHex ...
func NewAddressStringFormatHex() AddressStringFormatHex {
	return AddressStringFormatHex{Type: ASFHex}
}

// NewAddressStringFormatBase64 ...
func NewAddressStringFormatBase64() AddressStringFormatBase64 {
	return AddressStringFormatBase64{Type: ASFBase64}
}
