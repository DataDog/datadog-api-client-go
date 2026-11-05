// Set a model pricing override returns "OK" response

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
	body := datadogV2.AiGatewayModelPricingRequest{
		Data: datadogV2.AiGatewayModelPricingRequestData{
			Attributes: datadogV2.AiGatewayModelPricingRequestAttributes{
				PricingConfig: datadogV2.AiGatewayModelPricingRequestConfig{
					AiGatewayModelPricingCustomRequestConfig: &datadogV2.AiGatewayModelPricingCustomRequestConfig{
						Rates: datadogV2.AiGatewayModelPricingRequiredRates{
							CacheReadCostPerMtok:      0.0,
							CacheWriteCostPerMtok:     0.0,
							NonCachedInputCostPerMtok: 0.0,
							OutputCostPerMtok:         0.0,
						},
						Source: datadogV2.AiGatewayModelPricingCustomRequestSource{
							Type: datadogV2.AIGATEWAYMODELPRICINGCUSTOMREQUESTSOURCETYPE_CUSTOM,
						},
						Version: 1,
					}},
			},
			Id:   "gpt-4o",
			Type: datadogV2.AIGATEWAYMODELPRICINGTYPE_MODEL_PRICING,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAiGatewayModelPricing", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAIGatewayApi(apiClient)
	resp, r, err := api.UpdateAiGatewayModelPricing(ctx, 9223372036854775807, "canonical", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayApi.UpdateAiGatewayModelPricing`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AIGatewayApi.UpdateAiGatewayModelPricing`:\n%s\n", responseContent)
}
