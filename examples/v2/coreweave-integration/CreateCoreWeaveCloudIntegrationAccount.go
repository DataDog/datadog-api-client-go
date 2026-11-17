// Create a CoreWeave integration account returns "Created" response

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
	body := datadogV2.CoreWeaveCloudIntegrationAccountCreateRequest{
		Data: datadogV2.CoreWeaveCloudIntegrationAccountCreateData{
			Attributes: datadogV2.CoreWeaveCloudIntegrationAccountCreateAttributes{
				Authentication: datadogV2.CoreWeaveCloudIntegrationAccountAuthenticationRequest{
					CoreWeaveCloudIntegrationAccountAPITokenAuthRequest: &datadogV2.CoreWeaveCloudIntegrationAccountAPITokenAuthRequest{
						ApiToken: "CW-SECRET-xxxxxxxxxxxxxxxxxxxx",
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
				Name: "coreweave-prod",
				Settings: datadogV2.CoreWeaveCloudIntegrationAccountSettingsRequest{
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
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCoreWeaveCloudIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCoreWeaveIntegrationApi(apiClient)
	resp, r, err := api.CreateCoreWeaveCloudIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreWeaveIntegrationApi.CreateCoreWeaveCloudIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CoreWeaveIntegrationApi.CreateCoreWeaveCloudIntegrationAccount`:\n%s\n", responseContent)
}
