package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestLogsList(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewLogsApi(Client(ctx))

	now := tests.ClockFromContext(ctx).Now()
	nanoNow := now.UnixNano()
	source := fmt.Sprintf("go-client-test-%d", nanoNow)
	message := fmt.Sprintf("test-log-list-%d", nanoNow)
	hostname := *tests.UniqueEntityName(ctx, t)

	// Create log entry
	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1)*1000, message,
	)
	httpresp, respBody, err := SendRequest(ctx, "POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
	}
	assert.Equal(200, httpresp.StatusCode)
	time.Sleep(time.Duration(500) * time.Microsecond)

	secondMessage := fmt.Sprintf("second-test-log-list-%d", nanoNow)
	httpLog = fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix())*1000, secondMessage,
	)

	// Create second log entry
	httpresp, respBody, err = SendRequest(ctx, "POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		t.Fatalf("Error creating log: Response %s: %v", respBody, err)
	}
	assert.Equal(200, httpresp.StatusCode)

	logsRequest := datadogV1.LogsListRequest{
		Query: datadog.PtrString(fmt.Sprintf("source:%s", source)),
		Time: datadogV1.LogsListRequestTime{
			From: now.Add(time.Duration(-3600) * time.Second),
			To:   now.Add(time.Duration(3600) * time.Second),
		},
		Limit: datadog.PtrInt32(2),
		Sort:  datadogV1.LOGSSORT_TIME_ASCENDING.Ptr(),
	}

	var logsResponse datadogV1.LogsListResponse

	// Make sure that both log items are indexed
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		logsResponse, httpresp, err = api.ListLogs(ctx, logsRequest)
		if err != nil {
			t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200 && len(logsResponse.GetLogs()) == 2
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Find first log item
	logsRequest.SetLimit(1)

	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		logsResponse, httpresp, err = api.ListLogs(ctx, logsRequest)
		if err != nil {
			t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200 && len(logsResponse.GetNextLogId()) > 0
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpresp.StatusCode)

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(hostname, content.GetHost())
		assert.Equal(message, content.GetMessage())
	}

	// Find second log item
	assert.True(len(logsResponse.GetNextLogId()) > 0)
	logsRequest.SetStartAt(logsResponse.GetNextLogId())

	assert.Equal(200, httpresp.StatusCode)
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		logsResponse, httpresp, err = api.ListLogs(ctx, logsRequest)
		if err != nil {
			t.Fatalf("Error listing logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200 && len(logsResponse.GetLogs()) > 0
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, log := range logsResponse.GetLogs() {
		content := log.GetContent()
		assert.Equal(hostname, content.GetHost())
		assert.Equal(secondMessage, content.GetMessage())
	}
}

func TestLogsListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	req := *datadogV1.NewLogsListRequestWithDefaults()
	req.SetStartAt("notanid")
	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.LogsListRequest
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, req, 400},
		"403 Forbidden":   {WithFakeAuth, req, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewLogsApi(Client(ctx))

			_, httpresp, err := api.ListLogs(ctx, tc.Body)
			openAPIErr, ok := err.(datadog.GenericOpenAPIError)
			if !ok {
				t.Fatalf("Unexpected error %T: %v", err, err)
			}
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := openAPIErr.Model().(datadogV1.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := openAPIErr.Model().(datadogV1.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}
