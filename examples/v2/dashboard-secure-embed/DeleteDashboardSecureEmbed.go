// Delete a secure embed for a dashboard returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteDashboardSecureEmbed", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardSecureEmbedApi(apiClient)
	r, err := api.DeleteDashboardSecureEmbed(ctx, "dashboard_id", "token")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSecureEmbedApi.DeleteDashboardSecureEmbed`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
