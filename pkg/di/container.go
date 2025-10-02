package di

import (
	"github.com/goodylabs/awxhelper/internal/adapters/awxconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/httpconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/prompter"
	"github.com/goodylabs/awxhelper/internal/app"
	"github.com/goodylabs/awxhelper/internal/services"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	container := dig.New()

	container.Provide(httpconnector.NewHttpConnector)

	container.Provide(prompter.NewPrompter)
	container.Provide(awxconnector.NewAwxConnector)

	container.Provide(services.NewMonitorJobProgress)
	container.Provide(services.NewGetEndingInstruction)

	container.Provide(app.NewConfigureUseCase)
	container.Provide(app.NewRunTemplateUseCase)
	container.Provide(app.NewDownloadDB)
	container.Provide(app.NewCustomTemplateUseCase)

	return container
}
