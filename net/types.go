package net

const (
	SortDirectionASC  = "ASC"
	SortDirectionDESC = "DESC"
)

type ParamsOfQueryCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Order      []OrderBy   `json:"order,omitempty"`
	Limit      int         `json:"limit,omitempty"`
}

type ResultOfQueryCollection struct {
	Result []interface{} `json:"result"`
}

type OrderBy struct {
	Path      string `json:"path"`
	Direction string `json:"direction"`
}

// ParamsOfWaitForCollection ...
type ParamsOfWaitForCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Timeout    int         `json:"timeout,omitempty"`
}

// ResultOfWaitForCollection ...
type ResultOfWaitForCollection struct {
	Result interface{} `json:"result"`
}

type ResultOfSubscribeCollection struct {
	Handle int `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
}
