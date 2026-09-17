// Create an LLM Observability monitor returns "OK" response

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
		Type:    datadogV1.MONITORTYPE_LLM_OBSERVABILITY_ALERT,
		Query:   `llm-observability("*").rollup("count").last("2h") > 0`,
		Message: datadog.PtrString("LLM observability alert triggered"),
		Tags: []string{
			"test:examplemonitor",
			"env:ci",
		},
		Options: &datadogV1.MonitorOptions{
			Thresholds: &datadogV1.MonitorThresholds{
				Critical: datadog.PtrFloat64(0),
			},
			IncludeTags: datadog.PtrBool(true),
			NotifyAudit: datadog.PtrBool(false),
		},
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
