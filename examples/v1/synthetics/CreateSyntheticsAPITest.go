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
		Config: datadog.SyntheticsAPITestConfig{
			Assertions: []datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
						Target:   1000,
						Type:     datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
					}},
			},
			Request: &datadog.SyntheticsTestRequest{
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Url:    datadog.PtrString("https://example.com"),
			},
		},
		Locations: []string{
			"aws:eu-west-3",
		},
		Message: "Notification message",
		Name:    "Example test name",
		Options: datadog.SyntheticsTestOptions{
			Ci: &datadog.SyntheticsTestCiOptions{
				ExecutionRule: datadog.SYNTHETICSTESTEXECUTIONRULE_BLOCKING.Ptr(),
			},
			DeviceIds: []datadog.SyntheticsDeviceID{
				datadog.SYNTHETICSDEVICEID_LAPTOP_LARGE,
			},
			MonitorOptions: &datadog.SyntheticsTestOptionsMonitorOptions{},
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			Retry: &datadog.SyntheticsTestOptionsRetry{},
			RumSettings: &datadog.SyntheticsBrowserTestRumSettings{
				ApplicationId: datadog.PtrString("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
				ClientTokenId: datadog.PtrInt64(12345),
				IsEnabled:     true,
			},
		},
		Status:  datadog.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
		Tags: []string{
			"env:production",
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
