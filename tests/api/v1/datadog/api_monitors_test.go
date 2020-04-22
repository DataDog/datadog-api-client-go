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

func testMonitor(t *testing.T) datadog.Monitor {
	return datadog.Monitor{
		Name:    datadog.PtrString(fmt.Sprintf("datadog-api-client-go: %s %d", t.Name(), TESTCLOCK.Now().UnixNano())),
		Type:    datadog.MONITORTYPE_LOG_ALERT.Ptr(),
		Query:   datadog.PtrString("logs(\"service:foo AND type:error\").index(\"main\").rollup(\"count\").last(\"5m\") > 2"),
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: &[]string{
			"test",
			"client:go",
		},
		Options: &datadog.MonitorOptions{
			NotifyAudit:       datadog.PtrBool(false),
			Locked:            datadog.PtrBool(false),
			TimeoutH:          *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			RenotifyInterval:  *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			EnableLogsSample:  datadog.PtrBool(true),
			NoDataTimeframe:   *datadog.NewNullableInt64(nil),
			NewHostDelay:      *datadog.NewNullableInt64(datadog.PtrInt64(600)),
			RequireFullWindow: datadog.PtrBool(true),
			NotifyNoData:      datadog.PtrBool(false),
			IncludeTags:       datadog.PtrBool(true),
			EvaluationDelay:   *datadog.NewNullableInt64(datadog.PtrInt64(700)),
			EscalationMessage: datadog.PtrString("the situation has escalated"),
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *datadog.NewNullableFloat64(datadog.PtrFloat64(1)),
			},
		},
	}
}

var testUpdateMonitor = datadog.Monitor{
	Name: datadog.PtrString("updated name"),
	Options: &datadog.MonitorOptions{
		TimeoutH:         *datadog.NewNullableInt64(nil),
		RenotifyInterval: *datadog.NewNullableInt64(nil),
		NewHostDelay:     *datadog.NewNullableInt64(nil),
		EvaluationDelay:  *datadog.NewNullableInt64(nil),
		Thresholds: &datadog.MonitorThresholds{
			Critical: datadog.PtrFloat64(2),
			Warning:  *datadog.NewNullableFloat64(nil),
		},
	},
}

func TestMonitorValidation(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Monitor            datadog.Monitor
		ExpectedStatusCode int
	}{
		{"empty monitor", datadog.Monitor{}, 400},
		{"example monitor", testMonitor(t), 200},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.ValidateMonitor(TESTAUTH).Body(tc.Monitor).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode, "error: %v", err)
		})
	}
}

func TestMonitorLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	tm := testMonitor(t)

	// Create monitor
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, tm.GetName(), monitor.GetName())
	val, set := monitor.GetDeletedOk()
	assert.True(t, set)
	assert.Nil(t, val)

	// Edit a monitor
	updatedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.UpdateMonitor(TESTAUTH, monitor.GetId()).Body(testUpdateMonitor).Execute()
	if err != nil {
		t.Errorf("Error updating Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testUpdateMonitor.GetName(), updatedMonitor.GetName())

	// Assert Explicitly null fields
	var monitorOptions = updatedMonitor.GetOptions()
	var monitorOptionThresholds = monitorOptions.GetThresholds()
	th, ok1 := monitorOptions.GetTimeoutHOk()
	assert.Nil(t, th)
	assert.Equal(t, true, ok1)

	rn, ok2 := monitorOptions.GetRenotifyIntervalOk()
	assert.Nil(t, rn)
	assert.Equal(t, true, ok2)

	nhd, ok3 := monitorOptions.GetNewHostDelayOk()
	assert.Nil(t, nhd)
	assert.Equal(t, true, ok3)

	ed, ok4 := monitorOptions.GetEvaluationDelayOk()
	assert.Nil(t, ed)
	assert.Equal(t, true, ok4)
	// Warning isn't returned in the API response if its unset
	assert.Equal(t, false, monitorOptionThresholds.HasWarning())

	// Check monitor existence
	fetchedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.GetMonitor(TESTAUTH, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedMonitor.GetName(), fetchedMonitor.GetName())

	// Find our monitor in the full list
	monitors, httpresp, err := TESTAPICLIENT.MonitorsApi.ListMonitors(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, monitors, fetchedMonitor)

	// Can delete
	ids := []int64{monitor.GetId()}
	canDeleteResp, httpresp, err := TESTAPICLIENT.MonitorsApi.CheckCanDeleteMonitor(TESTAUTH).MonitorIds(ids).Execute()
	if err != nil {
		t.Errorf("Cannot delete Monitor %v: Response %s: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, ids, canDeleteResp.Data.GetOk())
	assert.Empty(t, canDeleteResp.GetErrors())

	// Delete
	deletedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error deleting Monitor %v: Response %s: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, monitor.GetId(), deletedMonitor.GetDeletedMonitorId())
}

func TestMonitorPagination(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	tm := testMonitor(t)

	// Create monitor
	monitor, _, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())

	monitors, httpresp, err := TESTAPICLIENT.MonitorsApi.ListMonitors(TESTAUTH).Page(0).PageSize(1).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 1, len(monitors))

	monitors, httpresp, err = TESTAPICLIENT.MonitorsApi.ListMonitors(TESTAUTH).IdOffset(monitor.GetId() - 1).PageSize(1).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 1, len(monitors))
	assert.Equal(t, monitor.GetId(), monitors[0].GetId())
}

func TestMonitorsCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.Monitor{}, 400},
		{"403 Forbidden", fake_auth, datadog.Monitor{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestMonitorsListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.ListMonitors(tc.Ctx).GroupStates("notagroupstate").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestMonitorUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	tm := testMonitor(t)
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	updateMonitor := *datadog.NewMonitorWithDefaults()
	updateMonitor.SetType(datadog.MONITORTYPE_COMPOSITE)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, monitor.GetId(), updateMonitor, 400},
		// Cannot trigger 401 for client. Need underrestricted creds.
		// {"401 Unauthorized", TESTAUTH, 1234, datadog.Monitor{}, 401},
		{"403 Forbidden", fake_auth, 1234, datadog.Monitor{}, 403},
		{"404 Not Found", TESTAUTH, 1234, datadog.Monitor{}, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.UpdateMonitor(tc.Ctx, tc.ID).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestMonitorUpdate401Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	// Cannot trigger 401 for client. Need underrestricted creds. Mock it.
	res, err := tests.ReadFixture("fixtures/monitors/error_401.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	gock.New("https://api.datadoghq.com").Put("/api/v1/monitor/121").Reply(401).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.MonitorsApi.UpdateMonitor(TESTAUTH, 121).Body(datadog.Monitor{}).Execute()
	assert.Equal(t, 401, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestMonitorsGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	tm := testMonitor(t)
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, monitor.GetId(), 400},
		{"403 Forbidden", fake_auth, 1234, 403},
		{"404 Not Found", TESTAUTH, 1234, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.GetMonitor(tc.Ctx, tc.ID).GroupStates("notagroupstate").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestMonitorDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 int64
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		// Cannot trigger 400 due to client side validations
		// {"400 Bad Request", TESTAUTH, monitor.GetId(), updateMonitor, 400},
		// Cannot trigger 401 for client. Need underrestricted creds.
		// {"401 Unauthorized", TESTAUTH, 1234, datadog.Monitor{}, 401},
		{"403 Forbidden", fake_auth, 1234, datadog.Monitor{}, 403},
		{"404 Not Found", TESTAUTH, 1234, datadog.Monitor{}, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.UpdateMonitor(tc.Ctx, tc.ID).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestMonitorDelete400Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	// Cannot trigger 400 due to client side validations, so mock it
	res, err := tests.ReadFixture("fixtures/monitors/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	gock.New("https://api.datadoghq.com").Delete("/api/v1/monitor/121").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, 121).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestMonitorDelete401Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	// Cannot trigger 401 for client. Need underrestricted creds. Mock it.
	res, err := tests.ReadFixture("fixtures/monitors/error_401.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	gock.New("https://api.datadoghq.com").Delete("/api/v1/monitor/121").Reply(401).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, 121).Execute()
	assert.Equal(t, 401, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestMonitorCanDeleteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor that can't be deleted
	tm := testMonitor(t)
	monitor := *datadog.NewMonitorWithDefaults()
	monitor.SetType(datadog.MONITORTYPE_QUERY_ALERT)
	monitor.SetQuery("avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100")
	monitor, _, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(monitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	composite := *datadog.NewMonitorWithDefaults()
	composite.SetType(datadog.MONITORTYPE_COMPOSITE)
	composite.SetQuery(fmt.Sprintf("%d", monitor.GetId()))
	composite, _, err = TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(composite).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(composite.GetId())

	testCases := []struct {
		Name               string
		Ctx                context.Context
		IDs                []int64
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, []int64{}, 400},
		{"403 Forbidden", fake_auth, []int64{1234}, 403},
		{"409 Conflict", TESTAUTH, []int64{monitor.GetId()}, 409},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.CheckCanDeleteMonitor(tc.Ctx).MonitorIds(tc.IDs).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 409 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.CheckCanDeleteMonitorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			}
		})
	}
}

func TestMonitorValidateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.Monitor{}, 400},
		{"403 Forbidden", fake_auth, datadog.Monitor{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.ValidateMonitor(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func deleteMonitor(monitorID int64) {
	_, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitorID).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting Monitor: %v failed with %v, Another test may have already deleted this monitor.", monitorID, httpresp.StatusCode)
	}
}
