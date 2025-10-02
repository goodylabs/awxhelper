package awxconnector

import (
	"encoding/json"
	"fmt"

	"github.com/goodylabs/awxhelper/internal/ports"
)

type jobEventsResponse struct {
	Results []ports.Event `json:"results"`
}

func (a *awxconnector) GetJobEvents(jobId int) ([]ports.Event, error) {
	path := fmt.Sprintf("/api/v2/jobs/%d/job_events?page_size=100", jobId)

	var err error

	for range 3 {
		resBody, statusCode, connectorErr := a.httpconnector.DoGet(a.httpCfg, path)

		if connectorErr != nil {
			err = fmt.Errorf("failed to ping AWX: %w", connectorErr)
			continue
		}
		if statusCode != 200 {
			err = fmt.Errorf("failed to ping AWX: unauthorized or AWX not reachable")
			continue
		}

		var response jobEventsResponse
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal job events: %w", err)
			continue
		}

		return response.Results, nil
	}

	return []ports.Event{}, err
}
