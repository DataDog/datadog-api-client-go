package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestLogsList(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	now := c.Clock.Now()
	nanoNow := now.UnixNano()
	source := fmt.Sprintf("go-client-test-%d", nanoNow)
	message := fmt.Sprintf("test-log-list-%d", nanoNow)
	hostname := fmt.Sprintf("datadog-api-client-go-test-%d", nanoNow)

	// Create log entry
	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1)*1000, message,
	)
	httpresp, respBody, err := c.SendRequest("POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	time.Sleep(time.Duration(500) * time.Microsecond)

	secondMessage := fmt.Sprintf("second-test-log-list-%d", nanoNow)
	httpLog = fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix())*1000, secondMessage,
	)

	// Create second log entry
	httpresp, respBody, err = c.SendRequest("POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

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
		logsResponse, httpresp, err = c.Client.LogsApi.ListLogs(c.Ctx).Body(logsRequest).Execute()
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

	logsResponse, httpresp, err = c.Client.LogsApi.ListLogs(c.Ctx).Body(logsRequest).Execute()
	if err != nil {
		t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(t, hostname, content.GetHost())
		assert.Equal(t, message, content.GetMessage())
	}

	// Find second log item
	assert.True(t, len(*logsResponse.NextLogId) > 0)
	logsRequest.SetStartAt(logsResponse.GetNextLogId())

	logsResponse, httpresp, err = c.Client.LogsApi.ListLogs(c.Ctx).Body(logsRequest).Execute()
	if err != nil {
		t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(t, hostname, content.GetHost())
		assert.Equal(t, secondMessage, content.GetMessage())
	}
}

func TestLogsListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	req := *datadog.NewLogsListRequestWithDefaults()
	req.SetStartAt("notanid")
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsListRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, req, 400},
		"403 Forbidden":   {WithFakeAuth, req, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.LogsApi.ListLogs(c.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetErrors())
			} else {
				apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(t, ok)
				assert.NotEmpty(t, apiError.GetError())
			}
		})
	}
}
