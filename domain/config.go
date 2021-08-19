package domain

import "github.com/move-ton/ton-client-go/util"

const (
	BaseCustomUrl = "https://tonos.freeton.surf/"
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
		NetworkRetriesCount      *int     `toml:"network_retries_count" json:"network_retries_count,omitempty"`
		MaxReconnectTimeOut      *int     `toml:"max_reconnect_timeout" json:"max_reconnect_timeout,omitempty"`
		ReconnectTimeout         *int     `toml:"reconnect_timeout" json:"reconnect_timeout,omitempty"`
		MessageRetriesCount      *int     `toml:"message_retries_count" json:"message_retries_count,omitempty"`
		MessageProcessingTimeout *int     `toml:"message_processing_timeout" json:"message_processing_timeout,omitempty"`
		WaitForTimeout           *int     `toml:"wait_for_timeout" json:"wait_for_timeout,omitempty"`
		OutOfSyncThreshold       *int     `toml:"out_of_sync_threshold" json:"out_of_sync_threshold,omitempty"`
		SendingEndpointCount     *int     `toml:"sending_endpoint_count,omitempty" json:"sending_endpoint_count,omitempty"`
		LatencyDetectionInterval *int     `toml:"latency_detection_interval,omitempty" json:"latency_detection_interval,omitempty"`
		MaxLatency               *int     `toml:"max_latency,omitempty" json:"max_latency,omitempty"`
		AccessKey                string   `toml:"access_key" json:"access_key,omitempty"`
	}

	// Crypto ...
	Crypto struct {
		MnemonicDictionary  *int   `toml:"mnemonic_dictionary" json:"mnemonic_dictionary,omitempty"`
		MnemonicWordCount   *int   `toml:"mnemonic_word_count" json:"mnemonic_word_count,omitempty"`
		HdKeyDerivationPath string `toml:"hdkey_derivation_path" json:"hdkey_derivation_path,omitempty"`
	}

	// AbiConfig ...
	AbiConfig struct {
		WorkChain                          *int     `toml:"workchain" json:"workchain,omitempty"`
		MessageExpirationTimeout           *int     `toml:"message_expiration_timeout" json:"message_expiration_timeout,omitempty"`
		MessageExpirationTimeoutGrowFactor *float32 `toml:"message_expiration_timeout_grow_factor" json:"message_expiration_timeout_grow_factor,omitempty"`
	}

	// BocConfig ...
	BocConfig struct {
		CacheMaxSize *int `toml:"cache_max_size" json:"cache_max_size,omitempty"`
	}
)

// NewDefaultConfig create new config for connect client.
func NewDefaultConfig(address string, endPoints []string) Config {
	config := Config{
		Network: &Network{
			ServerAddress:            address,
			Endpoints:                endPoints,
			NetworkRetriesCount:      util.IntToPointerInt(5),
			MessageRetriesCount:      util.IntToPointerInt(5),
			MessageProcessingTimeout: util.IntToPointerInt(40000), //ms
			WaitForTimeout:           util.IntToPointerInt(40000), //ms
			OutOfSyncThreshold:       util.IntToPointerInt(15000), //ms
			AccessKey:                "",
		},
		Crypto: &Crypto{
			MnemonicDictionary:  util.IntToPointerInt(1),
			MnemonicWordCount:   util.IntToPointerInt(12),
			HdKeyDerivationPath: "",
		},
		Abi: &AbiConfig{
			WorkChain:                          util.IntToPointerInt(0),
			MessageExpirationTimeout:           util.IntToPointerInt(40000), //ms
			MessageExpirationTimeoutGrowFactor: util.Float32ToPointerFloat32(1.5)},
	}

	return config
}

// WordCountList list length mnemonic phrases
func WordCountList() map[int]*int {
	return map[int]*int{12: util.IntToPointerInt(12), 15: util.IntToPointerInt(15), 18: util.IntToPointerInt(18), 21: util.IntToPointerInt(21), 24: util.IntToPointerInt(24)}
}

// DictionaryList list dictionary mnemonic phrase
func DictionaryList() map[string]*int {
	return map[string]*int{
		"TON":                 util.IntToPointerInt(0),
		"ENGLISH":             util.IntToPointerInt(1),
		"CHINESE_SIMPLIFIED":  util.IntToPointerInt(2),
		"CHINESE_TRADITIONAL": util.IntToPointerInt(3),
		"FRENCH":              util.IntToPointerInt(4),
		"ITALIAN":             util.IntToPointerInt(5),
		"JAPANESE":            util.IntToPointerInt(6),
		"KOREAN":              util.IntToPointerInt(7),
		"SPANISH":             util.IntToPointerInt(8),
	}
}

// GetDevNetBaseUrls ...
func GetDevNetBaseUrls() []string {
	return []string{"https://net1.ton.dev/", "https://net5.ton.dev/"}
}

// GetMainNetBaseUrls ...
func GetMainNetBaseUrls() []string {
	return []string{"https://main2.ton.dev/", "https://main3.ton.dev/", "https://main4.ton.dev/"}
}

// GetLocalNetBaseUrls ...
func GetLocalNetBaseUrls() []string {
	return []string{"http://0.0.0.0/", "http://127.0.0.1/", "http://localhost/"}
}
