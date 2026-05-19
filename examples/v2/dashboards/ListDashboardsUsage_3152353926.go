// Get usage stats for all dashboards returns "OK" response with pagination

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
	resp, _ := api.ListDashboardsUsageWithPagination(ctx, *datadogV2.NewListDashboardsUsageOptionalParameters().WithPageLimit(500))

	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboardsUsage`: %v\n", paginationResult.Error)
		}
		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
