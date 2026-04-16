// Scalar cross product query with container data source returns "OK" response

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
	body := datadogV2.ScalarFormulaQueryRequest{
		Data: datadogV2.ScalarFormulaRequest{
			Attributes: datadogV2.ScalarFormulaRequestAttributes{
				Formulas: []datadogV2.QueryFormula{
					{
						Formula: "a",
						Limit: &datadogV2.FormulaLimit{
							Count: datadog.PtrInt32(10),
							Order: datadogV2.QUERYSORTORDER_DESC.Ptr(),
						},
					},
				},
				From: 1636625471000,
				Queries: []datadogV2.ScalarQuery{
					datadogV2.ScalarQuery{
						ContainerScalarQuery: &datadogV2.ContainerScalarQuery{
							DataSource: datadogV2.CONTAINERDATASOURCE_CONTAINER,
							Name:       "a",
							Metric:     "process.stat.container.cpu.system_pct",
							Aggregator: datadogV2.METRICSAGGREGATOR_AVG.Ptr(),
							TagFilters: []string{},
							Limit:      datadog.PtrInt64(10),
							Sort:       datadogV2.QUERYSORTORDER_DESC.Ptr(),
						}},
				},
				To: 1636629071000,
			},
			Type: datadogV2.SCALARFORMULAREQUESTTYPE_SCALAR_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.QueryScalarData(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryScalarData`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.QueryScalarData`:\n%s\n", responseContent)
}
