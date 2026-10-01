// Create an Archive Search returns "OK" response

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
	body := datadogV2.ArchiveSearchCreateRequest{
		Data: datadogV2.ArchiveSearchCreateRequestData{
			Attributes: datadogV2.ArchiveSearchCreateRequestAttributes{
				ArchiveId:   "mhmyYmyLTOaFYKvhNadu1w",
				Description: datadog.PtrString("Investigating the checkout latency spike."),
				From:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "checkout-latency-investigation",
				Query:       "service:checkout status:error",
				Rehydration: &datadogV2.ArchiveSearchCreateRehydration{
					MaxRehydratedEvents: 1000000,
					RetentionDays:       15,
					Tier:                datadogV2.ARCHIVESEARCHREHYDRATIONTIER_STANDARD,
				},
				To: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			Type: datadogV2.ARCHIVESEARCHTYPE_ARCHIVE_SEARCH,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateArchiveSearch", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchiveSearchesApi(apiClient)
	resp, r, err := api.CreateArchiveSearch(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchiveSearchesApi.CreateArchiveSearch`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchiveSearchesApi.CreateArchiveSearch`:\n%s\n", responseContent)
}
