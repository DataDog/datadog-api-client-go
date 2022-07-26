// Search logs returns "OK" response with pagination

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
	body := datadog.LogsListRequest{
		Filter: &datadog.LogsQueryFilter{
			From: common.PtrString("now-15m"),
			Indexes: []string{
				"main",
			},
			To: common.PtrString("now"),
		},
		Options: &datadog.LogsQueryOptions{
			Timezone: common.PtrString("GMT"),
		},
		Page: &datadog.LogsListRequestPage{
			Limit: common.PtrInt32(2),
		},
		Sort: datadog.LOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsApi(apiClient)
	resp, _, err := api.ListLogsWithPagination(ctx, *datadog.NewListLogsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
	}

	for item := range resp {
		responseContent, _ := json.MarshalIndent(item, "", "  ")
		fmt.Fprintf(os.Stdout, "%s\n", responseContent)
	}
}
