// Get a RUM rate limit configuration returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetRumRateLimitConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumRateLimitApi(apiClient)
	resp, r, err := api.GetRumRateLimitConfig(ctx, datadogV2.RUMRATELIMITSCOPETYPE_APPLICATION, "cd73a516-a481-4af5-8352-9b577465c77b")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRateLimitApi.GetRumRateLimitConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumRateLimitApi.GetRumRateLimitConfig`:\n%s\n", responseContent)
}
