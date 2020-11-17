package processing

import (
	"github.com/markgenuine/ton-client-go/domain"
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

// SendMessage method processing.send_message
func (p *processing) Send(pOSM domain.ParamsOfSendMessage, ResponseHandler int) (int, error) { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return p.client.Request("processing.send_message", pOSM)
}

// WaitForTransaction method processing.wait_for_transaction
func (p *processing) WaitForTransaction(pOWFT domain.ParamsOfWaitForTransaction, ResponseHandler int) (int, error) { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return p.client.Request("processing.wait_for_transaction", pOWFT)
}

// ProcessMessage method processing.process_message
func (p *processing) Process(pOPM domain.ParamsOfProcessMessage, ResponseHandler int) (int, error) { ///responseHandler?: ResponseHandler – additional responses handler.### Result
	return p.client.Request("processing.process_message", pOPM)
}

// ConvertAddress ...
func (p *processing) ConvertAddress(pOCA domain.ParamsOfConvertAddress) (int, error) {
	return p.client.Request("utils.convert_address", pOCA)
}
