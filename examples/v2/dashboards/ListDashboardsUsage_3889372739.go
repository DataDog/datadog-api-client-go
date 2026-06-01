// Get usage stats for all dashboards with edited_before filter returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListDashboardsUsage", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardsApi(apiClient)
	resp, r, err := api.ListDashboardsUsage(ctx, *datadogV2.NewListDashboardsUsageOptionalParameters().WithFilterEditedBefore("2025-04-26T00:00:00Z"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboardsUsage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.ListDashboardsUsage`:\n%s\n", responseContent)
}
