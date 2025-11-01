// List rum segments returns "Successful response with list of segments" response

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
	configuration.SetUnstableOperationEnabled("v2.ListRumSegments", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSegmentsApi(apiClient)
	resp, r, err := api.ListRumSegments(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.ListRumSegments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.ListRumSegments`:\n%s\n", responseContent)
}
