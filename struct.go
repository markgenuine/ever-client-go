package goton

const (
	TONInputEncoding_text          = "text"
	TONInputEncoding_hex           = "hex"
	TONInputEncoding_base64        = "base64"
	TONOutputEncoding_Text         = "Text"
	TONOutputEncoding_Hex          = "Hex"
	TONOutputEncoding_HexUppercase = "HexUppercase"
	TONOutputEncoding_Base64       = "Base64"
)

var (
	chains                = getChains()
	lensMnemonic          = getlengthCounWordsInMnemonic()
	TONMnemonicDictionary = getTONMnemonicDictionary()
)

//TONKey...
type TONKey struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

type InputMessage struct {
	Text   string `json:"text,omitempty"`
	Hex    string `json:"hex,omitempty"`
	Base64 string `json:"base64,omitempty"`
}

type MaxFactorizeResult struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ModularPowerRequest struct {
	Base     string `json:"base"`
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

type RandomGenerateBytesRequest struct {
	Length         int    `json:"length"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

type MnemonicStructRequest struct {
	*InputMessage `json:"entropy,omitempty"`
	Dictionary    int    `json:"dictionary,omitempty"`
	WordCount     int    `json:"wordCount,omitempty"`
	Phrase        string `json:"phrase,omitempty"`
}

type MessageInputMessage struct {
	*InputMessage `json:"message"`
}

type ScryptDate struct {
	Password       *InputMessage `json:"password"`
	Salt           *InputMessage `json:"salt"`
	LogN           int           `json:"logN"`
	R              int           `json:"r"`
	P              int           `json:"p"`
	DkLen          int           `json:"dkLen"`
	OutputEncoding string        `json:"outputEncoding,omitempty"`
}

type NaclBoxIn struct {
	*InputMessage  `json:"message"`
	Nonce          string `json:"nonce"`
	TheirPublicKey string `json:"theirPublicKey"`
	SecretKey      string `json:"secretKey"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

type NaclSecretBox struct {
	*InputMessage  `json:"message"`
	Nonce          string `json:"nonce"`
	Key            string `json:"key"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

type NaclSign struct {
	*InputMessage  `json:"message"`
	Key            string `json:"key"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

type HDSerialized struct {
	Serialized string `json:"serialized"`
}

type HDDerivery struct {
	Serialized string `json:"serialized"`
	Index      int    `json:"index"`
	Hardened   bool   `json:"hardened"`
	Compliant  bool   `json:"compliant"`
}

type HDPathDerivery struct {
	Serialized string `json:"serialized"`
	Path       string `json:"path"`
	Compliant  bool   `json:"compliant"`
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

func getChains() map[int]string {
	return map[int]string{0: "net.ton.dev", 1: "main.ton.dev"}
}

func getlengthCounWordsInMnemonic() map[int]string {
	return map[int]string{12: "12", 15: "15", 18: "18", 21: "21", 24: "24"}
}

func getTONMnemonicDictionary() map[string]int {
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

func stringWithQuer(inStr string) string {
	return `"` + inStr + `"`
}
