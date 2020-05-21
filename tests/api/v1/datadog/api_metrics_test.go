/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
)

func TestMetrics(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	api := Client(ctx).MetricsApi
	now := tests.ClockFromContext(ctx).Now().Unix()

	// the API would replace everything by underscores anyway, making us unable to search for metric by this value
	// in the tests.Retry loop below
	testMetric := strings.ReplaceAll(*tests.UniqueEntityName(ctx, t), "-", "_")
	testQuery := fmt.Sprintf("avg:%s{bar:baz}by{host}", testMetric)

	metricsPayload := fmt.Sprintf(
		`{"series": [{"host": "go-client-test-host", "metric": "%s", "points": [[%f, 10.5], [%f, 11]], "tags": ["%s", "%s"]}]}`,
		testMetric, float64(now-60), float64(now), "tag:foo", "bar:baz",
	)
	httpresp, respBody, err := SendRequest(ctx, "POST", "/api/v1/series", []byte(metricsPayload))
	if err != nil {
		t.Fatalf("Error submitting metric: Response %s: %v", string(respBody), err)
	}
	assert.Equal(httpresp.StatusCode, 202)
	assert.Equal(`{"status": "ok"}`, string(respBody))

	// Check that the metric was submitted successfully
	err = tests.Retry(10*time.Second, 10, func() bool {
		metrics, httpresp, err := api.ListActiveMetrics(ctx).From(now).Execute()
		if err != nil {
			t.Logf("Error getting list of active metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		if httpresp.StatusCode != 200 {
			return false
		}

		found := false
		for _, metric := range metrics.GetMetrics() {
			if metric == testMetric {
				found = true
				break
			}
		}
		return found
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

	// Test query
	queryResult, httpresp, err := api.QueryMetrics(ctx).From(now - 100).To(now + 100).Query(testQuery).Execute()
	if err != nil {
		t.Errorf("Error making query %s: Response %s: %v", testQuery, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal([]string{"host"}, queryResult.GetGroupBy())
	assert.Equal(testQuery, queryResult.GetQuery())
	assert.Equal((now-100)*1000, queryResult.GetFromDate())
	assert.Equal((now+100)*1000, queryResult.GetToDate())
	assert.Equal("ok", queryResult.GetStatus())
	assert.Equal("time_series", queryResult.GetResType())
	assert.Equal(1, len(queryResult.GetSeries()))
	series := queryResult.GetSeries()[0]
	assert.Equal(int64(2), series.GetLength())
	assert.Equal("avg", series.GetAggr())
	assert.Equal(testMetric, series.GetDisplayName())
	assert.Equal(testMetric, series.GetMetric())
	assert.Equal(series.GetPointlist()[0][0], float64(series.GetStart()))
	assert.Equal(series.GetPointlist()[1][0], float64(series.GetEnd()))
	assert.Equal(10.5, series.GetPointlist()[0][1])
	assert.Equal(11., series.GetPointlist()[1][1])

	// Test search
	searchQuery := fmt.Sprintf("metrics:%s", testMetric)
	searchResult, httpresp, err := api.ListMetrics(ctx).Q(searchQuery).Execute()
	if err != nil {
		t.Errorf("Error searching metrics %s: Response %s: %v", searchQuery, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	metrics := searchResult.Results.GetMetrics()
	assert.Equal(1, len(metrics))
	assert.Equal(testMetric, metrics[0])

	// Test metric metadata
	metadata, httpresp, err := api.GetMetricMetadata(ctx, testMetric).Execute()
	if err != nil {
		t.Errorf("Error getting metric metadata for %s: Response %s: %v", testMetric, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Nil(metadata.Description)
	assert.Nil(metadata.Integration)
	assert.Nil(metadata.PerUnit)
	assert.Nil(metadata.Unit)
	assert.Nil(metadata.ShortName)
	assert.Nil(metadata.StatsdInterval)
	assert.Nil(metadata.Type)

	newMetadata := datadog.MetricMetadata{
		Description:    datadog.PtrString("description"),
		PerUnit:        datadog.PtrString("second"),
		Unit:           datadog.PtrString("byte"),
		ShortName:      datadog.PtrString("short_name"),
		StatsdInterval: datadog.PtrInt64(20),
		Type:           datadog.PtrString("count"),
	}

	metadata, httpresp, err = api.UpdateMetricMetadata(ctx, testMetric).Body(newMetadata).Execute()
	if err != nil {
		t.Errorf("Error editing metric metadata for %s: Response %s: %v", testMetric, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("description", metadata.GetDescription())
	assert.Nil(metadata.Integration)
	assert.Equal("second", metadata.GetPerUnit())
	assert.Equal("byte", metadata.GetUnit())
	assert.Equal("short_name", metadata.GetShortName())
	assert.Equal(int64(20), metadata.GetStatsdInterval())
	assert.Equal("count", metadata.GetType())
}

func TestMetricListActive(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	data := setupGock(ctx, t, "metrics/active_metrics.json", "GET", "metrics")
	defer gock.Off()

	var expected datadog.MetricsListResponse
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).MetricsApi

	_, _, err := api.ListActiveMetrics(ctx).Execute()
	assert.NotNil(err) // Parameter `from` required

	metrics, _, err := api.ListActiveMetrics(ctx).From(1).Host("host").Execute()
	assert.Nil(err)
	assert.Equal(expected.GetMetrics(), metrics.GetMetrics())
	assert.Equal(expected.GetFrom(), metrics.GetFrom())
}

func TestMetricsListActive400Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Error 400 cannot be triggered from the client due to client side validation, so mock it
	res, err := tests.ReadFixture("fixtures/metrics/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/metrics").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MetricsApi.ListActiveMetrics(ctx).From(-1).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMetricsListActiveErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// Error 400 cannot be triggered from the client due to client side validation, so mock it
		// "400 Bad Request": {WithTestAuth,400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MetricsApi.ListActiveMetrics(ctx).From(-1).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMetricsMetadataGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MetricsApi.GetMetricMetadata(ctx, "ametric").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMetricsMetadataUpdate400Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Error 400 cannot be triggered from the client due to client side validation, so mock it
	res, err := tests.ReadFixture("fixtures/metrics/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Put("/api/v1/metrics/ametric").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MetricsApi.UpdateMetricMetadata(ctx, "ametric").Body(datadog.MetricMetadata{}).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMetricsMetadataUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.MetricMetadata
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, datadog.MetricMetadata{}, 403},
		"404 Not Found": {WithTestAuth, datadog.MetricMetadata{}, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MetricsApi.UpdateMetricMetadata(ctx, "ametric").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMetricsList400Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Error 400 cannot be triggered from the client due to client side validation, so mock it
	res, err := tests.ReadFixture("fixtures/metrics/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/search").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MetricsApi.ListMetrics(ctx).Q("").Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMetricsListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// Error 400 cannot be triggered from the client due to client side validation
		// "400 Bad Request": {WithTestAuth,400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MetricsApi.ListMetrics(ctx).Q("somequery").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestMetricsQuery400Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Error 400 cannot be triggered from the client due to client side validation, so mock it
	res, err := tests.ReadFixture("fixtures/metrics/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/query").Reply(400).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).MetricsApi.QueryMetrics(ctx).Query("").From(0).To(0).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestMetricsQueryErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// Error 400 cannot be triggered from the client due to client side validation
		// "400 Bad Request": {WithTestAuth,400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).MetricsApi.QueryMetrics(ctx).Query("somequery").From(0).To(0).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
