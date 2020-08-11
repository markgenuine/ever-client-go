package goton

import (
	"fmt"
	"testing"
)

func TestClientMethods(t *testing.T) {

	config := NewConfig(0)

	t.Run("TestNewConfig", func(t *testing.T) {
		tomlConf := &TomlConfig{}
		tomlConf.BaseURL = chains[0]
		tomlConf.MessageRetriesCount = 10
		tomlConf.MessageExpirationTimeout = 10000
		tomlConf.MessageExpirationTimeoutGrowFactor = 1.5
		tomlConf.MessageProcessingTimeout = 40000
		tomlConf.WaitForTimeout = 40000
		tomlConf.OutOfSyncThreshold = 15000

		if fmt.Sprintf("%T", tomlConf) != fmt.Sprintf("%T", config) {
			t.Errorf("test Failed - Default config type: %T different new config type: %T", tomlConf, config)
		}
	})

	t.Run("TestVersion", func(t *testing.T) {
		client, err := InitClient(config)
		if err != nil {
			t.Errorf("test Failed - Init client error: %s", err)
		}
		defer client.Destroy()

		value, err := client.Version()
		if err != nil {
			t.Errorf("test Failed - Error get version, err: %s", err)
		}

		if value != VersionLibSDK {
			t.Errorf("test Failed - Version lib %s different and version go-ton-sdk %s", value, VersionLibSDK)
		}
	})
}
