package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type RunTemplateUseCase struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

func NewRunTemplateUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector) *RunTemplateUseCase {
	return &RunTemplateUseCase{
		prompter:     prompter,
		awxconnector: awxconnector,
	}
}

func (uc *RunTemplateUseCase) Execute(templatePrefix string) error {
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

	if _, err := uc.awxconnector.JobProgress(jobId); err != nil {
		return err
	}

	return nil
}
