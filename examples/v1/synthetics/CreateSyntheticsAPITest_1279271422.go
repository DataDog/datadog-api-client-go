// Create an API test with multi subtype returns "OK - Returns the created test details." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.SyntheticsAPITest{
		Config: datadogV1.SyntheticsAPITestConfig{
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Steps: []datadogV1.SyntheticsAPIStep{
				{
					AllowFailure: datadog.PtrBool(true),
					Assertions: []datadogV1.SyntheticsAssertion{
						datadogV1.SyntheticsAssertion{
							SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
								Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
								Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
								Target:   200,
							}},
					},
					ExtractedValues: []datadogV1.SyntheticsParsingOptions{
						{
							Field: datadog.PtrString("server"),
							Name:  datadog.PtrString("EXTRACTED_VALUE"),
							Parser: &datadogV1.SyntheticsVariableParser{
								Type: datadogV1.SYNTHETICSGLOBALVARIABLEPARSERTYPE_RAW,
							},
							Type: datadogV1.SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER.Ptr(),
						},
					},
					IsCritical: datadog.PtrBool(true),
					Name:       "request is sent",
					Request: datadogV1.SyntheticsTestRequest{
						Method:  datadog.PtrString("GET"),
						Timeout: datadog.PtrFloat64(10),
						Url:     datadog.PtrString("https://datadoghq.com"),
					},
					Retry: &datadogV1.SyntheticsTestOptionsRetry{
						Count:    datadog.PtrInt64(5),
						Interval: datadog.PtrFloat64(1000),
					},
					Subtype: datadogV1.SYNTHETICSAPISTEPSUBTYPE_HTTP,
				},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_multi_step_payload.json",
		Name:    "Example-Create_an_API_test_with_multi_subtype_returns_OK_Returns_the_created_test_details_response",
		Options: datadogV1.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Create_an_API_test_with_multi_subtype_returns_OK_Returns_the_created_test_details_response"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(1000),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadogV1.SYNTHETICSAPITESTTYPE_API,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsAPITest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsAPITest`:\n%s\n", responseContent)
}
