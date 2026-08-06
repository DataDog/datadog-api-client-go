// Create an Elastic Cloud monitoring account returns "Created" response

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
	body := datadogV2.ElasticCloudMonitoringAccountRequest{
		Data: datadogV2.ElasticCloudMonitoringAccountCreateData{
			Attributes: datadogV2.ElasticCloudMonitoringAccountAttributes{
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
				Name: "elastic-cloud-prod",
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
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateElasticCloudMonitoringAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewElasticCloudMonitoringApi(apiClient)
	resp, r, err := api.CreateElasticCloudMonitoringAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElasticCloudMonitoringApi.CreateElasticCloudMonitoringAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ElasticCloudMonitoringApi.CreateElasticCloudMonitoringAccount`:\n%s\n", responseContent)
}
