package client

import (
	"github.com/markgenuine/ever-client-go/domain"
)
import "C"

type SDKResponse struct {
	Result uint32              `json:"result"`
	Error  *domain.ClientError `json:"error"`
}
