package domain

import "encoding/json"

const (
	// SortDirectionASC ...
	SortDirectionASC = "ASC"

	// SortDirectionDESC ...
	SortDirectionDESC = "DESC"
)

// NetErrorCode ...
var NetErrorCode map[string]int

type (
	// OrderBy ...
	OrderBy struct {
		Path      string `json:"path"`
		Direction string `json:"direction"`
	}

	// ParamsOfQueryCollection ...
	ParamsOfQueryCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
		Order      []OrderBy   `json:"order,omitempty"`
		Limit      int         `json:"limit,omitempty"`
	}

	// ResultOfQueryCollection ...
	ResultOfQueryCollection struct {
		Result []json.RawMessage `json:"result"`
	}

	// ParamsOfWaitForCollection ...
	ParamsOfWaitForCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
		Timeout    int         `json:"timeout,omitempty"`
	}

	// ResultOfWaitForCollection ...
	ResultOfWaitForCollection struct {
		Result json.RawMessage `json:"result"`
	}

	// ResultOfSubscribeCollection ...
	ResultOfSubscribeCollection struct {
		Handle int `json:"handle"`
	}

	// ParamsOfSubscribeCollection ...
	ParamsOfSubscribeCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
	}

	// NetUseCase ...
	NetUseCase interface {
		QueryCollection(*ParamsOfQueryCollection) (*ResultOfQueryCollection, error)
		WaitForCollection(*ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error)
		Unsubscribe(*ResultOfSubscribeCollection) error
		SubscribeCollection(*ParamsOfSubscribeCollection) (<-chan interface{}, *ResultOfSubscribeCollection, error)
	}
)

func init() {
	
	NetErrorCode = map[string]int{
		"QueryFailed":  601,
		"SubscribeFailed": 602,
		"WaitForFailed": 603,
		"GetSubscriptionResultFailed": 604,
		"InvalidServerResponse": 605,
		"ClockOutOfSync": 606,
		"WaitForTimeout": 607,
		"GraphqlError": 608,
		"NetworkModuleSuspended": 609,
		"WebsocketDisconnected": 610,
		"NotSupported": 611,
		"NoEndpointsProvided": 612,
	}	
}

func getSortDirection() map[int]string {
	return map[int]string{
		0: "ASC",
		1: "DESC",
	}
}
