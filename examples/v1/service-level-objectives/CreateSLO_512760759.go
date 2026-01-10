// Create a new metric SLO object using sli_specification returns "OK" response

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
	body := datadogV1.ServiceLevelObjectiveRequest{
		Type:        datadogV1.SLOTYPE_METRIC,
		Description: *datadog.NewNullableString(datadog.PtrString("Metric SLO using sli_specification")),
		Name:        "Example-Service-Level-Objective",
		SliSpecification: &datadogV1.SLOSliSpec{
			SLOCountSpec: &datadogV1.SLOCountSpec{
				Count: datadogV1.SLOCountDefinition{
					GoodEventsFormula: datadogV1.SLOFormula{
						Formula: "query1 - query2",
					},
					TotalEventsFormula: datadogV1.SLOFormula{
						Formula: "query1",
					},
					Queries: []datadogV1.SLODataSourceQueryDefinition{
						datadogV1.SLODataSourceQueryDefinition{
							FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
								DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
								Name:       "query1",
								Query:      "sum:httpservice.hits{*}.as_count()",
							}},
						datadogV1.SLODataSourceQueryDefinition{
							FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
								DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
								Name:       "query2",
								Query:      "sum:httpservice.errors{*}.as_count()",
							}},
					},
				},
			}},
		Tags: []string{
			"env:prod",
			"type:count",
		},
		Thresholds: []datadogV1.SLOThreshold{
			{
				Target:         99.0,
				TargetDisplay:  datadog.PtrString("99.0"),
				Timeframe:      datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
				Warning:        datadog.PtrFloat64(99.5),
				WarningDisplay: datadog.PtrString("99.5"),
			},
		},
		Timeframe:        datadogV1.SLOTIMEFRAME_SEVEN_DAYS.Ptr(),
		TargetThreshold:  datadog.PtrFloat64(99.0),
		WarningThreshold: datadog.PtrFloat64(99.5),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.CreateSLO(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSLO`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.CreateSLO`:\n%s\n", responseContent)
}
