package datadog_test

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestGetGraphSnapshot(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"

	snapshot, httpresp, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).Start(start).End(end).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.Equal(t, snapshot.GetMetricQuery(), metricQuery)
	assert.Assert(t, snapshot.GetSnapshotUrl() != "")
}

func TestGetGraphSnapshotRequiredParams(t *testing.T) {
	var start int64 = 1
	var end int64 = 2
	metricQuery := "query"

	_, _, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).Start(start).End(end).Execute()
	assert.ErrorContains(t, err, "metricQuery is required")
	_, _, err = TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).End(end).Execute()
	assert.ErrorContains(t, err, "start is required")
	_, _, err = TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).Start(start).Execute()
	assert.ErrorContains(t, err, "end is required")
}
