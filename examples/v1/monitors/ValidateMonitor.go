// Validate a monitor returns "OK" response

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
		Name:    datadog.PtrString("Example-Monitor"),
		Type:    datadogV1.MONITORTYPE_LOG_ALERT,
		Query:   `logs("service:foo AND type:error").index("main").rollup("count").by("source").last("5m") > 2`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadogV1.MonitorOptions{
			EnableLogsSample:       datadog.PtrBool(true),
			EscalationMessage:      datadog.PtrString("the situation has escalated"),
			EvaluationDelay:        *datadog.NewNullableInt64(datadog.PtrInt64(700)),
			GroupbySimpleMonitor:   datadog.PtrBool(true),
			IncludeTags:            datadog.PtrBool(true),
			Locked:                 datadog.PtrBool(false),
			NewHostDelay:           *datadog.NewNullableInt64(datadog.PtrInt64(600)),
			NoDataTimeframe:        *datadog.NewNullableInt64(nil),
			NotifyAudit:            datadog.PtrBool(false),
			NotifyNoData:           datadog.PtrBool(false),
			OnMissingData:          datadogV1.ONMISSINGDATAOPTION_SHOW_AND_NOTIFY_NO_DATA.Ptr(),
			NotificationPresetName: datadogV1.MONITOROPTIONSNOTIFICATIONPRESETS_HIDE_HANDLES.Ptr(),
			RenotifyInterval:       *datadog.NewNullableInt64(datadog.PtrInt64(60)),
			RequireFullWindow:      datadog.PtrBool(true),
			TimeoutH:               *datadog.NewNullableInt64(datadog.PtrInt64(24)),
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *datadog.NewNullableFloat64(datadog.PtrFloat64(1)),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.ValidateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ValidateMonitor`:\n%s\n", responseContent)
}
