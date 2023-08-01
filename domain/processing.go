package domain

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const (
	AtLeastOne MonitorFetchWaitMode = "AtLeastOne"
	All        MonitorFetchWaitMode = "All"
	NoWait     MonitorFetchWaitMode = "NoWait"

	Finalized MessageMonitoringStatus = "Finalized"
	Timeout   MessageMonitoringStatus = "Timeout"
	Reserved  MessageMonitoringStatus = "Reserved"
)

var ProcessingErrorCode map[string]int

type (
	ProcessingEvent struct {
		ValueEnumType interface{}
	}

	ProcessingEventWillFetchFirstBlock struct {
		MessageID  string `json:"message_id"`
		MessageDst string `json:"message_dst"`
	}

	ProcessingEventFetchFirstBlockFailed struct {
		Error      ClientError `json:"error"`
		MessageID  string      `json:"message_id"`
		MessageDst string      `json:"message_dst"`
	}

	ProcessingEventWillSend struct {
		ShardBlockID string `json:"shard_block_id"`
		MessageID    string `json:"message_id"`
		MessageDst   string `json:"message_dst"`
		Message      string `json:"message"`
	}

	ProcessingEventDidSend struct {
		ShardBlockID string `json:"shard_block_id"`
		MessageID    string `json:"message_id"`
		MessageDst   string `json:"message_dst"`
		Message      string `json:"message"`
	}

	ProcessingEventSendFailed struct {
		ShardBlockID string      `json:"shard_block_id"`
		MessageID    string      `json:"message_id"`
		MessageDst   string      `json:"message_dst"`
		Message      string      `json:"message"`
		Error        ClientError `json:"error"`
	}

	ProcessingEventWillFetchNextBlock struct {
		ShardBlockID string `json:"shard_block_id"`
		MessageID    string `json:"message_id"`
		MessageDst   string `json:"message_dst"`
		Message      string `json:"message"`
	}

	ProcessingEventFetchNextBlockFailed struct {
		ShardBlockID string      `json:"shard_block_id"`
		MessageID    string      `json:"message_id"`
		MessageDst   string      `json:"message_dst"`
		Message      string      `json:"message"`
		Error        ClientError `json:"error"`
	}

	ProcessingEventMessageExpired struct {
		MessageID  string      `json:"message_id"`
		MessageDst string      `json:"message_dst"`
		Message    string      `json:"message"`
		Error      ClientError `json:"error"`
	}

	ProcessingRempSentToValidators struct {
		MessageID  string          `json:"message_id"`
		MessageDst string          `json:"message_dst"`
		TimeStamp  big.Int         `json:"timestamp"`
		JSON       json.RawMessage `json:"json"`
	}

	ProcessingRempIncludedIntoBlock struct {
		MessageID  string          `json:"message_id"`
		MessageDst string          `json:"message_dst"`
		Timestamp  big.Int         `json:"timestamp"`
		JSON       json.RawMessage `json:"json"`
	}

	ProcessingRempIncludedIntoAcceptedBlock struct {
		MessageID  string          `json:"message_id"`
		MessageDst string          `json:"message_dst"`
		Timestamp  big.Int         `json:"timestamp"`
		JSON       json.RawMessage `json:"json"`
	}

	ProcessingRempOther struct {
		MessageID  string          `json:"message_id"`
		MessageDst string          `json:"message_dst"`
		Timestamp  big.Int         `json:"timestamp"`
		JSON       json.RawMessage `json:"json"`
	}

	ProcessingRempError struct {
		MessageID  string      `json:"message_id"`
		MessageDst string      `json:"message_dst"`
		Error      ClientError `json:"error"`
	}

	ParamsOfMonitorMessages struct {
		Queue    string                     `json:"queue"`
		Messages []*MessageMonitoringParams `json:"messages"`
	}

	ParamsOfGetMonitorInfo struct {
		Queue string `json:"queue"`
	}

	MonitoringQueueInfo struct {
		Unresolved int `json:"unresolved"`
		Resolved   int `json:"resolved"`
	}

	ParamsOfFetchNextMonitorResults struct {
		Queue    string                `json:"queue"`
		WaitMode *MonitorFetchWaitMode `json:"wait_mode,omitempty"`
	}

	ResultOfFetchNextMonitorResults struct {
		Results []*MessageMonitoringResult `json:"results"`
	}

	ParamsOfCancelMonitor struct {
		Queue string `json:"queue"`
	}

	ParamsOfSendMessages struct {
		Messages     []*MessageSendingParams `json:"messages"`
		MonitorQueue string                  `json:"monitor_queue,omitempty"`
	}

	ResultOfSendMessages struct {
		Messages []*MessageMonitoringParams `json:"messages"`
	}

	ParamsOfSendMessage struct {
		Message    string `json:"message"`
		Abi        *Abi   `json:"abi,omitempty"`
		SendEvents bool   `json:"send_events"`
	}

	ResultOfSendMessage struct {
		ShardBlockID     string   `json:"shard_block_id"`
		SendingEndpoints []string `json:"sending_endpoints"`
	}

	ParamsOfWaitForTransaction struct {
		Abi              *Abi     `json:"abi,omitempty"`
		Message          string   `json:"message"`
		ShardBlockID     string   `json:"shard_block_id"`
		SendEvents       bool     `json:"send_events"`
		SendingEndpoints []string `json:"sending_endpoints,omitempty"`
	}

	ResultOfProcessMessage struct {
		Transaction json.RawMessage  `json:"transaction"`
		OutMessages []string         `json:"out_messages"`
		Decoded     *DecodedOutput   `json:"decoded,omitempty"`
		Fees        *TransactionFees `json:"fees"`
	}

	DecodedOutput struct {
		OutMessages []*DecodedMessageBody `json:"out_messages"`
		Output      json.RawMessage       `json:"output,omitempty"`
	}

	MessageMonitoringTransactionCompute struct {
		ExitCode int `json:"exit_code"`
	}

	MessageMonitoringTransaction struct {
		Hash    string                               `json:"hash,omitempty"`
		Aborted bool                                 `json:"aborted"`
		Compute *MessageMonitoringTransactionCompute `json:"compute"`
	}

	MessageMonitoringParams struct {
		Message   *MonitoredMessage `json:"message"`
		WaitUntil int               `json:"wait_until"`
		UserData  *json.RawMessage  `json:"user_data,omitempty"`
	}

	MessageMonitoringResult struct {
		Hash        string                        `json:"hash,omitempty"`
		Status      *MessageMonitoringStatus      `json:"status"`
		Transaction *MessageMonitoringTransaction `json:"transaction,omitempty"`
		Error       string                        `json:"error,omitempty"`
		UserData    *json.RawMessage              `json:"user_data,omitempty"`
	}

	MonitorFetchWaitMode string

	// MonitoredMessageBocVariant - BOC of the message.
	MonitoredMessageBocVariant struct {
		value struct {
			Boc string `json:"boc"`
		}
	}
	// MonitoredMessageHashAddressVariant - Message's hash and destination address.
	MonitoredMessageHashAddressVariant struct {
		value struct {
			Hash    string `json:"hash"`
			Address string `json:"address"`
		}
	}

	MonitoredMessage struct {
		ValueEnumType interface{}
	}

	MessageMonitoringStatus string

	MessageSendingParams struct {
		Boc       string           `json:"boc"`
		WaitUntil int              `json:"wait_until"`
		UserData  *json.RawMessage `json:"user_data,omitempty"`
	}

	ParamsOfProcessMessage struct {
		MessageEncodeParams *ParamsOfEncodeMessage `json:"message_encode_params"`
		SendEvents          bool                   `json:"send_events"`
	}

	EventCallback func(event *ProcessingEvent)

	ProcessingUseCase interface {
		MonitorMessages(*ParamsOfMonitorMessages) error
		GetMonitorInfo(*ParamsOfGetMonitorInfo) (*MonitoringQueueInfo, error)
		FetchNextMonitorResults(*ParamsOfFetchNextMonitorResults) (*ResultOfFetchNextMonitorResults, error)
		CancelMonitor(monitor *ParamsOfCancelMonitor) error
		SendMessages(messages *ParamsOfSendMessages) (*ResultOfSendMessages, error)
		SendMessage(*ParamsOfSendMessage, EventCallback) (*ResultOfSendMessage, error)
		WaitForTransaction(*ParamsOfWaitForTransaction, EventCallback) (*ResultOfProcessMessage, error)
		ProcessMessage(*ParamsOfProcessMessage, EventCallback) (*ResultOfProcessMessage, error)
	}
)

func init() {
	ProcessingErrorCode = map[string]int{
		"MessageAlreadyExpired":           501,
		"MessageHasNotDestinationAddress": 502,
		"CanNotBuildMessageCell":          503,
		"FetchBlockFailed":                504,
		"SendMessageFailed":               505,
		"InvalidMessageBoc":               506,
		"MessageExpired":                  507,
		"TransactionWaitTimeout":          508,
		"InvalidBlockReceived":            509,
		"CanNotCheckBlockShard":           510,
		"BlockNotFound":                   511,
		"InvalidData":                     512,
		"ExternalSignerMustNotBeUsed":     513,
		"MessageRejected":                 514,
		"InvalidRempStatus":               515,
		"NextRempStatusTimeout":           516,
	}
}

func (mm *MonitoredMessage) MarshalJSON() ([]byte, error) {
	switch value := (mm.ValueEnumType).(type) {
	case MonitoredMessageBocVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			MonitoredMessageBocVariant
		}{"Boc", value})
	case MonitoredMessageHashAddressVariant:
		return json.Marshal(struct {
			Type string `json:"type"`
			MonitoredMessageHashAddressVariant
		}{"HashAddress", value})
	default:
		return nil, fmt.Errorf("unsupported type for MonitoredMessage %v", mm.ValueEnumType)
	}
}

func (mm *MonitoredMessage) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Boc":
		var valueEnum MonitoredMessageBocVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		mm.ValueEnumType = valueEnum
	case "HashAddress":
		var valueEnum MonitoredMessageHashAddressVariant
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		mm.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for MonitoredMessage %v", typeD.Type)
	}
	return nil

}

func (pE *ProcessingEvent) MarshalJSON() ([]byte, error) {
	switch value := (pE.ValueEnumType).(type) {
	case ProcessingEventWillFetchFirstBlock:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventWillFetchFirstBlock
		}{"WillFetchFirstBlock", value})
	case ProcessingEventFetchFirstBlockFailed:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventFetchFirstBlockFailed
		}{"FetchFirstBlockFailed", value})
	case ProcessingEventWillSend:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventWillSend
		}{"WillSend", value})
	case ProcessingEventDidSend:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventDidSend
		}{"DidSend", value})
	case ProcessingEventSendFailed:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventSendFailed
		}{"SendFailed", value})
	case ProcessingEventWillFetchNextBlock:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventWillFetchNextBlock
		}{"WillFetchNextBlock", value})
	case ProcessingEventFetchNextBlockFailed:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventFetchNextBlockFailed
		}{"FetchNextBlockFailed", value})
	case ProcessingEventMessageExpired:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingEventMessageExpired
		}{"MessageExpired", value})
	case ProcessingRempSentToValidators:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingRempSentToValidators
		}{"RempSentToValidators", value})
	case ProcessingRempIncludedIntoBlock:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingRempIncludedIntoBlock
		}{"RempIncludedIntoBlock", value})
	case ProcessingRempIncludedIntoAcceptedBlock:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingRempIncludedIntoAcceptedBlock
		}{"RempIncludedIntoAcceptedBlock", value})
	case ProcessingRempOther:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingRempOther
		}{"RempOther", value})
	case ProcessingRempError:
		return json.Marshal(struct {
			Type string `json:"type"`
			ProcessingRempError
		}{"RempError", value})
	default:
		return nil, fmt.Errorf("unsupported type for ProcessingEvent %v", pE.ValueEnumType)
	}
}

func (pE *ProcessingEvent) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "WillFetchFirstBlock":
		var valueEnum ProcessingEventWillFetchFirstBlock
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "FetchFirstBlockFailed":
		var valueEnum ProcessingEventFetchFirstBlockFailed
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "WillSend":
		var valueEnum ProcessingEventWillSend
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "DidSend":
		var valueEnum ProcessingEventDidSend
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "SendFailed":
		var valueEnum ProcessingEventSendFailed
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "WillFetchNextBlock":
		var valueEnum ProcessingEventWillFetchNextBlock
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "FetchNextBlockFailed":
		var valueEnum ProcessingEventFetchNextBlockFailed
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "MessageExpired":
		var valueEnum ProcessingEventMessageExpired
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "RempSentToValidators":
		var valueEnum ProcessingRempSentToValidators
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "RempIncludedIntoBlock":
		var valueEnum ProcessingRempIncludedIntoBlock
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "RempIncludedIntoAcceptedBlock":
		var valueEnum ProcessingRempIncludedIntoAcceptedBlock
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "RempOther":
		var valueEnum ProcessingRempOther
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	case "RempError":
		var valueEnum ProcessingRempError
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		pE.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for ProcessingEvent %v", typeD.Type)
	}
	return nil
}
