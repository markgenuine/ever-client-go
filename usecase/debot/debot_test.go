package debot

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/move-ton/ton-client-go/util"
	"io/ioutil"
	"os"
	"testing"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/move-ton/ton-client-go/usecase/abi"
	"github.com/move-ton/ton-client-go/usecase/crypto"
	"github.com/move-ton/ton-client-go/usecase/processing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	config := domain.NewDefaultConfig(0)
	config.Network.ServerAddress = "https://tonos.freeton.surf"
	clientMain, err := client.NewClientGateway(config)
	assert.Equal(t, nil, err)

	debotUC := debot{
		config: config,
		client: clientMain,
	}

	cryptoUC := crypto.NewCrypto(debotUC.config, debotUC.client)
	keypair, err := cryptoUC.GenerateRandomSignKeys()
	assert.Equal(t, nil, err)

	// #Deploy addresses
	fileAbi1, err := os.Open("../samples/DebotTarget.abi.json")
	assert.Equal(t, nil, err)
	byteAbi1, err := ioutil.ReadAll(fileAbi1)
	assert.Equal(t, nil, err)

	targetAbi := domain.AbiContract{}
	err = json.Unmarshal(byteAbi1, &targetAbi)
	assert.Equal(t, nil, err)

	abiValue1 := domain.NewAbiContract()
	abiValue1.Value = targetAbi

	fileAbi2, err := os.Open("../samples/Debot.abi.json")
	assert.Equal(t, nil, err)
	byteAbi2, err := ioutil.ReadAll(fileAbi2)
	assert.Equal(t, nil, err)

	debotAbi := domain.AbiContract{}
	err = json.Unmarshal(byteAbi2, &debotAbi)
	assert.Equal(t, nil, err)

	abiValue2 := domain.NewAbiContract()
	abiValue2.Value = debotAbi

	fileTvc1, err := os.Open("../samples/DebotTarget.tvc")
	assert.Equal(t, nil, err)
	byteTvc1, err := ioutil.ReadAll(fileTvc1)
	assert.Equal(t, nil, err)
	tvcTarget := base64.StdEncoding.EncodeToString(byteTvc1)

	fileTvc2, err := os.Open("../samples/Debot.tvc")
	assert.Equal(t, nil, err)
	byteTvc2, err := ioutil.ReadAll(fileTvc2)
	assert.Equal(t, nil, err)
	tvcDebot := base64.StdEncoding.EncodeToString(byteTvc2)

	// Deploy debot and target
	signer := domain.NewSignerKeys()
	signer.Keys = *keypair

	// # Deploy target
	callSet := domain.CallSet{FunctionName: "constructor"}
	deploySet := domain.DeploySet{Tvc: tvcTarget}

	abiUC := abi.NewAbi(debotUC.config, debotUC.client)
	message, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue1, Signer: signer, DeploySet: &deploySet, CallSet: &callSet})
	assert.Equal(t, nil, err)

	// # Send grams
	fileAbiG, err := os.Open("../samples/Giver.abi.json")
	assert.Equal(t, nil, err)
	byteAbiG, err := ioutil.ReadAll(fileAbiG)
	assert.Equal(t, nil, err)

	eventsAbiG := domain.AbiContract{}
	err = json.Unmarshal(byteAbiG, &eventsAbiG)
	assert.Equal(t, nil, err)

	giverAbi := domain.NewAbiContract()
	giverAbi.Value = eventsAbiG
	callSetN := domain.CallSet{}
	callSetN.FunctionName = "grant"
	callSetN.Input = json.RawMessage(`{"dest":"` + message.Address + `"}`)
	assert.Equal(t, nil, err)
	procUC := processing.NewProcessing(debotUC.config, debotUC.client)
	_, err = procUC.ProcessMessage(domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Type:    "EncodingParams",
			Abi:     giverAbi,
			Signer:  domain.NewSignerNone(),
			Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
			CallSet: &callSetN}, SendEvents: false}, nil)
	assert.Equal(t, nil, err)

	// # Deploy account
	target, err := procUC.ProcessMessage(domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Type:      "EncodingParams",
			Abi:       abiValue1,
			Signer:    signer,
			DeploySet: &deploySet,
			CallSet:   &callSet,
		}, SendEvents: false}, nil)
	assert.Equal(t, nil, err)

	var (
		objmap                      map[string]json.RawMessage
		targerAddress, debotAddress string
	)

	err = json.Unmarshal(target.Transaction, &objmap)
	assert.Equal(t, nil, err)

	err = json.Unmarshal(objmap["account_addr"], &targerAddress)
	assert.Equal(t, nil, err)

	//# Deploy debot
	type inputs1 struct {
		DebotAbi   string `json:"debotAbi"`
		TargetAbi  string `json:"targetAbi"`
		TargetAddr string `json:"targetAddr"`
	}

	//debotAbiByte,err := json.Marshal(debotAbi)
	//assert.Equal(t, nil, err)

	//targetAbiByte,err := json.Marshal(targetAbi)
	//assert.Equal(t, nil, err)

	dataConv := inputs1{DebotAbi: string(util.ToHex(byteAbi1)), TargetAbi: string(util.ToHex(byteAbi2)), TargetAddr: targerAddress}
	inputs1Byte, err := json.Marshal(dataConv)
	assert.Equal(t, nil, err)
	callSet2 := domain.CallSet{}
	callSet2.FunctionName = "constructor"
	callSet2.Input = json.RawMessage(inputs1Byte)

	deploySet2 := domain.DeploySet{Tvc: tvcDebot}
	message2, err := abiUC.EncodeMessage(domain.ParamsOfEncodeMessage{Abi: abiValue2, Signer: signer, DeploySet: &deploySet2, CallSet: &callSet2})
	assert.Equal(t, nil, err)

	// # Send grams
	callSetN.Input = json.RawMessage(`{"dest":"` + message2.Address + `"}`)
	assert.Equal(t, nil, err)
	_, err = procUC.ProcessMessage(domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Type:    "EncodingParams",
			Abi:     giverAbi,
			Signer:  domain.NewSignerNone(),
			Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5",
			CallSet: &callSetN}, SendEvents: false}, nil)
	assert.Equal(t, nil, err)

	// # Deploy account
	debotDate, err := procUC.ProcessMessage(domain.ParamsOfProcessMessage{
		MessageEncodeParams: &domain.ParamsOfEncodeMessage{
			Type:      "EncodingParams",
			Abi:       abiValue2,
			Signer:    signer,
			DeploySet: &deploySet2,
			CallSet:   &callSet2,
		}, SendEvents: false}, nil)
	assert.Equal(t, nil, err)

	err = json.Unmarshal(debotDate.Transaction, &objmap)
	assert.Equal(t, nil, err)

	err = json.Unmarshal(objmap["account_addr"], &debotAddress)
	assert.Equal(t, nil, err)

	fmt.Println("address: ", debotAddress)
	// t.Run("TestDebot", func(t *testing.T) {
	// 	t.Run("TestVersion", func(t *testing.T) {

	// 	})
	// })

	// class TestTonDebotAsyncCore(unittest.TestCase):
	//     def setUp(self) -> None:
	//         self.keypair = async_custom_client.crypto.generate_random_sign_keys()
	//         self.target_abi = Abi.from_path(
	//             path=os.path.join(SAMPLES_DIR, 'DebotTarget.abi.json'))
	//         self.debot_abi = Abi.from_path(
	//             path=os.path.join(SAMPLES_DIR, 'Debot.abi.json'))
	//
	// 			with open(os.path.join(SAMPLES_DIR, 'DebotTarget.tvc'), 'rb') as fp:
	//             self.target_tvc = base64.b64encode(fp.read()).decode()
	//         with open(os.path.join(SAMPLES_DIR, 'Debot.tvc'), 'rb') as fp:
	//             self.debot_tvc = base64.b64encode(fp.read()).decode()

	//         self.debot_address, self.target_address = self.__deploy()

	//     def test_goto(self):
	//         steps = [
	//             {'choice': 0, 'inputs': [], 'outputs': ['Test Goto Action'], 'actions': 1},
	//             {'choice': 0, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     def test_print(self):
	//         steps = [
	//             {'choice': 1, 'inputs': [], 'outputs': ['Test Print Action', 'test2: instant print', 'test instant print'], 'actions': 3},
	//             {'choice': 0, 'inputs': [], 'outputs': ['test simple print'], 'actions': 3},
	//             {'choice': 1, 'inputs': [], 'outputs': [f'integer=1,addr={self.target_address},string=test_string_1'], 'actions': 3},
	//             {'choice': 2, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     def test_run(self):
	//         steps = [
	//             {'choice': 2, 'inputs': '-1:1111111111111111111111111111111111111111111111111111111111111111', 'outputs': ['Test Run Action', 'test1: instant run 1', 'test2: instant run 2'], 'actions': 3},
	//             {'choice': 0, 'inputs': 'hello', 'outputs': [], 'actions': 3},
	//             {'choice': 1, 'inputs': [], 'outputs': ['integer=2,addr=-1:1111111111111111111111111111111111111111111111111111111111111111,string=hello'], 'actions': 3},
	//             {'choice': 2, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     def test_run_method(self):
	//         steps = [
	//             {'choice': 3, 'inputs': [], 'outputs': ['Test Run Method Action'], 'actions': 3},
	//             {'choice': 0, 'inputs': [], 'outputs': [], 'actions': 3},
	//             {'choice': 1, 'inputs': [], 'outputs': ['data=64'], 'actions': 3},
	//             {'choice': 2, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     def test_send_message(self):
	//         steps = [
	//             {'choice': 4, 'inputs': [], 'outputs': ['Test Send Msg Action'], 'actions': 4},
	//             {'choice': 0, 'inputs': [], 'outputs': ['Sending message {}', 'Transaction succeeded.'], 'actions': 4},
	//             {'choice': 1, 'inputs': [], 'outputs': [], 'actions': 4},
	//             {'choice': 2, 'inputs': [], 'outputs': ['data=100'], 'actions': 4},
	//             {'choice': 3, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     def test_invoke(self):
	//         steps = [
	//             {'choice': 5, 'inputs': self.debot_address, 'outputs': ['Test Invoke Debot Action', 'enter debot address:'], 'actions': 2},
	//             {'choice': 0, 'inputs': [], 'outputs': [], 'actions': 2, 'invokes': [
	//                 {'choice': 0, 'inputs': [], 'outputs': ['Print test string', 'Debot is invoked'], 'actions': 0}
	//             ]},
	//             {'choice': 1, 'inputs': [], 'outputs': ['Debot Tests'], 'actions': 8},
	//             {'choice': 7, 'inputs': [], 'outputs': [], 'actions': 0}
	//         ]
	//         asyncio.get_event_loop().run_until_complete(
	//             self._debot_run(steps=steps))

	//     async def _debot_run(
	//             self, steps: List[Dict[str, Any]], start_fn: str = 'start',
	//             actions: List[DebotAction] = None):
	//         # Create initial state
	//         state = {
	//             'handle': None,
	//             'messages': [],
	//             'actions': actions or [],
	//             'steps': steps,
	//             'step': None
	//         }

	//         # Start debot browser and wait for handle
	//         asyncio.get_running_loop().create_task(
	//             self._debot_browser(state=state, start_fn=start_fn))
	//         await self._debot_handle_await(state=state)
	//         self._debot_print_state(state=state)

	//         while len(state['steps']):
	//             step = state['steps'].pop(0)
	//             action = state['actions'][step['choice']]
	//             logging.info(f'[ACTION SELECTED]\t{action}')
	//             state['messages'].clear()
	//             state['step'] = step
	//             await client.debot.execute(
	//                 debot_handle=state['handle'], action=action)
	//             while len(state['messages']) != len(step['outputs']) or \
	//                     len(state['actions']) != step['actions']:
	//                 await asyncio.sleep(1)
	//             self._debot_print_state(state=state)
	//         await client.debot.remove(debot_handle=state['handle'])

	//     async def _debot_browser(self, state: Dict[str, Any], start_fn: str):
	//         generator = getattr(client.debot, start_fn)(address=self.debot_address)
	//         async for event in generator:
	//             data = event['response_data']
	//             # Check for useful data to be present
	//             if data is None:
	//                 continue

	//             # Set state handle when received
	//             if event['response_type'] == TCResponseType.Success:
	//                 state['handle'] = data['debot_handle']

	//             # Process app notifications
	//             if event['response_type'] == TCResponseType.AppNotify:
	//                 params = ParamsOfAppDebotBrowser.from_dict(data=data)
	//                 if isinstance(params, ParamsOfAppDebotBrowser.Log):
	//                     state['messages'].append(params.msg)
	//                 if isinstance(params, ParamsOfAppDebotBrowser.Switch):
	//                     state['actions'].clear()
	//                     if params.context_id == DebotState.EXIT:
	//                         break
	//                 if isinstance(params, ParamsOfAppDebotBrowser.ShowAction):
	//                     state['actions'].append(params.action)

	//             # Process app requests
	//             if event['response_type'] == TCResponseType.AppRequest:
	//                 params = ParamsOfAppDebotBrowser.from_dict(
	//                     data=data['request_data'])
	//                 result = None

	//                 if isinstance(params, ParamsOfAppDebotBrowser.Input):
	//                     result = ResultOfAppDebotBrowser.Input(
	//                         value=state['step']['inputs'])
	//                 if isinstance(params, ParamsOfAppDebotBrowser.GetSigningBox):
	//                     box = await client.crypto.get_signing_box(
	//                         keypair=self.keypair)
	//                     result = ResultOfAppDebotBrowser.GetSigningBox(
	//                         signing_box=box)
	//                 if isinstance(params, ParamsOfAppDebotBrowser.InvokeDebot):
	//                     await self._debot_run(
	//                         steps=state['step']['invokes'], start_fn='fetch',
	//                         actions=[params.action])
	//                     result = ResultOfAppDebotBrowser.InvokeDebot()

	//                 # Resolve app request
	//                 result = AppRequestResult.Ok(result=result.dict)
	//                 await client.resolve_app_request(
	//                     app_request_id=data['app_request_id'], result=result)

	//     @staticmethod
	//     async def _debot_handle_await(state: Dict[str, Any]):
	//         while True:
	//             await asyncio.sleep(0.1)
	//             if state['handle']:
	//                 break

	//     @staticmethod
	//     def _debot_print_state(state: Dict[str, Any]):
	//         # Print messages
	//         for log in state['messages']:
	//             logging.info(f'[LOG]\t{log}')
	//         # Print available actions
	//         for action in state['actions']:
	//             logging.info(f'[ACTION]\t{action}')

	//
}
