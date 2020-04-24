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
	"github.com/stretchr/testify/assert"
)

func TestDowntimeLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := TESTCLOCK.Now()
	testDowntime := datadog.Downtime{
		Message:  datadog.PtrString("Testing downtime from Go client"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{"*"},
		Recurrence: *datadog.NewNullableDowntimeRecurrence(
			&datadog.DowntimeRecurrence{
				Type:      datadog.PtrString("weeks"),
				Period:    datadog.PtrInt32(1),
				WeekDays:  &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				UntilDate: *datadog.NewNullableInt64(datadog.PtrInt64(start.Unix() + 21*24*60*60)), // +21d
			}),
	}

	// Create downtime
	downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH).Body(testDowntime).Execute()
	if err != nil {
		t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	assert.Equal(t, testDowntime.GetMessage(), downtime.GetMessage())
	assert.True(t, downtime.GetActive())
	val, set := downtime.GetUpdaterIdOk()
	assert.True(t, set)
	assert.Nil(t, val)

	// Edit a downtime
	editedDowntime := datadog.Downtime{Message: datadog.PtrString("updated message")}
	updatedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.UpdateDowntime(TESTAUTH, downtime.GetId()).Body(editedDowntime).Execute()
	if err != nil {
		t.Errorf("Error updating Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, editedDowntime.GetMessage(), updatedDowntime.GetMessage())
	val, set = updatedDowntime.GetUpdaterIdOk()
	assert.True(t, set)
	assert.NotNil(t, val)

	// Check downtime existence
	fetchedDowntime, httpresp, err := TESTAPICLIENT.DowntimesApi.GetDowntime(TESTAUTH, downtime.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedDowntime.GetMessage(), fetchedDowntime.GetMessage())

	// Find our downtime in the full list
	downtimes, httpresp, err := TESTAPICLIENT.DowntimesApi.ListDowntimes(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error fetching downtimes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Contains(t, downtimes, fetchedDowntime)

	// Cancel downtime
	httpresp, err = TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtime.GetId()).Execute()
	if err != nil {
		t.Errorf("Error canceling Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)

	// Check downtime status
	fetchedDowntime, httpresp, err = TESTAPICLIENT.DowntimesApi.GetDowntime(TESTAUTH, downtime.GetId()).Execute()
	if httpresp.StatusCode != 200 {
		t.Errorf("Downtime %v should still exist: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.False(t, fetchedDowntime.GetActive())
	assert.True(t, fetchedDowntime.GetDisabled())
}

func TestMonitorDowntime(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create monitor
	tm := testMonitor(t)
	monitor, httpresp, err := TESTAPICLIENT.MonitorsApi.CreateMonitor(TESTAUTH).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	monitorID := monitor.GetId()
	defer deleteMonitor(monitorID)

	start := TESTCLOCK.Now()
	testDowntime := datadog.Downtime{
		Message:   datadog.PtrString("Testing downtime with monitor from Go client"),
		Start:     datadog.PtrInt64(start.Unix()),
		Timezone:  datadog.PtrString("Etc/UTC"),
		Scope:     &[]string{"*"},
		MonitorId: *datadog.NewNullableInt64(datadog.PtrInt64(monitorID)),
	}

	// Create downtime
	downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH).Body(testDowntime).Execute()
	if err != nil {
		t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(downtime.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	assert.Equal(t, monitorID, downtime.GetMonitorId())
}

func TestScopedDowntime(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := TESTCLOCK.Now()

	scopeClient := fmt.Sprintf("test:client-%s-%d", t.Name(), TESTCLOCK.Now().UnixNano())
	scopeGo := fmt.Sprintf("test:go-%s-%d", t.Name(), TESTCLOCK.Now().UnixNano())

	testDowntimes := []datadog.Downtime{{
		Message:  datadog.PtrString("Testing scope downtime: client, go"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeClient, scopeGo},
	}, {
		Message:  datadog.PtrString("Testing scope downtime: go"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeGo},
	}, {
		Message:  datadog.PtrString("Testing scope downtime: client"),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeClient},
	},
	}

	// Create downtimes
	downtimes := make([]datadog.Downtime, len(testDowntimes))
	for i, testDowntime := range testDowntimes {
		downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH).Body(testDowntime).Execute()
		if err != nil {
			t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		defer cancelDowntime(downtime.GetId())
		assert.Equal(t, 200, httpresp.StatusCode)

		assert.False(t, downtime.GetDisabled())
		downtimes[i] = downtime
	}

	// Cancel downtimes with the scopeGo
	cancelDowntimesByScopeRequest := datadog.CancelDowntimesByScopeRequest{
		Scope: scopeGo,
	}
	canceledDowntimesIds, httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntimesByScope(TESTAUTH).Body(cancelDowntimesByScopeRequest).Execute()
	if err != nil {
		t.Fatalf("Error canceling downtimes by scope %s: Response: %s: %v", scopeGo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	canceledIds := canceledDowntimesIds.GetCancelledIds()
	expectedCanceledIds := []int64{downtimes[1].GetId()}
	assert.Equal(t, expectedCanceledIds, canceledIds)
}

func TestDowntimeRecurrence(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	start := TESTCLOCK.Now()
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
			UntilDate: *datadog.NewNullableInt64(
				datadog.PtrInt64(start.Unix() + 21*24*60*60), // +21d
			)}, 200},
		{"until occurrences", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: *datadog.NewNullableInt32(
				datadog.PtrInt32(3),
			)}, 200},
		{"until occurences and until date are mutually exclusive", datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: *datadog.NewNullableInt32(
				datadog.PtrInt32(3),
			),
			UntilDate: *datadog.NewNullableInt64(
				datadog.PtrInt64(start.Unix() + 21*24*60*60), // +21d
			)}, 400},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testDowntime := datadog.Downtime{
				Message:    datadog.PtrString(tc.Name),
				Start:      datadog.PtrInt64(start.Unix()),
				Timezone:   datadog.PtrString("Etc/UTC"),
				Scope:      &[]string{"*"},
				Recurrence: *datadog.NewNullableDowntimeRecurrence(&tc.DowntimeRecurence),
			}

			downtime, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(TESTAUTH).Body(testDowntime).Execute()
			if tc.ExpectedStatusCode < 300 {
				defer cancelDowntime(downtime.GetId())
			}

			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode, "error: %v", err)
		})
	}
}

func TestDowntimeListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.DowntimesApi.ListDowntimes(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDowntimeCreateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.Downtime
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.Downtime{}, 400},
		{"403 Forbidden", fake_auth, datadog.Downtime{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.DowntimesApi.CreateDowntime(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelByScopeErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.CancelDowntimesByScopeRequest
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.CancelDowntimesByScopeRequest{}, 400},
		{"403 Forbidden", fake_auth, datadog.CancelDowntimesByScopeRequest{}, 403},
		{"404 Not Found", TESTAUTH, datadog.CancelDowntimesByScopeRequest{Scope: "nonexistent"}, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntimesByScope(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", TESTAUTH, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntime(tc.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDowntimeGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		{"404 Not Found", TESTAUTH, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.DowntimesApi.GetDowntime(tc.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestDowntimeUpdateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Endpoint will 400 if there are too many tags
	badDowntime := *datadog.NewDowntimeWithDefaults()
	tags := make([]string, 50)
	for i := 0; i < 50; i++ {
		tags[i] = fmt.Sprintf("tag%d", i)
	}
	badDowntime.MonitorTags = &tags

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.Downtime
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, badDowntime, 400},
		{"403 Forbidden", fake_auth, datadog.Downtime{}, 403},
		{"404 Not Found", TESTAUTH, datadog.Downtime{}, 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.DowntimesApi.UpdateDowntime(tc.Ctx, 1234).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func cancelDowntime(downtimeID int64) {
	httpresp, err := TESTAPICLIENT.DowntimesApi.CancelDowntime(TESTAUTH, downtimeID).Execute()
	if err != nil {
		log.Printf("Canceling Downtime: %v failed with %v, Another test may have already canceled this downtime: %v", downtimeID, httpresp.StatusCode, err)
	}
}
