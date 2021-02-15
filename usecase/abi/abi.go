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

// EncodeMessageBody Encode message body according to ABI function call.
func (a *abi) EncodeMessageBody(pOEMB *domain.ParamsOfEncodeMessageBody) (*domain.ResultOfEncodeMessageBody, error) {
	result := new(domain.ResultOfEncodeMessageBody)
	err := a.client.GetResult("abi.encode_message_body", pOEMB, result)
	return result, err
}

// AttachSignatureToMessageBody method attach_signature_to_message_body.
func (a *abi) AttachSignatureToMessageBody(pOASTMB *domain.ParamsOfAttachSignatureToMessageBody) (*domain.ResultOfAttachSignatureToMessageBody, error) {
	result := new(domain.ResultOfAttachSignatureToMessageBody)
	err := a.client.GetResult("abi.attach_signature_to_message_body", pOASTMB, result)
	return result, err
}

// EncodeMessage encodes an ABI-compatible message.
func (a *abi) EncodeMessage(pOEM *domain.ParamsOfEncodeMessage) (*domain.ResultOfEncodeMessage, error) {
	result := new(domain.ResultOfEncodeMessage)
	err := a.client.GetResult("abi.encode_message", pOEM, result)
	return result, err
}

// AttachSignature сombines hex-encoded signature with base64-encoded unsigned_message. Returns signed message encoded in base64.
func (a *abi) AttachSignature(pOAS *domain.ParamsOfAttachSignature) (*domain.ResultOfAttachSignature, error) {
	result := new(domain.ResultOfAttachSignature)
	err := a.client.GetResult("abi.attach_signature", pOAS, result)
	return result, err
}

// DecodeMessage Decodes message body using provided message BOC and ABI.
func (a *abi) DecodeMessage(pODM *domain.ParamsOfDecodeMessage) (*domain.DecodedMessageBody, error) {
	result := new(domain.DecodedMessageBody)
	err := a.client.GetResult("abi.decode_message", pODM, result)
	return result, err
}

// DecodeMessageBody Decodes message body using provided body BOC and ABI.
func (a *abi) DecodeMessageBody(pODMB *domain.ParamsOfDecodeMessageBody) (*domain.DecodedMessageBody, error) {
	result := new(domain.DecodedMessageBody)
	err := a.client.GetResult("abi.decode_message_body", pODMB, result)
	return result, err
}

// EncodeAccount Creates account state BOC.
func (a *abi) EncodeAccount(pOEA *domain.ParamsOfEncodeAccount) (*domain.ResultOfEncodeAccount, error) {
	result := new(domain.ResultOfEncodeAccount)
	err := a.client.GetResult("abi.encode_account", pOEA, result)
	return result, err
}
