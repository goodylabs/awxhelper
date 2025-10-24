package testutils

import (
	"os"
	"path/filepath"
)

func GetProjectDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("could not get working directory")
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			panic("could not find project root")
		}
		dir = parent
	}
}

func GetAwxHelperEventsDir() string {
	projectDir := GetProjectDir()
	return filepath.Join(projectDir, "tests", "mocks", "awxconnector", "events")
}

func GetTestingDir() string {
	projectDir := GetProjectDir()
	return filepath.Join(projectDir, ".testing")
}
