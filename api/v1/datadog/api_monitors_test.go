package datadog_test

import (
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

var testMonitor = datadog.Monitor{
	Name:    datadog.PtrString("name for Go client monitor foo"),
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
		TimeoutH:          datadog.PtrInt64(60),
		RenotifyInterval:  datadog.PtrInt64(60),
		EnableLogsSample:  datadog.PtrBool(true),
		NoDataTimeframe:   nil,
		NewHostDelay:      datadog.PtrInt64(600),
		RequireFullWindow: datadog.PtrBool(true),
		NotifyNoData:      datadog.PtrBool(false),
		IncludeTags:       datadog.PtrBool(true),
		EvaluationDelay:   datadog.PtrInt64(700),
		EscalationMessage: datadog.PtrString("the situation has escalated"),
		Thresholds: &datadog.MonitorThresholds{
			Critical: datadog.PtrFloat64(2),
			Warning:  datadog.PtrFloat64(1),
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
		{"example monitor", testMonitor, 200},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.ValidateMonitor(TESTAUTH).Body(tc.Monitor).Execute()
			assert.Equal(t, httpresp.StatusCode, tc.ExpectedStatusCode, "error: %v", err)
		})
	}
}

func TestMonitorLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(testMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testMonitor, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, monitor.GetName(), testMonitor.GetName())

	// Edit a monitor
	editedMonitor := datadog.Monitor{Name: datadog.PtrString("updated name")}
	updatedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.EditMonitor(TESTAUTH, monitor.GetId()).Body(editedMonitor).Execute()
	if err != nil {
		t.Errorf("Error updating Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, editedMonitor.GetName(), updatedMonitor.GetName())

	// Check monitor existence
	fetchedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.GetMonitor(TESTAUTH, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Monitor %v: Response %v: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, updatedMonitor.GetName(), fetchedMonitor.GetName())

	// Find our monitor in the full list
	monitors, httpresp, err := TESTAPICLIENT.MonitorsApi.GetAllMonitors(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error fetching monitors: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.Contains(monitors, fetchedMonitor))

	// Delete
	deletedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitor.GetId()).Execute()
	if err != nil {
		t.Errorf("Error deleting Monitor %v: Response %s: %v", monitor.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, monitor.GetId(), deletedMonitor.GetDeletedMonitorId())
}

func deleteMonitor(monitorID int64) {
	_, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitorID).Execute()
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting Monitor: %v failed with %v, Another test may have already deleted this monitor.", monitorID, httpresp.StatusCode)
	}
}
