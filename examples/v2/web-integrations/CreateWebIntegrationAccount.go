// Create integration account returns "Created" response

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
	body := datadogV2.WebIntegrationAccountCreateRequest{
		Data: datadogV2.WebIntegrationAccountCreateRequestData{
			Type: datadogV2.WEBINTEGRATIONACCOUNTTYPE_ACCOUNT,
			Attributes: datadogV2.WebIntegrationAccountCreateRequestAttributes{
				Name: "Example-Web-Integration",
				Settings: map[string]interface{}{
					"api_key":        "SKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					"account_sid":    "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					"events":         "True",
					"messages":       "True",
					"alerts":         "True",
					"call_summaries": "True",
					"ccm_enabled":    "True",
					"censor_logs":    "True",
				},
				Secrets: map[string]interface{}{
					"api_key_token": "test_secret_token",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebIntegrationsApi(apiClient)
	resp, r, err := api.CreateWebIntegrationAccount(ctx, "twilio", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebIntegrationsApi.CreateWebIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebIntegrationsApi.CreateWebIntegrationAccount`:\n%s\n", responseContent)
}
