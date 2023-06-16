// Update a shared dashboard returns "OK" response

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
	// there is a valid "shared_dashboard" in the system
	SharedDashboardToken := os.Getenv("SHARED_DASHBOARD_TOKEN")

	body := datadogV1.SharedDashboardUpdateRequest{
		GlobalTime: *datadogV1.NewNullableSharedDashboardUpdateRequestGlobalTime(&datadogV1.SharedDashboardUpdateRequestGlobalTime{
			LiveSpan: datadogV1.DASHBOARDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES.Ptr(),
		}),
		ShareList: *datadog.NewNullableList(&[]string{}),
		ShareType: *datadogV1.NewNullableDashboardShareType(datadogV1.DASHBOARDSHARETYPE_OPEN.Ptr()),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.UpdatePublicDashboard(ctx, SharedDashboardToken, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdatePublicDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.UpdatePublicDashboard`:\n%s\n", responseContent)
}
