package client

import (
	"github.com/markgenuine/ever-client-go/util"
	"testing"

	"github.com/markgenuine/ever-client-go/domain"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	configConn := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
	clientConn, err := NewClientGateway(configConn)
	assert.Equal(t, nil, err)
	defer clientConn.Destroy()

	t.Run("TestConfigFields", func(t *testing.T) {
		defConf := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
		defConf.Abi.MessageExpirationTimeout = util.IntToPointerInt(0)
		defConf.Network.MaxReconnectTimeOut = util.IntToPointerInt(100)
		assert.Equal(t, defConf.Crypto.MnemonicWordCount, util.IntToPointerInt(domain.DefaultWordCount))
		assert.Equal(t, defConf.Network.MaxReconnectTimeOut, util.IntToPointerInt(100))
	})

	t.Run("TestVersion", func(t *testing.T) {
		version, err := clientConn.Version()
		assert.Equal(t, nil, err)
		assert.Equal(t, VersionLibSDK, version.Version)
	})

	t.Run("TestGetApiReference", func(t *testing.T) {
		getAPIReference, err := clientConn.GetAPIReference()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, len(getAPIReference.API.Modules), 0)
		assert.Equal(t, VersionLibSDK, getAPIReference.API.Version)
	})

	t.Run("TestBuildInfo", func(t *testing.T) {
		buildInfo, err := clientConn.GetBuildInfo()
		assert.Equal(t, nil, err)
		assert.NotNil(t, buildInfo.BuildNumber)
	})
}
