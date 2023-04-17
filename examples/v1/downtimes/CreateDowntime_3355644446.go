// Schedule a monitor downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadogV1.Downtime{
		Message:  *datadog.NewNullableString(datadog.PtrString("Example-Downtime")),
		Start:    datadog.PtrInt64(time.Now().Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope: []string{
			"test:exampledowntime",
		},
		MonitorId: *datadog.NewNullableInt64(datadog.PtrInt64(MonitorID)),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DowntimesApi.CreateDowntime`:\n%s\n", responseContent)
}
