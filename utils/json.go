package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeResponseBody[T any](resp *http.Response) (T, error) {
	var v T
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		err := fmt.Errorf("error decoding response body: %v", err)
		return v, err
	}
	return v, nil
}
