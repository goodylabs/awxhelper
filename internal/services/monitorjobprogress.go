package services

import (
	"fmt"
	"os"
	"os/signal"
	"slices"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/goodylabs/awxhelper/internal/ports"
)

type MonitorJobProgress struct {
	awxconnector ports.AwxConnector
}

func NewMonitorJobProgress(awxconnector ports.AwxConnector) *MonitorJobProgress {
	return &MonitorJobProgress{
		awxconnector: awxconnector,
	}
}

var successStatuses = []string{"successful"}
var failureStatuses = []string{"failed", "canceled", "error"}
var pendingStatuses = []string{"pending", "waiting", "running"}

func (p *MonitorJobProgress) Execute(jobId int) ([]ports.Event, error) {
	var events []ports.Event

	var s *spinner.Spinner

	for {

		s = startSpinner(fmt.Sprintf("Job %d in progress...", jobId))
		defer stopSpinner(s)

		time.Sleep(3 * time.Second)

		allEvents, err := p.awxconnector.GetJobEvents(jobId)
		if err != nil {
			return []ports.Event{}, fmt.Errorf("failed to get job progress: %w", err)
		}

		stopSpinner(s)

		newEvents := allEvents[len(events):]
		for _, newEvent := range newEvents {
			if newEvent.Event != "runner_on_ok" {
				continue
			}
			p.printEvent(&newEvent)
		}

		if len(newEvents) == 0 {
			continue
		}

		events = allEvents

		status := newEvents[0].SummaryFields.Job.Status

		if slices.Contains(pendingStatuses, status) {
			continue
		}

		if slices.Contains(successStatuses, status) {
			return allEvents, nil
		}

		if slices.Contains(failureStatuses, status) {
			return allEvents, fmt.Errorf("Job stopped with status: '%s'", status)
		}

		return allEvents, fmt.Errorf("unknown job status: '%s'", status)
	}
}

func (p *MonitorJobProgress) printEvent(newEvent *ports.Event) {
	color := "\033[34m"
	if newEvent.Changed {
		color = "\033[32m"
	}
	if newEvent.Failed {
		color = "\033[31m"
	}
	date := newEvent.Created
	date = strings.Split(date, ".")[0]
	if strings.Contains(date, "T") {
		date = strings.Split(date, "T")[1]
	}
	fmt.Printf("%s - %s%s\033[0m\n", date, color, newEvent.Task)
}

func startSpinner(msg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + msg
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

func stopSpinner(s *spinner.Spinner) {
	s.Stop()
	fmt.Print("\r\033[K")
}
