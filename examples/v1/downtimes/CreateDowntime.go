// Schedule a downtime returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.Downtime{
		Message:  datadog.PtrString("Example-Schedule_a_downtime_returns_OK_response"),
		Start:    datadog.PtrInt64(time.Now().Unix()),
		Timezone: datadog.PtrString("Etc/UTC"),
		Scope: []string{
			"test:examplescheduleadowntimereturnsokresponse",
		},
		Recurrence: *datadogV1.NewNullableDowntimeRecurrence(&datadogV1.DowntimeRecurrence{
			Type:   datadog.PtrString("weeks"),
			Period: datadog.PtrInt32(1),
			WeekDays: []string{
				"Mon",
				"Tue",
				"Wed",
				"Thu",
				"Fri",
			},
			UntilDate: *datadog.NewNullableInt64(datadog.PtrInt64(time.Now().AddDate(0, 0, 21).Unix())),
		}),
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
