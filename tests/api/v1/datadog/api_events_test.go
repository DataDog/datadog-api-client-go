/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
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
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create event
	marshalledEvent, _ := json.Marshal(testEvent)
	httpresp, respBody, err := sendRequest("POST", "/api/v1/events", marshalledEvent)
	if err != nil {
		t.Fatalf("Error creating Event %v: Response %s: %v", testEvent, string(respBody), err)
	}
	assert.Equal(t, httpresp.StatusCode, 202)
	var unmarshaledBody createEventResponse
	err = json.Unmarshal(respBody, &unmarshaledBody)
	if err != nil {
		t.Fatalf("Failed unmarshalling event from response: %s", err)
	}
	event := unmarshaledBody.Event

	var fetchedEventResponse datadog.EventResponse

	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		// Check event existence
		fetchedEventResponse, httpresp, err = TESTAPICLIENT.EventsApi.GetEvent(TESTAUTH, event.GetId()).Execute()
		if err != nil {
			t.Logf("Error fetching Event %v: Response %s: %v", event.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		return true
	})

	fetchedEvent := fetchedEventResponse.GetEvent()
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, testEvent.GetTitle(), fetchedEvent.GetTitle())
	assert.Equal(t, testEvent.GetText(), fetchedEvent.GetText())
	// not the same!!! assert.Equal(t, testEvent.GetUrl(), fetchedEvent.GetUrl())
	assert.Assert(t, fetchedEvent.GetUrl() != "")

	// Find our event in the full list
	start := event.GetDateHappened() - 10
	end := start + 20

	var eventListResponse datadog.EventListResponse

	// Confirm that the fetchedEvent is inside the getEvents response
	// Use Retry instead of assert as the response may not be empty but may also require
	// some time for our event to show up in the response
	tests.Retry(time.Duration(5*time.Second), 20, func() bool {
		var matchedEvent = false
		eventListResponse, httpresp, err = TESTAPICLIENT.EventsApi.ListEvents(TESTAUTH).Start(start).End(end).Priority("normal").Sources("datadog-api-client-go").Tags("test,client:go").Unaggregated(true).Execute()
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

	assert.Equal(t, httpresp.StatusCode, 200)

	events := eventListResponse.GetEvents()
	assert.Assert(t, is.Contains(events, fetchedEvent))
}
