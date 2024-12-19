package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c client) makeCall(method, endpoint string, reqObj, resObj interface{}) {
	var req *http.Request
	var err error
	if reqObj != nil {
		jsonBody, err := json.Marshal(reqObj)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		byteBody := bytes.NewBuffer(jsonBody)
		req, err = http.NewRequest(method, endpoint, byteBody)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = json.Unmarshal(responseBody, &resObj)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}
