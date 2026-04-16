// Timeseries cross product query with apm_dependency_stats data source returns "OK" response

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
	body := datadogV2.TimeseriesFormulaQueryRequest{
		Data: datadogV2.TimeseriesFormulaRequest{
			Attributes: datadogV2.TimeseriesFormulaRequestAttributes{
				Formulas: []datadogV2.QueryFormula{
					{
						Formula: "a",
						Limit: &datadogV2.FormulaLimit{
							Count: datadog.PtrInt32(10),
							Order: datadogV2.QUERYSORTORDER_DESC.Ptr(),
						},
					},
				},
				From:     1636625471000,
				Interval: datadog.PtrInt64(5000),
				Queries: []datadogV2.TimeseriesQuery{
					datadogV2.TimeseriesQuery{
						ApmDependencyStatsQuery: &datadogV2.ApmDependencyStatsQuery{
							DataSource:      datadogV2.APMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS,
							Name:            "a",
							Env:             "ci",
							Service:         "cassandra",
							Stat:            datadogV2.APMDEPENDENCYSTATNAME_AVG_DURATION,
							OperationName:   "cassandra.query",
							ResourceName:    "DELETE FROM monitor_history.monitor_state_change_history WHERE org_id = ? AND monitor_id IN ? AND group = ?",
							PrimaryTagName:  datadog.PtrString("datacenter"),
							PrimaryTagValue: datadog.PtrString("edge-eu1.prod.dog"),
						}},
				},
				To: 1636629071000,
			},
			Type: datadogV2.TIMESERIESFORMULAREQUESTTYPE_TIMESERIES_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.QueryTimeseriesData(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryTimeseriesData`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.QueryTimeseriesData`:\n%s\n", responseContent)
}
