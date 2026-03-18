package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type RunJobUseCase struct {
	prompter             ports.Prompter
	awxconnector         ports.AwxConnector
	monitorJobProcessing *services.MonitorJobProgress
	connectToAwx         *services.ConnectToAwx
}

func NewRunJobUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector, monitorJobProcessing *services.MonitorJobProgress, connectToAwx *services.ConnectToAwx) *RunJobUseCase {
	return &RunJobUseCase{
		prompter:             prompter,
		awxconnector:         awxconnector,
		monitorJobProcessing: monitorJobProcessing,
		connectToAwx:         connectToAwx,
	}
}

func (uc *RunJobUseCase) Execute() error {
	var cfg ports.AwxConfig
	if err := uc.connectToAwx.Execute(&cfg); err != nil {
		return err
	}

	templates, err := uc.awxconnector.ListJobTemplates("")
	if err != nil {
		return err
	}

	template, err := uc.prompter.ChooseFromList(templates, "What do you want to do?")
	if err != nil {
		return nil
	}

	jobId, err := uc.awxconnector.LaunchJob(template.Value, nil)
	if err != nil {
		return err
	}

	if _, err := uc.monitorJobProcessing.Execute(jobId); err != nil {
		return err
	}

	return nil
}
