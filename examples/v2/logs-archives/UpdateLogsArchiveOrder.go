// Update archive order returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.LogsArchiveOrder{
		Data: &datadog.LogsArchiveOrderDefinition{
			Attributes: datadog.LogsArchiveOrderAttributes{
				ArchiveIds: []string{
					"a2zcMylnM4OCHpYusxIi1g",
					"a2zcMylnM4OCHpYusxIi2g",
					"a2zcMylnM4OCHpYusxIi3g",
				},
			},
			Type: datadog.LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsArchivesApi.UpdateLogsArchiveOrder(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchiveOrder`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.UpdateLogsArchiveOrder`:\n%s\n", responseContent)
}
