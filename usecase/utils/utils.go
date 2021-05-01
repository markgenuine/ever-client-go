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

// ConvertAddress - Converts address from any TON format to any TON format.
func (u *utils) ConvertAddress(pOCA *domain.ParamsOfConvertAddress) (*domain.ResultOfConvertAddress, error) {
	result := new(domain.ResultOfConvertAddress)
	err := u.client.GetResult("utils.convert_address", pOCA, result)
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
