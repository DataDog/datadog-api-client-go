/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetGraphSnapshot(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	start := c.Clock.Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"
	eventQuery := "successful builds"

	// Try to create a snapshot with a metric_query (and an optional event_query)
	snapshot, httpresp, err := c.Client.SnapshotsApi.GetGraphSnapshot(c.Ctx).MetricQuery(metricQuery).Start(start).End(end).EventQuery(eventQuery).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.Equal(t, snapshot.GetMetricQuery(), metricQuery)
	assert.NotEmpty(t, snapshot.GetSnapshotUrl())

	// Try to create a snapshot with a graph_def
	snapshot, httpresp, err = c.Client.SnapshotsApi.GetGraphSnapshot(c.Ctx).GraphDef(graphDef).Start(start).End(end).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, snapshot.GetGraphDef(), graphDef)
	assert.NotEmpty(t, snapshot.GetSnapshotUrl())
}

func TestGetGraphSnapshotRequiredParams(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	var start int64 = 1
	var end int64 = 2
	metricQuery := "query"

	_, _, err := c.Client.SnapshotsApi.GetGraphSnapshot(c.Ctx).MetricQuery(metricQuery).End(end).Execute()
	assert.Contains(t, err.Error(), "start is required")
	_, _, err = c.Client.SnapshotsApi.GetGraphSnapshot(c.Ctx).MetricQuery(metricQuery).Start(start).Execute()
	assert.Contains(t, err.Error(), "end is required")
}

func TestGraphGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.SnapshotsApi.GetGraphSnapshot((c.Ctx)).Start(345).End(123).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
