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

	"gopkg.in/h2non/gock.v1"
)

func getTestServiceCheckMonitor(ctx context.Context, t *testing.T) datadogV1.Monitor {
	return datadogV1.Monitor{
		Name:    tests.UniqueEntityName(ctx, t),
		Type:    datadogV1.MONITORTYPE_SERVICE_CHECK,
		Query:   "\"datadog.agent.check_status\".over(\"database\").by(\"*\").last(2).count_by_status()",
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test",
			"client:go",
		},
	}
}

func getTestMonitorSLO(ctx context.Context, t *testing.T) datadogV1.ServiceLevelObjectiveRequest {
	return datadogV1.ServiceLevelObjectiveRequest{
		Type:        "monitor",
		Name:        *tests.UniqueEntityName(ctx, t),
		Description: *datadog.NewNullableString(datadog.PtrString("Track the uptime of host foo which is critical to us.")),
		Tags:        []string{"app:core", "kpi"},
		Thresholds: []datadogV1.SLOThreshold{{
			Timeframe: datadogV1.SLOTIMEFRAME_THIRTY_DAYS,
			Target:    95.0,
			Warning:   datadog.PtrFloat64(98.0),
		}},
	}
}

func getTestEventSLO(ctx context.Context, t *testing.T) datadogV1.ServiceLevelObjectiveRequest {
	return datadogV1.ServiceLevelObjectiveRequest{
		Type:        "metric",
		Name:        *tests.UniqueEntityName(ctx, t),
		Description: *datadog.NewNullableString(datadog.PtrString("Make sure we don't have too many failed HTTP responses.")),
		Tags:        []string{"app:httpd"},
		Thresholds: []datadogV1.SLOThreshold{{
			Timeframe: datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
			Target:    95.0,
			Warning:   datadog.PtrFloat64(98.0),
		}},
		Query: &datadogV1.ServiceLevelObjectiveQuery{
			Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
			Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
		},
	}
}

func isSLOIDPresent(sloID string, slos []datadogV1.ServiceLevelObjective) error {
	for _, slo := range slos {
		if slo.GetId() == sloID {
			return nil
		}
	}
	return fmt.Errorf("SLO %s expected but not found", sloID)
}

func TestSLOMonitorLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewMonitorsApi(Client(ctx))
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

	// Create monitor to reference in the SLO
	testServiceCheckMonitor := getTestServiceCheckMonitor(ctx, t)
	monitor, httpresp, err := api.CreateMonitor(ctx, testServiceCheckMonitor)
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testMonitorSLO := getTestMonitorSLO(ctx, t)
	testMonitorSLO.MonitorIds = []int64{*monitor.Id}

	// Create SLO
	sloResp, httpresp, err := sloApi.CreateSLO(ctx, testMonitorSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, slo.GetId())
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testMonitorSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = sloApi.UpdateSLO(ctx, slo.GetId(), slo)
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal("Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadogV1.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = sloApi.CheckCanDeleteSLO(ctx, slo2.GetId())
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal([]string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadogV1.SLOResponse
	sloGetResp, httpresp, err = sloApi.GetSLO(ctx, slo2.GetId())
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(slo2.GetName(), slo3.GetName())

	// Get SLO history
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSLOHistory", true)
	now := tests.ClockFromContext(ctx).Now().Unix()
	_, httpresp, err = sloApi.GetSLOHistory(ctx, slo3.GetId(), now-11, now-1)
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadogV1.SLODeleteResponse
	sloDeleteResp, httpresp, err = sloApi.DeleteSLO(ctx, slo3.GetId())
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal([]string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOEventLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

	// Create SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := sloApi.CreateSLO(ctx, testEventSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, slo.GetId())
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testEventSLO.GetName(), slo.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = sloApi.UpdateSLO(ctx, slo.GetId(), slo)
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo2 := sloResp.GetData()[0]
	assert.Equal("Updated description", slo2.GetDescription())

	// Check that the SLO can be deleted
	var canDeleteResp datadogV1.CheckCanDeleteSLOResponse
	canDeleteResp, httpresp, err = sloApi.CheckCanDeleteSLO(ctx, slo2.GetId())
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	canDelete := canDeleteResp.GetData()
	assert.Equal([]string{slo2.GetId()}, canDelete.GetOk())

	// Get SLO
	var sloGetResp datadogV1.SLOResponse
	sloGetResp, httpresp, err = sloApi.GetSLO(ctx, slo2.GetId())
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slo3 := sloGetResp.GetData()
	assert.Equal(slo2.GetName(), slo3.GetName())

	// Get SLO history
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSLOHistory", true)
	now := tests.ClockFromContext(ctx).Now().Unix()
	_, httpresp, err = sloApi.GetSLOHistory(ctx, slo3.GetId(), now-11, now-1)
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// Delete SLO
	var sloDeleteResp datadogV1.SLODeleteResponse
	sloDeleteResp, httpresp, err = sloApi.DeleteSLO(ctx, slo3.GetId())
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal([]string{slo3.GetId()}, sloDeleteResp.GetData())
}

func TestSLOMultipleInstances(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewMonitorsApi(Client(ctx))
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

	// Create monitor to reference in the monitor SLO
	testServiceCheckMonitor := getTestServiceCheckMonitor(ctx, t)
	monitor, httpresp, err := api.CreateMonitor(ctx, testServiceCheckMonitor)
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testServiceCheckMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testMonitorSLO := getTestMonitorSLO(ctx, t)
	testMonitorSLO.MonitorIds = []int64{*monitor.Id}

	// Create monitor SLO
	sloResp, httpresp, err := sloApi.CreateSLO(ctx, testMonitorSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	monitorSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, monitorSLO.GetId())
	assert.Equal(200, httpresp.StatusCode)

	// Create event SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err = sloApi.CreateSLO(ctx, testEventSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	eventSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, eventSLO.GetId())
	assert.Equal(200, httpresp.StatusCode)

	// Get multiple SLOs
	var slosResp datadogV1.SLOListResponse
	slosResp, httpresp, err = sloApi.ListSLOs(ctx, *datadogV1.NewListSLOsOptionalParameters().
		WithIds(fmt.Sprintf("%s,%s", monitorSLO.GetId(), eventSLO.GetId())))
	if err != nil {
		t.Fatalf("Error getting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	slos := slosResp.GetData()
	assert.NoError(isSLOIDPresent(monitorSLO.GetId(), slos))
	assert.NoError(isSLOIDPresent(eventSLO.GetId(), slos))

	// Use bulk delete operation to delete the event SLO
	var deleteResp datadogV1.SLOBulkDeleteResponse
	deleteResp, httpresp, err = sloApi.DeleteSLOTimeframeInBulk(ctx,
		map[string][]datadogV1.SLOTimeframe{
			eventSLO.GetId(): {datadogV1.SLOTIMEFRAME_SEVEN_DAYS},
		})

	if err != nil {
		t.Fatalf("Error bulk deleting SLOs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	deleteData := deleteResp.GetData()
	assert.Equal([]string{eventSLO.GetId()}, deleteData.GetDeleted())
	assert.Empty(deleteData.GetUpdated())
}

func TestSLOCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.CreateSLO(ctx, datadogV1.ServiceLevelObjectiveRequest{})
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.ListSLOs(ctx, *datadogV1.NewListSLOsOptionalParameters().WithIds(tc.Ids))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.ServiceLevelObjective
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.ServiceLevelObjective{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.ServiceLevelObjective{}, 403},
		"404 Not Found": {WithTestAuth, datadogV1.ServiceLevelObjective{
			Type:        "metric",
			Name:        "Go client HTTP Return Codes",
			Description: *datadog.NewNullableString(datadog.PtrString("Make sure we don't have too many failed HTTP responses.")),
			Tags:        []string{"app:httpd"},
			Thresholds: []datadogV1.SLOThreshold{{
				Timeframe: datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
				Target:    95.0,
				Warning:   datadog.PtrFloat64(98.0),
			}},
			Query: &datadogV1.ServiceLevelObjectiveQuery{
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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.UpdateSLO(ctx, "id", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOGetErrors(t *testing.T) {
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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.GetSLO(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLODeleteErrors(t *testing.T) {
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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.DeleteSLO(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLODelete409Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

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
	//sloResp, httpresp, err := sloApi.CreateSLO(ctx).Body(testEventSLO)
	//if err != nil {
	//	t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//slo := sloResp.GetData()[0]
	//defer deleteSLOIfExists(ctx, t, slo.GetId())
	//assert.Equal(200, httpresp.StatusCode)
	//
	//// SLO Widget
	//sloWidgetDefinition := datadogV1.NewSLOWidgetDefinitionWithDefaults()
	//sloWidgetDefinition.SetSloId(slo.GetId())
	//sloWidgetDefinition.SetShowErrorBudget(true)
	//sloWidgetDefinition.SetViewMode(datadogV1.WIDGETVIEWMODE_BOTH)
	//sloWidgetDefinition.SetTimeWindows([]datadogV1.WidgetTimeWindows{
	//	datadogV1.WIDGETTIMEWINDOWS_SEVEN_DAYS,
	//})
	//
	//sloWidget := *datadogV1.NewWidget(sloWidgetDefinition.AsWidgetDefinition())
	//orderedWidgetList := []datadogV1.Widget{sloWidget}
	//
	//dashboard := *datadogV1.NewDashboardWithDefaults()
	//dashboard.SetLayoutType(datadogV1.DASHBOARDLAYOUTTYPE_ORDERED)
	//dashboard.SetWidgets(orderedWidgetList)
	//dashboard.SetTitle("Go Client Test SLO Widget Dashboard")
	//createdDashboard, httpresp, err := Client(ctx).DashboardsApi.CreateDashboard(ctx).Body(dashboard)
	//if err != nil {
	//	t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	//}
	//defer deleteDashboard(createdDashboard.GetId())
	//assert.Equal(200, httpresp.StatusCode)

	_, httpresp, err := sloApi.DeleteSLO(ctx, "id")
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.SLODeleteResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSLOHistoryGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := sloApi.CreateSLO(ctx, testEventSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.GreaterOrEqual(len(sloResp.GetData()), 1)
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, slo.GetId())
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
			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSLOHistory", true)
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.GetSLOHistory(ctx, tc.ID, 123, 12)
			assert.Error(err)
			assert.NotNil(httpresp, err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOCanDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.CheckCanDeleteSLO(ctx, "")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSLOCanDelete409Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

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

	_, httpresp, err := sloApi.CheckCanDeleteSLO(ctx, "id")
	assert.Equal(409, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.CheckCanDeleteSLOResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSLOBulkDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

			_, httpresp, err := sloApi.DeleteSLOTimeframeInBulk(ctx, map[string][]datadogV1.SLOTimeframe{})
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteSLOIfExists(ctx context.Context, t *testing.T, sloID string) {
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))

	_, httpresp, err := sloApi.DeleteSLO(ctx, sloID)
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting SLO: %v failed with %v, Another test may have already deleted this SLO: %v",
			sloID, httpresp.StatusCode, err)
	}
}

func TestSLOCorrectionsLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))
	sloCorrectionApi := datadogV1.NewServiceLevelObjectiveCorrectionsApi(Client(ctx))

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.CreateSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.ListSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.UpdateSLOCorrection", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.DeleteSLOCorrection", true)

	// Create SLO
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, _, err := sloApi.CreateSLO(ctx, testEventSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testEventSLO, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, slo.GetId())

	testSLOCorrectionCreateData := datadogV1.NewSLOCorrectionCreateDataWithDefaults()
	now := tests.ClockFromContext(ctx).Now().Unix()
	testTimezone := "UTC"
	testStart := now
	testEnd := now + 3600
	testCategory := datadogV1.SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE
	testSLOCorrectionCreateAttributes := datadogV1.SLOCorrectionCreateRequestAttributes{
		Timezone: &testTimezone,
		SloId:    slo.GetId(),
		Category: testCategory,
		Start:    testStart,
		End:      &testEnd,
	}
	testSLOCorrectionCreateData.SetAttributes(testSLOCorrectionCreateAttributes)
	testSLOCorrectionCreate := datadogV1.SLOCorrectionCreateRequest{
		Data: testSLOCorrectionCreateData,
	}

	sloCorrectionResp, httpresp, err := sloCorrectionApi.CreateSLOCorrection(ctx, testSLOCorrectionCreate)
	if err != nil {
		t.Fatalf("Error creating SLO Correction %v: Response %s: %v", testSLOCorrectionCreate, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrection := sloCorrectionResp.GetData()
	defer deleteSLOCorrectionIfExists(ctx, t, sloCorrection.GetId())

	sloCorrectionListResp, httpresp, err := sloCorrectionApi.ListSLOCorrection(ctx)
	if err != nil {
		t.Fatalf("Error getting SLO corrections: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	sloCorrections := sloCorrectionListResp.GetData()
	assert.NoError(isSLOCorrectionIDPresent(sloCorrection.GetId(), sloCorrections))

	sloCorrectionGetResp, httpresp, err := sloCorrectionApi.GetSLOCorrection(ctx, sloCorrection.GetId())
	if err != nil {
		t.Fatalf("Error getting SLO correction %s: Response %s: %v", sloCorrection.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrectionGetData := sloCorrectionGetResp.GetData()
	assert.Equal(sloCorrectionGetData.GetId(), sloCorrection.GetId())
	sloCorrectionAttributes := sloCorrectionGetData.GetAttributes()
	assert.Equal(sloCorrectionAttributes.GetTimezone(), testTimezone)
	assert.Equal(sloCorrectionAttributes.GetCategory(), testCategory)

	testSLOCorrectionUpdateData := datadogV1.NewSLOCorrectionUpdateData()
	testCategory = datadogV1.SLOCORRECTIONCATEGORY_OTHER
	testSLOCorrectionUpdateAttributes := datadogV1.SLOCorrectionUpdateRequestAttributes{
		Timezone: &testTimezone,
		Category: &testCategory,
		Start:    &testStart,
		End:      &testEnd,
	}
	testSLOCorrectionUpdateData.SetAttributes(testSLOCorrectionUpdateAttributes)
	testSLOCorrectionUpdate := datadogV1.SLOCorrectionUpdateRequest{
		Data: testSLOCorrectionUpdateData,
	}

	sloCorrectionUpdateResp, httpresp, err := sloCorrectionApi.UpdateSLOCorrection(ctx, sloCorrection.GetId(), testSLOCorrectionUpdate)
	if err != nil {
		t.Fatalf("Error updating SLO correction %v: Response %s: %v", testSLOCorrectionUpdate, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	sloCorrectionUpdateData := sloCorrectionUpdateResp.GetData()
	sloCorrectionAttributes = sloCorrectionUpdateData.GetAttributes()
	assert.Equal(sloCorrectionAttributes.GetCategory(), testCategory)

	httpresp, err = sloCorrectionApi.DeleteSLOCorrection(ctx, sloCorrection.GetId())
	if err != nil {
		t.Fatalf("Error deleting SLO correction %s: Response %s: %v", sloCorrection.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)
}

func deleteSLOCorrectionIfExists(ctx context.Context, t *testing.T, sloCorrectionID string) {
	sloCorrectionApi := datadogV1.NewServiceLevelObjectiveCorrectionsApi(Client(ctx))
	httpresp, err := sloCorrectionApi.DeleteSLOCorrection(ctx, sloCorrectionID)
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting SLO correction: %v failed with %v, Another test may have already deleted this SLO correction: %v",
			sloCorrectionID, httpresp.StatusCode, err)
	}
}

func isSLOCorrectionIDPresent(sloCorrectionID string, sloCorrections []datadogV1.SLOCorrection) error {
	for _, sloCorrection := range sloCorrections {
		if sloCorrection.GetId() == sloCorrectionID {
			return nil
		}
	}
	return fmt.Errorf("SLO correction %s expected but not found", sloCorrectionID)
}
