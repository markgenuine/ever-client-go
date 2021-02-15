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

	// ParamsOfQuery ...
	ParamsOfQuery struct {
		Query     string          `json:"query"`
		Variables json.RawMessage `json:"variables"`
	}

	//ResultOfQuery ...
	ResultOfQuery struct {
		Result json.RawMessage `json:"result"`
	}

	//ParamsOfQueryOperation ...
	ParamsOfQueryOperation json.RawMessage

	// ParamsOfBatchQuery
	ParamsOfBatchQuery struct {
		Operations []ParamsOfQueryOperation `json:"operations"`
	}

	// ResultOfBatchQuery
	ResultOfBatchQuery struct {
		Result []json.RawMessage `json:"result"`
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

	FieldsAggregation json.RawMessage

	// ParamsOfAggregateCollection ...
	ParamsOfAggregateCollection struct {
		Collection string              `json:"collection"`
		Filter     interface{}         `json:"filter,omitempty"`
		Fields     []FieldsAggregation `json:"fields,omitempty"`
	}
	// ResultOfAggregateCollection ...
	ResultOfAggregateCollection struct {
		Values json.RawMessage `json:"values"`
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

	// ParamsOfFindLastShardBlock ...
	ParamsOfFindLastShardBlock struct {
		Address string `json:"address""`
	}

	// ResultOfFindLastShardBlock ...
	ResultOfFindLastShardBlock struct {
		BlockID string `json:"block_id"`
	}

	// EndpointsSet ...
	EndpointsSet struct {
		Endpoints []string `json:"endpoints"`
	}

	// NetUseCase ...
	NetUseCase interface {
		Query(*ParamsOfQuery) (*ResultOfQuery, error)
		BatchQuery(*ParamsOfBatchQuery) (*ResultOfBatchQuery, error)
		QueryCollection(*ParamsOfQueryCollection) (*ResultOfQueryCollection, error)
		AggregateCollection(*ParamsOfAggregateCollection) (*ResultOfAggregateCollection, error)
		WaitForCollection(*ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error)
		Unsubscribe(*ResultOfSubscribeCollection) error
		SubscribeCollection(*ParamsOfSubscribeCollection) (<-chan interface{}, *ResultOfSubscribeCollection, error)
		Suspend() error
		Resume() error
		FindLastShardBlock(*ParamsOfFindLastShardBlock) (*ResultOfFindLastShardBlock, error)
		FetchEndpoints() (*EndpointsSet, error)
		SetEndpoints(*EndpointsSet) error
	}
)

func init() {

	NetErrorCode = map[string]int{
		"QueryFailed":                 601,
		"SubscribeFailed":             602,
		"WaitForFailed":               603,
		"GetSubscriptionResultFailed": 604,
		"InvalidServerResponse":       605,
		"ClockOutOfSync":              606,
		"WaitForTimeout":              607,
		"GraphqlError":                608,
		"NetworkModuleSuspended":      609,
		"WebsocketDisconnected":       610,
		"NotSupported":                611,
		"NoEndpointsProvided":         612,
		"GraphqlWebsocketInitError":   613,
		"NetworkModuleResumed":        614,
	}
}

func getSortDirection() map[int]string {
	return map[int]string{
		0: "ASC",
		1: "DESC",
	}
}
