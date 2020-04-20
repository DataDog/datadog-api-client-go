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
	"log"
	"path/filepath"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"
	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/stretchr/testify/assert"
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
	Message:   datadog.PtrString("Go client testing Synthetics API test - this is message"),
	Name:      datadog.PtrString("Go client testing Synthetics API test"),
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
	Message:   datadog.PtrString("Go client testing Synthetics Browser test - this is message"),
	Name:      datadog.PtrString("Go client testing Synthetics Browser test"),
	Options: &datadog.SyntheticsTestOptions{
		AcceptSelfSigned:  datadog.PtrBool(false),
		DeviceIds:         &[]datadog.SyntheticsDeviceID{datadog.SYNTHETICSDEVICEID_TABLET},
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
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicID)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := "go client updated name"
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.CreatedAt = nil
	synt.CreatedBy = nil
	synt.ModifiedAt = nil
	synt.PublicId = nil
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTest(TESTAUTH, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedName, synt.GetName())

	// Get API test
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetTest(TESTAUTH, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedName, synt.GetName())

	// NOTE: API tests are started by default, so we have to stop it first
	// Stop API test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetAPITestLatestResults(TESTAUTH, publicID).
		FromTs(0).
		ToTs(TESTCLOCK.Now().Unix() * 1000).
		ProbeDc(locs).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)

	// API tests sometimes have a delay before getting first result, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test recent response: %v", err)
	}

	// Delete API test
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestSyntheticsBrowserTestLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create Browser test
	synt, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicID)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testSyntheticsBrowser.GetName(), synt.GetName())

	// Update Browser test
	updatedName := "go client browser updated name"
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.CreatedAt = nil
	synt.CreatedBy = nil
	synt.ModifiedAt = nil
	synt.PublicId = nil
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTest(TESTAUTH, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedName, synt.GetName())

	// Get Browser test
	synt, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetTest(TESTAUTH, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, updatedName, synt.GetName())

	// NOTE: Browser tests are paused by default, so we have to run it first
	// Start Browser test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, true, pauseStatus)

	// Stop Browser test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = TESTAPICLIENT.SyntheticsApi.UpdateTestPauseStatus(TESTAUTH, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, true, pauseStatus)

	// Get the most recent Browser test results
	var latestResults datadog.SyntheticsGetBrowserTestLatestResultsResponse
	locs := synt.GetLocations()
	latestResults, httpresp, err = TESTAPICLIENT.SyntheticsApi.GetBrowserTestLatestResults(TESTAUTH, publicID).
		FromTs(0).
		ToTs(TESTCLOCK.Now().Unix() * 1000).
		ProbeDc(locs).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Empty(t, latestResults.GetResults())
	// Browser tests are asynchronous and take some time to run, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/browser_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test recent response: %v", err)
	}

	// Delete Browser test
	_, httpresp, err = TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
}

func TestSyntheticsGetBrowserTestResult(t *testing.T) {
	teardownUnitTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownUnitTest(t)

	// Test that the get browser result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsBrowserTestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/browser_test_single_result.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err := json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test single response: %v", err)
	}

	gock.New("https://api.datadoghq.com/api/v1").
		Get("/synthetics/tests/browser/test-synthetics-id/results/test-result-id").
		Reply(200).
		JSON(data)

	unitAPI := TESTAPICLIENT.SyntheticsApi
	browserResp, httpresp, err := unitAPI.GetBrowserTestResult(TESTAUTH, "test-synthetics-id", "test-result-id").Execute()
	if err != nil {
		t.Errorf("Failed to get synthetics browser test result: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, "5140738909114888212", browserResp.GetResultId())
	assert.Equal(t, datadog.SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED, browserResp.GetStatus())
	assert.Equal(t, float64(1579711893111), browserResp.GetCheckTime())
	assert.Equal(t, int64(5), browserResp.GetCheckVersion())
}

func TestSyntheticsGetApiTestResult(t *testing.T) {
	teardownUnitTest := setupUnitTest(t)
	defer gock.Off()
	defer teardownUnitTest(t)

	// Test that the get api result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsAPITestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_single_result.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err := json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test single response: %v", err)
	}

	gock.New("https://api.datadoghq.com/api/v1").
		Get("/synthetics/tests/test-synthetics-id/results/test-result-id").
		Reply(200).
		JSON(data)

	unitAPI := TESTAPICLIENT.SyntheticsApi
	apiResp, httpresp, err := unitAPI.GetAPITestResult(TESTAUTH, "test-synthetics-id", "test-result-id").Execute()
	if err != nil {
		t.Errorf("Failed to get synthetics api test result: %v", err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, datadog.SYNTHETICSTESTMONITORSTATUS_TRIGGERED, apiResp.GetStatus())
	assert.Equal(t, float64(1580204310361), apiResp.GetCheckTime())
	assert.Equal(t, int64(2), apiResp.GetCheckVersion())
	assert.Equal(t, "7761116396307201795", apiResp.GetResultId())
	apiResult := apiResp.GetResult()
	assert.Equal(t, datadog.SYNTHETICSTESTPROCESSSTATUS_FINISHED, apiResult.GetEventType())
}

func TestSyntheticsMultipleTestsOperations(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	var syntAPI, syntBrowser datadog.SyntheticsTestDetails
	// Create API test
	syntAPI, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDAPI)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testSyntheticsAPI.GetName(), syntAPI.GetName())

	// Create Browser test
	syntBrowser, httpresp, err = TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDBrowser := syntBrowser.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDBrowser)
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.Equal(t, testSyntheticsBrowser.GetName(), syntBrowser.GetName())

	// Test Getting multiple tests
	var allTests datadog.SyntheticsListTestsResponse
	allTests, httpresp, err = TESTAPICLIENT.SyntheticsApi.ListTests(TESTAUTH).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics tests: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	td := allTests.GetTests()
	assertPublicIDPresent(t, publicIDAPI, td)
	assertPublicIDPresent(t, publicIDBrowser, td)
}

func TestSyntheticsDeleteTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		Body               datadog.SyntheticsDeleteTestsPayload
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, datadog.SyntheticsDeleteTestsPayload{}, 400},
		{"403 Forbidden", fake_auth, datadog.SyntheticsDeleteTestsPayload{}, 403},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.DeleteTests(tc.Ctx).Body(tc.Body).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsDeleteTest404Error(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because not sure how to trigger the 404 response
	gock.New("https://api.datadoghq.com").Post("/api/v1/synthetics/tests/delete").Reply(404).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).Body(datadog.SyntheticsDeleteTestsPayload{}).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestSyntheticsUpdateStatusTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create API test
	syntAPI, _, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDAPI)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, publicIDAPI, 400},
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.UpdateTestPauseStatus(tc.Ctx, tc.ID).Body(datadog.SyntheticsUpdateTestPauseStatusPayload{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserResultsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetBrowserTestLatestResults(tc.Ctx, tc.ID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPIResultsErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetAPITestLatestResults(tc.Ctx, tc.ID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserSpecificResultErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetBrowserTestResult(tc.Ctx, tc.ID, "resultid").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPISpecificResultErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetAPITestResult(tc.Ctx, tc.ID, "resultid").Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.GetTest(tc.Ctx, tc.ID).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsUpdateTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create API test
	syntAPI, _, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(publicIDAPI)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		{"400 Bad Request", TESTAUTH, publicIDAPI, 400},
		{"403 Forbidden", fake_auth, "id", 403},
		{"404 Not Found", TESTAUTH, "aaa-aaa-aaa", 404},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.UpdateTest(tc.Ctx, tc.ID).Body(datadog.SyntheticsTestDetails{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	teardownTest := setupTest(t)
	defer teardownTest(t)

	testCases := []struct {
		Name               string
		Ctx                context.Context
		ExpectedStatusCode int
	}{
		{"403 Forbidden", fake_auth, 403},
		// Mock 404 because it's returned when synthetics is not enabled for the org
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.ListTests(tc.Ctx).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTest404Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mock 404 because it's returned when synthetics is not enabled for the org
	gock.New("https://api.datadoghq.com").Get("/api/v1/synthetics/test").Reply(404).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.SyntheticsApi.ListTests(TESTAUTH).Execute()
	assert.Equal(t, 404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func TestSyntheticsCreateTestErrors(t *testing.T) {
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
			_, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(tc.Ctx).Body(datadog.SyntheticsTestDetails{}).Execute()
			assert.Equal(t, tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(t, ok)
			assert.NotEmpty(t, apiError.GetErrors())
		})
	}
}

func TestSyntheticsCreateTest402Error(t *testing.T) {
	teardownTest := setupUnitTest(t)
	defer teardownTest(t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_402.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Triggered when the synthetics test quota is reached. Need to mock.
	gock.New("https://api.datadoghq.com").Post("/api/v1/synthetics/test").Reply(402).JSON(res)
	defer gock.Off()

	_, httpresp, err := TESTAPICLIENT.SyntheticsApi.CreateTest(TESTAUTH).Body(datadog.SyntheticsTestDetails{}).Execute()
	assert.Equal(t, 402, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(t, ok)
	assert.NotEmpty(t, apiError.GetErrors())
}

func assertPublicIDPresent(t *testing.T, publicID string, syntTests []datadog.SyntheticsTestDetails) {
	for _, st := range syntTests {
		if st.GetPublicId() == publicID {
			return
		}
	}
	assert.Nil(t, fmt.Errorf("Synthetics tests %s expected but not found", publicID))
}

func deleteSyntheticsTestIfExists(testID string) {
	_, httpresp, err := TESTAPICLIENT.SyntheticsApi.DeleteTests(TESTAUTH).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{testID}}).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting synthetics test %s failed with %v, Another test may have already deleted this entity: %v",
			testID, httpresp.StatusCode, err)
	}
}
