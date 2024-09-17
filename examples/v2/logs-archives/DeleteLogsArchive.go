// Delete an archive returns "OK" response

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
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchivesApi(apiClient)
	r, err := api.DeleteLogsArchive(ctx, "archive_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.DeleteLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
