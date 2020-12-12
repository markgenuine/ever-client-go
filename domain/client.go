package domain

import (
	"encoding/json"
	"fmt"
)

type (

	// ClientError ...
	ClientError struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}

	// ClientResponse ...
	ClientResponse struct {
		Code  int
		Data  []byte
		Error error
	}

	// AppRequestResult ...
	AppRequestResult struct {
		Type   string      `json:"type"`
		Text   string      `json:"text,omitempty"`
		Result interface{} `json:"result,omitempty"`
	}

	// ResultOfVersion ...
	ResultOfVersion struct {
		Version string `json:"version"`
	}

	// ResultOfGetAPIReference ...
	ResultOfGetAPIReference struct {
		API API `json:"api"`
	}

	// API ...
	API struct {
		Modules []struct {
			Description string `json:"description"`
			Functions   []struct {
				Description interface{}   `json:"description"`
				Errors      interface{}   `json:"errors"`
				Name        string        `json:"name"`
				Params      []interface{} `json:"params"`
				Result      struct {
					Ref string `json:"ref"`
				} `json:"result"`
				Summary interface{} `json:"summary"`
			} `json:"functions"`
			Name    string `json:"name"`
			Summary string `json:"summary"`
			Types   []struct {
				Description interface{} `json:"description"`
				Name        string      `json:"name"`
				Struct      []struct {
					Description interface{} `json:"description"`
					Name        string      `json:"name"`
					Ref         string      `json:"ref"`
					Summary     interface{} `json:"summary"`
				} `json:"struct"`
				Summary interface{} `json:"summary"`
			} `json:"types"`
		} `json:"modules"`
		Version string `json:"version"`
	}

	// ResultOfBuildInfo ...
	ResultOfBuildInfo struct {
		BuildNumber  int                   `json:"build_number"`
		Dependencies []BuildInfoDependency `json:"dependencies"`
	}

	// BuildInfoDependency ...
	BuildInfoDependency struct {
		Name      string `json:"name"`
		GitCommit string `json:"git_commit"`
	}

	// ParamsOfResolveAppRequest ...
	ParamsOfResolveAppRequest struct {
		AppRequestID int              `json:"app_request_id"`
		Result       AppRequestResult `json:"result"`
	}

	// ClientGateway ...
	ClientGateway interface {
		Destroy()
		GetResult(method string, paramIn interface{}, resultStruct interface{}) error
		Request(method string, paramsIn interface{}) (<-chan *ClientResponse, error)
		GetResponse(method string, paramIn interface{}) ([]byte, error)
		Version() (*ResultOfVersion, error)
		GetAPIReference() (*ResultOfGetAPIReference, error)
		GetBuildInfo() (*ResultOfBuildInfo, error)
		ResolveAppRequest(*ParamsOfResolveAppRequest) error
	}
)

// DynBufferForResponses ...
func DynBufferForResponses(in <-chan *ClientResponse) <-chan *ClientResponse {
	out := make(chan *ClientResponse, 1)
	var storage []*ClientResponse
	go func() {
		defer close(out)
		for {
			if len(storage) == 0 {
				item, ok := <-in
				if !ok {
					return
				}
				storage = append(storage, item)

				continue
			}

			select {
			case item, ok := <-in:
				if ok {
					storage = append(storage, item)
				} else {
					for _, item := range storage {
						out <- item
					}

					return
				}
			case out <- storage[0]:
				if len(storage) == 1 {
					storage = nil
				} else {
					storage = storage[1:]
				}
			}
		}
	}()

	return out
}

// HandleEvents ...
func HandleEvents(responses <-chan *ClientResponse, callback EventCallback, result interface{}) error {
	for r := range responses {
		switch r.Code {
		case 100:
			event := &ProcessingEvent{}
			if err := json.Unmarshal(r.Data, event); err != nil {
				panic(err)
			}
			callback(event)
		case 1:
			return r.Error
		case 0:
			if err := json.Unmarshal(r.Data, result); err != nil {
				panic(err)
			}

			return nil
		default:
			panic(fmt.Errorf("unknown response type code %v", r.Code))
		}
	}

	return nil
}
