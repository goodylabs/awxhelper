package services

import (
	"fmt"
	"strings"

	"github.com/goodylabs/awxhelper/internal/ports"
)

func NewGetEndingInstruction() *GetEndingInstruction {
	return new(GetEndingInstruction)
}

type GetEndingInstruction struct{}

const downloadDbTemplate = `
=================================================
To download the database, run the commands below:

curl -L "%s" -o /tmp/%s
gunzip -f /tmp/%s
echo "\nDatabase dump downloaded and unzipped to /tmp/%s"

Important note: the link will expire in 60 minutes!
===================================================
`

func (e *GetEndingInstruction) DownloadDb(events []ports.Event) (string, error) {
	bucketPhrase := "backupus/"
	for _, event := range events {
		if event.Task == "Print url" && event.Event == "runner_on_ok" {

			fmt.Println("ooo")

			url := event.EventData.Res.Msg

			if !strings.Contains(url, bucketPhrase) || !strings.Contains(url, ".gz") {
				return "", fmt.Errorf("Contact with devops team. Can not extract installation script with url: %s", url)
			}
			fileName := strings.Split(event.EventData.Res.Msg, bucketPhrase)[1]
			fileName = strings.Split(fileName, ".gz")[0]
			fileNameGz := fileName + ".gz"

			return fmt.Sprintf(downloadDbTemplate, url, fileNameGz, fileNameGz, fileName), nil
		}
		fmt.Println(event)
	}
	return "", fmt.Errorf("No URL found in the job events - please contact devops team.")
}
