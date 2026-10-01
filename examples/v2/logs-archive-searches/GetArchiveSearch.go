// Get an Archive Search returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetArchiveSearch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchiveSearchesApi(apiClient)
	resp, r, err := api.GetArchiveSearch(ctx, "archive_search_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchiveSearchesApi.GetArchiveSearch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchiveSearchesApi.GetArchiveSearch`:\n%s\n", responseContent)
}
