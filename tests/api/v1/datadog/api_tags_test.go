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
	is "gotest.tools/assert/cmp"
)

func TestTags(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	if tests.GetRecording() == tests.ModeIgnore {
		t.Skipf("Slow test")
	}
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	api := Client(ctx).TagsApi
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
		_, httpresp, err := api.GetHostTags(ctx, hostname)
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
	sentHostTags, httpresp, err := api.CreateHostTags(ctx, hostname, hostTags, *datadog.NewCreateHostTagsOptionalParameters().
		WithSource("datadog"))
	if err != nil {
		t.Fatalf("Error adding tags to %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(201, httpresp.StatusCode)
	assert.Equal(hostname, sentHostTags.GetHost())
	assert.Equal(hostTags.GetTags(), sentHostTags.GetTags())

	getHostTags, httpresp, err := api.GetHostTags(ctx, hostname, *datadog.NewGetHostTagsOptionalParameters().
		WithSource("datadog"))
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(hostTags.GetTags(), getHostTags.GetTags())

	getHostTags, httpresp, err = api.GetHostTags(ctx, hostname, *datadog.NewGetHostTagsOptionalParameters().
		WithSource("users"))
	if err != nil {
		t.Errorf("Error getting tags from %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(0, len(getHostTags.GetTags())) // filtering on a different source gives 0 tags

	err = tests.Retry(10*time.Second, 10, func() bool {
		hostTagsList, httpresp, err := api.ListHostTags(ctx, *datadog.NewListHostTagsOptionalParameters().
			WithSource("datadog"))
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

	hostTagsList, httpresp, err := api.ListHostTags(ctx, *datadog.NewListHostTagsOptionalParameters().
		WithSource("users"))
	if err != nil {
		t.Errorf("Error getting all tags: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.NotContains(hostTagsList.GetTags(), "test:client_go") // filtering on a different source gives 0 tags

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	updatedHostTags, httpresp, err := api.UpdateHostTags(ctx, hostname, hostTags, *datadog.NewUpdateHostTagsOptionalParameters().
		WithSource("datadog"))
	if err != nil {
		t.Errorf("Error updating tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(201, httpresp.StatusCode)
	assert.Equal(hostTags.GetTags(), updatedHostTags.GetTags())
	assert.Equal(hostname, updatedHostTags.GetHost())

	hostTags.Tags = &[]string{"foo:bar", "toto:tata"}
	httpresp, err = api.DeleteHostTags(ctx, hostname, *datadog.NewDeleteHostTagsOptionalParameters().WithSource("datadog"))
	if err != nil {
		t.Errorf("Error deleting tags for %s: Response %s: %v", hostname, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)
}

func TestTagsListErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).TagsApi.ListHostTags(ctx, *datadog.NewListHostTagsOptionalParameters().
				WithSource("nosource"))
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestTagsGetErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).TagsApi.GetHostTags(ctx, "notahostname1234")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestTagsCreateErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).TagsApi.CreateHostTags(ctx, "notahostname1234", datadog.HostTags{})
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestTagsUpdateErrors(t *testing.T) {
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

			_, httpresp, err := Client(ctx).TagsApi.UpdateHostTags(ctx, "notahostname1234", datadog.HostTags{})
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestTagsDeleteErrors(t *testing.T) {
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

			httpresp, err := Client(ctx).TagsApi.DeleteHostTags(ctx, "notahostname1234")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
