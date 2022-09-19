/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"

	"gopkg.in/h2non/gock.v1"
)

func getStartEndHr(ctx context.Context) (time.Time, time.Time) {
	year, month, _ := tests.ClockFromContext(ctx).Now().Date()
	start := time.Date(year, month-1, 12, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 13, 0, 0, 0, 0, time.UTC)
	return start, end
}

func TestUsageAnalyzedLogs(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageAnalyzedLogs(ctx, startHr, *datadogV1.NewGetUsageAnalyzedLogsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Analyzed Logs (Security Monitoring): Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageFargate(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageFargate(ctx, startHr, *datadogV1.NewGetUsageFargateOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Fargate: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageHosts(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageHosts(ctx, startHr, *datadogV1.NewGetUsageHostsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLogs(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageLogs(ctx, startHr, *datadogV1.NewGetUsageLogsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSynthetics(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageSynthetics(ctx, startHr, *datadogV1.NewGetUsageSyntheticsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Synthetics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSyntheticsAPI(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageSyntheticsAPI(ctx, startHr, *datadogV1.NewGetUsageSyntheticsAPIOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Synthetics API: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSyntheticsBrowser(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageSyntheticsBrowser(ctx, startHr, *datadogV1.NewGetUsageSyntheticsBrowserOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Synthetics Browser: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageTimeseries(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageTimeseries(ctx, startHr, *datadogV1.NewGetUsageTimeseriesOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Timeseries: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageTopAvgMetrics(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	usage, httpresp, err := api.GetUsageTopAvgMetrics(ctx, *datadogV1.NewGetUsageTopAvgMetricsOptionalParameters().WithMonth(tests.ClockFromContext(ctx).Now()))
	if err != nil {
		t.Errorf("Error getting Usage Avg Metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageIndexedSpans(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageIndexedSpans(ctx, startHr, *datadogV1.NewGetUsageIndexedSpansOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Indexed Spans: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLogsByIndex(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageLogsByIndex(ctx, startHr, *datadogV1.NewGetUsageLogsByIndexOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Logs by Index: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLambda(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageLambda(ctx, startHr, *datadogV1.NewGetUsageLambdaOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Lambda: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageNetworkHosts(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageNetworkHosts(ctx, startHr, *datadogV1.NewGetUsageNetworkHostsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Network Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageNetworkFlows(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageNetworkFlows(ctx, startHr, *datadogV1.NewGetUsageNetworkFlowsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Network Flows: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageRumSessions(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageRumSessions(ctx, startHr, *datadogV1.NewGetUsageRumSessionsOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageMobileRumSessions(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageRumSessions(ctx, startHr, *datadogV1.NewGetUsageRumSessionsOptionalParameters().WithEndHr(endHr).WithType("mobile"))
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageSNMP(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	usage, httpresp, err := api.GetUsageSNMP(ctx, startHr, *datadogV1.NewGetUsageSNMPOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage SNMP Devices: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageProfiling(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)

	usage, httpresp, err := api.GetUsageProfiling(ctx, startHr, *datadogV1.NewGetUsageProfilingOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageIngestedSpans(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)

	usage, httpresp, err := api.GetIngestedSpans(ctx, startHr, *datadogV1.NewGetIngestedSpansOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageIncidentManagement(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)

	usage, httpresp, err := api.GetIncidentManagement(ctx, startHr, *datadogV1.NewGetIncidentManagementOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageLogsByRetention(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)

	usage, httpresp, err := api.GetUsageLogsByRetention(ctx, startHr, *datadogV1.NewGetUsageLogsByRetentionOptionalParameters().WithEndHr(endHr))
	if err != nil {
		t.Errorf("Error getting logs usage by retention: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
}

func TestUsageBillableSummary(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// startMonth := time.Date(2020, 06, 01, 0, 0, 0, 0, time.UTC)
	// endMonth := time.Date(2020, 06, 28, 23, 0, 0, 0, time.UTC)

	fixturePath, err := filepath.Abs("fixtures/usage/usage_billable_summary.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := os.ReadFile(fixturePath)
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

	var expected datadogV1.UsageBillableSummaryResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewUsageMeteringApi(Client(ctx))
	usage, httpresp, err := api.GetUsageBillableSummary(ctx)
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
	assert.Equal(float64(1), usageItem.GetRatioInMonth())
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	currentDate := tests.ClockFromContext(ctx).Now().AddDate(0, 0, -1) // only have reports from day before
	reportID := currentDate.Format("2006-01-02")                       // this is only a date format example YYYY-MM-DD

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSpecifiedDailyCustomReports", true)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))
	usage, httpresp, err := api.GetSpecifiedDailyCustomReports(ctx, reportID)
	if err != nil {
		if tests.GetRecording() != tests.ModeReplaying || httpresp.StatusCode == 404 || httpresp.StatusCode == 403 {
			t.Skip("No reports are available yet or this org is forbidden")
		} else {
			t.Errorf("Error getting Specified Daily Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestSpecifiedMonthlyCustomReports(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	currentDate := tests.ClockFromContext(ctx).Now().AddDate(0, 0, -1) // We will only have monthly reports on 2020-08-15 for this org
	reportID := currentDate.Format("2006-01-02")                       // this is only a date format example YYYY-MM-DD

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSpecifiedMonthlyCustomReports", true)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))
	usage, httpresp, err := api.GetSpecifiedMonthlyCustomReports(ctx, reportID)
	if err != nil {
		if tests.GetRecording() != tests.ModeReplaying || httpresp.StatusCode == 404 || httpresp.StatusCode == 403 {
			t.Skip("No reports are available yet or this org is forbidden")
		} else {
			t.Errorf("Error getting Specified Monthly Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestDailyCustomReports(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetDailyCustomReports", true)
	usage, httpresp, err := api.GetDailyCustomReports(ctx)
	if err != nil {
		if tests.GetRecording() != tests.ModeReplaying || httpresp.StatusCode == 404 || httpresp.StatusCode == 403 {
			t.Skip("No reports are available yet or this org is forbidden")
		} else {
			t.Errorf("Error getting Daily Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

func TestMonthlyCustomReports(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetMonthlyCustomReports", true)
	usage, httpresp, err := api.GetMonthlyCustomReports(ctx)
	if err != nil {
		if tests.GetRecording() != tests.ModeReplaying || httpresp.StatusCode == 404 || httpresp.StatusCode == 403 {
			t.Skip("No reports are available yet or this org is forbidden ")
		} else {
			t.Errorf("Error getting Monthly Custom Reports Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasMeta())
	assert.True(usage.HasData())
}

// This test needs multi-org token so make it a unit test
func TestUsageSummary(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	startMonth := time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC)
	endMonth := time.Date(2019, 02, 02, 23, 0, 0, 0, time.UTC)
	includeOrgDetails := true

	fixturePath, err := filepath.Abs("fixtures/usage/usage_summary.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := os.ReadFile(fixturePath)
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

	var expected datadogV1.UsageSummaryResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewUsageMeteringApi(Client(ctx))
	usage, httpresp, err := api.GetUsageSummary(ctx, startMonth, *datadogV1.NewGetUsageSummaryOptionalParameters().WithEndMonth(endMonth).WithIncludeOrgDetails(includeOrgDetails))
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
	assert.Equal(int64(6), usage.GetProfilingHostCountTop99pSum())
	assert.Equal(int64(7), usage.GetProfilingContainerAgentCountAvg())
	assert.Equal(int64(8), usage.GetTwolIngestedEventsBytesAggSum())
	assert.Equal(int64(9), usage.GetMobileRumSessionCountAggSum())
	assert.Equal(int64(10), usage.GetIncidentManagementMonthlyActiveUsersHwmSum())
	assert.Equal(int64(11), usage.GetMobileRumSessionCountIosAggSum())
	assert.Equal(int64(12), usage.GetMobileRumSessionCountAndroidAggSum())
	assert.Equal(int64(13), usage.GetRumTotalSessionCountAggSum())
	assert.Equal(int64(14), usage.GetAzureAppServiceTop99pSum())
	assert.Equal(int64(15), usage.GetApmAzureAppServiceHostTop99pSum())

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
	assert.Equal(int64(10), usageItem.GetProfilingHostTop99p())
	assert.Equal(int64(11), usageItem.GetTwolIngestedEventsBytesSum())
	assert.Equal(int64(12), usageItem.GetMobileRumSessionCountSum())
	assert.Equal(int64(13), usageItem.GetIncidentManagementMonthlyActiveUsersHwm())
	assert.Equal(int64(14), usageItem.GetMobileRumSessionCountIosSum())
	assert.Equal(int64(15), usageItem.GetMobileRumSessionCountAndroidSum())
	assert.Equal(int64(16), usageItem.GetRumTotalSessionCountSum())
	assert.Equal(int64(17), usageItem.GetAzureAppServiceTop99p())
	assert.Equal(int64(18), usageItem.GetApmAzureAppServiceHostTop99p())

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
	assert.Equal(int64(10), usageOrgItem.GetProfilingHostTop99p())
	assert.Equal(int64(11), usageOrgItem.GetTwolIngestedEventsBytesSum())
	assert.Equal(int64(12), usageOrgItem.GetMobileRumSessionCountSum())
	assert.Equal(int64(13), usageOrgItem.GetIncidentManagementMonthlyActiveUsersHwm())
	assert.Equal(int64(14), usageItem.GetMobileRumSessionCountIosSum())
	assert.Equal(int64(15), usageItem.GetMobileRumSessionCountAndroidSum())
	assert.Equal(int64(16), usageItem.GetRumTotalSessionCountSum())
	assert.Equal(int64(17), usageOrgItem.GetAzureAppServiceTop99p())
	assert.Equal(int64(18), usageOrgItem.GetApmAzureAppServiceHostTop99p())
}

func TestUsageAttribution(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	startMonth := tests.ClockFromContext(ctx).Now().AddDate(0, 0, -1)

	Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetUsageAttribution", true)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))
	usage, httpresp, err := api.GetUsageAttribution(ctx, startMonth, "*")
	if err != nil {
		t.Errorf("Error getting Usage Attribution: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(200, httpresp.StatusCode)
	assert.True(usage.HasUsage())
	assert.True(usage.HasMetadata())
}

func TestUsageGetAnalyzedLogsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageAnalyzedLogs(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetHostsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageHosts(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageLogs(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetLogsByIndexErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageLogsByIndex(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetSNMPErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageSNMP(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetBillableSummaryErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageBillableSummary(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetBillableSummary400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

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
	_, httpresp, err := api.GetUsageBillableSummary(ctx)
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUsageTimeSeriesErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageTimeseries(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageTopAvgMetricsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageTopAvgMetrics(ctx, *datadogV1.NewGetUsageTopAvgMetricsOptionalParameters().WithMonth(tests.ClockFromContext(ctx).Now().AddDate(-2, 0, 0)))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageIndexedSpansErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageIndexedSpans(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageSynthetics(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsAPIErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageSyntheticsAPI(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSyntheticsBrowserErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageSyntheticsBrowser(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageFargateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageFargate(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageLambdaErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageLambda(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageRumSessionErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageRumSessions(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageRumSessionBadType(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

	startHr, endHr := getStartEndHr(ctx)
	_, httpresp, err := api.GetUsageRumSessions(ctx, startHr, *datadogV1.NewGetUsageRumSessionsOptionalParameters().WithEndHr(endHr).WithType("invalid"))
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestUsageNetworkHostsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageNetworkHosts(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageNetworkFlowsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageNetworkFlows(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageSummaryErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageSummary(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUsageGetSummary400Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewUsageMeteringApi(Client(ctx))

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
	_, httpresp, err := api.GetUsageSummary(ctx, time.Now())
	assert.Equal(400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestGetSpecifiedDailyCustomReportsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSpecifiedDailyCustomReports", true)
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetSpecifiedDailyCustomReports(ctx, "2010-01-01")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetSpecifiedMonthlyCustomReportsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSpecifiedMonthlyCustomReports", true)
			_, httpresp, err := api.GetSpecifiedMonthlyCustomReports(ctx, "whatever")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetSpecifiedMonthlyCustomReports404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetSpecifiedMonthlyCustomReports", true)
			_, httpresp, err := api.GetSpecifiedMonthlyCustomReports(ctx, "2010-01-01")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetDailyCustomReportsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetDailyCustomReports", true)
			_, httpresp, err := api.GetDailyCustomReports(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetMonthlyCustomReportsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetMonthlyCustomReports", true)
			_, httpresp, err := api.GetMonthlyCustomReports(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetUsageProfilingErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetUsageProfiling(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetIngestedSpansErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetIngestedSpans(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetIncidentManagementErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			_, httpresp, err := api.GetIncidentManagement(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetUsageAttributionErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewUsageMeteringApi(Client(ctx))

			Client(ctx).GetConfig().SetUnstableOperationEnabled("v1.GetUsageAttribution", true)
			_, httpresp, err := api.GetUsageAttribution(ctx, tests.ClockFromContext(ctx).Now().AddDate(0, 1, 0), "*")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
