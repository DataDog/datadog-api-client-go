// Create a ci-pipelines formula and functions monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Monitor{
		Name:    datadog.PtrString("Example-Create_a_ci_pipelines_formula_and_functions_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_CI_PIPELINES_ALERT,
		Query:   `formula("query1 / query2 * 100").last("15m") >= 0.8`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplecreateacipipelinesformulaandfunctionsmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(0.8),
			},
			Variables: []datadog.MonitorFormulaAndFunctionQueryDefinition{
				datadog.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadog.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadog.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
						Name:       "query1",
						Search: &datadog.MonitorFormulaAndFunctionEventQueryDefinitionSearch{
							Query: "@ci.status:error",
						},
						Indexes: []string{
							"*",
						},
						Compute: datadog.MonitorFormulaAndFunctionEventQueryDefinitionCompute{
							Aggregation: datadog.MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
						},
						GroupBy: []datadog.MonitorFormulaAndFunctionEventQueryGroupBy{},
					}},
				datadog.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadog.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadog.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
						Name:       "query2",
						Search: &datadog.MonitorFormulaAndFunctionEventQueryDefinitionSearch{
							Query: "",
						},
						Indexes: []string{
							"*",
						},
						Compute: datadog.MonitorFormulaAndFunctionEventQueryDefinitionCompute{
							Aggregation: datadog.MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
						},
						GroupBy: []datadog.MonitorFormulaAndFunctionEventQueryGroupBy{},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsApi.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
