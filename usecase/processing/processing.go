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
func (p *processing) SendMessage(pOSM domain.ParamsOfSendMessage) (int, error) {
	return p.client.Request("processing.send_message", pOSM)
}

// WaitForTransaction method processing.wait_for_transaction
func (p *processing) WaitForTransaction(pOWFT domain.ParamsOfWaitForTransaction) (int, error) {
	return p.client.Request("processing.wait_for_transaction", pOWFT)
}

// ProcessMessage method processing.process_message
func (p *processing) ProcessMessage(pOPM domain.ParamsOfProcessMessage) (int, error) {
	return p.client.Request("processing.process_message", pOPM)
}

// ConvertAddress ...
func (p *processing) ConvertAddress(pOCA domain.ParamsOfConvertAddress) (int, error) {
	return p.client.Request("utils.convert_address", pOCA)
}
