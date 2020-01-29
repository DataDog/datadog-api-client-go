package datadog_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestGetAllLogsIndexes(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	logIndexes, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetAllLogIndexes(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all log indexes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, 0 < len(logIndexes.GetIndexes()))
}

func TestGetLogsIndex(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	name := "main"
	logsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndex(TESTAUTH, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, name, logsIndex.GetName())
	filter := logsIndex.GetFilter()
	assert.Equal(t, "", filter.GetQuery())
	assert.Equal(t, 0, len(logsIndex.GetExclusionFilters()))
}

func TestLogsIndexOrder(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	name := "main"
	indexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndexOrder(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, 0 < len(indexOrder.GetIndexNames()))
	assert.Assert(t, is.Contains(indexOrder.GetIndexNames(), name))
}

func TestUpdateLogsIndex(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	name := "main"
	logsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndex(TESTAUTH, name).Execute()
	if err != nil {
		t.Fatalf("Error getting logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, name, logsIndex.GetName())

	updateLogsIndex := datadog.LogsIndex{
		Filter: logsIndex.GetFilter(),
		ExclusionFilters: []datadog.LogsExclusion{{
			Name:      fmt.Sprintf("datadog-api-client-go::%d", time.Now().UnixNano()/1000),
			IsEnabled: datadog.PtrBool(false),
			Filter: &datadog.LogsExclusionFilter{
				Query:      datadog.PtrString("hostname:datadog-api-client-go"),
				SampleRate: 1.0,
			},
		}},
	}

	updatedLogsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.UpdateLogsIndex(TESTAUTH, name).Body(updateLogsIndex).Execute()
	if err != nil {
		t.Fatalf("Error updating logs index '%s': Response %s: %v", name, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer resetExclusionFilters(name)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, name, updatedLogsIndex.GetName())
}

func TestUpdateLogsIndexOrder(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	indexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndexOrder(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting index order: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	newOrder := indexOrder.GetIndexNames()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	indexOrder.SetIndexNames(newOrder)

	newIndexOrder, httpresp, err := TESTAPICLIENT.LogsIndexesApi.UpdateLogsIndexOrder(TESTAUTH).Body(indexOrder).Execute()
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.DeepEqual(indexOrder.GetIndexNames(), newIndexOrder.GetIndexNames()))
}

func resetExclusionFilters(name string) {
	logsIndex, httpresp, err := TESTAPICLIENT.LogsIndexesApi.GetLogsIndex(TESTAUTH, name).Execute()
	if err != nil {
		log.Printf("Retrieving index: %s failed with %v: %v", name, httpresp.StatusCode, err)
	}

	updateLogsIndex := datadog.LogsIndex{
		Filter:           logsIndex.GetFilter(),
		ExclusionFilters: []datadog.LogsExclusion{},
	}
	_, httpresp, err = TESTAPICLIENT.LogsIndexesApi.UpdateLogsIndex(TESTAUTH, name).Body(updateLogsIndex).Execute()
	if err != nil {
		log.Printf("Reseting index: %s failed with %v: %v", name, httpresp.StatusCode, err)
	}
}
