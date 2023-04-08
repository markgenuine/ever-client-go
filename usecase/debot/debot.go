package debot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/move-ton/ever-client-go/domain"
)

type debot struct {
	config domain.ClientConfig
	client domain.ClientGateway
}

func NewDebot(config domain.ClientConfig, client domain.ClientGateway) domain.DebotUseCase {
	return &debot{
		config: config,
		client: client,
	}
}

// Init - Creates and instance of DeBot.
// Downloads debot smart contract (code and data) from blockchain and creates an instance of Debot Engine for it.
func (d *debot) Init(pOI *domain.ParamsOfInit, app domain.AppDebotBrowser) (*domain.RegisteredDebot, error) {
	result := new(domain.RegisteredDebot)
	responses, err := d.client.Request("debot.init", pOI)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == 1 {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == 3 {
				d.appRequestDebotInit(r.Data, app)
			}
			if r.Code == 4 {
				d.appNotifyDebotInit(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (d *debot) appRequestDebotInit(payload []byte, app domain.AppDebotBrowser) {
	var appRequest domain.ParamsOfAppRequest
	var appParams domain.ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	var appResponse interface{}

	switch value := (appParams.ValueEnumType).(type) {
	case domain.ParamsOfAppDebotBrowserInput:
		appResponse, err = app.Input(value)
	case domain.ParamsOfAppDebotBrowserGetSigningBox:
		appResponse, err = app.GetSigningBox(value)
	case domain.ParamsOfAppDebotBrowserInvokeDebot:
		appResponse, err = app.InvokeDebot(value)
	case domain.ParamsOfAppDebotBrowserApprove:
		appResponse, err = app.Approve(value)
	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.ValueEnumType)
	}

	appRequestResult := &domain.AppRequestResult{}
	if err != nil {
		appRequestResult.ValueEnumType = domain.AppRequestResultError{Text: err.Error()}
	} else {
		marshal, _ := json.Marshal(&domain.ResultOfAppDebotBrowser{ValueEnumType: appResponse})
		appRequestResult.ValueEnumType = domain.AppRequestResultOk{Result: marshal}
	}
	paramsResolved := &domain.ParamsOfResolveAppRequest{AppRequestID: appRequest.AppRequestID, Result: appRequestResult}
	err = d.client.ResolveAppRequest(paramsResolved)
	if err != nil || errors.Is(err, errors.New("channels is closed")) {
		return
	}
	panic(err)
}

func (d *debot) appNotifyDebotInit(payload []byte, app domain.AppDebotBrowser) {
	var appParams domain.ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}

	switch value := (appParams.ValueEnumType).(type) {
	case domain.ParamsOfAppDebotBrowserLog:
		_ = app.Log(value)
	case domain.ParamsOfAppDebotBrowserSwitch:
		_ = app.Switch(value)
	case domain.ParamsOfAppDebotBrowserSwitchCompleted:
		_ = app.SwitchCompleted(value)
	case domain.ParamsOfAppDebotBrowserShowAction:
		_ = app.ShowAction(value)
	case domain.ParamsOfAppDebotBrowserSend:
		_ = app.Send(value)
	default:
		panic(fmt.Errorf("unsupported type for request %v", appParams.ValueEnumType))
	}
}

// Start - Starts the DeBot.
// Downloads debot smart contract from blockchain and switches it to context zero.
// This function must be used by Debot Browser to start a dialog with debot. While the function is executing, several
// Browser Callbacks can be called, since the debot tries to display all actions from the context 0 to the user.
// When the debot starts SDK registers BrowserCallbacks AppObject. Therefore when debote.remove is called the debot is
// being deleted and the callback is called with finish=true which indicates that it will never be used again.
func (d *debot) Start(poS *domain.ParamsOfStart) error {
	_, err := d.client.GetResponse("debot.start", poS)
	return err
}

// Fetch - Fetches DeBot metadata from blockchain.
// Downloads DeBot from blockchain and creates and fetches its metadata.
func (d *debot) Fetch(pOF *domain.ParamsOfFetch) (*domain.ResultOfFetch, error) {
	result := new(domain.ResultOfFetch)
	err := d.client.GetResult("debot.fetch", pOF, result)
	return result, err
}

// Execute - Executes debot action.
// Calls debot engine referenced by debot handle to execute input action. Calls Debot Browser Callbacks if needed.
func (d *debot) Execute(pOE *domain.ParamsOfExecute) error {
	_, err := d.client.GetResponse("debot.execute", pOE)
	return err
}

// Send - Sends message to Debot.
// Used by Debot Browser to send response on Dinterface call or from other Debots.
func (d *debot) Send(pOS *domain.ParamsOfSend) error {
	_, err := d.client.GetResponse("debot.send", pOS)
	return err
}

// Remove - Destroys debot handle.
// Removes handle from Client Context and drops debot engine referenced by that handle.
func (d *debot) Remove(pOR *domain.ParamsOfRemove) error {
	_, err := d.client.GetResponse("debot.remove", pOR)
	return err
}
