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

import data.datadog.output as dd_output

import future.keywords.contains
import future.keywords.if
import future.keywords.in

milliseconds_in_a_day := ((1000 * 60) * 60) * 24

eval(iam_service_account_key) = "skip" if {
	iam_service_account_key.disabled
} else = "pass" if {
	(iam_service_account_key.resource_seen_at / milliseconds_in_a_day) - (iam_service_account_key.valid_after_time / milliseconds_in_a_day) <= 90
} else = "fail"

# This part remains unchanged for all rules
results contains result if {
	some resource in input.resources[input.main_resource_type]
	result := dd_output.format(resource, eval(resource))
}
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
				UserActivationStatus: *datadog.NewNullableBool(datadog.PtrBool(true)),
				UserGroupByFields: *datadog.NewNullableList(&[]string{
					"@account_id",
				}),
			},
			Filters: []datadogV2.SecurityMonitoringFilter{
				{
					Action: datadogV2.SECURITYMONITORINGFILTERACTION_REQUIRE.Ptr(),
					Query:  datadog.PtrString("resource_id:helo*"),
				},
				{
					Action: datadogV2.SECURITYMONITORINGFILTERACTION_SUPPRESS.Ptr(),
					Query:  datadog.PtrString("control:helo*"),
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
