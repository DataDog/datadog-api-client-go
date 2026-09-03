// Update a Databricks integration account returns "OK" response

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
	body := datadogV2.DatabricksIntegrationAccountUpdateRequest{
		Data: datadogV2.DatabricksIntegrationAccountUpdateData{
			Attributes: datadogV2.DatabricksIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.DatabricksIntegrationAccountAuthenticationUpdate{
					DatabricksIntegrationAccountOAuthAuthUpdate: &datadogV2.DatabricksIntegrationAccountOAuthAuthUpdate{
						AuthType:      datadogV2.DATABRICKSINTEGRATIONACCOUNTOAUTHAUTHTYPE_DATABRICKS_OAUTH,
						AzureTenantId: datadog.PtrString("4d3bac44-0230-4732-9e70-cc00736f0a97"),
						ClientId:      "5c10654a-b3a3-4840-b37f-f477590c70a0",
						ClientSecret:  datadog.PtrString("your-client-secret"),
					}},
				Dataflows: &datadogV2.DatabricksIntegrationDataflowsRequest{
					DatabricksCloudCostMetrics: &datadogV2.DatabricksCloudCostMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest{
							CcmCollectAllWorkspaces: datadog.PtrBool(true),
						},
					},
					DatabricksDataJobMonitoring: &datadogV2.DatabricksDataJobMonitoringIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest{
							DdApiKeyId:                 datadog.PtrString("fe383f4e-09fc-46bf-8e10-4efdd453a646"),
							DdApiKeySecret:             datadog.PtrString("your-datadog-api-key"),
							DjmGlobalInitScriptEnabled: datadog.PtrBool(true),
							ScriptGpumEnabled:          datadog.PtrBool(true),
							ScriptLogsEnabled:          datadog.PtrBool(true),
						},
					},
					DatabricksDataObservability: &datadogV2.DatabricksDataObservabilityIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
						Settings: &datadogV2.DatabricksDataObservabilityIntegrationDataflowSettingsRequest{
							DoCrawlersCron:    datadog.PtrString("0 * * * *"),
							SyncSystemCatalog: datadog.PtrBool(true),
						},
					},
					DatabricksModelServingMetrics: &datadogV2.DatabricksModelServingMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					DatabricksServerlessJobs: &datadogV2.DatabricksServerlessJobsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("My Databricks Workspace"),
				Settings: &datadogV2.DatabricksIntegrationAccountSettingsUpdate{
					SystemTablesSqlWarehouseId: datadog.PtrString("aba7c023d4172910"),
					WorkspaceUrl:               datadog.PtrString("https://dbc-1234abcd.cloud.databricks.com"),
				},
			},
			Id:   "a9a69c2e-4f8d-4e42-9c1a-2a7a2d3b7c6f",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDatabricksIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDatabricksIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateDatabricksIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabricksIntegrationAccountsApi.UpdateDatabricksIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DatabricksIntegrationAccountsApi.UpdateDatabricksIntegrationAccount`:\n%s\n", responseContent)
}
