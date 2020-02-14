// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func getStartEndHr() (time.Time, time.Time) {
	year, month, _ := time.Now().Date()
	start := time.Date(year, month-1, 12, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 13, 0, 0, 0, 0, time.UTC)
	return start, end
}

func getStartEndMonth() (time.Time, time.Time) {
	year, month, _ := time.Now().Date()
	start := time.Date(year, month-2, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 1, 0, 0, 0, 0, time.UTC)
	return start, end
}

func TestUsageFargate(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageFargate(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Fargate: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

func TestUsageHosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageHosts(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

func TestUsageLogs(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageLogs(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

/* FIXME needs multi-org token
func TestUsageSummary(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startMonth, endMonth := getStartEndMonth()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageSummary(TESTAUTH).StartMonth(startMonth).EndMonth(endMonth).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Summary: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}
*/

func TestUsageSynthetics(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageSynthetics(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

func TestUsageTimeseries(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageTimeseries(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Timeseries: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

func TestUsageTopAvgMetrics(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageTopAvgMetrics(TESTAUTH).Month(time.Now()).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Avg Metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

func TestUsageTrace(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageTrace(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Trace: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}
