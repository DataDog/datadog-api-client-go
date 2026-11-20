// Delete RUM replay playlist returns "No Content" response

package main

import (
	"context"
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
	r, err := api.DeleteRUMReplayPlaylist(ctx, 1234567)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMReplayPlaylistsApi.DeleteRUMReplayPlaylist`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
