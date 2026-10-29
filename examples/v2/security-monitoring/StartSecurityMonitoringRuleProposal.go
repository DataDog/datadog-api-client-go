// Start a rule proposal returns "Accepted" response

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
	body := datadogV2.SecurityMonitoringRuleProposalStartRequest{
		Data: datadogV2.SecurityMonitoringRuleProposalStartRequestData{
			Type: datadogV2.SECURITYMONITORINGRULEPROPOSALTYPE_RULE_PROPOSALS,
			Attributes: datadogV2.SecurityMonitoringRuleProposalStartRequestAttributes{
				InputType: datadogV2.SECURITYMONITORINGRULEPROPOSALINPUTTYPE_NATURAL_LANGUAGE,
				Input:     "Detect repeated failed Okta logins followed by a successful login from the same IP.",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.StartSecurityMonitoringRuleProposal", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.StartSecurityMonitoringRuleProposal(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.StartSecurityMonitoringRuleProposal`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.StartSecurityMonitoringRuleProposal`:\n%s\n", responseContent)
}
