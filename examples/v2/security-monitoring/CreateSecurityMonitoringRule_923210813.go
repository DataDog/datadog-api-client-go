// Create a cloud configuration rule returns "OK" response

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
			Name: "Example-Create_a_cloud_configuration_rule_returns_OK_response",
			Cases: []datadogV2.CloudConfigurationRuleCaseCreate{
				{
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{},
				},
			},
			Options: datadogV2.CloudConfigurationRuleOptions{},
			ComplianceSignalOptions: datadogV2.CloudConfigurationRuleComplianceSignalOptions{
				UserActivationStatus: datadog.PtrBool(false),
				UserGroupByFields:    []string{},
			},
			Message:   "Test rule",
			Tags:      []string{},
			IsEnabled: true,
			Type:      datadogV2.CLOUDCONFIGURATIONRULETYPE_CLOUD_CONFIGURATION.Ptr(),
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
