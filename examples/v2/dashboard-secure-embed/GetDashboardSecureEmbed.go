// Get a secure embed for a dashboard returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetDashboardSecureEmbed", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardSecureEmbedApi(apiClient)
	resp, r, err := api.GetDashboardSecureEmbed(ctx, "dashboard_id", "token")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardSecureEmbedApi.GetDashboardSecureEmbed`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardSecureEmbedApi.GetDashboardSecureEmbed`:\n%s\n", responseContent)
}
