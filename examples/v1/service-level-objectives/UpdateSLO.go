// Update an SLO returns "OK" response

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
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")
	SloData0Name := os.Getenv("SLO_DATA_0_NAME")

	body := datadogV1.ServiceLevelObjective{
		Type: datadogV1.SLOTYPE_METRIC,
		Name: SloData0Name,
		Thresholds: []datadogV1.SLOThreshold{
			{
				Target:    97.0,
				Timeframe: datadogV1.SLOTIMEFRAME_SEVEN_DAYS,
				Warning:   datadog.PtrFloat64(98.0),
			},
		},
		Timeframe:        datadogV1.SLOTIMEFRAME_SEVEN_DAYS.Ptr(),
		TargetThreshold:  datadog.PtrFloat64(97.0),
		WarningThreshold: datadog.PtrFloat64(98),
		Query: &datadogV1.ServiceLevelObjectiveQuery{
			Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
			Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.UpdateSLO(ctx, SloData0ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.UpdateSLO`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.UpdateSLO`:\n%s\n", responseContent)
}
