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

func TestDashboardListLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now()
	testDashboardList := datadog.DashboardList{
		Name: fmt.Sprintf("%s", start),
	}

	// Create downtime
	dashboardList, httpresp, err := TESTAPICLIENT.DashboardListsApi.CreateDashboardList(TESTAUTH, testDashboardList)
	if err != nil {
		t.Fatalf("Error creating dashboard list %v: Response %s: %v", testDashboardList, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboardList(dashboardList.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, dashboardList.GetName(), testDashboardList.GetName())

	// Edit a downtime
	editedDashboardList := datadog.DashboardList{Name: fmt.Sprintf("updated%s", start)}
	updatedDashboardList, httpresp, err := TESTAPICLIENT.DashboardListsApi.UpdateDashboardList(TESTAUTH, dashboardList.GetId(), editedDashboardList)
	if err != nil {
		t.Errorf("Error updating dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, editedDashboardList.GetName(), updatedDashboardList.GetName())

	// Check downtime existence
	fetchedDashboardList, httpresp, err := TESTAPICLIENT.DashboardListsApi.GetDashboardList(TESTAUTH, dashboardList.GetId())
	if err != nil {
		t.Errorf("Error fetching dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, editedDashboardList.GetName(), fetchedDashboardList.GetName())

	// Find our downtime in the full list
	dashboardLists, httpresp, err := TESTAPICLIENT.DashboardListsApi.GetAllDashboardLists(TESTAUTH)
	if err != nil {
		t.Errorf("Error fetching dashboard lists: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.Contains(dashboardLists, fetchedDashboardList))

	// Cancel downtime
	deletedDashboardListResponse, httpresp, err := TESTAPICLIENT.DashboardListsApi.DeleteDashboardList(TESTAUTH, dashboardList.GetId())
	if err != nil {
		t.Errorf("Error deleting dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)
	assert.Equal(t, deletedDashboardListResponse.GetDeletedDashboardListId(), dashboardList.GetId())
}

func deleteDashboardList(dashboardListID int64) {
	httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, dashboardListID)
	if err != nil {
		log.Printf("Deleting dashboard list: %v failed with %v, Another test may have already deleted this dashboard list: %v", dashboardListID, httpresp.StatusCode, err)
	}
}
