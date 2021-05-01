package domain

import "math/big"

// DebotErrorCode ...
var DebotErrorCode map[string]int

const (
	// DebotActivityTransactionType is DebotActivityType with value "Transaction"
	DebotActivityTransactionType DebotActivityType = "Transaction"
)

type (
	// DebotHandle - Handle of registered in SDK debot.
	DebotHandle = int

	// DebotActivityType - mode for DebotActivity
	DebotActivityType string

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

	// DebotActivity - Describes the operation that the DeBot wants to perform.
	DebotActivity struct {
		Type    DebotActivityType `json:"type"`
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
		Type      string         `json:"type"`
		Msg       string         `json:"msg,omitempty"`
		ContextID int            `json:"context_id,omitempty"`
		Action    *DebotAction   `json:"action,omitempty"`
		Prompt    string         `json:"prompt,omitempty"`
		DebotAddr string         `json:"debot_addr,omitempty"`
		Message   string         `json:"message,omitempty"`
		Activity    *DebotActivity `json:"activity,omitempty"`
	}

	// ResultOfAppDebotBrowser - Returning values from Debot Browser callbacks.
	ResultOfAppDebotBrowser struct {
		Type       string            `json:"type"`
		Value      string            `json:"value,omitmepty"`
		SigningBox *SigningBoxHandle `json:"signing_box,omitempty"`
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

// DebotActivityTransaction - Variant constructors transaction.
func DebotActivityTransaction(msg, dst string, out []Spending, fee *big.Int, setcode bool, signkey string) *DebotActivity {
	return &DebotActivity{Type: DebotActivityTransactionType, Msg: msg, Dst: dst, Out: out, Fee: fee, Setcode: setcode, Signkey: signkey}
}

// ParamsOfAppDebotBrowserLog variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserLog(msg string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Log", Msg: msg}
}

// ParamsOfAppDebotBrowserSwitch variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserSwitch(contextID int) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Switch", ContextID: contextID}
}

// ParamsOfAppDebotBrowserSwitchCompleted variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserSwitchCompleted() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "SwitchCompleted"}
}

// ParamsOfAppDebotBrowserShowAction variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserShowAction(action *DebotAction) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "ShowAction", Action: action}
}

// ParamsOfAppDebotBrowserInput variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserInput(prompt string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Input", Prompt: prompt}
}

// ParamsOfAppDebotBrowserGetSigningBox variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserGetSigningBox() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "GetSigningBox"}
}

// ParamsOfAppDebotBrowserInvokeDebot variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserInvokeDebot(debotAddr string, action *DebotAction) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "InvokeDebot", DebotAddr: debotAddr, Action: action}
}

// ParamsOfAppDebotBrowserSend variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserSend(message string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Send", Message: message}
}

// ParamsOfAppDebotBrowserApprove variant constructor ParamsOfAppDebotBrowser.
func ParamsOfAppDebotBrowserApprove(activity *DebotActivity) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Approve", Activity: activity}
}

// ResultOfAppDebotBrowserInput ...
func ResultOfAppDebotBrowserInput(value string) *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "Input", Value: value}
}

// ResultOfAppDebotBrowserGetSigningBox ...
func ResultOfAppDebotBrowserGetSigningBox(signingBox *SigningBoxHandle) *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "GetSigningBox", SigningBox: signingBox}
}

// ResultOfAppDebotBrowserInvokeDebot ...
func ResultOfAppDebotBrowserInvokeDebot() *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "InvokeDebot"}
}
