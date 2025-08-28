package awxconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/goodylabs/awxhelper/internal/services/ports"
	"github.com/goodylabs/awxhelper/pkg/config"
)

type awxconnector struct {
	httpconnector ports.HttpConnector
	httpCfg       ports.HttpConnOpts
}

func NewAwxConnector(httpconnector ports.HttpConnector) ports.AwxConnector {
	return &awxconnector{
		httpconnector: httpconnector,
	}
}

func (a *awxconnector) ConfigureConnection(cfg *dto.AwxConfig) error {
	a.httpCfg.BaseURL = cfg.URL
	a.httpCfg.Username = cfg.Username
	a.httpCfg.Password = cfg.Password

	respBody, statusCode, err := a.httpconnector.DoGet(a.httpCfg, "/api/v2/ping/")
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
	url := fmt.Sprintf("/api/v2/job_templates/?name__icontains=%s", prefix)
	respBody, statusCode, err := a.httpconnector.DoGet(a.httpCfg, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("failed to list job templates, status %d", statusCode)
	}

	var response struct {
		Results []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			SummaryFilds struct {
				Labels struct {
					Results []string `json:"results"`
				} `json:"labels"`
			} `json:"summary_fields"`
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
	templateIdInt, err := strconv.Atoi(templateId)
	if err != nil {
		return 0, fmt.Errorf("invalid template id: %w", err)
	}

	url := fmt.Sprintf("/api/v2/job_templates/%d/launch/", templateIdInt)

	launchBody := map[string]any{
		"inventory": config.INVENTORY_ID,
	}

	respBody, statusCode, err := a.httpconnector.DoPost(a.httpCfg, url, launchBody)
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
	path := fmt.Sprintf("/api/v2/jobs/%d/job_events?page_size=100", jobId)

	type Event struct {
		Event   string `json:"event"`
		Task    string `json:"task,omitempty"`
		Changed bool   `json:"changed,omitempty"`
		Failed  bool   `json:"failed,omitempty"`
		Created string `json:"created,omitempty"`
	}
	var events []Event

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	for {
		if len(events) > 0 && events[len(events)-1].Event == "playbook_on_stats" {
			s.Stop()
			fmt.Println("Job completed.")
			break
		}

		resBody, statusCode, err := a.httpconnector.DoGet(a.httpCfg, path)
		if err != nil {
			s.Stop()
			log.Fatal("failed to ping AWX:", err)
		}
		if statusCode != 200 {
			s.Stop()
			log.Fatal("failed to ping AWX: unauthorized or AWX not reachable")
		}

		type responseDTO struct {
			Results []Event `json:"results"`
		}
		var resp responseDTO
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			s.Stop()
			return fmt.Errorf("failed to unmarshal job events: %w", err)
		}

		newEvents := resp.Results[len(events):]

		for _, newEvent := range newEvents {
			if newEvent.Event != "runner_on_ok" {
				continue
			}

			var color = "\033[34m"
			if newEvent.Changed {
				color = "\033[32m"
			}
			if newEvent.Failed {
				color = "\033[31m"
			}

			date := strings.Split(newEvent.Created, ".")[0]
			date = strings.Split(date, "T")[1]
			fmt.Printf("%s - %s%s\033[0m\n", date, color, newEvent.Task)
		}

		events = resp.Results

		s.Stop()
		s.Suffix = " Waiting for new job events..."
		s.Start()
		defer s.Stop()

		time.Sleep(5 * time.Second)
	}
	return nil
}
