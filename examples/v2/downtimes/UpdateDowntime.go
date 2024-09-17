// Update a downtime returns "OK" response

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

	body := datadogV2.DowntimeUpdateRequest{
		Data: datadogV2.DowntimeUpdateRequestData{
			Attributes: datadogV2.DowntimeUpdateRequestAttributes{
				Message: *datadog.NewNullableString(datadog.PtrString("light speed")),
			},
			Id:   DowntimeV2DataID,
			Type: datadogV2.DOWNTIMERESOURCETYPE_DOWNTIME,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDowntimesApi(apiClient)
	resp, r, err := api.UpdateDowntime(ctx, DowntimeV2DataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.UpdateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.UpdateDowntime`:\n%s\n", responseContent)
}
