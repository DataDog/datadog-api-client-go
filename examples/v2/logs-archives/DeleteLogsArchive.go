// Delete an archive returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.LogsArchivesApi.DeleteLogsArchive(ctx, "archive_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.DeleteLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
