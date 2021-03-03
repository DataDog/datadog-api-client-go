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

	"gopkg.in/h2non/gock.v1"
)

func testMonitor(ctx context.Context, t *testing.T) datadog.Monitor {
	return datadog.Monitor{
		Name:    tests.UniqueEntityName(ctx, t),
		Type:    datadog.MONITORTYPE_LOG_ALERT.Ptr(),
		Query:   datadog.PtrString("logs(\"service:foo AND type:error\").index(\"main\").rollup(\"count\").by(\"source\").last(\"5m\") > 2"),
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: &[]string{
			"test",
			"client:go",
		},
		Priority: datadog.PtrInt64(3),
		Options: &datadog.MonitorOptions{
			NotifyAudit:          datadog.PtrBool(false),
			Locked:               datadog.PtrBool(false),
			TimeoutH:             *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			RenotifyInterval:     *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			EnableLogsSample:     datadog.PtrBool(true),
			GroupbySimpleMonitor: datadog.PtrBool(true),
			NoDataTimeframe:      *datadog.NewNullableInt64(nil),
			NewHostDelay:         *datadog.NewNullableInt64(datadog.PtrInt64(600)),
			RequireFullWindow:    datadog.PtrBool(true),
			NotifyNoData:         datadog.PtrBool(false),
			IncludeTags:          datadog.PtrBool(true),
			EvaluationDelay:      *datadog.NewNullableInt64(datadog.PtrInt64(700)),
			EscalationMessage:    datadog.PtrString("the situation has escalated"),
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *datadog.NewNullableFloat64(datadog.PtrFloat64(1)),
			},
		},
		RestrictedRoles: &[]string{
			"94172442-be03-11e9-a77a-3b7612558ac1",
		},
	}
}

var testUpdateMonitor = datadog.MonitorUpdateRequest{
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
	RestrictedRoles: &[]string{
		"94172442-be03-11e9-a77a-3b7612558ac1",
	},
}

func TestMonitorValidation(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	testCases := map[string]struct {
		Monitor            datadog.Monitor
		ExpectedStatusCode int
	}{
		"empty monitor":   {datadog.Monitor{}, 400},
		"example monitor": {testMonitor(ctx, t), 200},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(ctx, t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.ValidateMonitor(ctx).Body(tc.Monitor).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode, "error: %v", err)
		})
	}
}

func TestMonitorLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	tm := testMonitor(ctx, t)

	// Create monitor
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(tm.GetName(), monitor.GetName())
	val, set := monitor.GetDeletedOk()
	assert.True(set)
	assert.Nil(val)

	// Edit a monitor
	testUpdateMonitor.SetName(fmt.Sprintf("%s-updated", tm.GetName()))
	updatedMonitor, httpresp, err := Client(ctx).MonitorsApi.UpdateMonitor(ctx, monitor.GetId()).Body(testUpdateMonitor).Execute()
	if err != nil {
		t.Errorf("Error updating Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testUpdateMonitor.GetName(), updatedMonitor.GetName())

	// Assert Explicitly null fields
	var monitorOptions = updatedMonitor.GetOptions()
	var monitorOptionThresholds = monitorOptions.GetThresholds()
	th, ok1 := monitorOptions.GetTimeoutHOk()
	assert.Nil(th)
	assert.Equal(true, ok1)

	rn, ok2 := monitorOptions.GetRenotifyIntervalOk()
	assert.Nil(rn)
	assert.Equal(true, ok2)

	nhd, ok3 := monitorOptions.GetNewHostDelayOk()
	assert.Nil(nhd)
	assert.Equal(true, ok3)

	ed, ok4 := monitorOptions.GetEvaluationDelayOk()
	assert.Nil(ed)
	assert.Equal(true, ok4)
	// Warning isn't returned in the API response if its unset
	assert.Equal(false, monitorOptionThresholds.HasWarning())

	// Check monitor existence
	fetchedMonitor, httpresp, err := Client(ctx).MonitorsApi.GetMonitor(ctx, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedMonitor.GetName(), fetchedMonitor.GetName())

	// Find our monitor in the full list
	monitors, httpresp, err := Client(ctx).MonitorsApi.ListMonitors(ctx).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(monitors, fetchedMonitor)

	// Can delete
	ids := []int64{monitor.GetId()}
	canDeleteResp, httpresp, err := Client(ctx).MonitorsApi.CheckCanDeleteMonitor(ctx).MonitorIds(ids).Execute()
	if err != nil {
		t.Errorf("Cannot delete Monitor %v: Response %s: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(ids, canDeleteResp.Data.GetOk())
	assert.Empty(canDeleteResp.GetErrors())

	// Delete
	deletedMonitor, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error deleting Monitor %v: Response %s: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(monitor.GetId(), deletedMonitor.GetDeletedMonitorId())
}

func TestMonitorPagination(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	tm := testMonitor(ctx, t)

	// Create monitor
	monitor, _, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())

	monitors, httpresp, err := Client(ctx).MonitorsApi.ListMonitors(ctx).Page(0).PageSize(1).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(1, len(monitors))

	monitors, httpresp, err = Client(ctx).MonitorsApi.ListMonitors(ctx).IdOffset(monitor.GetId() - 1).PageSize(1).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(1, len(monitors))
	assert.Equal(monitor.GetId(), monitors[0].GetId())
}

func TestMonitorSyntheticsGet(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create a synthetics API test to retrieve its corresponding monitor via GetMonitor (we're borrowing these from api_synthetics_test.go)
	testSyntheticsAPI := getLegacyTestSyntheticsAPI(ctx, t)
	syntheticsTest, _, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", syntheticsTest, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteSyntheticsTestIfExists(ctx, t, syntheticsTest.GetPublicId())

	// Retrieve the corresponding synthetics Test monitor
	syntheticsMonitor, httpresp, err := Client(ctx).MonitorsApi.GetMonitor(ctx, syntheticsTest.GetMonitorId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Monitor %v: Response %v: %v", syntheticsTest.GetMonitorId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	syntheticsMonitorOptions := syntheticsMonitor.GetOptions()
	assert.Equal(syntheticsTest.GetPublicId(), syntheticsMonitorOptions.GetSyntheticsCheckId())
}

func TestMonitorsCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.Monitor{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Monitor{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMonitorsListErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).MonitorsApi.ListMonitors(ctx).GroupStates("notagroupstate").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMonitorUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create monitor
	tm := testMonitor(ctx, t)
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	updateMonitor := *datadog.NewMonitorUpdateRequestWithDefaults()
	updateMonitor.SetType(datadog.MONITORTYPE_COMPOSITE)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		Body               datadog.MonitorUpdateRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, monitor.GetId(), updateMonitor, 400},
		// Cannot trigger 401 for client. Need underrestricted creds.
		// "401 Unauthorized": {WithTestAuth,1234, datadog.Monitor{}, 401},
		"403 Forbidden": {WithFakeAuth, 1234, datadog.MonitorUpdateRequest{}, 403},
		"404 Not Found": {WithTestAuth, 1234, datadog.MonitorUpdateRequest{}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.UpdateMonitor(ctx, tc.ID).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMonitorUpdate401Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// Cannot trigger 401 for client. Need underrestricted creds. Mock it.
	res, err := tests.ReadFixture("fixtures/monitors/error_401.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Put("/api/v1/monitor/121").Reply(401).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MonitorsApi.UpdateMonitor(ctx, 121).Body(datadog.MonitorUpdateRequest{}).Execute()
	assert.Equal(401, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMonitorsGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create monitor
	tm := testMonitor(ctx, t)
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	assert.Equal(200, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, monitor.GetId(), 400},
		"403 Forbidden":   {WithFakeAuth, 1234, 403},
		"404 Not Found":   {WithTestAuth, 1234, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.GetMonitor(ctx, tc.ID).GroupStates("notagroupstate").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMonitorDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 int64
		ExpectedStatusCode int
	}{
		// Cannot trigger 400 due to client side validations
		// "400 Bad Request": {WithTestAuth,monitor.GetId(), updateMonitor, 400},
		// Cannot trigger 401 for client. Need underrestricted creds.
		// "401 Unauthorized": {WithTestAuth,1234, datadog.Monitor{}, 401},
		"403 Forbidden": {WithFakeAuth, 1234, 403},
		"404 Not Found": {WithTestAuth, 1234, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, tc.ID).Execute()
			assert.Error(err)
			assert.NotNil(httpresp, err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMonitorDelete400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// Cannot trigger 400 due to client side validations, so mock it
	res, err := tests.ReadFixture("fixtures/monitors/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/monitor/121").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, 121).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMonitorDelete401Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// Cannot trigger 401 for client. Need underrestricted creds. Mock it.
	res, err := tests.ReadFixture("fixtures/monitors/error_401.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Delete("/api/v1/monitor/121").Reply(401).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, 121).Execute()
	assert.Equal(401, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMonitorCanDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	// Create monitor that can't be deleted
	tm := testMonitor(ctx, t)
	monitor := *datadog.NewMonitorWithDefaults()
	monitor.SetType(datadog.MONITORTYPE_QUERY_ALERT)
	monitor.SetQuery("avg(last_5m):sum:system.net.bytes_rcvd{host:host0} > 100")
	monitor, _, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(monitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, monitor.GetId())
	composite := *datadog.NewMonitorWithDefaults()
	composite.SetType(datadog.MONITORTYPE_COMPOSITE)
	composite.SetQuery(fmt.Sprintf("%d", monitor.GetId()))
	composite, _, err = Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(composite).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(ctx, t, composite.GetId())

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		IDs                []int64
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, []int64{}, 400},
		"403 Forbidden":   {WithFakeAuth, []int64{1234}, 403},
		"409 Conflict":    {WithTestAuth, []int64{monitor.GetId()}, 409},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.CheckCanDeleteMonitor(ctx).MonitorIds(tc.IDs).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 409 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.CheckCanDeleteMonitorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			}
		})
	}
}

func TestMonitorValidateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Monitor
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.Monitor{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Monitor{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MonitorsApi.ValidateMonitor(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteMonitor(ctx context.Context, t *testing.T, monitorID int64) {
	_, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, monitorID).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Deleting Monitor: %v failed with %v, Another test may have already deleted this monitor.", monitorID, httpresp.StatusCode)
	}
}
