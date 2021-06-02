package debot

import (
	"encoding/json"
	"fmt"
	"github.com/move-ton/ton-client-go/domain"
)

type debot struct {
	config domain.Config
	client domain.ClientGateway
}

// NewDebot ...
func NewDebot(config domain.Config, client domain.ClientGateway) domain.DebotUseCase {
	return &debot{
		config: config,
		client: client,
	}
}

// Init - Creates and instance of DeBot.
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

// appRequestDebotInit ...
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
	if err != nil {
		panic(err)
	}
}

// appNotifyDebotInit ...
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

//Start - Starts the DeBot.
func (d *debot) Start(poS *domain.ParamsOfStart) error {
	_, err := d.client.GetResponse("debot.start", poS)
	return err
}

// Fetch - Fetches DeBot metadata from blockchain.
func (d *debot) Fetch(pOF *domain.ParamsOfFetch) (*domain.ResultOfFetch, error) {
	result := new(domain.ResultOfFetch)
	err := d.client.GetResult("debot.fetch", pOF, result)
	return result, err
}

// Execute - Executes debot action.
func (d *debot) Execute(pOE *domain.ParamsOfExecute) error {
	_, err := d.client.GetResponse("debot.execute", pOE)
	return err
}

// Send - Sends message to Debot.
func (d *debot) Send(pOS *domain.ParamsOfSend) error {
	_, err := d.client.GetResponse("debot.send", pOS)
	return err
}

// Remove - Destroys debot handle.
func (d *debot) Remove(pOR *domain.ParamsOfRemove) error {
	_, err := d.client.GetResponse("debot.remove", pOR)
	return err
}
