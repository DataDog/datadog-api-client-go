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

func getStartEndHr(c *Client) (time.Time, time.Time) {
	year, month, _ := c.Clock.Now().Date()
	start := time.Date(year, month-1, 12, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 13, 0, 0, 0, 0, time.UTC)
	return start, end
}

func getStartEndMonth(c *Client) (time.Time, time.Time) {
	year, month, _ := c.Clock.Now().Date()
	start := time.Date(year, month-2, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month-1, 1, 0, 0, 0, 0, time.UTC)
	return start, end
}

func TestUsageFargate(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageFargate(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Fargate: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageHosts(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageHosts(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLogs(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageLogs(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSynthetics(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageSynthetics(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSyntheticsAPI(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageSyntheticsAPI(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics API: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageSyntheticsBrowser(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageSyntheticsBrowser(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Synthetics Browser: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTimeseries(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageTimeseries(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Timeseries: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTopAvgMetrics(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageTopAvgMetrics(c.Ctx).Month(c.Clock.Now()).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Avg Metrics: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageTrace(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageTrace(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Trace: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLogsByIndex(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageLogsByIndex(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Logs by Index: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageLambda(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageLambda(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Lambda: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageNetworkHosts(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageNetworkHosts(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Hosts: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageNetworkFlows(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageNetworkFlows(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage Network Flows: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

func TestUsageRumSessions(t *testing.T) {
	c := NewClientWithRecording(WithTestAuth(context.Background()), t)
	defer c.Close()

	startHr, endHr := getStartEndHr(c)
	usage, httpresp, err := c.Client.UsageMeteringApi.GetUsageRumSessions(c.Ctx).StartHr(startHr).EndHr(endHr).Execute()
	if err != nil {
		t.Errorf("Error getting Usage RUM Sessions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, usage.HasUsage())
}

// This test needs multi-org token so make it a unit test
func TestUsageSummary(t *testing.T) {
	c := NewClient(WithFakeAuth(context.Background()), t)
	defer c.Close()

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

	api := c.Client.UsageMeteringApi
	usage, httpresp, err := api.GetUsageSummary(c.Ctx).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageHosts(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageLogs(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageLogsByIndex(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageTimeseries(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageTopAvgMetrics(c.Ctx).Month(c.Clock.Now().AddDate(-2, 0, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageTrace(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageSynthetics(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageSyntheticsAPI(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageSyntheticsBrowser(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageFargate(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageLambda(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageRumSessions(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageNetworkHosts(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageNetworkFlows(c.Ctx).StartHr(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
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
			c := NewClientWithRecording(tc.Ctx(ctx), t)
			defer c.Close()

			_, httpresp, err := c.Client.UsageMeteringApi.GetUsageSummary(c.Ctx).StartMonth(c.Clock.Now().AddDate(0, 1, 0)).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestUsageGetSummary400Error(t *testing.T) {
	c := NewClient(WithFakeAuth(context.Background()), t)
	defer c.Close()

	res, err := tests.ReadFixture("fixtures/usage/error_400.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because it is only returned when the aws integration is not installed, which is not the case on test org
	// and it can't be done through the API
	gock.New("https://api.datadoghq.com").Get("/api/v1/usage/summary").Reply(400).JSON(res)
	defer gock.Off()

	// 400 Bad Request
	_, httpresp, err := c.Client.UsageMeteringApi.GetUsageSummary(c.Ctx).StartMonth(time.Now()).Execute()
	assert.Equal(t, 400, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}
