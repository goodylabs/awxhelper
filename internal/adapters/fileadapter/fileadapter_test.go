package fileadapter_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/goodylabs/awxhelper/internal/adapters/fileadapter"
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/tests/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileAdapter_WriteJSONFile(t *testing.T) {

	adapter := fileadapter.NewFileAdapter()

	t.Run("Write JSON file successfully", func(t *testing.T) {
		sourceData := ports.AwxConfig{
			URL:      "http://example.com",
			Username: "testuser",
			Password: "testpass",
		}

		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "test.json")
		err := adapter.WriteJSONFile(filePath, sourceData)

		assert.NoError(t, err, "Writing to JSON file should not return an error")
		assert.FileExists(t, filePath, "The JSON file should have been created")

		content, readErr := os.ReadFile(filePath)
		require.NoError(t, readErr, "Reading the written file should not return an error")

		var fileData ports.AwxConfig
		unmarshalErr := json.Unmarshal(content, &fileData)
		require.NoError(t, unmarshalErr, "Unmarshaling the JSON content should not return an error")

		assert.Equal(t, sourceData, fileData, "The data read from the file should match the source data")
	})
}

func TestFileAdapter_ReadJSONFile(t *testing.T) {
	adapter := fileadapter.NewFileAdapter()
	testingDir := testutils.GetTestingDir()

	t.Run("Success - file exists and has valid content", func(t *testing.T) {

		filePath := filepath.Join(testingDir, "adapters", "fileadapter", "1.json")
		var target ports.AwxConfig
		err := adapter.ReadJSONFile(filePath, &target)

		expected := ports.AwxConfig{
			URL:      "https://lorem.ipsum",
			Username: "dolor",
			Password: "sit amet",
		}

		assert.NoError(t, err, "Reading a valid JSON file should not return an error")
		assert.Equal(t, expected, target, "The data read from the file should match the expected data")
	})
}
