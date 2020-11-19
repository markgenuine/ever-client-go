package client

import (
	"github.com/move-ton/ton-client-go/domain"
)
import "C"

// SDKResponse ...
type SDKResponse struct {
	Result uint32              `json:"result"`
	Error  *domain.ClientError `json:"error"`
}
