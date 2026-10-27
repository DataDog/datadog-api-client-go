// Rehydrate an Archive Search returns "OK" response

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
	body := datadogV2.ArchiveSearchRehydrateRequest{
		Data: datadogV2.ArchiveSearchRehydrateRequestData{
			Attributes: datadogV2.ArchiveSearchRehydrateRequestAttributes{
				Rehydration: datadogV2.ArchiveSearchRehydrateRequestRehydration{
					RetentionDays: 15,
					Tier:          datadogV2.ARCHIVESEARCHREHYDRATIONTIER_STANDARD,
				},
			},
			Type: datadogV2.ARCHIVESEARCHTYPE_ARCHIVE_SEARCH,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RehydrateArchiveSearch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchiveSearchesApi(apiClient)
	resp, r, err := api.RehydrateArchiveSearch(ctx, "archive_search_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchiveSearchesApi.RehydrateArchiveSearch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchiveSearchesApi.RehydrateArchiveSearch`:\n%s\n", responseContent)
}
