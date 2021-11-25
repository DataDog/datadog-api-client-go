// Get an SLO correction for an SLO returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "correction" for "slo"
	CORRECTION_DATA_ID := os.Getenv("CORRECTION_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetSLOCorrection", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection(ctx, CORRECTION_DATA_ID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection`:\n%s\n", responseContent)
}
