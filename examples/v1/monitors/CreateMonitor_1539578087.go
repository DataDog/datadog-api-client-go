// Create a metric monitor with a custom schedule returns "OK" response

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
	body := datadogV1.Monitor{
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Name:    datadog.PtrString("Example-Monitor"),
		Query:   "avg(current_1mo):avg:system.load.5{*} > 0.5",
		Tags:    []string{},
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(0.5),
			},
			NotifyAudit: datadog.PtrBool(false),
			IncludeTags: datadog.PtrBool(false),
			SchedulingOptions: &datadogV1.MonitorOptionsSchedulingOptions{
				EvaluationWindow: &datadogV1.MonitorOptionsSchedulingOptionsEvaluationWindow{
					DayStarts:   datadog.PtrString("04:00"),
					MonthStarts: datadog.PtrInt32(1),
				},
				CustomSchedule: &datadogV1.MonitorOptionsCustomSchedule{
					Recurrences: []datadogV1.MonitorOptionsCustomScheduleRecurrence{
						{
							Rrule:    datadog.PtrString("FREQ=DAILY;INTERVAL=1"),
							Timezone: datadog.PtrString("America/Los_Angeles"),
							Start:    datadog.PtrString("2024-10-26T09:13:00"),
						},
					},
				},
			},
		},
		Type: datadogV1.MONITORTYPE_QUERY_ALERT,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
