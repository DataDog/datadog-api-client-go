package datadog_test

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestCreateGraphSnapshot(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testGraphSnapshot := datadog.GraphSnapshot{
		GraphDef:    datadog.PtrString("Graph definition"),
		MetricQuery: datadog.PtrString("Metrics query"),
		SnapshotUrl: datadog.PtrString("Snapshot url"),
	}

	snapshot, httpresp, err := TESTAPICLIENT.SnapshotsApi.CreateGraphSnapshot(TESTAUTH, testGraphSnapshot)
	if err != nil {
		t.Errorf("Error creating Snapshot %v: Status: %s: %v", testGraphSnapshot, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), testGraphSnapshot.GetGraphDef())
	assert.Equal(t, snapshot.GetMetricQuery(), testGraphSnapshot.GetMetricQuery())
	assert.Equal(t, snapshot.GetSnapshotUrl(), testGraphSnapshot.GetSnapshotUrl())
}
