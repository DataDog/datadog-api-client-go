// Create a multi-step api test with every type of basicAuth returns "OK - Returns the created test details." response

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
			Steps: []datadogV1.SyntheticsAPIStep{
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthWeb: &datadogV1.SyntheticsBasicAuthWeb{
									Password: "password",
									Username: "username",
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthWeb: &datadogV1.SyntheticsBasicAuthWeb{
									Password: "password",
									Username: "username",
									Type:     datadogV1.SYNTHETICSBASICAUTHWEBTYPE_WEB.Ptr(),
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthSigv4: &datadogV1.SyntheticsBasicAuthSigv4{
									AccessKey: "accessKey",
									SecretKey: "secretKey",
									Type:      datadogV1.SYNTHETICSBASICAUTHSIGV4TYPE_SIGV4,
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthNTLM: &datadogV1.SyntheticsBasicAuthNTLM{
									Type: datadogV1.SYNTHETICSBASICAUTHNTLMTYPE_NTLM,
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthDigest: &datadogV1.SyntheticsBasicAuthDigest{
									Password: "password",
									Username: "username",
									Type:     datadogV1.SYNTHETICSBASICAUTHDIGESTTYPE_DIGEST,
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthOauthClient: &datadogV1.SyntheticsBasicAuthOauthClient{
									AccessTokenUrl:         "accessTokenUrl",
									TokenApiAuthentication: datadogV1.SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_HEADER,
									ClientId:               "clientId",
									ClientSecret:           "clientSecret",
									Type:                   datadogV1.SYNTHETICSBASICAUTHOAUTHCLIENTTYPE_OAUTH_CLIENT,
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
				datadogV1.SyntheticsAPIStep{
					SyntheticsAPITestStep: &datadogV1.SyntheticsAPITestStep{
						Assertions: []datadogV1.SyntheticsAssertion{
							datadogV1.SyntheticsAssertion{
								SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{
									Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_IS,
									Type:     datadogV1.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
									Target:   200,
								}},
						},
						Name: "request is sent",
						Request: datadogV1.SyntheticsTestRequest{
							Url:    datadog.PtrString("https://httpbin.org/status/200"),
							Method: datadog.PtrString("GET"),
							BasicAuth: &datadogV1.SyntheticsBasicAuth{
								SyntheticsBasicAuthOauthROP: &datadogV1.SyntheticsBasicAuthOauthROP{
									AccessTokenUrl:         "accessTokenUrl",
									Password:               "password",
									TokenApiAuthentication: datadogV1.SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_HEADER,
									Username:               "username",
									Type:                   datadogV1.SYNTHETICSBASICAUTHOAUTHROPTYPE_OAUTH_ROP,
								}},
						},
						Subtype: datadogV1.SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
					}},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: "BDD test payload: synthetics_api_test_multi_step_with_every_type_of_basic_auth.json",
		Name:    "Example-Synthetic",
		Options: datadogV1.SyntheticsTestOptions{
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadogV1.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Type:    datadogV1.SYNTHETICSAPITESTTYPE_API,
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
