package app

import (
	"fmt"

	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type RunDownloadDB struct {
	prompter             ports.Prompter
	awxconnector         ports.AwxConnector
	monitorJobProcessing *services.MonitorJobProgress
	getEndingInstruction *services.GetEndingInstruction
	connectToAwx         *services.ConnectToAwx
}

func NewDownloadDB(
	prompter ports.Prompter,
	awxconnector ports.AwxConnector,
	monitorJobProcessing *services.MonitorJobProgress,
	getEndingInstruction *services.GetEndingInstruction,
	connectToAwx *services.ConnectToAwx,
) *RunDownloadDB {
	return &RunDownloadDB{
		prompter:             prompter,
		awxconnector:         awxconnector,
		monitorJobProcessing: monitorJobProcessing,
		getEndingInstruction: getEndingInstruction,
		connectToAwx:         connectToAwx,
	}
}

func (uc *RunDownloadDB) Execute(templatePrefix string) error {
	var cfg ports.AwxConfig
	if err := uc.connectToAwx.Execute(&cfg); err != nil {
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

	events, err := uc.monitorJobProcessing.Execute(jobId)
	if err != nil {
		return err
	}

	hint, err := uc.getEndingInstruction.DownloadDb(events)
	if err != nil {
		return err
	}

	fmt.Println(hint)

	return nil
}
