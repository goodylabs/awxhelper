package mocks

import (
	"log"

	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/pkg/utils"
)

type MockAwxConnector struct {
	GetJobEventsResponseFile string
}

func (m *MockAwxConnector) ConfigureConnection(cfg *ports.AwxConfig) error {
	return nil
}

func (m *MockAwxConnector) ListJobTemplates(prefix string) ([]ports.PrompterItem, error) {
	return []ports.PrompterItem{}, nil
}

func (m *MockAwxConnector) LaunchJob(templateId string, params map[string]any) (int, error) {
	return 0, nil
}

func (m *MockAwxConnector) GetJobEvents(jobId int) ([]ports.Event, error) {
	events, err := utils.ReadJSON[[]ports.Event](m.GetJobEventsResponseFile)
	if err != nil {
		log.Fatal(err)
	}
	return events, nil
}
