package services

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/pkg/config"
)

type ConnectToAwx struct {
	fileadapter  ports.FileAdapter
	awxconnector ports.AwxConnector
}

func NewConnectToAwx(fileadapter ports.FileAdapter, awxconnector ports.AwxConnector) *ConnectToAwx {
	return &ConnectToAwx{
		fileadapter:  fileadapter,
		awxconnector: awxconnector,
	}
}

func (c *ConnectToAwx) Execute(cfg *ports.AwxConfig) error {
	fmt.Println("Connecting to AWX...")

	awxCfgPath := filepath.Join(config.GetAwxhelperDir(), "awxhelper-config.json")
	if err := c.fileadapter.ReadJSONFile(awxCfgPath, cfg); err != nil {
		return err
	}

	return c.awxconnector.ConfigureConnection(cfg)
}
