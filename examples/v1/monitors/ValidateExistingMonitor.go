// Validate an existing monitor returns "OK" response

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

	body := datadog.Monitor{
		Name:    datadog.PtrString("Example-Validate_an_existing_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_LOG_ALERT,
		Query:   `logs("service:foo AND type:error").index("main").rollup("count").by("source").last("5m") > 2`,
		Message: datadog.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplevalidateanexistingmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *common.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			EnableLogsSample:     datadog.PtrBool(true),
			EscalationMessage:    datadog.PtrString("the situation has escalated"),
			EvaluationDelay:      *common.NewNullableInt64(datadog.PtrInt64(700)),
			GroupbySimpleMonitor: datadog.PtrBool(true),
			IncludeTags:          datadog.PtrBool(true),
			Locked:               datadog.PtrBool(false),
			NewHostDelay:         *common.NewNullableInt64(datadog.PtrInt64(600)),
			NoDataTimeframe:      *common.NewNullableInt64(nil),
			NotifyAudit:          datadog.PtrBool(false),
			NotifyNoData:         datadog.PtrBool(false),
			RenotifyInterval:     *common.NewNullableInt64(datadog.PtrInt64(60)),
			RequireFullWindow:    datadog.PtrBool(true),
			TimeoutH:             *common.NewNullableInt64(datadog.PtrInt64(24)),
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(2),
				Warning:  *common.NewNullableFloat64(datadog.PtrFloat64(1)),
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMonitorsApi(apiClient)
	resp, r, err := api.ValidateExistingMonitor(ctx, MonitorID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateExistingMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ValidateExistingMonitor`:\n%s\n", responseContent)
}
