// Validate an existing monitor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadog.Monitor{
		Name:    common.PtrString("Example-Validate_an_existing_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_LOG_ALERT,
		Query:   `logs("service:foo AND type:error").index("main").rollup("count").by("source").last("5m") > 2`,
		Message: common.PtrString("some message Notify: @hipchat-channel"),
		Tags: []string{
			"test:examplevalidateanexistingmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *common.NewNullableInt64(common.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			EnableLogsSample:     common.PtrBool(true),
			EscalationMessage:    common.PtrString("the situation has escalated"),
			EvaluationDelay:      *common.NewNullableInt64(common.PtrInt64(700)),
			GroupbySimpleMonitor: common.PtrBool(true),
			IncludeTags:          common.PtrBool(true),
			Locked:               common.PtrBool(false),
			NewHostDelay:         *common.NewNullableInt64(common.PtrInt64(600)),
			NoDataTimeframe:      *common.NewNullableInt64(nil),
			NotifyAudit:          common.PtrBool(false),
			NotifyNoData:         common.PtrBool(false),
			RenotifyInterval:     *common.NewNullableInt64(common.PtrInt64(60)),
			RequireFullWindow:    common.PtrBool(true),
			TimeoutH:             *common.NewNullableInt64(common.PtrInt64(24)),
			Thresholds: &datadog.MonitorThresholds{
				Critical: common.PtrFloat64(2),
				Warning:  *common.NewNullableFloat64(common.PtrFloat64(1)),
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
