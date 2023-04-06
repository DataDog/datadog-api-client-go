// Create a cloud_configuration rule returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleCreatePayload{
		CloudConfigurationRuleCreatePayload: &datadogV2.CloudConfigurationRuleCreatePayload{
			Type:      datadogV2.CLOUDCONFIGURATIONRULETYPE_CLOUD_CONFIGURATION.Ptr(),
			Name:      "Example-Security-Monitoring_cloud",
			IsEnabled: false,
			Cases: []datadogV2.CloudConfigurationRuleCaseCreate{
				{
					Status: datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{
						"channel",
					},
				},
			},
			Options: datadogV2.CloudConfigurationRuleOptions{
				ComplianceRuleOptions: datadogV2.CloudConfigurationComplianceRuleOptions{
					ResourceType: datadog.PtrString("gcp_compute_disk"),
					ComplexRule:  datadog.PtrBool(false),
					RegoRule: &datadogV2.CloudConfigurationRegoRule{
						Policy: `package datadog
`,
						ResourceTypes: []string{
							"gcp_compute_disk",
						},
					},
				},
			},
			Message: "ddd",
			Tags: []string{
				"my:tag",
			},
			ComplianceSignalOptions: datadogV2.CloudConfigurationRuleComplianceSignalOptions{
				UserActivationStatus: datadog.PtrBool(true),
				UserGroupByFields: []string{
					"@account_id",
				},
			},
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
