// Create a Workload Protection agent rule (US1-FED) returns "OK" response

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
	body := datadogV2.CloudWorkloadSecurityAgentRuleCreateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentRuleCreateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentRuleCreateAttributes{
				Actions: []datadogV2.CloudWorkloadSecurityAgentRuleAction{
					{
						Coredump:      &datadogV2.CloudWorkloadSecurityAgentRuleActionCoreDump{},
						Disabled:      datadog.PtrBool(false),
						Hash:          &datadogV2.CloudWorkloadSecurityAgentRuleActionHash{},
						Kill:          &datadogV2.CloudWorkloadSecurityAgentRuleKill{},
						Log:           &datadogV2.CloudWorkloadSecurityAgentRuleActionLog{},
						Metadata:      &datadogV2.CloudWorkloadSecurityAgentRuleActionMetadata{},
						NetworkFilter: &datadogV2.CloudWorkloadSecurityAgentRuleActionNetworkFilter{},
						Set:           &datadogV2.CloudWorkloadSecurityAgentRuleActionSet{},
					},
				},
				Blocking:    []string{},
				Description: datadog.PtrString("My Agent rule"),
				Disabled:    []string{},
				Enabled:     datadog.PtrBool(true),
				Expression:  `exec.file.name == "sh"`,
				Filters:     []string{},
				Monitoring:  []string{},
				Name:        "my_agent_rule",
				PolicyId:    datadog.PtrString("a8c8e364-6556-434d-b798-a4c23de29c0b"),
				ProductTags: []string{},
				Silent:      datadog.PtrBool(false),
			},
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.CreateCloudWorkloadSecurityAgentRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.CreateCloudWorkloadSecurityAgentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.CreateCloudWorkloadSecurityAgentRule`:\n%s\n", responseContent)
}
