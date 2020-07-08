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
)

func getStartEndHr(ctx context.Context) (time.Time, time.Time) {
	year, month, _ := tests.ClockFromContext(ctx).Now().Date()
	start := time.Date(year, month-1, 12, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 13, 0, 0, 0, 0, time.UTC)
	return start, end
}

func getStartEndMonth(ctx context.Context) (time.Time, time.Time) {
	year, month, _ := tests.ClockFromContext(ctx).Now().Date()
	start := time.Date(year, month-2, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 1, 0, 0, 0, 0, time.UTC)
	return start, end
}

func TestUsageAnalyzedLogs(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageAnalyzedLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Analyzed Logs (Security Monitoring): Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageFargate(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageFargate(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Fargate: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageHosts(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLogs(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSynthetics(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSynthetics(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSyntheticsAPI(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSyntheticsAPI(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics API: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSyntheticsBrowser(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSyntheticsBrowser(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics Browser: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageTimeseries(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTimeseries(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Timeseries: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageTopAvgMetrics(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTopAvgMetrics(ctx).Month(tests.ClockFromContext(ctx).Now()).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Avg Metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageTrace(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTrace(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Trace: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLogsByIndex(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLogsByIndex(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs by Index: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLambda(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLambda(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Lambda: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageNetworkHosts(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageNetworkHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageNetworkFlows(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageNetworkFlows(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Flows: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageRumSessions(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageRumSessions(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSNMP(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSNMP(ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage SNMP Devices: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageBillableSummary(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// startMonth := time.Date(2020, 06, 01, 0, 0, 0, 0, time.UTC)
	// endMonth := time.Date(2020, 06, 28, 23, 0, 0, 0, time.UTC)

	fixturePath, err := filepath.Abs("fixtures/usage/usage_billable_summary.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	// The query params get normalized via autogenerated code, so just confirm that the keys exist in the request
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "UsageMeteringApiService.GetUsageBillableSummary")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/usage/billable-summary").
		Reply(200).
		JSON(data)
	defer gock.Off()

	var expected datadog.UsageBillableSummaryResponse
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).UsageMeteringApi
	usage, httpresp, err := api.GetUsageBillableSummary(ctx).Execute()
	if err != nil {
		t.Errorf("Failed to get Billable Usage Summary: %v", err)
	}

	assert.Equal(200, httpresp.StatusCode)
	var usageItem = usage.GetUsage()[0]
	assert.Equal("API - Test", usageItem.GetOrgName())
	assert.Equal("Pro", usageItem.GetBillingPlan())
	assert.Equal("123abcxyz", usageItem.GetPublicId())
	assert.Equal(time.Date(2020, 06, 01, 00, 0, 0, 0, time.UTC), usageItem.GetStartDate().UTC())
	assert.Equal(time.Date(2020, 06, 28, 23, 0, 0, 0, time.UTC), usageItem.GetEndDate().UTC())
	assert.Equal(int64(1), usageItem.GetRatioInMonth())
	assert.Equal(int64(235), usageItem.GetNumOrgs())

	var usageUsageItem = usageItem.GetUsage()
	var usageKeys = usageUsageItem.GetLogsIndexedSum()

	assert.Equal(int64(14514687), usageKeys.GetOrgBillableUsage())
	assert.Equal("logs", usageKeys.GetUsageUnit())
	assert.Equal(int64(1611132837), usageKeys.GetAccountBillableUsage())
	assert.Equal(time.Date(2020, 06, 01, 0, 0, 0, 0, time.UTC), usageKeys.GetFirstBillableUsageHour().UTC())
	assert.Equal(int64(672), usageKeys.GetElapsedUsageHours())
	assert.Equal(time.Date(2020, 06, 28, 23, 0, 0, 0, time.UTC), usageKeys.GetLastBillableUsageHour().UTC())
	assert.Equal(float64(0.9), usageKeys.GetPercentageInAccount())
}

func TestSpecifiedDailyCustomReports(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	reportID := "2019-10-02"

	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetSpecifiedDailyCustomReports(ctx, reportID).Execute()
	if err != nil {
		t.Errorf("Error getting Specified Daily Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestSpecifiedMonthlyCustomReports(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	reportID := "2019-10-02"

	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetSpecifiedMonthlyCustomReports(ctx, reportID).Execute()
	if err != nil {
		t.Errorf("Error getting Specified Monthly Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestDailyCustomReports(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetDailyCustomReports(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting Daily Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestMonthlyCustomReports(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	usage, httpresp, err := Client(ctx).UsageMeteringApi.GetMonthlyCustomReports(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting Monthly Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

// This test needs multi-org token so make it a unit test
func TestUsageSummary(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

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
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "UsageMeteringApiService.GetUsageSummary")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/usage/summary").
		MatchParam("include_org_details", strconv.FormatBool(includeOrgDetails)).
		ParamPresent("end_month").
		ParamPresent("start_month").
		Reply(200).
		JSON(data)
	defer gock.Off()

	var expected datadog.UsageSummaryResponse
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).UsageMeteringApi
	usage, httpresp, err := api.GetUsageSummary(ctx).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()
	if err != nil {
		t.Errorf("Failed to get Usage Summary: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC), usage.GetStartDate().UTC())
	assert.Equal(time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC), usage.GetEndDate().UTC())
	assert.Equal(int64(1), usage.GetApmHostTop99pSum())
	assert.Equal(int64(2), usage.GetInfraHostTop99pSum())
	assert.Equal(int64(1), usage.GetContainerHwmSum())
	assert.Equal(int64(4), usage.GetCustomTsSum())
	assert.Equal(int64(5), usage.GetRumSessionCountAggSum())

	var usageItem = usage.GetUsage()[0]
	assert.Equal(time.Date(2020, 02, 02, 23, 0, 0, 0, time.UTC), usageItem.GetDate().UTC())
	assert.Equal(int64(1), usageItem.GetAgentHostTop99p())
	assert.Equal(int64(2), usageItem.GetApmHostTop99p())
	assert.Equal(int64(3), usageItem.GetAwsHostTop99p())
	assert.Equal(int64(5), usageItem.GetContainerHwm())
	assert.Equal(int64(6), usageItem.GetCustomTsAvg())
	assert.Equal(int64(7), usageItem.GetGcpHostTop99p())
	assert.Equal(int64(8), usageItem.GetInfraHostTop99p())
	assert.Equal(int64(9), usageItem.GetRumSessionCountSum())

	var usageOrgItem = usageItem.GetOrgs()[0]
	assert.Equal("1b", usageOrgItem.GetId())
	assert.Equal("datadog", usageOrgItem.GetName())
	assert.Equal(int64(1), usageOrgItem.GetAgentHostTop99p())
	assert.Equal(int64(2), usageOrgItem.GetApmHostTop99p())
	assert.Equal(int64(3), usageOrgItem.GetAwsHostTop99p())
	assert.Equal(int64(5), usageOrgItem.GetContainerHwm())
	assert.Equal(int64(6), usageOrgItem.GetCustomTsAvg())
	assert.Equal(int64(7), usageOrgItem.GetGcpHostTop99p())
	assert.Equal(int64(8), usageOrgItem.GetInfraHostTop99p())
	assert.Equal(int64(9), usageOrgItem.GetRumSessionCountSum())
}

func TestUsageGetAnalyzedLogsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageAnalyzedLogs(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetHostsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageHosts(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLogs(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsByIndexErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLogsByIndex(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetSNMPErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSNMP(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetBillableSummaryErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		//"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageBillableSummary(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetSpecifiedDailyCustomReportsErrors(t *testing.T) {

	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden":   {WithFakeAuth, 403},
		"502 Bad Gateway": {WithTestAuth, 502},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetSpecifiedDailyCustomReports(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetSpecifiedMonthlyCustomReportsErrors(t *testing.T) {

	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden":   {WithFakeAuth, 403},
		"400 Bad Request": {WithTestAuth, 400},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetSpecifiedMonthlyCustomReports(ctx, "whatever").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetDailyCustomReportsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// "400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetDailyCustomReports(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetDailyCustomReports400Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/usage/no_authenticated_user_error.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/daily_custom_reports").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).UsageMeteringApi.GetDailyCustomReports(ctx).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestGetMonthlyCustomReportsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// "400 Bad Request": {WithTestAuth, 404},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetMonthlyCustomReports(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetMonthlyCustomReports400Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/usage/no_authenticated_user_error.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/monthly_custom_reports").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).UsageMeteringApi.GetMonthlyCustomReports(ctx).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUsageGetBillableSummary400Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/usage/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked as this call must be made from the parent organization
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/usage/billable-summary").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageBillableSummary(ctx).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUsageTimeSeriesErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTimeseries(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageTopAvgMetricsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTopAvgMetrics(ctx).Month(tests.ClockFromContext(ctx).Now().AddDate(-2, 0, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageTraceErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageTrace(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSynthetics(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsAPIErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSyntheticsAPI(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsBrowserErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSyntheticsBrowser(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageFargateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageFargate(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageLambdaErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageLambda(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageRumSessionErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageRumSessions(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageNetworkHostsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageNetworkHosts(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageNetworkFlowsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageNetworkFlows(ctx).StartHr(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSummaryErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		// Mocked because requires multi org feature
		// "400 Bad Request": {WithTestAuth,400},
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSummary(ctx).StartMonth(tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetSummary400Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/usage/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/usage/summary").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := Client(ctx).UsageMeteringApi.GetUsageSummary(ctx).StartMonth(time.Now()).Execute()
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}
