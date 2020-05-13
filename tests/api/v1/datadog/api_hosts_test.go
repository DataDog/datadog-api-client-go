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
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

func TestHosts(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	api := Client(ctx).HostsApi
	now := tests.ClockFromContext(ctx).Now().Unix()

	hostname := *tests.UniqueEntityName(ctx, t)

	// create host by sending a metric
	metricsPayload := fmt.Sprintf(
		`{"series": [{"host": "%s", "metric": "go.client.test.metric", "points": [[%f, 0]]}]}`,
		hostname, float64(now),
	)
	httpresp, respBody, err := SendRequest(ctx, "POST", "/api/v1/series", []byte(metricsPayload))
	if err != nil {
		t.Fatalf("Error submitting metric: Response %s: %v", string(respBody), err)
	}
	assert.Equal(202, httpresp.StatusCode)
	assert.Equal(`{"status": "ok"}`, string(respBody))

	// wait for host to appear
	err = tests.Retry(10*time.Second, 10, func() bool {
		_, httpresp, err := Client(ctx).TagsApi.GetHostTags(ctx, hostname).Execute()
		if err != nil {
			t.Logf("Error getting host tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	// test methods
	hostMuteSettings := datadog.HostMuteSettings{
		Message: datadog.PtrString("muting for test"),
		End:     datadog.PtrInt64(now + 60),
	}
	muteHostResp, httpresp, err := api.MuteHost(ctx, hostname).Body(hostMuteSettings).Execute()
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("muting for test", muteHostResp.GetMessage())
	assert.Equal(hostname, muteHostResp.GetHostname())
	assert.Equal(now+60, muteHostResp.GetEnd())
	assert.Equal("Muted", muteHostResp.GetAction())

	hostMuteSettings = datadog.HostMuteSettings{
		Message: datadog.PtrString("muting for test override"),
		End:     datadog.PtrInt64(now + 120),
	}
	muteHostResp, httpresp, err = api.MuteHost(ctx, hostname).Body(hostMuteSettings).Execute()
	assert.NotNil(err)
	assert.Equal(400, httpresp.StatusCode)
	hostMuteSettings.SetOverride(true)
	muteHostResp, httpresp, err = api.MuteHost(ctx, hostname).Body(hostMuteSettings).Execute()
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("muting for test override", muteHostResp.GetMessage())
	assert.Equal(hostname, muteHostResp.GetHostname())
	assert.Equal(now+120, muteHostResp.GetEnd())
	assert.Equal("Muted", muteHostResp.GetAction())

	muteHostResp, httpresp, err = api.UnmuteHost(ctx, hostname).Execute()
	if err != nil {
		t.Errorf("Error unmuting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("Unmuted", muteHostResp.GetAction())
	assert.Equal(hostname, muteHostResp.GetHostname())
}

func TestHostTotalsMocked(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/hosts/host_totals.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	gock.New("https://api.datadoghq.com/api/v1").
		Get("/hosts/totals").
		MatchParam("from", "123").
		Reply(200).
		JSON(data)

	var expected datadog.HostTotals
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).HostsApi
	hostListResp, httpresp, err := api.GetHostTotals(ctx).From(123).Execute()
	if err != nil {
		t.Errorf("Failed to get host totals: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMocked(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	params := map[string]string{
		"filter":     "filter string",
		"count":      "4",
		"from":       "123",
		"sort_dir":   "asc",
		"sort_field": "status",
		"start":      "3",
	}
	gock.New("https://api.datadoghq.com/api/v1").
		Get("/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadog.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := Client(ctx).HostsApi
	hostListResp, httpresp, err := api.ListHosts(ctx).Filter("filter string").Count(4).From(123).SortDir("asc").SortField("status").Start(3).Execute()
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsListErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).HostsApi.ListHosts(ctx).Count(-1).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsGetTotalsErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).HostsApi.GetHostTotals(ctx).From(tests.ClockFromContext(ctx).Now().Unix() + 60).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsMuteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Mute host a first time in order to trigger a 400
	hostname := *tests.UniqueEntityName(ctx, t)
	muteSettings := *datadog.NewHostMuteSettingsWithDefaults()
	muteSettings.SetOverride(true)
	_, httpresp, err := Client(ctx).HostsApi.MuteHost(ctx, hostname).Body(muteSettings).Execute()
	if err != nil {
		t.Fatalf("Error muting host %s: Response: %s", hostname, err.(datadog.GenericOpenAPIError).Body())
	}
	defer Client(ctx).HostsApi.UnmuteHost(ctx, hostname).Execute()
	assert.Equal(200, httpresp.StatusCode)

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

			_, httpresp, err := Client(ctx).HostsApi.MuteHost(ctx, hostname).Body(datadog.HostMuteSettings{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsUnmuteErrors(t *testing.T) {
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

			hostname := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := Client(ctx).HostsApi.UnmuteHost(ctx, hostname).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
