package client

import "github.com/move-ton/ton-client-go/domain"

func (c *clientGateway) Version() (*domain.ResultOfVersion, error) {
	result := new(domain.ResultOfVersion)
	err := c.GetResult("client.version", nil, result)
	return result, err
}

func (c *clientGateway) GetAPIReference() (*domain.ResultOfGetAPIReference, error) {
	result := new(domain.ResultOfGetAPIReference)
	err := c.GetResult("client.get_api_reference", nil, result)
	return result, err
}

func (c *clientGateway) GetBuildInfo() (*domain.ResultOfBuildInfo, error) {
	result := new(domain.ResultOfBuildInfo)
	err := c.GetResult("client.build_info", nil, result)
	return result, err
}
