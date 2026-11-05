// Override a budget instance's capacity returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.AiGatewayBudgetInstanceCapacityOverrideRequest{
		Data: datadogV2.AiGatewayBudgetInstanceCapacityOverrideData{
			Attributes: datadogV2.AiGatewayBudgetInstanceCapacityOverrideAttributes{
				CapacityMicrocents: 250000000,
				InstanceKey:        "v1|7:user_id=36:22222222-2222-2222-2222-222222222222",
				Reason:             datadog.PtrString("Temporary increase for a launch week"),
			},
			Type: datadogV2.AIGATEWAYBUDGETINSTANCETYPE_BUDGET_INSTANCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.OverrideAiGatewayBudgetInstanceCapacity", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.OverrideAiGatewayBudgetInstanceCapacity(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.OverrideAiGatewayBudgetInstanceCapacity`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.OverrideAiGatewayBudgetInstanceCapacity`:\n%s\n", responseContent)
}
