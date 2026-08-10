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
	body := datadogV2.ElasticCloudIntegrationAccountRequest{
		Data: datadogV2.ElasticCloudIntegrationAccountCreateData{
			Attributes: datadogV2.ElasticCloudIntegrationAccountAttributes{
				Interface: datadogV2.ElasticCloudInterface{
					ElasticCloudMonitoringInterface: &datadogV2.ElasticCloudMonitoringInterface{
						Authentication: datadogV2.ElasticCloudAuthentication{
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
						Settings: &datadogV2.ElasticCloudSettings{
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
							Url: "https://example.es.us-central1.gcp.cloud.es.io:9243",
						},
						Type: datadogV2.ELASTICCLOUDMONITORINGINTERFACETYPE_ELASTIC_CLOUD,
					}},
				Name: "elastic-cloud-prod",
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateElasticCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudIntegrationAccountsApi(apiClient)
	resp, r, err := api.CreateElasticCloudIntegrationAccount(ctx, datadogV2.ELASTICCLOUDINTERFACEID_ELASTIC_CLOUD, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudIntegrationAccountsApi.CreateElasticCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudIntegrationAccountsApi.CreateElasticCloudIntegrationAccount`:\n%s\n", responseContent)
}
