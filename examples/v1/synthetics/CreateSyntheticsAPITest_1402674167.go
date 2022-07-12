// Create an API GRPC test returns "OK - Returns the created test details." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
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
				Host:     common.PtrString("localhost"),
				Port:     common.PtrInt64(50051),
				Service:  common.PtrString("Hello"),
				Method:   datadog.HTTPMETHOD_GET.Ptr(),
				Message:  common.PtrString(""),
				Metadata: map[string]string{},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_grpc_test_payload.json",
		Name:    "Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response",
		Options: datadog.SyntheticsTestOptions{
			MinFailureDuration: common.PtrInt64(0),
			MinLocationFailed:  common.PtrInt64(1),
			MonitorOptions: &datadog.SyntheticsTestOptionsMonitorOptions{
				RenotifyInterval: common.PtrInt64(0),
			},
			MonitorName: common.PtrString("Example-Create_an_API_GRPC_test_returns_OK_Returns_the_created_test_details_response"),
			TickEvery:   common.PtrInt64(60),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_GRPC.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadog.SYNTHETICSAPITESTTYPE_API,
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsAPITest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsAPITest`:\n%s\n", responseContent)
}
