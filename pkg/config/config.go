package config

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	AWXHELPER_DIR string
)

func LoadConfig() {
	godotenv.Load(".env")
	switch os.Getenv("AWXHELPER_ENV") {
	case "development":
		rootDir := findProjectRoot()
		AWXHELPER_DIR = path.Join(rootDir, ".development")
	default:
		AWXHELPER_DIR = path.Join(os.Getenv("HOME"), ".awxhelper")
	}
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
