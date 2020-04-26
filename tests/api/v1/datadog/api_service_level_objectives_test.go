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
	"gopkg.in/h2non/gock.v1"
)

var testServiceCheckMonitor = datadog.Monitor{
	Name:    datadog.PtrString("name for service check monitor in Go client"),
	Type:    datadog.MONITORTYPE_SERVICE_CHECK.Ptr(),
	Query:   datadog.PtrString("\"datadog.agent.check_status\".over(\"database\").last(2).count_by_status()"),
	Message: datadog.PtrString("some message Notify: @hipchat-channel"),
	Tags: &[]string{
		"test",
		"client:go",
	},
}

var testMonitorSLO = datadog.ServiceLevelObjective{
	Type:        "monitor",
	Name:        "Go client Critical Foo Host Uptime",
	Description: *datadog.NewNullableString(datadog.PtrString("Track the uptime of host foo which is critical to us.")),
	Tags:        &[]string{"app:core", "kpi"},
	Thresholds: []datadog.SLOThreshold{{
		Timeframe: datadog.SLOTIMEFRAME_THIRTY_DAYS,
		Target:    95.0,
		Warning:   datadog.PtrFloat64(98.0),
	}},
}

var testEventSLO = datadog.ServiceLevelObjective{
	Type:        "metric",
	Name:        "Go client HTTP Return Codes",
	Description: *datadog.NewNullableString(datadog.PtrString("Make sure we don't have too many failed HTTP responses.")),
	Tags:        &[]string{"app:httpd"},
	Thresholds: []datadog.SLOThreshold{{
		Timeframe: datadog.SLOTIMEFRAME_SEVEN_DAYS,
		Target:    95.0,
		Warning:   datadog.PtrFloat64(98.0),
	}},
	Query: &datadog.ServiceLevelObjectiveQuery{
		Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
		Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
	},
}

func TestSLOMonitorLifecycle(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// Create monitor to reference in the SLO
	monitor, httpresp, err := c.Client.MonitorsApi.CreateMonitor(c.Ctx).Body(testServiceCheckMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create SLO
	sloResp, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(slo.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testMonitorSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.UpdateSLO(c.Ctx, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal(t, "Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(c.Ctx).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal(t, []string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadog.SLOResponse
	sloGetResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.GetSLO(c.Ctx, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(t, slo2.GetName(), slo3.GetName())

	// Get SLO history
	now := c.Clock.Now().Unix()
	_, httpresp, err = c.Client.ServiceLevelObjectivesApi.GetSLOHistory(c.Ctx, slo3.GetId()).
		FromTs(now - 11).ToTs(now - 1).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadog.SLODeleteResponse
	sloDeleteResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.DeleteSLO(c.Ctx, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, []string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOEventLifecycle(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// Create SLO
	sloResp, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(slo.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testEventSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.UpdateSLO(c.Ctx, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal(t, "Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(c.Ctx).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal(t, []string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadog.SLOResponse
	sloGetResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.GetSLO(c.Ctx, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(t, slo2.GetName(), slo3.GetName())

	// Get SLO history
	now := c.Clock.Now().Unix()
	_, httpresp, err = c.Client.ServiceLevelObjectivesApi.GetSLOHistory(c.Ctx, slo3.GetId()).
		FromTs(now - 11).ToTs(now - 1).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadog.SLODeleteResponse
	sloDeleteResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.DeleteSLO(c.Ctx, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, []string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOMultipleInstances(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// Create monitor to reference in the monitor SLO
	monitor, httpresp, err := c.Client.MonitorsApi.CreateMonitor(c.Ctx).Body(testServiceCheckMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create monitor SLO
	sloResp, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	monitorSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(monitorSLO.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Create event SLO
	sloResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	eventSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(eventSLO.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	// Get multiple SLOs
	var slosResp datadog.SLOListResponse
	slosResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.ListSLOs(c.Ctx).
		Ids(fmt.Sprintf("%s,%s", monitorSLO.GetId(), eventSLO.GetId())).Execute()
	if err != nil {
		t.Fatalf("Error getting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	slos := slosResp.GetData()
	assertSLOIDPresent(t, monitorSLO.GetId(), slos)
	assertSLOIDPresent(t, eventSLO.GetId(), slos)

	// Use bulk delete operation to delete the event SLO
	var deleteResp datadog.SLOBulkDeleteResponse
	deleteResp, httpresp, err = c.Client.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(c.Ctx).
		Body(map[string][]datadog.SLOTimeframe{
			eventSLO.GetId(): {datadog.SLOTIMEFRAME_SEVEN_DAYS},
		}).Execute()

	if err != nil {
		t.Fatalf("Error bulk deleting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	deleteData := deleteResp.GetData()
	assert.Equal(t, []string{eventSLO.GetId()}, deleteData.GetDeleted())
	assert.Empty(t, deleteData.GetUpdated())
}

func TestSLOCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(datadog.ServiceLevelObjective{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLOListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		Ids                string
		ExpectedStatusCode int
	}{
		"400 Bad Request", c.Ctx, "id1,id1": {400},
		"403 Forbidden", context.Background(), "id1,id1": {403},
		"404 Not Found", c.Ctx, "id1,id2": {404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.ListSLOs(c.Ctx).Ids(tc.Ids).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLOUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		Body               datadog.ServiceLevelObjective
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ServiceLevelObjective{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ServiceLevelObjective{}, 403},
		"404 Not Found":   {WithTestAuth, testEventSLO, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.UpdateSLO(c.Ctx, "id").Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLOGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.GetSLO(c.Ctx, "id").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLODeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.DeleteSLO(c.Ctx, "id").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLODelete409Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClient(WithFakeAuth(context.Background()), t)
	defer c.Close()

	// Mocked as the feature is in a broken state right now
	res, err := tests.ReadFixture("fixtures/slo/error_409_delete.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	gock.New("https://api.datadoghq.com").Delete("/api/v1/slo/id").Reply(409).JSON(res)
	defer gock.Off()

	// FIXME: Make it an integration test when feature is fixed
	//// Create SLO and reference it in a dashboard to trigger 409
	//sloResp, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testEventSLO).Execute()
	//if err != nil {
	//	t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//slo := sloResp.GetData()[0]
	//defer deleteSLOIfExists(slo.GetId())
	//assert.Equal(t, 200, httpresp.StatusCode)
	//
	//// SLO Widget
	//sloWidgetDefinition := datadog.NewSLOWidgetDefinitionWithDefaults()
	//sloWidgetDefinition.SetSloId(slo.GetId())
	//sloWidgetDefinition.SetShowErrorBudget(true)
	//sloWidgetDefinition.SetViewMode(datadog.WIDGETVIEWMODE_BOTH)
	//sloWidgetDefinition.SetTimeWindows([]datadog.WidgetTimeWindows{
	//	datadog.WIDGETTIMEWINDOWS_SEVEN_DAYS,
	//})
	//
	//sloWidget := *datadog.NewWidget(sloWidgetDefinition.AsWidgetDefinition())
	//orderedWidgetList := []datadog.Widget{sloWidget}
	//
	//dashboard := *datadog.NewDashboardWithDefaults()
	//dashboard.SetLayoutType(datadog.DASHBOARDLAYOUTTYPE_ORDERED)
	//dashboard.SetWidgets(orderedWidgetList)
	//dashboard.SetTitle("Go Client Test SLO Widget Dashboard")
	//createdDashboard, httpresp, err := c.Client.DashboardsApi.CreateDashboard(c.Ctx).Body(dashboard).Execute()
	//if err != nil {
	//	t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//defer deleteDashboard(createdDashboard.GetId())
	//assert.Equal(t, 200, httpresp.StatusCode)

	_, httpresp, err := c.Client.ServiceLevelObjectivesApi.DeleteSLO(c.Ctx, "id").Execute()
	assert.Equal(t, 409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.SLODeleteResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestSLOHistoryGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	sloResp, httpresp, err := c.Client.ServiceLevelObjectivesApi.CreateSLO(c.Ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(slo.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, slo.GetId(), 400},
		"403 Forbidden", context.Background(), "id": {403},
		"404 Not Found", c.Ctx, "id": {404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.GetSLOHistory(c.Ctx, tc.ID).FromTs(123).ToTs(12).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLOCanDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(c.Ctx).Ids("").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSLOCanDelete409Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClient(WithFakeAuth(context.Background()), t)
	defer c.Close()

	// Mocked as the feature is in a broken state right now
	res, err := tests.ReadFixture("fixtures/slo/error_409_can_delete.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// FIXME: Make it an integration test when feature is fixed
	gock.New("https://api.datadoghq.com").Get("/api/v1/slo/can_delete").Reply(409).JSON(res)
	defer gock.Off()

	_, httpresp, err := c.Client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(c.Ctx).Ids("id").Execute()
	assert.Equal(t, 409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.CheckCanDeleteSLOResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestSLOBulkDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Name               string
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(c.Ctx).Body(map[string][]datadog.SLOTimeframe{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func assertSLOIDPresent(t *testing.T, sloID string, slos []datadog.ServiceLevelObjective) {
	for _, slo := range slos {
		if slo.GetId() == sloID {
			return
		}
	}
	assert.Nil(t, fmt.Errorf("SLO %s expected but not found", sloID))
}

func deleteSLOIfExists(sloID string) {
	_, httpresp, err := c.Client.ServiceLevelObjectivesApi.DeleteSLO(c.Ctx, sloID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting SLO: %v failed with %v, Another test may have already deleted this SLO: %v",
			sloID, httpresp.StatusCode, err)
	}
}
