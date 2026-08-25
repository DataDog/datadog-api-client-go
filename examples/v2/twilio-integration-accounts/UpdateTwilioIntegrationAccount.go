// Update a Twilio integration account returns "OK" response

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
	body := datadogV2.TwilioIntegrationAccountUpdateRequest{
		Data: datadogV2.TwilioIntegrationAccountUpdateData{
			Attributes: datadogV2.TwilioIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.TwilioIntegrationAccountAuthenticationUpdate{
					IntegrationAccountBasicAuthUpdate: &datadogV2.IntegrationAccountBasicAuthUpdate{
						AuthType: datadogV2.INTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
						Password: datadog.PtrString("your-password"),
						Username: datadog.PtrString("datadog"),
					}},
				Dataflows: &datadogV2.TwilioIntegrationDataflowsRequest{
					TwilioAlertsLogs: &datadogV2.TwilioAlertsLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					TwilioCallSummariesLogs: &datadogV2.TwilioCallSummariesLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					TwilioCloudCostMetrics: &datadogV2.TwilioCloudCostMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					TwilioEventsLogs: &datadogV2.TwilioEventsLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					TwilioMessagesLogs: &datadogV2.TwilioMessagesLogsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("twilio-prod"),
				Settings: &datadogV2.TwilioIntegrationAccountSettingsUpdate{
					AccountSid: datadog.PtrString("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					CensorLogs: datadog.PtrBool(true),
				},
			},
			Id:   "953a0060-81ec-4221-aed4-d4733b59cd96",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTwilioIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTwilioIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateTwilioIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwilioIntegrationAccountsApi.UpdateTwilioIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TwilioIntegrationAccountsApi.UpdateTwilioIntegrationAccount`:\n%s\n", responseContent)
}
