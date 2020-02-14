// +build integration

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
)

func TestIPRanges(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Get IP ranges
	ipRanges, httpresp, err := TESTAPICLIENT.IPRangesApi.GetIPRanges(TESTAUTH).Execute()
	if err != nil {
		t.Errorf("Error getting IP ranges: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.NotEmpty(t, ipRanges.Agents.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Api.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Apm.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Logs.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Process.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Synthetics.GetPrefixesIpv4())
	assert.NotEmpty(t, ipRanges.Webhooks.GetPrefixesIpv4())
}
