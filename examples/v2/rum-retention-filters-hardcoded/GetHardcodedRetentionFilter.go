// Get a hardcoded retention filter returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMRetentionFiltersHardcodedApi(apiClient)
	resp, r, err := api.GetHardcodedRetentionFilter(ctx, "Example-RUM-Retention-Filters-Hardcoded", "Example-RUM-Retention-Filters-Hardcoded")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMRetentionFiltersHardcodedApi.GetHardcodedRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMRetentionFiltersHardcodedApi.GetHardcodedRetentionFilter`:\n%s\n", responseContent)
}
