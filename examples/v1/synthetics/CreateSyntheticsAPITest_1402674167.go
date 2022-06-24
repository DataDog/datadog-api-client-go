// Create an API GRPC test returns "OK - Returns the created test details." response

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
		Config: datadog.SyntheticsAPITestConfig{
			Assertions: []datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SYNTHETICSASSERTIONOPERATOR_IS,
						Target:   1,
						Type:     datadog.SYNTHETICSASSERTIONTYPE_GRPC_HEALTHCHECK_STATUS,
					}},
			},
			Request: &datadog.SyntheticsTestRequest{
				Host:     datadog.PtrString("localhost"),
				Port:     datadog.PtrInt64(50051),
				Service:  datadog.PtrString("Hello"),
				Method:   datadog.HTTPMETHOD_GET.Ptr(),
				Message:  datadog.PtrString(""),
				Metadata: map[string]string{},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_grpc_test_payload.json",
		Name:    "Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response",
		Options: datadog.SyntheticsTestOptions{
			MinFailureDuration: datadog.PtrInt64(0),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorOptions: &datadog.SyntheticsTestOptionsMonitorOptions{
				RenotifyInterval: datadog.PtrInt64(0),
			},
			MonitorName: datadog.PtrString("Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response"),
			TickEvery:   datadog.PtrInt64(60),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_GRPC.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadog.SYNTHETICSAPITESTTYPE_API,
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
