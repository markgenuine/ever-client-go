package net

import (
	"encoding/json"
)

// QueryCollection method net.query_collection
func QueryCollection(pOQC *ParamsOfQueryCollection) (string, string) {
	request, err := json.Marshal(pOQC)
	if err != nil {
		return "", ""
	}

	return "net.query_collection", string(request)
}

// QueryCollectionResult ...
func QueryCollectionResult(resp string, err error) (*ResultOfQueryCollection, error) {
	if err != nil {
		return nil, err
	}

	resStrt := &ResultOfQueryCollection{}
	err = json.Unmarshal([]byte(resp), resStrt)
	if err != nil {
		return nil, err
	}

	return resStrt, nil
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
	hndl := &ResultOfSubscribeCollection{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", ""
	}
	return "net.unsubscribe", string(request)
}

// SubscribeCollection method net.subscribe_collection
func SubscribeCollection(pOSC *ParamsOfSubscribeCollection) (string, string) {
	request, err := json.Marshal(pOSC)
	if err != nil {
		return "", ""
	}

	return "net.subscribe_collection", string(request)
}

// SubscribeCollectionResult response net.subscribe_collection method
func SubscribeCollectionResult(resp string, err error) (*ResultOfSubscribeCollection, error) {
	if err != nil {
		return nil, err
	}

	hndl := &ResultOfSubscribeCollection{}
	err = json.Unmarshal([]byte(resp), hndl)
	if err != nil {
		return nil, err
	}

	return hndl, nil
}
