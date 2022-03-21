// Bulk Delete SLO Timeframes returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := map[string][]datadog.SLOTimeframe{
		"id1": []datadog.SLOTimeframe{
			datadog.SLOTIMEFRAME_SEVEN_DAYS,
			datadog.SLOTIMEFRAME_THIRTY_DAYS,
		},
		"id2": []datadog.SLOTimeframe{
			datadog.SLOTIMEFRAME_SEVEN_DAYS,
			datadog.SLOTIMEFRAME_THIRTY_DAYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`:\n%s\n", responseContent)
}
