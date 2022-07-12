// Create a browser test returns "OK - Returns saved rumSettings." response

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
	body := datadog.SyntheticsBrowserTest{
		Config: datadog.SyntheticsBrowserTestConfig{
			Assertions: []datadog.SyntheticsAssertion{},
			ConfigVariables: []datadog.SyntheticsConfigVariable{
				{
					Example: common.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: common.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: datadog.SyntheticsTestRequest{
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Url:    common.PtrString("https://datadoghq.com"),
			},
			SetCookie: common.PtrString("name:test"),
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "Test message",
		Name:    "Example-Create_a_browser_test_returns_OK_Returns_saved_rumSettings_response",
		Options: datadog.SyntheticsTestOptions{
			AcceptSelfSigned: common.PtrBool(false),
			AllowInsecure:    common.PtrBool(true),
			DeviceIds: []datadog.SyntheticsDeviceID{
				datadog.SYNTHETICSDEVICEID_TABLET,
			},
			DisableCors:        common.PtrBool(true),
			FollowRedirects:    common.PtrBool(true),
			MinFailureDuration: common.PtrInt64(10),
			MinLocationFailed:  common.PtrInt64(1),
			NoScreenshot:       common.PtrBool(true),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    common.PtrInt64(3),
				Interval: common.PtrFloat64(10),
			},
			RumSettings: &datadog.SyntheticsBrowserTestRumSettings{
				IsEnabled:     true,
				ApplicationId: common.PtrString("mockApplicationId"),
				ClientTokenId: common.PtrInt64(123456),
			},
			TickEvery: common.PtrInt64(300),
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
				AllowFailure: common.PtrBool(false),
				IsCritical:   common.PtrBool(true),
				Name:         common.PtrString("Refresh page"),
				Params:       new(interface{}),
				Type:         datadog.SYNTHETICSSTEPTYPE_REFRESH.Ptr(),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsBrowserTest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsBrowserTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsBrowserTest`:\n%s\n", responseContent)
}
