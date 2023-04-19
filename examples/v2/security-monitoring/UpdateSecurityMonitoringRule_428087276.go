// Update a cloud configuration rule's details returns "OK" response

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
	// there is a valid "cloud_configuration_rule" in the system
	CloudConfigurationRuleID := os.Getenv("CLOUD_CONFIGURATION_RULE_ID")

	body := datadogV2.SecurityMonitoringRuleUpdatePayload{
		Name:      datadog.PtrString("Example-Security-Monitoring_cloud_updated"),
		IsEnabled: datadog.PtrBool(false),
		Cases: []datadogV2.SecurityMonitoringRuleCase{
			{
				Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO.Ptr(),
				Notifications: []string{},
			},
		},
		Options: &datadogV2.SecurityMonitoringRuleOptions{
			ComplianceRuleOptions: &datadogV2.CloudConfigurationComplianceRuleOptions{
				ResourceType: datadog.PtrString("gcp_compute_disk"),
				RegoRule: &datadogV2.CloudConfigurationRegoRule{
					Policy: `package datadog
`,
					ResourceTypes: []string{
						"gcp_compute_disk",
					},
				},
			},
		},
		Message: datadog.PtrString("ddd"),
		Tags:    []string{},
		ComplianceSignalOptions: &datadogV2.CloudConfigurationRuleComplianceSignalOptions{
			UserActivationStatus: *datadog.NewNullableBool(datadog.PtrBool(false)),
			UserGroupByFields:    []string{},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringRule(ctx, CloudConfigurationRuleID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`:\n%s\n", responseContent)
}
