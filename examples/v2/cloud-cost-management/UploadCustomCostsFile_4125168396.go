// Upload Custom Costs File returns "Accepted" response

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
	body := []datadogV2.CustomCostsFileLineItem{
		{
			ProviderName:      datadog.PtrString("my_provider"),
			ChargePeriodStart: datadog.PtrString("2023-05-06"),
			ChargePeriodEnd:   datadog.PtrString("2023-06-06"),
			ChargeDescription: datadog.PtrString("my_description"),
			BilledCost:        datadog.PtrFloat64(250),
			BillingCurrency:   datadog.PtrString("USD"),
			Tags: map[string]string{
				"key": "value",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCloudCostManagementApi(apiClient)
	resp, r, err := api.UploadCustomCostsFile(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCostManagementApi.UploadCustomCostsFile`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudCostManagementApi.UploadCustomCostsFile`:\n%s\n", responseContent)
}
