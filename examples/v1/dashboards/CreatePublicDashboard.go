// Create a shared dashboard returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	body := datadogV1.SharedDashboard{
		DashboardId:   DashboardID,
		DashboardType: datadogV1.DASHBOARDTYPE_CUSTOM_TIMEBOARD,
		ShareType:     *datadogV1.NewNullableDashboardShareType(datadogV1.DASHBOARDSHARETYPE_OPEN.Ptr()),
		GlobalTime: &datadogV1.DashboardGlobalTime{
			LiveSpan: datadogV1.DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_HOUR.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreatePublicDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreatePublicDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreatePublicDashboard`:\n%s\n", responseContent)
}
