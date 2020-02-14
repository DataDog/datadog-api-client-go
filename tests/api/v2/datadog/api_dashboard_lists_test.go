package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/stretchr/testify/assert"
)

const (
	integrationTimeboardId string = "1"
	customTimeboardId      string = "q5j-nti-fv6"
	customScreenboardId    string = "4n7-s4g-dqv"
)

var (
	dashboardListID int64
	testAuthV1      context.Context
	testAPIClientV1 *datadogV1.APIClient
)

func initializeClientV1() {
	testAuthV1 = context.WithValue(
		context.Background(),
		datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			"apiKeyAuth": datadogV1.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": datadogV1.APIKey{
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadogV1.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	testAPIClientV1 = datadogV1.NewAPIClient(config)
}

func createDashboardList() error {
	initializeClientV1()
	res, httpresp, err := testAPIClientV1.DashboardListsApi.CreateDashboardList(testAuthV1).
		Body(datadogV1.DashboardList{Name: fmt.Sprintf("go-client-test-v2-%d", time.Now().Unix())}).
		Execute()
	if err != nil || httpresp.StatusCode != 200 {
		return fmt.Errorf("error creating dashboard list: %v", err.(datadogV1.GenericOpenAPIError).Body())
	}
	dashboardListID = res.GetId()
	return nil
}

func deleteDashboardList() {
	testAPIClientV1.DashboardListsApi.DeleteDashboardList(testAuthV1, dashboardListID).Execute()
}

func TestDashboardListItemCRUD(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	err := createDashboardList()
	defer deleteDashboardList()
	if err != nil {
		t.Fatal(err)
	}

	integrationTimeboard := datadog.NewDashboardListItem(integrationTimeboardId, datadog.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboard := datadog.NewDashboardListItem(customTimeboardId, datadog.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboard := datadog.NewDashboardListItem(customScreenboardId, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	dashboards := []datadog.DashboardListItem{
		*integrationTimeboard,
		*customTimeboard,
		*customScreenboard,
	}
	body := datadog.NewDashboardListItems(dashboards)

	addResponse, httpresp, err := TestApiClient.DashboardListsApi.AddDashboardListItems(TestAuth, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error adding items to dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 3, len(addResponse.GetAddedDashboardsToList()))
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *integrationTimeboard)
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *customTimeboard)
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *customScreenboard)

	deleteResponse, httpresp, err := TestApiClient.DashboardListsApi.DeleteDashboardListItems(TestAuth, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 3, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *integrationTimeboard)
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *customTimeboard)
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *customScreenboard)

	getResponse, httpresp, err := TestApiClient.DashboardListsApi.GetDashboardListItems(TestAuth, dashboardListID).Execute()
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, int32(0), getResponse.GetTotal())
	assert.Equal(t, 0, len(getResponse.GetDashboards()))

	updateResponse, httpresp, err := TestApiClient.DashboardListsApi.UpdateDashboardListItems(TestAuth, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error updating items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 3, len(updateResponse.GetDashboards()))
	assert.Contains(t, updateResponse.GetDashboards(), *integrationTimeboard)
	assert.Contains(t, updateResponse.GetDashboards(), *customTimeboard)
	assert.Contains(t, updateResponse.GetDashboards(), *customScreenboard)

	// Leave only one dash in the list for easier assertion
	dashboards = []datadog.DashboardListItem{
		*integrationTimeboard,
		*customTimeboard,
	}
	body.SetDashboards(dashboards)
	deleteResponse, httpresp, err = TestApiClient.DashboardListsApi.DeleteDashboardListItems(TestAuth, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 2, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Equal(t, 200, httpresp.StatusCode)
	getResponse, httpresp, err = TestApiClient.DashboardListsApi.GetDashboardListItems(TestAuth, dashboardListID).Execute()
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 1, len(getResponse.GetDashboards()))
	assert.Equal(t, int32(1), getResponse.GetTotal())
	assert.True(t, getResponse.GetDashboards()[0].GetIsReadOnly())
	assert.True(t, getResponse.GetDashboards()[0].GetIsShared())
	assert.Equal(t, customScreenboardId, getResponse.GetDashboards()[0].GetId())
	assert.Equal(t, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD, getResponse.GetDashboards()[0].GetType())
	assert.Equal(t, "For dashboard list tests - DO NOT DELETE", getResponse.GetDashboards()[0].GetTitle())
	assert.Equal(t, "/dashboard/4n7-s4g-dqv/for-dashboard-list-tests---do-not-delete", getResponse.GetDashboards()[0].GetUrl())
	assert.True(t, getResponse.GetDashboards()[0].GetPopularity() >= 0)
	assert.NotNil(t, getResponse.GetDashboards()[0].Author)
	assert.NotNil(t, getResponse.GetDashboards()[0].Modified)
	assert.NotNil(t, getResponse.GetDashboards()[0].Created)
	assert.Nil(t, getResponse.GetDashboards()[0].Icon)
}
