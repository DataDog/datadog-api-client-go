// Get all invitations for a shared dashboard returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.GetPublicDashboardInvitations(ctx, SharedDashboardToken, *datadogV1.NewGetPublicDashboardInvitationsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetPublicDashboardInvitations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetPublicDashboardInvitations`:\n%s\n", responseContent)
}
