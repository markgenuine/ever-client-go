package domain

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
	}

	// ResultOfAppDebotBractionowser ...
	ResultOfAppDebotBractionowser struct {
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

	// DebotUseCase ...
	DebotUseCase interface{}
)
