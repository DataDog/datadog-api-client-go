// Create a Product Analytics monitor returns "Created" response

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
	body := datadogV2.ProductAnalyticsMonitorCreateRequest{
		Data: datadogV2.ProductAnalyticsMonitorCreateData{
			Attributes: datadogV2.ProductAnalyticsMonitorCreateAttributes{
				Message: `Conversion from /a to /b dropped over the last hour

@slack-my-alerts-channel`,
				Name: "Funnel conversion drop",
				Options: &datadogV2.ProductAnalyticsMonitorOptions{
					Thresholds: &datadogV2.ProductAnalyticsMonitorThresholds{
						Critical: datadog.PtrFloat64(1),
						Warning:  datadog.PtrFloat64(2),
					},
					Variables: []map[string]interface{}{
						map[string]interface{}{
							"data_source": "product_analytics_journey",
							"name":        "query",
							"search":      "{'expression': 'node_0->node_1', 'filters': {'graph_filters': [{'name': '__dd.time_to_convert', 'operator': '<=', 'value': 86400000}]}, 'join_keys': {'primary': '@usr.id'}, 'node_objects': {'node_0': {'data_source': 'product_analytics', 'search': {'query': '@type:view @view.name:/a'}}, 'node_1': {'data_source': 'product_analytics', 'search': {'query': '@type:view @view.name:/b'}}}}",
						},
					},
				},
				Priority: datadog.PtrInt32(3),
				Query:    `formula("query").rollup("sum").last("1h") < 1`,
				Tags: map[string]string{
					"chart_id":   "chart-1",
					"query_type": "funnel",
				},
			},
			Type: datadogV2.PRODUCTANALYTICSMONITORREQUESTTYPE_PA_MONITOR_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateProductAnalyticsMonitor", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.CreateProductAnalyticsMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.CreateProductAnalyticsMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.CreateProductAnalyticsMonitor`:\n%s\n", responseContent)
}
