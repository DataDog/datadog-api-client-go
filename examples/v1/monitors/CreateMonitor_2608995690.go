// Create a monitor with aggregate augmented query variables returns "OK" response

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
		Query:   `formula("query1").rollup("sum").last("5m") > 124`,
		Message: datadog.PtrString("test message"),
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(124),
			},
			Variables: []datadogV1.MonitorFormulaAndFunctionQueryDefinition{
				datadogV1.MonitorFormulaAndFunctionQueryDefinition{
					MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition: &datadogV1.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition{
						DataSource: datadogV1.MONITORFORMULAANDFUNCTIONAGGREGATEAUGMENTEDDATASOURCE_AGGREGATE_AUGMENTED_QUERY,
						Name:       datadog.PtrString("query1"),
						GroupBy: []datadogV1.MonitorFormulaAndFunctionEventQueryGroupBy{
							{
								Facet: "org_id",
							},
							{
								Facet: "name",
							},
						},
						Compute: []datadogV1.MonitorFormulaAndFunctionEventQueryDefinitionCompute{
							{
								Name:        datadog.PtrString("compute_result"),
								Aggregation: datadogV1.MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MAX,
							},
						},
						AugmentQuery: datadogV1.MonitorFormulaAndFunctionAggregateAugmentQuery{
							MonitorFormulaAndFunctionReferenceTableQueryDefinition: &datadogV1.MonitorFormulaAndFunctionReferenceTableQueryDefinition{
								Name:       datadog.PtrString("filter_query"),
								DataSource: datadogV1.MONITORFORMULAANDFUNCTIONREFERENCETABLEDATASOURCE_REFERENCE_TABLE,
								TableName:  "test_table",
								Columns: []datadogV1.MonitorFormulaAndFunctionReferenceTableColumn{
									{
										Name: "org_id",
									},
									{
										Name: "name",
									},
								},
							}},
						BaseQuery: datadogV1.MonitorFormulaAndFunctionAggregateBaseQuery{
							MonitorFormulaAndFunctionMetricsQueryDefinition: &datadogV1.MonitorFormulaAndFunctionMetricsQueryDefinition{
								DataSource: datadogV1.MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_METRICS,
								Name:       datadog.PtrString("query1"),
								Query:      "avg:dd{*} by {org_id}.as_count()",
							}},
						JoinCondition: datadogV1.MonitorFormulaAndFunctionAggregateQueryJoinCondition{
							AugmentAttribute: "org_id",
							BaseAttribute:    "org_id",
							JoinType:         datadogV1.MONITORFORMULAANDFUNCTIONAGGREGATEQUERYJOINTYPE_INNER,
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
