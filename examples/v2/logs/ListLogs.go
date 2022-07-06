// Search logs returns "OK" response

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
			Query: datadog.PtrString("datadog-agent"),
			Indexes: []string{
				"main",
			},
			From: datadog.PtrString("2020-09-17T11:48:36+01:00"),
			To:   datadog.PtrString("2020-09-17T12:48:36+01:00"),
		},
		Sort: datadog.LOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
		Page: &datadog.LogsListRequestPage{
			Limit: datadog.PtrInt32(5),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.LogsApi(apiClient)
	resp, r, err := api.ListLogs(ctx, *datadog.NewListLogsOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`:\n%s\n", responseContent)
}
