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
)

func TestDowntimeLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	start := tests.ClockFromContext(ctx).Now()
	testDowntime := datadog.Downtime{
		Message:  tests.UniqueEntityName(ctx, t),
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
	downtime, httpresp, err := Client(ctx).DowntimesApi.CreateDowntime(ctx).Body(testDowntime).Execute()
	if err != nil {
		t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(ctx, t, downtime.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(testDowntime.GetMessage(), downtime.GetMessage())
	assert.True(downtime.GetActive())
	val, set := downtime.GetUpdaterIdOk()
	assert.True(set)
	assert.Nil(val)

	// Edit a downtime
	editedDowntime := datadog.Downtime{Message: datadog.PtrString(fmt.Sprintf("%s-updated", testDowntime.GetMessage()))}
	updatedDowntime, httpresp, err := Client(ctx).DowntimesApi.UpdateDowntime(ctx, downtime.GetId()).Body(editedDowntime).Execute()
	if err != nil {
		t.Errorf("Error updating Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDowntime.GetMessage(), updatedDowntime.GetMessage())
	val, set = updatedDowntime.GetUpdaterIdOk()
	assert.True(set)
	assert.NotNil(val)

	// Check downtime existence
	fetchedDowntime, httpresp, err := Client(ctx).DowntimesApi.GetDowntime(ctx, downtime.GetId()).Execute()
	if err != nil {
		t.Errorf("Error fetching Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedDowntime.GetMessage(), fetchedDowntime.GetMessage())

	// Find our downtime in the full list
	downtimes, httpresp, err := Client(ctx).DowntimesApi.ListDowntimes(ctx).Execute()
	if err != nil {
		t.Errorf("Error fetching downtimes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(downtimes, fetchedDowntime)

	// Cancel downtime
	httpresp, err = Client(ctx).DowntimesApi.CancelDowntime(ctx, downtime.GetId()).Execute()
	if err != nil {
		t.Errorf("Error canceling Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)

	// Check downtime status
	fetchedDowntime, httpresp, err = Client(ctx).DowntimesApi.GetDowntime(ctx, downtime.GetId()).Execute()
	if httpresp.StatusCode != 200 {
		t.Errorf("Downtime %v should still exist: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.False(fetchedDowntime.GetActive())
	assert.True(fetchedDowntime.GetDisabled())
}

func TestMonitorDowntime(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create monitor
	tm := testMonitor(ctx, t)
	monitor, httpresp, err := Client(ctx).MonitorsApi.CreateMonitor(ctx).Body(tm).Execute()
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	monitorID := monitor.GetId()
	defer deleteMonitor(ctx, t, monitorID)

	start := tests.ClockFromContext(ctx).Now()
	testDowntime := datadog.Downtime{
		Message:   tests.UniqueEntityName(ctx, t),
		Start:     datadog.PtrInt64(start.Unix()),
		Timezone:  datadog.PtrString("Etc/UTC"),
		Scope:     &[]string{"*"},
		MonitorId: *datadog.NewNullableInt64(datadog.PtrInt64(monitorID)),
	}

	// Create downtime
	downtime, httpresp, err := Client(ctx).DowntimesApi.CreateDowntime(ctx).Body(testDowntime).Execute()
	if err != nil {
		t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(ctx, t, downtime.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(monitorID, downtime.GetMonitorId())
}

func TestScopedDowntime(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	start := tests.ClockFromContext(ctx).Now()

	scopeClient := fmt.Sprintf("test:client-%s-%d", t.Name(), tests.ClockFromContext(ctx).Now().UnixNano())
	scopeGo := fmt.Sprintf("test:go-%s-%d", t.Name(), tests.ClockFromContext(ctx).Now().UnixNano())

	testDowntimes := []datadog.Downtime{{
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeClient, scopeGo},
	}, {
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeGo},
	}, {
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    &[]string{scopeClient},
	},
	}

	// Create downtimes
	downtimes := make([]datadog.Downtime, len(testDowntimes))
	for i, testDowntime := range testDowntimes {
		downtime, httpresp, err := Client(ctx).DowntimesApi.CreateDowntime(ctx).Body(testDowntime).Execute()
		if err != nil {
			t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		defer cancelDowntime(ctx, t, downtime.GetId())
		assert.Equal(200, httpresp.StatusCode)

		assert.False(downtime.GetDisabled())
		downtimes[i] = downtime
	}

	// Cancel downtimes with the scopeGo
	cancelDowntimesByScopeRequest := datadog.CancelDowntimesByScopeRequest{
		Scope: scopeGo,
	}
	canceledDowntimesIds, httpresp, err := Client(ctx).DowntimesApi.CancelDowntimesByScope(ctx).Body(cancelDowntimesByScopeRequest).Execute()
	if err != nil {
		t.Fatalf("Error canceling downtimes by scope %s: Response: %s: %v", scopeGo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	canceledIds := canceledDowntimesIds.GetCancelledIds()
	expectedCanceledIds := []int64{downtimes[1].GetId()}
	assert.Equal(expectedCanceledIds, canceledIds)
}

func TestDowntimeRecurrence(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	start := tests.ClockFromContext(ctx).Now()
	testCases := map[string]struct {
		DowntimeRecurence  datadog.DowntimeRecurrence
		ExpectedStatusCode int
	}{
		"once a year": {datadog.DowntimeRecurrence{
			Type:   datadog.PtrString("years"),
			Period: datadog.PtrInt32(1),
		}, 200},
		"invalid type hours": {datadog.DowntimeRecurrence{
			Type:   datadog.PtrString("hours"), // Choose from: days, weeks, months, years.
			Period: datadog.PtrInt32(1),
		}, 400},
		"invalid weekdays": {datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"mon", "tue"},
		}, 400},
		/* Only applicable when type is weeks -> unclear error code
		"weekdays only with type weeks not days": { datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("days"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400},
		"weekdays only with type weeks not months": { datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("months"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400},
		"weekdays only with type weeks not years": { datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("years"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon"},
		}, 400}, */
		"until date": {datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilDate: *datadog.NewNullableInt64(
				datadog.PtrInt64(start.Unix() + 21*24*60*60), // +21d
			)}, 200},
		"until occurrences": {datadog.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: &[]string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: *datadog.NewNullableInt32(
				datadog.PtrInt32(3),
			)}, 200},
		"until occurences and until date are mutually exclusive": {datadog.DowntimeRecurrence{
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

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(ctx, t)
			defer finish()
			assert := tests.Assert(ctx, t)

			testDowntime := datadog.Downtime{
				Message:    datadog.PtrString(fmt.Sprintf("%s; %s", *tests.UniqueEntityName(ctx, t), name)),
				Start:      datadog.PtrInt64(start.Unix()),
				Timezone:   datadog.PtrString("Etc/UTC"),
				Scope:      &[]string{"*"},
				Recurrence: *datadog.NewNullableDowntimeRecurrence(&tc.DowntimeRecurence),
			}

			downtime, httpresp, err := Client(ctx).DowntimesApi.CreateDowntime(ctx).Body(testDowntime).Execute()
			if tc.ExpectedStatusCode < 300 {
				defer cancelDowntime(ctx, t, downtime.GetId())
			}

			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode, "error: %v", err)
		})
	}
}

func TestDowntimeListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DowntimesApi.ListDowntimes(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Downtime
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.Downtime{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Downtime{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DowntimesApi.CreateDowntime(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelByScopeErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.CancelDowntimesByScopeRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.CancelDowntimesByScopeRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.CancelDowntimesByScopeRequest{}, 403},
		"404 Not Found":   {WithTestAuth, datadog.CancelDowntimesByScopeRequest{Scope: "nonexistent"}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DowntimesApi.CancelDowntimesByScope(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelErrors(t *testing.T) {
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

			httpresp, err := Client(ctx).DowntimesApi.CancelDowntime(ctx, 1234).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeGetErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).DowntimesApi.GetDowntime(ctx, 1234).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	// Endpoint will 400 if there are too many tags
	badDowntime := *datadog.NewDowntimeWithDefaults()
	tags := make([]string, 50)
	for i := 0; i < 50; i++ {
		tags[i] = fmt.Sprintf("tag%d", i)
	}
	badDowntime.MonitorTags = &tags

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Downtime
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, badDowntime, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Downtime{}, 403},
		"404 Not Found":   {WithTestAuth, datadog.Downtime{}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DowntimesApi.UpdateDowntime(ctx, 1234).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func cancelDowntime(ctx context.Context, t *testing.T, downtimeID int64) {
	httpresp, err := Client(ctx).DowntimesApi.CancelDowntime(ctx, downtimeID).Execute()
	if err != nil {
		t.Logf("Canceling Downtime: %v failed with %v, Another test may have already canceled this downtime: %v", downtimeID, httpresp.StatusCode, err)
	}
}
