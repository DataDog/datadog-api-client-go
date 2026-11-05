// Create an HTTP Action Connection with an AWS Secrets Manager token

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "action_connection" in the system
	ActionConnectionDataID := os.Getenv("ACTION_CONNECTION_DATA_ID")

	body := datadogV2.CreateActionConnectionRequest{
		Data: datadogV2.ActionConnectionData{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributes{
				Name: "HTTP Secrets Manager Connection exampleactionconnection",
				ExternalSecretsManager: *datadogV2.NewNullableActionConnectionAWSSecretsManager(&datadogV2.ActionConnectionAWSSecretsManager{
					Type:         datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
					ConnectionId: ActionConnectionDataID,
					Region:       "us-east-1",
				}),
				Integration: datadogV2.ActionConnectionIntegration{
					HTTPIntegration: &datadogV2.HTTPIntegration{
						Type:    datadogV2.HTTPINTEGRATIONTYPE_HTTP,
						BaseUrl: "https://example.com",
						Credentials: datadogV2.HTTPCredentials{
							HTTPTokenAuth: &datadogV2.HTTPTokenAuth{
								Type: datadogV2.HTTPTOKENAUTHTYPE_HTTPTOKENAUTH,
								Tokens: []datadogV2.HTTPToken{
									{
										Name: "token",
										Type: datadogV2.TOKENTYPE_SECRET,
										Value: datadogV2.HTTPTokenValue{
											ActionConnectionAWSSecretsManagerReference: &datadogV2.ActionConnectionAWSSecretsManagerReference{
												Type:     datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
												SecretId: "example/http",
												Key:      datadog.PtrString("token"),
											}},
									},
								},
								Headers: []datadogV2.HTTPHeader{
									{
										Name:  "ApiKey",
										Value: "literal-value",
									},
								},
								UrlParameters: []datadogV2.UrlParam{
									{
										Name:  "apiKey",
										Value: "literal-value",
									},
								},
								Body: &datadogV2.HTTPBody{
									ContentType: datadog.PtrString("application/json"),
									Content:     datadog.PtrString(`{"example":"literal"}`),
								},
							}},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewActionConnectionApi(apiClient)
	resp, r, err := api.CreateActionConnection(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.CreateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.CreateActionConnection`:\n%s\n", responseContent)
}
