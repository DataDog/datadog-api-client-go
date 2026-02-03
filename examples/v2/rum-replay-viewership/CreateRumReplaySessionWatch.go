// Create rum replay session watch returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.Watch{
		Data: datadogV2.WatchData{
			Attributes: &datadogV2.WatchDataAttributes{
				ApplicationId: "aaaaaaaa-1111-2222-3333-bbbbbbbbbbbb",
				EventId:       "11111111-2222-3333-4444-555555555555",
				Timestamp:     time.Date(2026, 1, 13, 17, 15, 53, 208340, time.UTC),
			},
			Type: datadogV2.WATCHDATATYPE_RUM_REPLAY_WATCH,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayViewershipApi(apiClient)
	resp, r, err := api.CreateRumReplaySessionWatch(ctx, "00000000-0000-0000-0000-000000000001", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayViewershipApi.CreateRumReplaySessionWatch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayViewershipApi.CreateRumReplaySessionWatch`:\n%s\n", responseContent)
}
