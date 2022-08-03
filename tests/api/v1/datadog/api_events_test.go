/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	"github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

var testEvent = datadog.Event{
	Text: common.PtrString("example text"),
	Tags: []string{
		"test",
		"client:go",
	},
	Priority: *datadog.NewNullableEventPriority(datadog.EVENTPRIORITY_NORMAL.Ptr()),
}

type createEventResponse struct {
	Event datadog.Event `json:"event"`
}

func TestEventLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	testEvent.SetTitle(*tests.UniqueEntityName(ctx, t))
	api := datadog.NewEventsApi(Client(ctx))

	// Create event
	marshalledEvent, _ := json.Marshal(testEvent)
	httpresp, respBody, err := SendRequest(ctx, "POST", "/api/v1/events", marshalledEvent)
	if err != nil {
		t.Fatalf("Error creating Event %v: Response %s: %v", testEvent, string(respBody), err)
	}
	assert.Equal(202, httpresp.StatusCode)
	var unmarshaledBody createEventResponse
	err = json.Unmarshal(respBody, &unmarshaledBody)
	if err != nil {
		t.Fatalf("Failed unmarshalling event from response: %s", err)
	}
	event := unmarshaledBody.Event

	var fetchedEventResponse datadog.EventResponse

	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		// Check event existence
		fetchedEventResponse, httpresp, err = api.GetEvent(ctx, event.GetId())
		if err != nil {
			t.Logf("Error fetching Event %v: Response %s: %v", event.GetId(), err.(common.GenericOpenAPIError).Body(), err)
			return false
		}
		return true
	})

	fetchedEvent := fetchedEventResponse.GetEvent()
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(fetchedEvent.GetTitle(), testEvent.GetTitle())
	assert.Equal(fetchedEvent.GetText(), testEvent.GetText())
	// not the same!!! assert.Equal(testEvent.GetUrl(), fetchedEvent.GetUrl())
	assert.NotEmpty(fetchedEvent.GetUrl())

	// Find our event in the full list
	start := event.GetDateHappened() - 10
	end := start + 20

	var eventListResponse datadog.EventListResponse

	// Confirm that the fetchedEvent is inside the getEvents response
	// Use Retry instead of assert as the response may not be empty but may also require
	// some time for our event to show up in the response
	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		var matchedEvent = false
		eventListResponse, httpresp, err = api.ListEvents(ctx, start, end, *datadog.NewListEventsOptionalParameters().WithPriority("normal").
			WithTags("test,client:go").
			WithUnaggregated(true))
		if err != nil {
			t.Logf("Error fetching events: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
		} else {
			events := eventListResponse.GetEvents()
			for e := range events {
				if events[e].GetId() == fetchedEvent.GetId() {
					matchedEvent = true
				}
			}
		}
		return matchedEvent
	})

	assert.Equal(200, httpresp.StatusCode)

	events := eventListResponse.GetEvents()
	assert.Contains(events, fetchedEvent)
}

func TestEventListErrors(t *testing.T) {
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
			api := datadog.NewEventsApi(Client(ctx))

			_, httpresp, err := api.ListEvents(ctx, 345, 123)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestEventGetErrors(t *testing.T) {
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
			api := datadog.NewEventsApi(Client(ctx))

			_, httpresp, err := api.GetEvent(ctx, 1234)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
