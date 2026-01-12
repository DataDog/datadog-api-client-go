// Create a Data Quality monitor returns "OK" response

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
		Type:    datadogV1.MONITORTYPE_DATA_QUALITY_ALERT,
		Query:   `formula("query1").last("5m") > 100`,
		Message: datadog.PtrString("Data quality alert triggered"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(100),
			},
			Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionDataQualityQueryDefinition: &datadogV1.MonitorFormulaAndFunctionDataQualityQueryDefinition{
						Name:       "query1",
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONDATAQUALITYDATASOURCE_DATA_QUALITY_METRICS,
						Measure:    "row_count",
						Filter:     `search for column where ` + "`" + `database:production AND table:users` + "`",
						GroupBy: []string{
							"entity_id",
						},
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
