// Create a CSM Threats Agent rule returns "OK" response

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
	// there is a valid "policy_rc" in the system
	PolicyDataID := os.Getenv("POLICY_DATA_ID")

	body := datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentRuleCreateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentRuleCreateAttributes{
				Description: datadog.PtrString("My Agent rule"),
				Enabled:     datadog.PtrBool(true),
				Expression:  `exec.file.name == "sh"`,
				Filters:     []string{},
				Name:        "examplecsmthreat",
				PolicyId:    datadog.PtrString(PolicyDataID),
				ProductTags: []string{},
			},
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.CreateCSMThreatsAgentRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.CreateCSMThreatsAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.CreateCSMThreatsAgentRule`:\n%s\n", responseContent)
}
