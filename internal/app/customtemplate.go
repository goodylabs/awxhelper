package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type CustomTemplateUseCase struct {
	prompter             ports.Prompter
	awxconnector         ports.AwxConnector
	monitorJobProcessing *services.MonitorJobProgress
	connectToAwx         *services.ConnectToAwx
}

func NewCustomTemplateUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector, monitorJobProcessing *services.MonitorJobProgress, connectToAwx *services.ConnectToAwx) *CustomTemplateUseCase {
	return &CustomTemplateUseCase{
		prompter:             prompter,
		awxconnector:         awxconnector,
		monitorJobProcessing: monitorJobProcessing,
		connectToAwx:         connectToAwx,
	}
}

func (uc *CustomTemplateUseCase) Execute() error {
	var cfg ports.AwxConfig
	if err := uc.connectToAwx.Execute(&cfg); err != nil {
		return err
	}

	phrase, err := uc.prompter.PromptForString("Filter template name by phrase")
	if err != nil {
		return err
	}

	templates, err := uc.awxconnector.ListJobTemplates(phrase)
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
