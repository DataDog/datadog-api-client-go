package datadog_test

import (
	is "gotest.tools/assert/cmp"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var testEvent = datadog.Event{
	Title: "test event from go client",
	Text:  "example text",
	Tags: &[]string{
		"test",
		"client:go",
	},
	Priority:       datadog.PtrString("normal"),
	SourceTypeName: datadog.PtrString("datadog-api-client-go"),
}

func TestEventLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create event
	eventResponse, httpresp, err := TESTAPICLIENT.EventsApi.CreateEvent(TESTAUTH).Body(testEvent).Execute()
	if err != nil {
		t.Fatalf("Error creating Event %v: Response %s: %v", testEvent, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	// defer deleteEvent(eventResponse.GetEvent().GetId())
	assert.Equal(t, httpresp.StatusCode, 202)

	event := eventResponse.GetEvent()
	status := eventResponse.GetStatus()
	assert.Equal(t, status, "ok")
	assert.Equal(t, event.GetTitle(), testEvent.GetTitle())
	assert.Equal(t, event.GetText(), testEvent.GetText())
	assert.Assert(t, event.GetUrl() != "")

	var fetchedEventResponse datadog.EventResponse

	retry(time.Duration(5*time.Second), 20, func() bool {
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
	assert.Equal(t, event.GetTitle(), fetchedEvent.GetTitle())
	assert.Equal(t, event.GetText(), fetchedEvent.GetText())
	// not the same!!! assert.Equal(t, event.GetUrl(), fetchedEvent.GetUrl())
	assert.Assert(t, fetchedEvent.GetUrl() != "")

	// Find our event in the full list
	start := event.GetDateHappened() - 10
	end := start + 20

	var eventListResponse datadog.EventListResponse

	// Confirm that the fetchedEvent is inside the getEvents response
	// Use retry instead of assert as the response may not be empty but may also require
	// some time for our event to show up in the response
	retry(time.Duration(5*time.Second), 20, func() bool {
		eventListResponse, httpresp, err = TESTAPICLIENT.EventsApi.ListEvents(TESTAUTH).Start(start).End(end).Priority("normal").Sources("datadog-api-client-go").Tags("test,client:go").Unaggregated(true).Execute()
		if err != nil {
			t.Logf("Error fetching events: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		} else if len(eventListResponse.GetEvents()) == 0 {
			events := eventListResponse.GetEvents()
			var matchedEvent = false
			for e := range events {
				if events[e] == fetchedEvent {
					matchedEvent = true
				}
			}
			return matchedEvent
		}
		return true
	})

	assert.Equal(t, httpresp.StatusCode, 200)

	events := eventListResponse.GetEvents()
	assert.Assert(t, is.Contains(events, fetchedEvent))
}
