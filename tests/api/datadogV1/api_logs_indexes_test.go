/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"

	"gopkg.in/h2non/gock.v1"
)

func TestGetAllLogsIndexes(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-indexes.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.GetLogsIndexOrder")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/logs/config/indexes").
		Reply(200).
		JSON(data)

	logIndexes, httpresp, err := api.ListLogIndexes(ctx)
	if err != nil {
		t.Fatalf("Error getting all log indexes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(logIndexes.GetIndexes()) > 0)
}

func TestGetLogsIndex(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-index.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	logsIndex, httpresp, err := api.GetLogsIndex(ctx, name)
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(name, logsIndex.GetName())
	assert.Equal(int64(15), logsIndex.GetNumRetentionDays())
	assert.Equal(int64(200000000), logsIndex.GetDailyLimit())
	assert.Equal(false, logsIndex.GetIsRateLimited())
	filter := logsIndex.GetFilter()
	assert.Equal("host:test.log.index", filter.GetQuery())
	assert.Equal(0, len(logsIndex.GetExclusionFilters()))
}

func TestLogsIndexOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.GetLogsIndexOrder")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := api.GetLogsIndexOrder(ctx)
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(indexOrder.GetIndexNames()) > 0)
	assert.Contains(indexOrder.GetIndexNames(), name)
}

func TestUpdateLogsIndex(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-index.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.GetLogsIndex")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	logsIndex, httpresp, err := api.GetLogsIndex(ctx, name)
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(name, logsIndex.GetName())

	updateLogsIndex := datadogV1.LogsIndexUpdateRequest{
		Filter: logsIndex.GetFilter(),
		ExclusionFilters: []datadogV1.LogsExclusion{{
			Name:      "datadog-api-client-go",
			IsEnabled: datadog.PtrBool(false),
			Filter: &datadogV1.LogsExclusionFilter{
				Query:      datadog.PtrString("hostname:datadog-api-client-go"),
				SampleRate: 1.0,
			},
		}},
	}

	data, err = tests.ReadFixture("fixtures/logs-indexes/update-logs-index.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err = Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.UpdateLogsIndex")
	assert.NoError(err)
	gock.New(URL).
		Put(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	updatedLogsIndex, httpresp, err := api.UpdateLogsIndex(ctx, name, updateLogsIndex)
	if err != nil {
		t.Fatalf("Error updating logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(name, updatedLogsIndex.GetName())
}

func TestUpdateLogsIndexOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	defer gock.Off()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.GetLogsIndexOrder")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := api.GetLogsIndexOrder(ctx)
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	newOrder := indexOrder.GetIndexNames()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	indexOrder.SetIndexNames(newOrder)

	data, err = tests.ReadFixture("fixtures/logs-indexes/update-logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err = Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.UpdateLogsIndexOrder")
	assert.NoError(err)
	gock.New(URL).
		Put("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	newIndexOrder, httpresp, err := api.UpdateLogsIndexOrder(ctx, indexOrder)
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(indexOrder.GetIndexNames(), newIndexOrder.GetIndexNames())
}

func TestLogsIndexesListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsIndexesApi(Client(ctx))

			_, httpresp, err := api.ListLogIndexes(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsIndexesGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsIndexesApi(Client(ctx))

			_, httpresp, err := api.GetLogsIndex(ctx, "shrugs")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsIndexUpdateRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.LogsIndexUpdateRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.LogsIndexUpdateRequest{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsIndexesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsIndex(ctx, "shrugs", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdate429Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsIndexesApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/logs-indexes/error_429.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.UpdateLogsIndex")
	assert.NoError(err)
	gock.New(URL).
		Put("/api/v1/logs/config/indexes/name").
		Times(4). //default maxRetries=3, total 4 retries by default
		Reply(429).
		JSON(data)
	defer gock.Off()

	_, httpresp, err := api.UpdateLogsIndex(ctx, "name", datadogV1.LogsIndexUpdateRequest{})
	assert.Equal(429, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetError())

}

func TestLogsIndexesOrderGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsIndexesApi(Client(ctx))

			_, httpresp, err := api.GetLogsIndexOrder(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsIndexesOrderUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsIndexesOrder
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.LogsIndexesOrder{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.LogsIndexesOrder{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsIndexesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsIndexOrder(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}
