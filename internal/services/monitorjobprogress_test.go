package services_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/awxhelper/internal/services"
	"github.com/goodylabs/awxhelper/tests/mocks"
	"github.com/goodylabs/awxhelper/tests/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMonitorJobProgress(t *testing.T) {

	resFilesDir := testutils.GetAwxHelperEventsDir()

	t.Run("return events on finish", func(t *testing.T) {
		svc := services.NewMonitorJobProgress(&mocks.MockAwxConnector{
			GetJobEventsResponseFile: filepath.Join(resFilesDir, "canceled_1.json"),
		})

		events, err := svc.Execute(123)
		assert.ErrorContains(t, err, "Job stopped with status: 'canceled'")

		assert.Len(t, events, 11)
		assert.Equal(t, "career reduce especially", events[len(events)-1].Task)
	})

	t.Run("return events on finish", func(t *testing.T) {
		svc := services.NewMonitorJobProgress(&mocks.MockAwxConnector{
			GetJobEventsResponseFile: filepath.Join(resFilesDir, "failed_1.json"),
		})

		events, err := svc.Execute(123)
		assert.ErrorContains(t, err, "Job stopped with status: 'failed'")

		assert.Len(t, events, 11)
	})

}
