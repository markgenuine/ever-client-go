package goton

import (
	"encoding/hex"
	"encoding/json"

	"github.com/BurntSushi/toml"
)

// ParseConfigFile parse TOML config file
func ParseConfigFile(path string) (*TomlConfig, error) {
	var config TomlConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func structToJSON(structValue interface{}) string {
	request, _ := json.Marshal(structValue)
	return string(request)
}

func fromHex(value string) []byte {
	src := []byte(value)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, _ := hex.Decode(dst, src)

	return dst[:n]
}
