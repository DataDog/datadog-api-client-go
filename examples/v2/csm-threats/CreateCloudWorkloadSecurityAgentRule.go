// Create a Cloud Workload Security Agent rule returns "OK" response

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
				Description: datadog.PtrString("My Agent rule"),
				Enabled:     datadog.PtrBool(true),
				Expression:  `exec.file.name == "sh"`,
				Filters:     []string{},
				Name:        "examplecsmthreat",
			},
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
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
