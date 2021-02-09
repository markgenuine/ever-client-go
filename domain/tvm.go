package domain

import (
	"encoding/json"
	"math/big"
)

const (
	// AccountForExecutorTypeNone ...
	AccountForExecutorTypeNone = AccountForExecutorType("None")

	// AccountForExecutorTypeUninit ...
	AccountForExecutorTypeUninit = AccountForExecutorType("Uninit")

	// AccountForExecutorTypeAccount ...
	AccountForExecutorTypeAccount = AccountForExecutorType("Account")
)

// TVMErrorCode ...
var TVMErrorCode map[string]int

type (
	// ExecutionOptions ...
	ExecutionOptions struct {
		BlockchainConfig string   `json:"blockchain_config,omitempty"`
		BlockTime        int      `json:"block_time,omitempty"`
		BlockLt          *big.Int `json:"block_lt,omitempty"`
		TransactionLt    *big.Int `json:"transaction_lt,omitempty"`
	}

	// AccountForExecutorType ...
	AccountForExecutorType string

	// AccountForExecutorNone ...
	AccountForExecutorNone struct {
		Type AccountForExecutorType `json:"type"`
	}

	// AccountForExecutorUninit ...
	AccountForExecutorUninit struct {
		Type AccountForExecutorType `json:"type"`
	}

	// AccountForExecutorAccount ...
	AccountForExecutorAccount struct {
		Type             AccountForExecutorType `json:"type"`
		Boc              string                 `json:"boc,omitempty"`
		UnlimitedBalance bool                   `json:"unlimited_balance,omitempty"`
	}

	// ParamsOfRunExecutor ...
	ParamsOfRunExecutor struct {
		Message              string            `json:"message"`
		Account              interface{}       `json:"account"` // AccountForExecutor
		ExecutionOptions     *ExecutionOptions `json:"execution_options,omitempty"`
		Abi                  Abi               `json:"abi,omitempty"`
		SkipTransactionCheck bool              `json:"skip_transaction_check,omitempty"`
	}

	// ResultOfRunExecuteMessage ...
	ResultOfRunExecuteMessage struct {
		Transaction json.RawMessage `json:"transaction,omitempty"`
		OutMessages []string        `json:"out_messages"`
		Decoded     DecodedOutput   `json:"decoded,omitempty"`
		Account     string          `json:"account"`
		Fees        TransactionFees `json:"fees"`
	}

	// ParamsOfRunTvm ...
	ParamsOfRunTvm struct {
		Message          string            `json:"message"`
		Account          string            `json:"account"`
		ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
		Abi              Abi               `json:"abi,omitempty"`
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
		FunctionName     string            `json:"function_name,omitempty"`
		Input            interface{}       `json:"input,omitempty"`
		ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
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
