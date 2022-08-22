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
)

func TestDowntimeLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewDowntimesApi(Client(ctx))

	start := tests.ClockFromContext(ctx).Now()
	testDowntime := datadogV1.Downtime{
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		End:      *datadog.NewNullableInt64(datadog.PtrInt64(start.Unix() + 60*60)),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    []string{"*"},
		Recurrence: *datadogV1.NewNullableDowntimeRecurrence(
			&datadogV1.DowntimeRecurrence{
				Type:      datadog.PtrString("weeks"),
				Period:    datadog.PtrInt32(1),
				WeekDays:  []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				UntilDate: *datadog.NewNullableInt64(datadog.PtrInt64(start.Unix() + 21*24*60*60)), // +21d
			}),
	}

	// Create downtime
	downtime, httpresp, err := api.CreateDowntime(ctx, testDowntime)
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
	editedDowntime := datadogV1.Downtime{Message: datadog.PtrString(fmt.Sprintf("%s-updated", testDowntime.GetMessage()))}
	updatedDowntime, httpresp, err := api.UpdateDowntime(ctx, downtime.GetId(), editedDowntime)
	if err != nil {
		t.Errorf("Error updating Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(editedDowntime.GetMessage(), updatedDowntime.GetMessage())
	val, set = updatedDowntime.GetUpdaterIdOk()
	assert.True(set)
	assert.NotNil(val)

	// Check downtime existence
	fetchedDowntime, httpresp, err := api.GetDowntime(ctx, downtime.GetId())
	if err != nil {
		t.Errorf("Error fetching Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedDowntime.GetMessage(), fetchedDowntime.GetMessage())

	// Find our downtime in the full list
	downtimes, httpresp, err := api.ListDowntimes(ctx)
	if err != nil {
		t.Errorf("Error fetching downtimes: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(downtimes, fetchedDowntime)

	// Cancel downtime
	httpresp, err = api.CancelDowntime(ctx, downtime.GetId())
	if err != nil {
		t.Errorf("Error canceling Downtime %v: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)

	// Check downtime status
	fetchedDowntime, httpresp, err = api.GetDowntime(ctx, downtime.GetId())
	if httpresp.StatusCode != 200 {
		t.Errorf("Downtime %v should still exist: Response %s: %v", downtime.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.False(fetchedDowntime.GetActive())
	assert.True(fetchedDowntime.GetDisabled())
}

func TestMonitorDowntime(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewDowntimesApi(Client(ctx))
	monitorsApi := datadogV1.NewMonitorsApi(Client(ctx))

	// Create monitor
	tm := testMonitor(ctx, t)
	monitor, httpresp, err := monitorsApi.CreateMonitor(ctx, tm)
	if err != nil {
		t.Fatalf("Error creating Monitor %v: Response %s: %v", tm, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	monitorID := monitor.GetId()
	defer deleteMonitor(ctx, t, monitorID)

	start := tests.ClockFromContext(ctx).Now()
	testDowntime := datadogV1.Downtime{
		Message:   tests.UniqueEntityName(ctx, t),
		Start:     datadog.PtrInt64(start.Unix()),
		Timezone:  datadog.PtrString("Etc/UTC"),
		Scope:     []string{"*"},
		MonitorId: *datadog.NewNullableInt64(datadog.PtrInt64(monitorID)),
	}

	// Create downtime
	downtime, httpresp, err := api.CreateDowntime(ctx, testDowntime)
	if err != nil {
		t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer cancelDowntime(ctx, t, downtime.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(monitorID, downtime.GetMonitorId())
}

func TestScopedDowntime(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewDowntimesApi(Client(ctx))

	start := tests.ClockFromContext(ctx).Now()

	scopeClient := fmt.Sprintf("test:client-%s-%d", t.Name(), tests.ClockFromContext(ctx).Now().UnixNano())
	scopeGo := fmt.Sprintf("test:go-%s-%d", t.Name(), tests.ClockFromContext(ctx).Now().UnixNano())

	testDowntimes := []datadogV1.Downtime{{
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    []string{scopeClient, scopeGo},
	}, {
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    []string{scopeGo},
	}, {
		Message:  tests.UniqueEntityName(ctx, t),
		Start:    datadog.PtrInt64(start.Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope:    []string{scopeClient},
	},
	}

	// Create downtimes
	downtimes := make([]datadogV1.Downtime, len(testDowntimes))
	for i, testDowntime := range testDowntimes {
		downtime, httpresp, err := api.CreateDowntime(ctx, testDowntime)
		if err != nil {
			t.Fatalf("Error creating Downtime %v: Response %s: %v", testDowntime, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		defer cancelDowntime(ctx, t, downtime.GetId())
		assert.Equal(200, httpresp.StatusCode)

		assert.False(downtime.GetDisabled())
		downtimes[i] = downtime
	}

	// Cancel downtimes with the scopeGo
	cancelDowntimesByScopeRequest := datadogV1.CancelDowntimesByScopeRequest{
		Scope: scopeGo,
	}
	canceledDowntimesIds, httpresp, err := api.CancelDowntimesByScope(ctx, cancelDowntimesByScopeRequest)
	if err != nil {
		t.Fatalf("Error canceling downtimes by scope %s: Response: %s: %v", scopeGo, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	canceledIds := canceledDowntimesIds.GetCancelledIds()
	expectedCanceledIds := []int64{downtimes[1].GetId()}
	assert.Equal(expectedCanceledIds, canceledIds)
}

func TestDowntimeRecurrence(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	start := tests.ClockFromContext(ctx).Now()
	testCases := map[string]struct {
		DowntimeRecurence  datadogV1.DowntimeRecurrence
		ExpectedStatusCode int
	}{
		"once a year": {datadogV1.DowntimeRecurrence{
			Type:   datadog.PtrString("years"),
			Period: datadog.PtrInt32(1),
		}, 200},
		"invalid type hours": {datadogV1.DowntimeRecurrence{
			Type:   datadog.PtrString("hours"), // Choose from: days, weeks, months, years.
			Period: datadog.PtrInt32(1),
		}, 400},
		"invalid weekdays": {datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"mon", "tue"},
		}, 400},
		/* Only applicable when type is weeks -> unclear error code
		"weekdays only with type weeks not days": { datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("days"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon"},
		}, 400},
		"weekdays only with type weeks not months": { datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("months"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon"},
		}, 400},
		"weekdays only with type weeks not years": { datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("years"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon"},
		}, 400}, */
		"until date": {datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilDate: *datadog.NewNullableInt64(
				datadog.PtrInt64(start.Unix() + 21*24*60*60), // +21d
			)}, 200},
		"until occurrences": {datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			UntilOccurrences: *datadog.NewNullableInt32(
				datadog.PtrInt32(3),
			)}, 200},
		"until occurences and until date are mutually exclusive": {datadogV1.DowntimeRecurrence{
			Type:     datadog.PtrString("weeks"),
			Period:   datadog.PtrInt32(1),
			WeekDays: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
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
			api := datadogV1.NewDowntimesApi(Client(ctx))

			testDowntime := datadogV1.Downtime{
				Message:    datadog.PtrString(fmt.Sprintf("%s; %s", *tests.UniqueEntityName(ctx, t), name)),
				Start:      datadog.PtrInt64(start.Unix()),
				End:        *datadog.NewNullableInt64(datadog.PtrInt64(start.Unix() + 60*60)),
				Timezone:   datadog.PtrString("Etc/UTC"),
				Scope:      []string{"*"},
				Recurrence: *datadogV1.NewNullableDowntimeRecurrence(&tc.DowntimeRecurence),
			}

			downtime, httpresp, err := api.CreateDowntime(ctx, testDowntime)
			if tc.ExpectedStatusCode < 300 {
				defer cancelDowntime(ctx, t, downtime.GetId())
			}
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode, "error: %v", err)
		})
	}
}

func TestDowntimeListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewDowntimesApi(Client(ctx))

			_, httpresp, err := api.ListDowntimes(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.Downtime
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.Downtime{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.Downtime{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDowntimesApi(Client(ctx))

			_, httpresp, err := api.CreateDowntime(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelByScopeErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.CancelDowntimesByScopeRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.CancelDowntimesByScopeRequest{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.CancelDowntimesByScopeRequest{}, 403},
		"404 Not Found":   {WithTestAuth, datadogV1.CancelDowntimesByScopeRequest{Scope: "nonexistent"}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDowntimesApi(Client(ctx))

			_, httpresp, err := api.CancelDowntimesByScope(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeCancelErrors(t *testing.T) {
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
			api := datadogV1.NewDowntimesApi(Client(ctx))

			httpresp, err := api.CancelDowntime(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeGetErrors(t *testing.T) {
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
			api := datadogV1.NewDowntimesApi(Client(ctx))

			_, httpresp, err := api.GetDowntime(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDowntimeUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	// Endpoint will 400 if there are too many tags
	badDowntime := *datadogV1.NewDowntimeWithDefaults()
	tags := make([]string, 50)
	for i := 0; i < 50; i++ {
		tags[i] = fmt.Sprintf("tag%d", i)
	}
	badDowntime.MonitorTags = tags

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.Downtime
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, badDowntime, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.Downtime{}, 403},
		"404 Not Found":   {WithTestAuth, datadogV1.Downtime{}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDowntimesApi(Client(ctx))

			_, httpresp, err := api.UpdateDowntime(ctx, 1234, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func cancelDowntime(ctx context.Context, t *testing.T, downtimeID int64) {
	api := datadogV1.NewDowntimesApi(Client(ctx))

	_, err := api.CancelDowntime(ctx, downtimeID)
	if err != nil {
		t.Logf("Canceling Downtime: %v failed, Another test may have already canceled this downtime: %v", downtimeID, err)
	}
}
