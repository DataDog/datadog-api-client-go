/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

var targetTextHTML interface{} = "text/html"
var target2000 interface{} = 2000

var testSyntheticsAPI = datadog.SyntheticsTestDetails{
	Config: &datadog.SyntheticsTestConfig{
		Assertions: []datadog.SyntheticsAssertion{
			{
				Operator: datadog.SYNTHETICSASSERTIONOPERATOR_IS,
				Property: datadog.PtrString("content-type"),
				Target:   &targetTextHTML,
				Type:     datadog.SYNTHETICSASSERTIONTYPE_HEADER,
			},
			{
				Operator: datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
				Target:   &target2000,
				Type:     datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
			},
		},
		Request: datadog.SyntheticsTestRequest{
			Headers: &map[string]string{"testingGoClient": "true"},
			Method:  datadog.HTTPMETHOD_GET,
			Timeout: datadog.PtrFloat64(10),
			Url:     "https://datadoghq.com",
		},
	},
	Locations: &[]string{"aws:us-east-2"},
	Message:   datadog.PtrString("testing Synthetics API test - this is message"),
	Name:      datadog.PtrString("testing Synthetics API test"),
	Options: &datadog.SyntheticsTestOptions{
		AcceptSelfSigned:  datadog.PtrBool(false),
		FollowRedirects:   datadog.PtrBool(true),
		MinLocationFailed: datadog.PtrInt64(10),
		Retry: &datadog.SyntheticsTestOptionsRetry{
			Count:    datadog.PtrInt64(5),
			Interval: datadog.PtrFloat64(10),
		},
		TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
	},
	Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
	Tags:    &[]string{"testing:api"},
	Type:    datadog.SYNTHETICSTESTDETAILSTYPE_API.Ptr(),
}

var testSyntheticsBrowser = datadog.SyntheticsTestDetails{
	Config: &datadog.SyntheticsTestConfig{
		Assertions: []datadog.SyntheticsAssertion{},
		Request: datadog.SyntheticsTestRequest{
			Method: datadog.HTTPMETHOD_GET,
			Url:    "https://datadoghq.com",
		},
	},
	Locations: &[]string{"aws:us-east-2"},
	Message:   datadog.PtrString("testing Synthetics Browser test - this is message"),
	Name:      datadog.PtrString("testing Synthetics Browser test"),
	Options: &datadog.SyntheticsTestOptions{
		AcceptSelfSigned:  datadog.PtrBool(false),
		DeviceIds:         &[]datadog.SyntheticsDeviceId{datadog.SYNTHETICSDEVICEID_TABLET},
		FollowRedirects:   datadog.PtrBool(true),
		MinLocationFailed: datadog.PtrInt64(10),
		Retry: &datadog.SyntheticsTestOptionsRetry{
			Count:    datadog.PtrInt64(5),
			Interval: datadog.PtrFloat64(10),
		},
		TickEvery: datadog.SYNTHETICSTICKINTERVAL_FIVE_MINUTES.Ptr(),
	},
	Tags: &[]string{"testing:browser"},
	Type: datadog.SYNTHETICSTESTDETAILSTYPE_BROWSER.Ptr(),
}

func TestSyntheticsAPITestLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create API test
	synt, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.Error(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicID)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "testing Synthetics API test")

	// Update API test
	synt.SetName("updated name")
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.CreatedAt = nil
	synt.CreatedBy = nil
	synt.ModifiedAt = nil
	synt.PublicId = nil
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTest(TESTAUTH, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "updated name")

	// Get API test
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetTest(TESTAUTH, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "updated name")

	// NOTE: API tests are started by default, so we have to stop it first
	// Stop API test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.SetTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsSetTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pauseStatus, true)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.SetTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsSetTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pauseStatus, true)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetApiTestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetAPITestLatestResults(TESTAUTH, publicID).
		Body(datadog.SyntheticsGetTestLatestResultsPayload{
			FromTs:  0,
			ProbeDc: &locs,
			ToTs:    float64(time.Now().Unix()),
		}).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	// API tests sometimes have a delay before getting first result, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test recent response: %v", err)
	}

	// Get a specific API test result
	// Again, using a mock response to just test deserialization
	var singleResult datadog.SyntheticsApiTestResultFull
	fixturePath, _ = filepath.Abs("fixtures/synthetics/api_test_single_result.json")
	data, _ = ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test single response: %v", err)
	}

	// Delete API test
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestSyntheticsBrowserTestLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create Browser test
	synt, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.Error(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicID)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "testing Synthetics Browser test")

	// Update Browser test
	synt.SetName("updated name")
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.CreatedAt = nil
	synt.CreatedBy = nil
	synt.ModifiedAt = nil
	synt.PublicId = nil
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTest(TESTAUTH, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "updated name")

	// Get Browser test
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetTest(TESTAUTH, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, synt.GetName(), "updated name")

	// NOTE: Browser tests are paused by default, so we have to run it first
	// Start Browser test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.SetTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsSetTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pauseStatus, true)

	// Stop Browser test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.SetTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsSetTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, pauseStatus, true)

	// Get the most recent Browser test results
	var latestResults datadog.SyntheticsGetBrowserTestLatestResultsResponse
	locs := synt.GetLocations()
	latestResults, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetBrowserTestLatestResults(TESTAUTH, publicID).
		Body(datadog.SyntheticsGetTestLatestResultsPayload{
			FromTs:  0,
			ProbeDc: &locs,
			ToTs:    float64(time.Now().Unix()),
		}).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, len(latestResults.GetResults()) == 0)
	// Browser tests are asynchronous and take some time to run, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/browser_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test recent response: %v", err)
	}

	// Get a specific Browser test result
	// Again, using a mock response to just test deserialization
	var singleResult datadog.SyntheticsBrowserTestResultFull
	fixturePath, _ = filepath.Abs("fixtures/synthetics/browser_test_single_result.json")
	data, _ = ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test single response: %v", err)
	}

	// Delete Browser test
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
}

func TestSyntheticsMultipleTestsOperations(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	var syntAPI, syntBrowser datadog.SyntheticsTestDetails
	// Create API test
	syntAPI, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.Error(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDAPI)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, syntAPI.GetName(), "testing Synthetics API test")

	// Create Browser test
	syntBrowser, httpresp, err = TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.Error(), err)
	}
	publicIDBrowser := syntBrowser.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDBrowser)
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Equal(t, syntBrowser.GetName(), "testing Synthetics Browser test")

	// Test Getting multiple tests
	var allTests datadog.SyntheticsGetAllTestsResponse
	allTests, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetAllTests(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics tests: Response %s: %v", err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	td := allTests.GetTests()
	assertPublicIDPresent(t, publicIDAPI, td)
	assertPublicIDPresent(t, publicIDBrowser, td)
}

func TestSyntheticsGetAllLocations(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	locs, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetAllLocations(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics locations: Response %s: %v", err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, len(locs.GetLocations()) > 0)
}

func TestSyntheticsGetAllDevices(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	devices, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetAllDevices(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics devices: Response %s: %v", err.Error(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)
	assert.Assert(t, len(devices.GetDevices()) > 0)
}

func assertPublicIDPresent(t *testing.T, publicID string, syntTests []datadog.SyntheticsTestDetails) {
	for _, st := range syntTests {
		if st.GetPublicId() == publicID {
			return
		}
	}
	assert.NilError(t, fmt.Errorf("Synthetics tests %s expected but not found", publicID))
}

func deleteSyntheticsTestIfExists(testID string) {
	_, httpresp, err := TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{testID}}).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting synthetics test %s failed with %v, Another test may have already deleted this entity: %v",
			testID, httpresp.StatusCode, err)
	}
}
