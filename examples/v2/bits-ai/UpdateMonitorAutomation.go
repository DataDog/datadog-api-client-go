// Update monitor automatic investigation settings returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.MonitorAutomationRequest{
		Data: datadogV2.MonitorAutomationRequestData{
			Attributes: datadogV2.MonitorAutomationAttributes{
				Enabled: true,
			},
			Type: datadogV2.MONITORAUTOMATIONTYPE_MONITOR_AUTOMATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateMonitorAutomation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsAIApi(apiClient)
	resp, r, err := api.UpdateMonitorAutomation(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsAIApi.UpdateMonitorAutomation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsAIApi.UpdateMonitorAutomation`:\n%s\n", responseContent)
}
