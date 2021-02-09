package debot

import (
	"encoding/json"
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

// Start method debot.start
func (d *debot) Start(poS *domain.ParamsOfStart, app domain.AppDebotBrowser) (*domain.RegisteredDebot, error) {
	result := new(domain.RegisteredDebot)
	responses, err := d.client.Request("debot.start", poS)
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
				d.appRequestDebotStart(r.Data, app)
			}
			if r.Code == 4 {
				d.appNotifyDebotStart(r.Data, app)
			}
		}
	}()

	return result, nil
}

//appRequestDebotStart ...
func (d *debot) appRequestDebotStart(payload []byte, app domain.AppDebotBrowser) {
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
	appResponse, err := app.Request(appParams)
	appRequestResult := domain.AppRequestResult{}
	if err != nil {
		appRequestResult.Type = "Error"
		appRequestResult.Text = err.Error()
	} else {
		appRequestResult.Type = "Ok"
		appRequestResult.Result, _ = json.Marshal(appResponse)
	}
	err = d.client.ResolveAppRequest(&domain.ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

//appNotifyDebotStart ...
func (d *debot) appNotifyDebotStart(payload []byte, app domain.AppDebotBrowser) {
	var appParams domain.ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// Fetch method debot.fetch
func (d *debot) Fetch(pOF *domain.ParamsOfFetch, app domain.AppDebotBrowser) (*domain.RegisteredDebot, error) {
	result := new(domain.RegisteredDebot)
	responses, err := d.client.Request("debot.fetch", pOF)
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
				d.appRequestDebotFetch(r.Data, app)
			}
			if r.Code == 4 {
				d.appNotifyDebotFetch(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (d *debot) appRequestDebotFetch(payload []byte, app domain.AppDebotBrowser) {
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
	appResponse, err := app.Request(appParams)
	appRequestResult := domain.AppRequestResult{}
	if err != nil {
		appRequestResult.Type = "Error"
		appRequestResult.Text = err.Error()
	} else {
		appRequestResult.Type = "Ok"
		appRequestResult.Result, _ = json.Marshal(appResponse)
	}
	err = d.client.ResolveAppRequest(&domain.ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (d *debot) appNotifyDebotFetch(payload []byte, app domain.AppDebotBrowser) {
	var appParams domain.ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// Execute method debot.execute
func (d *debot) Execute(pOE *domain.ParamsOfExecute) error {
	_, err := d.client.GetResponse("debot.execute", pOE)
	return err
}

// Remove method debot.remove
func (d *debot) Remove(pOE *domain.RegisteredDebot) error {
	_, err := d.client.GetResponse("debot.remove", pOE)
	return err
}
