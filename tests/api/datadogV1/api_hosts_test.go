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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"

	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

func TestHosts(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	api := datadogV1.NewHostsApi(Client(ctx))
	tagsApi := datadogV1.NewTagsApi(Client(ctx))
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
		_, httpresp, err := tagsApi.GetHostTags(ctx, hostname)
		if err != nil {
			t.Logf("Error getting host tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	// test methods
	hostMuteSettings := datadogV1.HostMuteSettings{
		Message: datadog.PtrString("muting for test"),
		End:     datadog.PtrInt64(now + 60),
	}
	muteHostResp, httpresp, err := api.MuteHost(ctx, hostname, hostMuteSettings)
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("muting for test", muteHostResp.GetMessage())
	assert.Equal(hostname, muteHostResp.GetHostname())
	assert.Equal(now+60, muteHostResp.GetEnd())
	assert.Equal("Muted", muteHostResp.GetAction())

	hostMuteSettings = datadogV1.HostMuteSettings{
		Message: datadog.PtrString("muting for test override"),
		End:     datadog.PtrInt64(now + 120),
	}
	muteHostResp, httpresp, err = api.MuteHost(ctx, hostname, hostMuteSettings)
	assert.NotNil(err)
	assert.Equal(400, httpresp.StatusCode)
	hostMuteSettings.SetOverride(true)
	muteHostResp, httpresp, err = api.MuteHost(ctx, hostname, hostMuteSettings)
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("muting for test override", muteHostResp.GetMessage())
	assert.Equal(hostname, muteHostResp.GetHostname())
	assert.Equal(now+120, muteHostResp.GetEnd())
	assert.Equal("Muted", muteHostResp.GetAction())

	muteHostResp, httpresp, err = api.UnmuteHost(ctx, hostname)
	if err != nil {
		t.Errorf("Error unmuting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("Unmuted", muteHostResp.GetAction())
	assert.Equal(hostname, muteHostResp.GetHostname())
}

func TestHostTotalsMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.GetHostTotals")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts/totals").
		MatchParam("from", "123").
		Reply(200).
		JSON(data)

	var expected datadogV1.HostTotals
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.GetHostTotals(ctx, *datadogV1.NewGetHostTotalsOptionalParameters().WithFrom(123))
	if err != nil {
		t.Errorf("Failed to get host totals: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsListErrors(t *testing.T) {
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
			api := datadogV1.NewHostsApi(Client(ctx))

			_, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
				WithCount(-1))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsGetTotalsErrors(t *testing.T) {
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
			api := datadogV1.NewHostsApi(Client(ctx))

			_, httpresp, err := api.GetHostTotals(ctx, *datadogV1.NewGetHostTotalsOptionalParameters().
				WithFrom(tests.ClockFromContext(ctx).Now().Unix() + 60))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsMuteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewHostsApi(Client(ctx))

	// Mute host a first time in order to trigger a 400
	hostname := *tests.UniqueEntityName(ctx, t)
	muteSettings := *datadogV1.NewHostMuteSettingsWithDefaults()
	muteSettings.SetOverride(true)
	_, httpresp, err := api.MuteHost(ctx, hostname, muteSettings)
	if err != nil {
		t.Fatalf("Error muting host %s: Response: %s", hostname, err.(datadog.GenericOpenAPIError).Body())
	}
	defer api.UnmuteHost(ctx, hostname)
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
			api := datadogV1.NewHostsApi(Client(ctx))

			_, httpresp, err := api.MuteHost(ctx, hostname, datadogV1.HostMuteSettings{})
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsUnmuteErrors(t *testing.T) {
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
			api := datadogV1.NewHostsApi(Client(ctx))

			hostname := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := api.UnmuteHost(ctx, hostname)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestHostsSearchMockedIncludeMutedHostsDataFalse(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	params := map[string]string{
		"filter":                   "filter string",
		"count":                    "4",
		"from":                     "123",
		"sort_dir":                 "asc",
		"sort_field":               "status",
		"start":                    "3",
		"include_muted_hosts_data": "false",
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3).
		WithIncludeMutedHostsData(false))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMockedIncludeMutedHostsDataTrue(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search_with_mutes.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	params := map[string]string{
		"filter":                   "filter string",
		"count":                    "4",
		"from":                     "123",
		"sort_dir":                 "asc",
		"sort_field":               "status",
		"start":                    "3",
		"include_muted_hosts_data": "true",
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).
		WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3).
		WithIncludeMutedHostsData(true))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMockedIncludeMutedHostsDataDefault(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search_with_mutes.json")
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
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).
		WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMockedIncludeHostsMetadataFalse(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	params := map[string]string{
		"filter":                 "filter string",
		"count":                  "4",
		"from":                   "123",
		"sort_dir":               "asc",
		"sort_field":             "status",
		"start":                  "3",
		"include_hosts_metadata": "false",
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).
		WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3).
		WithIncludeHostsMetadata(false))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMockedIncludeHostsMetadataTrue(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search_with_metadata.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	params := map[string]string{
		"filter":                 "filter string",
		"count":                  "4",
		"from":                   "123",
		"sort_dir":               "asc",
		"sort_field":             "status",
		"start":                  "3",
		"include_hosts_metadata": "true",
	}
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).
		WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3).
		WithIncludeHostsMetadata(true))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMockedIncludeHostsMetadataDefault(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	data, err := tests.ReadFixture("fixtures/hosts/host_search_with_metadata.json")
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
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "HostsApiService.ListHosts")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/hosts").
		MatchParams(params).
		Reply(200).
		JSON(data)

	var expected datadogV1.HostListResponse
	json.Unmarshal([]byte(data), &expected)

	api := datadogV1.NewHostsApi(Client(ctx))
	hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
		WithFilter("filter string").
		WithCount(4).WithFrom(123).
		WithSortDir("asc").
		WithSortField("status").
		WithStart(3))
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsIncludeMutedHostsDataFunctional(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	if tests.GetRecording() == tests.ModeIgnore {
		t.Skipf("Slow test")
	}

	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	api := datadogV1.NewHostsApi(Client(ctx))
	tagsApi := datadogV1.NewTagsApi(Client(ctx))
	now := tests.ClockFromContext(ctx).Now().Unix()

	hostname := *tests.UniqueEntityName(ctx, t)

	var hostListResp datadogV1.HostListResponse

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
		_, httpresp, err := tagsApi.GetHostTags(ctx, hostname)
		if err != nil {
			t.Logf("Error getting host tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	// waiting for host to appear on infralist
	err = tests.Retry(10*time.Second, 10, func() bool {
		hostListResp, httpresp, err := api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().WithFilter(hostname))
		if err != nil {
			t.Errorf("Failed to get hosts: %v", err)
		}
		return httpresp.StatusCode == 200 && *hostListResp.TotalReturned == 1
	})

	// muting the host
	hostMuteSettings := datadogV1.HostMuteSettings{
		Message: datadog.PtrString("muting for test"),
		End:     datadog.PtrInt64(now + 1200),
	}
	muteHostResp, httpresp, err := api.MuteHost(ctx, hostname, hostMuteSettings)
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("muting for test", muteHostResp.GetMessage())
	assert.Equal(hostname, muteHostResp.GetHostname())
	assert.Equal(now+1200, muteHostResp.GetEnd())
	assert.Equal("Muted", muteHostResp.GetAction())

	// waiting for host to be muted
	err = tests.Retry(10*time.Second, 10, func() bool {
		hostListResp, httpresp, err = api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().WithFilter(hostname))
		if err != nil {
			t.Errorf("Failed to get hosts: %v", err)
		}
		return httpresp.StatusCode == 200 && *hostListResp.TotalReturned == 1 && *hostListResp.HostList[0].IsMuted
	})

	// this is the default case, should have the muted data
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(1), *hostListResp.TotalReturned)
	host := hostListResp.HostList[0]
	assert.True(*host.IsMuted)
	assert.NotEqual(host.MuteTimeout, nil)

	// this is the case where the include_muted_hosts_data is true, should have muted data
	err = tests.Retry(10*time.Second, 10, func() bool {
		hostListResp, httpresp, err = api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
			WithFilter(hostname).
			WithIncludeMutedHostsData(true))
		if err != nil {
			t.Errorf("Failed to get hosts: %v", err)
		}
		return httpresp.StatusCode == 200 && *hostListResp.TotalReturned == 1 && *hostListResp.HostList[0].IsMuted
	})
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(1), *hostListResp.TotalReturned)
	host = hostListResp.HostList[0]
	assert.True(*host.IsMuted)
	assert.NotEqual(host.MuteTimeout, nil)

	// this is the case where the include_muted_hosts_data is false, should not have muted data
	err = tests.Retry(10*time.Second, 10, func() bool {
		hostListResp, httpresp, err = api.ListHosts(ctx, *datadogV1.NewListHostsOptionalParameters().
			WithFilter(hostname).
			WithIncludeMutedHostsData(false))
		if err != nil {
			t.Errorf("Failed to get hosts: %v", err)
		}
		return httpresp.StatusCode == 200 && *hostListResp.TotalReturned == 1 && *hostListResp.HostList[0].IsMuted
	})
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(1), *hostListResp.TotalReturned)
	host = hostListResp.HostList[0]
	assert.False(*host.IsMuted)
	assert.Nil(host.MuteTimeout)
}
