// Delete a Cloud Workload Security Agent rule returns "OK" response

package main

import (
	"context"
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
	r, err := api.DeleteCloudWorkloadSecurityAgentRule(ctx, AgentRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.DeleteCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
