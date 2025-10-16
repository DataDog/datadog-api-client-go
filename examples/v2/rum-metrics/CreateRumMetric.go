// Create a rum-based metric returns "Created" response

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
	body := datadogV2.RumMetricCreateRequest{
Data: datadogV2.RumMetricCreateData{
Attributes: datadogV2.RumMetricCreateAttributes{
Compute: datadogV2.RumMetricCompute{
AggregationType: datadogV2.RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
IncludePercentiles: datadog.PtrBool(true),
Path: datadog.PtrString("@duration"),
},
EventType: datadogV2.RUMMETRICEVENTTYPE_SESSION,
Filter: &datadogV2.RumMetricFilter{
Query: "@service:web-ui",
},
GroupBy: []datadogV2.RumMetricGroupBy{
{
Path: "@browser.name",
TagName: datadog.PtrString("browser_name"),
},
},
Uniqueness: &datadogV2.RumMetricUniqueness{
When: datadogV2.RUMMETRICUNIQUENESSWHEN_WHEN_MATCH,
},
},
Id: "examplerummetric",
Type: datadogV2.RUMMETRICTYPE_RUM_METRICS,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumMetricsApi(apiClient)
	resp, r, err := api.CreateRumMetric(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumMetricsApi.CreateRumMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumMetricsApi.CreateRumMetric`:\n%s\n", responseContent)
}