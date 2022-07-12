// Search logs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.LogsListRequest{
		Index: datadog.PtrString("retention-3,retention-15"),
		Query: datadog.PtrString("service:web* AND @http.status_code:[200 TO 299]"),
		Sort:  datadog.LOGSSORT_TIME_ASCENDING.Ptr(),
		Time: datadog.LogsListRequestTime{
			From: time.Date(2020, 2, 2, 2, 2, 2, 202000, time.UTC),
			To:   time.Date(2020, 2, 20, 2, 2, 2, 202000, time.UTC),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsApi(apiClient)
	resp, r, err := api.ListLogs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`:\n%s\n", responseContent)
}
