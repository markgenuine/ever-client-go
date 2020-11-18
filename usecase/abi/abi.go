package abi

import (
	"github.com/move-ton/ton-client-go/domain"
)

type abi struct {
	config domain.Config
	client domain.ClientGateway
}

// NewAbi ...
func NewAbi(
	config domain.Config,
	client domain.ClientGateway,
) domain.AbiUseCase {
	return &abi{
		config: config,
		client: client,
	}
}

// EncodeMessageBody method encode_message_body
func (a *abi) EncodeMessageBody(pOEMB domain.ParamsOfEncodeMessageBody) (int, error) {
	return a.client.Request("abi.encode_message_body", pOEMB)
}

// AttachSignatureToMessageBody method attach_signature_to_message_body
func (a *abi) AttachSignatureToMessageBody(pOASTMB domain.ParamsOfAttachSignatureToMessageBody) (int, error) {
	return a.client.Request("abi.attach_signature_to_message_body", pOASTMB)
}

// EncodeMessage method encode_message
func (a *abi) EncodeMessage(pOEM domain.ParamsOfEncodeMessage) (int, error) {
	return a.client.Request("abi.encode_message", pOEM)
}

// AttachSignature method attach_signature
func (a *abi) AttachSignature(pOAS domain.ParamsOfAttachSignature) (int, error) {
	return a.client.Request("abi.attach_signature", pOAS)
}

// DecodeMessage method decode_message
func (a *abi) DecodeMessage(pODM domain.ParamsOfDecodeMessage) (int, error) {
	return a.client.Request("abi.decode_message", pODM)
}

// DecodeMessageBody method decode_message_body
func (a *abi) DecodeMessageBody(pODMB domain.ParamsOfDecodeMessageBody) (int, error) {
	return a.client.Request("abi.decode_message_body", pODMB)
}

// EncodeAccount method encode_account
func (a *abi) EncodeAccount(pOEA domain.ParamsOfEncodeAccount) (int, error) {
	return a.client.Request("abi.encode_account", pOEA)
}
