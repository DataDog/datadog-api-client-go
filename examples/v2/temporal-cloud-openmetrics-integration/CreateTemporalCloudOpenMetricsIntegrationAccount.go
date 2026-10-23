// Create a Temporal Cloud OpenMetrics account returns "Created" response

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
	body := datadogV2.TemporalCloudOpenMetricsIntegrationAccountCreateRequest{
		Data: datadogV2.TemporalCloudOpenMetricsIntegrationAccountCreateData{
			Attributes: datadogV2.TemporalCloudOpenMetricsIntegrationAccountCreateAttributes{
				Authentication: datadogV2.TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest{
					TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthRequest: &datadogV2.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.TEMPORALCLOUDOPENMETRICSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Name: "temporal-cloud-prod",
				Settings: datadogV2.TemporalCloudOpenMetricsIntegrationAccountSettingsRequest{
					IncludeOptInLabels:  datadog.PtrBool(false),
					NamespacesToInclude: datadog.PtrString("prod-*,billing.a1b2c"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTemporalCloudOpenMetricsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTemporalCloudOpenMetricsIntegrationApi(apiClient)
	resp, r, err := api.CreateTemporalCloudOpenMetricsIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalCloudOpenMetricsIntegrationApi.CreateTemporalCloudOpenMetricsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TemporalCloudOpenMetricsIntegrationApi.CreateTemporalCloudOpenMetricsIntegrationAccount`:\n%s\n", responseContent)
}
