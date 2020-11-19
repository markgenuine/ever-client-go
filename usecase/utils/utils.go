package utils

import "github.com/move-ton/ton-client-go/domain"

type utils struct {
	config domain.Config
	client domain.ClientGateway
}

// NewUtils ...
func NewUtils(
	config domain.Config,
	client domain.ClientGateway,
) domain.UtilsUseCase {
	return &utils{
		config: config,
		client: client,
	}
}

func (u *utils) ConverAddress(pOCA domain.ParamsOfConvertAddress) (*domain.ResultOfConvertAddress, error) {
	result := new(domain.ResultOfConvertAddress)
	err := u.client.GetResult("utils.convert_address", pOCA, result)
	return result, err
}
