// Create a RUM formula and functions monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.Monitor{
		Name:    common.PtrString("Example-Create_a_RUM_formula_and_functions_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_RUM_ALERT,
		Query:   `formula("query2 / query1 * 100").last("15m") >= 0.8`,
		Message: common.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplecreatearumformulaandfunctionsmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *common.NewNullableInt64(common.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			Thresholds: &datadog.MonitorThresholds{
				Critical: common.PtrFloat64(0.8),
			},
			Variables: []datadog.MonitorFormulaAndFunctionQueryDefinition{
				datadog.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadog.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadog.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
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
				datadog.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionEventQueryDefinition: &datadog.MonitorFormulaAndFunctionEventQueryDefinition{
						DataSource: datadog.MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
						Name:       "query1",
						Search: &datadog.MonitorFormulaAndFunctionEventQueryDefinitionSearch{
							Query: "status:error",
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
