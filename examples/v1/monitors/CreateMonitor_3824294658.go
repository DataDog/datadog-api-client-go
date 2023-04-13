// Create a ci-pipelines formula and functions monitor returns "OK" response

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
		Type:    datadogV1.MONITORTYPE_CI_PIPELINES_ALERT,
		Query:   `formula("query1 / query2 * 100").last("15m") >= 0.8`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(0.8),
			},
			Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadogV1.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
						Name:       "query1",
						Search: &datadogV1.MonitorFormulaAndFunctionEventQueryDefinitionSearch{
							Query: "@ci.status:error",
						},
						Indexes: []string{
							"*",
						},
						Compute: datadogV1.MonitorFormulaAndFunctionEventQueryDefinitionCompute{
							Aggregation: datadogV1.MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
						},
						GroupBy: []datadogV1.MonitorFormulaAndFunctionEventQueryGroupBy{},
					}},
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadogV1.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
						Name:       "query2",
						Search: &datadogV1.MonitorFormulaAndFunctionEventQueryDefinitionSearch{
							Query: "",
						},
						Indexes: []string{
							"*",
						},
						Compute: datadogV1.MonitorFormulaAndFunctionEventQueryDefinitionCompute{
							Aggregation: datadogV1.MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
						},
						GroupBy: []datadogV1.MonitorFormulaAndFunctionEventQueryGroupBy{},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
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
