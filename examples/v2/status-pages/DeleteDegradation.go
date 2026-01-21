// Delete degradation returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "status_page" in the system
	StatusPageDataID := uuid.MustParse(os.Getenv("STATUS_PAGE_DATA_ID"))

	// there is a valid "degradation" in the system
	DegradationDataID := uuid.MustParse(os.Getenv("DEGRADATION_DATA_ID"))

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
