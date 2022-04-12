// Create an archive returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

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
Type: datadog.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE,
}},
IncludeTags: datadog.PtrBool(false),
Name: "Nginx Archive",
Query: "source:nginx",
RehydrationTags: &[]string{
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
	resp, r, err := apiClient.LogsArchivesApi.CreateLogsArchive(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.CreateLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.CreateLogsArchive`:\n%s\n", responseContent)
}