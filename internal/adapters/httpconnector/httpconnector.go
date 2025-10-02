package httpconnector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/pkg/utils"
)

type httpconnector struct {
	client *http.Client
}

func NewHttpConnector() ports.HttpConnector {
	return &httpconnector{
		client: &http.Client{},
	}
}

func (h *httpconnector) DoGet(opts ports.HttpConnOpts, path string) ([]byte, int, error) {
	fullURL := opts.BaseURL + path

	var body []byte
	var status int
	var err error

	for range 3 {
		utils.OptionalLog(fmt.Sprintf("GET request %s username: %s, password: %s", fullURL, opts.Username, opts.Password))
		req, reqErr := http.NewRequest("GET", fullURL, nil)
		if reqErr != nil {
			utils.OptionalLog(fmt.Sprintf("GET response reqError: %s", reqErr))
			return nil, 0, reqErr
		}
		req.SetBasicAuth(opts.Username, opts.Password)

		resp, doErr := h.client.Do(req)
		if doErr != nil {
			utils.OptionalLog(fmt.Sprintf("Error making GET request to %s: %v", fullURL, doErr))
			err = doErr
			continue
		}

		body, err = io.ReadAll(resp.Body)
		status = resp.StatusCode

		utils.OptionalLog(fmt.Sprintf("GET response status: %s", resp.Status))

		if err == nil && status < 500 {
			return body, status, nil
		}
	}

	return body, status, err
}

func (h *httpconnector) DoPost(opts ports.HttpConnOpts, path string, bodyData any) ([]byte, int, error) {
	fullURL := opts.BaseURL + path
	jsonBody, err := json.Marshal(bodyData)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(opts.Username, opts.Password)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}
