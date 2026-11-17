// Update a Product Analytics launch returns "OK" response

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
	body := datadogV2.ProductAnalyticsLaunchPatchRequest{
		Data: datadogV2.ProductAnalyticsLaunchPatchData{
			Attributes: datadogV2.ProductAnalyticsLaunchPatchAttributes{
				Phase: datadogV2.PRODUCTANALYTICSLAUNCHPHASE_SETUP.Ptr(),
			},
			Id:   "5e3c7a1b-2d4f-4e6a-8b9c-0d1e2f3a4b5c",
			Type: datadogV2.PRODUCTANALYTICSLAUNCHTYPE_LIFTOFF,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateProductAnalyticsLaunch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.UpdateProductAnalyticsLaunch(ctx, "5e3c7a1b-2d4f-4e6a-8b9c-0d1e2f3a4b5c", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.UpdateProductAnalyticsLaunch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.UpdateProductAnalyticsLaunch`:\n%s\n", responseContent)
}
