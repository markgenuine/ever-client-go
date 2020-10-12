package queries

import (
	"encoding/json"

	goton "github.com/move-ton/ton-client-go"
)

// Query queries.query
func Query(sq *goton.ParamsOfQuery) (string, string) {
	request, err := json.Marshal(sq)
	if err != nil {
		return "", ""
	}

	return "queries.query", string(request)
}

// GetNext queries.get.next
func GetNext(handle int) (string, string) {
	hndl := &goton.HandleStruct{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", ""
	}
	return "queries.get.next", string(request)
}

// WaitFor queries.wait.for
func WaitFor(powf *goton.ParamsOfWaitFor) (string, string) {
	request, err := json.Marshal(powf)
	if err != nil {
		return "", ""
	}
	return "queries.wait.for", string(request)
}

// Unsubscribe queries.unsubscribe
func Unsubscribe(handle int) (string, string) {
	hndl := &goton.HandleStruct{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", ""
	}
	return "queries.unsubscribe", string(request)
}

// Subscribe queries.subscribe
func Subscribe(sq *goton.ParamsOfSubscribe) (string, string) {

	request, err := json.Marshal(sq)
	if err != nil {
		return "", ""
	}

	return "queries.subscribe", string(request)

}

// SubscribeResp response queries.subscribe method
func SubscribeResp(resp string, err error) (*goton.HandleStruct, error) {
	if err != nil {
		return nil, err
	}

	hndl := &goton.HandleStruct{}
	err = json.Unmarshal([]byte(resp), hndl)
	if err != nil {
		return nil, err
	}

	return hndl, nil
}
