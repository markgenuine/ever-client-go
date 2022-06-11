package domain

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// TVMErrorCode ...
var TVMErrorCode map[string]int

type (
	// ExecutionOptions ...
	ExecutionOptions struct {
		BlockchainConfig string   `json:"blockchain_config,omitempty"`
		BlockTime        *int     `json:"block_time,omitempty"`
		BlockLt          *big.Int `json:"block_lt,omitempty"`
		TransactionLt    *big.Int `json:"transaction_lt,omitempty"`
	}

	// AccountForExecutor ...
	AccountForExecutor struct {
		ValueEnumType interface{}
	}

	// AccountForExecutorNone ...
	AccountForExecutorNone struct{}

	// AccountForExecutorUninit ...
	AccountForExecutorUninit struct{}

	// AccountForExecutorAccount ...
	AccountForExecutorAccount struct {
		Boc              string `json:"boc"`
		UnlimitedBalance *bool  `json:"unlimited_balance,omitempty"`
	}

	// ParamsOfRunExecutor ...
	ParamsOfRunExecutor struct {
		Message              string             `json:"message"`
		Account              AccountForExecutor `json:"account"`
		ExecutionOptions     *ExecutionOptions  `json:"execution_options,omitempty"`
		Abi                  *Abi               `json:"abi,omitempty"`
		SkipTransactionCheck *bool              `json:"skip_transaction_check,omitempty"`
		BocCache             *BocCacheType      `json:"boc_cache,omitempty"`
		ReturnUpdatedAccount *bool              `json:"return_updated_account,omitempty"`
	}

	// ResultOfRunExecuteMessage ...
	ResultOfRunExecuteMessage struct {
		Transaction json.RawMessage  `json:"transaction,omitempty"`
		OutMessages []string         `json:"out_messages"`
		Decoded     *DecodedOutput   `json:"decoded,omitempty"`
		Account     string           `json:"account"`
		Fees        *TransactionFees `json:"fees"`
	}

	// ParamsOfRunTvm ...
	ParamsOfRunTvm struct {
		Message              string            `json:"message"`
		Account              string            `json:"account"`
		ExecutionOptions     *ExecutionOptions `json:"execution_options,omitempty"`
		Abi                  *Abi              `json:"abi,omitempty"`
		BocCache             *BocCacheType     `json:"boc_cache,omitempty"`
		ReturnUpdatedAccount *bool             `json:"return_updated_account,omitempty"`
	}

	// ResultOfRunTvm ...
	ResultOfRunTvm struct {
		OutMessages []string       `json:"out_messages"`
		Decoded     *DecodedOutput `json:"decoded,omitempty"`
		Account     string         `json:"account"`
	}

	// ParamsOfRunGet ...
	ParamsOfRunGet struct {
		Account          string            `json:"account"`
		FunctionName     string            `json:"function_name"`
		Input            interface{}       `json:"input,omitempty"`
		ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
		TupleListAsArray *bool             `json:"tuple_list_as_array,omitempty"`
	}

	// ResultOfRunGet ...
	ResultOfRunGet struct {
		Output json.RawMessage `json:"output"`
	}

	// TransactionFees ...
	TransactionFees struct {
		InMsgFwdFee      *big.Int `json:"in_msg_fwd_fee"`
		StorageFee       *big.Int `json:"storage_fee"`
		GasFee           *big.Int `json:"gas_fee"`
		OutMsgsFwdFee    *big.Int `json:"out_msgs_fwd_fee"`
		TotalAccountFees *big.Int `json:"total_account_fees"`
		TotalOutput      *big.Int `json:"total_output"`
		ExtInMsgFee      *big.Int `json:"ext_in_msg_fee"`
		TotalFwdFees     *big.Int `json:"total_fwd_fees"`
		AccountFees      *big.Int `json:"account_fees"`
	}

	// TvmUseCase ...
	TvmUseCase interface {
		RunExecutor(*ParamsOfRunExecutor) (*ResultOfRunExecuteMessage, error)
		RunTvm(*ParamsOfRunTvm) (*ResultOfRunTvm, error)
		RunGet(*ParamsOfRunGet) (*ResultOfRunGet, error)
	}
)

func init() {
	TVMErrorCode = map[string]int{
		"CanNotReadTransaction      ": 401,
		"CanNotReadBlockchainConfig ": 402,
		"TransactionAborted         ": 403,
		"InternalError              ": 404,
		"ActionPhaseFailed          ": 405,
		"AccountCodeMissing         ": 406,
		"LowBalance                 ": 407,
		"AccountFrozenOrDeleted     ": 408,
		"AccountMissing             ": 409,
		"UnknownExecutionError      ": 410,
		"InvalidInputStack          ": 411,
		"InvalidAccountBoc          ": 412,
		"InvalidMessageType         ": 413,
		"ContractExecutionError     ": 414,
	}
}

func (aFE *AccountForExecutor) MarshalJSON() ([]byte, error) {
	switch value := (aFE.ValueEnumType).(type) {
	case AccountForExecutorNone:
		return json.Marshal(struct {
			Type string `json:"type"`
			AccountForExecutorNone
		}{"None", value})
	case AccountForExecutorUninit:
		return json.Marshal(struct {
			Type string `json:"type"`
			AccountForExecutorUninit
		}{"Uninit", value})
	case AccountForExecutorAccount:
		return json.Marshal(struct {
			Type string `json:"type"`
			AccountForExecutorAccount
		}{"Account", value})
	default:
		return nil, fmt.Errorf("unsupported type for AccountForExecutor %v", aFE.ValueEnumType)
	}
}

func (aFE *AccountForExecutor) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "None":
		var valueEnum AccountForExecutorNone
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		aFE.ValueEnumType = valueEnum
	case "Uninit":
		var valueEnum AccountForExecutorUninit
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		aFE.ValueEnumType = valueEnum
	case "Account":
		var valueEnum AccountForExecutorAccount
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		aFE.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for AccountForExecutor %v", typeD.Type)
	}
	return nil
}
