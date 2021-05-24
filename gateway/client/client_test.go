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
		t.Run("TestVersion", func(t *testing.T) {
			version, err := client.Version()
			assert.Equal(t, nil, err)
			assert.Equal(t, VersionLibSDK, version.Version)
		})

		t.Run("TestGetApiReference", func(t *testing.T) {
			getAPIReference, err := client.GetAPIReference()
			assert.Equal(t, nil, err)
			assert.Equal(t, VersionLibSDK, getAPIReference.API.Version)
		})

		t.Run("TestBuildInfo", func(t *testing.T) {
			buildInfo, err := client.GetBuildInfo()
			assert.Equal(t, nil, err)
			assert.Equal(t, 0, buildInfo.BuildNumber)
		})
	})
}
