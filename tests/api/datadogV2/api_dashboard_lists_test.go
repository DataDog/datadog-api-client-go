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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

const (
	integrationTimeboardID string = "1"
	customTimeboardID      string = "q5j-nti-fv6"
	customScreenboardID    string = "4n7-s4g-dqv"
)

var (
	dashboardListID int64
	testAuthV1      context.Context
	testAPIClientV1 *datadog.APIClient
)

func initializeClientV1(ctx context.Context) {
	testAuthV1 = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	config.HTTPClient = Client(ctx).GetConfig().HTTPClient
	testAPIClientV1 = datadog.NewAPIClient(config)
}

func createDashboardList(ctx context.Context, t *testing.T) error {
	initializeClientV1(ctx)
	api := datadogV1.NewDashboardListsApi(testAPIClientV1)
	res, httpresp, err := api.CreateDashboardList(testAuthV1, datadogV1.DashboardList{Name: *tests.UniqueEntityName(ctx, t)})
	if err != nil || httpresp.StatusCode != 200 {
		return fmt.Errorf("error creating dashboard list: %v", err.(datadog.GenericOpenAPIError).Body())
	}
	dashboardListID = res.GetId()
	return nil
}

func deleteDashboardList() {
	api := datadogV1.NewDashboardListsApi(testAPIClientV1)
	api.DeleteDashboardList(testAuthV1, dashboardListID)
}

func TestDashboardListItemCRUD(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV2.NewDashboardListsApi(Client(ctx))

	err := createDashboardList(ctx, t)
	defer deleteDashboardList()
	if err != nil {
		t.Fatal(err)
	}

	integrationTimeboard := datadogV2.NewDashboardListItemRequest(integrationTimeboardID, datadogV2.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboard := datadogV2.NewDashboardListItemRequest(customTimeboardID, datadogV2.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboard := datadogV2.NewDashboardListItemRequest(customScreenboardID, datadogV2.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	dashboards := []datadogV2.DashboardListItemRequest{
		*integrationTimeboard,
		*customTimeboard,
		*customScreenboard,
	}

	integrationTimeboardResponse := datadogV2.NewDashboardListItemResponse(integrationTimeboardID, datadogV2.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboardResponse := datadogV2.NewDashboardListItemResponse(customTimeboardID, datadogV2.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboardResponse := datadogV2.NewDashboardListItemResponse(customScreenboardID, datadogV2.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	addRequest := datadogV2.NewDashboardListAddItemsRequest()
	addRequest.SetDashboards(dashboards)

	addResponse, httpresp, err := api.CreateDashboardListItems(ctx, dashboardListID, *addRequest)
	if err != nil {
		t.Fatalf("error adding items to dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(addResponse.GetAddedDashboardsToList()))
	assert.Contains(addResponse.GetAddedDashboardsToList(), *integrationTimeboardResponse)
	assert.Contains(addResponse.GetAddedDashboardsToList(), *customTimeboardResponse)
	assert.Contains(addResponse.GetAddedDashboardsToList(), *customScreenboardResponse)

	deleteRequest := datadogV2.NewDashboardListDeleteItemsRequest()
	deleteRequest.SetDashboards(dashboards)
	deleteResponse, httpresp, err := api.DeleteDashboardListItems(ctx, dashboardListID, *deleteRequest)
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *integrationTimeboardResponse)
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *customTimeboardResponse)
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *customScreenboardResponse)

	getResponse, httpresp, err := api.GetDashboardListItems(ctx, dashboardListID)
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(0), getResponse.GetTotal())
	assert.Equal(0, len(getResponse.GetDashboards()))

	updateRequest := datadogV2.NewDashboardListUpdateItemsRequest()
	updateRequest.SetDashboards(dashboards)
	updateResponse, httpresp, err := api.UpdateDashboardListItems(ctx, dashboardListID, *updateRequest)
	if err != nil {
		t.Fatalf("error updating items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(updateResponse.GetDashboards()))
	assert.Contains(updateResponse.GetDashboards(), *integrationTimeboardResponse)
	assert.Contains(updateResponse.GetDashboards(), *customTimeboardResponse)
	assert.Contains(updateResponse.GetDashboards(), *customScreenboardResponse)

	// Leave only one dash in the list for easier assertion
	dashboards = []datadogV2.DashboardListItemRequest{
		*integrationTimeboard,
		*customTimeboard,
	}
	deleteRequest.SetDashboards(dashboards)
	deleteResponse, httpresp, err = api.DeleteDashboardListItems(ctx, dashboardListID, *deleteRequest)
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(2, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Equal(200, httpresp.StatusCode)
	getResponse, _, err = api.GetDashboardListItems(ctx, dashboardListID)
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(1, len(getResponse.GetDashboards()))
	assert.Equal(int64(1), getResponse.GetTotal())
	assert.True(getResponse.GetDashboards()[0].GetIsReadOnly())
	assert.True(getResponse.GetDashboards()[0].GetIsShared())
	assert.Equal(customScreenboardID, getResponse.GetDashboards()[0].GetId())
	assert.Equal(datadogV2.DASHBOARDTYPE_CUSTOM_SCREENBOARD, getResponse.GetDashboards()[0].GetType())
	assert.Equal("For dashboard list tests - DO NOT DELETE", getResponse.GetDashboards()[0].GetTitle())
	assert.Equal("/dashboard/4n7-s4g-dqv/for-dashboard-list-tests---do-not-delete", getResponse.GetDashboards()[0].GetUrl())
	assert.True(getResponse.GetDashboards()[0].GetPopularity() >= 0)
	assert.NotNil(getResponse.GetDashboards()[0].Author)
	assert.NotNil(getResponse.GetDashboards()[0].Modified)
	assert.NotNil(getResponse.GetDashboards()[0].Created)
	assert.Nil(getResponse.GetDashboards()[0].Icon)
}

func TestDashboardListGetItemsErrors(t *testing.T) {
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
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()
			assert := tests.Assert(ctx, t)
			api := datadogV2.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.GetDashboardListItems(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV2.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListCreateItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	err := createDashboardList(ctx, t)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, dashboardListID, 400},
		"403 Forbidden":   {WithFakeAuth, 0, 403},
		"404 Not Found":   {WithTestAuth, 0, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()
			assert := tests.Assert(ctx, t)
			api := datadogV2.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.CreateDashboardListItems(ctx, tc.ID, *datadogV2.NewDashboardListAddItemsRequest())
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV2.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListUpdateItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	err := createDashboardList(ctx, t)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, dashboardListID, 400},
		"403 Forbidden":   {WithFakeAuth, 0, 403},
		"404 Not Found":   {WithTestAuth, 0, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()
			assert := tests.Assert(ctx, t)
			api := datadogV2.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.UpdateDashboardListItems(ctx, tc.ID, *datadogV2.NewDashboardListUpdateItemsRequest())
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV2.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListDeleteItemsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	err := createDashboardList(ctx, t)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDashboardList()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, dashboardListID, 400},
		"403 Forbidden":   {WithFakeAuth, 0, 403},
		"404 Not Found":   {WithTestAuth, 0, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()
			assert := tests.Assert(ctx, t)
			api := datadogV2.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.DeleteDashboardListItems(ctx, tc.ID, *datadogV2.NewDashboardListDeleteItemsRequest())
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV2.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
