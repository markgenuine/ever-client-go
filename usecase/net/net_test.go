package net

import (
	"encoding/json"
	"fmt"
	"github.com/move-ton/ton-client-go/util"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/move-ton/ton-client-go/domain"
	"github.com/move-ton/ton-client-go/gateway/client"
	"github.com/stretchr/testify/assert"
)

func TestNet(t *testing.T) {

	config := domain.NewDefaultConfig("", domain.GetDevNetBaseUrls())
	clientConn, err := client.NewClientGateway(config)
	assert.Equal(t, nil, err)
	defer clientConn.Destroy()

	netUC := net{
		config: config,
		client: clientConn,
	}
	defer netUC.client.Destroy()

	t.Run("TestQueryCollection", func(t *testing.T) {
		queryParams := &domain.ParamsOfQueryCollection{Collection: "blocks_signatures", Result: "id", Limit: util.IntToPointerInt(1)}
		result, err := netUC.QueryCollection(queryParams)
		assert.Equal(t, nil, err)
		assert.Greater(t, len(result.Result), 0)

		queryParams = &domain.ParamsOfQueryCollection{Collection: "accounts", Result: "id, balance", Limit: util.IntToPointerInt(5)}
		result, err = netUC.QueryCollection(queryParams)
		assert.Equal(t, nil, err)
		assert.Equal(t, 5, len(result.Result))

		queryParams = &domain.ParamsOfQueryCollection{Collection: "messages", Filter: json.RawMessage(`{"created_at":{"gt":1562342740}}`), Result: "body, created_at", Order: []*domain.OrderBy{{Path: "created_at", Direction: domain.SortDirectionASC}}, Limit: util.IntToPointerInt(10)}
		result, err = netUC.QueryCollection(queryParams)
		assert.Equal(t, nil, err)

		type resultStruct struct {
			Body      string `json:"body"`
			CreatedAt int    `json:"created_at"`
		}

		res := &resultStruct{}
		err = json.Unmarshal(result.Result[0], &res)
		assert.Equal(t, nil, err)
		assert.Greater(t, res.CreatedAt, 1562342740)

		_, err = netUC.QueryCollection(&domain.ParamsOfQueryCollection{Collection: "messages"})
		assert.NotEqual(t, nil, err)
	})

	t.Run("TestWaitCollection", func(t *testing.T) {
		nowTime := int(time.Now().Unix())
		queryParams := &domain.ParamsOfWaitForCollection{
			Collection: "transactions",
			Filter:     json.RawMessage(fmt.Sprintf(`{"now":{"gt":%d}}`, nowTime)),
			Result:     "id,now",
		}
		result, err := netUC.WaitForCollection(queryParams)
		assert.Equal(t, nil, err)

		type resultStruct struct {
			ID  string `json:"id"`
			Now int    `json:"now"`
		}

		res := &resultStruct{}
		err = json.Unmarshal(result.Result, res)
		assert.Equal(t, nil, err)
		assert.Greater(t, res.Now, nowTime)

		queryParams = &domain.ParamsOfWaitForCollection{
			Collection: "transactions",
			Result:     "",
			Timeout:    util.IntToPointerInt(1),
		}
		result, err = netUC.WaitForCollection(queryParams)
		assert.NotEqual(t, nil, err)
		assert.Equal(t, json.RawMessage(nil), result.Result)
	})

	t.Run("TestSubscribeCollection", func(t *testing.T) {

		// # Prepare query
		nowTime := int(time.Now().Unix())
		queryParams := &domain.ParamsOfSubscribeCollection{
			Collection: "messages",
			Filter:     json.RawMessage(fmt.Sprintf(`{"created_at":{"gt":%d}}`, nowTime)),
			Result:     "created_at",
		}

		// # Create generator
		generator, handle, err := netUC.SubscribeCollection(queryParams)
		assert.NotNil(t, generator)
		assert.Equal(t, nil, err)
		assert.NotNil(t, handle)

		swG := &sync.WaitGroup{}
		swG.Add(1)

		var slResult []json.RawMessage
		go func() {
			defer swG.Done()
			respCount := 1
			for g := range generator {
				slResult = append(slResult, g)
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
		assert.Greater(t, len(slResult), 0)
	})

	t.Run("TestQuery", func(t *testing.T) {
		variables := make(map[string]int)
		variables["time"] = int(time.Now().Unix()) - 60
		varBytes, err := json.Marshal(variables)
		assert.Equal(t, nil, err)
		queryParams := &domain.ParamsOfQuery{
			Query:     "query($time: Float){messages(filter:{created_at:{ge:$time}}limit:5){id}}",
			Variables: json.RawMessage(varBytes),
		}
		result, err := netUC.Query(queryParams)
		assert.Equal(t, nil, err)

		type resultStruct struct {
			Data struct {
				Messages []map[string]string `json:"messages"`
			} `json:"data"`
		}

		res := &resultStruct{}
		err = json.Unmarshal(result.Result, res)
		assert.Equal(t, nil, err)
		assert.Greater(t, len(res.Data.Messages), 0)
	})

	t.Run("TestFindLastShardBlock", func(t *testing.T) {
		findParams := &domain.ParamsOfFindLastShardBlock{Address: "0:b61cf024cda7dad90e556d0fafb72c08579d5ebf73a67737317d9f3fc73521c5"}
		result, err := netUC.FindLastShardBlock(findParams)
		assert.Equal(t, nil, err)
		assert.NotNil(t, result.BlockID)
	})

	t.Run("TestAggregateCollection", func(t *testing.T) {
		field := &domain.FieldAggregation{Field: "", Fn: domain.AggregationFnTypeCount}
		var fields []*domain.FieldAggregation
		fields = append(fields, field)
		params := &domain.ParamsOfAggregateCollection{
			Collection: "accounts",
			Fields:     fields}
		result, err := netUC.AggregateCollection(params)
		assert.Equal(t, nil, err)

		var resSl []string
		err = json.Unmarshal(result.Values, &resSl)
		resToInt, err := strconv.Atoi(resSl[0])
		assert.Equal(t, nil, err)
		assert.Greater(t, resToInt, 0)
	})
}
