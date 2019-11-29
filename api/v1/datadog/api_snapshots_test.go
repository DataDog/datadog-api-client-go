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
	testGraphSnapshot := datadog.GraphSnapshot{
		GraphDef:    datadog.PtrString(`{"requests": [{"q": "system.load.1{*}"}]}`),
		MetricQuery: datadog.PtrString("system.load.1{*}"),
	}

	snapshot, httpresp, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH, *testGraphSnapshot.MetricQuery, start, end, nil)
	if err != nil {
		t.Errorf("Error creating Snapshot %v: Status: %s: %v", testGraphSnapshot, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), testGraphSnapshot.GetGraphDef())
	assert.Equal(t, snapshot.GetMetricQuery(), testGraphSnapshot.GetMetricQuery())
	assert.Assert(t, snapshot.GetSnapshotUrl() != "")
}
