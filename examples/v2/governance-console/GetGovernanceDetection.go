// Get a detection returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetGovernanceDetection", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewGovernanceConsoleApi(apiClient)
	resp, r, err := api.GetGovernanceDetection(ctx, "detection_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceConsoleApi.GetGovernanceDetection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GovernanceConsoleApi.GetGovernanceDetection`:\n%s\n", responseContent)
}
