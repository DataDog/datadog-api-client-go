// Search logs (POST) returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.LogsListRequest{
Filter: &datadogV2.LogsQueryFilter{
From: datadog.PtrString("now-15m"),
Indexes: []string{
"main",
"web",
},
Query: datadog.PtrString("service:web* AND @http.status_code:[200 TO 299]"),
StorageTier: datadogV2.LOGSSTORAGETIER_INDEXES.Ptr(),
To: datadog.PtrString("now"),
},
Options: &datadogV2.LogsQueryOptions{
Timezone: datadog.PtrString("GMT"),
},
Page: &datadogV2.LogsListRequestPage{
Cursor: datadog.PtrString("eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ=="),
Limit: datadog.PtrInt32(25),
},
Sort: datadogV2.LOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)
	resp, r, err := api.ListLogs(ctx, *datadogV2.NewListLogsOptionalParameters().WithBody(body), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`:\n%s\n", responseContent)
}