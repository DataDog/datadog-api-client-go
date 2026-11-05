// Update an AI Gateway budget returns "OK" response

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
	body := datadogV2.AiGatewayBudgetUpdateRequest{
		Data: datadogV2.AiGatewayBudgetUpdateData{
			Attributes: datadogV2.AiGatewayBudgetUpdateAttributes{
				ApplyToCurrentPeriod: datadog.PtrBool(false),
				CapacityMicrocents:   datadog.PtrInt64(120000000),
				Enabled:              datadog.PtrBool(true),
				Enforcement:          datadogV2.AIGATEWAYBUDGETENFORCEMENT_BLOCK.Ptr(),
				Name:                 datadog.PtrString("Coding agents"),
				Reason:               datadog.PtrString("Temporary capacity increase approved for an incident"),
				TrackExisting:        datadog.PtrBool(false),
				WarningThreshold:     datadog.PtrFloat64(0.8),
			},
			Id:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Type: datadogV2.AIGATEWAYBUDGETTYPE_BUDGETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayBudget", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayBudget(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayBudget`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayBudget`:\n%s\n", responseContent)
}
