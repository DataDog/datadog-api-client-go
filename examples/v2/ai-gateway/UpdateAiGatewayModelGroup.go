// Update an AI Gateway model group returns "OK" response

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
	body := datadogV2.AiGatewayModelGroupUpdateRequest{
		Data: datadogV2.AiGatewayModelGroupUpdateData{
			Attributes: datadogV2.AiGatewayModelGroupUpdateAttributes{
				AllowFastMode:     datadog.PtrBool(true),
				MaxReasoningLevel: *datadogV2.NewNullableAiGatewayNullableMaxReasoningLevel(datadogV2.AIGATEWAYNULLABLEMAXREASONINGLEVEL_HIGH.Ptr()),
				ModelIds: []int64{
					101,
				},
				Name:        datadog.PtrString("Coding models"),
				ProviderIds: []int64{},
				RoutingRuleSlugs: []string{
					"coding-route",
				},
			},
			Id:   "7",
			Type: datadogV2.AIGATEWAYMODELGROUPTYPE_MODEL_GROUPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayModelGroup", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayModelGroup(ctx, "id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayModelGroup`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayModelGroup`:\n%s\n", responseContent)
}
