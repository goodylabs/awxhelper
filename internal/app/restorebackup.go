package app

import (
	services "github.com/goodylabs/awxhelper/internal/services/helpers"
	"github.com/goodylabs/awxhelper/internal/services/ports"
)

type RestoreBackupUseCase struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

func NewRestoreBackup(prompter ports.Prompter, awxconnector ports.AwxConnector) *RestoreBackupUseCase {
	uc := new(RestoreBackupUseCase)
	uc.prompter = prompter
	uc.awxconnector = awxconnector
	return uc
}

func (uc *RestoreBackupUseCase) Execute() error {
	if err := services.ConnectAwx(uc.awxconnector); err != nil {
		return err
	}

	// templates, err := uc.awxconnector.ListJobTemplates("backup_restore")
	templates, err := uc.awxconnector.ListJobTemplates("")
	if err != nil {
		return err
	}

	template, err := uc.prompter.ChooseFromList(templates, "Select action:")

	jobId, err := uc.awxconnector.LaunchJob(template.Value, map[string]any{})
	if err != nil {
		return err
	}

	if err := uc.awxconnector.JobProgress(jobId); err != nil {
		return err
	}

	return nil
}
