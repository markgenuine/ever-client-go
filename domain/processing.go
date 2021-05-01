package domain

import "encoding/json"

const (
	// WillFetchFirstBlock ...
	WillFetchFirstBlock ProcessingEventType = "WillFetchFirstBlock"

	// FetchFirstBlockFailed ...
	FetchFirstBlockFailed ProcessingEventType = "FetchFirstBlockFailed"

	// WillSend ...
	WillSend ProcessingEventType = "WillSend"

	// DidSend ...
	DidSend ProcessingEventType = "DidSend"

	// SendFailed ...
	SendFailed ProcessingEventType = "SendFailed"

	// WillFetchNextBlock ...
	WillFetchNextBlock ProcessingEventType = "WillFetchNextBlock"

	// FetchNextBlockFailed ...
	FetchNextBlockFailed ProcessingEventType = "FetchNextBlockFailed"

	// MessageExpired ...
	MessageExpired ProcessingEventType = "MessageExpired"
)

// ProcessingErrorCode ...
var ProcessingErrorCode map[string]int

type (
	// ProcessingEventType ...
	ProcessingEventType string

	// ProcessingEvent ...
	ProcessingEvent struct {
		Type         ProcessingEventType `json:"type"`
		Error        *ClientError        `json:"error,omitempty"`
		ShardBlockID string              `json:"shard_block_id,omitempty"`
		MessageID    string              `json:"message_id,omitempty"`
		Message      string              `json:"message,omitempty"`
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
	}

	// ParamsOfWaitForTransaction ...
	ParamsOfWaitForTransaction struct {
		Abi          *Abi   `json:"abi,omitempty"`
		Message      string `json:"message"`
		ShardBlockID string `json:"shard_block_id"`
		SendEvents   bool   `json:"send_events"`
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
		Output      json.RawMessage      `json:"output,omitempty"`
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

// ProcessingEventWillFetchFirstBlock variant constructor ProcessingEvent.
func ProcessingEventWillFetchFirstBlock() *ProcessingEvent {
	return &ProcessingEvent{Type: WillFetchFirstBlock}
}

// ProcessingEventFetchFirstBlockFailed variant constructor ProcessingEvent.
func ProcessingEventFetchFirstBlockFailed(err *ClientError) *ProcessingEvent {
	return &ProcessingEvent{Type: FetchFirstBlockFailed, Error: err}
}

// ProcessingEventWillSend variant constructor ProcessingEvent.
func ProcessingEventWillSend(shardBlockId, messageID, message string) *ProcessingEvent {
	return &ProcessingEvent{Type: WillSend, ShardBlockID: shardBlockId, MessageID: messageID, Message: message}
}

// ProcessingEventDidSend variant constructor ProcessingEvent.
func ProcessingEventDidSend(shardBlockId, messageID, message string) *ProcessingEvent {
	return &ProcessingEvent{Type: DidSend, ShardBlockID: shardBlockId, MessageID: messageID, Message: message}
}

// ProcessingEventSendFailed variant constructor ProcessingEvent.
func ProcessingEventSendFailed(shardBlockId, messageID, message string, err *ClientError) *ProcessingEvent {
	return &ProcessingEvent{Type: SendFailed, ShardBlockID: shardBlockId, MessageID: messageID, Message: message, Error: err}
}

// ProcessingEventWillFetchNextBlock variant constructor ProcessingEvent.
func ProcessingEventWillFetchNextBlock(shardBlockId, messageID, message string) *ProcessingEvent {
	return &ProcessingEvent{Type: WillFetchNextBlock, ShardBlockID: shardBlockId, MessageID: messageID, Message: message}
}

// ProcessingEventFetchNextBlockFailed variant constructor ProcessingEvent.
func ProcessingEventFetchNextBlockFailed(shardBlockId, messageID, message string, err *ClientError) *ProcessingEvent {
	return &ProcessingEvent{Type: FetchNextBlockFailed, ShardBlockID: shardBlockId, MessageID: messageID, Message: message, Error: err}
}

// ProcessingEventMessageExpired variant constructor ProcessingEvent.
func ProcessingEventMessageExpired(messageID, message string, err *ClientError) *ProcessingEvent {
	return &ProcessingEvent{Type: MessageExpired, MessageID: messageID, Message: message, Error: err}
}
