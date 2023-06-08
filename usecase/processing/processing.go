package processing

import (
	"errors"

	"github.com/markgenuine/ever-client-go/domain"
)

type processing struct {
	config domain.ClientConfig
	client domain.ClientGateway
}

func NewProcessing(
	config domain.ClientConfig,
	client domain.ClientGateway,
) domain.ProcessingUseCase {
	return &processing{
		config: config,
		client: client,
	}
}

// MonitorMessages - Starts monitoring for the processing results of the specified messages.
//
// Message monitor performs background monitoring for a message processing results for the specified set of messages.
//
// Message monitor can serve several isolated monitoring queues. Each monitor queue has a unique application defined identifier (or name) used to separate several queue's.
//
// There are two important lists inside of the monitoring queue:
//
// unresolved messages: contains messages requested by the application for monitoring and not yet resolved;
//
// resolved results: contains resolved processing results for monitored messages.
//
// Each monitoring queue tracks own unresolved and resolved lists. Application can add more messages to the monitoring queue at any time.
//
// Message monitor accumulates resolved results. Application should fetch this results with fetchNextMonitorResults function.
//
// When both unresolved and resolved lists becomes empty, monitor stops any background activity and frees all allocated internal memory.
//
// If monitoring queue with specified name already exists then messages will be added to the unresolved list.
//
// If monitoring queue with specified name does not exist then monitoring queue will be created with specified unresolved messages.
func (p *processing) MonitorMessages(messages *domain.ParamsOfMonitorMessages) error {
	_, err := p.client.GetResponse("processing.monitor_messages", messages)
	return err
}

// GetMonitorInfo - Returns summary information about current state of the specified monitoring queue.
func (p *processing) GetMonitorInfo(info *domain.ParamsOfGetMonitorInfo) (*domain.MonitoringQueueInfo, error) {
	result := new(domain.MonitoringQueueInfo)
	err := p.client.GetResult("processing.get_monitor_info", info, result)
	return result, err
}

// FetchNextMonitorResults - Fetches next resolved results from the specified monitoring queue.
// Results and waiting options are depends on the wait parameter. All returned results will be removed from the queue's resolved list.
func (p *processing) FetchNextMonitorResults(results *domain.ParamsOfFetchNextMonitorResults) (*domain.ResultOfFetchNextMonitorResults, error) {
	result := new(domain.ResultOfFetchNextMonitorResults)
	err := p.client.GetResult("processing.fetch_next_monitor_results", results, result)
	return result, err
}

// CancelMonitor - Cancels all background activity and releases all allocated system resources for the specified monitoring queue.
func (p *processing) CancelMonitor(monitor *domain.ParamsOfCancelMonitor) error {
	_, err := p.client.GetResponse("processing.cancel_monitor", monitor)
	return err
}

// SendMessages - Sends specified messages to the blockchain.
func (p *processing) SendMessages(messages *domain.ParamsOfSendMessages) (*domain.ResultOfSendMessages, error) {
	result := new(domain.ResultOfSendMessages)
	err := p.client.GetResult("processing.send_messages", messages, result)
	return result, err
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
