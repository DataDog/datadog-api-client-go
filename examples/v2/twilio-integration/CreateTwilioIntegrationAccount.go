// Create a Twilio integration account returns "Created" response

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
	body := datadogV2.TwilioIntegrationAccountCreateRequest{
		Data: datadogV2.TwilioIntegrationAccountCreateData{
			Attributes: datadogV2.TwilioIntegrationAccountCreateAttributes{
				Authentication: datadogV2.TwilioIntegrationAccountAuthenticationRequest{
					TwilioIntegrationAccountBasicAuthRequest: &datadogV2.TwilioIntegrationAccountBasicAuthRequest{
						AuthType: datadogV2.TWILIOINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
						Password: "your-password",
						Username: "datadog",
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
				Name: "twilio-prod",
				Settings: datadogV2.TwilioIntegrationAccountSettingsRequest{
					AccountSid: "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					CensorLogs: datadog.PtrBool(true),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTwilioIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTwilioIntegrationApi(apiClient)
	resp, r, err := api.CreateTwilioIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwilioIntegrationApi.CreateTwilioIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TwilioIntegrationApi.CreateTwilioIntegrationAccount`:\n%s\n", responseContent)
}
