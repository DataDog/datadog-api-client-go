// Create an API test returns "OK - Returns the created test details." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
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
				Headers: &map[string]string{
					"unique": "examplecreateanapitestreturnsokreturnsthecreatedtestdetailsresponse",
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
		Name:    datadog.PtrString("Example-Create_an_API_test_returns_OK_Returns_the_created_test_details_response"),
		Options: &datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Create_an_API_test_returns_OK_Returns_the_created_test_details_response"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadog.SyntheticsTestDetailsSubType("http").Ptr(),
		Tags: &[]string{
			"testing:api",
		},
		Type: datadog.SyntheticsAPITestType("api").Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.CreateSyntheticsAPITest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsAPITest`:\n%s\n", responseContent)
}
