// Search logs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.LogsListRequest{
		Index: datadog.PtrString("retention-3,retention-15"),
		Query: datadog.PtrString("service:web* AND @http.status_code:[200 TO 299]"),
		Sort:  datadogV1.LOGSSORT_TIME_ASCENDING.Ptr(),
		Time: datadogV1.LogsListRequestTime{
			From: time.Date(2020, 2, 2, 2, 2, 2, 202000, time.UTC),
			To:   time.Date(2020, 2, 20, 2, 2, 2, 202000, time.UTC),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsApi(apiClient)
	resp, r, err := api.ListLogs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`:\n%s\n", responseContent)
}
