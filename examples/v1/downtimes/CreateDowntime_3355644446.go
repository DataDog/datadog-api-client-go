// Schedule a monitor downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadog.Downtime{
		Message:  datadog.PtrString("Example-Schedule_a_monitor_downtime_returns_OK_response"),
		Start:    datadog.PtrInt64(time.Now().Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope: []string{
			"test:examplescheduleamonitordowntimereturnsokresponse",
		},
		MonitorId: *datadog.NewNullableInt64(datadog.PtrInt64(MonitorID)),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DowntimesApi.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
