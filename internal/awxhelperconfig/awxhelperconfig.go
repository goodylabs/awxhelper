package awxhelperconfig

import "github.com/goodylabs/awxhelper/pkg/utils"

type AwxHelperConfig struct {
	LastVersionCheck string `json:"lastVersionCheck"`
	AwxToken         string `json:"awxToken"`
}

func RetriveAwxHelperConfig(path string) (*AwxHelperConfig, error) {
	config, err := utils.ReadJSON[AwxHelperConfig](path)
	if err != err {
		return nil, err
	}
	return &config, err
}
