// Create a browser test returns "OK - Returns the created test details." response

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
	body := datadogV1.SyntheticsBrowserTest{
		Config: datadogV1.SyntheticsBrowserTestConfig{
			Assertions: []datadogV1.SyntheticsAssertion{},
			Variables: []datadogV1.SyntheticsBrowserVariable{
				{
					Type:    datadogV1.SYNTHETICSBROWSERVARIABLETYPE_TEXT,
					Name:    "TEST_VARIABLE",
					Pattern: datadog.PtrString("secret"),
					Secure:  datadog.PtrBool(true),
					Example: datadog.PtrString("secret"),
				},
			},
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
					Secure:  datadog.PtrBool(true),
				},
			},
			Request: datadogV1.SyntheticsTestRequest{
				Method: datadog.PtrString("GET"),
				Url:    datadog.PtrString("https://datadoghq.com"),
			},
			SetCookie: datadog.PtrString("name:test"),
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "Test message",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			AcceptSelfSigned: datadog.PtrBool(false),
			AllowInsecure:    datadog.PtrBool(true),
			DeviceIds: []datadogV1.SyntheticsDeviceID{
				datadogV1.SYNTHETICSDEVICEID_CHROME_LAPTOP_LARGE,
			},
			DisableCors:        datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			NoScreenshot:       datadog.PtrBool(true),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(2),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery:             datadog.PtrInt64(300),
			EnableProfiling:       datadog.PtrBool(true),
			EnableSecurityTesting: datadog.PtrBool(true),
		},
		Tags: []string{
			"testing:browser",
		},
		Type: datadogV1.SYNTHETICSBROWSERTESTTYPE_BROWSER,
		Steps: []datadogV1.SyntheticsStep{
			{
				AllowFailure: datadog.PtrBool(false),
				IsCritical:   datadog.PtrBool(true),
				Name:         datadog.PtrString("Refresh page"),
				Params:       new(interface{}),
				Type:         datadogV1.SYNTHETICSSTEPTYPE_REFRESH.Ptr(),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.CreateSyntheticsBrowserTest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsBrowserTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsBrowserTest`:\n%s\n", responseContent)
}
