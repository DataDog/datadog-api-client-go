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
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

func TestHosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	api := TESTAPICLIENT.HostsApi
	now := TESTCLOCK.Now().Unix()

	hostname := fmt.Sprintf("go-client-test-host-%d", now)

	// create host by sending a metric
	metricsPayload := fmt.Sprintf(
		`{"series": [{"host": "%s", "metric": "go.client.test.metric", "points": [[%f, 0]]}]}`,
		hostname, float64(now),
	)
	httpresp, respBody, err := sendRequest("POST", "/api/v1/series", []byte(metricsPayload))
	if err != nil {
		t.Fatalf("Error submitting metric: Response %s: %v", string(respBody), err)
	}
	assert.Equal(t, 202, httpresp.StatusCode)
	assert.Equal(t, `{"status": "ok"}`, string(respBody))

	// wait for host to appear
	err = tests.Retry(10*time.Second, 10, func() bool {
		_, httpresp, err := TESTAPICLIENT.TagsApi.GetHostTags(TESTAUTH, hostname).Execute()
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
	muteHostResp, httpresp, err := api.MuteHost(TESTAUTH, hostname).Body(hostMuteSettings).Execute()
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, "muting for test", muteHostResp.GetMessage())
	assert.Equal(t, hostname, muteHostResp.GetHostname())
	assert.Equal(t, now+60, muteHostResp.GetEnd())
	assert.Equal(t, "Muted", muteHostResp.GetAction())

	hostMuteSettings = datadog.HostMuteSettings{
		Message: datadog.PtrString("muting for test override"),
		End:     datadog.PtrInt64(now + 120),
	}
	muteHostResp, httpresp, err = api.MuteHost(TESTAUTH, hostname).Body(hostMuteSettings).Execute()
	assert.NotNil(t, err)
	assert.Equal(t, 400, httpresp.StatusCode)
	hostMuteSettings.SetOverride(true)
	muteHostResp, httpresp, err = api.MuteHost(TESTAUTH, hostname).Body(hostMuteSettings).Execute()
	if err != nil {
		t.Errorf("Error muting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, "muting for test override", muteHostResp.GetMessage())
	assert.Equal(t, hostname, muteHostResp.GetHostname())
	assert.Equal(t, now+120, muteHostResp.GetEnd())
	assert.Equal(t, "Muted", muteHostResp.GetAction())

	muteHostResp, httpresp, err = api.UnmuteHost(TESTAUTH, hostname).Execute()
	if err != nil {
		t.Errorf("Error unmuting host %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, "Unmuted", muteHostResp.GetAction())
	assert.Equal(t, hostname, muteHostResp.GetHostname())
}

func TestHostTotalsMocked(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
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

	api := TESTAPICLIENT.HostsApi
	hostListResp, httpresp, err := api.GetHostTotals(TESTAUTH).From(123).Execute()
	if err != nil {
		t.Errorf("Failed to get host totals: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsSearchMocked(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)
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

	api := TESTAPICLIENT.HostsApi
	hostListResp, httpresp, err := api.ListHosts(TESTAUTH).Filter("filter string").Count(4).From(123).SortDir("asc").SortField("status").Start(3).Execute()
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, is.DeepEqual(expected, hostListResp)().Success())
}

func TestHostsListErrors(t *testing.T) {
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
			_, httpresp, err := TESTAPICLIENT.HostsApi.GetAllHosts(tc.Ctx).Count(-1).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestHostsGetTotalsErrors(t *testing.T) {
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
			_, httpresp, err := TESTAPICLIENT.HostsApi.GetHostTotals(tc.Ctx).From(TESTCLOCK.Now().Unix() + 60).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestHostsMuteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Mute host a first time in order to trigger a 400
	hostname := fmt.Sprintf("go-client-test-hostname-%d", TESTCLOCK.Now().Unix())
	muteSettings := *datadog.NewHostMuteSettingsWithDefaults()
	muteSettings.SetOverride(true)
	_, httpresp, err := TESTAPICLIENT.HostsApi.MuteHost(TESTAUTH, hostname).Body(muteSettings).Execute()
	if err != nil {
		t.Fatalf("Error muting host %s: Response: %s", hostname, err.(datadog.GenericOpenAPIError).Body())
	}
	defer TESTAPICLIENT.HostsApi.UnmuteHost(TESTAUTH, hostname).Execute()
	assert.Equal(t, 200, httpresp.StatusCode)

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
			_, httpresp, err := TESTAPICLIENT.HostsApi.MuteHost(tc.Ctx, hostname).Body(datadog.HostMuteSettings{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestHostsUnmuteErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	hostname := fmt.Sprintf("go-client-test-hostname-%d", TESTCLOCK.Now().Unix())

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
			_, httpresp, err := TESTAPICLIENT.HostsApi.UnmuteHost(tc.Ctx, hostname).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
