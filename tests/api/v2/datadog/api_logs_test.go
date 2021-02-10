package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func TestLogsList(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	suffix := tests.UniqueEntityName(ctx, t)

	err := sendLogs(ctx, t, client, *suffix)
	if err != nil {
		t.Fatalf("Error sending logs: %v", err)
	}

	var response datadog.LogsListResponse
	var httpResp *http.Response

	filter := datadog.NewLogsQueryFilter()
	filter.SetQuery(*suffix)
	filter.SetFrom("now-2h")
	filter.SetTo("now+2h")

	request := datadog.NewLogsListRequestWithDefaults()
	request.SetFilter(*filter)

	// Make sure both logs are indexed
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogs(ctx).Body(*request).Execute()
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent logs: %v", err)
	}

	// Sort works correctly
	request.SetSort(datadog.LOGSSORT_TIMESTAMP_ASCENDING)
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogs(ctx).Body(*request).Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	attributes := response.GetData()[0].GetAttributes()
	assert.Equal("test-log-list-1 "+*suffix, attributes.GetMessage())

	attributes = response.GetData()[1].GetAttributes()
	assert.Equal("test-log-list-2 "+*suffix, attributes.GetMessage())

	request.SetSort(datadog.LOGSSORT_TIMESTAMP_DESCENDING)

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogs(ctx).Body(*request).Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	attributes = response.GetData()[0].GetAttributes()
	assert.Equal("test-log-list-2 "+*suffix, attributes.GetMessage())

	attributes = response.GetData()[1].GetAttributes()
	assert.Equal("test-log-list-1 "+*suffix, attributes.GetMessage())

	// Paging
	page := datadog.NewLogsListRequestPage()
	page.SetLimit(1)
	request.SetPage(*page)
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogs(ctx).Body(*request).Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 1 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	respMeta := response.GetMeta()
	respPage := respMeta.GetPage()
	cursor := respPage.GetAfter()
	firstId := response.GetData()[0].GetId()

	request.Page.SetCursor(cursor)
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogs(ctx).Body(*request).Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 1 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	secondId := response.GetData()[0].GetId()

	assert.NotEqual(firstId, secondId)
}

func TestLogsListGet(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	now := tests.ClockFromContext(ctx).Now()
	suffix := tests.UniqueEntityName(ctx, t)

	err := sendLogs(ctx, t, client, *suffix)
	if err != nil {
		t.Fatalf("Error sending logs: %v", err)
	}

	var response datadog.LogsListResponse
	var httpResp *http.Response

	from := now.Add(time.Duration(-2) * time.Hour)
	to := now.Add(time.Duration(2) * time.Hour)

	// Make sure both logs are indexed
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogsGet(ctx).
			FilterQuery(*suffix).
			FilterFrom(from).
			FilterTo(to).
			Execute()
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent logs: %v", err)
	}

	// Sort works correctly
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogsGet(ctx).
			FilterQuery(*suffix).
			FilterFrom(from).
			FilterTo(to).
			Sort(datadog.LOGSSORT_TIMESTAMP_ASCENDING).
			Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	attributes := response.GetData()[0].GetAttributes()
	assert.Equal("test-log-list-1 "+*suffix, attributes.GetMessage())

	attributes = response.GetData()[1].GetAttributes()
	assert.Equal("test-log-list-2 "+*suffix, attributes.GetMessage())

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogsGet(ctx).
			FilterQuery(*suffix).
			FilterFrom(from).
			FilterTo(to).
			Sort(datadog.LOGSSORT_TIMESTAMP_DESCENDING).
			Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	attributes = response.GetData()[0].GetAttributes()
	assert.Equal("test-log-list-2 "+*suffix, attributes.GetMessage())

	attributes = response.GetData()[1].GetAttributes()
	assert.Equal("test-log-list-1 "+*suffix, attributes.GetMessage())

	// Paging
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogsGet(ctx).
			FilterQuery(*suffix).
			FilterFrom(from).
			FilterTo(to).
			PageLimit(1).
			Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 1 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	respMeta := response.GetMeta()
	respPage := respMeta.GetPage()
	cursor := respPage.GetAfter()
	firstId := response.GetData()[0].GetId()

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = client.LogsApi.ListLogsGet(ctx).
			FilterQuery(*suffix).
			FilterFrom(from).
			FilterTo(to).
			PageLimit(1).
			PageCursor(cursor).
			Execute()
		if err != nil {
			t.Fatalf("Could not list logs: %v", err)
		}
		return 200 == httpResp.StatusCode && 1 == len(response.GetData())
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	secondId := response.GetData()[0].GetId()

	assert.NotEqual(firstId, secondId)
}

func sendLogs(ctx context.Context, t *testing.T, client *datadog.APIClient, suffix string) error {
	now := tests.ClockFromContext(ctx).Now()
	source := fmt.Sprintf("go-client-test-%s", suffix)
	firstMessage := fmt.Sprintf("test-log-list-1 %s", suffix)
	secondMessage := fmt.Sprintf("test-log-list-2 %s", suffix)
	hostname := *tests.UniqueEntityName(ctx, t)

	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1000)*1000, firstMessage,
	)

	domain, err := GetTestDomain(ctx, client)
	if err != nil {
		return fmt.Errorf("parsing domain: %v", err)
	}
	intakeUrl := fmt.Sprintf("https://http-intake.logs.%s/v1/input", domain)

	httpresp, respBody, err := SendRequest(ctx, "POST", intakeUrl, []byte(httpLog))
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

	httpresp, respBody, err = SendRequest(ctx, "POST", intakeUrl, []byte(httpLog))
	if err != nil {
		return fmt.Errorf("error creating log: Response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("unable to send log: Status=%d Response=%v", httpresp.StatusCode, respBody)
	}

	return nil
}
