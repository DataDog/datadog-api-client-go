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
)

func TestGetGraphSnapshot(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	start := tests.ClockFromContext(ctx).Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"
	eventQuery := "successful builds"

	// Try to create a snapshot with a metric_query (and an optional event_query)
	snapshot, httpresp, err := Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx).MetricQuery(metricQuery).Start(start).End(end).EventQuery(eventQuery).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	assert.Equal(snapshot.GetGraphDef(), graphDef)
	assert.Equal(snapshot.GetMetricQuery(), metricQuery)
	assert.NotEmpty(snapshot.GetSnapshotUrl())

	// Try to create a snapshot with a graph_def
	snapshot, httpresp, err = Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx).GraphDef(graphDef).Start(start).End(end).Execute()
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	assert.Equal(snapshot.GetGraphDef(), graphDef)
	assert.NotEmpty(snapshot.GetSnapshotUrl())
}

func TestGetGraphSnapshotRequiredParams(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	var start int64 = 1
	var end int64 = 2
	metricQuery := "query"

	_, _, err := Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx).MetricQuery(metricQuery).End(end).Execute()
	assert.Contains(err.Error(), "start is required")
	_, _, err = Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx).MetricQuery(metricQuery).Start(start).Execute()
	assert.Contains(err.Error(), "end is required")
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
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SnapshotsApi.GetGraphSnapshot((ctx)).Start(345).End(123).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
