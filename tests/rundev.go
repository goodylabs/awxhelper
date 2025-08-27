package main

import (
	"fmt"
	"os"

	"github.com/goodylabs/awxhelper/internal/adapters/awxconnector"
	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/joho/godotenv"
)

func main() {
	instance := awxconnector.NewAwxConnector()
	godotenv.Load(".env")
	instance.ConfigureConnection(&dto.AwxConfig{
		URL:      os.Getenv("AWX_URL"),
		Username: os.Getenv("AWX_USER"),
		Password: os.Getenv("AWX_PASSWORD"),
	})
	items, err := instance.ListJobTemplates("")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
