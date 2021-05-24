package domain

import (
	"encoding/json"
	"fmt"
)

const (
	// SortDirectionASC ...
	SortDirectionASC SortDirection = "ASC"

	// SortDirectionDESC ...
	SortDirectionDESC SortDirection = "DESC"

	AggregationFnTypeCount   AggregationFnType = "COUNT"
	AggregationFnTypeMin     AggregationFnType = "MIN"
	AggregationFnTypeMax     AggregationFnType = "MAX"
	AggregationFnTypeSum     AggregationFnType = "SUM"
	AggregationFnTypeAverage AggregationFnType = "AVERAGE"
)

// NetErrorCode ...
var NetErrorCode map[string]int

type (

	// SortDirection ...
	SortDirection string

	// OrderBy ...
	OrderBy struct {
		Path      string        `json:"path"`
		Direction SortDirection `json:"direction"`
	}

	// ParamsOfQueryOperation
	ParamsOfQueryOperation struct {
		ValueEnumType interface{}
	}

	// AggregationFnType ...
	AggregationFnType string

	// FieldAggregation ...
	FieldAggregation struct {
		Field string            `json:"field"`
		Fn    AggregationFnType `json:"fn"`
	}

	// ParamsOfQuery ...
	ParamsOfQuery struct {
		Query     string          `json:"query"`
		Variables json.RawMessage `json:"variables,omitempty"`
	}

	//ResultOfQuery ...
	ResultOfQuery struct {
		Result json.RawMessage `json:"result"`
	}

	// ParamsOfBatchQuery ...
	ParamsOfBatchQuery struct {
		Operations []ParamsOfQueryOperation `json:"operations"`
	}

	// ResultOfBatchQuery ...
	ResultOfBatchQuery struct {
		Result []json.RawMessage `json:"result"`
	}

	// ParamsOfQueryCollection ...
	ParamsOfQueryCollection struct {
		Collection string      `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string      `json:"result"`
		Order      []*OrderBy  `json:"order,omitempty"`
		Limit      int         `json:"limit,omitempty"`
	}

	// ResultOfQueryCollection ...
	ResultOfQueryCollection struct {
		Result []json.RawMessage `json:"result"`
	}

	// ParamsOfAggregateCollection ...
	ParamsOfAggregateCollection struct {
		Collection string              `json:"collection"`
		Filter     json.RawMessage         `json:"filter,omitempty"`
		Fields     []*FieldAggregation `json:"fields,omitempty"`
	}
	// ResultOfAggregateCollection ...
	ResultOfAggregateCollection struct {
		Values json.RawMessage `json:"values"`
	}

	// ParamsOfWaitForCollection ...
	ParamsOfWaitForCollection struct {
		Collection string      `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
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
		Filter     json.RawMessage `json:"filter,omitempty"`
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

	// ParamsOfQueryCounterparties
	ParamsOfQueryCounterparties struct {
		Account string `json:"account"`
		Result string `json:"result"`
		First int `json:"first,omitempty"`
		After string `json:"after,omitempty"`
	}

	// NetUseCase ...
	NetUseCase interface {
		Query(*ParamsOfQuery) (*ResultOfQuery, error)
		BatchQuery(*ParamsOfBatchQuery) (*ResultOfBatchQuery, error)
		QueryCollection(*ParamsOfQueryCollection) (*ResultOfQueryCollection, error)
		AggregateCollection(*ParamsOfAggregateCollection) (*ResultOfAggregateCollection, error)
		WaitForCollection(*ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error)
		Unsubscribe(*ResultOfSubscribeCollection) error
		SubscribeCollection(*ParamsOfSubscribeCollection) (<-chan json.RawMessage, *ResultOfSubscribeCollection, error)
		Suspend() error
		Resume() error
		FindLastShardBlock(*ParamsOfFindLastShardBlock) (*ResultOfFindLastShardBlock, error)
		FetchEndpoints() (*EndpointsSet, error)
		SetEndpoints(*EndpointsSet) error
		QueryCounterparties(*ParamsOfQueryCounterparties) (*ResultOfQueryCollection, error)
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

func (pOQO *ParamsOfQueryOperation) MarshalJSON()([]byte, error){
	switch value := (pOQO.ValueEnumType).(type) {
	case ParamsOfQueryCollection:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfQueryCollection
		}{"QueryCollection", value})
	case ParamsOfWaitForCollection:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfWaitForCollection
		}{"WaitForCollection", value})
	case ParamsOfAggregateCollection:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfAggregateCollection
		}{"AggregateCollection", value})
	case ParamsOfQueryCounterparties:
		return json.Marshal(struct {
			Type string `json:"type"`
			ParamsOfQueryCounterparties
		}{"QueryCounterparties", value})
	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfQueryOperation %v", pOQO.ValueEnumType)
	}
}

func (pOQO *ParamsOfQueryOperation) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "QueryCollection":
		var valueEnum ParamsOfQueryCollection
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOQO.ValueEnumType = valueEnum
	case "WaitForCollection":
		var valueEnum ParamsOfWaitForCollection
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOQO.ValueEnumType = valueEnum
	case "AggregateCollection":
		var valueEnum ParamsOfAggregateCollection
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOQO.ValueEnumType = valueEnum
	case "QueryCounterparties":
		var valueEnum ParamsOfQueryCounterparties
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pOQO.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ParamsOfQueryOperation %v", typeD.Type)
	}
	return nil
}

// NewParamsOfQueryOperation ...
func NewParamsOfQueryOperation(value interface{}) ParamsOfQueryOperation {
	return ParamsOfQueryOperation{ValueEnumType: value}
}