// Get a permanent RUM retention filter returns "OK" response

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
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionFiltersApi(apiClient)
	resp, r, err := api.GetPermanentRetentionFilter(ctx, "app_id", datadogV2.RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionFiltersApi.GetPermanentRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRetentionFiltersApi.GetPermanentRetentionFilter`:\n%s\n", responseContent)
}
