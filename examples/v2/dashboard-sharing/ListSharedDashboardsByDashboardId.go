// List shared dashboards for a dashboard returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListSharedDashboardsByDashboardId", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardSharingApi(apiClient)
	resp, r, err := api.ListSharedDashboardsByDashboardId(ctx, "abc-def-ghi")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSharingApi.ListSharedDashboardsByDashboardId`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardSharingApi.ListSharedDashboardsByDashboardId`:\n%s\n", responseContent)
}
