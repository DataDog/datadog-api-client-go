// Create a monitor with a metrics formula variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.Monitor{
		Name:    datadog.PtrString("Example-Monitor"),
		Type:    datadogV1.MONITORTYPE_QUERY_ALERT,
		Query:   `formula("query1 / query2").rollup("avg").last("5m") > threshold`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				CriticalQuery: datadog.PtrString(`formula("query1 / query2").rollup("p99").last("1w").multiplier(1.5)`),
			},
			Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionMetricQueryDefinition: &datadogV1.MonitorFormulaAndFunctionMetricQueryDefinition{
						Name:       "query1",
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONMETRICQUERYDATASOURCE_METRICS,
						Query:      "sum:trace.http.request.errors{*}.as_count()",
					}},
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionMetricQueryDefinition: &datadogV1.MonitorFormulaAndFunctionMetricQueryDefinition{
						Name:       "query2",
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONMETRICQUERYDATASOURCE_METRICS,
						Query:      "sum:trace.http.request.hits{*}.as_count()",
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
