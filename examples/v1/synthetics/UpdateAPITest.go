// Edit an API test returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "synthetics_api_test" in the system
	SYNTHETICS_API_TEST_PUBLIC_ID := os.Getenv("SYNTHETICS_API_TEST_PUBLIC_ID")

	body := datadog.SyntheticsAPITest{
		Config: &datadog.SyntheticsAPITestConfig{
			Assertions: &[]datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SyntheticsAssertionOperator("is"),
						Property: datadog.PtrString("{{ PROPERTY }}"),
						Target:   "text/html",
						Type:     datadog.SyntheticsAssertionType("header"),
					}},
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SyntheticsAssertionOperator("lessThan"),
						Target:   2000,
						Type:     datadog.SyntheticsAssertionType("responseTime"),
					}},
				datadog.SyntheticsAssertion{
					SyntheticsAssertionJSONPathTarget: &datadog.SyntheticsAssertionJSONPathTarget{
						Operator: datadog.SyntheticsAssertionJSONPathOperator("validatesJSONPath"),
						Target: &datadog.SyntheticsAssertionJSONPathTargetTarget{
							JsonPath:    datadog.PtrString("topKey"),
							Operator:    datadog.PtrString("isNot"),
							TargetValue: "0",
						},
						Type: datadog.SyntheticsAssertionType("body"),
					}},
			},
			ConfigVariables: &[]datadog.SyntheticsConfigVariable{
				datadog.SyntheticsConfigVariable{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadog.SyntheticsConfigVariableType("text"),
				},
			},
			Request: &datadog.SyntheticsTestRequest{
				Certificate: &datadog.SyntheticsTestRequestCertificate{
					Cert: &datadog.SyntheticsTestRequestCertificateItem{
						Filename:  datadog.PtrString("cert-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
					Key: &datadog.SyntheticsTestRequestCertificateItem{
						Filename:  datadog.PtrString("key-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
				},
				Headers: &map[string]string{
					"unique": "exampleeditanapitestreturnsokresponse",
				},
				Method:  datadog.HTTPMethod("GET").Ptr(),
				Timeout: datadog.PtrFloat64(10),
				Url:     datadog.PtrString("https://datadoghq.com"),
			},
		},
		Locations: &[]string{
			"aws:us-east-2",
		},
		Message: datadog.PtrString("BDD test payload: synthetics_api_test_payload.json"),
		Name:    datadog.PtrString("Example-Edit_an_API_test_returns_OK_response-updated"),
		Options: &datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Test-TestSyntheticsAPITestLifecycle-1623076664"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Status:  datadog.SyntheticsTestPauseStatus("live").Ptr(),
		Subtype: datadog.SyntheticsTestDetailsSubType("http").Ptr(),
		Tags: &[]string{
			"testing:api",
		},
		Type: datadog.SyntheticsAPITestType("api").Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.UpdateAPITest(ctx, SYNTHETICS_API_TEST_PUBLIC_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateAPITest`:\n%s\n", responseContent)
}
