package helpers

import (
	"path/filepath"

	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/goodylabs/awxhelper/pkg/config"
	"github.com/goodylabs/awxhelper/pkg/utils"
)

func RetriveConfig(path string) (*dto.AwxConfig, error) {
	config, err := utils.ReadJSON[dto.AwxConfig](path)
	if err != err {
		return nil, err
	}
	return &config, err
}

func GetConfigPath() string {
	return filepath.Join(config.AWXHELPER_DIR, "awxhelper-config.json")
}
