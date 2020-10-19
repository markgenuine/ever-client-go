package net

import (
	"encoding/json"

	goton "github.com/move-ton/ton-client-go"
)

// QueryCollection method net.query_collection
func QueryCollection(sq *goton.ParamsOfQuery) (string, string) {
	request, err := json.Marshal(sq)
	if err != nil {
		return "", ""
	}

	return "net.query_collection", string(request)
}

// WaitForCollection net.wait_for_collection
func WaitForCollection(pOWFC *ParamsOfWaitForCollection) (string, string) {
	request, err := json.Marshal(pOWFC)
	if err != nil {
		return "", ""
	}
	return "net.wait_for_collection", string(request)
}

// WaitForCollectionResult net.wait_for_collection
func WaitForCollectionResult(resp string, err error) (*ResultOfWaitForCollection, error) {
	if err != nil {
		return nil, err
	}

	rOWFC := &ResultOfWaitForCollection{}
	err = json.Unmarshal([]byte(resp), rOWFC)
	if err != nil {
		return nil, err
	}

	return rOWFC, nil
}

// Unsubscribe net.unsubscribe
func Unsubscribe(handle int) (string, string) {
	hndl := &goton.HandleStruct{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", ""
	}
	return "net.unsubscribe", string(request)
}

// Subscribe net.subscribe
func Subscribe(sq *goton.ParamsOfSubscribe) (string, string) {
	request, err := json.Marshal(sq)
	if err != nil {
		return "", ""
	}

	return "net.subscribe_collection", string(request)
}

// // SubscribeResp response net.subscribe method
// func SubscribeResp(resp string, err error) (*goton.HandleStruct, error) {
// 	if err != nil {
// 		return nil, err
// 	}

// 	hndl := &goton.HandleStruct{}
// 	err = json.Unmarshal([]byte(resp), hndl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return hndl, nil
// }
