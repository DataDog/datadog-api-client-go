// Get a CSM Threats Agent rule returns "OK" response

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
	// there is a valid "agent_rule_rc" in the system
	AgentRuleDataID := os.Getenv("AGENT_RULE_DATA_ID")

	// there is a valid "policy_rc" in the system
	PolicyDataID := os.Getenv("POLICY_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.GetCSMThreatsAgentRule(ctx, AgentRuleDataID, *datadogV2.NewGetCSMThreatsAgentRuleOptionalParameters().WithPolicyId(PolicyDataID))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.GetCSMThreatsAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.GetCSMThreatsAgentRule`:\n%s\n", responseContent)
}
