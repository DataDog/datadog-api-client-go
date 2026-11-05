// Create an AI Gateway budget returns "Created" response

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
	body := datadogV2.AiGatewayBudgetCreateRequest{
		Data: datadogV2.AiGatewayBudgetCreateData{
			Attributes: datadogV2.AiGatewayBudgetCreateAttributes{
				CapacityMicrocents: 100000000,
				Enforcement:        datadogV2.AIGATEWAYBUDGETENFORCEMENT_BLOCK,
				Name:               "Coding agents",
				Selector: []datadogV2.AiGatewayBudgetSelectorEntry{
					{
						Dim:      "user_id",
						PerValue: true,
						Values: []string{
							"11111111-1111-1111-1111-111111111111",
						},
					},
				},
				Tags: map[string]string{
					"model": "gpt-5.2-codex",
				},
				TrackExisting:    datadog.PtrBool(false),
				WarningThreshold: datadog.PtrFloat64(0.9),
				WindowPolicy:     datadogV2.AIGATEWAYBUDGETWINDOWPOLICY_MONTHLY,
			},
			Type: datadogV2.AIGATEWAYBUDGETTYPE_BUDGETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAiGatewayBudget", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.CreateAiGatewayBudget(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.CreateAiGatewayBudget`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.CreateAiGatewayBudget`:\n%s\n", responseContent)
}
