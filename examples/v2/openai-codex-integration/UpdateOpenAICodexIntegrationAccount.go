// Update an OpenAI Codex integration account returns "OK" response

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
	body := datadogV2.OpenAICodexIntegrationAccountUpdateRequest{
		Data: datadogV2.OpenAICodexIntegrationAccountUpdateData{
			Attributes: datadogV2.OpenAICodexIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.OpenAICodexIntegrationAccountAuthenticationUpdate{
					OpenAICodexIntegrationAccountBearerTokenAuthUpdate: &datadogV2.OpenAICodexIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.OpenAICodexIntegrationDataflowsRequest{
					OpenaiCodexDora: &datadogV2.OpenAICodexDoraIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiCodexTurnMetrics: &datadogV2.OpenAICodexTurnMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("openai-codex-acme"),
				Settings: &datadogV2.OpenAICodexIntegrationAccountSettingsUpdate{
					CodexWorkspaceId: datadog.PtrString("your-workspace-id"),
				},
			},
			Id:   "7d2f9c4e-3a1b-4c8e-9f6d-2b5a8e1c0d47",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOpenAICodexIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAICodexIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpenAICodexIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICodexIntegrationApi.UpdateOpenAICodexIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAICodexIntegrationApi.UpdateOpenAICodexIntegrationAccount`:\n%s\n", responseContent)
}
