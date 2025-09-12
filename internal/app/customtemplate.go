package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
)

type CustomTemplateUseCase struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

func NewCustomTemplateUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector) *CustomTemplateUseCase {
	return &CustomTemplateUseCase{
		prompter:     prompter,
		awxconnector: awxconnector,
	}
}

func (uc *CustomTemplateUseCase) Execute() error {
	if err := services.ConnectAwx(uc.awxconnector); err != nil {
		return err
	}

	phrase, err := uc.prompter.PromptForString("Filter template name by phrase:")
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

	if _, err := uc.awxconnector.JobProgress(jobId); err != nil {
		return err
	}

	return nil
}
