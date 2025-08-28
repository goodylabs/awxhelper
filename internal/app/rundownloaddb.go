package app

import (
	"fmt"
	"strings"

	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type RunDownloadDB struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

func NewDownloadDB(prompter ports.Prompter, awxconnector ports.AwxConnector) *RunDownloadDB {
	return &RunDownloadDB{
		prompter:     prompter,
		awxconnector: awxconnector,
	}
}

func (uc *RunDownloadDB) Execute(templatePrefix string) error {
	if err := services.ConnectAwx(uc.awxconnector); err != nil {
		return err
	}

	templates, err := uc.awxconnector.ListJobTemplates(templatePrefix)
	if err != nil {
		return err
	}

	template, err := uc.prompter.ChooseFromList(templates, "What do you want to do?")
	if err != nil {
		return nil
	}

	jobId, err := uc.awxconnector.LaunchJob(template.Value, map[string]any{})
	if err != nil {
		return err
	}

	events, err := uc.awxconnector.JobProgress(jobId)
	if err != nil {
		return err
	}

	uc.DisplayInstruction(events)

	return nil
}

const downloadDbHint = `
=================================================
To download the database, run the commands below:

curl -L %s -o /tmp/%s
gunzip -f /tmp/%s
echo "\nDatabase dump downloaded and unzipped to /tmp/%s"

Important note: the link will expire in 60 minutes!
===================================================
`

func (uc *RunDownloadDB) DisplayInstruction(events []ports.Event) {
	for _, event := range events {
		if event.Task == "Print url" && event.Event == "runner_on_ok" {

			url := event.EventData.Res.Msg

			fileName := strings.Split(event.EventData.Res.Msg, "archivus/")[1]
			fileName = strings.Split(fileName, ".gz")[0]
			fileNameGz := fileName + ".gz"

			fmt.Printf(downloadDbHint, url, fileNameGz, fileNameGz, fileName)
			return
		}
	}
	fmt.Println("No URL found in the job events - please contact devops team.")
}
