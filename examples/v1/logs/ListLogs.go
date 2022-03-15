// Search logs returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.LogsListRequest{
		Index: datadog.PtrString("main"),
		Limit: datadog.PtrInt32(5),
		Query: datadog.PtrString("service:web* AND @http.status_code:[200 TO 299]"),
		Sort:  datadog.LOGSSORT_TIME_ASCENDING.Ptr(),
		Time: datadog.LogsListRequestTime{
			From:     time.Now().AddDate(0, 0, -1).Unix(),
			Timezone: datadog.PtrString("Europe/Paris"),
			To:       time.Now().Unix(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsApi.ListLogs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`:\n%s\n", responseContent)
}
