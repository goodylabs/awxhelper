package main

import (
	"log"
	"os"

	"github.com/goodylabs/awxhelper/internal/adapters/awxconnector"
	"github.com/goodylabs/awxhelper/internal/adapters/httpconnector"
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	URL := os.Getenv("AWX_URL")
	USERNAME := os.Getenv("AWX_USER")
	PASSWORD := os.Getenv("AWX_PASSWORD")

	connector := awxconnector.NewAwxConnector(httpconnector.NewHttpConnector())
	if err := connector.ConfigureConnection(&ports.AwxConfig{
		URL:      URL,
		Username: USERNAME,
		Password: PASSWORD,
	}); err != nil {
		log.Fatalf("failed to verify connection: %+v", err)
	}
	if _, err := connector.JobProgress(2804); err != nil {
		log.Fatalf("failed to get job progress: %+v", err)
	}

	// PATH := "/api/v2/jobs/2741/job_events?page_size=100"
	// connector := httpconnector.NewHttpConnector()
	// respBody, statusCode, err := connector.DoGet(ports.HttpConnOpts{
	// 	BaseURL:  URL,
	// 	Username: USERNAME,
	// 	Password: PASSWORD,
	// }, PATH)
	// if err != nil {
	// 	log.Fatal("failed to get job progress: %", err)
	// }
	// if statusCode != 200 {
	// 	log.Fatalf("failed to get job progress, status %d", statusCode)
	// }
	// var prettyJSON map[string]any
	// if err := json.Unmarshal(respBody, &prettyJSON); err != nil {
	// 	log.Fatal("failed to unmarshal json: %", err)
	// }
	// var out []byte
	// out, err = json.MarshalIndent(prettyJSON, "", "  ")
	// if err != nil {
	// 	log.Fatal("failed to marshal json: %", err)
	// }
	// fmt.Println(string(out))
}
