// Get a RUM segment returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetRumSegment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMUserSegmentsApi(apiClient)
	resp, r, err := api.GetRumSegment(ctx, "a1b2c3d4-1234-5678-9abc-123456789abc")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMUserSegmentsApi.GetRumSegment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMUserSegmentsApi.GetRumSegment`:\n%s\n", responseContent)
}
