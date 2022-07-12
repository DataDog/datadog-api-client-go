// Delete an archive returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsArchivesApi(apiClient)
	r, err := api.DeleteLogsArchive(ctx, "archive_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.DeleteLogsArchive`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
