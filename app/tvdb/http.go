package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c client) makeCall(method, endpoint string, reqObj, resObj interface{}) error {
	var req *http.Request
	var err error
	if reqObj != nil {
		jsonBody, err := json.Marshal(reqObj)
		if err != nil {
			return fmt.Errorf("error marshaling JSON (%s): %s", jsonBody, err)
		}
		byteBody := bytes.NewBuffer(jsonBody)
		req, err = http.NewRequest(method, endpoint, byteBody)
		if err != nil {
			return fmt.Errorf("error creating request: %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}

	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	defer func() { _ = resp.Body.Close() }()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err)
	}

	err = json.Unmarshal(responseBody, &resObj)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON (%s): %s", responseBody, err)
	}

	return nil
}
