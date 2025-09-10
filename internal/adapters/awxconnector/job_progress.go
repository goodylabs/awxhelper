package awxconnector

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/goodylabs/awxhelper/internal/ports"
)

func (a *awxconnector) JobProgress(jobId int) ([]ports.Event, error) {
	path := fmt.Sprintf("/api/v2/jobs/%d/job_events?page_size=100", jobId)

	var events []ports.Event

	s := a.setupSpinner(jobId)
	s.Start()
	defer func() {
		s.Stop()
		fmt.Print("\r\033[K")
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
			fmt.Print("\r\033[K")

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

			s.Start()
		}

		events = resp.Results

		if len(events) > 0 && events[len(events)-1].SummaryFields.Job.Status != "running" {
			s.Stop()
			fmt.Print("\r\033[K")
			fmt.Println("Job completed.")
			break
		}

		time.Sleep(5 * time.Second)
	}
	return events, nil
}

func (a *awxconnector) setupSpinner(jobId int) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Processing job %d...", jobId)
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
