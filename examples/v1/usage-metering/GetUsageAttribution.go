// Get usage attribution returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetUsageAttribution", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageMeteringApi.GetUsageAttribution(ctx, time.Now().AddDate(0, 0, -3), datadog.USAGEATTRIBUTIONSUPPORTEDMETRICS_ALL, *datadog.NewGetUsageAttributionOptionalParameters().WithOffset(0).WithLimit(1))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAttribution`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageAttribution`:\n%s\n", responseContent)
}
