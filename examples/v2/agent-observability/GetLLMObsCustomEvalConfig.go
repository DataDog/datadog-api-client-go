// Get a custom evaluator configuration returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetLLMObsCustomEvalConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentObservabilityApi(apiClient)
	resp, r, err := api.GetLLMObsCustomEvalConfig(ctx, "eval_name")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentObservabilityApi.GetLLMObsCustomEvalConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentObservabilityApi.GetLLMObsCustomEvalConfig`:\n%s\n", responseContent)
}
