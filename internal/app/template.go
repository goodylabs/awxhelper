package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type RunTemplateUseCase struct {
	prompter             ports.Prompter
	awxconnector         ports.AwxConnector
	monitorJobProcessing *services.MonitorJobProgress
	fileadapter          ports.FileAdapter
	connectToAwx         *services.ConnectToAwx
}

func NewRunTemplateUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector, monitorJobProcessing *services.MonitorJobProgress, fileadapter ports.FileAdapter, connectToAwx *services.ConnectToAwx) *RunTemplateUseCase {
	return &RunTemplateUseCase{
		prompter:             prompter,
		awxconnector:         awxconnector,
		monitorJobProcessing: monitorJobProcessing,
		fileadapter:          fileadapter,
		connectToAwx:         connectToAwx,
	}
}

func (uc *RunTemplateUseCase) Execute(templatePrefix string) error {
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

	if _, err := uc.monitorJobProcessing.Execute(jobId); err != nil {
		return err
	}

	return nil
}
