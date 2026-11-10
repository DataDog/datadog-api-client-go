// Update a Nebius Cloud integration account returns "OK" response

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
	body := datadogV2.NebiusCloudIntegrationAccountUpdateRequest{
		Data: datadogV2.NebiusCloudIntegrationAccountUpdateData{
			Attributes: datadogV2.NebiusCloudIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.NebiusCloudIntegrationAccountAuthenticationUpdate{
					NebiusCloudIntegrationAccountServiceAccountAuthUpdate: &datadogV2.NebiusCloudIntegrationAccountServiceAccountAuthUpdate{
						AuthType: datadogV2.NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT,
						PrivateKey: datadog.PtrString(`-----BEGIN PRIVATE KEY-----
MIIE...
-----END PRIVATE KEY-----`),
						PublicKeyId:      *datadog.NewNullableString(datadog.PtrString("publickey-yyyyyyyy")),
						ServiceAccountId: datadog.PtrString("serviceaccount-xxxxxxxx"),
						StaticApiKey:     datadog.PtrString("eyJhbGciOi...static-api-key"),
					}},
				Dataflows: &datadogV2.NebiusCloudIntegrationDataflowsRequest{
					NebiusCloudCostCollection: &datadogV2.NebiusCloudCostCollectionIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.NebiusCloudCostCollectionIntegrationDataflowSettingsRequest{
							AccessKeyId:      datadog.PtrString("AKID0123456789ABCDEF"),
							AccessKeySecret:  datadog.PtrString("wJalrXUtnFEMIstaticaccesskey"),
							BucketName:       datadog.PtrString("nebius-cost-exports"),
							BucketPrefixPath: datadog.PtrString("cost-export/2026/"),
							BucketRegion:     datadog.PtrString("eu-north1"),
						},
					},
					NebiusCloudLogs: &datadogV2.NebiusCloudLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.NebiusCloudLogsIntegrationDataflowSettingsRequest{
							EnabledServices: []string{
								"nebius-cloud-compute",
								"nebius-cloud-kubernetes",
								"default",
							},
						},
					},
					NebiusCloudResources: &datadogV2.NebiusCloudResourcesIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					NebiusCloudServiceQuotaCollection: &datadogV2.NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("nebius-prod"),
				Settings: &datadogV2.NebiusCloudIntegrationAccountSettingsUpdate{
					AccountTags: []string{
						"env:prod",
						"team:ai",
					},
					TenantId: datadog.PtrString("tenant-12345678"),
				},
			},
			Id:   "e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7081",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateNebiusCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewNebiusCloudIntegrationApi(apiClient)
	resp, r, err := api.UpdateNebiusCloudIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NebiusCloudIntegrationApi.UpdateNebiusCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NebiusCloudIntegrationApi.UpdateNebiusCloudIntegrationAccount`:\n%s\n", responseContent)
}
