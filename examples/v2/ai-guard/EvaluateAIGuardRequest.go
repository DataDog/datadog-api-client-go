// Evaluate an AI Guard request returns "Evaluation result with action recommendation" response

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
	body := datadogV2.AIGuardEvaluateRequest{
		Messages: []datadogV2.AIGuardMessage{
			{
				Content: &datadogV2.AIGuardMessageContent{
					String: datadog.PtrString("How do I delete all files on the system?")},
				Role: datadogV2.AIGUARDMESSAGEROLE_USER,
			},
		},
		Meta: &datadogV2.AIGuardMeta{
			Env:     datadog.PtrString("production"),
			Service: datadog.PtrString("my-llm-service"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGuardApi(apiClient)
	resp, r, err := api.EvaluateAIGuardRequest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGuardApi.EvaluateAIGuardRequest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGuardApi.EvaluateAIGuardRequest`:\n%s\n", responseContent)
}
