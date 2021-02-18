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

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestDashboardListLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testDashboardList := datadog.DashboardList{
		Name: *tests.UniqueEntityName(ctx, t),
	}

	// Create downtime
	dashboardList, httpresp, err := Client(ctx).DashboardListsApi.CreateDashboardList(ctx).Body(testDashboardList).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard list %v: Response %s: %v", testDashboardList, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboardList(ctx, t, dashboardList.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(testDashboardList.GetName(), dashboardList.GetName())

	// Edit a downtime
	editedDashboardList := datadog.DashboardList{Name: fmt.Sprintf("%s-updated", testDashboardList.GetName())}
	updatedDashboardList, httpresp, err := Client(ctx).DashboardListsApi.UpdateDashboardList(ctx, dashboardList.GetId()).Body(editedDashboardList).Execute()
	if err != nil {
		t.Errorf("Error updating dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDashboardList.GetName(), updatedDashboardList.GetName())

	// Check downtime existence
	fetchedDashboardList, httpresp, err := Client(ctx).DashboardListsApi.GetDashboardList(ctx, dashboardList.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDashboardList.GetName(), fetchedDashboardList.GetName())

	// Find our downtime in the full list
	dashboardLists, httpresp, err := Client(ctx).DashboardListsApi.ListDashboardLists(ctx).Execute()
	if err != nil {
		t.Errorf("Error fetching dashboard lists: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(dashboardLists.GetDashboardLists(), fetchedDashboardList)

	// Cancel downtime
	deletedDashboardListResponse, httpresp, err := Client(ctx).DashboardListsApi.DeleteDashboardList(ctx, dashboardList.GetId()).Execute()
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

			_, httpresp, err := Client(ctx).DashboardListsApi.ListDashboardLists(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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
		Body               datadog.DashboardList
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.DashboardList{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.DashboardList{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DashboardListsApi.CreateDashboardList(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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

			_, httpresp, err := Client(ctx).DashboardListsApi.GetDashboardList(ctx, 1234).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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
		Body               datadog.DashboardList
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.DashboardList{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.DashboardList{}, 403},
		"404 Not Found":   {WithTestAuth, datadog.DashboardList{Name: "nonexistent"}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DashboardListsApi.UpdateDashboardList(ctx, 1234).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
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

			_, httpresp, err := Client(ctx).DashboardListsApi.DeleteDashboardList(ctx, 1234).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteDashboardList(ctx context.Context, t *testing.T, dashboardListID int64) {
	_, httpresp, err := Client(ctx).DashboardListsApi.DeleteDashboardList(ctx, dashboardListID).Execute()
	if err != nil {
		t.Logf("Deleting dashboard list: %v failed with %v, Another test may have already deleted this dashboard list: %v", dashboardListID, httpresp.StatusCode, err)
	}
}
