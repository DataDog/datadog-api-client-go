// List control detections returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListGovernanceControlDetections", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceConsoleApi(apiClient)
	resp, r, err := api.ListGovernanceControlDetections(ctx, "detection_type", *datadogV2.NewListGovernanceControlDetectionsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceConsoleApi.ListGovernanceControlDetections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceConsoleApi.ListGovernanceControlDetections`:\n%s\n", responseContent)
}
