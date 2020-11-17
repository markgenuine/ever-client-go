package util

import "encoding/json"

// StructToJSON ...
func StructToJSON(structValue interface{}) (string, error) {
	request, err := json.Marshal(structValue)
	if err != nil {
		return "", err
	}
	return string(request), nil
}
