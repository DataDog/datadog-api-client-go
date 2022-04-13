// Validate a monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Monitor{
		Name:    datadog.PtrString("Example-Validate_a_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_LOG_ALERT,
		Query:   `logs("service:foo AND type:error").index("main").rollup("count").by("source").last("5m") > 2`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplevalidateamonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			EnableLogsSample:     datadog.PtrBool(true),
			EscalationMessage:    datadog.PtrString("the situation has escalated"),
			EvaluationDelay:      *datadog.NewNullableInt64(datadog.PtrInt64(700)),
			GroupbySimpleMonitor: datadog.PtrBool(true),
			IncludeTags:          datadog.PtrBool(true),
			Locked:               datadog.PtrBool(false),
			NewHostDelay:         *datadog.NewNullableInt64(datadog.PtrInt64(600)),
			NoDataTimeframe:      *datadog.NewNullableInt64(nil),
			NotifyAudit:          datadog.PtrBool(false),
			NotifyNoData:         datadog.PtrBool(false),
			RenotifyInterval:     *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			RequireFullWindow:    datadog.PtrBool(true),
			TimeoutH:             *datadog.NewNullableInt64(datadog.PtrInt64(24)),
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *datadog.NewNullableFloat64(datadog.PtrFloat64(1)),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsApi.ValidateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ValidateMonitor`:\n%s\n", responseContent)
}
