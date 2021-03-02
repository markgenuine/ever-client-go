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
	}

	// ParamsOfAppDebotBrowser ...
	ParamsOfAppDebotBrowser struct {
		Type      string       `json:"type"`
		Msg       string       `json:"msg,omitempty"`
		ContextID int          `json:"context_id,omitempty"`
		Action    *DebotAction `json:"action,omitempty"`
		Prompt    string       `json:"prompt,omitempty"`
		DebotAddr string       `json:"debot_addr,omitempty"`
		Message string `json:"message,omitempty"`
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
		Source string `json:"source"`
		FuncID int `json:"func_id"`
		Params string `json:"params"`
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
		"DebotStartFailed":     801,
		"DebotFetchFailed":     802,
		"DebotExecutionFailed": 803,
		"DebotInvalidHandle":   804,
		"DebotInvalidJsonParams" : 805,
		"DebotInvalidFunctionId" : 806,
		"DebotInvalidAbi" : 807,
		"DebotGetMethodFailed" : 808,
		"DebotInvalidMsg" : 809,
		"DebotExternaCallFailed" : 810,
	}
}

// ParamsOfAppDebotBrowserLogNew ...
func ParamsOfAppDebotBrowserLogNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Log"}
}

// ParamsOfAppDebotBrowserSwitchNew ...
func ParamsOfAppDebotBrowserSwitchNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Switch"}
}

// ParamsOfAppDebotBrowserSwitchCompletedNew ...
func ParamsOfAppDebotBrowserSwitchCompletedNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "SwitchCompleted"}
}

// ParamsOfAppDebotBrowserShowActionNew...
func ParamsOfAppDebotBrowserShowActionNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "ShowAction"}
}

// ParamsOfAppDebotBrowserInputNew
func ParamsOfAppDebotBrowserInputNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Input"}
}

// ParamsOfAppDebotBrowserGetSigningBoxNew
func ParamsOfAppDebotBrowserGetSigningBoxNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "GetSigningBox"}
}

// ParamsOfAppDebotBrowserInvokeDebotNew
func ParamsOfAppDebotBrowserInvokeDebotNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "InvokeDebot"}
}

// ParamsOfAppDebotBrowserSendNew
func ParamsOfAppDebotBrowserSendNew() *ParamsOfAppDebotBrowser {
	return &ParamsOfAppDebotBrowser{Type: "Send"}
}

// ResultOfAppDebotBrowserInput
func ResultOfAppDebotBrowserInput() *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "Input"}
}

// ResultOfAppDebotBrowserGetSigningBox
func ResultOfAppDebotBrowserGetSigningBox() *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "GetSigningBox"}
}

// ResultOfAppDebotBrowserInvokeDebot
func ResultOfAppDebotBrowserInvokeDebot() *ResultOfAppDebotBrowser {
	return &ResultOfAppDebotBrowser{Type: "InvokeDebot"}
}

