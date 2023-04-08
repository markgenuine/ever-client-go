package domain

import (
	"encoding/json"
	"fmt"
)

const (
	SortDirectionASC  SortDirection = "ASC"
	SortDirectionDESC SortDirection = "DESC"

	AggregationFnTypeCount   AggregationFnType = "COUNT"
	AggregationFnTypeMin     AggregationFnType = "MIN"
	AggregationFnTypeMax     AggregationFnType = "MAX"
	AggregationFnTypeSum     AggregationFnType = "SUM"
	AggregationFnTypeAverage AggregationFnType = "AVERAGE"
)

var NetErrorCode map[string]int

type (
	SortDirection string

	OrderBy struct {
		Path      string        `json:"path"`
		Direction SortDirection `json:"direction"`
	}

	ParamsOfQueryOperation struct {
		ValueEnumType interface{}
	}

	AggregationFnType string

	FieldAggregation struct {
		Field string            `json:"field"`
		Fn    AggregationFnType `json:"fn"`
	}

	TransactionNode struct {
		ID          string   `json:"id"`
		InMsg       string   `json:"in_msg"`
		OutMsgs     []string `json:"out_msgs"`
		AccountAddr string   `json:"account_addr"`
		TotalFees   string   `json:"total_fees"`
		Aborted     bool     `json:"aborted"`
		ExitCode    int      `json:"exit_code,omitempty"`
	}

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

	ParamsOfQuery struct {
		Query     string          `json:"query"`
		Variables json.RawMessage `json:"variables,omitempty"`
	}

	ResultOfQuery struct {
		Result json.RawMessage `json:"result"`
	}

	ParamsOfBatchQuery struct {
		Operations []ParamsOfQueryOperation `json:"operations"`
	}

	ResultOfBatchQuery struct {
		Result []json.RawMessage `json:"result"`
	}

	ParamsOfQueryCollection struct {
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
		Order      []*OrderBy      `json:"order,omitempty"`
		Limit      *int            `json:"limit,omitempty"`
	}

	ResultOfQueryCollection struct {
		Result []json.RawMessage `json:"result"`
	}

	ParamsOfAggregateCollection struct {
		Collection string              `json:"collection"`
		Filter     json.RawMessage     `json:"filter,omitempty"`
		Fields     []*FieldAggregation `json:"fields,omitempty"`
	}

	ResultOfAggregateCollection struct {
		Values json.RawMessage `json:"values"`
	}

	ParamsOfWaitForCollection struct {
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
		Timeout    *int            `json:"timeout,omitempty"`
	}

	ResultOfWaitForCollection struct {
		Result json.RawMessage `json:"result"`
	}

	ParamsOfSubscribeCollection struct {
		Collection string          `json:"collection"`
		Filter     json.RawMessage `json:"filter,omitempty"`
		Result     string          `json:"result"`
	}

	ResultOfSubscribeCollection struct {
		Handle int `json:"handle"`
	}

	ParamsOfSubscribe struct {
		Subscription string          `json:"subscription"`
		Variables    json.RawMessage `json:"variables,omitempty"`
	}

	ParamsOfFindLastShardBlock struct {
		Address string `json:"address"`
	}

	ResultOfFindLastShardBlock struct {
		BlockID string `json:"block_id"`
	}

	EndpointsSet struct {
		Endpoints []string `json:"endpoints"`
	}

	ResultOfGetEndpoints struct {
		Query     string   `json:"query"`
		Endpoints []string `json:"endpoints"`
	}

	ParamsOfQueryCounterparties struct {
		Account string `json:"account"`
		Result  string `json:"result"`
		First   *int   `json:"first,omitempty"`
		After   string `json:"after,omitempty"`
	}

	ParamsOfQueryTransactionTree struct {
		InMsg               string `json:"in_msg"`
		AbiRegistry         []*Abi `json:"abi_registry,omitempty"`
		TimeOut             *int   `json:"timeout,omitempty"`
		TransactionMaxCount *int   `json:"transaction_max_count, omitempty"`
	}

	ResultOfQueryTransactionTree struct {
		Messages     []MessageNode     `json:"messages"`
		Transactions []TransactionNode `json:"transactions"`
	}

	ParamsOfCreateBlockIterator struct {
		StartTime   *int     `json:"start_time,omitempty"`
		EndTime     *int     `json:"end_time,omitempty"`
		ShardFilter []string `json:"shard_filter,omitempty"`
		Result      string   `json:"result,omitempty"`
	}

	RegisteredIterator struct {
		Handle int `json:"handle"`
	}

	ParamsOfResumeBlockIterator struct {
		ResumeState json.RawMessage `json:"resume_state"`
	}

	ParamsOfCreateTransactionIterator struct {
		StartTime        *int     `json:"start_time,omitempty"`
		EndTime          *int     `json:"end_time,omitempty"`
		ShardFilter      []string `json:"shard_filter,omitempty"`
		AccountsFilter   []string `json:"accounts_filter,omitempty"`
		Result           string   `json:"result,omitempty"`
		IncludeTransfers *bool    `json:"include_transfers,omitempty"`
	}

	ParamsOfResumeTransactionIterator struct {
		ResumeState    json.RawMessage `json:"resume_state"`
		AccountsFilter []string        `json:"accounts_filter,omitempty"`
	}

	ParamsOfIteratorNext struct {
		Iterator          int   `json:"iterator"`
		Limit             *int  `json:"limit,omitempty"`
		ReturnResumeState *bool `json:"return_resume_state,omitempty"`
	}

	ResultOfIteratorNext struct {
		Items       []json.RawMessage `json:"items"`
		HasMore     bool              `json:"has_more"`
		ResumeState json.RawMessage   `json:"resume_state,omitempty"`
	}

	ResultOfGetSignatureId struct {
		SignatureID *int `json:"signature_id,omitempty"`
	}

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
		GetSignatureID() (*ResultOfGetSignatureId, error)
	}
)

func init() {
	NetErrorCode = map[string]int{
		"QueryFailed":                    601,
		"SubscribeFailed":                602,
		"WaitForFailed":                  603,
		"GetSubscriptionResultFailed":    604,
		"InvalidServerResponse":          605,
		"ClockOutOfSync":                 606,
		"WaitForTimeout":                 607,
		"GraphqlError":                   608,
		"NetworkModuleSuspended":         609,
		"WebsocketDisconnected":          610,
		"NotSupported":                   611,
		"NoEndpointsProvided":            612,
		"GraphqlWebsocketInitError":      613,
		"NetworkModuleResumed":           614,
		"Unauthorized":                   615,
		"QueryTransactionTreeTimeout":    616,
		"GraphqlConnectionError":         617,
		"WrongWebsocketProtocolSequence": 618,
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

func NewParamsOfQueryOperation(value interface{}) ParamsOfQueryOperation {
	return ParamsOfQueryOperation{ValueEnumType: value}
}
