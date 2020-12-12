package debot

import (
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
func (d *debot) Start(poS domain.ParamsOfStart) (*domain.RegisteredDebot, error) { //?callback 2 params  obj: AppDebotBrowser,
	result := new(domain.RegisteredDebot)
	err := d.client.GetResult("debot.start", poS, result)
	return result, err
}

// Fetch method debot.fetch
func (d *debot) Fetch(pOF domain.ParamsOfFetch) (*domain.RegisteredDebot, error) { //?callback 2 params  obj: AppDebotBrowser,
	result := new(domain.RegisteredDebot)
	err := d.client.GetResult("debot.fetch", pOF, result)
	return result, err
}

// Execute method debot.execute
func (d *debot) Execute(pOE domain.ParamsOfExecute) error {
	_, err := d.client.GetResponse("debot.execute", pOE)
	return err
}

// Remove method debot.remove
func (d *debot) Remove(pOE domain.RegisteredDebot) error {
	_, err := d.client.GetResponse("debot.remove", pOE)
	return err
}
