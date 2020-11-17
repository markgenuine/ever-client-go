package domain

import "encoding/json"

const (
	// SortDirectionASC ...
	SortDirectionASC = "ASC"

	// SortDirectionDESC ...
	SortDirectionDESC = "DESC"
)

type (
	// OrderBy ...
	OrderBy struct {
		Path      string `json:"path"`
		Direction string `json:"direction"`
	}

	// ParamsOfQueryCollection ...
	ParamsOfQueryCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
		Order      []OrderBy   `json:"order,omitempty"`
		Limit      int         `json:"limit,omitempty"`
	}

	// ResultOfQueryCollection ...
	ResultOfQueryCollection struct {
		Result []json.RawMessage `json:"result"`
	}

	// ParamsOfWaitForCollection ...
	ParamsOfWaitForCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
		Timeout    int         `json:"timeout,omitempty"`
	}

	// ResultOfWaitForCollection ...
	ResultOfWaitForCollection struct {
		Result json.RawMessage `json:"result"`
	}

	// ResultOfSubscribeCollection ...
	ResultOfSubscribeCollection struct {
		Handle int `json:"handle"`
	}

	// ParamsOfSubscribeCollection ...
	ParamsOfSubscribeCollection struct {
		Collection string      `json:"collection"`
		Filter     interface{} `json:"filter,omitempty"`
		Result     string      `json:"result"`
	}

	// NetUseCase ...
	NetUseCase interface {
		QueryCollection(ParamsOfQueryCollection) (int, error)
		WaitForCollection(ParamsOfWaitForCollection) (int, error)
		Unsubscribe(ResultOfSubscribeCollection) (int, error)
		SubscribeCollection(ParamsOfSubscribeCollection) (int, error)
	}
)

func getSortDirection() map[int]string {
	return map[int]string{
		0: "ASC",
		1: "DESC",
	}
}
