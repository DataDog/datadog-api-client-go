package datadog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
	is "gotest.tools/assert/cmp"
)

func TestTags(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	api := TESTAPICLIENT.TagsApi
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
		_, httpresp, err := api.GetHostTags(TESTAUTH, hostname).Execute()
		if err != nil {
			t.Logf("Error getting host tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
		}
		return httpresp.StatusCode == 200
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	// test methods
	hostTags := datadog.HostTags{
		Tags: &[]string{"test:client_go"},
	}
	sentHostTags, httpresp, err := api.AddToHostTags(TESTAUTH, hostname).Body(hostTags).Source("datadog").Execute()
	if err != nil {
		t.Fatalf("Error adding tags to %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, hostname, sentHostTags.GetHost())
	assert.Equal(t, hostTags.GetTags(), sentHostTags.GetTags())

	getHostTags, httpresp, err := api.GetHostTags(TESTAUTH, hostname).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, hostTags.GetTags(), getHostTags.GetTags())

	getHostTags, httpresp, err = api.GetHostTags(TESTAUTH, hostname).Source("users").Execute()
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 0, len(getHostTags.GetTags())) // filtering on a different source gives 0 tags

	err = retry(10*time.Second, 10, func() bool {
		hostTagsList, httpresp, err := api.GetAllHostTags(TESTAUTH).Source("datadog").Execute()
		if err != nil {
			t.Logf("Error getting all tags: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
		}
		if httpresp.StatusCode != 200 {
			return false
		}
		hosts, ok := hostTagsList.GetTags()["test:client_go"]
		if !ok {
			return false
		}
		return is.Contains(hosts, hostname)().Success()
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	hostTagsList, httpresp, err := api.GetAllHostTags(TESTAUTH).Source("users").Execute()
	if err != nil {
		t.Errorf("Error getting all tags: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotContains(t, hostTagsList.GetTags(), "test:client_go") // filtering on a different source gives 0 tags

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	updatedHostTags, httpresp, err := api.UpdateHostTags(TESTAUTH, hostname).Body(hostTags).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error updating tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, hostTags.GetTags(), updatedHostTags.GetTags())
	assert.Equal(t, hostname, updatedHostTags.GetHost())

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	httpresp, err = api.RemoveHostTags(TESTAUTH, hostname).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error deleting tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}
