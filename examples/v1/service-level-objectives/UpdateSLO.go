// Update an SLO returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")
	SloData0Name := os.Getenv("SLO_DATA_0_NAME")

	body := datadog.ServiceLevelObjective{
		Type: datadog.SLOTYPE_METRIC,
		Name: SloData0Name,
		Thresholds: []datadog.SLOThreshold{
			{
				Target:    97.0,
				Timeframe: datadog.SLOTIMEFRAME_SEVEN_DAYS,
				Warning:   datadog.PtrFloat64(98.0),
			},
		},
		Query: &datadog.ServiceLevelObjectiveQuery{
			Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
			Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelObjectivesApi.UpdateSLO(ctx, SloData0ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.UpdateSLO`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.UpdateSLO`:\n%s\n", responseContent)
}
