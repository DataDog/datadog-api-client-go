// Update a GCP Action Connection with a versioned AWS Secrets Manager reference

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
	// there is a valid "gcp_action_connection" in the system
	GcpActionConnectionDataID := os.Getenv("GCP_ACTION_CONNECTION_DATA_ID")

	// there is a valid "action_connection" in the system
	ActionConnectionDataID := os.Getenv("ACTION_CONNECTION_DATA_ID")

	body := datadogV2.UpdateActionConnectionRequest{
		Data: datadogV2.ActionConnectionDataUpdate{
			Type: datadogV2.ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
			Attributes: datadogV2.ActionConnectionAttributesUpdate{
				ExternalSecretsManager: *datadogV2.NewNullableActionConnectionAWSSecretsManager(&datadogV2.ActionConnectionAWSSecretsManager{
					Type:         datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
					ConnectionId: ActionConnectionDataID,
					Region:       "us-east-1",
				}),
				Integration: &datadogV2.ActionConnectionIntegrationUpdate{
					GCPIntegrationUpdate: &datadogV2.GCPIntegrationUpdate{
						Type: datadogV2.GCPINTEGRATIONTYPE_GCP,
						Credentials: &datadogV2.GCPCredentialsUpdate{
							GCPServiceAccountUpdate: &datadogV2.GCPServiceAccountUpdate{
								Type: datadogV2.GCPSERVICEACCOUNTCREDENTIALTYPE_GCPSERVICEACCOUNT,
								PrivateKey: &datadogV2.GCPServiceAccountUpdatePrivateKey{
									ActionConnectionAWSSecretsManagerReference: &datadogV2.ActionConnectionAWSSecretsManagerReference{
										Type:         datadogV2.ACTIONCONNECTIONAWSSECRETSMANAGERTYPE_AWS_SECRETS_MANAGER,
										SecretId:     "example/gcp/service-account",
										Key:          datadog.PtrString("private_key"),
										VersionStage: datadog.PtrString("AWSCURRENT"),
									}},
								ServiceAccountEmail: &datadogV2.GCPServiceAccountUpdateServiceAccountEmail{
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
	resp, r, err := api.UpdateActionConnection(ctx, GcpActionConnectionDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionConnectionApi.UpdateActionConnection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ActionConnectionApi.UpdateActionConnection`:\n%s\n", responseContent)
}
