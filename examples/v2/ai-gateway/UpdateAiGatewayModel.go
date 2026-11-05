// Update a model for an AI Gateway provider returns "OK" response

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
	body := datadogV2.AiGatewayModelUpdateRequest{
		Data: datadogV2.AiGatewayModelUpdateData{
			Attributes: datadogV2.AiGatewayModelUpdateAttributes{
				ModelIdWire: datadog.PtrString("gpt-5.2-codex-20260801"),
				SupportedOperations: []datadogV2.AiGatewaySupportedOperation{
					datadogV2.AIGATEWAYSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
				},
			},
			Id:   "101",
			Type: datadogV2.AIGATEWAYMODELCATALOGENTRYTYPE_MODEL_CATALOG_ENTRIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayModel", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayModel(ctx, 9223372036854775807, "canonical", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayModel`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayModel`:\n%s\n", responseContent)
}
