package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestLogsList(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := datadog.NewLogsApi(client)

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
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(*request))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent logs: %v", err)
	}

	// Sort works correctly
	request.SetSort(datadog.LOGSSORT_TIMESTAMP_ASCENDING)
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(*request))
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
		response, httpResp, err = api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(*request))
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
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(*request))
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
	firstID := response.GetData()[0].GetId()

	request.Page.SetCursor(cursor)
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(*request))
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
	secondID := response.GetData()[0].GetId()

	assert.NotEqual(firstID, secondID)
}

func TestGetLogsNilBody(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := datadog.NewLogsApi(client)

	_, httpResp, err := api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters())
	if err != nil {
		t.Fatalf("Could not list logs: %v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
}

func TestLogsListGet(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := datadog.NewLogsApi(client)

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
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters().
			WithFilterQuery(*suffix).
			WithFilterFrom(from).
			WithFilterTo(to))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent logs: %v", err)
	}

	// Sort works correctly
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters().
			WithFilterQuery(*suffix).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithSort(datadog.LOGSSORT_TIMESTAMP_ASCENDING))
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

	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters().
			WithFilterQuery(*suffix).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithSort(datadog.LOGSSORT_TIMESTAMP_DESCENDING))
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
	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters().
			WithFilterQuery(*suffix).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithPageLimit(1))
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
	firstID := response.GetData()[0].GetId()

	err = tests.Retry(time.Duration(15)*time.Second, 10, func() bool {
		response, httpResp, err = api.ListLogsGet(ctx, *datadog.NewListLogsGetOptionalParameters().
			WithFilterQuery(*suffix).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithPageLimit(1).
			WithPageCursor(cursor))
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
	secondID := response.GetData()[0].GetId()

	assert.NotEqual(firstID, secondID)
}

func sendLogs(ctx context.Context, t *testing.T, client *common.APIClient, suffix string) error {
	now := tests.ClockFromContext(ctx).Now()
	source := fmt.Sprintf("go-client-test-%s", suffix)
	firstMessage := fmt.Sprintf("test-log-list-1 %s", suffix)
	secondMessage := fmt.Sprintf("test-log-list-2 %s", suffix)
	hostname := *tests.UniqueEntityName(ctx, t)

	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,list", "hostname": "%s", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, hostname, (now.Unix()-1000)*1000, firstMessage,
	)

	domain, err := getTestDomain(ctx, client)
	if err != nil {
		return fmt.Errorf("parsing domain: %v", err)
	}
	intakeURL := fmt.Sprintf("https://http-intake.logs.%s/v1/input", domain)

	httpresp, respBody, err := SendRequest(ctx, "POST", intakeURL, []byte(httpLog))
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

	httpresp, respBody, err = SendRequest(ctx, "POST", intakeURL, []byte(httpLog))
	if err != nil {
		return fmt.Errorf("error creating log: Response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("unable to send log: Status=%d Response=%v", httpresp.StatusCode, respBody)
	}

	return nil
}
