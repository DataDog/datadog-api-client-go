// Disable automatic investigations for a monitor

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "monitor" in the system
	MonitorID, _ := strconv.ParseInt(os.Getenv("MONITOR_ID"), 10, 64)

	body := datadogV2.MonitorAutomationRequest{
		Data: datadogV2.MonitorAutomationRequestData{
			Type: datadogV2.MONITORAUTOMATIONTYPE_MONITOR_AUTOMATION,
			Attributes: datadogV2.MonitorAutomationAttributes{
				Enabled: false,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateMonitorAutomation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsAIApi(apiClient)
	resp, r, err := api.UpdateMonitorAutomation(ctx, MonitorID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsAIApi.UpdateMonitorAutomation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsAIApi.UpdateMonitorAutomation`:\n%s\n", responseContent)
}
