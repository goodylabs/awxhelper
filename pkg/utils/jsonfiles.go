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
