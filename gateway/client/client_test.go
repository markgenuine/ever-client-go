package client

import (
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	client, err := NewClientGateway(domain.NewDefaultConfig(2))
	assert.Equal(t, nil, err)
	defer client.Destroy()

	t.Run("TestClient", func(t *testing.T) {

		idVersion, err := client.Version()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, idVersion, 0)

		idGetAPIReference, err := client.GetAPIReference()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, idGetAPIReference, 0)

		idBuildInfo, err := client.GetBuildInfo()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, idBuildInfo, 0)

		t.Run("TestVersion", func(t *testing.T) {
			valueVersion, err := client.GetResp(idVersion)
			assert.Equal(t, nil, err)
			resultValue := valueVersion.(domain.ResultOfVersion)
			assert.Equal(t, VersionLibSDK, resultValue.Version)
		})

		t.Run("TestGetApiReference", func(t *testing.T) {
			valueAPIReference, err := client.GetResp(idGetAPIReference)
			assert.Equal(t, nil, err)
			resultValue := valueAPIReference.(domain.ResultOfGetAPIReference)
			assert.Equal(t, "1.0.0", resultValue.API.Version)
		})

		t.Run("TestBuildInfo", func(t *testing.T) {
			valueBuildInfo, err := client.GetResp(idBuildInfo)
			assert.Equal(t, nil, err)
			resultValue := valueBuildInfo.(domain.ResultOfBuildInfo)
			assert.Equal(t, "BuildNumber", resultValue.BuildNumber)
		})
	})
}
