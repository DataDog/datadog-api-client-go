// Create a time-slice SLO object returns "OK" response

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
		Type:        datadogV1.SLOTYPE_TIME_SLICE,
		Description: *datadog.NewNullableString(datadog.PtrString("string")),
		Name:        "Example-Service-Level-Objective",
		SliSpecification: &datadogV1.SLOSliSpec{
			SLOTimeSliceSpec: &datadogV1.SLOTimeSliceSpec{
				TimeSlice: datadogV1.SLOTimeSliceCondition{
					Query: datadogV1.SLOTimeSliceQuery{
						Formulas: []datadogV1.SLOFormula{
							{
								Formula: "query1",
							},
						},
						Queries: []datadogV1.SLODataSourceQueryDefinition{
							datadogV1.SLODataSourceQueryDefinition{
								FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
									DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
									Name:       "query1",
									Query:      "trace.servlet.request{env:prod}",
								}},
						},
					},
					Comparator: datadogV1.SLOTIMESLICECOMPARATOR_GREATER,
					Threshold:  5,
				},
			}},
		Tags: []string{
			"env:prod",
		},
		Thresholds: []datadogV1.SLOThreshold{
			{
				Target:         97.0,
				TargetDisplay:  datadog.PtrString("97.0"),
				Timeframe:      datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
				Warning:        datadog.PtrFloat64(98),
				WarningDisplay: datadog.PtrString("98.0"),
			},
		},
		Timeframe:        datadogV1.SLOTIMEFRAME_SEVEN_DAYS.Ptr(),
		TargetThreshold:  datadog.PtrFloat64(97.0),
		WarningThreshold: datadog.PtrFloat64(98),
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
