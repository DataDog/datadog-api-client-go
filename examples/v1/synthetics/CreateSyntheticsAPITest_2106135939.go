// Create a multistep test with subtest returns "OK" response

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
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestPublicID := os.Getenv("SYNTHETICS_API_TEST_PUBLIC_ID")

	body := datadogV1.SyntheticsAPITest{
		Config: datadogV1.SyntheticsAPITestConfig{
			Steps: []datadogV1.SyntheticsAPIStep{
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target: datadogV1.SyntheticsAssertionTargetValue{
										SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(200)},
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthWeb: &datadogV1.SyntheticsBasicAuthWeb{
									Password: datadog.PtrString("password"),
									Username: datadog.PtrString("username"),
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPISubtestStep: &datadogV1.SyntheticsAPISubtestStep{
						Subtype:         datadogV1.SYNTHETICSAPISUBTESTSTEPSUBTYPE_PLAY_SUB_TEST,
						SubtestPublicId: SyntheticsAPITestPublicID,
						Name:            "subtest step",
					}},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_multi_step_with_subtest.json",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Type:    datadogV1.SYNTHETICSAPITESTTYPE_API,
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
