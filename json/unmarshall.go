package json

import (
	"bytes"
	"encoding/json"
	"godmr/radioid"
	"net/http"
)

type Results struct {
	Count   int               `json:"count"`
	Results []radioid.Contact `json:"results"`
}

func ReadJSON(response *http.Response) ([]radioid.Contact, error) {
	var results Results
	buffer := new(bytes.Buffer)
	_, _ = buffer.ReadFrom(response.Body)

	respByte := buffer.Bytes()
	if err := json.Unmarshal(respByte, &results); err != nil {
		return nil, err
	}

	return results.Results, nil
}
