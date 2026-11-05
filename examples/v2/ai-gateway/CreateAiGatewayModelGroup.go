// Create an AI Gateway model group returns "Created" response

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
	body := datadogV2.AiGatewayModelGroupCreateRequest{
		Data: datadogV2.AiGatewayModelGroupCreateData{
			Attributes: datadogV2.AiGatewayModelGroupCreateAttributes{
				AllowFastMode: datadog.PtrBool(true),
				Bindings: []datadogV2.AiGatewayModelGroupBinding{
					{
						Principals: []string{
							"user:55555555-5555-5555-5555-555555555555",
						},
						Relation: datadogV2.AIGATEWAYMODELGROUPBINDINGRELATION_RUNNER,
					},
				},
				MaxReasoningLevel: datadogV2.AIGATEWAYMAXREASONINGLEVEL_HIGH.Ptr(),
				ModelIds: []int64{
					101,
				},
				Name:        "Coding models",
				ProviderIds: []int64{},
				RoutingRuleSlugs: []string{
					"coding-route",
				},
			},
			Type: datadogV2.AIGATEWAYMODELGROUPTYPE_MODEL_GROUPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayModelGroup", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayModelGroup(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayModelGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayModelGroup`:\n%s\n", responseContent)
}
