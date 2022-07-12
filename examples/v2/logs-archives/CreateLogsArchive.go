// Create an archive returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.LogsArchiveCreateRequest{
		Data: &datadog.LogsArchiveCreateRequestDefinition{
			Attributes: &datadog.LogsArchiveCreateRequestAttributes{
				Destination: datadog.LogsArchiveCreateRequestDestination{
					LogsArchiveDestinationAzure: &datadog.LogsArchiveDestinationAzure{
						Container: "container-name",
						Integration: datadog.LogsArchiveIntegrationAzure{
							ClientId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
							TenantId: "aaaaaaaa-1a1a-1a1a-1a1a-aaaaaaaaaaaa",
						},
						StorageAccount: "account-name",
						Type:           datadog.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE,
					}},
				IncludeTags:                common.PtrBool(false),
				Name:                       "Nginx Archive",
				Query:                      "source:nginx",
				RehydrationMaxScanSizeInGb: *common.NewNullableInt64(common.PtrInt64(100)),
				RehydrationTags: []string{
					"team:intake",
					"team:app",
				},
			},
			Type: "archives",
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsArchivesApi(apiClient)
	resp, r, err := api.CreateLogsArchive(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.CreateLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.CreateLogsArchive`:\n%s\n", responseContent)
}
