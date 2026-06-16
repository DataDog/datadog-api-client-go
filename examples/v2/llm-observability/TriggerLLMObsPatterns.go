// Trigger a patterns run returns "Accepted" response

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
	body := datadogV2.LLMObsPatternsTriggerRequest{
		Data: datadogV2.LLMObsPatternsTriggerRequestData{
			Attributes: datadogV2.LLMObsPatternsTriggerRequestAttributes{
				ConfigId: "a7c8d9e0-1234-5678-9abc-def012345678",
			},
			Type: datadogV2.LLMOBSPATTERNSREQUESTTYPE_TOPIC_DISCOVERY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.TriggerLLMObsPatterns", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.TriggerLLMObsPatterns(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.TriggerLLMObsPatterns`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.TriggerLLMObsPatterns`:\n%s\n", responseContent)
}
