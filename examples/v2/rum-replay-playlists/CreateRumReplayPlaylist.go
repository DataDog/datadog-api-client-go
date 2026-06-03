// Create rum replay playlist returns "Created" response

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
	body := datadogV2.Playlist{
		Data: datadogV2.PlaylistData{
			Attributes: &datadogV2.PlaylistDataAttributes{
				CreatedBy: &datadogV2.PlaylistDataAttributesCreatedBy{
					Handle: "john.doe@example.com",
					Id:     "00000000-0000-0000-0000-000000000001",
					Uuid:   "00000000-0000-0000-0000-000000000001",
				},
				Name: "My Playlist",
			},
			Type: datadogV2.PLAYLISTDATATYPE_RUM_REPLAY_PLAYLIST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayPlaylistsApi(apiClient)
	resp, r, err := api.CreateRumReplayPlaylist(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayPlaylistsApi.CreateRumReplayPlaylist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayPlaylistsApi.CreateRumReplayPlaylist`:\n%s\n", responseContent)
}
