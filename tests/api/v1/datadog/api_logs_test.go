package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"gotest.tools/assert"
)

func TestLogsList(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	now := TESTCLOCK.Now()
	nanoNow := now.UnixNano()
	source := fmt.Sprintf("go-client-test-%d", nanoNow)
	message := fmt.Sprintf("test-log-list-%d", nanoNow)
	hostname := fmt.Sprintf("datadog-api-client-go-test-%d", nanoNow)

	// Create log entry
	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1)*1000, message,
	)
	httpresp, respBody, err := sendRequest("POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	time.Sleep(time.Duration(500) * time.Microsecond)

	secondMessage := fmt.Sprintf("second-test-log-list-%d", nanoNow)
	httpLog = fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix())*1000, secondMessage,
	)

	// Create second log entry
	httpresp, respBody, err = sendRequest("POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
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
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		logsResponse, httpresp, err = TESTAPICLIENT.LogsApi.ListLogs(TESTAUTH).Body(logsRequest).Execute()
		if err != nil {
			t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return 200 == httpresp.StatusCode && 2 == len(logsResponse.GetLogs())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

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
