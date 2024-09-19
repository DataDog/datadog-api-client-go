// Delete dashboards returns "No Content" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")


	body := datadogV1.DashboardBulkDeleteRequest{
Data: []datadogV1.DashboardBulkActionData{
{
Id: DashboardID,
Type: datadogV1.DASHBOARDRESOURCETYPE_DASHBOARD,
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	r, err := api.DeleteDashboards(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboards`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}