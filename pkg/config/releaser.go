package config

import (
	"fmt"
	"os"
	"path/filepath"

	releaserApi "github.com/goodylabs/releaser/api"
	releaser "github.com/goodylabs/releaser/releaser"
)

var releaserInstance *releaser.ReleaserInstance

func GetReleaser() *releaser.ReleaserInstance {
	if releaserInstance == nil {
		homeDir, err := os.UserHomeDir()
		awxhelperDir := filepath.Join(homeDir, ".awxhelper")
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			os.Exit(1)
		}

		opts := &releaserApi.GithubAppOpts{
			User: "goodylabs",
			Repo: "awxhelper",
		}

		releaserInstance = releaserApi.ConfigureGithubApp(awxhelperDir, opts)
	}
	return releaserInstance
}
