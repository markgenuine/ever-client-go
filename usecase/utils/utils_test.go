package utils

import (
	"encoding/base64"
	"github.com/move-ton/ton-client-go/util"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	configConn := domain.NewDefaultConfig(domain.BaseUrl)
	client, err := client.NewClientGateway(configConn)
	assert.Equal(t, nil, err)
	defer client.Destroy()

	utilsUC := utils{
		config: configConn,
		client: client,
	}

	t.Run("TestUtils", func(t *testing.T) {
		accountID := "fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
		hex := "-1:fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
		hexWorkchain0 := "0:fcb91a3a3816d0f7b8c2c76108b8a9bc5a6b7a55bd79f8ab101c52db29232260"
		base64 := "Uf/8uRo6OBbQ97jCx2EIuKm8Wmt6Vb15+KsQHFLbKSMiYG+9"
		base64url := "kf_8uRo6OBbQ97jCx2EIuKm8Wmt6Vb15-KsQHFLbKSMiYIny"

		converted, err := utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{
			Address:      accountID,
			OutputFormat: domain.NewAddressStringFormat(domain.AddressStringFormatHex{}),
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, converted.Address, hexWorkchain0)

		converted, err = utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{
			Address:      accountID,
			OutputFormat: domain.NewAddressStringFormat(domain.AddressStringFormatAccountID{}),
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, converted.Address, accountID)

		converted, err = utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{
			Address: hex,
			OutputFormat: domain.NewAddressStringFormat(domain.AddressStringFormatBase64{
				Bounce: false,
				Test:   false,
				URL:    false,
			}),
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, converted.Address, base64)

		converted, err = utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{
			Address: base64,
			OutputFormat: domain.NewAddressStringFormat(domain.AddressStringFormatBase64{
				Bounce: true,
				Test:   true,
				URL:    true,
			}),
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, converted.Address, base64url)

		converted, err = utilsUC.ConvertAddress(&domain.ParamsOfConvertAddress{
			Address:      base64url,
			OutputFormat: domain.NewAddressStringFormat(domain.AddressStringFormatHex{}),
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, converted.Address, hex)
	})

	t.Run("TestCalcStorageFee", func(t *testing.T) {
		account := "te6ccgECHQEAA/wAAnfAArtKDoOR5+qId/SCUGSSS9Qc4RD86X6TnTMjmZ4e+7EyOobmQvsHNngAAAg6t/34DgJWKJuuOehjU0ADAQFBlcBqp0PR+QAN1kt1SY8QavS350RCNNfeZ+ommI9hgd/gAgBToB6t2E3E7a7aW2YkvXv2hTmSWVRTvSYmCVdH4HjgZ4Z94AAAAAvsHNwwAib/APSkICLAAZL0oOGK7VNYMPShBgQBCvSkIPShBQAAAgEgCgcBAv8IAf5/Ie1E0CDXScIBn9P/0wD0Bfhqf/hh+Gb4Yo4b9AVt+GpwAYBA9A7yvdcL//hicPhjcPhmf/hh4tMAAY4SgQIA1xgg+QFY+EIg+GX5EPKo3iP4RSBukjBw3vhCuvLgZSHTP9MfNCD4I7zyuSL5ACD4SoEBAPQOIJEx3vLQZvgACQA2IPhKI8jLP1mBAQD0Q/hqXwTTHwHwAfhHbvJ8AgEgEQsCAVgPDAEJuOiY/FANAdb4QW6OEu1E0NP/0wD0Bfhqf/hh+Gb4Yt7RcG1vAvhKgQEA9IaVAdcLP3+TcHBw4pEgjjJfM8gizwv/Ic8LPzExAW8iIaQDWYAg9ENvAjQi+EqBAQD0fJUB1ws/f5NwcHDiAjUzMehfAyHA/w4AmI4uI9DTAfpAMDHIz4cgzo0EAAAAAAAAAAAAAAAAD3RMfijPFiFvIgLLH/QAyXH7AN4wwP+OEvhCyMv/+EbPCwD4SgH0AMntVN5/+GcBCbkWq+fwEAC2+EFujjbtRNAg10nCAZ/T/9MA9AX4an/4Yfhm+GKOG/QFbfhqcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLe+Ebyc3H4ZtH4APhCyMv/+EbPCwD4SgH0AMntVH/4ZwIBIBUSAQm7Fe+TWBMBtvhBbo4S7UTQ0//TAPQF+Gp/+GH4Zvhi3vpA1w1/ldTR0NN/39cMAJXU0dDSAN/RVHEgyM+FgMoAc89AzgH6AoBrz0DJc/sA+EqBAQD0hpUB1ws/f5NwcHDikSAUAISOKCH4I7ubIvhKgQEA9Fsw+GreIvhKgQEA9HyVAdcLP3+TcHBw4gI1MzHoXwb4QsjL//hGzwsA+EoB9ADJ7VR/+GcCASAYFgEJuORhh1AXAL74QW6OEu1E0NP/0wD0Bfhqf/hh+Gb4Yt7U0fhFIG6SMHDe+EK68uBl+AD4QsjL//hGzwsA+EoB9ADJ7VT4DyD7BCDQ7R7tU/ACMPhCyMv/+EbPCwD4SgH0AMntVH/4ZwIC2hsZAQFIGgAs+ELIy//4Rs8LAPhKAfQAye1U+A/yAAEBSBwAWHAi0NYCMdIAMNwhxwDcIdcNH/K8UxHdwQQighD////9vLHyfAHwAfhHbvJ8"
		result, err := utilsUC.CalcStorageFee(&domain.ParamsOfCalcStorageFee{
			Account: account,
			Period:  1000,
		})
		assert.Equal(t, nil, err)
		assert.Equal(t, result.Fee, "330")
	})

	t.Run("TestCompression", func(t *testing.T) {
		uncompressed := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor \
		incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud \
		exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure \
		dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. \
		Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit \
		anim id est laborum.`

		compressed, err := utilsUC.CompressZstd(&domain.ParamsOfCompressZstd{
			Uncompressed: base64.StdEncoding.EncodeToString([]byte(uncompressed)),
			Level:        util.IntToPointerInt(21)})
		assert.Equal(t, nil, err)
		assert.Equal(t, compressed.Compressed, "KLUv/QCA5QgANpc6GpC3mAOgiezaLbKxxNZJRHFGTt9+Xd31hqoTMwAyADIAaElOPuwEAkBB0R1jLDmT8u6REqcW8MVF2Ef6OamiRONwshstoE0Vn4quHKnJdF5jdAMBC/ccs0+/YaqUrN2x8AN5YZII01nTM9GCqZsVNyVVc3ORoIbKxEu1DFGmRtfoDKoUswFYR3VapaQV64g8NG2ylI86JRdfrKix8GQ5BR6ZmA+MjlHtRoleY1ZADViZOHLG3Gi5AQaTySWvuRW+lDPIyaT18cPh+ky0TslhmVpAM/mhUzss8pOjWMM1xhXtqSiNx2MhYEAZEggQaLkbawsKwzLoDR9YZ83lhYylJkYz7F/bhZDLYL7Xu5gKhBGjBVgsXBVHMAM=")

		decompressed, err := utilsUC.DecompressZstd(&domain.ParamsOfDecompressZstd{Compressed: compressed.Compressed})
		assert.Equal(t, nil, err)

		decompressedS, err := base64.StdEncoding.DecodeString(decompressed.Decompressed)
		assert.Equal(t, nil, err)
		assert.Equal(t, string(decompressedS), uncompressed)
	})
}
