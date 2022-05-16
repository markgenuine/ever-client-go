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

	// AggregationFnTypeCount - Count value.
	AggregationFnTypeCount AggregationFnType = "COUNT"

	// AggregationFnTypeMin - Min value.
	AggregationFnTypeMin AggregationFnType = "MIN"

	// AggregationFnTypeMax - MAX value.
	AggregationFnTypeMax AggregationFnType = "MAX"

	// AggregationFnTypeSum - SUM value.
	AggregationFnTypeSum AggregationFnType = "SUM"

	// AggregationFnTypeAverage - AVERAGE value.
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

	// ParamsOfQueryOperation ...
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

	// TransactionNode ...
	TransactionNode struct {
		ID          string   `json:"id"`
		InMsg       string   `json:"in_msg"`
		OutMsgs     []string `json:"out_msgs"`
		AccountAddr string   `json:"account_addr"`
		TotalFees   string   `json:"total_fees"`
		Aborted     bool     `json:"aborted"`
		ExitCode    int      `json:"exit_code,omitempty"`
	}

	// MessageNode ...
	MessageNode struct {
		ID               string              `json:"id"`
		SrcTransactionID string              `json:"src_transaction_id,omitempty"`
		DstTransactionID string              `json:"dst_transaction_id,omitempty"`
		Src              string              `json:"src,omitempty"`
		Dst              string              `json:"dst,omitempty"`
		Value            string              `json:"value,omitempty"`
		Bounce           bool                `json:"bounce"`
		DecodedBody      *DecodedMessageBody `json:"decoded_body,omitempty"`
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
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
		Order      []*OrderBy      `json:"order,omitempty"`
		Limit      *int            `json:"limit,omitempty"`
	}

	// ResultOfQueryCollection ...
	ResultOfQueryCollection struct {
		Result []json.RawMessage `json:"result"`
	}

	// ParamsOfAggregateCollection ...
	ParamsOfAggregateCollection struct {
		Collection string              `json:"collection"`
		Filter     json.RawMessage     `json:"filter,omitempty"`
		Fields     []*FieldAggregation `json:"fields,omitempty"`
	}
	// ResultOfAggregateCollection ...
	ResultOfAggregateCollection struct {
		Values json.RawMessage `json:"values"`
	}

	// ParamsOfWaitForCollection ...
	ParamsOfWaitForCollection struct {
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
		Timeout    *int            `json:"timeout,omitempty"`
	}

	// ResultOfWaitForCollection ...
	ResultOfWaitForCollection struct {
		Result json.RawMessage `json:"result"`
	}

	// ParamsOfSubscribeCollection ...
	ParamsOfSubscribeCollection struct {
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
	}

	// ResultOfSubscribeCollection ...
	ResultOfSubscribeCollection struct {
		Handle int `json:"handle"`
	}

	// ParamsOfSubscribe ...
	ParamsOfSubscribe struct {
		Subscription string          `json:"subscription"`
		Variables    json.RawMessage `json:"variables,omitempty"`
	}

	// ParamsOfFindLastShardBlock ...
	ParamsOfFindLastShardBlock struct {
		Address string `json:"address"`
	}

	// ResultOfFindLastShardBlock ...
	ResultOfFindLastShardBlock struct {
		BlockID string `json:"block_id"`
	}

	// EndpointsSet ...
	EndpointsSet struct {
		Endpoints []string `json:"endpoints"`
	}

	// ResultOfGetEndpoints ...
	ResultOfGetEndpoints struct {
		Query     string   `json:"query"`
		Endpoints []string `json:"endpoints"`
	}

	// ParamsOfQueryCounterparties ...
	ParamsOfQueryCounterparties struct {
		Account string `json:"account"`
		Result  string `json:"result"`
		First   *int   `json:"first,omitempty"`
		After   string `json:"after,omitempty"`
	}

	// ParamsOfQueryTransactionTree ...
	ParamsOfQueryTransactionTree struct {
		InMsg       string `json:"in_msg"`
		AbiRegistry []*Abi `json:"abi_registry,omitempty"`
		TimeOut     *int   `json:"timeout,omitempty"`
	}

	// ResultOfQueryTransactionTree ...
	ResultOfQueryTransactionTree struct {
		Messages     []MessageNode     `json:"messages"`
		Transactions []TransactionNode `json:"transactions"`
	}

	// ParamsOfCreateBlockIterator ...
	ParamsOfCreateBlockIterator struct {
		StartTime   *int     `json:"start_time,omitempty"`
		EndTime     *int     `json:"end_time,omitempty"`
		ShardFilter []string `json:"shard_filter,omitempty"`
		Result      string   `json:"result,omitempty"`
	}

	// RegisteredIterator ...
	RegisteredIterator struct {
		Handle int `json:"handle"`
	}

	// ParamsOfResumeBlockIterator ...
	ParamsOfResumeBlockIterator struct {
		ResumeState json.RawMessage `json:"resume_state"`
	}

	// ParamsOfCreateTransactionIterator ...
	ParamsOfCreateTransactionIterator struct {
		StartTime        *int     `json:"start_time,omitempty"`
		EndTime          *int     `json:"end_time,omitempty"`
		ShardFilter      []string `json:"shard_filter,omitempty"`
		AccountsFilter   []string `json:"accounts_filter,omitempty"`
		Result           string   `json:"result,omitempty"`
		IncludeTransfers *bool    `json:"include_transfers,omitempty"`
	}

	// ParamsOfResumeTransactionIterator ...
	ParamsOfResumeTransactionIterator struct {
		ResumeState    json.RawMessage `json:"resume_state"`
		AccountsFilter []string        `json:"accounts_filter,omitempty"`
	}

	// ParamsOfIteratorNext ...
	ParamsOfIteratorNext struct {
		Iterator          int   `json:"iterator"`
		Limit             *int  `json:"limit,omitempty"`
		ReturnResumeState *bool `json:"return_resume_state,omitempty"`
	}

	// ResultOfIteratorNext ...
	ResultOfIteratorNext struct {
		Items       []json.RawMessage `json:"items"`
		HasMore     bool              `json:"has_more"`
		ResumeState json.RawMessage   `json:"resume_state,omitempty"`
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
		Subscribe(*ParamsOfSubscribe) (<-chan json.RawMessage, *ResultOfSubscribeCollection, error)
		Suspend() error
		Resume() error
		FindLastShardBlock(*ParamsOfFindLastShardBlock) (*ResultOfFindLastShardBlock, error)
		FetchEndpoints() (*EndpointsSet, error)
		SetEndpoints(*EndpointsSet) error
		GetEndpoints() (*ResultOfGetEndpoints, error)
		QueryCounterparties(*ParamsOfQueryCounterparties) (*ResultOfQueryCollection, error)
		QueryTransactionTree(*ParamsOfQueryTransactionTree) (*ResultOfQueryTransactionTree, error)
		CreateBlockIterator(*ParamsOfCreateBlockIterator) (*RegisteredIterator, error)
		ResumeBlockIterator(*ParamsOfResumeBlockIterator) (*RegisteredIterator, error)
		CreateTransactionIterator(*ParamsOfCreateTransactionIterator) (*RegisteredIterator, error)
		ResumeTransactionIterator(*ParamsOfResumeTransactionIterator) (*RegisteredIterator, error)
		IteratorNext(*ParamsOfIteratorNext) (*ResultOfIteratorNext, error)
		RemoveIterator(*RegisteredIterator) error
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

func (pOQO *ParamsOfQueryOperation) MarshalJSON() ([]byte, error) {
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
