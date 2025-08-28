package app

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/internal/services"
	"github.com/goodylabs/awxhelper/pkg/utils"
)

type ConfigureUseCase struct {
	prompter     ports.Prompter
	awxconnector ports.AwxConnector
}

type ConfigureOpts struct {
	URL      string
	Username string
	Password string
}

func NewConfigureUseCase(prompter ports.Prompter, awxconnector ports.AwxConnector) *ConfigureUseCase {
	uc := new(ConfigureUseCase)
	uc.prompter = prompter
	uc.awxconnector = awxconnector
	return uc
}

func (uc *ConfigureUseCase) Execute(opts *ConfigureOpts) error {
	var cfg ports.AwxConfig
	var err error

	cfg.URL, err = uc.getOrPrompt(opts.URL, "Enter AWX url")
	if err != nil {
		return err
	}

	cfg.Username, err = uc.getOrPrompt(opts.Username, "Enter AWX username")
	if err != nil {
		return err
	}

	cfg.Password, err = uc.getOrPrompt(opts.Password, "Enter AWX password")
	if err != nil {
		return err
	}

	if err = uc.awxconnector.ConfigureConnection(&cfg); err != nil {
		return err
	}

	configPath := services.GetConfigPath()
	return utils.WriteJSON(configPath, cfg)
}

func (uc *ConfigureUseCase) getOrPrompt(value, prompt string) (string, error) {
	if value != "" {
		return value, nil
	}
	return uc.prompter.PromptForString(prompt)
}
