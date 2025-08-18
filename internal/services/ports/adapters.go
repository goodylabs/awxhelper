package ports

import (
	"github.com/goodylabs/awxhelper/internal/awxhelperconfig"
	"github.com/goodylabs/awxhelper/internal/services/dto"
)

type Prompter interface {
	ChooseFromList([]dto.PrompterItem, string) (dto.PrompterItem, error)
	PromptForString(message string) (string, error)
}

type AwxConnector interface {
	ConfigureConnection(cfg *awxhelperconfig.Config) error
	ListJobTemplates(prefix string) ([]dto.PrompterItem, error)
	LaunchJob(templateId string, params map[string]any) (int, error)
	JobProgress(jobId int) error
}
