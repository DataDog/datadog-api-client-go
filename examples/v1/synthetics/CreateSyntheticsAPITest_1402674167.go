// Create an API GRPC test returns "OK - Returns the created test details." response

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
			Assertions: []datadogV1.SyntheticsAssertion{
				datadogV1.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
						Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
						Target:   1,
						Type:     datadogV1.SYNTHETICSASSERTIONTYPE_GRPC_HEALTHCHECK_STATUS,
					}},
			},
			Request: &datadogV1.SyntheticsTestRequest{
				Host:     datadog.PtrString("localhost"),
				Port:     datadog.PtrInt64(50051),
				Service:  datadog.PtrString("Hello"),
				Method:   datadog.PtrString("GET"),
				Message:  datadog.PtrString(""),
				Metadata: map[string]string{},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_grpc_test_payload.json",
		Name:    "Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response",
		Options: datadogV1.SyntheticsTestOptions{
			MinFailureDuration: datadog.PtrInt64(0),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorOptions: &datadogV1.SyntheticsTestOptionsMonitorOptions{
				RenotifyInterval: datadog.PtrInt64(0),
			},
			MonitorName: datadog.PtrString("Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response"),
			TickEvery:   datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_GRPC.Ptr(),
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
