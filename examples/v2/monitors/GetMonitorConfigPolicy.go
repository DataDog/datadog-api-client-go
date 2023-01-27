// Get a monitor configuration policy returns "OK" response

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
	// there is a valid "monitor_configuration_policy" in the system
	MonitorConfigurationPolicyDataID := os.Getenv("MONITOR_CONFIGURATION_POLICY_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.GetMonitorConfigPolicy(ctx, MonitorConfigurationPolicyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitorConfigPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitorConfigPolicy`:\n%s\n", responseContent)
}
