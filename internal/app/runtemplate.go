package app

import (
	"github.com/goodylabs/awxhelper/internal/services/helpers"
	"github.com/goodylabs/awxhelper/internal/services/ports"
)

type RunTemplateUseCase struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

func NewRestoreBackup(prompter ports.Prompter, awxconnector ports.AwxConnector) *RunTemplateUseCase {
	uc := new(RunTemplateUseCase)
	uc.prompter = prompter
	uc.awxconnector = awxconnector
	return uc
}

func (uc *RunTemplateUseCase) Execute(templatePrefix string) error {
	if err := helpers.ConnectAwx(uc.awxconnector); err != nil {
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

	if err := uc.awxconnector.JobProgress(jobId); err != nil {
		return err
	}

	return nil
}
