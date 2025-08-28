package awxconnector

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
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

func (a *awxconnector) ConfigureConnection(cfg *ports.AwxConfig) error {
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

func (a *awxconnector) ListJobTemplates(prefix string) ([]ports.PrompterItem, error) {
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

	var jobTemplates []ports.PrompterItem
	for _, r := range response.Results {
		jobTemplates = append(jobTemplates, ports.PrompterItem{
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

func (a *awxconnector) JobProgress(jobId int) ([]ports.Event, error) {
	path := fmt.Sprintf("/api/v2/jobs/%d/job_events?page_size=100", jobId)

	var events []ports.Event

	s := a.setupSpinner()
	s.Start()
	defer func() {
		s.Stop()
		fmt.Print("\r\033[K") // usuwa spinner z ostatniej linii
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		s.Stop()
		fmt.Print("\r\033[K")
		os.Exit(1)
	}()

	for {
		if len(events) > 0 && events[len(events)-1].Event == "playbook_on_stats" {
			s.Stop()
			fmt.Print("\r\033[K")
			fmt.Println("Job completed.")
			break
		}

		type eventRes struct {
			Results []ports.Event `json:"results"`
		}
		var resp eventRes
		var err error
		for range 3 {
			var resBody []byte
			var statusCode int
			resBody, statusCode, connectorErr := a.httpconnector.DoGet(a.httpCfg, path)
			if connectorErr != nil {
				err = fmt.Errorf("failed to ping AWX: %w", connectorErr)
				continue
			}
			if statusCode != 200 {
				err = fmt.Errorf("failed to ping AWX: unauthorized or AWX not reachable")
				continue
			}

			err = json.Unmarshal(resBody, &resp)
			if err != nil {
				err = fmt.Errorf("failed to unmarshal job events: %w", err)
				continue
			}
			err = nil
			break
		}
		if err != nil {
			return []ports.Event{}, err
		}

		newEvents := resp.Results[len(events):]
		for _, newEvent := range newEvents {
			if newEvent.Event != "runner_on_ok" {
				continue
			}

			s.Stop()
			fmt.Print("\r\033[K") // czyści linię spinnera

			color := "\033[34m"
			if newEvent.Changed {
				color = "\033[32m"
			}
			if newEvent.Failed {
				color = "\033[31m"
			}

			date := strings.Split(newEvent.Created, ".")[0]
			date = strings.Split(date, "T")[1]
			fmt.Printf("%s - %s%s\033[0m\n", date, color, newEvent.Task)

			s.Start() // spinner wraca na dół
		}

		events = resp.Results
		time.Sleep(5 * time.Second)
	}
	return events, nil
}

func (a *awxconnector) setupSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Waiting for new job events..."
	s.Start()

	go func(sp *spinner.Spinner) {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		sp.Stop()
		fmt.Print("\r")
		os.Exit(1)
	}(s)

	return s
}
