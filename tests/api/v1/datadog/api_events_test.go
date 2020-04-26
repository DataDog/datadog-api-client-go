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

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

var testEvent = datadog.Event{
	Title: "test event from go client",
	Text:  "example text",
	Tags: &[]string{
		"test",
		"client:go",
	},
	Priority:       datadog.EVENTPRIORITY_NORMAL.Ptr(),
	SourceTypeName: datadog.PtrString("datadog-api-client-go"),
}

type createEventResponse struct {
	Event datadog.Event `json:"event"`
}

func TestEventLifecycle(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	// Create event
	marshalledEvent, _ := json.Marshal(testEvent)
	httpresp, respBody, err := c.SendRequest("POST", "/api/v1/events", marshalledEvent)
	if err != nil {
		t.Fatalf("Error creating Event %v: Response %s: %v", testEvent, string(respBody), err)
	}
	assert.Equal(t, 202, httpresp.StatusCode)
	var unmarshaledBody createEventResponse
	err = json.Unmarshal(respBody, &unmarshaledBody)
	if err != nil {
		t.Fatalf("Failed unmarshalling event from response: %s", err)
	}
	event := unmarshaledBody.Event

	var fetchedEventResponse datadog.EventResponse

	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		// Check event existence
		fetchedEventResponse, httpresp, err = c.Client.EventsApi.GetEvent(c.Ctx, event.GetId()).Execute()
		if err != nil {
			t.Logf("Error fetching Event %v: Response %s: %v", event.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		return true
	})

	fetchedEvent := fetchedEventResponse.GetEvent()
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, fetchedEvent.GetTitle(), testEvent.GetTitle())
	assert.Equal(t, fetchedEvent.GetText(), testEvent.GetText())
	// not the same!!! assert.Equal(t, testEvent.GetUrl(), fetchedEvent.GetUrl())
	assert.NotEmpty(t, fetchedEvent.GetUrl())

	// Find our event in the full list
	start := event.GetDateHappened() - 10
	end := start + 20

	var eventListResponse datadog.EventListResponse

	// Confirm that the fetchedEvent is inside the getEvents response
	// Use Retry instead of assert as the response may not be empty but may also require
	// some time for our event to show up in the response
	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		var matchedEvent = false
		eventListResponse, httpresp, err = c.Client.EventsApi.ListEvents(c.Ctx).Start(start).End(end).Priority("normal").Sources("datadog-api-client-go").Tags("test,client:go").Unaggregated(true).Execute()
		if err != nil {
			t.Logf("Error fetching events: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
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

	assert.Equal(t, 200, httpresp.StatusCode)

	events := eventListResponse.GetEvents()
	assert.Contains(t, events, fetchedEvent)
}

func TestEventListErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.EventsApi.ListEvents(c.Ctx).Start(345).End(123).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestEventGetErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, httpresp, err := c.Client.EventsApi.GetEvent(c.Ctx, 1234).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
