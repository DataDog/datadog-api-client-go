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
)

func TestDashboardListLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewDashboardListsApi(Client(ctx))

	testDashboardList := datadogV1.DashboardList{
		Name: *tests.UniqueEntityName(ctx, t),
	}

	// Create downtime
	dashboardList, httpresp, err := api.CreateDashboardList(ctx, testDashboardList)
	if err != nil {
		t.Fatalf("Error creating dashboard list %v: Response %s: %v", testDashboardList, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboardList(ctx, t, dashboardList.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(testDashboardList.GetName(), dashboardList.GetName())

	// Edit a downtime
	editedDashboardList := datadogV1.DashboardList{Name: fmt.Sprintf("%s-updated", testDashboardList.GetName())}
	updatedDashboardList, httpresp, err := api.UpdateDashboardList(ctx, dashboardList.GetId(), editedDashboardList)
	if err != nil {
		t.Errorf("Error updating dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDashboardList.GetName(), updatedDashboardList.GetName())

	// Check downtime existence
	fetchedDashboardList, httpresp, err := api.GetDashboardList(ctx, dashboardList.GetId())
	if err != nil {
		t.Errorf("Error fetching dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDashboardList.GetName(), fetchedDashboardList.GetName())

	// Find our downtime in the full list
	dashboardLists, httpresp, err := api.ListDashboardLists(ctx)
	if err != nil {
		t.Errorf("Error fetching dashboard lists: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(dashboardLists.GetDashboardLists(), fetchedDashboardList)

	// Cancel downtime
	deletedDashboardListResponse, httpresp, err := api.DeleteDashboardList(ctx, dashboardList.GetId())
	if err != nil {
		t.Errorf("Error deleting dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(dashboardList.GetId(), deletedDashboardListResponse.GetDeletedDashboardListId())
}

func TestDashboardListListErrors(t *testing.T) {
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
			api := datadogV1.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.ListDashboardLists(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.DashboardList
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.DashboardList{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.DashboardList{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.CreateDashboardList(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListGetErrors(t *testing.T) {
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
			api := datadogV1.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.GetDashboardList(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.DashboardList
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.DashboardList{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.DashboardList{}, 403},
		"404 Not Found":   {WithTestAuth, datadogV1.DashboardList{Name: "nonexistent"}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.UpdateDashboardList(ctx, 1234, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListDeleteErrors(t *testing.T) {
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
			api := datadogV1.NewDashboardListsApi(Client(ctx))

			_, httpresp, err := api.DeleteDashboardList(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteDashboardList(ctx context.Context, t *testing.T, dashboardListID int64) {
	api := datadogV1.NewDashboardListsApi(Client(ctx))

	_, httpresp, err := api.DeleteDashboardList(ctx, dashboardListID)
	if err != nil {
		t.Logf("Deleting dashboard list: %v failed with %v, Another test may have already deleted this dashboard list: %v", dashboardListID, httpresp.StatusCode, err)
	}
}
