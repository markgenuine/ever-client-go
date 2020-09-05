package goton

import (
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
