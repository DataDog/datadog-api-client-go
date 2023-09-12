// Get a downtime returns "OK" response

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
	// there is a valid "downtime_v2" in the system
	DowntimeV2DataID := os.Getenv("DOWNTIME_V2_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDowntimesApi(apiClient)
	resp, r, err := api.GetDowntime(ctx, DowntimeV2DataID, *datadogV2.NewGetDowntimeOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.GetDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.GetDowntime`:\n%s\n", responseContent)
}
