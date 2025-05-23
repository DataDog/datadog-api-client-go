// Create an API test returns "OK - Returns the created test details." response

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
						Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
						Target: datadogV1.SyntheticsAssertionTargetValue{
							SyntheticsAssertionTargetValueNumber: datadog.PtrFloat64(1000)},
						Type: datadogV1.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
					}},
			},
			Request: &datadogV1.SyntheticsTestRequest{
				Method: datadog.PtrString("GET"),
				Url:    datadog.PtrString("https://example.com"),
			},
		},
		Locations: []string{
			"aws:eu-west-3",
		},
		Message: "Notification message",
		Name:    "Example test name",
		Options: datadogV1.SyntheticsTestOptions{
			Ci: &datadogV1.SyntheticsTestCiOptions{
				ExecutionRule: datadogV1.SYNTHETICSTESTEXECUTIONRULE_BLOCKING,
			},
			DeviceIds: []string{
				"chrome.laptop_large",
			},
			HttpVersion: datadogV1.SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP1.Ptr(),
			MonitorOptions: &datadogV1.SyntheticsTestOptionsMonitorOptions{
				NotificationPresetName: datadogV1.SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL.Ptr(),
			},
			RestrictedRoles: []string{
				"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			Retry: &datadogV1.SyntheticsTestOptionsRetry{},
			RumSettings: &datadogV1.SyntheticsBrowserTestRumSettings{
				ApplicationId: datadog.PtrString("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
				ClientTokenId: datadog.PtrInt64(12345),
				IsEnabled:     true,
			},
			Scheduling: &datadogV1.SyntheticsTestOptionsScheduling{
				Timeframes: []datadogV1.SyntheticsTestOptionsSchedulingTimeframe{
					{
						Day:  1,
						From: "07:00",
						To:   "16:00",
					},
					{
						Day:  3,
						From: "07:00",
						To:   "16:00",
					},
				},
				Timezone: "America/New_York",
			},
		},
		Status:  datadogV1.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
		Tags: []string{
			"env:production",
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
