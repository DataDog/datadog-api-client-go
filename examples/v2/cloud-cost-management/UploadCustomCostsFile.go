// Upload Custom Costs file returns "Accepted" response

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
			BilledCost:        datadog.PtrFloat64(100.5),
			BillingCurrency:   datadog.PtrString("USD"),
			ChargeDescription: datadog.PtrString("Monthly usage charge for my service"),
			ChargePeriodEnd:   datadog.PtrString("2023-02-28"),
			ChargePeriodStart: datadog.PtrString("2023-02-01"),
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
