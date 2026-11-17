// Update a CoreWeave integration account returns "OK" response

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
	body := datadogV2.CoreWeaveCloudIntegrationAccountUpdateRequest{
		Data: datadogV2.CoreWeaveCloudIntegrationAccountUpdateData{
			Attributes: datadogV2.CoreWeaveCloudIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.CoreWeaveCloudIntegrationAccountAuthenticationUpdate{
					CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate: &datadogV2.CoreWeaveCloudIntegrationAccountAPITokenAuthUpdate{
						ApiToken: datadog.PtrString("CW-SECRET-xxxxxxxxxxxxxxxxxxxx"),
						AuthType: datadogV2.COREWEAVECLOUDINTEGRATIONACCOUNTAPITOKENAUTHTYPE_COREWEAVE_API_TOKEN,
					}},
				Dataflows: &datadogV2.CoreWeaveCloudIntegrationDataflowsRequest{
					CoreweaveCloudLogs: &datadogV2.CoreWeaveCloudLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					CoreweaveCloudMetrics: &datadogV2.CoreWeaveCloudMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("coreweave-prod"),
				Settings: &datadogV2.CoreWeaveCloudIntegrationAccountSettingsUpdate{
					AccountTags: []string{
						"env:prod",
						"team:gpu",
					},
					ClusterFilters: []string{
						"cluster:prod-*",
						"!cluster:dev",
					},
				},
			},
			Id:   "f7a8b9c0-d1e2-4f3a-9b4c-5d6e7f8091a2",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCoreWeaveCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCoreWeaveIntegrationApi(apiClient)
	resp, r, err := api.UpdateCoreWeaveCloudIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreWeaveIntegrationApi.UpdateCoreWeaveCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CoreWeaveIntegrationApi.UpdateCoreWeaveCloudIntegrationAccount`:\n%s\n", responseContent)
}
