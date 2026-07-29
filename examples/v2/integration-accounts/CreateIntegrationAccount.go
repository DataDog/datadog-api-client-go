// Create an integration account returns "Created" response

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
	body := datadogV2.IntegrationAccountRequest{
		Data: datadogV2.IntegrationAccountCreateData{
			Attributes: datadogV2.IntegrationAccountAttributes{
				Integration: datadogV2.IntegrationAccountIntegration{
					TwilioIntegration: &datadogV2.TwilioIntegration{
						Interface: datadogV2.TwilioInterface{
							Authentication: datadogV2.TwilioAuthentication{
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
							Settings: &datadogV2.TwilioSettings{
								AccountSid: "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
								CensorLogs: datadog.PtrBool(true),
							},
							Type: datadogV2.TWILIOINTERFACETYPE_TWILIO,
						},
						Type: datadogV2.TWILIOINTEGRATIONTYPE_TWILIO,
					}},
				Name: "twilio-prod",
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	resp, r, err := api.CreateIntegrationAccount(ctx, datadogV2.INTEGRATIONACCOUNTINTEGRATIONID_TWILIO, datadogV2.INTEGRATIONACCOUNTINTERFACEID_TWILIO, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.CreateIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAccountsApi.CreateIntegrationAccount`:\n%s\n", responseContent)
}
