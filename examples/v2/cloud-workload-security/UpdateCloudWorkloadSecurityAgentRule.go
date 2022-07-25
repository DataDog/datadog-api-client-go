// Update a Cloud Workload Security Agent rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	// there is a valid "agent_rule" in the system
	AgentRuleDataID := os.Getenv("AGENT_RULE_DATA_ID")

	body := datadog.CloudWorkloadSecurityAgentRuleUpdateRequest{
		Data: datadog.CloudWorkloadSecurityAgentRuleUpdateData{
			Attributes: datadog.CloudWorkloadSecurityAgentRuleUpdateAttributes{
				Description: common.PtrString("Test Agent rule"),
				Enabled:     common.PtrBool(true),
				Expression:  common.PtrString(`exec.file.name == "sh"`),
			},
			Type: datadog.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewCloudWorkloadSecurityApi(apiClient)
	resp, r, err := api.UpdateCloudWorkloadSecurityAgentRule(ctx, AgentRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule`:\n%s\n", responseContent)
}
