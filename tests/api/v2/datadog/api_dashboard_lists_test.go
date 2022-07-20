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

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

const (
	integrationTimeboardID string = "1"
	customTimeboardID      string = "q5j-nti-fv6"
	customScreenboardID    string = "4n7-s4g-dqv"
)

var (
	dashboardListID int64
	testAuthV1      context.Context
	testAPIClientV1 *common.APIClient
)

func initializeClientV1(ctx context.Context) {
	testAuthV1 = context.WithValue(
		context.Background(),
		common.ContextAPIKeys,
		map[string]common.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := common.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	config.HTTPClient = Client(ctx).GetConfig().HTTPClient
	testAPIClientV1 = common.NewAPIClient(config)
}

func createDashboardList(ctx context.Context, t *testing.T) error {
	initializeClientV1(ctx)
	api := datadogV1.NewDashboardListsApi(testAPIClientV1)
	res, httpresp, err := api.CreateDashboardList(testAuthV1, datadogV1.DashboardList{Name: *tests.UniqueEntityName(ctx, t)})
	if err != nil || httpresp.StatusCode != 200 {
		return fmt.Errorf("error creating dashboard list: %v", err.(common.GenericOpenAPIError).Body())
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
	api := datadog.NewDashboardListsApi(Client(ctx))

	err := createDashboardList(ctx, t)
	defer deleteDashboardList()
	if err != nil {
		t.Fatal(err)
	}

	integrationTimeboard := datadog.NewDashboardListItemRequest(integrationTimeboardID, datadog.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboard := datadog.NewDashboardListItemRequest(customTimeboardID, datadog.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboard := datadog.NewDashboardListItemRequest(customScreenboardID, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	dashboards := []datadog.DashboardListItemRequest{
		*integrationTimeboard,
		*customTimeboard,
		*customScreenboard,
	}

	integrationTimeboardResponse := datadog.NewDashboardListItemResponse(integrationTimeboardID, datadog.DASHBOARDTYPE_INTEGRATION_TIMEBOARD)
	customTimeboardResponse := datadog.NewDashboardListItemResponse(customTimeboardID, datadog.DASHBOARDTYPE_CUSTOM_TIMEBOARD)
	customScreenboardResponse := datadog.NewDashboardListItemResponse(customScreenboardID, datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD)

	addRequest := datadog.NewDashboardListAddItemsRequest()
	addRequest.SetDashboards(dashboards)

	addResponse, httpresp, err := api.CreateDashboardListItems(ctx, dashboardListID, *addRequest)
	if err != nil {
		t.Fatalf("error adding items to dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(addResponse.GetAddedDashboardsToList()))
	assert.Contains(addResponse.GetAddedDashboardsToList(), *integrationTimeboardResponse)
	assert.Contains(addResponse.GetAddedDashboardsToList(), *customTimeboardResponse)
	assert.Contains(addResponse.GetAddedDashboardsToList(), *customScreenboardResponse)

	deleteRequest := datadog.NewDashboardListDeleteItemsRequest()
	deleteRequest.SetDashboards(dashboards)
	deleteResponse, httpresp, err := api.DeleteDashboardListItems(ctx, dashboardListID, *deleteRequest)
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *integrationTimeboardResponse)
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *customTimeboardResponse)
	assert.Contains(deleteResponse.GetDeletedDashboardsFromList(), *customScreenboardResponse)

	getResponse, httpresp, err := api.GetDashboardListItems(ctx, dashboardListID)
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(0), getResponse.GetTotal())
	assert.Equal(0, len(getResponse.GetDashboards()))

	updateRequest := datadog.NewDashboardListUpdateItemsRequest()
	updateRequest.SetDashboards(dashboards)
	updateResponse, httpresp, err := api.UpdateDashboardListItems(ctx, dashboardListID, *updateRequest)
	if err != nil {
		t.Fatalf("error updating items from dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(3, len(updateResponse.GetDashboards()))
	assert.Contains(updateResponse.GetDashboards(), *integrationTimeboardResponse)
	assert.Contains(updateResponse.GetDashboards(), *customTimeboardResponse)
	assert.Contains(updateResponse.GetDashboards(), *customScreenboardResponse)

	// Leave only one dash in the list for easier assertion
	dashboards = []datadog.DashboardListItemRequest{
		*integrationTimeboard,
		*customTimeboard,
	}
	deleteRequest.SetDashboards(dashboards)
	deleteResponse, httpresp, err = api.DeleteDashboardListItems(ctx, dashboardListID, *deleteRequest)
	if err != nil {
		t.Fatalf("error deleting items from dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(2, len(deleteResponse.GetDeletedDashboardsFromList()))
	assert.Equal(200, httpresp.StatusCode)
	getResponse, httpresp, err = api.GetDashboardListItems(ctx, dashboardListID)
	if err != nil {
		t.Fatalf("error getting items from dashboard list %d: Response %s: %v", dashboardListID, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(1, len(getResponse.GetDashboards()))
	assert.Equal(int64(1), getResponse.GetTotal())
	assert.True(getResponse.GetDashboards()[0].GetIsReadOnly())
	assert.True(getResponse.GetDashboards()[0].GetIsShared())
	assert.Equal(customScreenboardID, getResponse.GetDashboards()[0].GetId())
	assert.Equal(datadog.DASHBOARDTYPE_CUSTOM_SCREENBOARD, getResponse.GetDashboards()[0].GetType())
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
			api := datadog.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.GetDashboardListItems(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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
			api := datadog.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.CreateDashboardListItems(ctx, tc.ID, *datadog.NewDashboardListAddItemsRequest())
			assert.IsType(common.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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
			api := datadog.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.UpdateDashboardListItems(ctx, tc.ID, *datadog.NewDashboardListUpdateItemsRequest())
			assert.IsType(common.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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
			api := datadog.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.DeleteDashboardListItems(ctx, tc.ID, *datadog.NewDashboardListDeleteItemsRequest())
			assert.IsType(common.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
