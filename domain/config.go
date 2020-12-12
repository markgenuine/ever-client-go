package domain

type (
	// Config ...
	Config struct {
		Network *Network   `toml:"network" json:"network,omitempty"`
		Crypto  *Crypto    `toml:"crypto" json:"crypto,omitempty"`
		Abi     *AbiConfig `toml:"abi" json:"abi,omitempty"`
	}

	// Network ...
	Network struct {
		ServerAddress            string `toml:"server_address" json:"server_address,omitempty"`
		NetworkRetriesCount      int    `toml:"network_retries_count" json:"network_retries_count,omitempty"`
		MessageRetriesCount      int    `toml:"message_retries_count" json:"message_retries_count,omitempty"`
		MessageProcessingTimeout int    `toml:"message_processing_timeout" json:"message_processing_timeout,omitempty"`
		WaitForTimeout           int    `toml:"wait_for_timeout" json:"wait_for_timeout,omitempty"`
		OutOfSyncThreshold       int    `toml:"out_of_sync_threshold" json:"out_of_sync_threshold,omitempty"`
		AccessKey                string `toml:"access_key" json:"access_key,omitempty"`
	}

	// Crypto ...
	Crypto struct {
		MnemonicDictionary  int    `toml:"mnemonic_dictionary" json:"mnemonic_dictionary,omitempty"`
		MnemonicWordCount   int    `toml:"mnemonic_word_count" json:"mnemonic_word_count,omitempty"`
		HdkeyDerivationPath string `toml:"hdkey_derivation_path" json:"hdkey_derivation_path,omitempty"`
		HdkeyCompliant      bool   `toml:"hdkey_compliant" json:"hdkey_compliant,omitempty"`
	}

	// AbiConfig ...
	AbiConfig struct {
		Workchain                          int     `toml:"workchain" json:"workchain,omitempty"`
		MessageExpirationTimeout           int     `toml:"message_expiration_timeout" json:"message_expiration_timeout"`
		MessageExpirationTimeoutGrowFactor float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor"`
	}
)

// NewDefaultConfig create new config for connect client. ChanID 0-localhost, 1-devnet, 2-mainnet,
func NewDefaultConfig(chainID int) Config {
	config := Config{
		Network: &Network{
			ServerAddress:            Chains()[chainID],
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
			HdkeyDerivationPath: "",
			HdkeyCompliant:      true,
		},
		Abi: &AbiConfig{
			Workchain:                          0,
			MessageExpirationTimeout:           40000, //ms
			MessageExpirationTimeoutGrowFactor: 1.5},
	}

	return config
}

// Chains list endpoints
func Chains() map[int]string {
	return map[int]string{0: "localhost", 1: "net.ton.dev", 2: "main.ton.dev"}
}

// WordCounList list length mnemonic phrases
func WordCounList() map[int]int {
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
