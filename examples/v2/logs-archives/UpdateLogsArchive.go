// Update an archive returns "OK" response

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
	body := datadogV2.LogsArchiveCreateRequest{
		Data: &datadogV2.LogsArchiveCreateRequestDefinition{
			Attributes: &datadogV2.LogsArchiveCreateRequestAttributes{
				Destination: datadogV2.LogsArchiveCreateRequestDestination{
					LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{
						Container: "container-name",
						Integration: datadogV2.LogsArchiveIntegrationAzure{
							ClientId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
							TenantId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
						},
						StorageAccount: "account-name",
						Type:           datadogV2.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE,
					}},
				IncludeTags:                datadog.PtrBool(false),
				Name:                       "Nginx Archive",
				Query:                      "source:nginx",
				RehydrationMaxScanSizeInGb: *datadog.NewNullableInt64(datadog.PtrInt64(100)),
				RehydrationTags: []string{
					"team:intake",
					"team:app",
				},
			},
			Type: "archives",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchivesApi(apiClient)
	resp, r, err := api.UpdateLogsArchive(ctx, "archive_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.UpdateLogsArchive`:\n%s\n", responseContent)
}
