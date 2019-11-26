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
		Recurrence: &datadog.NullableDowntimeRecurrence{Value: datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilDate: &datadog.NullableInt64{
				Value: start.Unix() + 21*24*60*60, // +21d
			},
		}},
	}

	// Create downtime
	downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH, testDowntime)
	if err != nil {
		t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, downtime.GetMessage(), testDowntime.GetMessage())
	assert.Assert(t, downtime.GetActive())

	// Edit a downtime
	editedDowntime := datadog.Downtime{Message: datadog.PtrString("updated message")}
	updatedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.UpdateDowntime(TESTAUTH, downtime.GetId(), editedDowntime)
	if err != nil {
		t.Errorf("Error updating Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, editedDowntime.GetMessage(), updatedDowntime.GetMessage())

	// Check downtime existence
	fetchedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.GetDowntime(TESTAUTH, downtime.GetId())
	if err != nil {
		t.Errorf("Error fetching Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, updatedDowntime.GetMessage(), fetchedDowntime.GetMessage())

	// Find our downtime in the full list
	downtimes, httpresp, err := TESTAPICLIENT.DowntimesApi.GetAllDowntimes(TESTAUTH, nil)
	if err != nil {
		t.Errorf("Error fetching downtimes; Status: %v: %v", httpresp.StatusCode, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, is.Contains(downtimes, fetchedDowntime))

	// Cancel downtime
	httpresp, err = TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtime.GetId())
	if err != nil {
		t.Errorf("Error canceling Downtime %v: Status: %v: %v", downtime.GetId(), httpresp.StatusCode, err)
	}
	assert.Equal(t, httpresp.StatusCode, 204)

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
	if err != nil {
		t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())
	assert.Equal(t, httpresp.StatusCode, 200)

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
		if err != nil {
			t.Errorf("Error creating Downtime %v: Status: %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		defer cancelDowntime(downtime.GetId())
		assert.Equal(t, httpresp.StatusCode, 200)

		assert.Assert(t, !downtime.GetDisabled())
		downtimes[i] = downtime
	}

	// Cancel downtimes with a scope "test:go"
	scope := "test:go"
	cancelDowntimesByScopeRequest := datadog.CancelDowntimesByScopeRequest{
		Scope: scope,
	}
	canceledDowntimesIds, httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntimesByScope(TESTAUTH, cancelDowntimesByScopeRequest)
	if err != nil {
		t.Errorf("Error canceling downtimes by scope %s: %v", scope, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	canceledIds := canceledDowntimesIds.GetCancelledIds()
	expectedCanceledIds := []int64{downtimes[1].GetId()}
	assert.DeepEqual(t, canceledIds, expectedCanceledIds)
}

func TestDowntimeRecurrence(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := time.Now()
	testCases := []struct {
		Name               string
		DowntimeRecurence  datadog.DowntimeRecurrence
		ExpectedStatusCode int
	}{
		{"once a year", datadog.DowntimeRecurrence{
			Type:   datadog.PtrString("years"),
			Period: datadog.PtrInt32(1),
		}, 200},
		{"invalid type hours", datadog.DowntimeRecurrence{
			Type:   datadog.PtrString("hours"), // Choose from: days, weeks, months, years.
			Period: datadog.PtrInt32(1),
		}, 400},
		{"invalid weekdays", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"mon", "tue"},
		}, 400},
		/* Only applicable when type is weeks -> unclear error code
		{"weekdays only with type weeks not days", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("days"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400},
		{"weekdays only with type weeks not months", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("months"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400},
		{"weekdays only with type weeks not years", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("years"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400}, */
		{"until date", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilDate: &datadog.NullableInt64{
				Value: start.Unix() + 21*24*60*60, // +21d
			}}, 200},
		{"until occurrences", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: &datadog.NullableInt32{
				Value: 3,
			}}, 200},
		{"until occurences and until date are mutually exclusive", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: &datadog.NullableInt32{
				Value: 3,
			},
			UntilDate: &datadog.NullableInt64{
				Value: start.Unix() + 21*24*60*60, // +21d
			}}, 400},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testDowntime := datadog.Downtime{
				Message:    datadog.PtrString(tc.Name),
				Start:      datadog.PtrInt64(start.Unix()),
				Timezone:   datadog.PtrString("Etc/UTC"),
				Scope:      &[]string{"*"},
				Recurrence: &datadog.NullableDowntimeRecurrence{Value: tc.DowntimeRecurence},
			}

			downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH, testDowntime)
			defer cancelDowntime(downtime.GetId())

			assert.Equal(t, httpresp.StatusCode, tc.ExpectedStatusCode, "error: %v", err)
		})
	}
}

func cancelDowntime(downtimeId int64) {
	httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtimeId)
	if err != nil {
		log.Printf("Canceling Downtime: %v failed with %v, Another test may have already canceled this downtime: %v", downtimeId, httpresp.StatusCode, err)
	}
}
