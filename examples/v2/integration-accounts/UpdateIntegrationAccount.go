// Update an integration account returns "OK" response

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
	body := datadogV2.IntegrationAccountUpdateRequest{
		Data: datadogV2.IntegrationAccountUpdateData{
			Attributes: datadogV2.IntegrationAccountUpdateAttributes{
				Integration: &datadogV2.IntegrationAccountIntegrationUpdate{
					TwilioIntegrationUpdate: &datadogV2.TwilioIntegrationUpdate{
						Interface: &datadogV2.TwilioInterfaceUpdate{
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
							Settings: &datadogV2.TwilioSettingsUpdate{
								AccountSid: datadog.PtrString("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
								CensorLogs: datadog.PtrBool(true),
							},
							Type: datadogV2.TWILIOINTERFACETYPE_TWILIO,
						},
						Type: datadogV2.TWILIOINTEGRATIONTYPE_TWILIO,
					}},
				Name: datadog.PtrString("twilio-prod"),
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateIntegrationAccount(ctx, datadogV2.INTEGRATIONACCOUNTINTEGRATIONID_TWILIO, datadogV2.INTEGRATIONACCOUNTINTERFACEID_TWILIO, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.UpdateIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAccountsApi.UpdateIntegrationAccount`:\n%s\n", responseContent)
}
