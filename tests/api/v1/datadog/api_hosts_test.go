// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
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
