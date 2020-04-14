/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
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
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageFargate(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Fargate: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageHosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageHosts(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLogs(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLogs(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSynthetics(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSynthetics(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSyntheticsAPI(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSyntheticsAPI(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics API: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSyntheticsBrowser(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSyntheticsBrowser(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics Browser: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTimeseries(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTimeseries(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Timeseries: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTopAvgMetrics(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTopAvgMetrics(TESTAUTH).Month(TESTCLOCK.Now()).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Avg Metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTrace(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTrace(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Trace: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLogsByIndex(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLogsByIndex(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs by Index: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLambda(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLambda(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Lambda: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageNetworkHosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageNetworkHosts(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageNetworkFlows(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageNetworkFlows(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Flows: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageRumSessions(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	startHr, endHr := getStartEndHr()
	usage, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageRumSessions(TESTAUTH).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
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

	api := TESTAPICLIENT.UsageMeteringApi
	usage, httpresp, err := api.GetUsageSummary(TESTAUTH).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()
	if err != nil {
		t.Errorf("Failed to get Usage Summary: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC), usage.GetStartDate().UTC())
	assert.Equal(t, time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC), usage.GetEndDate().UTC())
	assert.Equal(t, int64(1), usage.GetApmHostTop99pSum())
	assert.Equal(t, int64(2), usage.GetInfraHostTop99pSum())
	assert.Equal(t, int64(1), usage.GetContainerHwmSum())
	assert.Equal(t, int64(4), usage.GetCustomTsSum())
	assert.Equal(t, int64(5), usage.GetRumSessionCountAggSum())

	var usageItem = usage.GetUsage()[0]
	assert.Equal(t, time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC), usageItem.GetDate().UTC())
	assert.Equal(t, int64(1), usageItem.GetAgentHostTop99p())
	assert.Equal(t, int64(2), usageItem.GetApmHostTop99p())
	assert.Equal(t, int64(3), usageItem.GetAwsHostTop99p())
	assert.Equal(t, int64(5), usageItem.GetContainerHwm())
	assert.Equal(t, int64(6), usageItem.GetCustomTsAvg())
	assert.Equal(t, int64(7), usageItem.GetGcpHostTop99p())
	assert.Equal(t, int64(8), usageItem.GetInfraHostTop99p())
	assert.Equal(t, int64(9), usageItem.GetRumSessionCountSum())

	var usageOrgItem = usageItem.GetOrgs()[0]
	assert.Equal(t, "1b", usageOrgItem.GetId())
	assert.Equal(t, "datadog", usageOrgItem.GetName())
	assert.Equal(t, int64(1), usageOrgItem.GetAgentHostTop99p())
	assert.Equal(t, int64(2), usageOrgItem.GetApmHostTop99p())
	assert.Equal(t, int64(3), usageOrgItem.GetAwsHostTop99p())
	assert.Equal(t, int64(5), usageOrgItem.GetContainerHwm())
	assert.Equal(t, int64(6), usageOrgItem.GetCustomTsAvg())
	assert.Equal(t, int64(7), usageOrgItem.GetGcpHostTop99p())
	assert.Equal(t, int64(8), usageOrgItem.GetInfraHostTop99p())
	assert.Equal(t, int64(9), usageOrgItem.GetRumSessionCountSum())
}

func TestUsageGetHostsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageHosts(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLogs(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsByIndexErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLogsByIndex(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageTimeSeriesErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTimeseries(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageTopAvgMetricsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTopAvgMetrics(tc.Ctx).Month(TESTCLOCK.Now().AddDate(-2, 0, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageTraceErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageTrace(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSynthetics(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsAPIErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSyntheticsAPI(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsBrowserErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSyntheticsBrowser(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageFargateErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageFargate(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageLambdaErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageLambda(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageRumSessionErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageRumSessions(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageNetworkHostsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageNetworkHosts(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageNetworkFlowsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageNetworkFlows(tc.Ctx).StartHr(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageSummaryErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		// Mocked because requires multi org feature
		// {"400 Bad Request", TESTAUTH, 400},
		{"403 Forbidden", fake_auth, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSummary(tc.Ctx).StartMonth(TESTCLOCK.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageGetSummary400Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/usage/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/usage/summary").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := TESTAPICLIENT.UsageMeteringApi.GetUsageSummary(TESTAUTH).StartMonth(time.Now()).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}
