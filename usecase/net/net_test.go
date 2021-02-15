package net

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/stretchr/testify/assert"
)

func TestNet(t *testing.T) {

	client, err := client.NewClientGateway(domain.NewDefaultConfig(2))
	assert.Equal(t, nil, err)
	defer client.Destroy()

	netUC := net{
		config: domain.NewDefaultConfig(2),
		client: client,
	}
	defer netUC.client.Destroy()

	t.Run("TestQueryCollection", func(t *testing.T) {
		valueRes1, err := netUC.QueryCollection(&domain.ParamsOfQueryCollection{Collection: "blocks_signatures", Result: "id", Limit: 1})
		assert.Equal(t, nil, err)
		assert.Greater(t, len(valueRes1.Result), 0)

		valueRes2, err := netUC.QueryCollection(&domain.ParamsOfQueryCollection{Collection: "accounts", Result: "id, balance", Limit: 5})
		assert.Equal(t, nil, err)
		assert.Equal(t, 5, len(valueRes2.Result))

		valueRes3, err := netUC.QueryCollection(&domain.ParamsOfQueryCollection{Collection: "messages", Filter: json.RawMessage(`{"created_at":{"gt":1562342740}}`), Result: "body, created_at", Order: []domain.OrderBy{{Path: "created_at", Direction: domain.SortDirectionASC}}, Limit: 10})
		assert.Equal(t, nil, err)
		var (
			objmap    map[string]json.RawMessage
			createdAt int
		)
		err = json.Unmarshal(valueRes3.Result[0], &objmap)
		assert.Equal(t, nil, err)
		bytesRes, err := objmap["created_at"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesRes, &createdAt)
		assert.Equal(t, nil, err)
		assert.Greater(t, createdAt, 1562342740)

		_, err = netUC.QueryCollection(&domain.ParamsOfQueryCollection{Collection: "messages"})
		assert.NotEqual(t, nil, err)
	})

	t.Run("TestWaitCollection", func(t *testing.T) {
		nowTime := int(time.Now().Unix())
		filter := fmt.Sprintf(`{"now":{"gt":%d}}`, nowTime)
		query := &domain.ParamsOfWaitForCollection{Collection: "transactions", Filter: json.RawMessage(filter), Result: "id, now"}
		valueRes1, err := netUC.WaitForCollection(query)
		assert.Equal(t, nil, err)
		var (
			objmap  map[string]json.RawMessage
			dateNow int
		)
		err = json.Unmarshal(valueRes1.Result, &objmap)
		assert.Equal(t, nil, err)
		bytesNow, err := objmap["now"].MarshalJSON()
		assert.Equal(t, nil, err)
		err = json.Unmarshal(bytesNow, &dateNow)
		assert.Equal(t, nil, err)
		assert.Greater(t, dateNow, nowTime)

		query.Timeout = 1
		valueRes2, err := netUC.WaitForCollection(query)
		assert.NotEqual(t, nil, err)
		assert.Equal(t, json.RawMessage(nil), valueRes2.Result)
	})

	t.Run("TestSubscribeCollection", func(t *testing.T) {

		// # Prepare query
		nowTime := int(time.Now().Unix())
		filter := fmt.Sprintf(`{"created_at":{"gt":%d}}`, nowTime)
		query := &domain.ParamsOfSubscribeCollection{Collection: "messages", Filter: json.RawMessage(filter), Result: "created_at"}

		// # Create generator
		generator, handle, err := netUC.SubscribeCollection(query)
		assert.NotNil(t, generator)
		assert.Equal(t, nil, err)
		assert.NotNil(t, handle)

		swG := &sync.WaitGroup{}
		swG.Add(1)

		go func() {
			defer swG.Done()
			respCount := 1
			for g := range generator {
				if respCount > 10 {
					err = netUC.Unsubscribe(&domain.ResultOfSubscribeCollection{Handle: handle.Handle})
					assert.Equal(t, nil, err)
					break
				}
				assert.NotEqual(t, nil, g)
				respCount++
			}
		}()

		swG.Wait()
	})
}
