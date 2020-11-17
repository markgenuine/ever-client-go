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

	//ProcessingEventWFFB - WillFetchFirstBlock
	ProcessingEventWFFB struct {
		ProcessingEventType `json:"type"`
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
		Message      string              `json:"Message"`
	}

	//ProcessingEventDS - DidSend
	ProcessingEventDS struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"Message"`
	}

	//ProcessingEventSF - SendFailed
	ProcessingEventSF struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"Message"`
		Error        ClientError         `json:"error"`
	}

	//ProcessingEventWFNB - WillFetchNextBlock
	ProcessingEventWFNB struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"Message"`
	}

	//ProcessingEventFNBF - FetchNextBlockFailed
	ProcessingEventFNBF struct {
		Type         ProcessingEventType `json:"type"`
		ShardBlockID string              `json:"shard_block_id"`
		MessageID    string              `json:"message_id"`
		Message      string              `json:"Message"`
		Error        ClientError         `json:"error"`
	}

	//ProcessingEventME - MessageExpired
	ProcessingEventME struct {
		Type      ProcessingEventType `json:"type"`
		MessageID string              `json:"message_id"`
		Message   string              `json:"Message"`
		Error     ClientError         `json:"error"`
	}

	// ParamsOfSendMessage ...
	ParamsOfSendMessage struct {
		Message    string      `json:"message"`
		Abi        interface{} `json:"abi,omitempty"` //ABI??? AbiS and AbiH
		SendEvents bool        `json:"send_events"`
	}

	// ResultOfSendMessage ...
	ResultOfSendMessage struct {
		ShardBlockID string `json:"shard_block_id"`
	}

	// ParamsOfWaitForTransaction ...
	ParamsOfWaitForTransaction struct {
		Abi          interface{} `json:"abi,omitempty"` //ABI??? AbiS and AbiH
		Message      string      `json:"message"`
		ShardBlockID string      `json:"shard_block_id"`
		SendEvents   bool        `json:"send_events"`
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

	// ProcessingUseCase ...
	ProcessingUseCase interface {
		Send(ParamsOfSendMessage, int) (int, error)
		WaitForTransaction(ParamsOfWaitForTransaction, int) (int, error)
		Process(ParamsOfProcessMessage, int) (int, error)
		ConvertAddress(ParamsOfConvertAddress) (int, error)
	}
)
