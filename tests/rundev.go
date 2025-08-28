package main

import (
	"log"
	"os"

	"github.com/goodylabs/awxhelper/internal/adapters/awxconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/httpconnector"
	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	URL := os.Getenv("AWX_URL")
	USERNAME := os.Getenv("AWX_USER")
	PASSWORD := os.Getenv("AWX_PASSWORD")

	// PATH := "/api/v2/jobs/2702/job_events?page_size=100"
	connector := awxconnector.NewAwxConnector(httpconnector.NewHttpConnector())
	connector.ConfigureConnection(&dto.AwxConfig{
		URL:      URL,
		Username: USERNAME,
		Password: PASSWORD,
	})
	if err := connector.JobProgress(2702); err != nil {
		// if err := connector.JobProgress(2725); err != nil {
		log.Fatal("failed to get job progress: %", err)
	}
}
