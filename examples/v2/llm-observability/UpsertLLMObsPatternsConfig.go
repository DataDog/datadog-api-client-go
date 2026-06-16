// Create or update a patterns configuration returns "OK" response

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
	body := datadogV2.LLMObsPatternsConfigUpsertRequest{
		Data: datadogV2.LLMObsPatternsConfigUpsertRequestData{
			Attributes: datadogV2.LLMObsPatternsConfigUpsertRequestAttributes{
				AccountId:           datadog.PtrString("1000000001"),
				ConfigId:            datadog.PtrString("a7c8d9e0-1234-5678-9abc-def012345678"),
				EvpQuery:            "@ml_app:support-bot",
				HierarchyDepth:      2,
				IntegrationProvider: datadog.PtrString("openai"),
				ModelName:           datadog.PtrString("gpt-4o"),
				Name:                "Support chatbot topics",
				NumRecords:          1000,
				SamplingRatio:       0.1,
				Scope:               datadog.PtrString(""),
				Template:            datadog.PtrString(""),
			},
			Type: datadogV2.LLMOBSPATTERNSCONFIGTYPE_TOPIC_DISCOVERY_CONFIGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpsertLLMObsPatternsConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpsertLLMObsPatternsConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpsertLLMObsPatternsConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpsertLLMObsPatternsConfig`:\n%s\n", responseContent)
}
