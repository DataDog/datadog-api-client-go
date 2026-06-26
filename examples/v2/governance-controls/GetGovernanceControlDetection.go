// Get a governance control detection returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetGovernanceControlDetection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceControlsApi(apiClient)
	resp, r, err := api.GetGovernanceControlDetection(ctx, "detection_type", "detection_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceControlsApi.GetGovernanceControlDetection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceControlsApi.GetGovernanceControlDetection`:\n%s\n", responseContent)
}
