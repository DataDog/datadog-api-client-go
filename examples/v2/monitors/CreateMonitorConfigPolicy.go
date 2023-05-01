// Create a monitor configuration policy returns "OK" response

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
	body := datadogV2.MonitorConfigPolicyCreateRequest{
		Data: datadogV2.MonitorConfigPolicyCreateData{
			Attributes: datadogV2.MonitorConfigPolicyAttributeCreateRequest{
				PolicyType: datadogV2.MONITORCONFIGPOLICYTYPE_TAG,
				Policy: datadogV2.MonitorConfigPolicyPolicyCreateRequest{
					MonitorConfigPolicyTagPolicyCreateRequest: &datadogV2.MonitorConfigPolicyTagPolicyCreateRequest{
						TagKey:         "examplemonitor",
						TagKeyRequired: false,
						ValidTagValues: []string{
							"prod",
							"staging",
						},
					}},
			},
			Type: datadogV2.MONITORCONFIGPOLICYRESOURCETYPE_MONITOR_CONFIG_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitorConfigPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitorConfigPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitorConfigPolicy`:\n%s\n", responseContent)
}
