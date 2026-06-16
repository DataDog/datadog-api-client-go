// List network health insights returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListNetworkHealthInsights", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewNetworkHealthInsightsApi(apiClient)
	resp, r, err := api.ListNetworkHealthInsights(ctx, *datadogV2.NewListNetworkHealthInsightsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkHealthInsightsApi.ListNetworkHealthInsights`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NetworkHealthInsightsApi.ListNetworkHealthInsights`:\n%s\n", responseContent)
}
