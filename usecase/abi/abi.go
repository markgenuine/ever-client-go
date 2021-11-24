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

// EncodeMessageBody - Encode message body according to ABI function call.
func (a *abi) EncodeMessageBody(pOEMB *domain.ParamsOfEncodeMessageBody) (*domain.ResultOfEncodeMessageBody, error) {
	result := new(domain.ResultOfEncodeMessageBody)
	err := a.client.GetResult("abi.encode_message_body", pOEMB, result)
	return result, err
}

// AttachSignatureToMessageBody - method attach_signature_to_message_body
func (a *abi) AttachSignatureToMessageBody(pOASTMB *domain.ParamsOfAttachSignatureToMessageBody) (*domain.ResultOfAttachSignatureToMessageBody, error) {
	result := new(domain.ResultOfAttachSignatureToMessageBody)
	err := a.client.GetResult("abi.attach_signature_to_message_body", pOASTMB, result)
	return result, err
}

// EncodeMessage - Encodes an ABI-compatible message.
// Allows to encode deploy and function call messages, both signed and unsigned.
func (a *abi) EncodeMessage(pOEM *domain.ParamsOfEncodeMessage) (*domain.ResultOfEncodeMessage, error) {
	result := new(domain.ResultOfEncodeMessage)
	err := a.client.GetResult("abi.encode_message", pOEM, result)
	return result, err
}

// EncodeInternalMessage - Encodes an internal ABI-compatible message
// Allows to encode deploy and function call messages.
func (a *abi) EncodeInternalMessage(pOEIM *domain.ParamsOfEncodeInternalMessage) (*domain.ResultOfEncodeInternalMessage, error) {
	result := new(domain.ResultOfEncodeInternalMessage)
	err := a.client.GetResult("abi.encode_internal_message", pOEIM, result)
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

// DecodeAccountData - Decodes account data using provided data BOC and ABI.
// Note: this feature requires ABI 2.1 or higher.
func (a *abi) DecodeAccountData(pODAD *domain.ParamsOfDecodeAccountData) (*domain.ResultOfDecodeData, error) {
	result := new(domain.ResultOfDecodeData)
	err := a.client.GetResult("abi.decode_account_data", pODAD, result)
	return result, err
}

// UpdateInitialData - Updates initial account data with initial values for the contract's static variables and owner's public key.
// This operation is applicable only for initial account data (before deploy). If the contract is already deployed, its data doesn't contain
// this data section any more.
func (a *abi) UpdateInitialData(pOUID *domain.ParamsOfUpdateInitialData) (*domain.ResultOfUpdateInitialData, error) {
	result := new(domain.ResultOfUpdateInitialData)
	err := a.client.GetResult("abi.update_initial_data", pOUID, result)
	return result, err
}

// DecodeInitialData - Decodes initial values of a contract's static variables and owner's public key from account initial data
// This operation is applicable only for initial account data (before deploy). If the contract is already deployed, its data doesn't
// contain this data section any more.
func (a *abi) DecodeInitialData(pODID *domain.ParamsOfDecodeInitialData) (*domain.ResultOfDecodeInitialData, error) {
	result := new(domain.ResultOfDecodeInitialData)
	err := a.client.GetResult("abi.decode_initial_data", pODID, result)
	return result, err
}

// DecodeBoc - Decodes BOC into JSON as a set of provided parameters.
//
// Solidity functions use ABI types for builder encoding. The simplest way to decode such a BOC is to use ABI decoding.
// ABI has it own rules for fields layout in cells so manually encoded BOC can not be described in terms of ABI rules.
//
// To solve this problem we introduce a new ABI type Ref(<ParamType>) which allows to store ParamType ABI parameter in
// cell reference and, thus, decode manually encoded BOCs. This type is available only in decode_boc function and will
// not be available in ABI messages encoding until it is included into some ABI revision.
//
// Such BOC descriptions covers most users needs. If someone wants to decode some BOC which can not be described by these
// rules (i.e. BOC with TLB containing constructors of flags defining some parsing conditions) then they can decode the
// fields up to fork condition, check the parsed data manually, expand the parsing schema and then decode the whole BOC
// with the full schema.
func (a *abi) DecodeBoc(boc *domain.ParamsOfDecodeBoc) (*domain.ResultOfDecodeBoc, error) {
	result := new(domain.ResultOfDecodeBoc)
	err := a.client.GetResult("abi.decode_boc", boc, result)
	return result, err
}
