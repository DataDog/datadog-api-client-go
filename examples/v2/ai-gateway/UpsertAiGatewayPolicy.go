// Create or replace an AI Gateway routing policy returns "OK" response

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
	body := datadogV2.AiGatewayPolicyRequest{
		Data: datadogV2.AiGatewayPolicyRequestData{
			Attributes: datadogV2.AiGatewayPolicyRequestAttributes{
				Inbound:  "coding-agent",
				Priority: datadog.PtrInt64(1),
				Tiers: []datadogV2.AiGatewayPolicyTierRequest{
					{
						Priority: datadog.PtrInt64(1),
						Targets: []datadogV2.AiGatewayPolicyTargetRequest{
							{
								Id:       datadog.PtrInt64(901),
								Model:    datadog.PtrString("gpt-5.2-codex"),
								Provider: datadog.PtrString("prod-us"),
							},
						},
					},
				},
			},
			Type: datadogV2.AIGATEWAYPOLICYTYPE_ROUTING_DOCUMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertAiGatewayPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpsertAiGatewayPolicy(ctx, body, *datadogV2.NewUpsertAiGatewayPolicyOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpsertAiGatewayPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpsertAiGatewayPolicy`:\n%s\n", responseContent)
}
