// Create an API test with WEBSOCKET subtype returns "OK - Returns the created test details." response

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
						Target:   "message",
						Type:     datadogV1.SYNTHETICSASSERTIONTYPE_RECEIVED_MESSAGE,
					}},
				datadogV1.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
						Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
						Target:   2000,
						Type:     datadogV1.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
					}},
			},
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{},
			Request: &datadogV1.SyntheticsTestRequest{
				Url:     datadog.PtrString("ws://datadoghq.com"),
				Message: datadog.PtrString("message"),
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_websocket_payload.json",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Synthetic"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_WEBSOCKET.Ptr(),
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
