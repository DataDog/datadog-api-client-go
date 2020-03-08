/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestGetGraphSnapshot(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := TESTCLOCK.Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"
	eventQuery := "successful builds"

	// Try to create a snapshot with a metric_query (and an optional event_query)
	snapshot, httpresp, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).Start(start).End(end).EventQuery(eventQuery).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.Equal(t, snapshot.GetMetricQuery(), metricQuery)
	assert.Assert(t, snapshot.GetSnapshotUrl() != "")

	// Try to create a snapshot with a graph_def
	snapshot, httpresp, err = TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).GraphDef(graphDef).Start(start).End(end).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.Assert(t, snapshot.GetSnapshotUrl() != "")
}

func TestGetGraphSnapshotRequiredParams(t *testing.T) {
	var start int64 = 1
	var end int64 = 2
	metricQuery := "query"

	_, _, err := TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).End(end).Execute()
	assert.ErrorContains(t, err, "start is required")
	_, _, err = TESTAPICLIENT.SnapshotsApi.GetGraphSnapshot(TESTAUTH).MetricQuery(metricQuery).Start(start).Execute()
	assert.ErrorContains(t, err, "end is required")
}
