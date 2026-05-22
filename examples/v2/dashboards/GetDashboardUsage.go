// Get usage stats for a dashboard returns "OK" response

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
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetDashboardUsage", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardsApi(apiClient)
	resp, r, err := api.GetDashboardUsage(ctx, DashboardID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboardUsage`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboardUsage`:\n%s\n", responseContent)
}
