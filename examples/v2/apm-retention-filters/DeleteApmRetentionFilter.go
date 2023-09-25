// Delete a retention filter returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "retention_filter" in the system
	RetentionFilterDataID := os.Getenv("RETENTION_FILTER_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMRetentionFiltersApi(apiClient)
	r, err := api.DeleteApmRetentionFilter(ctx, RetentionFilterDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMRetentionFiltersApi.DeleteApmRetentionFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
