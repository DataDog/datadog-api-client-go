// Create an Elastic Cloud integration account returns "Created" response

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
	body := datadogV2.ElasticCloudIntegrationAccountCreateRequest{
		Data: datadogV2.ElasticCloudIntegrationAccountCreateData{
			Attributes: datadogV2.ElasticCloudIntegrationAccountCreateAttributes{
				Authentication: datadogV2.ElasticCloudIntegrationAccountAuthenticationRequest{
					IntegrationAccountBasicAuthRequest: &datadogV2.IntegrationAccountBasicAuthRequest{
						AuthType: datadogV2.INTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
						Password: "your-password",
						Username: "datadog",
					}},
				Dataflows: &datadogV2.ElasticCloudIntegrationDataflowsRequest{
					ElasticCloudDetailedIndexStats: &datadogV2.ElasticCloudDetailedIndexStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudIndexStats: &datadogV2.ElasticCloudIndexStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudPendingTaskStats: &datadogV2.ElasticCloudPendingTaskStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudPrimaryShardGracefulTimeout: &datadogV2.ElasticCloudPrimaryShardGracefulTimeoutIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudPrimaryShardStats: &datadogV2.ElasticCloudPrimaryShardStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudShardAllocationStats: &datadogV2.ElasticCloudShardAllocationStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					ElasticCloudSlmStats: &datadogV2.ElasticCloudSlmStatsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "elastic-cloud-prod",
				Settings: datadogV2.ElasticCloudIntegrationAccountSettingsRequest{
					Tags: datadog.PtrString("env:prod,team:saasint"),
					Url:  "https://example.es.us-central1.gcp.cloud.es.io:9243",
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateElasticCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationAccountsApi(apiClient)
	resp, r, err := api.CreateElasticCloudIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationAccountsApi.CreateElasticCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudIntegrationAccountsApi.CreateElasticCloudIntegrationAccount`:\n%s\n", responseContent)
}
