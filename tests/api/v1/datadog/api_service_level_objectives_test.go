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

	"gopkg.in/h2non/gock.v1"
)

func getTestServiceCheckMonitor(ctx context.Context, t *testing.T) datadog.Monitor {
	return datadog.Monitor{
		Name:    tests.UniqueEntityName(ctx, t),
		Type:    datadog.MONITORTYPE_SERVICE_CHECK.Ptr(),
		Query:   datadog.PtrString("\"datadog.agent.check_status\".over(\"database\").last(2).count_by_status()"),
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: &[]string{
			"test",
			"client:go",
		},
	}
}

func getTestMonitorSLO(ctx context.Context, t *testing.T) datadog.ServiceLevelObjectiveRequest {
	return datadog.ServiceLevelObjectiveRequest{
		Type:        "monitor",
		Name:        *tests.UniqueEntityName(ctx, t),
		Description: *datadog.NewNullableString(datadog.PtrString("Track the uptime of host foo which is critical to us.")),
		Tags:        &[]string{"app:core", "kpi"},
		Thresholds: []datadog.SLOThreshold{{
			Timeframe: datadog.SLOTIMEFRAME_THIRTY_DAYS,
			Target:    95.0,
			Warning:   datadog.PtrFloat64(98.0),
		}},
	}
}

func getTestEventSLO(ctx context.Context, t *testing.T) datadog.ServiceLevelObjectiveRequest {
	return datadog.ServiceLevelObjectiveRequest{
		Type:        "metric",
		Name:        *tests.UniqueEntityName(ctx, t),
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
}

func isSLOIDPresent(sloID string, slos []datadog.ServiceLevelObjective) error {
	for _, slo := range slos {
		if slo.GetId() == sloID {
			return nil
		}
	}
	return fmt.Errorf("SLO %s expected but not found", sloID)
}

func TestSLOMonitorLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create monitor to reference in the SLO
	testServiceCheckMonitor := getTestServiceCheckMonitor(ctx, t)
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(testServiceCheckMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testMonitorSLO := getTestMonitorSLO(ctx, t)
	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create SLO
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, slo.GetId())
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testMonitorSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.UpdateSLO(ctx, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal("Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal([]string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadog.SLOResponse
	sloGetResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.GetSLO(ctx, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(slo2.GetName(), slo3.GetName())

	// Get SLO history
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetSLOHistory", true)
	now := tests.ClockFromContext(ctx).Now().Unix()
	_, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.GetSLOHistory(ctx, slo3.GetId()).
		FromTs(now - 11).ToTs(now - 1).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadog.SLODeleteResponse
	sloDeleteResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.DeleteSLO(ctx, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal([]string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOEventLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, slo.GetId())
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testEventSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.UpdateSLO(ctx, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal("Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal([]string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadog.SLOResponse
	sloGetResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.GetSLO(ctx, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(slo2.GetName(), slo3.GetName())

	// Get SLO history
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetSLOHistory", true)
	now := tests.ClockFromContext(ctx).Now().Unix()
	_, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.GetSLOHistory(ctx, slo3.GetId()).
		FromTs(now - 11).ToTs(now - 1).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadog.SLODeleteResponse
	sloDeleteResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.DeleteSLO(ctx, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal([]string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOMultipleInstances(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create monitor to reference in the monitor SLO
	testServiceCheckMonitor := getTestServiceCheckMonitor(ctx, t)
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(testServiceCheckMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testMonitorSLO := getTestMonitorSLO(ctx, t)
	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create monitor SLO
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	monitorSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, monitorSLO.GetId())
	assert.Equal(200, httpresp.StatusCode)

	// Create event SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	eventSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, eventSLO.GetId())
	assert.Equal(200, httpresp.StatusCode)

	// Get multiple SLOs
	var slosResp datadog.SLOListResponse
	slosResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.ListSLOs(ctx).
		Ids(fmt.Sprintf("%s,%s", monitorSLO.GetId(), eventSLO.GetId())).Execute()
	if err != nil {
		t.Fatalf("Error getting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slos := slosResp.GetData()
	assert.NoError(isSLOIDPresent(monitorSLO.GetId(), slos))
	assert.NoError(isSLOIDPresent(eventSLO.GetId(), slos))

	// Use bulk delete operation to delete the event SLO
	var deleteResp datadog.SLOBulkDeleteResponse
	deleteResp, httpresp, err = Client(ctx).ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(ctx).
		Body(map[string][]datadog.SLOTimeframe{
			eventSLO.GetId(): {datadog.SLOTIMEFRAME_SEVEN_DAYS},
		}).Execute()

	if err != nil {
		t.Fatalf("Error bulk deleting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	deleteData := deleteResp.GetData()
	assert.Equal([]string{eventSLO.GetId()}, deleteData.GetDeleted())
	assert.Empty(deleteData.GetUpdated())
}

func TestSLOCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(datadog.ServiceLevelObjectiveRequest{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Ids                string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, "id1,id1", 400},
		"403 Forbidden":   {WithFakeAuth, "id1,id1", 403},
		"404 Not Found":   {WithTestAuth, "id1,id2", 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.ListSLOs(ctx).Ids(tc.Ids).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.ServiceLevelObjective
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.ServiceLevelObjective{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.ServiceLevelObjective{}, 403},
		"404 Not Found": {WithTestAuth, datadog.ServiceLevelObjective{
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
		}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.UpdateSLO(ctx, "id").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOGetErrors(t *testing.T) {
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
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.GetSLO(ctx, "id").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLODeleteErrors(t *testing.T) {
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
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.DeleteSLO(ctx, "id").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLODelete409Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Mocked as the feature is in a broken state right now
	res, err := tests.ReadFixture("fixtures/slo/error_409_delete.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/slo/id").Reply(409).JSON(res)
	defer gock.Off()

	// FIXME: Make it an integration test when feature is fixed
	//// Create SLO and reference it in a dashboard to trigger 409
	//sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	//if err != nil {
	//	t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//slo := sloResp.GetData()[0]
	//defer deleteSLOIfExists(ctx, slo.GetId())
	//assert.Equal(200, httpresp.StatusCode)
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
	//createdDashboard, httpresp, err := Client(ctx).DashboardsApi.CreateDashboard(ctx).Body(dashboard).Execute()
	//if err != nil {
	//	t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//defer deleteDashboard(createdDashboard.GetId())
	//assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.DeleteSLO(ctx, "id").Execute()
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.SLODeleteResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSLOHistoryGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.GreaterOrEqual(len(sloResp.GetData()), 1)
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, slo.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, slo.GetId(), 400},
		"403 Forbidden":   {WithFakeAuth, "id", 403},
		"404 Not Found":   {WithTestAuth, "id", 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			Client(ctx).GetConfig().SetUnstableOperationEnabled("GetSLOHistory", true)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.GetSLOHistory(ctx, tc.ID).FromTs(123).ToTs(12).Execute()
			assert.Error(err)
			assert.NotNil(httpresp, err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOCanDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx).Ids("").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOCanDelete409Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Mocked as the feature is in a broken state right now
	res, err := tests.ReadFixture("fixtures/slo/error_409_can_delete.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// FIXME: Make it an integration test when feature is fixed
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/slo/can_delete").Reply(409).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx).Ids("id").Execute()
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.CheckCanDeleteSLOResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSLOBulkDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(ctx).Body(map[string][]datadog.SLOTimeframe{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteSLOIfExists(ctx context.Context, sloID string) {
	_, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.DeleteSLO(ctx, sloID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting SLO: %v failed with %v, Another test may have already deleted this SLO: %v",
			sloID, httpresp.StatusCode, err)
	}
}

func TestSLOCorrectionsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	Client(ctx).GetConfig().SetUnstableOperationEnabled("CreateSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("GetSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("ListSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("UpdateSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("DeleteSLOCorrection", true)

	// Create SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, slo.GetId())

	testSLOCorrectionCreateData := datadog.NewSLOCorrectionCreateRequestData()
	now := tests.ClockFromContext(ctx).Now().Unix()
	testSLOCorrectionCreateAttributes := datadog.SLOCorrectionCreateRequestAttributes{
		Timezone: "UTC",
		SloId: slo.GetId(),
		Category: datadog.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE,
		Start: now,
		End: now + 3600,
	}
	testSLOCorrectionCreateData.SetAttributes(testSLOCorrectionCreateAttributes)
	testSLOCorrectionCreate := datadog.SLOCorrectionCreateRequest{
		Data:        testSLOCorrectionCreateData,
	}

	sloCorrectionResp, httpresp, err := Client(ctx).ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection(ctx).Body(testSLOCorrectionCreate).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO Correction %v: Response %s: %v", testSLOCorrectionCreate, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrection := sloCorrectionResp.GetData()
	defer deleteSLOCorrectionIfExists(ctx, sloCorrection.GetId())

	sloCorrectionListResp, httpresp, err := Client(ctx).ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO corrections: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	sloCorrections := sloCorrectionListResp.GetData()
	assert.NoError(isSLOCorrecionIDPresent(sloCorrection.GetId(), sloCorrections))

	sloCorrectionGetResp, httpresp, err := Client(ctx).ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection(ctx, sloCorrection.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO correction %s: Response %s: %v", sloCorrection.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrectionGetData := sloCorrectionGetResp.GetData()
	assert.Equal(sloCorrectionGetData.GetId(), sloCorrection.GetId())
	sloCorrectionAttributes := sloCorrectionGetData.GetAttributes()
	assert.Equal(sloCorrectionAttributes.GetTimezone(), "UTC")
	assert.Equal(sloCorrectionAttributes.GetCategory(), datadog.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE)

	testSLOCorrectionUpdateData := datadog.NewSLOCorrectionUpdateRequestData()
	testSLOCorrectionUpdateAttributes := datadog.SLOCorrectionUpdateRequestAttributes{
		Timezone: "UTC",
		Category: datadog.SLOCORRECTIONCATEGORY_OTHER,
		Start: now,
		End: now + 3600,
	}
	testSLOCorrectionUpdateData.SetAttributes(testSLOCorrectionUpdateAttributes)
	testSLOCorrectionUpdate := datadog.SLOCorrectionUpdateRequest{
		Data:        testSLOCorrectionUpdateData,
	}

	sloCorrectionUpdateResp, httpresp, err := Client(ctx).ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection(ctx, sloCorrection.GetId()).Body(testSLOCorrectionUpdate).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO correction %v: Response %s: %v", testSLOCorrectionUpdate, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrectionUpdateData := sloCorrectionUpdateResp.GetData()
	sloCorrectionAttributes = sloCorrectionUpdateData.GetAttributes()
	assert.Equal(sloCorrectionAttributes.GetCategory(), datadog.SLOCORRECTIONCATEGORY_OTHER)

	httpresp, err = Client(ctx).ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection(ctx, sloCorrection.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO correction %s: Response %s: %v", sloCorrection.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)
}

func deleteSLOCorrectionIfExists(ctx context.Context, sloCorrectionID string) {
	httpresp, err := Client(ctx).ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection(ctx, sloCorrectionID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting SLO correction: %v failed with %v, Another test may have already deleted this SLO correction: %v",
			sloCorrectionID, httpresp.StatusCode, err)
	}
}

func isSLOCorrecionIDPresent(sloCorrectionID string, sloCorrections []datadog.SLOCorrectionListResponseData) error {
	for _, sloCorrection := range sloCorrections {
		if sloCorrection.GetId() == sloCorrectionID {
			return nil
		}
	}
	return fmt.Errorf("SLO correction %s expected but not found", sloCorrectionID)
}
