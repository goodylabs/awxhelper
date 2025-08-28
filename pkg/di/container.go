package di

import (
	"github.com/goodylabs/awxhelper/internal/adapters/awxconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/httpconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/prompter"
	"github.com/goodylabs/awxhelper/internal/app"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	container := dig.New()

	container.Provide(httpconnector.NewHttpConnector)

	container.Provide(prompter.NewPrompter)
	container.Provide(awxconnector.NewAwxConnector)

	container.Provide(app.NewConfigureUseCase)
	container.Provide(app.NewRunTemplateUseCase)
	container.Provide(app.NewDownloadDB)

	return container
}
