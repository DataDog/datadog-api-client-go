// Get a Cloud Workload Security Agent rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "agent_rule" in the system
	AgentRuleDataID := os.Getenv("AGENT_RULE_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.CloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.GetCloudWorkloadSecurityAgentRule(ctx, AgentRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule`:\n%s\n", responseContent)
}
