// Update a RUM application with Product Scales returns "OK" response

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
	// there is a valid "rum_application" in the system
	RumApplicationDataID := os.Getenv("RUM_APPLICATION_DATA_ID")

	body := datadogV2.RUMApplicationUpdateRequest{
		Data: datadogV2.RUMApplicationUpdate{
			Attributes: &datadogV2.RUMApplicationUpdateAttributes{
				Name:                           datadog.PtrString("updated_rum_with_product_scales"),
				RumEventProcessingState:        datadogV2.RUMEVENTPROCESSINGSTATE_ALL.Ptr(),
				ProductAnalyticsRetentionState: datadogV2.RUMPRODUCTANALYTICSRETENTIONSTATE_MAX.Ptr(),
			},
			Id:   RumApplicationDataID,
			Type: datadogV2.RUMAPPLICATIONUPDATETYPE_RUM_APPLICATION_UPDATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, r, err := api.UpdateRUMApplication(ctx, RumApplicationDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.UpdateRUMApplication`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.UpdateRUMApplication`:\n%s\n", responseContent)
}
