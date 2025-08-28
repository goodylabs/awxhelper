package utils

import (
	"encoding/json"
	"os"
)

func ReadJSON[T any](path string) (T, error) {
	var result T

	data, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	return result, nil
}

func WriteJSON[T any](path string, data T) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
