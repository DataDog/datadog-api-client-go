/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"gopkg.in/h2non/gock.v1"

	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func getStartEndHr() (time.Time, time.Time) {
	year, month, _ := TESTCLOCK.Now().Date()
	start := time.Date(year, month-1, 12, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 13, 0, 0, 0, 0, time.UTC)
	return start, end
}

func getStartEndMonth() (time.Time, time.Time) {
	year, month, _ := TESTCLOCK.Now().Date()
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

	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageTopAvgMetrics(TESTAUTH).Month(TESTCLOCK.Now()).Execute()
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

func TestUsageRumSessions(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageApi.GetUsageRumSessions(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	assert.Equal(t, usage.HasUsage(), true)
}

// This test needs multi-org token so make it a unit test
func TestUsageSummary(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	startMonth := time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC)
	endMonth := time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC)
	includeOrgDetails := true

	fixturePath, err := filepath.Abs("fixtures/usage/usage_summary.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	// The query params get normalized via autogenerated code, so just confirm that the keys exist in the request
	gock.New("https://api.datadoghq.com/api/v1").
		Get("/usage/summary").
		MatchParam("include_org_details", strconv.FormatBool(includeOrgDetails)).
		ParamPresent("end_month").
		ParamPresent("start_month").
		Reply(200).
		JSON(data)
	defer gock.Off()

	var expected datadog.UsageSummaryResponse
	json.Unmarshal([]byte(data), &expected)

	api := TESTAPICLIENT.UsageApi
	usage, httpresp, err := api.GetUsageSummary(TESTAUTH).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()
	if err != nil {
		t.Errorf("Failed to get Usage Summary: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.DeepEqual(t, usage.GetStartDate(), time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC))
	assert.DeepEqual(t, usage.GetEndDate(), time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, usage.GetApmHostTop99pSum(), int64(1))
	assert.Equal(t, usage.GetInfraHostTop99pSum(), int64(2))
	assert.Equal(t, usage.GetContainerHwmSum(), int64(1))
	assert.Equal(t, usage.GetCustomTsSum(), int64(4))

	var usageItem = usage.GetUsage()[0]
	assert.DeepEqual(t, usageItem.GetDate(), time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, usageItem.GetAgentHostTop99p(), int64(1))
	assert.Equal(t, usageItem.GetApmHostTop99p(), int64(2))
	assert.Equal(t, usageItem.GetAwsHostTop99p(), int64(3))
	assert.Equal(t, usageItem.GetContainerHwm(), int64(5))
	assert.Equal(t, usageItem.GetCustomTsAvg(), int64(6))
	assert.Equal(t, usageItem.GetGcpHostTop99p(), int64(7))
	assert.Equal(t, usageItem.GetInfraHostTop99p(), int64(8))

	var usageOrgItem = usageItem.GetOrgs()[0]
	assert.Equal(t, usageOrgItem.GetId(), "1b")
	assert.Equal(t, usageOrgItem.GetName(), "datadog")
	assert.Equal(t, usageOrgItem.GetAgentHostTop99p(), int64(1))
	assert.Equal(t, usageOrgItem.GetApmHostTop99p(), int64(2))
	assert.Equal(t, usageOrgItem.GetAwsHostTop99p(), int64(3))
	assert.Equal(t, usageOrgItem.GetContainerHwm(), int64(5))
	assert.Equal(t, usageOrgItem.GetCustomTsAvg(), int64(6))
	assert.Equal(t, usageOrgItem.GetGcpHostTop99p(), int64(7))
	assert.Equal(t, usageOrgItem.GetInfraHostTop99p(), int64(8))
}
