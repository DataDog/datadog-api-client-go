// Update an Elastic Cloud integration account returns "OK" response

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
	body := datadogV2.ElasticCloudIntegrationAccountUpdateRequest{
		Data: datadogV2.ElasticCloudIntegrationAccountUpdateData{
			Attributes: datadogV2.ElasticCloudIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.ElasticCloudIntegrationAccountAuthenticationUpdate{
					ElasticCloudIntegrationAccountBasicAuthUpdate: &datadogV2.ElasticCloudIntegrationAccountBasicAuthUpdate{
						AuthType: datadogV2.ELASTICCLOUDINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
						Password: datadog.PtrString("your-password"),
						Username: datadog.PtrString("datadog"),
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
				Name: datadog.PtrString("elastic-cloud-prod"),
				Settings: &datadogV2.ElasticCloudIntegrationAccountSettingsUpdate{
					Tags: datadog.PtrString("env:prod,team:saasint"),
					Url:  datadog.PtrString("https://example.es.us-central1.gcp.cloud.es.io:9243"),
				},
			},
			Id:   "953a0060-81ec-4221-aed4-d4733b59cd96",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateElasticCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationApi(apiClient)
	resp, r, err := api.UpdateElasticCloudIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationApi.UpdateElasticCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudIntegrationApi.UpdateElasticCloudIntegrationAccount`:\n%s\n", responseContent)
}
