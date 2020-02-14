/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

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
	hostListResp, httpresp, err := api.GetAllHosts(TESTAUTH).Filter("filter string").Count(4).From(123).SortDir("asc").SortField("status").Start(3).Execute()
	if err != nil {
		t.Errorf("Failed to get hosts: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.True(t, is.DeepEqual(expected, hostListResp)().Success())
}
