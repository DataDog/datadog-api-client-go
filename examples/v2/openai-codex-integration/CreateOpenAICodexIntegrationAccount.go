// Create an OpenAI Codex integration account returns "Created" response

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
	body := datadogV2.OpenAICodexIntegrationAccountCreateRequest{
		Data: datadogV2.OpenAICodexIntegrationAccountCreateData{
			Attributes: datadogV2.OpenAICodexIntegrationAccountCreateAttributes{
				Authentication: datadogV2.OpenAICodexIntegrationAccountAuthenticationRequest{
					OpenAICodexIntegrationAccountBearerTokenAuthRequest: &datadogV2.OpenAICodexIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.OpenAICodexIntegrationDataflowsRequest{
					OpenaiCodexDora: &datadogV2.OpenAICodexDoraIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiCodexTurnMetrics: &datadogV2.OpenAICodexTurnMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "openai-codex-acme",
				Settings: datadogV2.OpenAICodexIntegrationAccountSettingsRequest{
					CodexWorkspaceId: "your-workspace-id",
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOpenAICodexIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAICodexIntegrationApi(apiClient)
	resp, r, err := api.CreateOpenAICodexIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICodexIntegrationApi.CreateOpenAICodexIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAICodexIntegrationApi.CreateOpenAICodexIntegrationAccount`:\n%s\n", responseContent)
}
