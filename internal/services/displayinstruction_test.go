package services_test

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/goodylabs/awxhelper/internal/services"
	"github.com/goodylabs/awxhelper/tests/mocks"
	"github.com/goodylabs/awxhelper/tests/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetEndingInstruction(t *testing.T) {

	resFilesDir := testutils.GetAwxHelperEventsDir()

	awxconnector := &mocks.MockAwxConnector{
		GetJobEventsResponseFile: filepath.Join(resFilesDir, "canceled_1.json"),
	}

	events, err := awxconnector.GetJobEvents(123)
	if err != nil {
		log.Fatal(err)
	}

	svc := services.NewGetEndingInstruction()

	hint, err := svc.DownloadDb(events)
	assert.NoError(t, err)
	assert.Contains(t, hint, "ipsum.gz")
}
