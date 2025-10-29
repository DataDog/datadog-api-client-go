// Initialize rum segments returns "Default segments created successfully" response

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
	configuration.SetUnstableOperationEnabled("v2.InitializeRumSegments", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSegmentsApi(apiClient)
	r, err := api.InitializeRumSegments(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.InitializeRumSegments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
