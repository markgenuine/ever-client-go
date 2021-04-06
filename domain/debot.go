package domain

// DebotErrorCode ...
var DebotErrorCode map[string]int

type (
	// DebotHandle ...
	DebotHandle = int

	// DebotAction ...
	DebotAction struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		ActionType  int    `json:"action_type"`
		To          int    `json:"to"`
		Attributes  string `json:"attributes"`
		Misc        string `json:"misc"`
	}

	// ParamsOfStart ...
	ParamsOfStart struct {
		Address string `json:"address"`
	}

	// RegisteredDebot ...
	RegisteredDebot struct {
		DebotHandle DebotHandle `json:"debot_handle"`
		DebotAbi    string      `json:"debot_abi"`
	}

	// ParamsOfAppDebotBrowser ...
	ParamsOfAppDebotBrowser struct {
		Type      string       `json:"type"`
		Msg       string       `json:"msg,omitempty"`
		ContextID int          `json:"context_id,omitempty"`
		Action    *DebotAction `json:"action,omitempty"`
		Prompt    string       `json:"prompt,omitempty"`
		DebotAddr string       `json:"debot_addr,omitempty"`
		Message   string       `json:"message,omitempty"`
	}

	// ResultOfAppDebotBrowser ...
	ResultOfAppDebotBrowser struct {
		Type       string            `json:"type"`
		Value      string            `json:"value,omitmepty"`
		SigningBox *SigningBoxHandle `json:"signing_box,omitempty"`
	}

	// ParamsOfFetch ...
	ParamsOfFetch struct {
		Address string `json:"address"`
	}

	// ParamsOfExecute ...
	ParamsOfExecute struct {
		DebotHandle DebotHandle `json:"debot_handle"`
		Action      DebotAction `json:"action"`
	}

	// ParamsOfSend ...
	ParamsOfSend struct {
		DebotHandle DebotHandle `json:"debot_handle"`
		Message     string      `json:"message"`
	}

	// DebotUseCase ...
	DebotUseCase interface {
		Start(*ParamsOfStart, AppDebotBrowser) (*RegisteredDebot, error)
		Fetch(*ParamsOfFetch, AppDebotBrowser) (*RegisteredDebot, error)
		Execute(*ParamsOfExecute) error
		Send(*ParamsOfSend) error
		Remove(*RegisteredDebot) error
	}
)

func init() {
	DebotErrorCode = map[string]int{
		"DebotStartFailed":        801,
		"DebotFetchFailed":        802,
		"DebotExecutionFailed":    803,
		"DebotInvalidHandle":      804,
		"DebotInvalidJsonParams":  805,
		"DebotInvalidFunctionId":  806,
		"DebotInvalidAbi":         807,
		"DebotGetMethodFailed":    808,
		"DebotInvalidMsg":         809,
		"DebotExternalCallFailed": 810,
	}
}

// ParamsOfAppDebotBrowserLog ...
func ParamsOfAppDebotBrowserLog(msg string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Log", Msg: msg}
}

// ParamsOfAppDebotBrowserSwitch ...
func ParamsOfAppDebotBrowserSwitch(contextID int) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Switch", ContextID: contextID}
}

// ParamsOfAppDebotBrowserSwitchCompleted ...
func ParamsOfAppDebotBrowserSwitchCompleted() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "SwitchCompleted"}
}

// ParamsOfAppDebotBrowserShowAction...
func ParamsOfAppDebotBrowserShowAction(action *DebotAction) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "ShowAction", Action: action}
}

// ParamsOfAppDebotBrowserInput ...
func ParamsOfAppDebotBrowserInput(prompt string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Input", Prompt: prompt}
}

// ParamsOfAppDebotBrowserGetSigningBox ...
func ParamsOfAppDebotBrowserGetSigningBox() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "GetSigningBox"}
}

// ParamsOfAppDebotBrowserInvokeDebot ...
func ParamsOfAppDebotBrowserInvokeDebot(debotAddr string, action *DebotAction) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "InvokeDebot", DebotAddr: debotAddr, Action: action}
}

// ParamsOfAppDebotBrowserSend ...
func ParamsOfAppDebotBrowserSend(message string) *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Send", Message: message}
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
