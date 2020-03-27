/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var testMonitorSLO = datadog.ServiceLevelObjective{
	Type:        "monitor",
	Name:        "Critical Foo Host Uptime",
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
	Name:        "HTTP Return Codes",
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
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor to reference in the SLO
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(testMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testMonitor, err.Error(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create SLO
	sloResp, httpresp, err := TESTAPICLIENT.SLOApi.CreateSLO(TESTAUTH).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.Error(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(slo.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, slo.GetName(), testMonitorSLO.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = TESTAPICLIENT.SLOApi.EditSLO(TESTAUTH, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	slo2 := sloResp.GetData()[0]
	assert.Equal(t, slo2.GetDescription(), "Updated description")

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteServiceLevelObjectiveResponse
	canDeleteResp, httpresp, err = TESTAPICLIENT.SLOApi.CheckCanDeleteSLO(TESTAUTH).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	canDelete := canDeleteResp.GetData()
	assert.DeepEqual(t, canDelete.GetOk(), []string{slo2.GetId()})

	// Get SLO
	var sloGetResp datadog.ServiceLevelObjectiveResponse
	sloGetResp, httpresp, err = TESTAPICLIENT.SLOApi.GetSLO(TESTAUTH, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	slo3 := sloGetResp.GetData()
	assert.Equal(t, slo2.GetName(), slo3.GetName())

	// Get SLO history
	now := TESTCLOCK.Now().Unix()
	_, httpresp, err = TESTAPICLIENT.SLOApi.HistoryForSLO(TESTAUTH, slo3.GetId()).
		FromTs(fmt.Sprintf("%d", now-11)).ToTs(fmt.Sprintf("%d", now-1)).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// Delete SLO
	var sloDeleteResp datadog.ServiceLevelObjectiveDeleted
	sloDeleteResp, httpresp, err = TESTAPICLIENT.SLOApi.DeleteSLO(TESTAUTH, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.DeepEqual(t, sloDeleteResp.GetData(), []string{slo3.GetId()})
}

func TestSLOEventLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create SLO
	sloResp, httpresp, err := TESTAPICLIENT.SLOApi.CreateSLO(TESTAUTH).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.Error(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(slo.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, slo.GetName(), testEventSLO.GetName())

	// Edit SLO
	slo.SetDescription("Updated description")
	sloResp, httpresp, err = TESTAPICLIENT.SLOApi.EditSLO(TESTAUTH, slo.GetId()).Body(slo).Execute()
	if err != nil {
		t.Fatalf("Error updating SLO %v: Response %s: %v", slo, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	slo2 := sloResp.GetData()[0]
	assert.Equal(t, slo2.GetDescription(), "Updated description")

	// Check that the SLO can be deleted
	var canDeleteResp datadog.CheckCanDeleteServiceLevelObjectiveResponse
	canDeleteResp, httpresp, err = TESTAPICLIENT.SLOApi.CheckCanDeleteSLO(TESTAUTH).Ids(slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error checking that SLO %s can be deleted: Response %s: %v", slo.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	canDelete := canDeleteResp.GetData()
	assert.DeepEqual(t, canDelete.GetOk(), []string{slo2.GetId()})

	// Get SLO
	var sloGetResp datadog.ServiceLevelObjectiveResponse
	sloGetResp, httpresp, err = TESTAPICLIENT.SLOApi.GetSLO(TESTAUTH, slo2.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting SLO %s: Response %s: %v", slo2.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	slo3 := sloGetResp.GetData()
	assert.Equal(t, slo2.GetName(), slo3.GetName())

	// Get SLO history
	now := TESTCLOCK.Now().Unix()
	_, httpresp, err = TESTAPICLIENT.SLOApi.HistoryForSLO(TESTAUTH, slo3.GetId()).
		FromTs(fmt.Sprintf("%d", now-11)).ToTs(fmt.Sprintf("%d", now-1)).Execute()
	// the contents of history really depend on the org that this test is running in, so we just ensure
	// that the structure deserialized properly and no error was returned
	if err != nil {
		t.Fatalf("Error getting history for SLO %s: Response %s: %v", slo3.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// Delete SLO
	var sloDeleteResp datadog.ServiceLevelObjectiveDeleted
	sloDeleteResp, httpresp, err = TESTAPICLIENT.SLOApi.DeleteSLO(TESTAUTH, slo3.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting SLO %s: Response %s: %v", slo3.GetId(), err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.DeepEqual(t, sloDeleteResp.GetData(), []string{slo3.GetId()})
}

func TestSLOMultipleInstances(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor to reference in the monitor SLO
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(testMonitor).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", testMonitor, err.Error(), err)
	}
	defer deleteMonitor(monitor.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	testMonitorSLO.MonitorIds = &[]int64{*monitor.Id}

	// Create monitor SLO
	sloResp, httpresp, err := TESTAPICLIENT.SLOApi.CreateSLO(TESTAUTH).Body(testMonitorSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.Error(), err)
	}
	monitorSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(monitorSLO.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	// Create event SLO
	sloResp, httpresp, err = TESTAPICLIENT.SLOApi.CreateSLO(TESTAUTH).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v: Response %s: %v", testMonitorSLO, err.Error(), err)
	}
	eventSLO := sloResp.GetData()[0]
	defer deleteSLOIfExists(eventSLO.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	// Get multiple SLOs
	var slosResp datadog.ServiceLevelObjectiveListResponse
	slosResp, httpresp, err = TESTAPICLIENT.SLOApi.GetSLOs(TESTAUTH).
		Ids(fmt.Sprintf("%s,%s", monitorSLO.GetId(), eventSLO.GetId())).Execute()
	if err != nil {
		t.Fatalf("Error getting SLOs: Response %s: %v", err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	slos := slosResp.GetData()
	assertSLOIDPresent(t, monitorSLO.GetId(), slos)
	assertSLOIDPresent(t, eventSLO.GetId(), slos)

	// Use bulk delete operation to delete the event SLO
	var deleteResp datadog.ServiceLevelObjectivesBulkDeleted
	deleteResp, httpresp, err = TESTAPICLIENT.SLOApi.BulkPartialDeleteSLO(TESTAUTH).
		Body(map[string][]datadog.SLOTimeframe{
			eventSLO.GetId(): {datadog.SLOTIMEFRAME_SEVEN_DAYS},
		}).Execute()

	if err != nil {
		t.Fatalf("Error bulk deleting SLOs: Response %s: %v", err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	deleteData := deleteResp.GetData()
	assert.DeepEqual(t, deleteData.GetDeleted(), []string{eventSLO.GetId()})
	assert.Equal(t, len(deleteData.GetUpdated()), 0)
}

func assertSLOIDPresent(t *testing.T, sloID string, slos []datadog.ServiceLevelObjective) {
	for _, slo := range slos {
		if slo.GetId() == sloID {
			return
		}
	}
	assert.NilError(t, fmt.Errorf("SLO %s expected but not found", sloID))
}

func deleteSLOIfExists(sloID string) {
	_, httpresp, err := TESTAPICLIENT.SLOApi.DeleteSLO(TESTAUTH, sloID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting SLO: %v failed with %v, Another test may have already deleted this SLO: %v",
			sloID, httpresp.StatusCode, err)
	}
}
