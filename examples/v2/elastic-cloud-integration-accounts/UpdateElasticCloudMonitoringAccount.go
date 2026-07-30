// Update an Elastic Cloud monitoring account returns "OK" response

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
	body := datadogV2.ElasticCloudMonitoringAccountUpdateRequest{
		Data: datadogV2.ElasticCloudMonitoringAccountUpdateData{
			Attributes: datadogV2.ElasticCloudMonitoringAccountUpdateAttributes{
				Authentication: &datadogV2.ElasticCloudAuthentication{
					ElasticCloudBasicAuth: &datadogV2.ElasticCloudBasicAuth{
						Password: "your-password",
						Type:     datadogV2.ELASTICCLOUDBASICAUTHTYPE_BASIC,
						Username: "datadog",
					}},
				Dataflows: []datadogV2.ElasticCloudDataflow{
					{
						Enabled: datadog.PtrBool(true),
						Id:      datadogV2.ELASTICCLOUDDATAFLOWID_METRICS,
					},
				},
				Name: datadog.PtrString("elastic-cloud-prod"),
				Settings: &datadogV2.ElasticCloudSettingsUpdate{
					CatAllocationStatsEnabled: datadog.PtrBool(false),
					DetailedIndexStatsEnabled: datadog.PtrBool(false),
					IndexStatsEnabled:         datadog.PtrBool(false),
					PendingTaskStatsEnabled:   datadog.PtrBool(false),
					PshardGracefulToEnabled:   datadog.PtrBool(false),
					PshardStatsEnabled:        datadog.PtrBool(false),
					SlmStatsEnabled:           datadog.PtrBool(false),
					Tags: []string{
						"env:prod",
					},
					Url: datadog.PtrString("https://example.es.us-central1.gcp.cloud.es.io:9243"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateElasticCloudMonitoringAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateElasticCloudMonitoringAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationAccountsApi.UpdateElasticCloudMonitoringAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudIntegrationAccountsApi.UpdateElasticCloudMonitoringAccount`:\n%s\n", responseContent)
}
