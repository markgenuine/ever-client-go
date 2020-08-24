package goton

import (
	"encoding/json"
)

type ParamsOfQuery struct {
	Table   string  `json:"table"`
	Filter  string  `json:"filter"`
	Result  string  `json:"result"`
	OrderBy OrderBy `json:"order_by,omitempty"`
	Limit   int     `json:"limit,omitempty"` //max:50
}

type ParamsOfSubscribe struct {
	Table  string `json:"table"`
	Filter string `json:"filter"`
	Result string `json:"result"`
}

type ParamsOfWaitFor struct {
	Table   string `json:"table"`
	Filter  string `json:"filter"`
	Result  string `json:"result"`
	Timeout int    `json:"timeout"`
}

type OrderBy struct {
	Path      string `json:"path"`
	Direction string `json:"direction"`
}

type HandleStruct struct {
	Handle int `json:"handle"`
}

//Query queries.query
func (client *Client) Query(sq *ParamsOfQuery) (string, error) {
	request, err := json.Marshal(sq)
	if err != nil {
		return "", err
	}

	return client.request("queries.query", string(request))
}

//GetNext queries.get.next
func (client *Client) GetNext(handle int) (string, error) {
	hndl := &HandleStruct{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", err
	}
	return client.request("queries.get.next", string(request))
}

//WaitFor queries.wait.for
func (client *Client) WaitFor(powf *ParamsOfWaitFor) (string, error) {
	request, err := json.Marshal(powf)
	if err != nil {
		return "", err
	}
	return client.request("queries.wait.for", string(request))
}

//Unsubscribe queries.unsubscribe
func (client *Client) Unsubscribe(handle int) (string, error) {
	hndl := &HandleStruct{}
	hndl.Handle = handle
	request, err := json.Marshal(hndl)
	if err != nil {
		return "", err
	}
	return client.request("queries.unsubscribe", string(request))
}

//Subscribe queries.subscribe
func (client *Client) Subscribe(sq *ParamsOfSubscribe) (*HandleStruct, error) {

	request, err := json.Marshal(sq)
	if err != nil {
		return nil, err
	}

	result, err := client.request("queries.subscribe", string(request))
	hndl := &HandleStruct{}
	err = json.Unmarshal([]byte(result), hndl)
	return hndl, err
}
