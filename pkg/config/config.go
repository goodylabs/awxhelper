package config

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	awxhelperDir string
)

func GetAwxhelperDir() string {
	if awxhelperDir == "" {
		godotenv.Load(".env")
		envValue := getEnvOrError("AWXHELPER_ENV")
		if envValue == "development" {
			rootDir := findProjectRoot()
			awxhelperDir = path.Join(rootDir, ".development")
		} else {
			awxhelperDir = path.Join(getEnvOrError("HOME"), ".awxhelper")
		}
	}
	return awxhelperDir
}

func getEnvOrError(key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return envValue
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("Could not find project root with go.mod file")
		}
		dir = parent
	}
}
