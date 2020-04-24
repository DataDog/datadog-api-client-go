/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/stretchr/testify/assert"
)

const (
	integrationTimeboardID string = "1"
	customTimeboardID      string = "q5j-nti-fv6"
	customScreenboardID    string = "4n7-s4g-dqv"
)

var (
	dashboardListID int64
	testAuthV1      context.Context
	testAPIClientV1 *datadogV1.APIClient
)

func initializeClientV1(c *Client) {
	testAuthV1 = context.WithValue(
		context.Background(),
		datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadogV1.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	config.HTTPClient = c.Client.GetConfig().HTTPClient
	testAPIClientV1 = datadogV1.NewAPIClient(config)
}

func createDashboardList(c *Client) error {
	initializeClientV1(c)
	res, httpresp, err := testAPIClientV1.DashboardListsApi.CreateDashboardList(testAuthV1).
		Body(datadogV1.DashboardList{Name: fmt.Sprintf("go-client-test-v2-%d", c.Clock.Now().Unix())}).
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
	c := NewClientWithRecording(NewClientAuthContext(), t)
	defer c.Close()
	err := createDashboardList(c)
	defer deleteDashboardList()
	if err != nil {
		t.Fatal(err)
	}

	integrationTimeboard := datadog.NewDashboardListItem(integrationTimeboardID, datadog.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboard := datadog.NewDashboardListItem(customTimeboardID, datadog.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboard := datadog.NewDashboardListItem(customScreenboardID, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	dashboards := []datadog.DashboardListItem{
		*integrationTimeboard,
		*customTimeboard,
		*customScreenboard,
	}
	body := datadog.NewDashboardListItems(dashboards)

	addResponse, httpresp, err := c.Client.DashboardListsApi.CreateDashboardListItems(c.Ctx, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error adding items to dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 3, len(addResponse.GetAddedDashboardsToList()))
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *integrationTimeboard)
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *customTimeboard)
	assert.Contains(t, addResponse.GetAddedDashboardsToList(), *customScreenboard)

	deleteResponse, httpresp, err := c.Client.DashboardListsApi.DeleteDashboardListItems(c.Ctx, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 3, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *integrationTimeboard)
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *customTimeboard)
	assert.Contains(t, deleteResponse.GetDeletedDashboardsFromList(), *customScreenboard)

	getResponse, httpresp, err := c.Client.DashboardListsApi.GetDashboardListItems(c.Ctx, dashboardListID).Execute()
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, int64(0), getResponse.GetTotal())
	assert.Equal(t, 0, len(getResponse.GetDashboards()))

	updateResponse, httpresp, err := c.Client.DashboardListsApi.UpdateDashboardListItems(c.Ctx, dashboardListID).Body(*body).Execute()
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
	deleteResponse, httpresp, err = c.Client.DashboardListsApi.DeleteDashboardListItems(c.Ctx, dashboardListID).Body(*body).Execute()
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 2, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Equal(t, 200, httpresp.StatusCode)
	getResponse, httpresp, err = c.Client.DashboardListsApi.GetDashboardListItems(c.Ctx, dashboardListID).Execute()
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 1, len(getResponse.GetDashboards()))
	assert.Equal(t, int64(1), getResponse.GetTotal())
	assert.True(t, getResponse.GetDashboards()[0].GetIsReadOnly())
	assert.True(t, getResponse.GetDashboards()[0].GetIsShared())
	assert.Equal(t, customScreenboardID, getResponse.GetDashboards()[0].GetId())
	assert.Equal(t, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD, getResponse.GetDashboards()[0].GetType())
	assert.Equal(t, "For dashboard list tests - DO NOT DELETE", getResponse.GetDashboards()[0].GetTitle())
	assert.Equal(t, "/dashboard/4n7-s4g-dqv/for-dashboard-list-tests---do-not-delete", getResponse.GetDashboards()[0].GetUrl())
	assert.True(t, getResponse.GetDashboards()[0].GetPopularity() >= 0)
	assert.NotNil(t, getResponse.GetDashboards()[0].Author)
	assert.NotNil(t, getResponse.GetDashboards()[0].Modified)
	assert.NotNil(t, getResponse.GetDashboards()[0].Created)
	assert.Nil(t, getResponse.GetDashboards()[0].Icon)
}

func TestDashboardListGetItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(NewClientAuthContext(), t)
	defer c.Close()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", context.Background(), 403},
		{"404 Not Found", c.Ctx, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.DashboardListsApi.GetDashboardListItems(tc.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListCreateItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(NewClientAuthContext(), t)
	defer c.Close()

	err := createDashboardList(c)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, dashboardListID, 400},
		{"403 Forbidden", context.Background(), 0, 403},
		{"404 Not Found", c.Ctx, 0, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.DashboardListsApi.CreateDashboardListItems(tc.Ctx, tc.ID).Body(datadog.DashboardListItems{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListUpdateItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(NewClientAuthContext(), t)
	defer c.Close()

	err := createDashboardList(c)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, dashboardListID, 400},
		{"403 Forbidden", context.Background(), 0, 403},
		{"404 Not Found", c.Ctx, 0, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.DashboardListsApi.UpdateDashboardListItems(tc.Ctx, tc.ID).Body(datadog.DashboardListItems{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListDeleteItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(NewClientAuthContext(), t)
	defer c.Close()

	err := createDashboardList(c)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		{"400 Bad Request", c.Ctx, dashboardListID, 400},
		{"403 Forbidden", context.Background(), 0, 403},
		{"404 Not Found", c.Ctx, 0, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := c.Client.DashboardListsApi.DeleteDashboardListItems(tc.Ctx, tc.ID).Body(datadog.DashboardListItems{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
