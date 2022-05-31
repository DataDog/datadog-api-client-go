// Create an Error Tracking monitor returns "OK" response

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
		Name:    datadog.PtrString("Example-Create_an_Error_Tracking_monitor_returns_OK_response"),
		Type:    datadog.MONITORTYPE_ERROR_TRACKING_ALERT,
		Query:   `error-tracking-rum("service:foo AND @error.source:source").rollup("count").by("@issue.id").last("1h") >= 1`,
		Message: datadog.PtrString("some message"),
		Tags: []string{
			"test:examplecreateanerrortrackingmonitorreturnsokresponse",
			"env:ci",
		},
		Priority: *datadog.NewNullableInt64(datadog.PtrInt64(3)),
		Options: &datadog.MonitorOptions{
			Thresholds: &datadog.MonitorThresholds{
				Critical: datadog.PtrFloat64(1),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsApi.CreateMonitor(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`:\n%s\n", responseContent)
}
