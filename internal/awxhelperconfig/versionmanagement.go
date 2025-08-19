package awxhelperconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const owner = "goodylabs"
const repo = "awxhelper"
const latestReleaseUrl = "https://api.github.com/repos/" + owner + "/" + repo + "/releases/latest"

type githubReleaseResponse struct {
	Name string `json:"name"`
}

func GetLatestReleaseName() (string, error) {
	resp, err := http.Get(latestReleaseUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data *githubReleaseResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if data.Name == "" {
		return "", errors.New("release name not found in response")
	}

	return data.Name, nil
}
