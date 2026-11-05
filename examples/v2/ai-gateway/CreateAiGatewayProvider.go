// Create an AI Gateway provider returns "Created" response

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
	body := datadogV2.AiGatewayProviderCreateRequest{
		Data: datadogV2.AiGatewayProviderCreateData{
			Attributes: datadogV2.AiGatewayProviderCreateAttributes{
				CostMultiplier: datadog.PtrFloat64(1.0),
				InitialBindings: []datadogV2.AiGatewayProviderInitialBinding{
					{
						Principals: []string{
							"org:11111111-1111-1111-1111-111111111111",
						},
						Relation: datadogV2.AIGATEWAYPROVIDERINITIALBINDINGRELATION_VIEWER,
					},
				},
				IntegrationAccountId: datadog.PtrString("account-123"),
				IntegrationConfig: &datadogV2.AiGatewayProviderIntegrationConfig{
					BedrockRegion:     datadog.PtrString("us-east-1"),
					FoundryProject:    datadog.PtrString("my-project"),
					FoundryResourceId: datadog.PtrString("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ai/providers/Microsoft.CognitiveServices/accounts/my-foundry"),
					Location:          datadog.PtrString("us-central1"),
					ProjectId:         datadog.PtrString("my-gcp-project"),
					SupportedOperations: []datadogV2.AiGatewayCustomProviderSupportedOperation{
						datadogV2.AIGATEWAYCUSTOMPROVIDERSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
					},
				},
				KnownProvider: datadogV2.AIGATEWAYKNOWNPROVIDER_OPENAI.Ptr(),
				Models: []datadogV2.AiGatewayModelInput{
					{
						ModelIdCanonical: "gpt-5.2-codex",
						ModelIdWire:      "gpt-5.2-codex-20260801",
						SupportedOperations: []datadogV2.AiGatewaySupportedOperation{
							datadogV2.AIGATEWAYSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
						},
					},
				},
				Name: "production-us",
				Slug: "prod-us",
			},
			Type: datadogV2.AIGATEWAYPROVIDERTYPE_PROVIDERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayProvider", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayProvider(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayProvider`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayProvider`:\n%s\n", responseContent)
}
