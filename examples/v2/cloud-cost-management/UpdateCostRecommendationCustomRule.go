// Update a cost recommendation custom rule returns "OK" response

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
	body := datadogV2.CostRecommendationCustomRuleUpdateRequest{
		Data: datadogV2.CostRecommendationCustomRuleUpdateData{
			Attributes: datadogV2.CostRecommendationCustomRuleUpdateAttributes{
				Enabled: datadog.PtrBool(false),
				Filters: []datadogV2.CostRecommendationCustomRuleFilter{
					{
						FilterType: datadogV2.COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS,
						Tag:        "env",
						Values: []string{
							"staging",
						},
					},
				},
				Name:  datadog.PtrString("Staging and development instances"),
				Notes: datadog.PtrString("Staging instances can run with less headroom."),
				Settings: map[string]float64{
					"lookback_days":      14,
					"max_cpu_user_limit": 0.05,
				},
			},
			Id:   "8a1b6f3e-2c4d-4e5f-9a7b-1c2d3e4f5a6b",
			Type: datadogV2.COSTRECOMMENDATIONCUSTOMRULETYPE_CUSTOM_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateCostRecommendationCustomRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCostRecommendationCustomRule(ctx, "custom_rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCostRecommendationCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCostRecommendationCustomRule`:\n%s\n", responseContent)
}
