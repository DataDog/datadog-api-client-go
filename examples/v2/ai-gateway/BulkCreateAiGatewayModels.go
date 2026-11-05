// Add multiple models to an AI Gateway provider returns "Created" response

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
	body := datadogV2.AiGatewayModelsBulkCreateRequest{
		Data: datadogV2.AiGatewayModelsBulkCreateData{
			Attributes: datadogV2.AiGatewayModelsBulkCreateAttributes{
				Models: []datadogV2.AiGatewayModelCreateAttributes{
					{
						ModelIdCanonical: "gpt-5.2-codex",
						ModelIdWire:      "gpt-5.2-codex-20260801",
						SupportedOperations: []datadogV2.AiGatewaySupportedOperation{
							datadogV2.AIGATEWAYSUPPORTEDOPERATION_OPENAI_CHAT_COMPLETIONS,
						},
					},
				},
			},
			Type: datadogV2.AIGATEWAYMODELCATALOGENTRYTYPE_MODEL_CATALOG_ENTRIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BulkCreateAiGatewayModels", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.BulkCreateAiGatewayModels(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.BulkCreateAiGatewayModels`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.BulkCreateAiGatewayModels`:\n%s\n", responseContent)
}
