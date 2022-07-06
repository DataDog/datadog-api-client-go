// Create an SLO object returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.ServiceLevelObjectiveRequest{
		Type:        datadog.SLOTYPE_METRIC,
		Description: *common.NewNullableString(datadog.PtrString("string")),
		Groups: []string{
			"env:test",
			"role:mysql",
		},
		MonitorIds: []int64{},
		Name:       "Example-Create_an_SLO_object_returns_OK_response",
		Query: &datadog.ServiceLevelObjectiveQuery{
			Denominator: "sum:httpservice.hits{!code:3xx}.as_count()",
			Numerator:   "sum:httpservice.hits{code:2xx}.as_count()",
		},
		Tags: []string{
			"env:prod",
			"app:core",
		},
		Thresholds: []datadog.SLOThreshold{
			{
				Target:         95.0,
				TargetDisplay:  datadog.PtrString("95.0"),
				Timeframe:      datadog.SLOTIMEFRAME_SEVEN_DAYS,
				Warning:        datadog.PtrFloat64(98),
				WarningDisplay: datadog.PtrString("98.0"),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.ServiceLevelObjectivesApi(apiClient)
	resp, r, err := api.CreateSLO(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSLO`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.CreateSLO`:\n%s\n", responseContent)
}
