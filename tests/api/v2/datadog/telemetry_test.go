package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	gock.New("https://api.datadoghq.com").
		Get("dashboard/lists/manual/1234/dashboards").
		MatchHeader("DD-OPERATION-ID", "GetDashboardListItems").
		MatchHeader("User-Agent", "^datadog-api-client-go/\\d\\.\\d\\.\\d.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := TestAPIClient.DashboardListsApi.GetDashboardListItems(TestAuth, 1234).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 299, httpresp.StatusCode)
}
