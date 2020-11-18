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
		idConv, err := utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: AccountID, OutputFormat: domain.NewAddressStringFormatHex()})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err := utilsUC.client.GetResp(idConv)
		assert.Equal(t, nil, err)
		assert.Equal(t, HexWorkchain0, valueConv.(domain.ResultOfConvertAddress).Address)

		idConv = 0
		idConv, err = utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: valueConv.(domain.ResultOfConvertAddress).Address, OutputFormat: domain.NewAddressStringFormatAccountID()})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err = utilsUC.client.GetResp(idConv)
		assert.Equal(t, nil, err)
		assert.Equal(t, AccountID, valueConv.(domain.ResultOfConvertAddress).Address)

		idConv = 0
		idConv, err = utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: Hex, OutputFormat: domain.NewAddressStringFormatBase64()})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err = utilsUC.client.GetResp(idConv)
		assert.Equal(t, nil, err)
		assert.Equal(t, Base64, valueConv.(domain.ResultOfConvertAddress).Address)

		idConv = 0
		structAddressFormat := domain.NewAddressStringFormatBase64()
		structAddressFormat.URL = true
		structAddressFormat.Test = true
		structAddressFormat.Bounce = true
		idConv, err = utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: Hex, OutputFormat: structAddressFormat})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err = utilsUC.client.GetResp(idConv)
		assert.Equal(t, nil, err)
		assert.Equal(t, Base64url, valueConv.(domain.ResultOfConvertAddress).Address)

		idConv = 0
		idConv, err = utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: Base64url, OutputFormat: domain.NewAddressStringFormatHex()})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err = utilsUC.client.GetResp(idConv)
		assert.Equal(t, nil, err)
		assert.Equal(t, Hex, valueConv.(domain.ResultOfConvertAddress).Address)

		idConv = 0
		idConv, err = utilsUC.ConverAddress(domain.ParamsOfConvertAddress{Address: "-1:00", OutputFormat: domain.NewAddressStringFormatHex()})
		assert.Equal(t, nil, err)
		assert.NotEqual(t, 0, idConv)
		valueConv, err = utilsUC.client.GetResp(idConv)
		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, valueConv)
	})
}
