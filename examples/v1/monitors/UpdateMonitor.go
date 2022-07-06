// Edit a monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadog.MonitorUpdateRequest{
		Name: datadog.PtrString("My monitor-updated"),
		Options: &datadog.MonitorOptions{
			EvaluationDelay:  *common.NewNullableInt64(nil),
			NewGroupDelay:    *common.NewNullableInt64(datadog.PtrInt64(600)),
			NewHostDelay:     *common.NewNullableInt64(nil),
			RenotifyInterval: *common.NewNullableInt64(nil),
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *common.NewNullableFloat64(nil),
			},
			TimeoutH: *common.NewNullableInt64(nil),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.MonitorsApi(apiClient)
	resp, r, err := api.UpdateMonitor(ctx, MonitorID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitor`:\n%s\n", responseContent)
}
