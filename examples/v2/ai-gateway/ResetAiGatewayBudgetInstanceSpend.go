// Reset a budget instance's spend returns "OK" response

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
	body := datadogV2.AiGatewayBudgetInstanceResetRequest{
		Data: datadogV2.AiGatewayBudgetInstanceResetData{
			Attributes: datadogV2.AiGatewayBudgetInstanceResetAttributes{
				InstanceKey: "v1|7:user_id=36:22222222-2222-2222-2222-222222222222",
				Reason:      datadog.PtrString("Credit approved after incorrect usage attribution"),
			},
			Type: datadogV2.AIGATEWAYBUDGETINSTANCETYPE_BUDGET_INSTANCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ResetAiGatewayBudgetInstanceSpend", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.ResetAiGatewayBudgetInstanceSpend(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.ResetAiGatewayBudgetInstanceSpend`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.ResetAiGatewayBudgetInstanceSpend`:\n%s\n", responseContent)
}
