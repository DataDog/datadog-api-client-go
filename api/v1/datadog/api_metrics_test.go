package datadog_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestMetricSubmissionMock(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()

	// Test that a normal submission works
	t.Run("normalSubmission", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		testType := "count"
		testInterval := datadog.NullableInt32{Value: 20}
		testMetric := "hello.world"
		testPoints := [][]float64{{5, 10.5}}
		metricsPayload := datadog.MetricsPayload{
			Series: &[]datadog.Series{
				datadog.Series{
					Host:     &testHost,
					Type:     &testType,
					Interval: &testInterval,
					Metric:   testMetric,
					Points:   testPoints,
					Tags:     &testTags,
				},
			},
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// Read request
				defer req.Body.Close()
				body, err := ioutil.ReadAll(req.Body)
				assert.NoError(t, err)
				var payload datadog.MetricsPayload
				json.Unmarshal(body, &payload)

				// Check equality
				assert.Equal(t, *payload.GetSeries()[0].Host, testHost)
				assert.Equal(t, *payload.GetSeries()[0].Type, testType)
				assert.Equal(t, payload.GetSeries()[0].Interval.Value, testInterval.Value)
				assert.Equal(t, payload.GetSeries()[0].Metric, testMetric)
				assert.Equal(t, payload.GetSeries()[0].Points, testPoints)
				assert.Equal(t, *payload.GetSeries()[0].Tags, testTags)

				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).MetricsPayload(metricsPayload).Execute()
		cancel()
		assert.NoError(t, err)
	})

	// Test that we can send multiple points
	t.Run("multiplePoints", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		metricsPayload := datadog.MetricsPayload{
			Series: &[]datadog.Series{
				datadog.Series{
					Host:   &testHost,
					Metric: "hello.world",
					Points: [][]float64{{5, 10.5}, {6, 11}},
					Tags:   &testTags,
				},
			},
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// Read request
				defer req.Body.Close()
				body, err := ioutil.ReadAll(req.Body)
				assert.NoError(t, err)
				var payload datadog.MetricsPayload
				json.Unmarshal(body, &payload)

				// Check equality
				assert.Equal(t, *payload.GetSeries()[0].Host, testHost)
				assert.Equal(t, payload.GetSeries()[0].Metric, "hello.world")
				assert.Equal(t, payload.GetSeries()[0].Points, [][]float64{{5, 10.5}, {6, 11}})
				assert.Equal(t, *payload.GetSeries()[0].Tags, testTags)

				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).MetricsPayload(metricsPayload).Execute()
		cancel()
		assert.NoError(t, err)
	})
}

func TestMetricSubmissionIntegration(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	api := TESTAPICLIENT.MetricsApi
	now := time.Now().Unix()

	testMetric := fmt.Sprintf("go.client.test.%d", now)
	testPoints := [][]float64{{float64(now - 60), 10.5}, {float64(now), 11}}
	metricsPayload := datadog.MetricsPayload{
		Series: &[]datadog.Series{
			datadog.Series{
				Metric: testMetric,
				Points: testPoints,
			},
		},
	}

	r, httpresp, err := api.SubmitMetrics(TESTAUTH).MetricsPayload(metricsPayload).Execute()
	if err != nil {
		t.Fatalf("Error submitting metric %v: Response %s: %v", metricsPayload, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 202)
	assert.Equal(t, "ok", r.GetStatus())

	err = retry(func() bool {
		metrics, httpresp, err := api.GetAllActiveMetrics(TESTAUTH).From(now).Execute()
		if err != nil {
			t.Logf("Error getting list of active metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
			return false
		}
		assert.Equal(t, 200, httpresp.StatusCode)

		found := false
		for _, metric := range metrics.GetMetrics() {
			if metric == testMetric {
				found = true
				break
			}
		}
		return found
	}, 10, 10)

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestMetricListActive(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	data := setupGock(t, "metrics/active_metrics.json", "GET", "metrics")
	var expected datadog.MetricsListResponse
	json.Unmarshal([]byte(data), &expected)

	api := TESTAPICLIENT.MetricsApi

	_, _, err := api.GetAllActiveMetrics(TESTAUTH).Execute()
	assert.NotNil(t, err) // Parameter `from` required

	metrics, _, err := api.GetAllActiveMetrics(TESTAUTH).From(1).Host("host").Execute()
	assert.Nil(t, err)
	assert.Equal(t, expected.GetMetrics(), metrics.GetMetrics())
	assert.Equal(t, expected.GetFrom(), metrics.GetFrom())
}
