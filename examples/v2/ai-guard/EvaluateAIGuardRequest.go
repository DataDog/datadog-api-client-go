// Evaluate an AI Guard request returns "OK" response

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
		Data: datadogV2.AIGuardEvaluateRequestData{
			Type: datadogV2.AIGUARDEVALUATIONTYPE_EVALUATIONS,
			Attributes: datadogV2.AIGuardEvaluateRequestDataAttributes{
				Messages: []datadogV2.AIGuardMessage{
					{
						Content: &datadogV2.AIGuardMessageContent{
							AIGuardMessageContentText: datadog.PtrString("Hello world!")},
						Role: datadogV2.AIGUARDMESSAGEROLE_USER,
					},
				},
				Meta: &datadogV2.AIGuardMeta{
					ConfidenceThreshold: &datadogV2.AIGuardConfidenceThresholdOverride{
						AIGuardGlobalConfidenceThreshold: datadog.PtrFloat64(0.8)},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.EvaluateAIGuardRequest", true)
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
