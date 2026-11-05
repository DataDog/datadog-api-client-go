// Create a GCP Action Connection with AWS Secrets Manager references

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
				Name: "GCP Secrets Manager Connection exampleactionconnection",
				ExternalSecretsManager: *datadogV2.NewNullableActionConnectionAWSSecretsManager(&datadogV2.ActionConnectionAWSSecretsManager{
					Type:         datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
					ConnectionId: ActionConnectionDataID,
					Region:       "us-east-1",
				}),
				Integration: datadogV2.ActionConnectionIntegration{
					GCPIntegration: &datadogV2.GCPIntegration{
						Type: datadogV2.GCPINTEGRATIONTYPE_GCP,
						Credentials: datadogV2.GCPCredentials{
							GCPServiceAccount: &datadogV2.GCPServiceAccount{
								Type: datadogV2.GCPSERVICEACCOUNTCREDENTIALTYPE_GCPSERVICEACCOUNT,
								PrivateKey: datadogV2.GCPServiceAccountPrivateKey{
									ActionConnectionAWSSecretsManagerReference: &datadogV2.ActionConnectionAWSSecretsManagerReference{
										Type:     datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
										SecretId: "example/gcp/service-account",
										Key:      datadog.PtrString("private_key"),
									}},
								ServiceAccountEmail: datadogV2.GCPServiceAccountServiceAccountEmail{
									ActionConnectionAWSSecretsManagerReference: &datadogV2.ActionConnectionAWSSecretsManagerReference{
										Type:     datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
										SecretId: "example/gcp/service-account",
										Key:      datadog.PtrString("client_email"),
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
