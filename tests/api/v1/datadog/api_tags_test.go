/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
	is "gotest.tools/assert/cmp"
)

func TestTags(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	api := Client(ctx).TagsApi
	now := tests.ClockFromContext(ctx).Now().Unix()

	hostname := fmt.Sprintf("go-client-test-host-%d", now)

	// create host by sending a metric
	metricsPayload := fmt.Sprintf(
		`{"series": [{"host": "%s", "metric": "go.client.test.metric", "points": [[%f, 0]]}]}`,
		hostname, float64(now),
	)
	httpresp, respBody, err := SendRequest(ctx, "POST", "/api/v1/series", []byte(metricsPayload))
	if err != nil {
		t.Fatalf("Error submitting metric: Response %s: %v", string(respBody), err)
	}
	assert.Equal(t, 202, httpresp.StatusCode)
	assert.Equal(t, `{"status": "ok"}`, string(respBody))

	// wait for host to appear
	err = tests.Retry(10*time.Second, 10, func() bool {
		_, httpresp, err := api.GetHostTags(ctx, hostname).Execute()
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
	sentHostTags, httpresp, err := api.CreateHostTags(ctx, hostname).Body(hostTags).Source("datadog").Execute()
	if err != nil {
		t.Fatalf("Error adding tags to %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, hostname, sentHostTags.GetHost())
	assert.Equal(t, hostTags.GetTags(), sentHostTags.GetTags())

	getHostTags, httpresp, err := api.GetHostTags(ctx, hostname).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, hostTags.GetTags(), getHostTags.GetTags())

	getHostTags, httpresp, err = api.GetHostTags(ctx, hostname).Source("users").Execute()
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, 0, len(getHostTags.GetTags())) // filtering on a different source gives 0 tags

	err = tests.Retry(10*time.Second, 10, func() bool {
		hostTagsList, httpresp, err := api.ListHostTags(ctx).Source("datadog").Execute()
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

	hostTagsList, httpresp, err := api.ListHostTags(ctx).Source("users").Execute()
	if err != nil {
		t.Errorf("Error getting all tags: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotContains(t, hostTagsList.GetTags(), "test:client_go") // filtering on a different source gives 0 tags

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	updatedHostTags, httpresp, err := api.UpdateHostTags(ctx, hostname).Body(hostTags).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error updating tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 201, httpresp.StatusCode)
	assert.Equal(t, hostTags.GetTags(), updatedHostTags.GetTags())
	assert.Equal(t, hostname, updatedHostTags.GetHost())

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	httpresp, err = api.DeleteHostTags(ctx, hostname).Source("datadog").Execute()
	if err != nil {
		t.Errorf("Error deleting tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 204, httpresp.StatusCode)
}

func TestTagsListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).TagsApi.ListHostTags(ctx).Source("nosource").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestTagsGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).TagsApi.GetHostTags(ctx, "notahostname1234").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestTagsCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).TagsApi.CreateHostTags(ctx, "notahostname1234").Body(datadog.HostTags{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestTagsUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			_, httpresp, err := Client(ctx).TagsApi.UpdateHostTags(ctx, "notahostname1234").Body(datadog.HostTags{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestTagsDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, stop := WithRecorder(tc.Ctx(ctx), t)
			defer stop()

			httpresp, err := Client(ctx).TagsApi.DeleteHostTags(ctx, "notahostname1234").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}
