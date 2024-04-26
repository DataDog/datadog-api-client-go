// Get a Cloud Workload Security Agent rule returns "OK" response

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
	// there is a valid "agent_rule" in the system
	AgentRuleDataID := os.Getenv("AGENT_RULE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.GetCloudWorkloadSecurityAgentRule(ctx, AgentRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.GetCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.GetCloudWorkloadSecurityAgentRule`:\n%s\n", responseContent)
}
