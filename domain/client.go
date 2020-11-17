package domain

import (
	"encoding/json"
)

type (

	// ClientError ...
	ClientError struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
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

	// ClientGateway ...
	ClientGateway interface {
		Request(method string, paramsIn interface{}) (int, error)
		Version() (int, error)
		GetAPIReference() (int, error)
		GetBuildInfo() (int, error)
		GetResp(resp int) (interface{}, error)
		Destroy()
	}
)
