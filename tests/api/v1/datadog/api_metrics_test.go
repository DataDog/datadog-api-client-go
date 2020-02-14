// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	api := TESTAPICLIENT.MetricsApi
	now := time.Now().Unix()

	testMetric := fmt.Sprintf("go.client.test.%d", now)
	testPoints := [][]float64{{float64(now - 60), 10.5}, {float64(now), 11}}
	testTags := []string{"tag:foo", "bar:baz"}
	testHost := "go-client-test-host"
	testQuery := fmt.Sprintf("avg:%s{bar:baz}by{host}", testMetric)
	metricsPayload := datadog.MetricsPayload{
		Series: &[]datadog.Series{
			datadog.Series{
				Host:   &testHost,
				Metric: testMetric,
				Points: testPoints,
				Tags:   &testTags,
			},
		},
	}

	r, httpresp, err := api.SubmitMetrics(TESTAUTH).Body(metricsPayload).Execute()
	if err != nil {
		t.Fatalf("Error submitting metric %v: Response %s: %v", metricsPayload, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 202)
	assert.Equal(t, "ok", r.GetStatus())

	// Check that the metric was submitted successfully
	err = tests.Retry(10*time.Second, 10, func() bool {
		metrics, httpresp, err := api.GetAllActiveMetrics(TESTAUTH).From(now).Execute()
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
	queryResult, httpresp, err := api.QueryMetrics(TESTAUTH).From(now - 100).To(now + 100).Query(testQuery).Execute()
	if err != nil {
		t.Errorf("Error making query %s: Response %s: %v", testQuery, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, []string{"host"}, queryResult.GetGroupBy())
	assert.Equal(t, testQuery, queryResult.GetQuery())
	assert.Equal(t, (now-100)*1000, queryResult.GetFromDate())
	assert.Equal(t, (now+100)*1000, queryResult.GetToDate())
	assert.Equal(t, "ok", queryResult.GetStatus())
	assert.Equal(t, "time_series", queryResult.GetResType())
	assert.Equal(t, 1, len(queryResult.GetSeries()))
	series := queryResult.GetSeries()[0]
	assert.Equal(t, int64(2), series.GetLength())
	assert.Equal(t, "avg", series.GetAggr())
	assert.Equal(t, testMetric, series.GetDisplayName())
	assert.Equal(t, testMetric, series.GetMetric())
	assert.Equal(t, series.GetPointlist()[0][0], float64(series.GetStart()))
	assert.Equal(t, series.GetPointlist()[1][0], float64(series.GetEnd()))
	assert.Equal(t, 10.5, series.GetPointlist()[0][1])
	assert.Equal(t, 11., series.GetPointlist()[1][1])

	// Test search
	searchQuery := fmt.Sprintf("metrics:%s", testMetric)
	searchResult, httpresp, err := api.SearchMetrics(TESTAUTH).Q(searchQuery).Execute()
	if err != nil {
		t.Errorf("Error searching metrics %s: Response %s: %v", searchQuery, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	metrics := searchResult.Results.GetMetrics()
	assert.Equal(t, 1, len(metrics))
	assert.Equal(t, testMetric, metrics[0])

	// Test metric metadata
	metadata, httpresp, err := api.GetMetricMetadata(TESTAUTH, testMetric).Execute()
	if err != nil {
		t.Errorf("Error getting metric metadata for %s: Response %s: %v", testMetric, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Nil(t, metadata.Description)
	assert.Nil(t, metadata.Integration)
	assert.Nil(t, metadata.PerUnit)
	assert.Nil(t, metadata.Unit)
	assert.Nil(t, metadata.ShortName)
	assert.Nil(t, metadata.StatsdInterval)
	assert.Equal(t, "", metadata.Type)

	newMetadata := datadog.MetricMetadata{
		Description:    datadog.PtrString("description"),
		PerUnit:        datadog.PtrString("second"),
		Unit:           datadog.PtrString("byte"),
		ShortName:      datadog.PtrString("short_name"),
		StatsdInterval: datadog.PtrInt64(20),
		Type:           "count",
	}

	metadata, httpresp, err = api.EditMetricMetadata(TESTAUTH, testMetric).Body(newMetadata).Execute()
	if err != nil {
		t.Errorf("Error editing metric metadata for %s: Response %s: %v", testMetric, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, "description", metadata.GetDescription())
	assert.Nil(t, metadata.Integration)
	assert.Equal(t, "second", metadata.GetPerUnit())
	assert.Equal(t, "byte", metadata.GetUnit())
	assert.Equal(t, "short_name", metadata.GetShortName())
	assert.Equal(t, int64(20), metadata.GetStatsdInterval())
	assert.Equal(t, "count", metadata.GetType())
}
