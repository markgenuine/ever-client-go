package domain

const(
	BaseUrl = "https://net.ton.dev/"
	BaseCustomUrl = "https://tonos.freeton.surf"
	BaseMainUrl = "https://main.ton.dev"
)

type (
	// Config ...
	Config struct {
		Network *Network   `toml:"network" json:"network,omitempty"`
		Crypto  *Crypto    `toml:"crypto" json:"crypto,omitempty"`
		Abi     *AbiConfig `toml:"abi" json:"abi,omitempty"`
		Boc     *BocConfig `toml:"boc" json:"boc,omitempty"`
	}

	// Network - Network config.
	Network struct {
		ServerAddress            string   `toml:"server_address" json:"server_address,omitempty"`
		Endpoints                []string `toml:"endpoints" json:"endpoints,omitempty"`
		NetworkRetriesCount      int      `toml:"network_retries_count" json:"network_retries_count,omitempty"`
		MaxReconnectTimeOut      int      `toml:"max_reconnect_timeout" json:"max_reconnect_timeout,omitempty"`
		ReconnectTimeout         int      `toml:"reconnect_timeout" json:"reconnect_timeout,omitempty"`
		MessageRetriesCount      int      `toml:"message_retries_count" json:"message_retries_count,omitempty"`
		MessageProcessingTimeout int      `toml:"message_processing_timeout" json:"message_processing_timeout,omitempty"`
		WaitForTimeout           int      `toml:"wait_for_timeout" json:"wait_for_timeout,omitempty"`
		OutOfSyncThreshold       int      `toml:"out_of_sync_threshold" json:"out_of_sync_threshold,omitempty"`
		SendingEndpointCount     *int      `toml:"sending_endpoint_count,omitempty" json:"sending_endpoint_count"`
		AccessKey                string   `toml:"access_key" json:"access_key,omitempty"`
	}

	// Crypto ...
	Crypto struct {
		MnemonicDictionary  int    `toml:"mnemonic_dictionary" json:"mnemonic_dictionary,omitempty"`
		MnemonicWordCount   int    `toml:"mnemonic_word_count" json:"mnemonic_word_count,omitempty"`
		HdKeyDerivationPath string `toml:"hdkey_derivation_path" json:"hdkey_derivation_path,omitempty"`
	}

	// AbiConfig ...
	AbiConfig struct {
		WorkChain                          int     `toml:"workchain" json:"workchain,omitempty"`
		MessageExpirationTimeout           int     `toml:"message_expiration_timeout" json:"message_expiration_timeout"`
		MessageExpirationTimeoutGrowFactor float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor"`
	}

	// BocConfig ...
	BocConfig struct {
		CacheMaxSize int `toml:"cache_max_size" json:"cache_max_size,omitempty"`
	}
)

// NewDefaultConfig create new config for connect client. ChanID 0-localhost, 1-devnet, 2-mainnet,
func NewDefaultConfig(address string) Config {
	config := Config{
		Network: &Network{
			ServerAddress:           address,
			NetworkRetriesCount:      5,
			MessageRetriesCount:      5,
			MessageProcessingTimeout: 40000, //ms
			WaitForTimeout:           40000, //ms
			OutOfSyncThreshold:       15000, //ms
			AccessKey:                "",
		},
		Crypto: &Crypto{
			MnemonicDictionary:  1,
			MnemonicWordCount:   12,
			HdKeyDerivationPath: "",
		},
		Abi: &AbiConfig{
			WorkChain:                          0,
			MessageExpirationTimeout:           40000, //ms
			MessageExpirationTimeoutGrowFactor: 1.5},
		Boc: &BocConfig{
			CacheMaxSize: 10,
		},
	}

	return config
}

// WordCountList list length mnemonic phrases
func WordCountList() map[int]int {
	return map[int]int{12: 12, 15: 15, 18: 18, 21: 21, 24: 24}
}

// DictionaryList list dictionary mnemonic phrase
func DictionaryList() map[string]int {
	return map[string]int{
		"TON":                 0,
		"ENGLISH":             1,
		"CHINESE_SIMPLIFIED":  2,
		"CHINESE_TRADITIONAL": 3,
		"FRENCH":              4,
		"ITALIAN":             5,
		"JAPANESE":            6,
		"KOREAN":              7,
		"SPANISH":             8,
	}
}
