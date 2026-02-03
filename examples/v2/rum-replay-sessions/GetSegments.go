// Get segments returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplaySessionsApi(apiClient)
	r, err := api.GetSegments(ctx, "00000000-0000-0000-0000-000000000002", "00000000-0000-0000-0000-000000000001", *datadogV2.NewGetSegmentsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplaySessionsApi.GetSegments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
