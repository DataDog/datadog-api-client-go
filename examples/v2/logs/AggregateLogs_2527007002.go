// Aggregate compute events returns "OK" response

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
	body := datadogV2.LogsAggregateRequest{
Compute: []datadogV2.LogsCompute{
{
Aggregation: datadogV2.LOGSAGGREGATIONFUNCTION_COUNT,
Interval: datadog.PtrString("5m"),
Type: datadogV2.LOGSCOMPUTETYPE_TIMESERIES.Ptr(),
},
},
Filter: &datadogV2.LogsQueryFilter{
From: datadog.PtrString("now-15m"),
Indexes: []string{
"main",
},
Query: datadog.PtrString("*"),
To: datadog.PtrString("now"),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)
	resp, r, err := api.AggregateLogs(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.AggregateLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.AggregateLogs`:\n%s\n", responseContent)
}