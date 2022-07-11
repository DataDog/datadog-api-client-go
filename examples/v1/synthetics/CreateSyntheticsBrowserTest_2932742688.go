// Create a browser test returns "OK - Returns saved rumSettings." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SyntheticsBrowserTest{
		Config: datadog.SyntheticsBrowserTestConfig{
			Assertions: []datadog.SyntheticsAssertion{},
			ConfigVariables: []datadog.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: datadog.SyntheticsTestRequest{
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Url:    datadog.PtrString("https://datadoghq.com"),
			},
			SetCookie: datadog.PtrString("name:test"),
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "Test message",
		Name:    "Example-Create_a_browser_test_returns_OK_Returns_saved_rumSettings_response",
		Options: datadog.SyntheticsTestOptions{
			AcceptSelfSigned: datadog.PtrBool(false),
			AllowInsecure:    datadog.PtrBool(true),
			DeviceIds: []datadog.SyntheticsDeviceID{
				datadog.SYNTHETICSDEVICEID_TABLET,
			},
			DisableCors:        datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			NoScreenshot:       datadog.PtrBool(true),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			RumSettings: &datadog.SyntheticsBrowserTestRumSettings{
				IsEnabled:     true,
				ApplicationId: datadog.PtrString("mockApplicationId"),
				ClientTokenId: datadog.PtrInt64(12345),
			},
			TickEvery: datadog.PtrInt64(300),
			Ci: &datadog.SyntheticsTestCiOptions{
				ExecutionRule: datadog.SYNTHETICSTESTEXECUTIONRULE_SKIPPED.Ptr(),
			},
		},
		Tags: []string{
			"testing:browser",
		},
		Type: datadog.SYNTHETICSBROWSERTESTTYPE_BROWSER,
		Steps: []datadog.SyntheticsStep{
			{
				AllowFailure: datadog.PtrBool(false),
				IsCritical:   datadog.PtrBool(true),
				Name:         datadog.PtrString("Refresh page"),
				Params:       new(interface{}),
				Type:         datadog.SYNTHETICSSTEPTYPE_REFRESH.Ptr(),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.CreateSyntheticsBrowserTest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsBrowserTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsBrowserTest`:\n%s\n", responseContent)
}
