// Update a Temporal Cloud OpenMetrics account returns "OK" response

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
	body := datadogV2.TemporalCloudOpenMetricsIntegrationAccountUpdateRequest{
		Data: datadogV2.TemporalCloudOpenMetricsIntegrationAccountUpdateData{
			Attributes: datadogV2.TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate{
					TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate: &datadogV2.TemporalCloudOpenMetricsIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.TEMPORALCLOUDOPENMETRICSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Name: datadog.PtrString("temporal-cloud-prod"),
				Settings: &datadogV2.TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate{
					IncludeOptInLabels:  datadog.PtrBool(false),
					NamespacesToInclude: datadog.PtrString("prod-*,billing.a1b2c"),
				},
			},
			Id:   "7c2d4e91-5a3b-4f68-b0e2-9d1f6a8c3e57",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTemporalCloudOpenMetricsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTemporalCloudOpenMetricsIntegrationApi(apiClient)
	resp, r, err := api.UpdateTemporalCloudOpenMetricsIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalCloudOpenMetricsIntegrationApi.UpdateTemporalCloudOpenMetricsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TemporalCloudOpenMetricsIntegrationApi.UpdateTemporalCloudOpenMetricsIntegrationAccount`:\n%s\n", responseContent)
}
