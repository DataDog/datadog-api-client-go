// Create a mute rule returns "Successfully created the mute rule" response

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
	body := datadogV2.MuteRuleCreateRequest{
		Data: datadogV2.MuteRuleDataCreate{
			Attributes: datadogV2.MuteRuleAttributesCreate{
				Action: datadogV2.MuteRuleAction{
					Reason: datadogV2.MUTEREASON_RISK_ACCEPTED,
				},
				Enabled: datadog.PtrBool(true),
				Name:    "Example-Security-Monitoring",
				Rule: datadogV2.AutomationRuleScope{
					FindingTypes: []datadogV2.SecurityFindingType{
						datadogV2.SECURITYFINDINGTYPE_MISCONFIGURATION,
					},
					Query: datadog.PtrString("env:staging"),
				},
			},
			Type: datadogV2.MUTERULETYPE_MUTE_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityFindingsAutomationMuteRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFindingsAutomationMuteRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFindingsAutomationMuteRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFindingsAutomationMuteRule`:\n%s\n", responseContent)
}
