// Delete degradation returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "status_page" in the system
	StatusPageDataID := os.Getenv("STATUS_PAGE_DATA_ID")

	// there is a valid "degradation" in the system
	DegradationDataID := os.Getenv("DEGRADATION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewStatusPagesApi(apiClient)
	r, err := api.DeleteDegradation(ctx, StatusPageDataID, DegradationDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusPagesApi.DeleteDegradation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
