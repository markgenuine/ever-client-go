package proofs

import (
	"github.com/move-ton/ever-client-go/domain"
)

type proofs struct {
	config domain.Config
	client domain.ClientGateway
}

// NewProofs ...
func NewProofs(
	config domain.Config,
	client domain.ClientGateway,
) domain.ProofsUseCase {
	return &proofs{
		config: config,
		client: client,
	}
}

// ProofBlockData - Proves that a given block's data, which is queried from EverOS API, can be trusted.
//
// This function checks block proofs and compares given data with the proven. If the given data differs from the proven,
// the exception will be thrown. The input param is a single block's JSON object, which was queried from DApp server
// using functions such as net.query, net.query_collection or net.wait_for_collection. If block's BOC is not provided in
// the JSON, it will be queried from DApp server (in this case it is required to provide at least id of block).
//
// Please note, that joins (like signatures in Block) are separated entities and not supported, so function will throw an
// exception in a case if JSON being checked has such entities in it.
//
// If cache_in_local_storage in config is set to true (default), downloaded proofs and master-chain BOCs are saved into
// the persistent local storage (e.g. file system for native environments or browser's IndexedDB for the web); otherwise
// all the data is cached only in memory in current client's context and will be lost after destruction of the client.
func (pr *proofs) ProofBlockData(data *domain.ParamsOfProofBlockData) error {
	_, err := pr.client.GetResponse("proofs.proof_block_data", data)
	return err
}

// ProofTransactionData - Proves that a given transaction's data, which is queried from EverOS API, can be trusted.
//
// This function requests the corresponding block, checks block proofs, ensures that given transaction exists in the
// proven block and compares given data with the proven. If the given data differs from the proven, the exception will
// be thrown. The input parameter is a single transaction's JSON object (see params description), which was queried from
// EverOS API using functions such as net.query, net.query_collection or net.wait_for_collection.
//
// If transaction's BOC and/or block_id are not provided in the JSON, they will be queried from EverOS API (in this case
// it is required to provide at least id of transaction).
//
// Please note, that joins (like account, in_message, out_messages, etc. in Transaction entity) are separated entities
// and not supported, so function will throw an exception in a case if JSON being checked has such entities in it.
//
// If cache_in_local_storage in config is set to true (default), downloaded proofs and master-chain BOCs are saved into
// the persistent local storage (e.g. file system for native environments or browser's IndexedDB for the web); otherwise
// all the data is cached only in memory in current client's context and will be lost after destruction of the client.
func (pr *proofs) ProofTransactionData(data *domain.ParamsOfProofTransactionData) error {
	_, err := pr.client.GetResponse("proofs.proof_transaction_data", data)
	return err
}

// ParamsMessageData - Proves that a given message's data, which is queried from EverOS API, can be trusted.
//
// This function first proves the corresponding transaction, ensures that the proven transaction refers to the given
// message and compares given data with the proven. If the given data differs from the proven, the exception will be
// thrown. The input parameter is a single message's JSON object (see params description), which was queried from EverOS
// API using functions such as net.query, net.query_collection or net.wait_for_collection.
//
// If message's BOC and/or non-null src_transaction.id or dst_transaction.id are not provided in the JSON, they will be
// queried from EverOS API.
//
// Please note, that joins (like block, dst_account, dst_transaction, src_account, src_transaction, etc. in Message
// entity) are separated entities and not supported, so function will throw an exception in a case if JSON being checked has such entities in it.
//
// For more information about proofs checking, see description of proof_block_data function.
func (pr *proofs) ParamsMessageData(data *domain.ParamsOfProofMessageData) error {
	_, err := pr.client.GetResponse("proofs.proof_message_data", data)
	return err
}
