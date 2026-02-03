// Update replay heatmap snapshot returns "OK" response

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
	body := datadogV2.SnapshotUpdateRequest{
		Data: datadogV2.SnapshotUpdateRequestData{
			Attributes: &datadogV2.SnapshotUpdateRequestDataAttributes{
				EventId:                    "11111111-2222-3333-4444-555555555555",
				IsDeviceTypeSelectedByUser: false,
				Start:                      0,
			},
			Id:   datadog.PtrString("00000000-0000-0000-0000-000000000001"),
			Type: datadogV2.SNAPSHOTUPDATEREQUESTDATATYPE_SNAPSHOTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayHeatmapsApi(apiClient)
	resp, r, err := api.UpdateReplayHeatmapSnapshot(ctx, "00000000-0000-0000-0000-000000000001", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayHeatmapsApi.UpdateReplayHeatmapSnapshot`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayHeatmapsApi.UpdateReplayHeatmapSnapshot`:\n%s\n", responseContent)
}
