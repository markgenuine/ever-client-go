package queries

import (
	"encoding/json"
	"testing"

	goton "github.com/move-ton/go-ton-sdk"
)

func TestQueriesMethod(t *testing.T) {

	client, err := goton.InitClient(goton.NewConfig(0))
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

	t.Run("TestQuery", func(t *testing.T) {
		sq := &goton.ParamsOfQuery{}
		sq.Table = "messages"
		sq.Filter = "{\"dst\":{\"eq\":\"0:668b5c83056ebf1852cc7af4e61c8a421056c0311f035a39e5baf7ce28b14728\"}}"
		sq.Result = "id msg_type status"
		sq.OrderBy.Path = "created_at"
		sq.OrderBy.Direction = "ASC"
		sq.Limit = 1
		result, _ := client.Request(Query(sq))
		type resultStruct struct {
			Result []struct {
				ID      string `json:"id"`
				MsgType int    `json:"msg_type"`
				Status  int    `json:"status"`
			} `json:"result"`
		}

		dd := resultStruct{}
		err = json.Unmarshal([]byte(result), &dd)

		if dd.Result[0].ID != "0ab35fd0b5ccb4dc0598f5cdf64364b7d0ad5fa8bf292d8368ddd498c8e90070" {
			t.Errorf("test Query Failed")
		}
	})

	t.Run("TestWaitFor", func(t *testing.T) {
		powf := &goton.ParamsOfWaitFor{}
		powf.Table = "accounts"
		powf.Filter = "{\"id\":{\"eq\":\"0:668b5c83056ebf1852cc7af4e61c8a421056c0311f035a39e5baf7ce28b14728\"}}"
		powf.Result = "balance"
		powf.Timeout = 5
		result, _ := client.Request(WaitFor(powf))

		type resssS struct {
			Result struct {
				Balance string
			}
		}

		rr := &resssS{}
		json.Unmarshal([]byte(result), rr)

		if rr.Result.Balance != "0x12a3873d4" {
			t.Errorf("test WaitFor Failed")
		}

		t.Run("TestSubscribeAndEvent", func(t *testing.T) {
			powf := &goton.ParamsOfSubscribe{}
			powf.Table = "blocks"
			powf.Filter = "{\"workchain_id\":{\"eq\":0}}"
			powf.Result = "id gen_utime"

			//subscribe
			subscr, _ := SubscribeResp(client.Request(Subscribe(powf)))
			if subscr.Handle == 0 {
				t.Errorf("test subscribe in Sunbsrcibe and Event is Failed")
			}

			//get next
			for i := 0; i < 10; i++ {
				result, err = client.Request(GetNext(subscr.Handle))
				if err != nil {
					t.Errorf("test get next in Sunbscribe and Event is Failed")
				}
			}

			// unsubscribe
			_, err = client.Request(Unsubscribe(subscr.Handle))
			if err != nil {
				t.Errorf("test unsubscribe in Sunbsrcibe and Event is Failed")
			}
		})
	})
}
