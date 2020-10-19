package net

import "encoding/json"

// type ParamsOfQueryCollection struct {
//     Collection string
//     filter?: any
//     result: string
//     order?: OrderBy[]
//     limit?: number
// };

// type ResultOfQueryCollection struct {
//     result: any[]
// };

// type OrderBy = {
//     path: string,
//     direction: SortDirection
// };

// type SortDirection = 'ASC' | 'DESC';

// ParamsOfWaitForCollection ...
type ParamsOfWaitForCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Timeout    int         `json:"timeout,omitempty"`
}

// ResultOfWaitForCollection ...
type ResultOfWaitForCollection struct {
	Result json.RawMessage `json:"result"`
}

// type ResultOfSubscribeCollection = {
//     handle: number
// };

// type unit = void;

// type ParamsOfSubscribeCollection = {
//     collection: string,
//     filter?: any,
//     result: string
// };
