package contracts

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	goton "github.com/move-ton/ton-client-go"
	crypto "github.com/move-ton/ton-client-go/crypto"
)

func TestContracts(t *testing.T) {

	client, err := goton.InitClient(goton.NewConfig(0))
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

	tKey := &goton.TONKey{}
	tKey.Public = "1c1a1900b3ed6bd0401ac2315d570288b85823d639b66ad0173bcb6573154962"
	tKey.Secret = "a45aa3652d25c151c6c09ae8f561e1459593374bc8e566e1f460218a3dd2b7e9"
	address := "0:7c75c7c5b5a92858fa1d16b3f3a94acf9ca49b5a282b1412c9f3f598531eb64a"

	t.Run("TestSimpleWalletContract", func(t *testing.T) {
		abiByteWC, err := openFile("./samples/SimpleWallet/Wallet.abi.json")
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		abiSW := &ABI{}
		err = json.Unmarshal(abiByteWC, abiSW)

		//Получаем contracts binary
		binaryByte, err := openFile("./samples/SimpleWallet/Wallet.tvc")
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		t.Run("TestLoad", func(t *testing.T) {
			lR, err := LoadResp(client.Request(Load(&LoadParams{Address: address})))
			if err != nil {
				t.Errorf("test Failed - error method Load: %s", err)
			}

			if lR.ID == "" || lR.BalanceGrams == "" {
				t.Errorf("test Failed - error response method Load")
			}
		})

		t.Run("TestDeployAddress", func(t *testing.T) {
			dA, err := client.Request(DeployAddress(&ParamsOfGetDeployAddress{Abi: *abiSW, ImageBase64: base64.StdEncoding.EncodeToString(binaryByte), KeyPair: tKey}))
			if err != nil {
				t.Errorf("test Failed - error method DeployAddress: %s", err)
			}

			if dA != address {
				t.Errorf("test Failed - error response method DeployAddress")
			}
		})

		t.Run("TestAddressConvert", func(t *testing.T) {
			resp, err := AddressConvertResp(client.Request(AddressConvert(&ParamsOfConvertAddress{Address: address, ConvertTo: "Base64", Base64Params: Base64AddressParams{Url: false, Test: true, Bounce: true}})))
			if err != nil {
				t.Errorf("test Failed - error method AddressConvert: %s", err)
			}

			if resp.Address != "kQB8dcfFtakoWPodFrPzqUrPnKSbWigrFBLJ8/WYUx62ShSE" {
				t.Errorf("test Failed - error response addres incorrect in method AddressConvert")
			}

			resp, err = AddressConvertResp(client.Request(AddressConvert(&ParamsOfConvertAddress{Address: address, ConvertTo: "Base"})))
			if err == nil {
				t.Errorf("test Failed - error method AddressConvert: %s", err)
			}

			resp, err = AddressConvertResp(client.Request(AddressConvert(&ParamsOfConvertAddress{Address: address, ConvertTo: "Base64", Base64Params: Base64AddressParams{Url: false, Test: true}})))
			if err != nil {
				t.Errorf("test Failed - error method AddressConvert: %s", err)
			}

			if resp.Address != "0QB8dcfFtakoWPodFrPzqUrPnKSbWigrFBLJ8/WYUx62SklB" {
				t.Errorf("test Failed - error response addres incorrect in method AddressConvert")
			}
		})

		t.Run("TestDeployMessage", func(t *testing.T) {
			result, err := RunMessageResp(client.Request(DeployMessage(&ParamsOfDeploy{Abi: *abiSW, ImageBase64: base64.StdEncoding.EncodeToString(binaryByte), KeyPair: tKey, ConstructorParams: []byte("{}")})))
			if err != nil {
				t.Errorf("test Failed - error method DeployMessage: %s", err)
			}

			if address != result.Address {
				t.Errorf("test Failed - error response value method DeployMessage")
			}
		})

		t.Run("TestDeployEncodeUnsignedMessage", func(t *testing.T) {
			result, err := DeployEncodeUnsignedMessageResp(client.Request(DeployEncodeUnsignedMessage(&ParamsOfEncodeUnsignedDeployMessage{Abi: *abiSW, ConstructorParams: []byte("{}"), ImageBase64: base64.StdEncoding.EncodeToString(binaryByte), PublicKeyHex: tKey.Public})))
			if err != nil {
				t.Errorf("test Failed - error method DeployEncodeUnsignedMessage: %s", err)
			}

			if result.Encoded.UnsignedBytesBase64 == "" {
				t.Errorf("test Failed - error method DeployEncodeUnsignedMessage response UnsignedBytesBase64 incorected")
			}

			if result.Encoded.BytesToSignBase64 == "" {
				t.Errorf("test Failed - error method DeployEncodeUnsignedMessage response BytesToSignBase64 incorected")
			}

			if result.AddressHex != address {
				t.Errorf("test Failed - error method DeployEncodeUnsignedMessage response AddressHex incorected")
			}
		})

		t.Run("TestDeployData", func(t *testing.T) {
			result, err := DeployDataResp(client.Request(DeployData(&ParamsOfGetDeployData{Abi: *abiSW, ImageBase64: base64.StdEncoding.EncodeToString(binaryByte), PublicKeyHex: tKey.Public})))
			if err != nil {
				t.Errorf("test Failed - error method DeployData: %s", err)
			}

			if result.DataBase64 != "te6ccgEBAgEAKAABAcABAEPQBwaGQCz7WvQQBrCMV1XAoi4WCPWObZq0Bc7y2VzFUlig" {
				t.Errorf("test Failed - error responce method DeployData, value incorrected")
			}
		})

		t.Run("TestDeploy", func(t *testing.T) {
			result, err := DeployResp(client.Request(Deploy(&ParamsOfDeploy{Abi: *abiSW, ImageBase64: base64.StdEncoding.EncodeToString(binaryByte), KeyPair: tKey})))
			if err != nil {
				t.Errorf("test Failed - error method Deploy: %s", err)
			}

			if result.Address != address {
				t.Errorf("test Failed - error responce method Deploy, value incorrected")
			}
		})

		//don't work?!?
		t.Run("TestResolveError", func(t *testing.T) {
			type Temp struct {
				Dest   string `json:"dest"`
				Value  int    `json:"value"`
				Bounce bool   `json:"bounce"`
			}

			tt := Temp{Dest: "0:1ab22c364214e24b782bc4966e23874b1c0cc682e8dba2d24a0561bb27d04221", Value: 1000000000, Bounce: false}
			//req,_ := json.Marshal(tt)
			//inputs := `{"dest": "0:1ab22c364214e24b782bc4966e23874b1c0cc682e8dba2d24a0561bb27d04221", "value": 1000000000, "bounce": False}`
			rr, rr1 := RunMessage(&ParamsOfRun{Address: address, Abi: *abiSW, FunctionName: "sendTransaction", Input: tt, KeyPair: tKey})
			result, err := client.Request(rr, rr1)
			if err != nil {
				t.Errorf("test Failed - error method ResolveError: %s", err)
			}

			fmt.Println(result)

			//error = {'core_version': '0.25.3', 'source': 'node', 'code': 3025, 'message': 'Contract execution was terminated with error', 'message_processing_state': None, 'data': {'account_address': '0:2df86dd43c3fcd8cd9704126a6ecb6439116b39f0f9fb97c239dd67bdb6896b8', 'config_server': 'net.ton.dev', 'exit_code': 100, 'original_error': {'code': 1012, 'core_version': '0.25.3', 'data': {'message_id': '5794586d5c51015684445eee4b0ab4a40924ad55b82a81199ab787395a58b1ba', 'sending_time': 'Thu, 20 Aug 2020 16:47:12 +0300 (1597931232)', 'timeout': 50000}, 'message': 'Transaction was not produced during the specified timeout', 'message_processing_state': {'lastBlockId': '71761738a2ab096413441d36ed68f47cf523d7eb4e67007bc7e4fbb23012c946', 'sendingTime': 1597931232}, 'source': 'node'}, 'phase': 'computeVm', 'query_url': 'https://net.ton.dev', 'transaction_id': None}}
			//account = {'acc_type': 1, 'balance': '0x1409beb9a', 'balance_other': None, 'boc': 'te6ccgECDwEAAooAAnHAAt+G3UPD/NjNlwQSam7LZDkRaznw+fuXwjndZ722iWuCHoRewvnzvugAAAS0AhvKDUBQJvrmk0ACAQCRwHj7/2pNrtirjXjSNxqZd4xKOgWP1OQHngk+Whm0RyIAAAF0DAQwDuA8ff+1JtdsVca8aRuNTLvGJR0Cx+pyA88Eny0M2iORQAIm/wD0pCAiwAGS9KDhiu1TWDD0oQUDAQr0pCD0oQQAAAIBIAkGAQL/BwH+fyHtRNAg10nCAY4U0//TP9MA1wv/+Gp/+GH4Zvhj+GKOG/QFcPhqcAGAQPQO8r3XC//4YnD4Y3D4Zn/4YeLTAAGOEoECANcYIPkBWPhCIPhl+RDyqN7TPwGOHvhDIbkgnzAg+COBA+iogggbd0Cgud6S+GPggDTyNNjTHwHwAQgADvhHbpLyPN4CASALCgDdvUWq+f/CC3Rx52omgQa6ThAMcKaf/pn+mAa4X//DU//DD8M3wx/DFHDfoCuHw1OADAIHoHeV7rhf/8MTh8Mbh8Mz/8MPFvfCN5Obj8M2j8AHwhfDV8IWRl//wh54Wf/CNnhYB8JQDl/+T2qj/8M8AgEgDgwB8bsV75NfhBbo4X7UTQ0//TP9MA1wv/+Gp/+GH4Zvhj+GLe+kDXDX+V1NHQ03/f1wwAldTR0NIA39H4RSBukjBw3vhKuvLgZPgAIcIAIJcwIfgnbxC53vLgZSEjIsjPhYDKAHPPQM4B+gKAac9Az4HPgclw+wBfA8D/gNADqOF/hCyMv/+EPPCz/4Rs8LAPhKAcv/ye1U3n/4ZwBq3XAi0NYCMdIAMNwhxwCQ4CHXDR+S8jzhUxGQ4cEEIoIQ/////byxkvI84AHwAfhHbpLyPN4=', 'code_hash': '84e5d7c2abb4bff6d7f3402eba34232de6c534b45cb6b4aef0d30a6a031e51e2', 'data_hash': '8a69a508698697e5d82abc45d76a31c24c7c1bdfb84fa6ffe8e6ac9590c4834e', 'id': '0:2df86dd43c3fcd8cd9704126a6ecb6439116b39f0f9fb97c239dd67bdb6896b8', 'last_paid': 1597929437}
			//
			//def __call():
			//client.contracts.resolve_error(
			//	address=self.contract_address, account=account,
			//	message_b64=message["messageBodyBase64"], time=0, error=error)
			//self.assertRaises(TonException, __call)
			//res, err := client.Request()
			//if err != nil {
			//	fmt.Println("err: ", err)
			//}
			//
			//fmt.Println(res)
		})

	})

	t.Run("TestPiggyBankContract", func(t *testing.T) {

		address = "0:668b5c83056ebf1852cc7af4e61c8a421056c0311f035a39e5baf7ce28b14728"
		//owner := "0:1ab22c364214e24b782bc4966e23874b1c0cc682e8dba2d24a0561bb27d04221"
		//self.constructor_params = {
		//	"pb_owner": self.owner_address,
		//		"pb_limit": 5 * 10 ** 9
		//}

		abiByte, err := openFile("./samples/PiggyBank/PiggyBank.abi.json")
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		abiS := &ABI{}
		err = json.Unmarshal(abiByte, abiS)

		type T struct{}
		pOR := &ParamsOfRun{}
		pOR.Abi = *abiS
		pOR.Address = address
		pOR.FunctionName = "getVersion"
		pOR.KeyPair = tKey
		pOR.Input = T{}

		pOLR := &ParamsOfLocalRun{}
		pOLR.Abi = *abiS
		pOLR.Address = address
		pOLR.FunctionName = "getVersion"
		//pOLR.KeyPair = tKey
		pOLR.Input = T{}

		type tempStruct struct {
			Value0 string
		}

		t.Run("TestRunMessage", func(t *testing.T) {
			rMR, err := RunMessageResp(client.Request(RunMessage(pOR)))
			if err != nil {
				t.Errorf("test Failed - error method RunMessage: %s", err)
			}

			if rMR.Address != address {
				t.Errorf("test Failed - error method RunMessage in value")
			}
		})

		t.Run("TestRunEncodeUnsignedMessage", func(t *testing.T) {
			eUM, err := RunEncodeUnsignedMessageResp(client.Request(RunEncodeUnsignedMessage(&ParamsOfEncodeUnsignedRunMessage{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			if err != nil {
				t.Errorf("test Failed - error method RunEncodeUnsignedMessage: %s", err)
			}

			if eUM.UnsignedBytesBase64 == "" || eUM.BytesToSignBase64 == "" {
				t.Errorf("test Failed - error method RunEncodeUnsignedMessage in value")
			}

		})

		t.Run("TestRunBody", func(t *testing.T) {
			rOGRB, err := RunBodyResp(client.Request(RunBody(&ParamsOfGetRunBody{Abi: *abiS, KeyPair: tKey, FunctionName: "getVersion", Header: T{}, Params: T{}, Internal: true})))
			if err != nil {
				t.Errorf("test Failed - error method RunBody: %s", err)
			}

			if rOGRB.BodyBase64 != "te6ccgEBAQEABgAACFoZzZI=" {
				t.Errorf("test Failed - error method RunBody in value")
			}

		})

		t.Run("TestRun", func(t *testing.T) {
			rOR, err := RunResp(client.Request(Run(&ParamsOfRun{Abi: *abiS, Address: address, KeyPair: tKey, FunctionName: "getVersion", Input: T{}})))
			if err != nil {
				t.Errorf("test Failed - error method Run: %s", err)
			}

			tS := &tempStruct{}
			_ = json.Unmarshal(*rOR.Output, tS)
			if tS.Value0 == "" {
				t.Errorf("test Failed - error responce method Run, value incorrected")
			}
		})

		t.Run("TestRunLocalMessage", func(t *testing.T) {

			rMR, err := RunMessageResp(client.Request(RunMessage(pOR)))
			if err != nil {
				t.Errorf("test Failed - error method RunMessage in RunLocalMessage: %s", err)
			}

			pOFRWM := &ParamsOfLocalRunWithMsg{}
			pOFRWM.Address = address
			pOFRWM.Abi = *abiS
			pOFRWM.FunctionName = "getVersion"
			pOFRWM.MessageBase64 = rMR.MessageBodyBase64

			rLR, err := RunLocalResp(client.Request(RunLocalMsg(pOFRWM)))
			if err != nil {
				fmt.Println("error: ", err)
			}

			type tempStruct struct {
				Value0 string
			}
			tS := &tempStruct{}
			_ = json.Unmarshal(*rLR.Output, tS)

			if tS.Value0 != "0x1" {
				t.Errorf("test Failed - RunLocalMessage incorrect value: %s", err)
			}

			emptyFees := RunFees{}
			if rLR.Fees != emptyFees {
				t.Errorf("test Failed - RunLocalMessage fees must empty struct incorrect value: %s", err)
			}

			pOFRWM.FullRun = true

			rLR2, err := RunLocalResp(client.Request(RunLocalMsg(pOFRWM)))
			if err != nil {
				fmt.Println("error: ", err)
			}

			if rLR2.Fees == emptyFees {
				t.Errorf("test Failed - RunLocalMessage fees don't must empty struct incorrect value: %s", err)
			}
		})

		t.Run("TestRunLocal", func(t *testing.T) {

			result, err := RunLocalResp(client.Request(RunLocal(pOLR)))
			if err != nil {
				t.Errorf("test Failed - error method RunLocal: %s", err)
			}

			var rF RunFees
			if result.Fees != rF {
				t.Errorf("test Failed - error method RunLocal and value params: %s", err)
			}
		})

		t.Run("TestRunOutput", func(t *testing.T) {
			rOD, err := RunOutputResp(client.Request(RunOutput(&ParamsOfDecodeRunOutput{Abi: *abiS, FunctionName: "getVersion", BodyBase64: "te6ccgEBAQEAJgAASNoZzZIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQ"})))
			if err != nil {
				t.Errorf("test Failed - error method RunOutput: %s", err)
			}

			tS := &tempStruct{}
			_ = json.Unmarshal(*rOD.Output, tS)
			if tS.Value0 != "0x1" {
				t.Errorf("test Failed - incorrect value")
			}
		})

		t.Run("TestRunFee", func(t *testing.T) {
			result, err := RunLocalResp(client.Request(RunFee(pOLR)))
			if err != nil {
				t.Errorf("test Failed - error method RunLocal: %s", err)
			}

			tS := &tempStruct{}
			_ = json.Unmarshal(*result.Output, tS)

			if tS.Value0 != "0x1" {
				t.Errorf("test Failed - incorrect value")
			}
		})

		//don't work
		t.Run("TestRunFeeMsg", func(t *testing.T) {
			message, err := RunMessageResp(client.Request(RunMessage(&ParamsOfRun{Abi: *abiS, Address: address, FunctionName: "getVersion", Input: T{}})))
			if err != nil {
				t.Errorf("test Failed - error method RunMessage in RunFeeMsg: %s", err)
			}

			result, err := client.Request(RunFeeMsg(&ParamsOfLocalRunWithMsg{Address: address, FunctionName: "getVersion", MessageBase64: message.MessageBodyBase64}))
			if err != nil {
				t.Errorf("test Failed - error method RunFeeMsg: %s", err)
			}

			fmt.Println(result)
		})

		t.Run("TestRunUnknownInput", func(t *testing.T) {
			body, err := RunBodyResp(client.Request(RunBody(&ParamsOfGetRunBody{Abi: *abiS, FunctionName: "getVersion", Header: T{}, Params: T{}})))
			if err != nil {
				t.Errorf("test Failed - error method RunBody in RunUnknownInput: %s", err)
			}

			unknownInput, err := UnknownOutputResp(client.Request(RunUnknownInput(&ParamsOfDecodeUnknownRun{Abi: *abiS, BodyBase64: body.BodyBase64})))
			if err != nil {
				t.Errorf("test Failed - error method RunUnknownInput: %s", err)
			}

			if unknownInput.Function != "getVersion" || string(*unknownInput.Output) != "{}" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("RunUnknownOutput", func(t *testing.T) {
			body := "te6ccgEBAQEAJgAASNoZzZIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQ=="
			unknownOutput, err := UnknownOutputResp(client.Request(RunUnknownOutput(&ParamsOfDecodeUnknownRun{Abi: *abiS, BodyBase64: body})))
			if err != nil {
				t.Errorf("test Failed - error method RunUnknownOutput: %s", err)
			}

			if unknownOutput.Function != "getVersion" || string(*unknownOutput.Output) != "{\"value0\":\"0x1\"}" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestEncodeMessageWithSign", func(t *testing.T) {
			//# Create unsigned message
			unsigned, err := RunEncodeUnsignedMessageResp(client.Request(RunEncodeUnsignedMessage(&ParamsOfEncodeUnsignedRunMessage{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			if err != nil {
				t.Errorf("test Failed - error method EncodeUnsignedMessage in EncodeMessageWithSign: %s", err)
			}

			//# Create signature
			dd := &goton.NaclSign{InputMessage: &goton.InputMessage{Base64: unsigned.BytesToSignBase64}, Key: tKey.Secret + tKey.Public, OutputEncoding: goton.TONOutputEncodingBase64}
			signature, err := client.Request(crypto.NaclSignDetached(dd))

			//# Sign message
			signed, err := EncodeMessageWithSignResp(client.Request(EncodeMessageWithSign(&ParamsOfEncodeMessageWithSign{Abi: *abiS, UnsignedBytesBase64: unsigned.UnsignedBytesBase64, SignBytesBase64: signature})))
			if err != nil {
				t.Errorf("test Failed - incorrect value in response Sign Message, %s", err)
			}
			if signed.Address != address {
				t.Errorf("test Failed - incorrect value in response")
			}

			//# Run message
			result, err := RunResp(client.Request(ProcessMessage(&ParamsOfProcessMessage{Abi: *abiS, Message: *signed, FunctionName: "getVersion"})))
			if result == nil || string(*result.Output) != "{\"value0\":\"0x1\"}" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestParseMessage", func(t *testing.T) {
			messageBoc := "te6ccgEBAQEAcQAA3YgAzRa5BgrdfjClmPXpzDkUhCCtgGI+BrRzy3XvnFFijlAGkFJ1s8KnJdQgxR+kLP+yGcqn44lVZeU8uxDkYRvny3R/yxeAzUDFMudyk6jKu2fqeazMGmcUKztS4MSgNFscKAAABc8AC9m5TL3Gng=="
			result, err := client.Request(ParseMessage(&InputBoc{BocBase64: messageBoc}))
			if err != nil {
				t.Errorf("test Failed - error method ParseMessage: %s", err)
			}

			if result != `{"dst":"0:668b5c83056ebf1852cc7af4e61c8a421056c0311f035a39e5baf7ce28b14728"}` {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestFunctionID", func(t *testing.T) {
			result, _ := FunctionIDResp(client.Request(FunctionID(&ParamsOfGetFunctionId{Abi: *abiS, Function: "getVersion"})))
			if result.ID != 1511640466 {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestFindShard", func(t *testing.T) {
			shards := []byte(`{"workchain_id": 0, "shard": "0800000000000000"},
				{"workchain_id": 0, "shard": "1800000000000000"},`)
			res := (*json.RawMessage)(&shards)
			result, _ := client.Request(FindShard(&ParamsOfFindShard{Address: address, Shards: res}))
			if result != "" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestSendMessage", func(t *testing.T) {
			message, _ := RunMessageResp(client.Request(RunMessage(&ParamsOfRun{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			result, _ := SendMessageResp(client.Request(SendMessage(message)))
			if result.LastBlockID == "" || result.SendingTime == 0 {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestProcessMessage", func(t *testing.T) {
			//message = client.contracts.run_message(
			message, _ := RunMessageResp(client.Request(RunMessage(&ParamsOfRun{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			resultUnknown, _ := RunResp(client.Request(ProcessMessage(&ParamsOfProcessMessage{Message: *message, FunctionName: "getVersion", Abi: *abiS})))
			if string(*resultUnknown.Output) == "" {
				t.Errorf("test Failed - incorrect value in response")
			}
			var objmap map[string]interface{}
			_ = json.Unmarshal(*resultUnknown.Transaction, &objmap)
			output, _ := UnknownOutputResp(client.Request(RunUnknownOutput(&ParamsOfDecodeUnknownRun{Abi: *abiS, BodyBase64: (objmap["out_messages"].([]interface{})[0].(map[string]interface{})["body"].(string))})))
			if output.Function != "getVersion" || string(*output.Output) != "{\"value0\":\"0x1\"}" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestProcessTransaction", func(t *testing.T) {
			result, _ := RunResp(client.Request(Run(&ParamsOfRun{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			processed, _ := RunResp(client.Request(ProcessTransaction(&ParamsOfProcessTransaction{Address: address, Transaction: result.Transaction, Abi: *abiS, FunctionName: "getVersion"})))
			if string(*processed.Output) != "{\"value0\":\"0x1\"}" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestWaitTransaction", func(t *testing.T) {
			message, _ := RunMessageResp(client.Request(RunMessage(&ParamsOfRun{Address: address, Abi: *abiS, FunctionName: "getVersion", Input: T{}})))
			state, _ := SendMessageResp(client.Request(SendMessage(message)))
			unknownOutput, _ := RunResp(client.Request(WaitTransaction(&ParamsOfWaitTransaction{Message: *message, MessageProcessingState: *state})))
			if string(*unknownOutput.Output) != "" {
				t.Errorf("test Failed - incorrect value in response")
			}
		})

		t.Run("TestTvmGet", func(t *testing.T) {
			dd := &ParamsOfLocalRunGet{}
			dd.Address = "-1:3333333333333333333333333333333333333333333333333333333333333333"
			dd.FunctionName = "active_election_id"
			result, err := RunLocalResp(client.Request(TvmGet(dd)))
			if err != nil {
				t.Errorf("test Failed - error method TvmGet is function active_election_id: %s", err)
			}

			fmt.Println("Active election id is: ", result.Output)

			dd.FunctionName = "participant_list"
			result, err = RunLocalResp(client.Request(TvmGet(dd)))
			if err != nil {
				t.Errorf("test Failed - error method TvmGet is function participant_list: %s", err)
			}
			fmt.Println("List of participants: ", result.Output)
		})
	})
}

func openFile(fileWay string) ([]byte, error) {
	file, err := os.Open(fileWay)
	if err != nil {
		return nil, errors.New("test Failed - error open abi: " + err.Error())
	}

	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("test Failed - error read abi: " + err.Error())
	}

	return byteContent, nil
}
