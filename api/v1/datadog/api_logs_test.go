package datadog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestLogHTTPIntake(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	httpLog := datadog.HttpLog{
		Ddsource: datadog.PtrString("source"),
		Ddtags:   datadog.PtrString("go,test,intake"),
		Hostname: datadog.PtrString("datadog-api-client-go-test"),
		Message:  datadog.PtrString("testing message"),
	}

	// Create log entry
	_, httpresp, err := TESTAPICLIENT.LogsApi.SendLog(TESTAUTH).Body(httpLog).Execute()
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestLogsList(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	now := time.Now()
	nanoNow := now.UnixNano()
	source := fmt.Sprintf("go-client-test-%d", nanoNow)
	message := fmt.Sprintf("test-log-list-%d", nanoNow)
	hostname := fmt.Sprintf("datadog-api-client-go-test-%d", nanoNow)

	httpLog := datadog.HttpLog{
		Ddsource: &source,
		Ddtags:   datadog.PtrString("go,test,list"),
		Hostname: &hostname,
		Message:  &message,
	}

	// Create log entry
	_, httpresp, err := TESTAPICLIENT.LogsApi.SendLog(TESTAUTH).Body(httpLog).Execute()
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	secondMessage := "second-" + message
	httpLog.SetMessage(secondMessage)

	// Create second log entry
	_, httpresp, err = TESTAPICLIENT.LogsApi.SendLog(TESTAUTH).Body(httpLog).Execute()
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	logsRequest := datadog.LogsListRequest{
		Query: fmt.Sprintf("source:%s", source),
		Time: datadog.LogsListRequestTime{
			From: now.Add(time.Duration(-3600) * time.Second),
			To:   now.Add(time.Duration(3600) * time.Second),
		},
		Limit: datadog.PtrInt32(2),
		Sort:  datadog.LOGSSORT_TIME_ASCENDING.Ptr(),
	}

	var logsResponse datadog.LogsListResponse

	// Make sure that both log items are indexed
	retry(time.Duration(5)*time.Second, 10, func() bool {
		logsResponse, httpresp, err = TESTAPICLIENT.LogsApi.ListLogs(TESTAUTH).Body(logsRequest).Execute()
		if err != nil {
			t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return 200 == httpresp.StatusCode && 2 == len(logsResponse.GetLogs())
	})

	// Find first log item
	logsRequest.SetLimit(1)

	logsResponse, httpresp, err = TESTAPICLIENT.LogsApi.ListLogs(TESTAUTH).Body(logsRequest).Execute()
	if err != nil {
		t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(t, hostname, content.GetHost())
		assert.Equal(t, message, log.Content.GetMessage())
	}

	// Find second log item
	assert.Assert(t, len(*logsResponse.NextLogId) > 0)
	logsRequest.SetStartAt(logsResponse.GetNextLogId())

	logsResponse, httpresp, err = TESTAPICLIENT.LogsApi.ListLogs(TESTAUTH).Body(logsRequest).Execute()
	if err != nil {
		t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(t, hostname, content.GetHost())
		assert.Equal(t, secondMessage, log.Content.GetMessage())
	}
}
