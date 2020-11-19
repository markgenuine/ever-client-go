package processing

import (
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

// SendMessage method processing.send_message
func (p *processing) SendMessage(pOSM domain.ParamsOfSendMessage) (*domain.ResultOfSendMessage, error) {
	result := new(domain.ResultOfSendMessage)
	err := p.client.GetResult("processing.send_message", pOSM, result)
	return result, err

}

// WaitForTransaction method processing.wait_for_transaction
func (p *processing) WaitForTransaction(pOWFT domain.ParamsOfWaitForTransaction) (*domain.ResultOfProcessMessage, error) {
	result := new(domain.ResultOfProcessMessage)
	err := p.client.GetResult("processing.wait_for_transaction", pOWFT, result)
	return result, err
}

// ProcessMessage method processing.process_message
func (p *processing) ProcessMessage(pOPM domain.ParamsOfProcessMessage) (*domain.ResultOfProcessMessage, error) {
	result := new(domain.ResultOfProcessMessage)
	err := p.client.GetResult("processing.process_message", pOPM, result)
	return result, err
}
