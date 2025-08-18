package awxconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/goodylabs/awxhelper/internal/awxhelperconfig"
	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/goodylabs/awxhelper/internal/services/ports"
)

type awxconnector struct {
	client   *http.Client
	baseURL  string
	username string
	password string
}

func NewAwxConnector() ports.AwxConnector {
	return &awxconnector{
		client: &http.Client{},
	}
}

func (a *awxconnector) ConfigureConnection(cfg *awxhelperconfig.Config) error {
	a.baseURL = cfg.URL
	a.username = cfg.Username
	a.password = cfg.Password

	respBody, statusCode, err := a.doGet("/api/v2/ping/")
	if err != nil {
		return fmt.Errorf("failed to ping AWX: %w", err)
	}
	if statusCode != 200 {
		return errors.New("failed to ping AWX: unauthorized or AWX not reachable")
	}
	_ = respBody
	return nil
}

func (a *awxconnector) ListJobTemplates(prefix string) ([]dto.PrompterItem, error) {
	if a.client == nil {
		return nil, errors.New("AWX client is not configured")
	}
	url := fmt.Sprintf("/api/v2/job_templates/?name__icontains=%s", prefix)
	respBody, statusCode, err := a.doGet(url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("failed to list job templates, status %d", statusCode)
	}

	var response struct {
		Results []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"results"`
	}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	var jobTemplates []dto.PrompterItem
	for _, r := range response.Results {
		jobTemplates = append(jobTemplates, dto.PrompterItem{
			Label: r.Name,
			Value: strconv.Itoa(r.ID),
		})
	}

	return jobTemplates, nil
}

func (a *awxconnector) LaunchJob(templateId string, params map[string]any) (int, error) {
	if a.client == nil {
		return 0, errors.New("AWX client is not configured")
	}
	templateIdInt, err := strconv.Atoi(templateId)
	if err != nil {
		return 0, fmt.Errorf("invalid template id: %w", err)
	}

	url := fmt.Sprintf("/api/v2/job_templates/%d/launch/", templateIdInt)

	launchBody := map[string]any{
		"inventory": 13,
	}

	respBody, statusCode, err := a.doPost(url, launchBody)
	if err != nil {
		return 0, err
	}
	if statusCode != 201 {
		return 0, fmt.Errorf("failed to launch job, status %d: %s", statusCode, string(respBody))
	}

	var response struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal(respBody, &response)
	return response.ID, err
}

func (a *awxconnector) JobProgress(jobId int) error {
	url := fmt.Sprintf("/api/v2/jobs/%d/", jobId)

	start := time.Now()

	for {
		respBody, statusCode, err := a.doGet(url)
		if err != nil {
			return fmt.Errorf("failed to get job status: %w", err)
		}
		if statusCode != 200 {
			return fmt.Errorf("unexpected status code %d when getting job status", statusCode)
		}

		var job struct {
			Status string `json:"status"`
			Failed bool   `json:"failed"`
		}
		err = json.Unmarshal(respBody, &job)
		if err != nil {
			return fmt.Errorf("failed to unmarshal job status json: %w", err)
		}

		elapsed := time.Since(start).Round(time.Second)

		fmt.Print("\033[H\033[2J")
		fmt.Printf("Job ID %d running for %v, status: %s\n", jobId, elapsed, job.Status)

		switch job.Status {
		case "canceled", "failed":
			fmt.Printf("Job ID %d failed or errored\n", jobId)
			return nil
		case "successful":
			fmt.Printf("Job ID %d completed successfully\n", jobId)
			return nil
		}

		time.Sleep(5 * time.Second)
	}
}
