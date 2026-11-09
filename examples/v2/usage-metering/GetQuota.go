// Get a usage quota returns "OK" response

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
	// there is a valid "quota" in the system
	QuotaData0ID := os.Getenv("QUOTA_DATA_0_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateQuotas", true)
	configuration.SetUnstableOperationEnabled("v2.DeleteQuota", true)
	configuration.SetUnstableOperationEnabled("v2.GetQuota", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetQuota(ctx, "ai_credits", QuotaData0ID, *datadogV2.NewGetQuotaOptionalParameters().WithIncludeDescendants(false))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetQuota`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetQuota`:\n%s\n", responseContent)
}
