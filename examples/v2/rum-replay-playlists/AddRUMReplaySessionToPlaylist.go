// Add RUM replay session to playlist returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMReplayPlaylistsApi(apiClient)
	resp, r, err := api.AddRUMReplaySessionToPlaylist(ctx, 1704067200000, 1234567, "00000000-0000-0000-0000-000000000001", *datadogV2.NewAddRUMReplaySessionToPlaylistOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMReplayPlaylistsApi.AddRUMReplaySessionToPlaylist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMReplayPlaylistsApi.AddRUMReplaySessionToPlaylist`:\n%s\n", responseContent)
}
