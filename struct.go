package goton

var chains = map[int]string{0: "net.ton.dev", 1: "main.ton.dev"}

//TONKey struct with public and secret keys
type TONKey struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

//TomlConfig struct with config data
type TomlConfig struct {
	BaseURL                            string  `toml:"base_URL" json:"base_URL"`
	MessageRetriesCount                int     `toml:"message_retries_count" json:"message_retries_count"`
	MessageExpirationTimeout           int     `toml:"message_expiration_timeout" json:"message_expiration_timeout"`
	MessageExpirationTimeoutGrowFactor float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor"`
	MessageProcessingTimeout           int     `toml:"message_processing_timeout" json:"message_processing_timeout"`
	WaitForTimeout                     int     `toml:"wait_for_timeout" json:"wait_for_timeout"`
	AccessKey                          string  `toml:"access_key" json:"access_key"`
	OutOfSyncThreshold                 int     `toml:"out_of_sync_threshold" json:"out_of_sync_threshold"`
}

//NewConfig create new config for connect client
//chanID 0-devnet, 1-mainnet,
func NewConfig(chanID int) *TomlConfig {
	config := TomlConfig{}
	config.BaseURL = chains[chanID]
	config.MessageRetriesCount = 10
	config.MessageExpirationTimeout = 10000 //ms
	config.MessageExpirationTimeoutGrowFactor = 1.5
	config.MessageProcessingTimeout = 40000 //ms
	config.WaitForTimeout = 40000           //ms
	config.OutOfSyncThreshold = 15000       //ms

	return &config
}
