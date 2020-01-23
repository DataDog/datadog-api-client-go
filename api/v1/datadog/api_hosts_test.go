package datadog_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

func TestHosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	api := TESTAPICLIENT.HostsApi
	now := time.Now().Unix()

	hostname := fmt.Sprintf("go-client-test-host-%d", now)

	// create host by sending a metric
	metricsPayload := datadog.MetricsPayload{
		Series: &[]datadog.Series{{
			Host:   &hostname,
			Metric: "go.client.test.metric",
			Points: [][]float64{{float64(now), 0}},
		}},
	}
	r, httpresp, err := TESTAPICLIENT.MetricsApi.SubmitMetrics(TESTAUTH).Body(metricsPayload).Execute()
	if err != nil {
		t.Fatalf("Error submitting metric: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 202, httpresp.StatusCode)
	assert.Equal(t, "ok", r.GetStatus())

	// wait for host to appear
	err = retry(10*time.Second, 10, func() bool {
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

	data, err := readFixture("fixtures/hosts/host_search.json")
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
	hostListResp, httpresp, err := api.GetAllHosts(TESTAUTH).Filter("filter string").Count(4).From(123).SortDir("asc").SortField("status").Start(3).Execute()
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, is.DeepEqual(expected, hostListResp)().Success())
}
