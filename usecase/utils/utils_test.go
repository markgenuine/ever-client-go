package utils

import (
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {

	client, err := client.NewClientGateway(domain.NewDefaultConfig(2))
	assert.Equal(t, nil, err)
	defer client.Destroy()

	utilsUC := utils{
		config: domain.NewDefaultConfig(2),
		client: client,
	}

	AccountID := "fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
	Hex := "-1:fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
	HexWorkchain0 := "0:fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
	Base64 := "Uf/8uRo6OBbQ97jCx2EIuKm8Wmt6Vb15+KsQHFLbKSMiYG+9"
	Base64url := "kf_8uRo6OBbQ97jCx2EIuKm8Wmt6Vb15-KsQHFLbKSMiYIny"

	t.Run("TestConvertAddress", func(t *testing.T) {
		valueConv1, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: AccountID, OutputFormat: domain.AddressStringFormatHex()})
		assert.Equal(t, nil, err)
		assert.Equal(t, HexWorkchain0, valueConv1.Address)

		valueConv2, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: valueConv1.Address, OutputFormat: domain.AddressStringFormatAccountId()})
		assert.Equal(t, nil, err)
		assert.Equal(t, AccountID, valueConv2.Address)

		valueConv3, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: Hex, OutputFormat: domain.AddressStringFormatBase64(false,false,false)})
		assert.Equal(t, nil, err)
		assert.Equal(t, Base64, valueConv3.Address)

		valueConv4, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: Hex, OutputFormat: domain.AddressStringFormatBase64(true, true, true)})
		assert.Equal(t, nil, err)
		assert.Equal(t, Base64url, valueConv4.Address)

		valueConv5, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: Base64url, OutputFormat: domain.AddressStringFormatHex()})
		assert.Equal(t, nil, err)
		assert.Equal(t, Hex, valueConv5.Address)

		valueConv6, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{Address: "-1:00", OutputFormat: domain.AddressStringFormatHex()})
		assert.NotEqual(t, nil, err)
		assert.Equal(t, "", valueConv6.Address)
	})
}
