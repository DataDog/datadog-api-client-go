// List all RUM segments returns "OK" response

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
	api := datadogV2.NewRUMUserSegmentsApi(apiClient)
	resp, r, err := api.ListRumSegments(ctx, *datadogV2.NewListRumSegmentsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMUserSegmentsApi.ListRumSegments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMUserSegmentsApi.ListRumSegments`:\n%s\n", responseContent)
}
