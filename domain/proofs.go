package domain

import "encoding/json"

var ProofsErrorCode map[string]int

type (
	ParamsOfProofBlockData struct {
		Block json.RawMessage `json:"block"`
	}

	ParamsOfProofTransactionData struct {
		Transaction json.RawMessage `json:"transaction"`
	}

	ParamsOfProofMessageData struct {
		Message json.RawMessage `json:"message"`
	}

	ProofsUseCase interface {
		ProofBlockData(*ParamsOfProofBlockData) error
		ProofTransactionData(*ParamsOfProofTransactionData) error
		ParamsMessageData(*ParamsOfProofMessageData) error
	}
)

func init() {
	AbiErrorCode = map[string]int{
		"InvalidData":           901,
		"ProofCheckFailed":      902,
		"InternalError":         903,
		"DataDiffersFromProven": 904,
	}
}
