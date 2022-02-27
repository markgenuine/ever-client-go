package debot

import (
	"github.com/move-ton/ever-client-go/domain"
	"github.com/move-ton/ever-client-go/gateway/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebot(t *testing.T) {

	config := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
	clientConn, err := client.NewClientGateway(config)
	assert.Equal(t, nil, err)
	defer clientConn.Destroy()

	debotUC := debot{
		config: config,
		client: clientConn,
	}
	defer debotUC.client.Destroy()
}
