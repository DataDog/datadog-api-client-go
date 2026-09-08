// Investigate a timeseries anomaly returns "OK" response

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
	body := datadogV2.TimeseriesAnomalyInvestigationRequest{
		Data: datadogV2.TimeseriesAnomalyInvestigationRequestData{
			Attributes: datadogV2.TimeseriesAnomalyInvestigationRequestAttributes{
				Requests: []datadogV2.TimeseriesAnomalyInvestigationTimeseriesRequest{
					{
						Formulas: []datadogV2.TimeseriesAnomalyInvestigationFormula{
							{
								Formula: "anomalies(query1, 'agile', 3)",
							},
						},
						From: 1754406000000,
						Queries: []datadogV2.TimeseriesAnomalyInvestigationMetricQuery{
							{
								DataSource: datadogV2.TIMESERIESANOMALYINVESTIGATIONDATASOURCE_METRICS,
								Name:       "query1",
								Query:      "avg:system.cpu.user{env:prod} by {service}",
							},
						},
						To: 1754423940000,
					},
				},
			},
			Type: datadogV2.TIMESERIESANOMALYINVESTIGATIONTYPE_TIMESERIES_ANOMALY_INVESTIGATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTimeseriesAnomalyInvestigation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTimeseriesAnomalyInvestigationsApi(apiClient)
	resp, r, err := api.CreateTimeseriesAnomalyInvestigation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesAnomalyInvestigationsApi.CreateTimeseriesAnomalyInvestigation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TimeseriesAnomalyInvestigationsApi.CreateTimeseriesAnomalyInvestigation`:\n%s\n", responseContent)
}
