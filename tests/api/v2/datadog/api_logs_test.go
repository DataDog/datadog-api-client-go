package test

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/tests"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestLogsList(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)

	defer finish()
	assert := tests.Assert(ctx, t)

	now := tests.ClockFromContext(ctx).Now()
	suffix := fmt.Sprintf("post-%v", now.Unix())

	err := sendLogs(ctx, t, suffix)
	if err != nil {
		t.Fatalf("Error sending logs: %v", err)
	}

	var response datadog.LogsListResponse
	var httpResp *http.Response

	filter := datadog.NewLogsListPayloadFilter()
	filter.SetQuery(suffix)
	filter.SetFrom(now.Add(time.Duration(-2)*time.Hour))
	filter.SetTo(now.Add(time.Duration(2)*time.Hour))

	request := datadog.NewLogsListPayloadWithDefaults()
	request.SetFilter(*filter)

	// Make sure both logs are indexed
	err = tests.Retry(time.Duration(5) * time.Second, 30, func() bool {
		response, httpResp, err = Client(ctx).LogsApi.ListLogs(ctx).Body(*request).Execute()
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent logs: %v", err)
	}

	// Sort works correctly
	request.SetSort(datadog.LOGSSORT_TIMESTAMP_ASCENDING)

	response, httpResp, err = Client(ctx).LogsApi.ListLogs(ctx).Body(*request).Execute()
	if err != nil {
		t.Fatalf("Could not list logs: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	content := response.GetData()[0].GetContent()
	assert.Equal("test-log-list-1 " + suffix, content.GetMessage())

	content = response.GetData()[1].GetContent()
	assert.Equal("test-log-list-2 " + suffix, content.GetMessage())

	request.SetSort(datadog.LOGSSORT_TIMESTAMP_DESCENDING)

	response, httpResp, err = Client(ctx).LogsApi.ListLogs(ctx).Body(*request).Execute()
	if err != nil {
		t.Fatalf("Could not list logs: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	content = response.GetData()[0].GetContent()
	assert.Equal("test-log-list-2 " + suffix, content.GetMessage())

	content = response.GetData()[1].GetContent()
	assert.Equal("test-log-list-1 " + suffix, content.GetMessage())

	// Paging
	page := datadog.NewLogsListPayloadPage()
	page.SetLimit(1)
	request.SetPage(*page)
	response, httpResp, err = Client(ctx).LogsApi.ListLogs(ctx).Body(*request).Execute()
	if err != nil {
		t.Fatalf("Could not list logs: %v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	respMeta := response.GetMeta()
	respPage := respMeta.GetPage()
	cursor := respPage.GetAfter()
	firstId := response.GetData()[0].GetId()

	request.Page.SetCursor(cursor)
	response, httpResp, err = Client(ctx).LogsApi.ListLogs(ctx).Body(*request).Execute()
	if err != nil {
		t.Fatalf("Could not list logs: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	secondId := response.GetData()[0].GetId()

	assert.NotEqual(firstId, secondId)
}

func sendLogs(ctx context.Context, t *testing.T, suffix string) error {
	now := tests.ClockFromContext(ctx).Now()
	source := fmt.Sprintf("go-client-test-%s", suffix)
	firstMessage := fmt.Sprintf("test-log-list-1 %s", suffix)
	secondMessage := fmt.Sprintf("test-log-list-2 %s", suffix)
	hostname := *tests.UniqueEntityName(ctx, t)

	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1000)*1000, firstMessage,
	)
	httpresp, respBody, err := SendRequest(ctx, "POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		return fmt.Errorf("response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("status=%d response=%v", httpresp.StatusCode, respBody)
	}

	httpLog = fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1)*1000, secondMessage,
	)
	httpresp, respBody, err = SendRequest(ctx, "POST", "https://http-intake.logs.datadoghq.com/v1/input", []byte(httpLog))
	if err != nil {
		return fmt.Errorf("error creating log: Response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("unable to send log: Status=%d Response=%v", httpresp.StatusCode, respBody)
	}

	return  nil
}
