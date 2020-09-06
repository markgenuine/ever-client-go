package contracts

import (
	"encoding/json"
)

// RunLocalMsg method contracts.run.local.msg
func RunLocalMsg(pOLEWM *ParamsOfLocalRunWithMsg) (string, string) {
	request, err := json.Marshal(pOLEWM)
	if err != nil {
		return "", ""
	}

	return "contracts.run.local.msg", string(request)
}

// RunLocal method contracts.run.local
func RunLocal(pOLR *ParamsOfLocalRun) (string, string) {
	request, err := json.Marshal(pOLR)
	if err != nil {
		return "", ""
	}

	return "contracts.run.local", string(request)
}

// RunLocalResp method contracts.run.local response
func RunLocalResp(resp string, err error) (*ResultOfLocalRun, error) {
	if err != nil {
		return nil, err
	}

	rOLR := &ResultOfLocalRun{}
	err = json.Unmarshal([]byte(resp), rOLR)
	if err != nil {
		return nil, err
	}

	return rOLR, nil
}

// RunMessage method contracts.run.message
func RunMessage(pOR *ParamsOfRun) (string, string) {
	request, err := json.Marshal(pOR)
	if err != nil {
		return "", ""
	}

	return "contracts.run.message", string(request)
}

func RunMessageResp(resp string, err error) (*EncodedMessage, error) {
	if err != nil {
		return nil, err
	}

	eM := &EncodedMessage{}
	err = json.Unmarshal([]byte(resp), eM)
	if err != nil {
		return nil, err
	}

	return eM, nil
}

// RunFee method contracts.run.fee
func RunFee(pOLR *ParamsOfLocalRun) (string, string) {
	request, err := json.Marshal(pOLR)
	if err != nil {
		return "", ""
	}

	return "contracts.run.fee", string(request)
}

// TvmGet method tvm.get
func TvmGet(pTVM *ParamsOfLocalRunGet) (string, string) {
	request, err := json.Marshal(pTVM)
	if err != nil {
		return "", ""
	}

	return "tvm.get", string(request)
}

// DeployAddress method contracts.deploy.address
func DeployAddress(pOD *ParamsOfGetDeployAddress) (string, string) {
	request, err := json.Marshal(pOD)
	if err != nil {
		return "", ""
	}

	return "contracts.deploy.address", string(request)
}

// AddressConvert method contracts.address.convert
func AddressConvert(pOfCA *ParamsOfConvertAddress) (string, string) {
	request, err := json.Marshal(pOfCA)
	if err != nil {
		return "", ""
	}

	return "contracts.address.convert", string(request)
}

//AddressConvertResp result contracts.address.convert method
func AddressConvertResp(resp string, err error) (*ResultOfConvertAddress, error) {
	if err != nil {
		return nil, err
	}

	strROCA := &ResultOfConvertAddress{}
	err = json.Unmarshal([]byte(resp), strROCA)
	if err != nil {
		return nil, err
	}

	return strROCA, nil
}

//DeployMessage method contracts.deploy.message
func DeployMessage(pOD *ParamsOfDeploy) (string, string) {
	request, err := json.Marshal(pOD)
	if err != nil {
		return "", ""
	}
	return "contracts.deploy.message", string(request)
}

//DeployEncodeUnsignedMessage method contracts.deploy.encode_unsigned_message
func DeployEncodeUnsignedMessage(pOD *ParamsOfEncodeUnsignedDeployMessage) (string, string) {
	request, err := json.Marshal(pOD)
	if err != nil {
		return "", ""
	}

	return "contracts.deploy.encode_unsigned_message", string(request)
}

//DeployEncodeUnsignedMessage method contracts.deploy.encode_unsigned_message response
func DeployEncodeUnsignedMessageResp(resp string, err error) (*ResultOfEncodeUnsignedDeployMessage, error) {
	if err != nil {
		return nil, err
	}

	rOEUDM := &ResultOfEncodeUnsignedDeployMessage{}
	err = json.Unmarshal([]byte(resp), rOEUDM)
	if err != nil {
		return nil, err
	}

	return rOEUDM, nil
}

//DeployData method contracts.deploy.data
func DeployData(pOD *ParamsOfGetDeployData) (string, string) {
	request, err := json.Marshal(pOD)
	if err != nil {
		return "", ""
	}
	return "contracts.deploy.data", string(request)
}

//DeployDataResp method contracts.deploy.data response
func DeployDataResp(resp string, err error) (*ResultOfGetDeployData, error) {
	if err != nil {
		return nil, err
	}

	rOGDD := &ResultOfGetDeployData{}
	err = json.Unmarshal([]byte(resp), rOGDD)
	if err != nil {
		return nil, err
	}

	return rOGDD, nil
}

// Deploy method contracts.deploy
func Deploy(pOD *ParamsOfDeploy) (string, string) {
	request, err := json.Marshal(pOD)
	if err != nil {
		return "", ""
	}

	return "contracts.deploy", string(request)
}

// Deploy method contracts.deploy
func DeployResp(resp string, err error) (*ResultOfDeploy, error) {
	if err != nil {
		return nil, err
	}

	rOD := &ResultOfDeploy{}
	err = json.Unmarshal([]byte(resp), rOD)
	if err != nil {
		return nil, err
	}

	return rOD, nil
}

// Load method contracts.load
func Load(lp *LoadParams) (string, string) {
	request, err := json.Marshal(lp)
	if err != nil {
		return "", ""
	}
	return "contracts.load", string(request)
}

//LoadResp result contracts.load method
func LoadResp(resp string, err error) (*LoadResult, error) {
	if err != nil {
		return nil, err
	}

	lR := &LoadResult{}
	err = json.Unmarshal([]byte(resp), lR)
	if err != nil {
		return nil, err
	}

	return lR, nil
}

// ResolveError method contracts.resolve.error
func ResolveError(pORE *ParamsOfResolveError) (string, string) {
	request, err := json.Marshal(pORE)
	if err != nil {
		return "", ""
	}
	return "contracts.resolve.error", string(request)
}

// RunEncodeUnsignedMessage method contracts.run.encode_unsigned_message
func RunEncodeUnsignedMessage(pOEURM *ParamsOfEncodeUnsignedRunMessage) (string, string) {
	request, err := json.Marshal(pOEURM)
	if err != nil {
		return "", ""
	}

	return "contracts.run.encode_unsigned_message", string(request)
}

//RunEncodeUnsignedMessageResp response method contracts.run.encode_unsigned_message
func RunEncodeUnsignedMessageResp(resp string, err error) (*EncodedUnsignedMessage, error) {
	if err != nil {
		return nil, err
	}

	eUM := &EncodedUnsignedMessage{}
	err = json.Unmarshal([]byte(resp), eUM)
	if err != nil {
		return nil, err
	}

	return eUM, nil
}

// RunBody method contracts.run.body
func RunBody(pOGRB *ParamsOfGetRunBody) (string, string) {
	request, err := json.Marshal(pOGRB)
	if err != nil {
		return "", ""
	}

	return "contracts.run.body", string(request)
}

func RunBodyResp(resp string, err error) (*ResultOfGetRunBody, error) {
	if err != nil {
		return nil, err
	}

	rOGRB := &ResultOfGetRunBody{}
	err = json.Unmarshal([]byte(resp), rOGRB)
	if err != nil {
		return nil, err
	}

	return rOGRB, nil
}

// Run method contracts.run
func Run(pOR *ParamsOfRun) (string, string) {
	request, err := json.Marshal(pOR)
	if err != nil {
		return "", ""
	}
	return "contracts.run", string(request)
}

//RunResp
func RunResp(resp string, err error) (*ResultOfRun, error) {
	if err != nil {
		return nil, err
	}

	rOR := &ResultOfRun{}
	err = json.Unmarshal([]byte(resp), rOR)
	if err != nil {
		return nil, err
	}

	return rOR, nil
}

// RunOutput method contracts.run.output
func RunOutput(pODRO *ParamsOfDecodeRunOutput) (string, string) {
	request, err := json.Marshal(pODRO)
	if err != nil {
		return "", ""
	}
	return "contracts.run.output", string(request)
}

//RunOutputResp
func RunOutputResp(resp string, err error) (*ResultOfDecode, error) {
	if err != nil {
		return nil, err
	}

	rOD := &ResultOfDecode{}
	err = json.Unmarshal([]byte(resp), rOD)
	if err != nil {
		return nil, err
	}

	return rOD, nil
}

// RunFeeMsg method contracts.run.fee.msg
func RunFeeMsg(pOLRWM *ParamsOfLocalRunWithMsg) (string, string) {
	request, err := json.Marshal(pOLRWM)
	if err != nil {
		return "", ""
	}
	return "contracts.run.fee.msg", string(request)
}

// RunUnknownInput method contracts.run.unknown.input
func RunUnknownInput(pODUR *ParamsOfDecodeUnknownRun) (string, string) {
	request, err := json.Marshal(pODUR)
	if err != nil {
		return "", ""
	}
	return "contracts.run.unknown.input", string(request)
}

// RunUnknownOutput method contracts.run.unknown.output
func RunUnknownOutput(pODUR *ParamsOfDecodeUnknownRun) (string, string) {
	request, err := json.Marshal(pODUR)
	if err != nil {
		return "", ""
	}
	return "contracts.run.unknown.output", string(request)
}

func UnknownOutputResp(resp string, err error) (*ResultOfDecodeUnknownRun, error) {
	if err != nil {
		return nil, err
	}

	rODUR := &ResultOfDecodeUnknownRun{}
	err = json.Unmarshal([]byte(resp), rODUR)
	if err != nil {
		return nil, err
	}

	return rODUR, nil
}

//EncodeMessageWithSign method contracts.encode_message_with_sign
func EncodeMessageWithSign(pOEMWS *ParamsOfEncodeMessageWithSign) (string, string) {
	request, err := json.Marshal(pOEMWS)
	if err != nil {
		return "", ""
	}
	return "contracts.encode_message_with_sign", string(request)
}

//EncodeMessageWithSignResp result contracts.encode_message_with_sign method
func EncodeMessageWithSignResp(resp string, err error) (*EncodedMessage, error) {
	if err != nil {
		return nil, err
	}

	eM := &EncodedMessage{}
	err = json.Unmarshal([]byte(resp), eM)
	if err != nil {
		return nil, err
	}

	return eM, nil
}

// ParseMessage method contracts.parse.message
func ParseMessage(iB *InputBoc) (string, string) {
	request, err := json.Marshal(iB)
	if err != nil {
		return "", ""
	}
	return "contracts.parse.message", string(request)
}

//FunctionID method contracts.function.id
func FunctionID(pOGFI *ParamsOfGetFunctionId) (string, string) {
	request, err := json.Marshal(pOGFI)
	if err != nil {
		return "", ""
	}
	return "contracts.function.id", string(request)
}

//FunctionIDResp ...
func FunctionIDResp(resp string, err error) (*ResultOfGetFunctionId, error) {
	if err != nil {
		return nil, err
	}

	rOGFI := &ResultOfGetFunctionId{}
	err = json.Unmarshal([]byte(resp), rOGFI)
	if err != nil {
		return nil, err
	}

	return rOGFI, nil
}

// FindShard method contracts.find.shard
func FindShard(pOFS *ParamsOfFindShard) (string, string) {
	request, err := json.Marshal(pOFS)
	if err != nil {
		return "", ""
	}
	return "contracts.find.shard", string(request)
}

// SendMessage method contracts.send.message
func SendMessage(eM *EncodedMessage) (string, string) {
	request, err := json.Marshal(eM)
	if err != nil {
		return "", ""
	}
	return "contracts.send.message", string(request)
}

// SendMessageResp
func SendMessageResp(resp string, err error) (*MessageProcessingState, error) {
	if err != nil {
		return nil, err
	}

	mPS := &MessageProcessingState{}
	err = json.Unmarshal([]byte(resp), mPS)
	if err != nil {
		return nil, err
	}

	return mPS, nil
}

// ProcessMessage method contracts.process.message
func ProcessMessage(pOPM *ParamsOfProcessMessage) (string, string) {
	request, err := json.Marshal(pOPM)
	if err != nil {
		return "", ""
	}
	return "contracts.process.message", string(request)
}

// ProcessTransaction method contracts.process.transaction
func ProcessTransaction(pOPT *ParamsOfProcessTransaction) (string, string) {
	request, err := json.Marshal(pOPT)
	if err != nil {
		return "", ""
	}
	return "contracts.process.transaction", string(request)
}

// WaitTransaction method contracts.wait.transaction
func WaitTransaction(pOWT *ParamsOfWaitTransaction) (string, string) {
	request, err := json.Marshal(pOWT)
	if err != nil {
		return "", ""
	}
	return "contracts.wait.transaction", string(request)
}
