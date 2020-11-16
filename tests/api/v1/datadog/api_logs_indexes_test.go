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

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
)

func enableLogsIndexesUnstableOperations(ctx context.Context) func() {
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetLogsIndex", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("ListLogIndexes", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetLogsIndexOrder", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("UpdateLogsIndexOrder", true)
	return func() { disableLogsIndexesUnstableOperations(ctx) }
}

func disableLogsIndexesUnstableOperations(ctx context.Context) {
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetLogsIndex", false)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("ListLogIndexes", false)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", false)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetLogsIndexOrder", false)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("UpdateLogsIndexOrder", false)
}

func TestGetAllLogsIndexes(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	defer enableLogsIndexesUnstableOperations(ctx)()
	assert := tests.Assert(ctx, t)

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

	logIndexes, httpresp, err := Client(ctx).LogsIndexesApi.ListLogIndexes(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all log indexes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(logIndexes.GetIndexes()) > 0)
}

func TestGetLogsIndex(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	assert := tests.Assert(ctx, t)
	defer enableLogsIndexesUnstableOperations(ctx)()

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

	logsIndex, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndex(ctx, name).Execute()
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
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	assert := tests.Assert(ctx, t)

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

	defer enableLogsIndexesUnstableOperations(ctx)()
	indexOrder, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndexOrder(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(indexOrder.GetIndexNames()) > 0)
	assert.Contains(indexOrder.GetIndexNames(), name)
}

func TestUpdateLogsIndex(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	assert := tests.Assert(ctx, t)
	defer enableLogsIndexesUnstableOperations(ctx)()

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

	logsIndex, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndex(ctx, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(name, logsIndex.GetName())

	updateLogsIndex := datadog.LogsIndexUpdateRequest{
		Filter: logsIndex.GetFilter(),
		ExclusionFilters: &[]datadog.LogsExclusion{{
			Name:      "datadog-api-client-go",
			IsEnabled: datadog.PtrBool(false),
			Filter: &datadog.LogsExclusionFilter{
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

	updatedLogsIndex, httpresp, err := Client(ctx).LogsIndexesApi.UpdateLogsIndex(ctx, name).Body(updateLogsIndex).Execute()
	if err != nil {
		t.Fatalf("Error updating logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(name, updatedLogsIndex.GetName())
}

func TestUpdateLogsIndexOrder(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	assert := tests.Assert(ctx, t)

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

	defer enableLogsIndexesUnstableOperations(ctx)()
	indexOrder, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndexOrder(ctx).Execute()
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

	newIndexOrder, httpresp, err := Client(ctx).LogsIndexesApi.UpdateLogsIndexOrder(ctx).Body(indexOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(indexOrder.GetIndexNames(), newIndexOrder.GetIndexNames())
}

func TestLogsIndexesListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			defer enableLogsIndexesUnstableOperations(ctx)()

			_, httpresp, err := Client(ctx).LogsIndexesApi.ListLogIndexes(ctx).Execute()
			fmt.Println(err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsIndexesGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			defer enableLogsIndexesUnstableOperations(ctx)()

			_, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndex(ctx, "shrugs").Execute()
			fmt.Println(err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsIndexUpdateRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsIndexUpdateRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsIndexUpdateRequest{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			defer enableLogsIndexesUnstableOperations(ctx)()

			_, httpresp, err := Client(ctx).LogsIndexesApi.UpdateLogsIndex(ctx, "shrugs").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdate429Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer enableLogsIndexesUnstableOperations(ctx)()

	data, err := tests.ReadFixture("fixtures/logs-indexes/error_429.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "LogsIndexesApiService.UpdateLogsIndex")
	assert.NoError(err)
	gock.New(URL).
		Put("/api/v1/logs/config/indexes/name").
		Reply(429).
		JSON(data)
	defer gock.Off()

	_, httpresp, err := Client(ctx).LogsIndexesApi.UpdateLogsIndex(ctx, "name").Body(datadog.LogsIndexUpdateRequest{}).Execute()
	assert.Equal(429, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetError())

}

func TestLogsIndexesOrderGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			defer enableLogsIndexesUnstableOperations(ctx)()

			_, httpresp, err := Client(ctx).LogsIndexesApi.GetLogsIndexOrder(ctx).Execute()
			fmt.Println(err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsIndexesOrderUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsIndexesOrder
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsIndexesOrder{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsIndexesOrder{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			defer enableLogsIndexesUnstableOperations(ctx)()

			_, httpresp, err := Client(ctx).LogsIndexesApi.UpdateLogsIndexOrder(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}
