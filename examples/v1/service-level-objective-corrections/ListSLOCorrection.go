// Get all SLO corrections returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectiveCorrectionsApi(apiClient)
	resp, r, err := api.ListSLOCorrection(ctx, *datadogV1.NewListSLOCorrectionOptionalParameters().WithOffset(1).WithLimit(1))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection`:\n%s\n", responseContent)
}
