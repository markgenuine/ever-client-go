package domain

import "math/big"

// DebotErrorCode ...
var DebotErrorCode map[string]int

type (
	// DebotHandle - Handle of registered in SDK debot.
	DebotHandle = int

	// DebotAction - Describes a debot action in a Debot Context.
	DebotAction struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		ActionType  int    `json:"action_type"`
		To          int    `json:"to"`
		Attributes  string `json:"attributes"`
		Misc        string `json:"misc"`
	}

	// DebotInfo - Describes DeBot metadata.
	DebotInfo struct {
		Name       string   `json:"name,omitempty"`
		Version    string   `json:"version,omitempty"`
		Publisher  string   `json:"publisher,omitempty"`
		Key        string   `json:"key,omitempty"`
		Author     string   `json:"author,omitempty"`
		Support    string   `json:"support,omitempty"`
		Hello      string   `json:"hello,omitempty"`
		Language   string   `json:"language,omitempty"`
		Dabi       string   `json:"dabi,omitempty"`
		Icon       string   `json:"icon,omitempty"`
		interfaces []string `json:"interfaces"`
	}

	DebotActivity struct {
		ValueEnumType interface{}
	}

	// DebotActivityTransaction ...
	DebotActivityTransaction struct {
		Msg     string            `json:"msg"`
		Dst     string            `json:"dst"`
		Out     []Spending        `json:"Out"`
		Fee     *big.Int          `json:"fee"`
		Setcode bool              `json:"setcode"`
		Signkey string            `json:"signkey"`
	}

	// Spending - Describes how much funds will be debited from the target contract balance as a result of the transaction.
	Spending struct {
		Amount *big.Int
		Dst    string
	}

	// ParamsOfInit - Parameters to init DeBot.
	ParamsOfInit struct {
		Address string `json:"address"`
	}

	// ParamsOfStart - Parameters to start DeBot. DeBot must be already initialized with init() function.
	ParamsOfStart struct {
		DebotHandle DebotHandle `json:"debot_handle"`
	}

	// RegisteredDebot - Structure for storing debot handle returned from init function.
	RegisteredDebot struct {
		DebotHandle DebotHandle `json:"debot_handle"`
		DebotAbi    string      `json:"debot_abi"`
		Info        *DebotInfo  `json:"info"`
	}

	// ParamsOfAppDebotBrowser - Debot Browser callbacks.
	ParamsOfAppDebotBrowser struct {
		ValueEnumType interface{}
	}

	// ParamsOfAppDebotBrowserLog ...
	ParamsOfAppDebotBrowserLog struct {
		Msg       string         `json:"msg"`
	}

	// ParamsOfAppDebotBrowserSwitch ...
	ParamsOfAppDebotBrowserSwitch struct {
		ContextID int            `json:"context_id"`
	}

	// ParamsOfAppDebotBrowserSwitchCompleted ...
	ParamsOfAppDebotBrowserSwitchCompleted struct {}

	// ParamsOfAppDebotBrowserShowAction ...
	ParamsOfAppDebotBrowserShowAction struct {
		Action    *DebotAction   `json:"action"`
	}

	// ParamsOfAppDebotBrowserInput ...
	ParamsOfAppDebotBrowserInput struct {
		Prompt    string         `json:"prompt"`
	}

	// ParamsOfAppDebotBrowserGetSigningBox ...
	ParamsOfAppDebotBrowserGetSigningBox struct {}

	// ParamsOfAppDebotBrowserInvokeDebot ...
	ParamsOfAppDebotBrowserInvokeDebot struct {
		DebotAddr string         `json:"debot_addr"`
		Action    *DebotAction   `json:"action"`
	}

	// ParamsOfAppDebotBrowserSend ...
	ParamsOfAppDebotBrowserSend struct {
		Message   string         `json:"message"`
	}

	// ParamsOfAppDebotBrowserApprove ...
	ParamsOfAppDebotBrowserApprove struct {
		Activity  *DebotActivity `json:"activity"`
	}

	// ResultOfAppDebotBrowser - Returning values from Debot Browser callbacks.
	ResultOfAppDebotBrowser struct {
		ValueEnumType interface{}
	}

	// ResultOfAppDebotBrowserInput ...
	ResultOfAppDebotBrowserInput struct {
		Value      string            `json:"value"`
	}

	// ResultOfAppDebotBrowserGetSigningBox ...
	ResultOfAppDebotBrowserGetSigningBox struct {
		SigningBox SigningBoxHandle `json:"signing_box"`
	}

	// ResultOfAppDebotBrowserInvokeDebot ...
	ResultOfAppDebotBrowserInvokeDebot struct {}

	// ResultOfAppDebotBrowserApprove ...
	ResultOfAppDebotBrowserApprove struct {
		Approved bool `json:"approved"`
	}

	// ParamsOfFetch - Parameters to fetch DeBot metadata.
	ParamsOfFetch struct {
		Address string `json:"address"`
	}

	// ResultOfFetch ...
	ResultOfFetch struct {
		Info *DebotInfo `json:"info"`
	}

	// ParamsOfExecute - Parameters for executing debot action.
	ParamsOfExecute struct {
		DebotHandle DebotHandle  `json:"debot_handle"`
		Action      *DebotAction `json:"action"`
	}

	// ParamsOfSend - Parameters of send function.
	ParamsOfSend struct {
		DebotHandle DebotHandle `json:"debot_handle"`
		Message     string      `json:"message"`
	}

	// ParamsOfRemove ...
	ParamsOfRemove struct {
		DebotHandle DebotHandle `json:"debot_handle"`
	}

	// DebotUseCase ...
	DebotUseCase interface {
		Init(*ParamsOfInit, AppDebotBrowser) (*RegisteredDebot, error)
		Start(*ParamsOfStart) error
		Fetch(*ParamsOfFetch) (*ResultOfFetch, error)
		Execute(*ParamsOfExecute) error
		Send(*ParamsOfSend) error
		Remove(*ParamsOfRemove) error
	}
)

func init() {
	DebotErrorCode = map[string]int{
		"DebotStartFailed":           801,
		"DebotFetchFailed":           802,
		"DebotExecutionFailed":       803,
		"DebotInvalidHandle":         804,
		"DebotInvalidJsonParams":     805,
		"DebotInvalidFunctionId":     806,
		"DebotInvalidAbi":            807,
		"DebotGetMethodFailed":       808,
		"DebotInvalidMsg":            809,
		"DebotExternalCallFailed":    810,
		"DebotBrowserCallbackFailed": 811,
		"DebotOperationRejected":     812,
	}
}