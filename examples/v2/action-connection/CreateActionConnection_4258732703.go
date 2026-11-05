// Create an Anthropic Action Connection with an AWS Secrets Manager API token

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
				Name: "Anthropic Secrets Manager Connection exampleactionconnection",
				ExternalSecretsManager: *datadogV2.NewNullableActionConnectionAWSSecretsManager(&datadogV2.ActionConnectionAWSSecretsManager{
					Type:         datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
					ConnectionId: ActionConnectionDataID,
					Region:       "us-east-1",
				}),
				Integration: datadogV2.ActionConnectionIntegration{
					AnthropicIntegration: &datadogV2.AnthropicIntegration{
						Type: datadogV2.ANTHROPICINTEGRATIONTYPE_ANTHROPIC,
						Credentials: datadogV2.AnthropicCredentials{
							AnthropicAPIKey: &datadogV2.AnthropicAPIKey{
								Type: datadogV2.ANTHROPICAPIKEYTYPE_ANTHROPICAPIKEY,
								ApiToken: datadogV2.AnthropicAPIKeyApiToken{
									ActionConnectionAWSSecretsManagerReference: &datadogV2.ActionConnectionAWSSecretsManagerReference{
										Type:     datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
										SecretId: "example/anthropic",
										Key:      datadog.PtrString("api_token"),
									}},
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
