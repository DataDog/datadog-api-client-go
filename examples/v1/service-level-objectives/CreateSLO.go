// Create an SLO object returns "OK" response

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
		Description: *datadog.NewNullableString(datadog.PtrString("string")),
		Groups: []string{
			"env:test",
			"role:mysql",
		},
		MonitorIds: []int64{},
		Name:       "Example-Service-Level-Objective",
		Query: &datadogV1.ServiceLevelObjectiveQuery{
			Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
			Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
		},
		Tags: []string{
			"env:prod",
			"app:core",
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
