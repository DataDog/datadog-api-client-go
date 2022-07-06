// Restore deleted dashboards returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	body := datadog.DashboardRestoreRequest{
		Data: []datadog.DashboardBulkActionData{
			{
				Id:   DashboardID,
				Type: datadog.DASHBOARDRESOURCETYPE_DASHBOARD,
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.DashboardsApi(apiClient)
	r, err := api.RestoreDashboards(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.RestoreDashboards`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
