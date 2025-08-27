package ports

import (
	"github.com/goodylabs/awxhelper/internal/services/dto"
)

type Prompter interface {
	ChooseFromList([]dto.PrompterItem, string) (dto.PrompterItem, error)
	PromptForString(message string) (string, error)
}

type AwxConnector interface {
	ConfigureConnection(cfg *dto.AwxConfig) error
	ListJobTemplates(prefix string) ([]dto.PrompterItem, error)
	LaunchJob(templateId string, params map[string]any) (int, error)
	JobProgress(jobId int) error
}

type HttpConnOpts struct {
	BaseURL  string
	Username string
	Password string
}

type HttpConnector interface {
	DoGet(opts HttpConnOpts, path string) ([]byte, int, error)
	DoPost(opts HttpConnOpts, path string, bodyData any) ([]byte, int, error)
}
