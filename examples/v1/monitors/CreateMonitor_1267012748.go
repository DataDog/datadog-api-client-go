// Create a Data Jobs duration monitor returns "OK" response

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
		Type:    datadogV1.MONITORTYPE_DATA_JOBS_ALERT,
		Query:   `formula("job_duration(run_query)").detection(anomaly).by(name,workspace_name,entity_id) > 0`,
		Message: datadog.PtrString("Data jobs duration alert triggered"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Options: &datadogV1.MonitorOptions{
			Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionDataQualityQueryDefinition: &datadogV1.MonitorFormulaAndFunctionDataQualityQueryDefinition{
						Name:          "run_query",
						DataSource:    datadogV1.MONITORFORMULAANDFUNCTIONDATAQUALITYDATASOURCE_DATA_QUALITY_METRICS,
						SchemaVersion: datadog.PtrString("0.0.1"),
						Measure:       "duration",
						Filter:        "search for databricks_job",
						RunFilter:     datadog.PtrString("env:prod"),
						GroupBy: []string{
							"name",
							"workspace_name",
							"entity_id",
						},
						MonitorOptions: &datadogV1.MonitorFormulaAndFunctionDataQualityMonitorOptions{
							DetectionType:   datadogV1.MONITORFORMULAANDFUNCTIONDATAQUALITYDETECTIONTYPE_ANOMALY.Ptr(),
							ConsecutiveRuns: datadog.PtrInt64(2),
						},
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
