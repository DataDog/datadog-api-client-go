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
)

var targetTextHTML interface{} = "text/html"
var target2000 interface{} = 2000
var targetValue0 interface{} = "0"

func getTestSyntheticsAPI(ctx context.Context, t *testing.T) datadog.SyntheticsTestDetails {
	assertionTextHTML := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_IS, datadog.SYNTHETICSASSERTIONTYPE_HEADER)
	assertionTextHTML.Property = datadog.PtrString("content-type")
	assertionTextHTML.Target = &targetTextHTML

	assertion2000 := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME)
	assertion2000.Target = &target2000

	targetJSONPath := datadog.NewSyntheticsAssertionJSONPathTarget(datadog.SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH, datadog.SYNTHETICSASSERTIONTYPE_BODY)
	targetJSONPath.SetTarget(
		datadog.SyntheticsAssertionJSONPathTargetTarget{
			JsonPath: datadog.PtrString("topKey"),
			Operator: datadog.PtrString("isNot"),
			TargetValue: &targetValue0,
		},
	)

	return datadog.SyntheticsTestDetails{
		Config: &datadog.SyntheticsTestConfig{
			Assertions: []datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertionTextHTML),
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertion2000),
				datadog.SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion(targetJSONPath),
			},
			Request: datadog.SyntheticsTestRequest{
				Headers: &map[string]string{"testingGoClient": "true"},
				Method:  datadog.HTTPMETHOD_GET.Ptr(),
				Timeout: datadog.PtrFloat64(10),
				Url:     datadog.PtrString("https://datadoghq.com"),
			},
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   datadog.PtrString("Go client testing Synthetics API test - this is message"),
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(10),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
		Tags:    &[]string{"testing:api"},
		Type:    datadog.SYNTHETICSTESTDETAILSTYPE_API.Ptr(),
	}
}

func getTestSyntheticsSubtypeTcpAPI(ctx context.Context, t *testing.T) datadog.SyntheticsTestDetails {
	assertion2000 := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME)
	assertion2000.SetTarget(target2000)

	return datadog.SyntheticsTestDetails{
		Config: &datadog.SyntheticsTestConfig{
			Assertions: []datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertion2000),
			},
			Request: datadog.SyntheticsTestRequest{
				Host: datadog.PtrString("agent-intake.logs.datadoghq.com"),
				Port: datadog.PtrInt64(443),
			},
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   datadog.PtrString("Go client testing Synthetics API test Subtype TCP - this is message"),
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_TCP.Ptr(),
		Tags:    &[]string{"testing:api-tcp"},
		Type:    datadog.SYNTHETICSTESTDETAILSTYPE_API.Ptr(),
	}
}

func getTestSyntheticsBrowser(ctx context.Context, t *testing.T) datadog.SyntheticsTestDetails {
	return datadog.SyntheticsTestDetails{
		Config: &datadog.SyntheticsTestConfig{
			Assertions: []datadog.SyntheticsAssertion{},
			Request: datadog.SyntheticsTestRequest{
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Url:    datadog.PtrString("https://datadoghq.com"),
			},
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   datadog.PtrString("Go client testing Synthetics Browser test - this is message"),
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			DeviceIds:          &[]datadog.SyntheticsDeviceID{datadog.SYNTHETICSDEVICEID_TABLET},
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(10),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_FIVE_MINUTES.Ptr(),
		},
		Tags: &[]string{"testing:browser"},
		Type: datadog.SYNTHETICSTESTDETAILSTYPE_BROWSER.Ptr(),
	}
}

func TestSyntheticsAPITestLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateTest(ctx, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	synt, httpresp, err = Client(ctx).SyntheticsApi.GetTest(ctx, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	config := synt.GetConfig()

	assert.Equal(3, len(config.GetAssertions()))

	for _, assertion := range config.GetAssertions() {
		if assertion.SyntheticsAssertionJSONPathTarget != nil {
			assert.Equal(datadog.SYNTHETICSASSERTIONTYPE_BODY, assertion.SyntheticsAssertionJSONPathTarget.Type)
			assert.Equal(datadog.SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH, assertion.SyntheticsAssertionJSONPathTarget.Operator)
			target := assertion.SyntheticsAssertionJSONPathTarget.GetTarget()
			assert.Equal("0", target.GetTargetValue().(string))
			assert.Equal("isNot", target.GetOperator())
			assert.Equal("topKey", target.GetJsonPath())
		} else if assertion.SyntheticsAssertionTarget != nil {
			if assertion.SyntheticsAssertionTarget.Type == datadog.SYNTHETICSASSERTIONTYPE_HEADER {
				assert.Equal(datadog.SYNTHETICSASSERTIONOPERATOR_IS, assertion.SyntheticsAssertionTarget.Operator)
				assert.Equal("content-type", assertion.SyntheticsAssertionTarget.GetProperty())
				assert.Equal("text/html", assertion.SyntheticsAssertionTarget.GetTarget().(string))
			} else if assertion.SyntheticsAssertionTarget.Type == datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME {
				assert.Equal(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, assertion.SyntheticsAssertionTarget.Operator)
				assert.Equal(float64(2000), assertion.SyntheticsAssertionTarget.GetTarget().(float64))
			} else {
				assert.Fail("Unexpected type")
			}
		} else {
			assert.Fail("Unexpected target")
		}
	}

	// NOTE: API tests are started by default, so we have to stop it first
	// Stop API test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, publicID).
		FromTs(0).
		ToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		ProbeDc(locs).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// API tests sometimes have a delay before getting first result, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test recent response: %v", err)
	}

	// Delete API test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsSubtypeTcpAPITestLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsSubtypeTcpAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateTest(ctx, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	synt, httpresp, err = Client(ctx).SyntheticsApi.GetTest(ctx, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	config := synt.GetConfig()

	assert.Equal(1, len(config.GetAssertions()))

	for _, assertion := range config.GetAssertions() {
		if assertion.SyntheticsAssertionTarget.Type == datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME {
			assert.Equal(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, assertion.SyntheticsAssertionTarget.Operator)
			assert.Equal(float64(2000), assertion.SyntheticsAssertionTarget.GetTarget().(float64))
		} else {
			assert.Fail("Unexpected type")
		}
	}

	// NOTE: API tests are started by default, so we have to stop it first
	// Stop API test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, publicID).
		FromTs(0).
		ToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		ProbeDc(locs).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// API tests sometimes have a delay before getting first result, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_subtype_tcp_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test recent response: %v", err)
	}

	// Delete API test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsBrowserTestLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create Browser test
	testSyntheticsBrowser := getTestSyntheticsBrowser(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsBrowser.GetName(), synt.GetName())

	// Update Browser test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsBrowser.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateTest(ctx, publicID).Body(synt).Execute()
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get Browser test
	synt, httpresp, err = Client(ctx).SyntheticsApi.GetTest(ctx, publicID).Execute()
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// NOTE: Browser tests are paused by default, so we have to run it first
	// Start Browser test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Stop Browser test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID).
		Body(datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus}).Execute()
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent Browser test results
	var latestResults datadog.SyntheticsGetBrowserTestLatestResultsResponse
	locs := synt.GetLocations()
	latestResults, httpresp, err = Client(ctx).SyntheticsApi.GetBrowserTestLatestResults(ctx, publicID).
		FromTs(0).
		ToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		ProbeDc(locs).
		Execute()
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Empty(latestResults.GetResults())
	// Browser tests are asynchronous and take some time to run, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/browser_test_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test recent response: %v", err)
	}

	// Delete Browser test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}}).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsGetBrowserTestResult(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Test that the get browser result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsBrowserTestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/browser_test_single_result.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err := json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics browser test single response: %v", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsBrowserTestResultFullService.GetBrowserTestResult")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/synthetics/tests/browser/test-synthetics-id/results/test-result-id").
		Reply(200).
		JSON(data)
	defer gock.Off()

	unitAPI := Client(ctx).SyntheticsApi
	browserResp, httpresp, err := unitAPI.GetBrowserTestResult(ctx, "test-synthetics-id", "test-result-id").Execute()
	if err != nil {
		t.Errorf("Failed to get synthetics browser test result: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal("5140738909114888212", browserResp.GetResultId())
	assert.Equal(datadog.SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED, browserResp.GetStatus())
	assert.Equal(float64(1579711893111), browserResp.GetCheckTime())
	assert.Equal(int64(5), browserResp.GetCheckVersion())
}

func TestSyntheticsGetApiTestResult(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Test that the get api result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsAPITestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_single_result.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err := json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test single response: %v", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsBrowserTestResultFullService.GetAPITestResult")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/synthetics/tests/test-synthetics-id/results/test-result-id").
		Reply(200).
		JSON(data)
	defer gock.Off()

	unitAPI := Client(ctx).SyntheticsApi
	apiResp, httpresp, err := unitAPI.GetAPITestResult(ctx, "test-synthetics-id", "test-result-id").Execute()
	if err != nil {
		t.Errorf("Failed to get synthetics api test result: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(datadog.SYNTHETICSTESTMONITORSTATUS_TRIGGERED, apiResp.GetStatus())
	assert.Equal(float64(1580204310361), apiResp.GetCheckTime())
	assert.Equal(int64(2), apiResp.GetCheckVersion())
	assert.Equal("7761116396307201795", apiResp.GetResultId())
	apiResult := apiResp.GetResult()
	assert.Equal(datadog.SYNTHETICSTESTPROCESSSTATUS_FINISHED, apiResult.GetEventType())
}

func TestSyntheticsGetSubtypeTcpApiTestResult(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Test that the get api result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsAPITestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_subtype_tcp_single_result.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err := json.Unmarshal(data, &singleResult)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test single response: %v", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SyntheticsBrowserTestResultFullService.GetAPITestResult")
	assert.NoError(err)
	gock.New(URL).
		Get("/api/v1/synthetics/tests/test-synthetics-id/results/test-result-id").
		Reply(200).
		JSON(data)
	defer gock.Off()

	unitAPI := Client(ctx).SyntheticsApi
	apiResp, httpresp, err := unitAPI.GetAPITestResult(ctx, "test-synthetics-id", "test-result-id").Execute()
	if err != nil {
		t.Errorf("Failed to get synthetics api test result: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(datadog.SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED, apiResp.GetStatus())
	assert.Equal(float64(1594759676042), apiResp.GetCheckTime())
	assert.Equal(int64(2), apiResp.GetCheckVersion())
	assert.Equal("8846149531307597251", apiResp.GetResultId())
	apiResult := apiResp.GetResult()
	assert.Equal(datadog.SYNTHETICSTESTPROCESSSTATUS_FINISHED, apiResult.GetEventType())
}

func TestSyntheticsMultipleTestsOperations(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	var syntAPI, syntBrowser datadog.SyntheticsTestDetails
	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicIDAPI)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), syntAPI.GetName())

	// Create Browser test
	testSyntheticsBrowser := getTestSyntheticsBrowser(ctx, t)
	syntBrowser, httpresp, err = Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsBrowser).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDBrowser := syntBrowser.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicIDBrowser)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsBrowser.GetName(), syntBrowser.GetName())

	// Test Getting multiple tests
	var allTests datadog.SyntheticsListTestsResponse
	allTests, httpresp, err = Client(ctx).SyntheticsApi.ListTests(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics tests: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	td := allTests.GetTests()
	assert.NoError(isPublicIDPresent(publicIDAPI, td))
	assert.NoError(isPublicIDPresent(publicIDBrowser, td))
}

func TestSyntheticsDeleteTestErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.SyntheticsDeleteTestsPayload
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.SyntheticsDeleteTestsPayload{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.SyntheticsDeleteTestsPayload{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsDeleteTest404Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because not sure how to trigger the 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/synthetics/tests/delete").Reply(404).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx).Body(datadog.SyntheticsDeleteTestsPayload{}).Execute()
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSyntheticsUpdateStatusTestErrors(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, _, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicIDAPI)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, publicIDAPI, 400},
		"403 Forbidden":   {WithFakeAuth, "id", 403},
		"404 Not Found":   {WithTestAuth, "aaa-aaa-aaa", 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, tc.ID).Body(datadog.SyntheticsUpdateTestPauseStatusPayload{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserResultsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, "id", 403},
		"404 Not Found": {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.GetBrowserTestLatestResults(ctx, tc.ID).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPIResultsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, "id", 403},
		"404 Not Found": {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, tc.ID).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserSpecificResultErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, "id", 403},
		"404 Not Found": {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.GetBrowserTestResult(ctx, tc.ID, "resultid").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPISpecificResultErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, "id", 403},
		"404 Not Found": {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.GetAPITestResult(ctx, tc.ID, "resultid").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetTestErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, "id", 403},
		"404 Not Found": {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, tc.ID).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsUpdateTestErrors(t *testing.T) {
	// Setup the Client we'll use to interact with the Test account
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, _, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(testSyntheticsAPI).Execute()
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, publicIDAPI)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ID                 string
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, publicIDAPI, 400},
		"403 Forbidden":   {WithFakeAuth, "id", 403},
		"404 Not Found":   {WithTestAuth, "aaa-aaa-aaa", 404}, // TODO assert not exists
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.UpdateTest(ctx, tc.ID).Body(datadog.SyntheticsTestDetails{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTestErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		// Mock 404 because it's returned when synthetics is not enabled for the org
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.ListTests(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTest404Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mock 404 because it's returned when synthetics is not enabled for the org
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Get("/api/v1/synthetics/test").Reply(404).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SyntheticsApi.ListTests(ctx).Execute()
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSyntheticsCreateTestErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(datadog.SyntheticsTestDetails{}).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsCreateTest402Error(t *testing.T) {
	ctx, finish := WithClient(WithFakeAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	res, err := tests.ReadFixture("fixtures/synthetics/error_402.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Triggered when the synthetics test quota is reached. Need to mock.
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
	assert.NoError(err)
	gock.New(URL).Post("/api/v1/synthetics/test").Reply(402).JSON(res)
	defer gock.Off()

	_, httpresp, err := Client(ctx).SyntheticsApi.CreateTest(ctx).Body(datadog.SyntheticsTestDetails{}).Execute()
	assert.Equal(402, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func isPublicIDPresent(publicID string, syntTests []datadog.SyntheticsTestDetails) error {
	for _, st := range syntTests {
		if st.GetPublicId() == publicID {
			return nil
		}
	}
	return fmt.Errorf("Synthetics tests %s expected but not found", publicID)
}

func deleteSyntheticsTestIfExists(ctx context.Context, testID string) {
	_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx).
		Body(datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{testID}}).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Deleting synthetics test %s failed with %v, Another test may have already deleted this entity: %v",
			testID, httpresp.StatusCode, err)
	}
}

func TestSyntheticsListLocations(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	locs, httpresp, err := Client(ctx).SyntheticsApi.ListLocations(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all Synthetics locations: Response %s: %v", err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.Greater(len(locs.GetLocations()), 0)
}

func TestSyntheticsVariableLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	variable := datadog.SyntheticsGlobalVariable{
		Name: tests.UniqueEntityName(ctx, t),
		Description: datadog.PtrString("variable description"),
		Tags: &[]string{"synthetics"},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(false),
			Value: datadog.PtrString("VARIABLE_VALUE"),
		},
	}

	// Create variable
	result, httpresp, err := Client(ctx).SyntheticsApi.CreateGlobalVariable(ctx).Body(variable).Execute()

	if err != nil {
		t.Fatalf("Error creating Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), variable.GetName())

	// Edit variable
	updatedName := fmt.Sprintf("%s-updated", variable.GetName())
	variable.SetName(updatedName)

	result, httpresp, err = Client(ctx).SyntheticsApi.EditGlobalVariable(ctx, result.GetId()).Body(variable).Execute()

	if err != nil {
		t.Fatalf("Error editing Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), updatedName)

	// Delete variable
	httpresp, err = Client(ctx).SyntheticsApi.DeleteGlobalVariable(ctx, result.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting Synthetics global variable %s: Response %s: %v", result.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}
