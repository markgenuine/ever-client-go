package client

import (
	"github.com/move-ton/ever-client-go/domain"
)

// GetAPIReference - Returns Core Library API reference.
func (c *clientGateway) GetAPIReference() (*domain.ResultOfGetAPIReference, error) {
	result := new(domain.ResultOfGetAPIReference)
	err := c.GetResult("client.get_api_reference", nil, result)
	return result, err
}

// Version - Returns Core Library version.
func (c *clientGateway) Version() (*domain.ResultOfVersion, error) {
	result := new(domain.ResultOfVersion)
	err := c.GetResult("client.version", nil, result)
	return result, err
}

// GetBuildInfo - Returns detailed information about this build.
func (c *clientGateway) GetBuildInfo() (*domain.ResultOfBuildInfo, error) {
	result := new(domain.ResultOfBuildInfo)
	err := c.GetResult("client.build_info", nil, result)
	return result, err
}

// ResolveAppRequest - Resolves application request processing result.
func (c *clientGateway) ResolveAppRequest(pORAR *domain.ParamsOfResolveAppRequest) error {
	_, err := c.GetResponse("client.resolve_app_request", pORAR)
	return err
}
