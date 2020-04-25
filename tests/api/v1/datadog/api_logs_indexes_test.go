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
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func enableLogsIndexesUnstableOperations() {
	c.Client.GetConfig().SetUnstableOperationEnabled("GetLogsIndex", true)
	c.Client.GetConfig().SetUnstableOperationEnabled("ListLogIndexes", true)
	c.Client.GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", true)
}

func disableLogsIndexesUnstableOperations() {
	c.Client.GetConfig().SetUnstableOperationEnabled("GetLogsIndex", false)
	c.Client.GetConfig().SetUnstableOperationEnabled("ListLogIndexes", false)
	c.Client.GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", false)
}

func TestGetAllLogsIndexes(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-indexes.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	gock.New("https://api.datadoghq.com").
		Get("/api/v1/logs/config/indexes").
		Reply(200).
		JSON(data)

	logIndexes, httpresp, err := c.Client.LogsIndexesApi.ListLogIndexes(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all log indexes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, len(logIndexes.GetIndexes()) > 0)
}

func TestGetLogsIndex(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-index.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	gock.New("https://api.datadoghq.com").
		Get(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	logsIndex, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndex(c.Ctx, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, name, logsIndex.GetName())
	assert.Equal(t, int64(15), logsIndex.GetNumRetentionDays())
	assert.Equal(t, int64(200000000), logsIndex.GetDailyLimit())
	assert.Equal(t, false, logsIndex.GetIsRateLimited())
	filter := logsIndex.GetFilter()
	assert.Equal(t, "host:test.log.index", filter.GetQuery())
	assert.Equal(t, 0, len(logsIndex.GetExclusionFilters()))
}

func TestLogsIndexOrder(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer c.Close()

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	gock.New("https://api.datadoghq.com").
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndexOrder(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, len(indexOrder.GetIndexNames()) > 0)
	assert.Contains(t, indexOrder.GetIndexNames(), name)
}

func TestUpdateLogsIndex(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	data, err := tests.ReadFixture("fixtures/logs-indexes/log-index.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	gock.New("https://api.datadoghq.com").
		Get(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	logsIndex, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndex(c.Ctx, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, name, logsIndex.GetName())

	updateLogsIndex := datadog.LogsIndex{
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

	gock.New("https://api.datadoghq.com").
		Put(fmt.Sprintf("/api/v1/logs/config/indexes/%s", name)).
		Reply(200).
		JSON(data)

	updatedLogsIndex, httpresp, err := c.Client.LogsIndexesApi.UpdateLogsIndex(c.Ctx, name).Body(updateLogsIndex).Execute()
	if err != nil {
		t.Fatalf("Error updating logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, name, updatedLogsIndex.GetName())
}

func TestUpdateLogsIndexOrder(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer c.Close()

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	gock.New("https://api.datadoghq.com").
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndexOrder(c.Ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	newOrder := indexOrder.GetIndexNames()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	indexOrder.SetIndexNames(newOrder)

	data, err = tests.ReadFixture("fixtures/logs-indexes/update-logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	gock.New("https://api.datadoghq.com").
		Put("/api/v1/logs/config/index-orde").
		Reply(200).
		JSON(data)

	newIndexOrder, httpresp, err := c.Client.LogsIndexesApi.UpdateLogsIndexOrder(c.Ctx).Body(indexOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, indexOrder.GetIndexNames(), newIndexOrder.GetIndexNames())
}

func TestLogsIndexesListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.LogsIndexesApi.ListLogIndexes(tc.Ctx).Execute()
			fmt.Println(err)
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestLogsIndexesGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", c.Ctx, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndex(tc.Ctx, "shrugs").Execute()
			fmt.Println(err)
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.LogsIndex
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.LogsIndex{}, 400},
		{"403 Forbidden", fake_auth, datadog.LogsIndex{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.LogsIndexesApi.UpdateLogsIndex(tc.Ctx, "shrugs").Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetError())
			}
		})
	}
}

func TestLogsIndexesUpdate429Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	data, err := tests.ReadFixture("fixtures/logs-indexes/error_429.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	gock.New("https://api.datadoghq.com").
		Put("/api/v1/logs/config/indexes/name").
		Reply(429).
		JSON(data)
	defer gock.Off()

	_, httpresp, err := c.Client.LogsIndexesApi.UpdateLogsIndex(c.Ctx, "name").Body(datadog.LogsIndex{}).Execute()
	assert.Equal(t, 429, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetError())

}

func TestLogsIndexesOrderGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.LogsIndexesApi.GetLogsIndexOrder(tc.Ctx).Execute()
			fmt.Println(err)
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestLogsIndexesOrderUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()
	enableLogsIndexesUnstableOperations()
	defer disableLogsIndexesUnstableOperations()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.LogsIndexesOrder
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, datadog.LogsIndexesOrder{}, 400},
		{"403 Forbidden", fake_auth, datadog.LogsIndexesOrder{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.LogsIndexesApi.UpdateLogsIndexOrder(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetError())
			}
		})
	}
}
