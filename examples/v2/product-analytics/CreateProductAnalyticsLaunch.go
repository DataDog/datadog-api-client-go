// Create a Product Analytics launch returns "Created" response

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
	body := datadogV2.ProductAnalyticsLaunchCreateRequest{
		Data: datadogV2.ProductAnalyticsLaunchCreateData{
			Attributes: datadogV2.ProductAnalyticsLaunchRequestAttributes{
				Context:     "Let users apply promo codes at checkout.",
				Description: *datadog.NewNullableString(datadog.PtrString("Promo codes at checkout.")),
				Kpis: []datadogV2.ProductAnalyticsLaunchKPI{
					{
						Id:    "7c1e6a52-3a8f-4d4e-9a55-2f1d1c8b6a10",
						Title: "Checkout conversion rate",
					},
				},
				Name:  "checkout_promo_codes",
				Phase: datadogV2.PRODUCTANALYTICSLAUNCHPHASE_PLAN.Ptr(),
				Steps: datadogV2.ProductAnalyticsLaunchStepsRequest{
					TrackingPlan: datadogV2.ProductAnalyticsLaunchTrackingPlanRequest{
						Events: []datadogV2.ProductAnalyticsLaunchEvent{
							{
								Kpis: []string{
									"7c1e6a52-3a8f-4d4e-9a55-2f1d1c8b6a10",
								},
								Name:    "promo_code_applied",
								Trigger: datadog.PtrString("The user applies a valid promo code."),
								Views: []string{
									"/checkout",
								},
							},
						},
					},
				},
			},
			Type: datadogV2.PRODUCTANALYTICSLAUNCHTYPE_LIFTOFF,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateProductAnalyticsLaunch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.CreateProductAnalyticsLaunch(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.CreateProductAnalyticsLaunch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.CreateProductAnalyticsLaunch`:\n%s\n", responseContent)
}
