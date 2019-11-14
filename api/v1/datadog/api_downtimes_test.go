package datadog_test

import (
	"log"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestDowntimeLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now()
	testDowntime := datadog.Downtime{
		Message:  datadog.PtrString("Testing downtime from Go client"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"*"},
	}

	// Create downtime
	downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH, testDowntime)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())

	assert.Equal(t, downtime.GetMessage(), testDowntime.GetMessage())
	assert.Assert(t, downtime.GetActive())

	// Edit a downtime
	editedDowntime := datadog.Downtime{Message: datadog.PtrString("updated message")}
	updatedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.UpdateDowntime(TESTAUTH, downtime.GetId(), editedDowntime)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error updating Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, editedDowntime.GetMessage(), updatedDowntime.GetMessage())

	// Check downtime existence
	fetchedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.GetDowntime(TESTAUTH, downtime.GetId())
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error fetching Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, updatedDowntime.GetMessage(), fetchedDowntime.GetMessage())

	// Find our downtime in the full list
	downtimes, httpresp, err := TESTAPICLIENT.DowntimesApi.GetAllDowntimes(TESTAUTH, nil)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error fetching downtimes; Status: %v: %v", httpresp.StatusCode, err)
	}
	assert.Assert(t, is.Contains(downtimes, fetchedDowntime))

	// Cancel downtime
	httpresp, err = TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtime.GetId())
	if err != nil {
		t.Errorf("Error canceling Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}

	// Check downtime status
	fetchedDowntime, httpresp, err = TESTAPICLIENT.DowntimesApi.GetDowntime(TESTAUTH, downtime.GetId())
	if httpresp.StatusCode != 200 {
		t.Errorf("Downtime %v should still exist: Status: %v: %v", fetchedDowntime.GetId(), httpresp.StatusCode, err)
	}
	assert.Assert(t, !fetchedDowntime.GetActive())
	assert.Assert(t, fetchedDowntime.GetDisabled())
}

func TestMonitorDowntime(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH, testMonitor)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error creating Monitor %v: Status: %v: %v", testMonitor, httpresp.StatusCode, err)
	}
	monitorId := monitor.GetId()
	defer deleteMonitor(monitorId)

	start := time.Now()
	testDowntime := datadog.Downtime{
		Message:  datadog.PtrString("Testing downtime with monitor from Go client"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"*"},
		MonitorId: &datadog.NullableInt64{
			Value: monitorId,
		},
	}

	// Create downtime
	downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH, testDowntime)
	if err != nil || httpresp.StatusCode != 200 {
		t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())

	assert.Equal(t, downtime.GetMonitorId().Value, monitorId)
}

func TestScopedDowntime(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now()
	testDowntimes := []datadog.Downtime{{
		Message:  datadog.PtrString("Testing scope downtime: client, go"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"test:client", "test:go"},
	}, {
		Message:  datadog.PtrString("Testing scope downtime: go"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"test:go"},
	}, {
		Message:  datadog.PtrString("Testing scope downtime: client"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"test:client"},
	},
	}

	// Create downtimes
	downtimes := make([]datadog.Downtime, len(testDowntimes))
	for i, testDowntime := range testDowntimes {
		downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH, testDowntime)
		if err != nil || httpresp.StatusCode != 200 {
			t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		defer cancelDowntime(downtime.GetId())

		assert.Assert(t, !downtime.GetDisabled())
		downtimes[i] = downtime
	}

	// Cancel downtimes with a scope "test:go"
	scope := "test:go"
	cancelDowntimesByScopeRequest := datadog.CancelDowntimesByScopeRequest{
		Scope: scope,
	}
	canceledDowntimesIds, httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntimesByScope(TESTAUTH, cancelDowntimesByScopeRequest)
	if httpresp.StatusCode != 200 || err != nil {
		t.Errorf("Error canceling downtimes by scope %s: %v", scope, err)
	}

	canceledIds := canceledDowntimesIds.GetCancelledIds()
	expectedCanceledIds := []int64{downtimes[1].GetId()}
	assert.DeepEqual(t, canceledIds, expectedCanceledIds)
}

func cancelDowntime(downtimeId int64) {
	httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtimeId)
	if err != nil {
		log.Printf("Canceling Downtime: %v failed with %v, Another test may have already canceled this downtime: %v", downtimeId, httpresp.StatusCode, err)
	}
}
