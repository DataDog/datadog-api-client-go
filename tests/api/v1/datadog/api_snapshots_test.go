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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	start := tests.ClockFromContext(ctx).Now().Unix()
	end := start + 24*60*60
	graphDef := `{"requests": [{"q": "system.load.1{*}"}]}`
	metricQuery := "system.load.1{*}"
	eventQuery := "successful builds"

	// Try to create a snapshot with a metric_query (and an optional event_query)
	snapshot, httpresp, err := Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx, start, end, *datadog.NewGetGraphSnapshotOptionalParameters().WithMetricQuery(metricQuery).WithEventQuery(eventQuery))
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	assert.Equal(snapshot.GetGraphDef(), graphDef)
	assert.Equal(snapshot.GetMetricQuery(), metricQuery)
	assert.NotEmpty(snapshot.GetSnapshotUrl())

	// Try to create a snapshot with a graph_def
	snapshot, httpresp, err = Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx, start, end, *datadog.NewGetGraphSnapshotOptionalParameters().WithGraphDef(graphDef))
	if err != nil {
		t.Fatalf("Error creating Snapshot: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)

	assert.Equal(snapshot.GetGraphDef(), graphDef)
	assert.NotEmpty(snapshot.GetSnapshotUrl())
}

func TestGraphGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SnapshotsApi.GetGraphSnapshot(ctx, 346, 123)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
