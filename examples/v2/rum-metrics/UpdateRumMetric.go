// Update a rum-based metric returns "OK" response

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
	// there is a valid "rum_metric" in the system
	RumMetricDataID := os.Getenv("RUM_METRIC_DATA_ID")


	body := datadogV2.RumMetricUpdateRequest{
Data: datadogV2.RumMetricUpdateData{
Id: datadog.PtrString(RumMetricDataID),
Type: datadogV2.RUMMETRICTYPE_RUM_METRICS,
Attributes: datadogV2.RumMetricUpdateAttributes{
Compute: &datadogV2.RumMetricUpdateCompute{
IncludePercentiles: datadog.PtrBool(false),
},
Filter: &datadogV2.RumMetricFilter{
Query: "@service:rum-config",
},
GroupBy: []datadogV2.RumMetricGroupBy{
{
Path: "@browser.version",
TagName: datadog.PtrString("browser_version"),
},
},
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumMetricsApi(apiClient)
	resp, r, err := api.UpdateRumMetric(ctx, RumMetricDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumMetricsApi.UpdateRumMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumMetricsApi.UpdateRumMetric`:\n%s\n", responseContent)
}