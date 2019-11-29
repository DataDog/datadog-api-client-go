package datadog_test

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestCreateGraphSnapshot(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"

	snapshot, httpresp, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH, metricQuery, start, end, nil)
	if err != nil {
		t.Errorf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.Equal(t, snapshot.GetMetricQuery(), metricQuery)
	assert.Assert(t, snapshot.GetSnapshotUrl() != "")
}
