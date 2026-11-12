// Create an AI Gateway routing rule returns "Created" response

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
	body := datadogV2.AiGatewayRoutingRuleCreateRequest{
		Data: datadogV2.AiGatewayRoutingRuleCreateData{
			Attributes: datadogV2.AiGatewayRoutingRuleCreateAttributes{
				Slug: "coding",
				Tiers: []datadogV2.AiGatewayRoutingRuleTierRequest{
					{
						Classifier: *datadogV2.NewNullableAiGatewayRoutingClassifierRequest(&datadogV2.AiGatewayRoutingClassifierRequest{
							ConfidenceThreshold: *datadog.NewNullableFloat64(datadog.PtrFloat64(0.5)),
							Context:             datadog.PtrString("Requests come from engineers."),
							Model:               datadog.PtrString("gpt-6-luna"),
							Provider:            datadog.PtrString("openai"),
						}),
						Ordering: *datadogV2.NewNullableAiGatewayNullableRoutingRuleTierOrdering(datadogV2.AIGATEWAYNULLABLEROUTINGRULETIERORDERING_CONFIGURED.Ptr()),
						Targets: []datadogV2.AiGatewayRoutingRuleTargetRequest{
							{
								Default:        *datadog.NewNullableBool(datadog.PtrBool(false)),
								Id:             datadog.PtrInt64(901),
								Model:          datadog.PtrString("gpt-5.2-codex"),
								ModelId:        datadog.PtrInt64(101),
								Provider:       datadog.PtrString("prod-us"),
								ReasoningLevel: *datadogV2.NewNullableAiGatewayRoutingReasoningLevel(datadogV2.AIGATEWAYROUTINGREASONINGLEVEL_MEDIUM.Ptr()),
								Rubric:         *datadog.NewNullableString(datadog.PtrString("Short factual questions.")),
								Slug:           datadog.PtrString("coding-fallback"),
								Type:           datadogV2.AIGATEWAYROUTINGRULETARGETTYPE_MODEL,
								Weight:         *datadog.NewNullableInt64(datadog.PtrInt64(70)),
							},
						},
					},
				},
			},
			Type: datadogV2.AIGATEWAYROUTINGRULETYPE_ROUTING_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayRoutingRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayRoutingRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayRoutingRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayRoutingRule`:\n%s\n", responseContent)
}
