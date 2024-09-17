// Edit a monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadogV1.MonitorUpdateRequest{
		Name:     datadog.PtrString("My monitor-updated"),
		Priority: *datadog.NewNullableInt64(nil),
		Options: &datadogV1.MonitorOptions{
			EvaluationDelay:  *datadog.NewNullableInt64(nil),
			NewGroupDelay:    *datadog.NewNullableInt64(datadog.PtrInt64(600)),
			NewHostDelay:     *datadog.NewNullableInt64(nil),
			RenotifyInterval: *datadog.NewNullableInt64(nil),
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *datadog.NewNullableFloat64(nil),
			},
			TimeoutH: *datadog.NewNullableInt64(nil),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.UpdateMonitor(ctx, MonitorID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitor`:\n%s\n", responseContent)
}
