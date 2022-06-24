// Edit a browser test returns "OK" response

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
					Name: "VARIABLE_NAME",
					Type: datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: datadog.SyntheticsTestRequest{
				BasicAuth: &datadog.SyntheticsBasicAuth{
					SyntheticsBasicAuthWeb: &datadog.SyntheticsBasicAuthWeb{
						Password: "PaSSw0RD!",
						Type:     datadog.SYNTHETICSBASICAUTHWEBTYPE_WEB.Ptr(),
						Username: "my_username",
					}},
				Certificate: &datadog.SyntheticsTestRequestCertificate{
					Cert: &datadog.SyntheticsTestRequestCertificateItem{},
					Key:  &datadog.SyntheticsTestRequestCertificateItem{},
				},
				Method: datadog.HTTPMETHOD_GET.Ptr(),
				Proxy: &datadog.SyntheticsTestRequestProxy{
					Url: "https://example.com",
				},
				Url: datadog.PtrString("https://example.com"),
			},
			Variables: []datadog.SyntheticsBrowserVariable{
				{
					Name: "VARIABLE_NAME",
					Type: datadog.SYNTHETICSBROWSERVARIABLETYPE_TEXT,
				},
			},
		},
		Locations: []string{
			"aws:eu-west-3",
		},
		Message: "",
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
		Status: datadog.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Steps: []datadog.SyntheticsStep{
			{
				Type: datadog.SYNTHETICSSTEPTYPE_ASSERT_ELEMENT_CONTENT.Ptr(),
			},
		},
		Tags: []string{
			"env:prod",
		},
		Type: datadog.SYNTHETICSBROWSERTESTTYPE_BROWSER,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.UpdateBrowserTest(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateBrowserTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateBrowserTest`:\n%s\n", responseContent)
}
