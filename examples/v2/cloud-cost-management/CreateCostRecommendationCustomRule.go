// Create a cost recommendation custom rule returns "Created" response

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
	body := datadogV2.CostRecommendationCustomRuleCreateRequest{
		Data: datadogV2.CostRecommendationCustomRuleCreateData{
			Attributes: datadogV2.CostRecommendationCustomRuleCreateAttributes{
				Enabled: datadog.PtrBool(true),
				Filters: []datadogV2.CostRecommendationCustomRuleFilter{
					{
						FilterType: datadogV2.COSTRECOMMENDATIONCUSTOMRULEFILTERTYPE_IS,
						Tag:        "env",
						Values: []string{
							"staging",
						},
					},
				},
				Name:          "Staging instances",
				Notes:         datadog.PtrString("Staging instances can run with less headroom."),
				RecommenderId: "aws_ec2_instance_unused",
				Settings: map[string]float64{
					"lookback_days":      14,
					"max_cpu_user_limit": 0.05,
				},
			},
			Type: datadogV2.COSTRECOMMENDATIONCUSTOMRULETYPE_CUSTOM_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateCostRecommendationCustomRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.CreateCostRecommendationCustomRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.CreateCostRecommendationCustomRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.CreateCostRecommendationCustomRule`:\n%s\n", responseContent)
}
