package domain

import "encoding/json"

const (

	// PEWillFetchFirstBlock ...
	PEWillFetchFirstBlock ProcessingEventType = "WillFetchFirstBlock"

	// PEWillFetchFirstBlockFailed ...
	PEWillFetchFirstBlockFailed ProcessingEventType = "FetchFirstBlockFailed"

	// PEWillSend ...
	PEWillSend ProcessingEventType = "WillSend"

	// PEDidSend ...
	PEDidSend ProcessingEventType = "DidSend"

	// PESendFailed ...
	PESendFailed ProcessingEventType = "SendFailed"

	// PEWillFetchNextBlock ...
	PEWillFetchNextBlock ProcessingEventType = "WillFetchNextBlock"

	// PEFetchNextBlockFailed ...
	PEFetchNextBlockFailed ProcessingEventType = "FetchNextBlockFailed"

	// PEMessageExpired ...
	PEMessageExpired ProcessingEventType = "MessageExpired"
)

type (
	// ProcessingEventType ...
	ProcessingEventType string

	// ProcessingEvent ...
	ProcessingEvent struct {
		Type         ProcessingEventType `json:"type"`
		Error        ClientError         `json:"error"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	//ProcessingEventWFFB - WillFetchFirstBlock
	ProcessingEventWFFB struct {
		Type ProcessingEventType `json:"type"`
	}

	//ProcessingEventFFBF - FetchFirstBlockFailed
	ProcessingEventFFBF struct {
		Type  ProcessingEventType `json:"type"`
		Error ClientError         `json:"error"`
	}

	//ProcessingEventWS - WillSend
	ProcessingEventWS struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	//ProcessingEventDS - DidSend
	ProcessingEventDS struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	//ProcessingEventSF - SendFailed
	ProcessingEventSF struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
		Error        ClientError         `json:"error"`
	}

	//ProcessingEventWFNB - WillFetchNextBlock
	ProcessingEventWFNB struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
	}

	//ProcessingEventFNBF - FetchNextBlockFailed
	ProcessingEventFNBF struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"message"`
		Error        ClientError         `json:"error"`
	}

	//ProcessingEventME - MessageExpired
	ProcessingEventME struct {
		Type      ProcessingEventType `json:"type"`
		MessageID string              `json:"message_id"`
		Message   string              `json:"message"`
		Error     ClientError         `json:"error"`
	}

	// ParamsOfSendMessage ...
	ParamsOfSendMessage struct {
		Message    string `json:"message"`
		Abi        Abi    `json:"abi,omitempty"`
		SendEvents bool   `json:"send_events"`
	}

	// ResultOfSendMessage ...
	ResultOfSendMessage struct {
		ShardBlockID string `json:"shard_block_id"`
	}

	// ParamsOfWaitForTransaction ...
	ParamsOfWaitForTransaction struct {
		Abi          Abi    `json:"abi,omitempty"`
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
		OutMessages []DecodedMessageBody `json:"out_messages,omitempty"`
		Output      json.RawMessage      `json:"output,omitempty"`
	}

	// ParamsOfProcessMessage ...
	ParamsOfProcessMessage struct {
		MessageEncodeParams *ParamsOfEncodeMessage `json:"message_encode_params"`
		SendEvents          bool                   `json:"send_events"`
	}

	// EventCallback ...
	//change!!!!
	EventCallback func(event *ProcessingEvent) //*ProcessingEvent ???

	// ProcessingUseCase ...
	ProcessingUseCase interface {
		SendMessage(ParamsOfSendMessage, EventCallback) (*ResultOfSendMessage, error)
		WaitForTransaction(ParamsOfWaitForTransaction, EventCallback) (*ResultOfProcessMessage, error)
		ProcessMessage(ParamsOfProcessMessage, EventCallback) (*ResultOfProcessMessage, error)
	}
)

// NewProcEventWFFB - ProcessingEvent, type: WillFetchFirstBlock
func NewProcEventWFFB() ProcessingEventWFFB {
	return ProcessingEventWFFB{Type: PEWillFetchFirstBlock}
}

// NewProcEventFFBF - ProcessingEvent, type: FetchFirstBlockFailed
func NewProcEventFFBF() ProcessingEventFFBF {
	return ProcessingEventFFBF{Type: PEWillFetchFirstBlockFailed}
}

// NewProcEventWS - ProcessingEvent, type: WillSend
func NewProcEventWS() ProcessingEventWS {
	return ProcessingEventWS{Type: PEWillSend}
}

// NewProcEventDS - ProcessingEvent, type: DidSend
func NewProcEventDS() ProcessingEventDS {
	return ProcessingEventDS{Type: PEDidSend}
}

// NewProcEventSF - ProcessingEvent, type: SendFailed
func NewProcEventSF() ProcessingEventSF {
	return ProcessingEventSF{Type: PESendFailed}
}

// NewProcEventWFNB - ProcessingEvent, type: WillFetchNextBlock
func NewProcEventWFNB() ProcessingEventWFNB {
	return ProcessingEventWFNB{Type: PEWillFetchNextBlock}
}

// NewProcEventFNBF - ProcessingEvent, type: FetchNextBlockFailed
func NewProcEventFNBF() ProcessingEventFNBF {
	return ProcessingEventFNBF{Type: PEFetchNextBlockFailed}
}

// NewProcEventME - ProcessingEvent, type: MessageExpired
func NewProcEventME() ProcessingEventME {
	return ProcessingEventME{Type: PEMessageExpired}
}
