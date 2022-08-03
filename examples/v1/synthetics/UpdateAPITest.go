// Edit an API test returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	// there is a valid "synthetics_api_test" in the system
	SyntheticsAPITestPublicID := os.Getenv("SYNTHETICS_API_TEST_PUBLIC_ID")

	body := datadog.SyntheticsAPITest{
		Config: datadog.SyntheticsAPITestConfig{
			Assertions: []datadog.SyntheticsAssertion{
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SYNTHETICSASSERTIONOPERATOR_IS,
						Property: common.PtrString("{{ PROPERTY }}"),
						Target:   "text/html",
						Type:     datadog.SYNTHETICSASSERTIONTYPE_HEADER,
					}},
				datadog.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
						Operator: datadog.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
						Target:   2000,
						Type:     datadog.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
					}},
				datadog.SyntheticsAssertion{
					SyntheticsAssertionJSONPathTarget: &datadog.SyntheticsAssertionJSONPathTarget{
						Operator: datadog.SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH,
						Target: &datadog.SyntheticsAssertionJSONPathTargetTarget{
							JsonPath:    common.PtrString("topKey"),
							Operator:    common.PtrString("isNot"),
							TargetValue: "0",
						},
						Type: datadog.SYNTHETICSASSERTIONTYPE_BODY,
					}},
			},
			ConfigVariables: []datadog.SyntheticsConfigVariable{
				{
					Example: common.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: common.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: &datadog.SyntheticsTestRequest{
				Certificate: &datadog.SyntheticsTestRequestCertificate{
					Cert: &datadog.SyntheticsTestRequestCertificateItem{
						Filename:  common.PtrString("cert-filename"),
						UpdatedAt: common.PtrString("2020-10-16T09:23:24.857Z"),
					},
					Key: &datadog.SyntheticsTestRequestCertificateItem{
						Filename:  common.PtrString("key-filename"),
						UpdatedAt: common.PtrString("2020-10-16T09:23:24.857Z"),
					},
				},
				Headers: map[string]string{
					"unique": "exampleeditanapitestreturnsokresponse",
				},
				Method:  datadog.HTTPMETHOD_GET.Ptr(),
				Timeout: common.PtrFloat64(10),
				Url:     common.PtrString("https://datadoghq.com"),
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_payload.json",
		Name:    "Example-Edit_an_API_test_returns_OK_response-updated",
		Options: datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   common.PtrBool(false),
			AllowInsecure:      common.PtrBool(true),
			FollowRedirects:    common.PtrBool(true),
			MinFailureDuration: common.PtrInt64(10),
			MinLocationFailed:  common.PtrInt64(1),
			MonitorName:        common.PtrString("Test-TestSyntheticsAPITestLifecycle-1623076664"),
			MonitorPriority:    common.PtrInt32(5),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    common.PtrInt64(3),
				Interval: common.PtrFloat64(10),
			},
			TickEvery: common.PtrInt64(60),
		},
		Status:  datadog.SYNTHETICSTESTPAUSESTATUS_LIVE.Ptr(),
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadog.SYNTHETICSAPITESTTYPE_API,
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSyntheticsApi(apiClient)
	resp, r, err := api.UpdateAPITest(ctx, SyntheticsAPITestPublicID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateAPITest`:\n%s\n", responseContent)
}
