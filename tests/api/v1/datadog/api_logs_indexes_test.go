/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func enableLogsIndexesUnstableOperations() {
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("GetLogsIndex", true)
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("GetAllLogIndexes", true)
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", true)
}

func disableLogsIndexesUnstableOperations() {
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("GetLogsIndex", false)
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("GetAllLogIndexes", false)
    TESTAPICLIENT.GetConfig().SetUnstableOperationEnabled("UpdateLogsIndex", false)
}

func TestGetAllLogsIndexes(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownTest(t)
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

	logIndexes, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetAllLogIndexes(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all log indexes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, 0 < len(logIndexes.GetIndexes()))
}

func TestGetLogsIndex(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownTest(t)
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

	logsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndex(TESTAUTH, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
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
	defer teardownTest(t)

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	name := "main"

	gock.New("https://api.datadoghq.com").
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndexOrder(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, 0 < len(indexOrder.GetIndexNames()))
	assert.Assert(t, is.Contains(indexOrder.GetIndexNames(), name))
}

func TestUpdateLogsIndex(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownTest(t)
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

	logsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndex(TESTAUTH, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
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

	updatedLogsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.UpdateLogsIndex(TESTAUTH, name).Body(updateLogsIndex).Execute()
	if err != nil {
		t.Fatalf("Error updating logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, name, updatedLogsIndex.GetName())
}

func TestUpdateLogsIndexOrder(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownTest(t)

	data, err := tests.ReadFixture("fixtures/logs-indexes/logs-index-order.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	gock.New("https://api.datadoghq.com").
		Get("/api/v1/logs/config/index-order").
		Reply(200).
		JSON(data)

	indexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndexOrder(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

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

	newIndexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.UpdateLogsIndexOrder(TESTAUTH).Body(indexOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.DeepEqual(indexOrder.GetIndexNames(), newIndexOrder.GetIndexNames()))
}
