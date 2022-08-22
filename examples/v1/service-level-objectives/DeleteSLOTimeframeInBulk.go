// Bulk Delete SLO Timeframes returns "OK" response

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
	body := map[string][]datadogV1.SLOTimeframe{
		"id1": []datadogV1.SLOTimeframe{
			datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
			datadogV1.SLOTIMEFRAME_THIRTY_DAYS,
		},
		"id2": []datadogV1.SLOTimeframe{
			datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
			datadogV1.SLOTIMEFRAME_THIRTY_DAYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.DeleteSLOTimeframeInBulk(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`:\n%s\n", responseContent)
}
