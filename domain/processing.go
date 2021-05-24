package domain

import (
	"encoding/json"
	"fmt"
)

// ProcessingErrorCode ...
var ProcessingErrorCode map[string]int

type (
	// ProcessingEvent ...
	ProcessingEvent struct {
		ValueEnumType interface{}
	}

	// ProcessingEventWillFetchFirstBlock ...
	ProcessingEventWillFetchFirstBlock struct {}

	// ProcessingEventFetchFirstBlockFailed ...
	ProcessingEventFetchFirstBlockFailed struct {
		Error        ClientError        `json:"error"`
	}

	// ProcessingEventWillSend ...
	ProcessingEventWillSend struct {
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	// ProcessingEventDidSend ...
	ProcessingEventDidSend struct {
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	// ProcessingEventSendFailed ...
	ProcessingEventSendFailed struct {
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
		Error        ClientError        `json:"error"`
	}

	// ProcessingEventWillFetchNextBlock ...
	ProcessingEventWillFetchNextBlock struct {
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	// ProcessingEventFetchNextBlockFailed ...
	ProcessingEventFetchNextBlockFailed struct {
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
		Error        ClientError        `json:"error"`
	}

	// ProcessingEventMessageExpired ...
	ProcessingEventMessageExpired struct {
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
		Error        ClientError        `json:"error"`
	}

	// ParamsOfSendMessage ...
	ParamsOfSendMessage struct {
		Message    string `json:"message"`
		Abi        *Abi   `json:"abi,omitempty"`
		SendEvents bool   `json:"send_events"`
	}

	// ResultOfSendMessage ...
	ResultOfSendMessage struct {
		ShardBlockID string `json:"shard_block_id"`
		SendingEndpoints []string `json:"sending_endpoints"`
	}

	// ParamsOfWaitForTransaction ...
	ParamsOfWaitForTransaction struct {
		Abi          *Abi   `json:"abi,omitempty"`
		Message      string `json:"message"`
		ShardBlockID string `json:"shard_block_id"`
		SendEvents   bool   `json:"send_events"`
		SendingEndpoints []string `json:"sending_endpoints,omitempty"`
	}

	// ResultOfProcessMessage ...
	ResultOfProcessMessage struct {
		Transaction json.RawMessage  `json:"transaction"`
		OutMessages []string         `json:"out_messages"`
		Decoded     *DecodedOutput   `json:"decoded,omitempty"`
		Fees        *TransactionFees `json:"fees"`
	}

	// DecodedOutput ...
	DecodedOutput struct {
		OutMessages []*DecodedMessageBody `json:"out_messages,omitempty"`
		Output      json.RawMessage       `json:"output,omitempty"`
	}

	// ParamsOfProcessMessage ...
	ParamsOfProcessMessage struct {
		MessageEncodeParams *ParamsOfEncodeMessage `json:"message_encode_params"`
		SendEvents          bool                   `json:"send_events"`
	}

	// EventCallback
	EventCallback func(event *ProcessingEvent)

	// ProcessingUseCase ...
	ProcessingUseCase interface {
		SendMessage(*ParamsOfSendMessage, EventCallback) (*ResultOfSendMessage, error)
		WaitForTransaction(*ParamsOfWaitForTransaction, EventCallback) (*ResultOfProcessMessage, error)
		ProcessMessage(*ParamsOfProcessMessage, EventCallback) (*ResultOfProcessMessage, error)
	}
)

func init() {
	ProcessingErrorCode = map[string]int{
		"MessageAlreadyExpired          ": 501,
		"MessageHasNotDestinationAddress": 502,
		"CanNotBuildMessageCell         ": 503,
		"FetchBlockFailed               ": 504,
		"SendMessageFailed              ": 505,
		"InvalidMessageBoc              ": 506,
		"MessageExpired                 ": 507,
		"TransactionWaitTimeout         ": 508,
		"InvalidBlockReceived           ": 509,
		"CanNotCheckBlockShard          ": 510,
		"BlockNotFound                  ": 511,
		"InvalidData                    ": 512,
		"ExternalSignerMustNotBeUsed    ": 513,
	}
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
	default:
		return fmt.Errorf("unsupported type for ProcessingEvent %v", typeD.Type)
	}
	return nil
}