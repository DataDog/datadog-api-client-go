// Bulk remove rum replay playlist sessions returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SessionIdArray{
		Data: []datadogV2.SessionIdData{
			{
				Id:   datadog.PtrString("00000000-0000-0000-0000-000000000001"),
				Type: datadogV2.VIEWERSHIPHISTORYSESSIONDATATYPE_RUM_REPLAY_SESSION,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayPlaylistsApi(apiClient)
	r, err := api.BulkRemoveRumReplayPlaylistSessions(ctx, 1234567, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayPlaylistsApi.BulkRemoveRumReplayPlaylistSessions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
