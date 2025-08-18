package awxconnector

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (a *awxconnector) doGet(path string) ([]byte, int, error) {
	fullURL := a.baseURL + path

	var body []byte
	var status int
	var err error

	for range 3 {
		req, reqErr := http.NewRequest("GET", fullURL, nil)
		if reqErr != nil {
			return nil, 0, reqErr
		}
		req.SetBasicAuth(a.username, a.password)

		resp, doErr := a.client.Do(req)
		if doErr != nil {
			err = doErr
			continue
		}

		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		status = resp.StatusCode

		if err == nil && status < 500 {
			return body, status, nil
		}
	}

	return body, status, err
}

func (a *awxconnector) doPost(path string, bodyData any) ([]byte, int, error) {
	fullURL := a.baseURL + path
	jsonBody, err := json.Marshal(bodyData)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.username, a.password)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}
