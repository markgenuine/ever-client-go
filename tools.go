package goton

import (
	"github.com/BurntSushi/toml"
)

//ParseConfigFile parse TOML config file
func ParseConfigFile(path string, chainID int) (*TomlConfig, error) {
	var config TomlConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		//log.Fatal("Error read config file", err)
		return nil, err
	}

	return &config, nil
}
