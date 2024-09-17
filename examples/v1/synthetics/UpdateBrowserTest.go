// Edit a browser test returns "OK" response

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
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{
				{
					Name:   "VARIABLE_NAME",
					Secure: datadog.PtrBool(false),
					Type:   datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: datadogV1.SyntheticsTestRequest{
				BasicAuth: &datadogV1.SyntheticsBasicAuth{
					SyntheticsBasicAuthWeb: &datadogV1.SyntheticsBasicAuthWeb{
						Password: "PaSSw0RD!",
						Type:     datadogV1.SYNTHETICSBASICAUTHWEBTYPE_WEB.Ptr(),
						Username: "my_username",
					}},
				BodyType: datadogV1.SYNTHETICSTESTREQUESTBODYTYPE_TEXT_PLAIN.Ptr(),
				CallType: datadogV1.SYNTHETICSTESTCALLTYPE_UNARY.Ptr(),
				Certificate: &datadogV1.SyntheticsTestRequestCertificate{
					Cert: &datadogV1.SyntheticsTestRequestCertificateItem{},
					Key:  &datadogV1.SyntheticsTestRequestCertificateItem{},
				},
				CertificateDomains: []string{},
				Files: []datadogV1.SyntheticsTestRequestBodyFile{
					{},
				},
				HttpVersion: datadogV1.SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP1.Ptr(),
				Proxy: &datadogV1.SyntheticsTestRequestProxy{
					Url: "https://example.com",
				},
				Service: datadog.PtrString("Greeter"),
				Url:     datadog.PtrString("https://example.com"),
			},
			Variables: []datadogV1.SyntheticsBrowserVariable{
				{
					Name: "VARIABLE_NAME",
					Type: datadogV1.SYNTHETICSBROWSERVARIABLETYPE_TEXT,
				},
			},
		},
		Locations: []string{
			"aws:eu-west-3",
		},
		Message: "",
		Name:    "Example test name",
		Options: datadogV1.SyntheticsTestOptions{
			Ci: &datadogV1.SyntheticsTestCiOptions{
				ExecutionRule: datadogV1.SYNTHETICSTESTEXECUTIONRULE_BLOCKING.Ptr(),
			},
			DeviceIds: []datadogV1.SyntheticsDeviceID{
				datadogV1.SYNTHETICSDEVICEID_CHROME_LAPTOP_LARGE,
			},
			HttpVersion:    datadogV1.SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP1.Ptr(),
			MonitorOptions: &datadogV1.SyntheticsTestOptionsMonitorOptions{},
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
						Day:  datadog.PtrInt32(1),
						From: datadog.PtrString("07:00"),
						To:   datadog.PtrString("16:00"),
					},
					{
						Day:  datadog.PtrInt32(3),
						From: datadog.PtrString("07:00"),
						To:   datadog.PtrString("16:00"),
					},
				},
				Timezone: datadog.PtrString("America/New_York"),
			},
		},
		Status: datadogV1.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Steps: []datadogV1.SyntheticsStep{
			{
				Type: datadogV1.SYNTHETICSSTEPTYPE_ASSERT_ELEMENT_CONTENT.Ptr(),
			},
		},
		Tags: []string{
			"env:prod",
		},
		Type: datadogV1.SYNTHETICSBROWSERTESTTYPE_BROWSER,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewSyntheticsApi(apiClient)
	resp, r, err := api.UpdateBrowserTest(ctx, "public_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateBrowserTest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateBrowserTest`:\n%s\n", responseContent)
}
