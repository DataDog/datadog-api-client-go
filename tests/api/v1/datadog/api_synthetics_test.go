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
	"path/filepath"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

var targetTextHTML interface{} = "text/html"
var target2000 interface{} = 2000
var targetValue0 interface{} = "0"

func getTestSyntheticsAPI(ctx context.Context, t *testing.T) datadog.SyntheticsAPITest {
	assertionTextHTML := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_IS, datadog.SYNTHETICSASSERTIONTYPE_HEADER)
	assertionTextHTML.Property = datadog.PtrString("{{ PROPERTY }}")
	assertionTextHTML.Target = &targetTextHTML

	assertion2000 := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME)
	assertion2000.Target = &target2000

	targetJSONPath := datadog.NewSyntheticsAssertionJSONPathTarget(datadog.SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH, datadog.SYNTHETICSASSERTIONTYPE_BODY)
	targetJSONPath.SetTarget(
		datadog.SyntheticsAssertionJSONPathTargetTarget{
			JsonPath:    datadog.PtrString("topKey"),
			Operator:    datadog.PtrString("isNot"),
			TargetValue: &targetValue0,
		},
	)

	return datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			Assertions: &[]datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertionTextHTML),
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertion2000),
				datadog.SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion(targetJSONPath),
			},
			ConfigVariables: &[]datadog.SyntheticsConfigVariable{
				datadog.SyntheticsConfigVariable{
					Name:    "PROPERTY",
					Example: datadog.PtrString("content-type"),
					Pattern: datadog.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: &datadog.SyntheticsTestRequest{
				Headers: &map[string]string{"testingGoClient": "true"},
				Method:  datadog.HTTPMETHOD_GET.Ptr(),
				Timeout: datadog.PtrFloat64(10),
				Url:     datadog.PtrString("https://datadoghq.com"),
				Certificate: &datadog.SyntheticsTestRequestCertificate{
					Cert: &datadog.SyntheticsTestRequestCertificateItem{
						Content:   datadog.PtrString("cert-content"),
						Filename:  datadog.PtrString("cert-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
					Key: &datadog.SyntheticsTestRequestCertificateItem{
						Content:   datadog.PtrString("key-content"),
						Filename:  datadog.PtrString("key-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
				},
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
			MinLocationFailed:  datadog.PtrInt64(1),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
			MonitorName: tests.UniqueEntityName(ctx, t),
			MonitorPriority: datadog.PtrInt32(5),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
		Tags:    &[]string{"testing:api"},
		Type:    datadog.SYNTHETICSAPITESTTYPE_API.Ptr(),
	}
}

func getTestSyntheticsAPIMultistep(ctx context.Context, t *testing.T, globalVariable datadog.SyntheticsGlobalVariable) datadog.SyntheticsAPITest {
	return datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			ConfigVariables: &[]datadog.SyntheticsConfigVariable{
				datadog.SyntheticsConfigVariable{
					Name:    "PROPERTY",
					Example: datadog.PtrString("content-type"),
					Pattern: datadog.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
				datadog.SyntheticsConfigVariable{
					Name: "VARIABLE_NAME",
					Id:   datadog.PtrString(globalVariable.GetId()),
					Type: datadog.SYNTHETICSCONFIGVARIABLETYPE_GLOBAL,
				},
			},
			Steps: &[]datadog.SyntheticsAPIStep{
				datadog.SyntheticsAPIStep{
					Assertions: &[]datadog.SyntheticsAssertion{},
					Name:       datadog.PtrString("First step"),
					Request: &datadog.SyntheticsTestRequest{
						Headers: &map[string]string{"testingGoClient": "true"},
						Method:  datadog.HTTPMETHOD_GET.Ptr(),
						Timeout: datadog.PtrFloat64(10),
						Url:     datadog.PtrString("https://datadoghq.com"),
					},
					Subtype: datadog.SYNTHETICSAPISTEPSUBTYPE_HTTP.Ptr(),
					ExtractedValues: &[]datadog.SyntheticsParsingOptions{
						datadog.SyntheticsParsingOptions{
							Name: datadog.PtrString("EXTRACTED_VALUE"),
							Parser: &datadog.SyntheticsVariableParser{
								Type: datadog.SYNTHETICSGLOBALVARIABLEPARSERTYPE_RAW,
							},
							Type:  datadog.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER.Ptr(),
							Field: datadog.PtrString("content-type"),
						},
					},
					AllowFailure: datadog.PtrBool(true),
					IsCritical:   datadog.PtrBool(false),
				},
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
			MinLocationFailed:  datadog.PtrInt64(1),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Tags:    &[]string{"testing:api"},
		Type:    datadog.SYNTHETICSAPITESTTYPE_API.Ptr(),
	}
}

func getTestSyntheticsSubtypeTCPAPI(ctx context.Context, t *testing.T) datadog.SyntheticsAPITest {
	assertion2000 := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME)
	assertion2000.SetTarget(target2000)

	return datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			Assertions: &[]datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(assertion2000),
			},
			Request: &datadog.SyntheticsTestRequest{
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
		Type:    datadog.SYNTHETICSAPITESTTYPE_API.Ptr(),
	}
}

func getTestSyntheticsSubtypeDNSAPI(ctx context.Context, t *testing.T) datadog.SyntheticsAPITest {
	recordAssertion := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_IS, datadog.SYNTHETICSASSERTIONTYPE_RECORD_SOME)
	var target interface{} = "0.0.0.0"
	recordAssertion.SetProperty("A")
	recordAssertion.SetTarget(target)

	return datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			Assertions: &[]datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(recordAssertion),
			},
			Request: &datadog.SyntheticsTestRequest{
				Host:          datadog.PtrString("https://www.datadoghq.com"),
				DnsServer:     datadog.PtrString("8.8.8.8"),
				DnsServerPort: datadog.PtrInt32(53),
			},
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   datadog.PtrString("Go client testing Synthetics API test Subtype DNS - this is message"),
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_DNS.Ptr(),
		Tags:    &[]string{"testing:api-dns"},
		Type:    datadog.SYNTHETICSAPITESTTYPE_API.Ptr(),
	}
}

func getTestSyntheticsSubtypeICMPAPI(ctx context.Context, t *testing.T) datadog.SyntheticsAPITest {
	latencyAssertion := datadog.NewSyntheticsAssertionTarget(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, datadog.SYNTHETICSASSERTIONTYPE_LATENCY)
	var target interface{} = 200
	latencyAssertion.SetProperty("avg")
	latencyAssertion.SetTarget(target)

	return datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			Assertions: &[]datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertionTargetAsSyntheticsAssertion(latencyAssertion),
			},
			Request: &datadog.SyntheticsTestRequest{
				Host:            datadog.PtrString("www.datadoghq.com"),
				NumberOfPackets: datadog.PtrInt32(2),
			},
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   datadog.PtrString("Go client testing Synthetics API test Subtype ICMP - this is message"),
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			TickEvery: datadog.SYNTHETICSTICKINTERVAL_MINUTE.Ptr(),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_ICMP.Ptr(),
		Tags:    &[]string{"testing:api-icmp"},
		Type:    datadog.SYNTHETICSAPITESTTYPE_API.Ptr(),
	}
}

func getTestSyntheticsBrowser(ctx context.Context, t *testing.T) datadog.SyntheticsBrowserTest {
	return datadog.SyntheticsBrowserTest{
		Config: &datadog.SyntheticsBrowserTestConfig{
			Assertions: []datadog.SyntheticsAssertion{},
			Request: datadog.SyntheticsTestRequest{
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Url:    datadog.PtrString("https://datadoghq.com"),
			},
			SetCookie: datadog.PtrString("name:test"),
		},
		Locations: &[]string{"aws:us-east-2"},
		Message:   "Go client testing Synthetics Browser test - this is message",
		Name:      tests.UniqueEntityName(ctx, t),
		Options: &datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			DeviceIds:          &[]datadog.SyntheticsDeviceID{datadog.SYNTHETICSDEVICEID_TABLET},
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery:    datadog.SYNTHETICSTICKINTERVAL_FIVE_MINUTES.Ptr(),
			NoScreenshot: datadog.PtrBool(true),
			DisableCors:  datadog.PtrBool(true),
		},
		Tags: &[]string{"testing:browser"},
		Type: datadog.SYNTHETICSBROWSERTESTTYPE_BROWSER.Ptr(),
	}
}

func TestSyntheticsAPITestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics API test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateAPITest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics API test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	syntTest, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, syntTest.GetName())

	config := syntTest.GetConfig()

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
				assert.Equal("{{ PROPERTY }}", assertion.SyntheticsAssertionTarget.GetProperty())
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
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, publicID, *datadog.NewGetAPITestLatestResultsOptionalParameters().
		WithFromTs(0).
		WithToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		WithProbeDc(locs))
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
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsSubtypeTcpAPITestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsSubtypeTCPAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateAPITest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	syntTest, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, syntTest.GetName())

	config := syntTest.GetConfig()

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
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, publicID, *datadog.NewGetAPITestLatestResultsOptionalParameters().
		WithFromTs(0).
		WithToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		WithProbeDc(locs))
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
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsSubtypeDnsAPITestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsSubtypeDNSAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateAPITest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	syntTest, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, syntTest.GetName())

	config := syntTest.GetConfig()

	assert.Equal(1, len(config.GetAssertions()))

	for _, assertion := range config.GetAssertions() {
		if assertion.SyntheticsAssertionTarget.Type == datadog.SYNTHETICSASSERTIONTYPE_RECORD_SOME {
			assert.Equal(datadog.SYNTHETICSASSERTIONOPERATOR_IS, assertion.SyntheticsAssertionTarget.Operator)
			assert.Equal("0.0.0.0", assertion.SyntheticsAssertionTarget.GetTarget().(string))
		} else {
			assert.Fail("Unexpected type")
		}
	}

	// NOTE: API tests are started by default, so we have to stop it first
	// Stop API test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Start API test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent API test results
	var latestResults datadog.SyntheticsGetAPITestLatestResultsResponse
	locs := synt.GetLocations()
	_, httpresp, err = Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, publicID, *datadog.NewGetAPITestLatestResultsOptionalParameters().
		WithFromTs(0).
		WithToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		WithProbeDc(locs))
	if err != nil {
		t.Fatalf("Error getting latest results for Synthetics test %s: Response %s: %v",
			publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	// API tests sometimes have a delay before getting first result, so we use a mock response to verify
	// that deserialization works properly
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_subtype_dns_recent_results.json")
	data, _ := ioutil.ReadFile(fixturePath)
	err = json.Unmarshal(data, &latestResults)
	if err != nil {
		t.Fatalf("Failed deserializing Synthetics API test recent response: %v", err)
	}

	// Delete API test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsSubtypeIcmpAPITestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsSubtypeICMPAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateAPITest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get API test
	synt, httpresp, err = Client(ctx).SyntheticsApi.GetAPITest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	config := synt.GetConfig()

	assert.Equal(1, len(config.GetAssertions()))

	for _, assertion := range config.GetAssertions() {
		if assertion.SyntheticsAssertionTarget.Type == datadog.SYNTHETICSASSERTIONTYPE_LATENCY {
			assert.Equal(datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN, assertion.SyntheticsAssertionTarget.Operator)
			assert.Equal(float64(200), assertion.SyntheticsAssertionTarget.GetTarget())
			assert.Equal("avg", assertion.SyntheticsAssertionTarget.GetProperty())
		} else {
			assert.Fail("Unexpected type")
		}
	}

	// Delete API test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsBrowserTestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create Browser test
	testSyntheticsBrowser := getTestSyntheticsBrowser(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsBrowserTest(ctx, testSyntheticsBrowser)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsBrowser.GetName(), synt.GetName())

	// Update Browser test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsBrowser.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateBrowserTest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get Browser test
	syntTest, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, syntTest.GetName())

	// NOTE: Browser tests are paused by default, so we have to run it first
	// Start Browser test
	var pauseStatus bool
	newStatus := datadog.SYNTHETICSTESTPAUSESTATUS_LIVE
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s live: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Stop Browser test
	newStatus = datadog.SYNTHETICSTESTPAUSESTATUS_PAUSED
	pauseStatus, httpresp, err = Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, publicID, datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: &newStatus})
	if err != nil {
		t.Fatalf("Error making Synthetics test %s paused: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(true, pauseStatus)

	// Get the most recent Browser test results
	var latestResults datadog.SyntheticsGetBrowserTestLatestResultsResponse
	locs := synt.GetLocations()
	latestResults, httpresp, err = Client(ctx).SyntheticsApi.GetBrowserTestLatestResults(ctx, publicID, *datadog.NewGetBrowserTestLatestResultsOptionalParameters().
		WithFromTs(0).
		WithToTs(tests.ClockFromContext(ctx).Now().Unix() * 1000).
		WithProbeDc(locs))
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
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsGetBrowserTestResult(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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
	browserResp, httpresp, err := unitAPI.GetBrowserTestResult(ctx, "test-synthetics-id", "test-result-id")
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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
	apiResp, httpresp, err := unitAPI.GetAPITestResult(ctx, "test-synthetics-id", "test-result-id")
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
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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
	apiResp, httpresp, err := unitAPI.GetAPITestResult(ctx, "test-synthetics-id", "test-result-id")
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

func TestSyntheticsGetSubtypeDnsApiTestResult(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)

	// Test that the get api result test data can be properly unmarshalled and takes the expected elements in the path
	var singleResult datadog.SyntheticsAPITestResultFull
	fixturePath, _ := filepath.Abs("fixtures/synthetics/api_test_subtype_dns_single_result.json")
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
	apiResp, httpresp, err := unitAPI.GetAPITestResult(ctx, "test-synthetics-id", "test-result-id")
	if err != nil {
		t.Errorf("Failed to get synthetics api test result: %v", err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(datadog.SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED, apiResp.GetStatus())
	assert.Equal(float64(1597741443749), apiResp.GetCheckTime())
	assert.Equal(int64(2), apiResp.GetCheckVersion())
	assert.Equal("1006122437899386104", apiResp.GetResultId())
	apiResult := apiResp.GetResult()
	assert.Equal(datadog.SYNTHETICSTESTPROCESSSTATUS_FINISHED, apiResult.GetEventType())
}

func TestSyntheticsMultipleTestsOperations(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	var syntAPI datadog.SyntheticsAPITest
	var syntBrowser datadog.SyntheticsBrowserTest
	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicIDAPI)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), syntAPI.GetName())

	// Create Browser test
	testSyntheticsBrowser := getTestSyntheticsBrowser(ctx, t)
	syntBrowser, httpresp, err = Client(ctx).SyntheticsApi.CreateSyntheticsBrowserTest(ctx, testSyntheticsBrowser)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDBrowser := syntBrowser.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicIDBrowser)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsBrowser.GetName(), syntBrowser.GetName())

	// Test Getting multiple tests
	var allTests datadog.SyntheticsListTestsResponse
	allTests, httpresp, err = Client(ctx).SyntheticsApi.ListTests(ctx)
	if err != nil {
		t.Fatalf("Error getting all Synthetics tests: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	td := allTests.GetTests()
	assert.NoError(isTestPublicIDPresent(publicIDAPI, td))
	assert.NoError(isTestPublicIDPresent(publicIDBrowser, td))
}

func TestSyntheticsDeleteTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsDeleteTest404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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

	_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{})
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSyntheticsUpdateStatusTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, _, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicIDAPI)

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

			_, httpresp, err := Client(ctx).SyntheticsApi.UpdateTestPauseStatus(ctx, tc.ID, datadog.SyntheticsUpdateTestPauseStatusPayload{})
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserResultsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.GetBrowserTestLatestResults(ctx, tc.ID)
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPIResultsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.GetAPITestLatestResults(ctx, tc.ID)
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsBrowserSpecificResultErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.GetBrowserTestResult(ctx, tc.ID, "resultid")
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetAPISpecificResultErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.GetAPITestResult(ctx, tc.ID, "resultid")
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsGetTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.GetTest(ctx, tc.ID)
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsUpdateTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	// Setup the Client we'll use to interact with the Test account
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	syntAPI, _, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicIDAPI := syntAPI.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicIDAPI)

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

			_, httpresp, err := Client(ctx).SyntheticsApi.UpdateAPITest(ctx, tc.ID, datadog.SyntheticsAPITest{})
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.ListTests(ctx)
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsListTest404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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

	_, httpresp, err := Client(ctx).SyntheticsApi.ListTests(ctx)
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSyntheticsCreateTestErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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

			_, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, datadog.SyntheticsAPITest{})
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSyntheticsCreateTest402Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
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

	_, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, datadog.SyntheticsAPITest{})
	assert.Equal(402, httpresp.StatusCode)
	apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func isTestPublicIDPresent(publicID string, syntTests []datadog.SyntheticsTestDetails) error {
	for _, st := range syntTests {
		if st.GetPublicId() == publicID {
			return nil
		}
	}
	return fmt.Errorf("Synthetics tests %s expected but not found", publicID)
}

func isGlobalVariableIdPresent(id string, syntGlobalVariables []datadog.SyntheticsGlobalVariable) error {
	for _, st := range syntGlobalVariables {
		if st.GetId() == id {
			return nil
		}
	}
	return fmt.Errorf("Synthetics global variable %s expected but not found", id)
}

func deleteSyntheticsTestIfExists(ctx context.Context, t *testing.T, testID string) {
	_, httpresp, err := Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{testID}})
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Deleting synthetics test %s failed with %v, Another test may have already deleted this entity: %v",
			testID, httpresp.StatusCode, err)
	}
}

func TestSyntheticsListLocations(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	locs, httpresp, err := Client(ctx).SyntheticsApi.ListLocations(ctx)
	if err != nil {
		t.Fatalf("Error getting all Synthetics locations: Response %s: %v", err.Error(), err)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.Greater(len(locs.GetLocations()), 0)
}

func TestSyntheticsVariableLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	variable := datadog.SyntheticsGlobalVariable{
		Name:        strings.Replace(strings.ToUpper(*tests.UniqueEntityName(ctx, t)), "-", "_", -1),
		Description: "variable description",
		Tags:        []string{"synthetics"},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(false),
			Value:  datadog.PtrString("VARIABLE_VALUE"),
		},
	}

	// Create variable
	result, httpresp, err := Client(ctx).SyntheticsApi.CreateGlobalVariable(ctx, variable)

	if err != nil {
		t.Fatalf("Error creating Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), variable.GetName())

	// Get variable
	result, httpresp, err = Client(ctx).SyntheticsApi.GetGlobalVariable(ctx, result.GetId())

	if err != nil {
		t.Fatalf("Error getting Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), variable.GetName())

	var allVariables datadog.SyntheticsListGlobalVariablesResponse
	allVariables, httpresp, err = Client(ctx).SyntheticsApi.ListGlobalVariables(ctx)

	if err != nil {
		t.Fatalf("Error getting all Synthetics global variable: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	variables := allVariables.GetVariables()
	assert.NoError(isGlobalVariableIdPresent(result.GetId(), variables))

	// Edit variable
	updatedName := fmt.Sprintf("%s_UPDATED", variable.GetName())
	variable.SetName(updatedName)

	result, httpresp, err = Client(ctx).SyntheticsApi.EditGlobalVariable(ctx, result.GetId(), variable)

	if err != nil {
		t.Fatalf("Error editing Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), updatedName)

	// Delete variable
	httpresp, err = Client(ctx).SyntheticsApi.DeleteGlobalVariable(ctx, result.GetId())
	if err != nil {
		t.Fatalf("Error deleting Synthetics global variable %s: Response %s: %v", result.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsVariableFromTestLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// create test to use
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)

	variable := datadog.SyntheticsGlobalVariable{
		Name:              strings.Replace(strings.ToUpper(*tests.UniqueEntityName(ctx, t)), "-", "_", -1),
		Description:       "variable from test description",
		ParseTestPublicId: datadog.PtrString(publicID),
		ParseTestOptions: &datadog.SyntheticsGlobalVariableParseTestOptions{
			Parser: datadog.SyntheticsVariableParser{
				Type: datadog.SYNTHETICSGLOBALVARIABLEPARSERTYPE_RAW,
			},
			Type:  datadog.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER,
			Field: datadog.PtrString("content-type"),
		},
		Tags: []string{"synthetics"},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(false),
			Value:  datadog.PtrString(""),
		},
	}

	// Create variable
	result, httpresp, err := Client(ctx).SyntheticsApi.CreateGlobalVariable(ctx, variable)

	if err != nil {
		t.Fatalf("Error creating Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), variable.GetName())

	// Get variable
	result, httpresp, err = Client(ctx).SyntheticsApi.GetGlobalVariable(ctx, result.GetId())

	if err != nil {
		t.Fatalf("Error getting Synthetics global variable %v: Response %s: %v", variable, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(result.GetName(), variable.GetName())

	// Delete variable
	httpresp, err = Client(ctx).SyntheticsApi.DeleteGlobalVariable(ctx, result.GetId())
	if err != nil {
		t.Fatalf("Error deleting Synthetics global variable %s: Response %s: %v", result.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
}

func TestSyntheticsTriggerCITests(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// create api test to trigger later
	testSyntheticsAPI := getTestSyntheticsAPI(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()

	// trigger the test
	test := datadog.SyntheticsCITest{
		Locations: &[]string{"aws:us-east-2"},
		PublicId:  publicID,
	}
	tests := []datadog.SyntheticsCITest{test}

	fullResult, httpresp, err := Client(ctx).SyntheticsApi.TriggerCITests(ctx, datadog.SyntheticsCITestBody{
		Tests: &tests,
	})
	if err != nil {
		t.Fatalf("Error triggering ci tests: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	results := fullResult.GetResults()
	triggeredCheckIds := fullResult.GetTriggeredCheckIds()
	assert.Equal(publicID, results[0].GetPublicId())
	assert.Equal(publicID, triggeredCheckIds[0])

	// delete the test
	_, httpresp, err = Client(ctx).SyntheticsApi.DeleteTests(ctx, datadog.SyntheticsDeleteTestsPayload{PublicIds: &[]string{publicID}})
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
}

func TestSyntheticsPrivateLocationLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// create a private location
	privateLocationRequest := datadog.SyntheticsPrivateLocation{
		Description: "Private Location description",
		Name:        *tests.UniqueEntityName(ctx, t),
		Tags:        []string{"testing:private-location"},
	}

	privateLocationFull, httpresp, err := Client(ctx).SyntheticsApi.CreatePrivateLocation(ctx, privateLocationRequest)
	if err != nil {
		t.Fatalf("Error creating Synthetics private location %v: Response %s: %v", privateLocationRequest, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	pl := privateLocationFull.GetPrivateLocation()

	assert.Equal(pl.GetName(), privateLocationRequest.GetName())

	privateLocationID := pl.GetId()

	// edit private location
	privateLocationUpdateRequest := datadog.SyntheticsPrivateLocation{
		Description: "Private Location description",
		Name:        fmt.Sprintf("%s-updated", pl.GetName()),
		Tags:        []string{"testing:private-location"},
	}

	privateLocation, httpresp, err := Client(ctx).SyntheticsApi.UpdatePrivateLocation(ctx, privateLocationID, privateLocationUpdateRequest)
	if err != nil {
		t.Fatalf("Error editing Synthetics private location %v: Response %s: %v", privateLocationRequest, err.(datadog.GenericOpenAPIError).Body(), err)
	}

	assert.Equal(privateLocation.GetName(), privateLocationUpdateRequest.GetName())

	// get private location
	privateLocationResponse, httpresp, err := Client(ctx).SyntheticsApi.GetPrivateLocation(ctx, privateLocationID)
	if err != nil {
		t.Fatalf("Error getting Synthetics private location %s: Response %s: %v", privateLocationID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(privateLocationResponse.GetName(), privateLocation.GetName())

	// delete private location
	httpresp, err = Client(ctx).SyntheticsApi.DeletePrivateLocation(ctx, privateLocationID)
	if err != nil {
		t.Fatalf("Error deleting Synthetics test %s: Response %s: %v", privateLocationID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)
}

func TestSyntheticsBrowserTestEndpointLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create Browser test
	testSyntheticsBrowser := getTestSyntheticsBrowser(ctx, t)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsBrowserTest(ctx, testSyntheticsBrowser)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsBrowser, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsBrowser.GetName(), synt.GetName())

	// Update Browser test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsBrowser.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateBrowserTest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())

	// Get Browser test
	synt, httpresp, err = Client(ctx).SyntheticsApi.GetBrowserTest(ctx, publicID)
	if err != nil {
		t.Fatalf("Error getting Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())
}

func TestSyntheticsAPIMultistepTestEndpointLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Create global variable for test config
	variable := datadog.SyntheticsGlobalVariable{
		Name:        strings.Replace(strings.ToUpper(*tests.UniqueEntityName(ctx, t)), "-", "_", -1),
		Description: "variable description",
		Tags:        []string{"synthetics"},
		Value: datadog.SyntheticsGlobalVariableValue{
			Secure: datadog.PtrBool(false),
			Value:  datadog.PtrString("VARIABLE_VALUE"),
		},
	}

	// Create variable
	globalVariable, httpresp, err := Client(ctx).SyntheticsApi.CreateGlobalVariable(ctx, variable)

	// Create API test
	testSyntheticsAPI := getTestSyntheticsAPIMultistep(ctx, t, globalVariable)
	synt, httpresp, err := Client(ctx).SyntheticsApi.CreateSyntheticsAPITest(ctx, testSyntheticsAPI)
	if err != nil {
		t.Fatalf("Error creating Synthetics test %v: Response %s: %v", testSyntheticsAPI, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	publicID := synt.GetPublicId()
	defer deleteSyntheticsTestIfExists(ctx, t, publicID)
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(testSyntheticsAPI.GetName(), synt.GetName())

	// Update API test
	updatedName := fmt.Sprintf("%s-updated", testSyntheticsAPI.GetName())
	synt.SetName(updatedName)
	// if we want to reuse the entity returned by the API, we must set these field to nil, as they can't be sent back
	synt.MonitorId = nil
	synt.PublicId = nil
	synt, httpresp, err = Client(ctx).SyntheticsApi.UpdateAPITest(ctx, publicID, synt)
	if err != nil {
		t.Fatalf("Error updating Synthetics test %s: Response %s: %v", publicID, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(updatedName, synt.GetName())
}
