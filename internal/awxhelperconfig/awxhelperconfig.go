package awxhelperconfig

import (
	"path/filepath"

	"github.com/goodylabs/awxhelper/pkg/config"
	"github.com/goodylabs/awxhelper/pkg/utils"
)

type Config struct {
	LastVersionCheck string `json:"lastVersionCheck"`
	URL              string `json:"url"`
	Username         string `json:"username"`
	Password         string `json:"password"`
}

func RetriveConfig(path string) (*Config, error) {
	config, err := utils.ReadJSON[Config](path)
	if err != err {
		return nil, err
	}
	return &config, err
}

func GetConfigPath() string {
	return filepath.Join(config.AWXHELPER_DIR, "awxhelper-config.json")
}

// func CheckShouldUpdate() bool {
// 	path := GetConfigPath()

// 	config, err := RetriveConfig(path)
// 	if err != nil {
// 		return err
// 	}

// }
