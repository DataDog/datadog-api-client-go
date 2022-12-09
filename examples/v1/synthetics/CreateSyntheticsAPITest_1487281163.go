// Create an API HTTP test returns "OK - Returns the created test details." response

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
						Property: datadog.PtrString("{{ PROPERTY }}"),
						Target:   "text/html",
						Type:     datadogV1.SYNTHETICSASSERTIONTYPE_HEADER,
					}},
				datadogV1.SyntheticsAssertion{
					SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
						Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
						Target:   2000,
						Type:     datadogV1.SYNTHETICSASSERTIONTYPE_RESPONSE_TIME,
					}},
				datadogV1.SyntheticsAssertion{
					SyntheticsAssertionJSONPathTarget: &datadogV1.SyntheticsAssertionJSONPathTarget{
						Operator: datadogV1.SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH,
						Target: &datadogV1.SyntheticsAssertionJSONPathTargetTarget{
							JsonPath:    datadog.PtrString("topKey"),
							Operator:    datadog.PtrString("isNot"),
							TargetValue: "0",
						},
						Type: datadogV1.SYNTHETICSASSERTIONTYPE_BODY,
					}},
				datadogV1.SyntheticsAssertion{
					SyntheticsAssertionXPathTarget: &datadogV1.SyntheticsAssertionXPathTarget{
						Operator: datadogV1.SYNTHETICSASSERTIONXPATHOPERATOR_VALIDATES_X_PATH,
						Target: &datadogV1.SyntheticsAssertionXPathTargetTarget{
							XPath:       datadog.PtrString("target-xpath"),
							TargetValue: "0",
							Operator:    datadog.PtrString("contains"),
						},
						Type: datadogV1.SYNTHETICSASSERTIONTYPE_BODY,
					}},
			},
			ConfigVariables: []datadogV1.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadogV1.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Request: &datadogV1.SyntheticsTestRequest{
				Certificate: &datadogV1.SyntheticsTestRequestCertificate{
					Cert: &datadogV1.SyntheticsTestRequestCertificateItem{
						Content:   datadog.PtrString("cert-content"),
						Filename:  datadog.PtrString("cert-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
					Key: &datadogV1.SyntheticsTestRequestCertificateItem{
						Content:   datadog.PtrString("key-content"),
						Filename:  datadog.PtrString("key-filename"),
						UpdatedAt: datadog.PtrString("2020-10-16T09:23:24.857Z"),
					},
				},
				Headers: map[string]string{
					"unique": "examplecreateanapihttptestreturnsokreturnsthecreatedtestdetailsresponse",
				},
				Method:  datadog.PtrString("GET"),
				Timeout: datadog.PtrFloat64(10),
				Url:     datadog.PtrString("https://datadoghq.com"),
				Proxy: &datadogV1.SyntheticsTestRequestProxy{
					Url:     "https://datadoghq.com",
					Headers: map[string]string{},
				},
				BasicAuth: &datadogV1.SyntheticsBasicAuth{
					SyntheticsBasicAuthOauthClient: &datadogV1.SyntheticsBasicAuthOauthClient{
						AccessTokenUrl:         "https://datadog-token.com",
						Audience:               datadog.PtrString("audience"),
						ClientId:               "client-id",
						ClientSecret:           "client-secret",
						Resource:               datadog.PtrString("resource"),
						Scope:                  datadog.PtrString("yoyo"),
						TokenApiAuthentication: datadogV1.SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_HEADER,
						Type:                   datadogV1.SYNTHETICSBASICAUTHOAUTHCLIENTTYPE_OAUTH_CLIENT.Ptr(),
					}},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_http_test_payload.json",
		Name:    "Example-Create_an_API_HTTP_test_returns_OK_Returns_the_created_test_details_response",
		Options: datadogV1.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Create_an_API_HTTP_test_returns_OK_Returns_the_created_test_details_response"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadogV1.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(10),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_HTTP.Ptr(),
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
