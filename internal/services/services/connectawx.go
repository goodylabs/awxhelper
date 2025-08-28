package services

import (
	"fmt"

	"github.com/goodylabs/awxhelper/internal/services/ports"
)

func ConnectAwx(awxconnector ports.AwxConnector) error {
	fmt.Println("Connecting to AWX...")
	path := GetConfigPath()
	cfg, err := RetriveConfig(path)
	if err != nil {
		return err
	}
	return awxconnector.ConfigureConnection(cfg)
}
