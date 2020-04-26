/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestDashboardListLifecycle(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	start := c.Clock.Now().Unix()
	testDashboardList := datadog.DashboardList{
		Name: fmt.Sprintf("go dashboard list %d", start),
	}

	// Create downtime
	dashboardList, httpresp, err := c.Client.DashboardListsApi.CreateDashboardList(c.Ctx).Body(testDashboardList).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard list %v: Response %s: %v", testDashboardList, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboardList(c, dashboardList.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	assert.Equal(t, testDashboardList.GetName(), dashboardList.GetName())

	// Edit a downtime
	editedDashboardList := datadog.DashboardList{Name: fmt.Sprintf("go dashboard list updated %d", start)}
	updatedDashboardList, httpresp, err := c.Client.DashboardListsApi.UpdateDashboardList(c.Ctx, dashboardList.GetId()).Body(editedDashboardList).Execute()
	if err != nil {
		t.Errorf("Error updating dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, editedDashboardList.GetName(), updatedDashboardList.GetName())

	// Check downtime existence
	fetchedDashboardList, httpresp, err := c.Client.DashboardListsApi.GetDashboardList(c.Ctx, dashboardList.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, editedDashboardList.GetName(), fetchedDashboardList.GetName())

	// Find our downtime in the full list
	dashboardLists, httpresp, err := c.Client.DashboardListsApi.ListDashboardLists(c.Ctx).Execute()
	if err != nil {
		t.Errorf("Error fetching dashboard lists: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, dashboardLists.GetDashboardLists(), fetchedDashboardList)

	// Cancel downtime
	deletedDashboardListResponse, httpresp, err := c.Client.DashboardListsApi.DeleteDashboardList(c.Ctx, dashboardList.GetId()).Execute()
	if err != nil {
		t.Errorf("Error deleting dashboard list %v: Response %s: %v", dashboardList.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, dashboardList.GetId(), deletedDashboardListResponse.GetDeletedDashboardListId())
}

func TestDashboardListListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.DashboardListsApi.ListDashboardLists(c.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.DashboardListsApi.CreateDashboardList(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.DashboardListsApi.GetDashboardList(c.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.DashboardListsApi.UpdateDashboardList(c.Ctx, 1234).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDashboardListDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.DashboardListsApi.DeleteDashboardList(c.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func deleteDashboardList(c *Client, dashboardListID int64) {
	_, httpresp, err := c.Client.DashboardListsApi.DeleteDashboardList(c.Ctx, dashboardListID).Execute()
	if err != nil {
		log.Printf("Deleting dashboard list: %v failed with %v, Another test may have already deleted this dashboard list: %v", dashboardListID, httpresp.StatusCode, err)
	}
}
