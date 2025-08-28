package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goodylabs/releaser"
	releaserGithub "github.com/goodylabs/releaser/providers/github"
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

		releaserInstance = releaser.ConfigureGithubApp(
			awxhelperDir,
			&releaserGithub.GithubOpts{
				User: "goodylabs",
				Repo: "awxhelper",
			})
	}
	return releaserInstance
}
