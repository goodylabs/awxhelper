package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/goodylabs/awxhelper/internal/services/ports"
	"github.com/goodylabs/awxhelper/pkg/di"
	"github.com/joho/godotenv"
)

func main() {
	container := di.CreateContainer()
	if err := container.Invoke(func(awxconnector ports.AwxConnector) error {
		godotenv.Load(".env")
		awxconnector.ConfigureConnection(&dto.AwxConfig{
			URL:      os.Getenv("AWX_URL"),
			Username: os.Getenv("AWX_USER"),
			Password: os.Getenv("AWX_PASSWORD"),
		})
		items, err := awxconnector.ListJobTemplates("")
		for _, item := range items {
			fmt.Println(item)
		}
		return err
	}); err != nil {
		log.Fatal(err)
	}

}
