package datadog_v1_test

import (
	"log"
	"testing"

	datadog "github.com/DataDog/datadog-api-client-go/datadog_v1"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

var testMonitor = datadog.Monitor{
	Name:    datadog.PtrString("name for Go client monitor foo"),
	Type:    datadog.PtrString("log alert"),
	Query:   datadog.PtrString("logs(\"service:foo AND type:error\").index(\"main\").rollup(\"count\").last(\"5m\") > 2"),
	Message: datadog.PtrString("some message Notify: @hipchat-channel"),
	Tags: &[]string{
		"test",
		"client:go",
	},
	Options: &datadog.MonitorOptions{
		NotifyAudit:       datadog.PtrBool(false),
		Locked:            datadog.PtrBool(false),
		TimeoutH:          datadog.PtrInt32(60),
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
			_, httpresp, err := TESTAPICLIENT.MonitorsApi.ValidateMonitor(TESTAUTH, tc.Monitor)
			assert.Equal(t, httpresp.StatusCode, tc.ExpectedStatusCode, "error: %v", err)
		})
	}
}

func TestMonitorLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH, testMonitor)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error creating Monitor %v: Status: %v: %v", testMonitor, httpresp.StatusCode, err)
	}
	defer deleteMonitor(monitor.GetId())

	assert.Equal(t, monitor.GetName(), testMonitor.GetName())

	// Edit a monitor
	editedMonitor := datadog.Monitor{Name: datadog.PtrString("updated name")}
	updatedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.EditMonitor(TESTAUTH, monitor.GetId(), editedMonitor)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error updating Monitor %v: Status: %v: %v", monitor.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, editedMonitor.GetName(), updatedMonitor.GetName())

	// Check monitor existence
	fetchedMonitor, httpresp, err := TESTAPICLIENT.MonitorsApi.GetMonitor(TESTAUTH, monitor.GetId(), nil)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error fetching Monitor %v: Status: %v: %v", monitor.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, updatedMonitor.GetName(), fetchedMonitor.GetName())

	// Find our monitor in the full list
	monitors, httpresp, err := TESTAPICLIENT.MonitorsApi.GetAllMonitors(TESTAUTH, nil)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error fetching monitors; Status: %v: %v", httpresp.StatusCode, err)
	}
	assert.Assert(t, is.Contains(monitors, fetchedMonitor))

	// Delete
	_, httpresp, err = TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitor.GetId())
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error deleting Monitor %v: Status: %v: %v", monitor.GetId(), httpresp.StatusCode, err)
	}

	// Check monitor non existence
	fetchedMonitor, httpresp, err = TESTAPICLIENT.MonitorsApi.GetMonitor(TESTAUTH, monitor.GetId(), nil)
	if httpresp.StatusCode != 404 {
		t.Errorf("Monitor %v should be deleted: Status: %v: %v", fetchedMonitor.GetId(), httpresp.StatusCode, err)
	}
}

func deleteMonitor(monitorId int64) {
	_, httpresp, err := TESTAPICLIENT.MonitorsApi.DeleteMonitor(TESTAUTH, monitorId)
	if httpresp.StatusCode != 200 || err != nil {
		log.Printf("Deleting Monitor: %v failed with %v, Another test may have already deleted this monitor.", monitorId, httpresp.StatusCode)
	}
}
