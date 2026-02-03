// Create replay heatmap snapshot returns "Created" response

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
	body := datadogV2.SnapshotCreateRequest{
		Data: datadogV2.SnapshotCreateRequestData{
			Attributes: &datadogV2.SnapshotCreateRequestDataAttributes{
				ApplicationId:              "aaaaaaaa-1111-2222-3333-bbbbbbbbbbbb",
				DeviceType:                 "desktop",
				EventId:                    "11111111-2222-3333-4444-555555555555",
				IsDeviceTypeSelectedByUser: false,
				SnapshotName:               "My Snapshot",
				Start:                      0,
				ViewName:                   "/home",
			},
			Type: datadogV2.SNAPSHOTUPDATEREQUESTDATATYPE_SNAPSHOTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayHeatmapsApi(apiClient)
	resp, r, err := api.CreateReplayHeatmapSnapshot(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayHeatmapsApi.CreateReplayHeatmapSnapshot`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayHeatmapsApi.CreateReplayHeatmapSnapshot`:\n%s\n", responseContent)
}
