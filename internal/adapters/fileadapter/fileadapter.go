package fileadapter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/goodylabs/awxhelper/internal/ports"
)

type fileAdapter struct{}

func NewFileAdapter() ports.FileAdapter {
	return &fileAdapter{}
}

func (fa *fileAdapter) ReadJSONFile(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("CRITICAL ERROR: Failed to read file: path=%s, error=%v, data=%v", path, err, data)
		return fmt.Errorf("error reading file %s: %w", path, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		log.Printf("CRITICAL ERROR: Failed to unmarshal JSON data from file: path=%s, error=%v, data=%v", path, err, data)
		return fmt.Errorf("error unmarshaling JSON from file %s: %w", path, err)
	}

	return nil
}

func (fa *fileAdapter) WriteJSONFile(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("CRITICAL ERROR: Failed to create file for writing: path=%s, error=%v", path, err)
		return fmt.Errorf("error creating file %s: %w", path, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Printf("CRITICAL ERROR: Failed to encode JSON data to file: path=%s, error=%v, data=%v", path, err, data)
		return fmt.Errorf("error encoding JSON data to file %s: %w", path, err)
	}

	return nil
}
