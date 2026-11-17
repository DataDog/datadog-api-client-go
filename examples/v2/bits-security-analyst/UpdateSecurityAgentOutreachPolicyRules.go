// Update outreach policy rules returns "OK" response

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
	body := datadogV2.SecurityAgentOutreachRulesUpdateRequest{
		Data: datadogV2.SecurityAgentOutreachRulesUpdateRequestData{
			Attributes: datadogV2.SecurityAgentOutreachRulesUpdateAttributes{
				Rules: []datadogV2.SecurityAgentOutreachRuleUpdate{
					{
						IsEnabled: true,
						RuleId:    "abc-def-ghi",
					},
				},
			},
			Type: datadogV2.SECURITYAGENTOUTREACHRULESUPDATETYPE_OUTREACH_POLICY_RULES_UPDATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityAgentOutreachPolicyRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewBitsSecurityAnalystApi(apiClient)
	resp, r, err := api.UpdateSecurityAgentOutreachPolicyRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BitsSecurityAnalystApi.UpdateSecurityAgentOutreachPolicyRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `BitsSecurityAnalystApi.UpdateSecurityAgentOutreachPolicyRules`:\n%s\n", responseContent)
}
