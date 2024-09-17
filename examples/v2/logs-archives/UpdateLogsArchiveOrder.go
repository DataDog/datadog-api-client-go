// Update archive order returns "OK" response

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
	body := datadogV2.LogsArchiveOrder{
		Data: &datadogV2.LogsArchiveOrderDefinition{
			Attributes: datadogV2.LogsArchiveOrderAttributes{
				ArchiveIds: []string{
					"a2zcMylnM4OCHpYusxIi1g",
					"a2zcMylnM4OCHpYusxIi2g",
					"a2zcMylnM4OCHpYusxIi3g",
				},
			},
			Type: datadogV2.LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsArchivesApi(apiClient)
	resp, r, err := api.UpdateLogsArchiveOrder(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchiveOrder`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.UpdateLogsArchiveOrder`:\n%s\n", responseContent)
}
