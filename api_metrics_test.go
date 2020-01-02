package datadog_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestMetricSubmission(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
	defer gock.Off()
	gock.Observe(gock.DumpRequest) // ToDo: Remove

	// Test that a normal submission works
	t.Run("normalSubmission", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		reqSeries := datadog.Series{
			Host:   &testHost,
			Metric: "hello.world",
			Points: [][]float64{{5, 10.5}},
			Tags:   &testTags,
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// Read request
				defer req.Body.Close()
				body, err := ioutil.ReadAll(req.Body)
				assert.NoError(t, err)
				var series datadog.Series
				json.Unmarshal(body, &series)

				// Check equality
				assert.Equal(t, *series.Host, testHost)
				assert.Equal(t, series.Metric, "hello.world")
				assert.Equal(t, series.Points, [][]float64{{5, 10.5}})
				assert.Equal(t, *series.Tags, testTags)

				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).Series(reqSeries).Execute()
		cancel()
		assert.NoError(t, err)
	})

	// Test that metric name is mandatory
	t.Run("metricRequired", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		reqSeries := datadog.Series{
			Host:   &testHost,
			Points: [][]float64{{5, 10.5}},
			Tags:   &testTags,
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// The call should not come to us since the metric is missing
				t.Error("We should not have received a call with a missing metric name")
				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).Series(reqSeries).Execute()
		cancel()
		assert.NoError(t, err)
	})

	// Test that a point is mandatory
	t.Run("pointsRequired", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		reqSeries := datadog.Series{
			Host:   &testHost,
			Metric: "hello.world",
			Tags:   &testTags,
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// The call should not come to us since the metric is missing
				t.Error("We should not have received a call with no points")
				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).Series(reqSeries).Execute()
		cancel()
		assert.NoError(t, err)
	})

	// Test that we can send multiple points
	t.Run("multiplePoints", func(t *testing.T) {
		testHost := "test"
		testTags := []string{"tagA", "tagB"}
		reqSeries := datadog.Series{
			Host:   &testHost,
			Metric: "hello.world",
			Points: [][]float64{{5, 10.5}, {6, 11}},
			Tags:   &testTags,
		}
		// Check that request is correct
		gock.Clean()
		gock.New("https://api.datadoghq.com/api/v1").Post("/series").
			AddMatcher(func(req *http.Request, _ *gock.Request) (bool, error) {
				// Read request
				defer req.Body.Close()
				body, err := ioutil.ReadAll(req.Body)
				assert.NoError(t, err)
				var series datadog.Series
				json.Unmarshal(body, &series)

				// Check equality
				assert.Equal(t, *series.Host, testHost)
				assert.Equal(t, series.Metric, "hello.world")
				assert.Equal(t, series.Points, [][]float64{{5, 10.5}, {6, 11}})
				assert.Equal(t, *series.Tags, testTags)

				return true, nil
			}).Reply(200).JSON(map[string]string{"status": "ok"})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, _, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(ctx).Series(reqSeries).Execute()
		cancel()
		assert.NoError(t, err)
	})

}
