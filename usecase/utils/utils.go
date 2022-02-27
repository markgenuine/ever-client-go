package utils

import "github.com/move-ton/ever-client-go/domain"

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

// ConvertAddress - Converts address from any Ever format to any Ever format.
func (u *utils) ConvertAddress(pOCA *domain.ParamsOfConvertAddress) (*domain.ResultOfConvertAddress, error) {
	result := new(domain.ResultOfConvertAddress)
	err := u.client.GetResult("utils.convert_address", pOCA, result)
	return result, err
}

// GetAddressType - Validates and returns the type of any Ever address.
// Address types are the following
// 0:919db8e740d50bf349df2eea03fa30c385d846b991ff5542e67098ee833fc7f7 - standart Ever address most commonly used in all cases. Also called as hex addres 919db8e740d50bf349df2eea03fa30c385d846b991ff5542e67098ee833fc7f7 - account ID. A part of full address. Identifies account inside particular workchain EQCRnbjnQNUL80nfLuoD+jDDhdhGuZH/VULmcJjugz/H9wam - base64 address. Also called "user-friendly". Was used at the beginning of Ever. Now it is supported for compatibility
func (u *utils) GetAddressType(pOGAT *domain.ParamsOfGetAddressType) (*domain.ResultOfGetAddressType, error) {
	result := new(domain.ResultOfGetAddressType)
	err := u.client.GetResult("utils.get_address_type", pOGAT, result)
	return result, err
}

// CalcStorageFee - Calculates storage fee for an account over a specified time period.
func (u *utils) CalcStorageFee(pOCA *domain.ParamsOfCalcStorageFee) (*domain.ResultOfCalcStorageFee, error) {
	result := new(domain.ResultOfCalcStorageFee)
	err := u.client.GetResult("utils.calc_storage_fee", pOCA, result)
	return result, err
}

// CompressZstd - Compresses data using Zstandard algorithm.
func (u *utils) CompressZstd(pOCA *domain.ParamsOfCompressZstd) (*domain.ResultOfCompressZstd, error) {
	result := new(domain.ResultOfCompressZstd)
	err := u.client.GetResult("utils.compress_zstd", pOCA, result)
	return result, err
}

// DecompressZstd - Decompresses data using Zstandard algorithm.
func (u *utils) DecompressZstd(pOCA *domain.ParamsOfDecompressZstd) (*domain.ResultOfDecompressZstd, error) {
	result := new(domain.ResultOfDecompressZstd)
	err := u.client.GetResult("utils.decompress_zstd", pOCA, result)
	return result, err
}
