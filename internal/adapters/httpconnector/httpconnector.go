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
	utils.OptionalLog(fmt.Sprintf("--> [DoGet] Start: path='%s'", path))
	fullURL := opts.BaseURL + path
	utils.OptionalLog(fmt.Sprintf("[DoGet] Full request URL: %s", fullURL))

	var body []byte
	var status int
	var err error

	for i := 1; i <= 3; i++ {
		utils.OptionalLog(fmt.Sprintf("[DoGet] Attempt %d/3 for GET %s", i, fullURL))

		utils.OptionalLog(fmt.Sprintf("[DoGet] Creating a new GET request. User: '%s'", opts.Username))
		req, reqErr := http.NewRequest("GET", fullURL, nil)
		if reqErr != nil {
			utils.OptionalLog(fmt.Sprintf("!!! [DoGet] CRITICAL ERROR while creating request object: %s", reqErr))
			return nil, 0, reqErr
		}
		utils.OptionalLog("[DoGet] GET request object successfully created.")

		req.SetBasicAuth(opts.Username, opts.Password)
		utils.OptionalLog("[DoGet] Set Basic Auth headers.")

		utils.OptionalLog(fmt.Sprintf("--- [DoGet] Sending GET request to %s ---", fullURL))
		resp, doErr := h.client.Do(req)
		if doErr != nil {
			utils.OptionalLog(fmt.Sprintf("!!! [DoGet] ERROR while executing request to %s: %v", fullURL, doErr))
			err = doErr
			utils.OptionalLog("[DoGet] Proceeding to the next attempt due to an HTTP client error...")
			continue
		}

		utils.OptionalLog(fmt.Sprintf("<-- [DoGet] Received response from %s", fullURL))
		utils.OptionalLog(fmt.Sprintf("[DoGet] Response status: %s | Status code: %d", resp.Status, resp.StatusCode))
		status = resp.StatusCode

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			utils.OptionalLog(fmt.Sprintf("!!! [DoGet] ERROR while reading response body: %v", err))
			continue
		}

		utils.OptionalLog(fmt.Sprintf("[DoGet] Successfully read response body. Length: %d bytes.", len(body)))
		utils.OptionalLog(fmt.Sprintf("[DoGet] Response body content: %s", string(body)))

		if status < 500 {
			utils.OptionalLog(fmt.Sprintf("[DoGet] Status code %d (< 500) indicates success. Returning result.", status))
			return body, status, nil
		}

		utils.OptionalLog(fmt.Sprintf("[DoGet] Status code %d (>= 500) is a server error. Retrying...", status))
	}

	utils.OptionalLog(fmt.Sprintf("!!! [DoGet] All 3 attempts for GET %s have failed.", fullURL))
	utils.OptionalLog(fmt.Sprintf("<-- [DoGet] End: Returning last error: %v", err))
	return body, status, err
}

func (h *httpconnector) DoPost(opts ports.HttpConnOpts, path string, bodyData any) ([]byte, int, error) {
	utils.OptionalLog(fmt.Sprintf("--> [DoPost] Start: path='%s'", path))
	fullURL := opts.BaseURL + path
	utils.OptionalLog(fmt.Sprintf("[DoPost] Full request URL: %s", fullURL))
	utils.OptionalLog(fmt.Sprintf("[DoPost] Input data (before JSON conversion): %+v", bodyData))

	jsonBody, err := json.Marshal(bodyData)
	if err != nil {
		utils.OptionalLog(fmt.Sprintf("!!! [DoPost] CRITICAL ERROR during body to JSON conversion: %v", err))
		return nil, 0, err
	}
	utils.OptionalLog(fmt.Sprintf("[DoPost] Successfully converted body to JSON: %s", string(jsonBody)))

	utils.OptionalLog(fmt.Sprintf("[DoPost] Creating new POST request for %s", fullURL))
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		utils.OptionalLog(fmt.Sprintf("!!! [DoPost] CRITICAL ERROR while creating request object: %v", err))
		return nil, 0, err
	}
	utils.OptionalLog("[DoPost] POST request object successfully created.")

	req.Header.Set("Content-Type", "application/json")
	utils.OptionalLog("[DoPost] Set 'Content-Type: application/json' header.")
	req.SetBasicAuth(opts.Username, opts.Password)
	utils.OptionalLog("[DoPost] Set Basic Auth headers.")

	utils.OptionalLog(fmt.Sprintf("--- [DoPost] Sending POST request to %s ---", fullURL))
	resp, err := h.client.Do(req)
	if err != nil {
		utils.OptionalLog(fmt.Sprintf("!!! [DoPost] ERROR while executing POST request: %v", err))
		return nil, 0, err
	}
	defer resp.Body.Close()
	utils.OptionalLog(fmt.Sprintf("<-- [DoPost] Received response from %s", fullURL))

	utils.OptionalLog(fmt.Sprintf("[DoPost] Response status: %s | Status code: %d", resp.Status, resp.StatusCode))
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.OptionalLog(fmt.Sprintf("!!! [DoPost] ERROR while reading response body: %v", err))
		return nil, resp.StatusCode, err
	}

	utils.OptionalLog(fmt.Sprintf("[DoPost] Successfully read response body. Length: %d bytes.", len(body)))
	utils.OptionalLog(fmt.Sprintf("[DoPost] Response body content: %s", string(body)))

	utils.OptionalLog(fmt.Sprintf("<-- [DoPost] End: Returning status %d and error: %v", resp.StatusCode, err))
	return body, resp.StatusCode, err
}
