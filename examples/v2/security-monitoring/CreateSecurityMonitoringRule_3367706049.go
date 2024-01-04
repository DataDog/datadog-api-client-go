// Create a detection rule with detection method 'third_party' returns "OK" response

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
		SecurityMonitoringStandardRuleCreatePayload: &datadogV2.SecurityMonitoringStandardRuleCreatePayload{
			Name:      "Example-Security-Monitoring",
			Type:      datadogV2.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
			IsEnabled: true,
			ThirdPartyCases: []datadogV2.SecurityMonitoringThirdPartyRuleCaseCreate{
				{
					Query:  datadog.PtrString("status:error"),
					Name:   datadog.PtrString("high"),
					Status: datadogV2.SECURITYMONITORINGRULESEVERITY_HIGH,
				},
				{
					Query:  datadog.PtrString("status:info"),
					Name:   datadog.PtrString("low"),
					Status: datadogV2.SECURITYMONITORINGRULESEVERITY_LOW,
				},
			},
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{},
			Cases:   []datadogV2.SecurityMonitoringRuleCaseCreate{},
			Message: "This is a third party rule",
			Options: datadogV2.SecurityMonitoringRuleOptions{
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THIRD_PARTY.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ZERO_MINUTES.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES.Ptr(),
				ThirdPartyRuleOptions: &datadogV2.SecurityMonitoringRuleThirdPartyOptions{
					DefaultStatus: datadogV2.SECURITYMONITORINGRULESEVERITY_INFO.Ptr(),
					RootQueries: []datadogV2.SecurityMonitoringThirdPartyRootQuery{
						{
							Query: datadog.PtrString("source:guardduty @details.alertType:*EC2*"),
							GroupByFields: []string{
								"instance-id",
							},
						},
						{
							Query:         datadog.PtrString("source:guardduty"),
							GroupByFields: []string{},
						},
					},
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
