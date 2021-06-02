package processing

import (
	"errors"

	"github.com/move-ton/ton-client-go/domain"
)

type processing struct {
	config domain.Config
	client domain.ClientGateway
}

// NewProcessing ...
func NewProcessing(
	config domain.Config,
	client domain.ClientGateway,
) domain.ProcessingUseCase {
	return &processing{
		config: config,
		client: client,
	}
}

// SendMessage - Sends message to the network.
// Sends message to the network and returns the last
// generated shard block of the destination account before
// the message was sent. It will be required later for message processing.
func (p *processing) SendMessage(pOSM *domain.ParamsOfSendMessage, callback domain.EventCallback) (*domain.ResultOfSendMessage, error) {
	if pOSM.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.send_message", pOSM)
	if err != nil {
		return nil, err
	}

	if pOSM.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfSendMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}

// WaitForTransaction - Performs monitoring of the network for the result transaction of the external inbound message processing.
// send_events enables intermediate events, such as WillFetchNextBlock, FetchNextBlockFailed
// that may be useful for logging of new shard blocks creation during message processing.
// Note, that presence of the abi parameter is critical for ABI compliant contracts.
// Message processing uses drastically different strategy for processing message for contracts which ABI includes "expire" header.
func (p *processing) WaitForTransaction(pOWFT *domain.ParamsOfWaitForTransaction, callback domain.EventCallback) (*domain.ResultOfProcessMessage, error) {
	if pOWFT.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.wait_for_transaction", pOWFT)
	if err != nil {
		return nil, err
	}

	if pOWFT.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfProcessMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}

// ProcessMessage - Creates message, sends it to the network and monitors its processing.
// Creates ABI-compatible message, sends it to the network and monitors for the result
// transaction. Decodes the output messages' bodies.
// If contract's ABI includes "expire" header, then SDK implements retries in case of
// unsuccessful message delivery within the expiration timeout: SDK recreates the message, sends it and processes it again.
// The intermediate events, such as WillFetchFirstBlock, WillSend, DidSend, WillFetchNextBlock, etc -
// are switched on/off by send_events flag and logged into the supplied callback function.
// The retry configuration parameters are defined in the client's NetworkConfig and AbiConfig.
// If contract's ABI does not include "expire" header then, if no transaction is found within
// the network timeout (see config parameter ), exits with error.
func (p *processing) ProcessMessage(pOPM *domain.ParamsOfProcessMessage, callback domain.EventCallback) (*domain.ResultOfProcessMessage, error) {
	if pOPM.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.process_message", pOPM)
	if err != nil {
		return nil, err
	}

	if pOPM.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfProcessMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}
