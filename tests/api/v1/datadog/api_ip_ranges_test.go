/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gopkg.in/h2non/gock.v1"
)

func TestIPRanges(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Get IP ranges
	ipRanges, httpresp, err := Client(ctx).IPRangesApi.GetIPRanges(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting IP ranges: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.NotEmpty(ipRanges.Agents.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Api.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Apm.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Logs.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Process.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Synthetics.GetPrefixesIpv4())
	assert.NotEmpty(ipRanges.Webhooks.GetPrefixesIpv4())
}

func TestIPRangesMocked(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer gock.Off()
	defer stop()
	assert := tests.Assert(ctx, t)

	data, err := tests.ReadFixture("fixtures/ip-ranges/ip-ranges.json")
	if err != nil {
		t.Errorf("Failed to read fixture: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "IPRangesApiService.GetIPRanges")
	assert.NoError(err)
	gock.New(URL).
		Get("/").
		Reply(200).
		JSON(data)

	// Get IP ranges
	ipRanges, httpresp, err := Client(ctx).IPRangesApi.GetIPRanges(ctx).Execute()
	if err != nil {
		t.Errorf("Error getting IP ranges: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(int64(11), ipRanges.GetVersion())
	assert.Equal("2019-07-29-11-48-00", ipRanges.GetModified())
	assert.Equal([]string{"ipv4"}, ipRanges.Agents.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Api.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Apm.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Logs.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Process.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Synthetics.GetPrefixesIpv4())
	assert.Equal([]string{"ipv4"}, ipRanges.Webhooks.GetPrefixesIpv4())
	assert.Equal([]string{"ipv6"}, ipRanges.Agents.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Api.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Apm.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Logs.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Process.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Synthetics.GetPrefixesIpv6())
	assert.Equal([]string{"ipv6"}, ipRanges.Webhooks.GetPrefixesIpv6())
}
