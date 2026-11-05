// Replace an AI Gateway routing rule returns "OK" response

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
	body := datadogV2.AiGatewayRoutingRuleReplaceRequest{
		Data: datadogV2.AiGatewayRoutingRuleReplaceData{
			Attributes: datadogV2.AiGatewayRoutingRuleReplaceAttributes{
				Tiers: []datadogV2.AiGatewayRoutingRuleTierRequest{
					{
						Ordering: *datadogV2.NewNullableAiGatewayNullableRoutingRuleTierOrdering(datadogV2.AIGATEWAYNULLABLEROUTINGRULETIERORDERING_CONFIGURED.Ptr()),
						Targets: []datadogV2.AiGatewayRoutingRuleTargetRequest{
							{
								Id:             datadog.PtrInt64(901),
								Model:          datadog.PtrString("gpt-5.2-codex"),
								ModelId:        datadog.PtrInt64(101),
								Provider:       datadog.PtrString("prod-us"),
								ReasoningLevel: *datadogV2.NewNullableAiGatewayRoutingReasoningLevel(datadogV2.AIGATEWAYROUTINGREASONINGLEVEL_MEDIUM.Ptr()),
								Slug:           datadog.PtrString("coding-fallback"),
								Type:           datadogV2.AIGATEWAYROUTINGRULETARGETTYPE_MODEL,
								Weight:         *datadog.NewNullableInt64(datadog.PtrInt64(70)),
							},
						},
					},
				},
			},
			Id:   "coding",
			Type: datadogV2.AIGATEWAYROUTINGRULETYPE_ROUTING_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayRoutingRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayRoutingRule(ctx, "slug", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayRoutingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayRoutingRule`:\n%s\n", responseContent)
}
