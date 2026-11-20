// List RUM replay session watchers returns "OK" response

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
	api := datadogV2.NewRUMReplayViewershipApi(apiClient)
	resp, r, err := api.ListRUMReplaySessionWatchers(ctx, "00000000-0000-0000-0000-000000000001", *datadogV2.NewListRUMReplaySessionWatchersOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMReplayViewershipApi.ListRUMReplaySessionWatchers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMReplayViewershipApi.ListRUMReplaySessionWatchers`:\n%s\n", responseContent)
}
