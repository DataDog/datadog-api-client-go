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
	body := datadogV2.TwilioAccountUpdateRequest{
		Data: datadogV2.TwilioAccountUpdateData{
			Attributes: datadogV2.TwilioAccountUpdateAttributes{
				Authentication: &datadogV2.TwilioAuthentication{
					TwilioBasicAuth: &datadogV2.TwilioBasicAuth{
						ApiKey:      "SKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
						ApiKeyToken: "your-api-key-secret",
						Type:        datadogV2.TWILIOBASICAUTHTYPE_BASIC,
					}},
				Dataflows: []datadogV2.TwilioDataflow{
					{
						Enabled: datadog.PtrBool(true),
						Id:      datadogV2.TWILIODATAFLOWID_MESSAGES_LOGS,
					},
				},
				Name: datadog.PtrString("twilio-prod"),
				Settings: &datadogV2.TwilioSettingsUpdate{
					AccountSid: datadog.PtrString("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					CensorLogs: datadog.PtrBool(true),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateTwilioAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTwilioIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateTwilioAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwilioIntegrationAccountsApi.UpdateTwilioAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TwilioIntegrationAccountsApi.UpdateTwilioAccount`:\n%s\n", responseContent)
}
