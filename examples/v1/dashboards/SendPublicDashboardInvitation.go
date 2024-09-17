// Send shared dashboard invitation email returns "OK" response

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

	body := datadogV1.SharedDashboardInvites{
		Data: datadogV1.SharedDashboardInvitesData{
			SharedDashboardInvitesDataObject: &datadogV1.SharedDashboardInvitesDataObject{
				Attributes: datadogV1.SharedDashboardInvitesDataObjectAttributes{
					Email: datadog.PtrString("exampledashboard@datadoghq.com"),
				},
				Type: datadogV1.DASHBOARDINVITETYPE_PUBLIC_DASHBOARD_INVITATION,
			}},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.SendPublicDashboardInvitation(ctx, SharedDashboardToken, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.SendPublicDashboardInvitation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.SendPublicDashboardInvitation`:\n%s\n", responseContent)
}
