// Update Custom Costs file returns "OK" response

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
	// there is a valid "custom_costs_file" in the system
	CustomCostsFileDataID := os.Getenv("CUSTOM_COSTS_FILE_DATA_ID")

	body := datadogV2.CustomCostsFileUpdateRequest{
		Data: datadogV2.CustomCostsFileUpdateData{
			Attributes: datadogV2.CustomCostsFileUpdateAttributes{
				AiProvider: true,
			},
			Id:   CustomCostsFileDataID,
			Type: datadogV2.CUSTOMCOSTSFILEUPDATETYPE_COST_METADATA,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UpdateCustomCostsFile(ctx, CustomCostsFileDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UpdateCustomCostsFile`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UpdateCustomCostsFile`:\n%s\n", responseContent)
}
