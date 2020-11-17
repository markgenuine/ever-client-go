package client

import (
	"sync"
)

var (
	// State temp store for response
	State = &state{&sync.Mutex{}, map[int]*AsyncResponse{}}
)

type (
	// AsyncResponse ...
	AsyncResponse struct {
		ReqID        int
		MethodName   string
		Params       string
		ResponseType int
		Finished     bool
	}

	state struct {
		*sync.Mutex
		stores map[int]*AsyncResponse
	}
)
