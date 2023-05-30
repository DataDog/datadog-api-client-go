// Edit a monitor configuration policy returns "OK" response

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

	body := datadogV2.MonitorConfigPolicyEditRequest{
		Data: datadogV2.MonitorConfigPolicyEditData{
			Attributes: datadogV2.MonitorConfigPolicyAttributeEditRequest{
				Policy: datadogV2.MonitorConfigPolicyPolicy{
					MonitorConfigPolicyTagPolicy: &datadogV2.MonitorConfigPolicyTagPolicy{
						TagKey:         datadog.PtrString("examplemonitor"),
						TagKeyRequired: datadog.PtrBool(false),
						ValidTagValues: []string{
							"prod",
							"staging",
						},
					}},
				PolicyType: datadogV2.MONITORCONFIGPOLICYTYPE_TAG,
			},
			Id:   MonitorConfigurationPolicyDataID,
			Type: datadogV2.MONITORCONFIGPOLICYRESOURCETYPE_MONITOR_CONFIG_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.UpdateMonitorConfigPolicy(ctx, MonitorConfigurationPolicyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitorConfigPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitorConfigPolicy`:\n%s\n", responseContent)
}
