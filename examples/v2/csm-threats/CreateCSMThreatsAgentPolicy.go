// Create a CSM Threats Agent policy returns "OK" response

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
	body := datadogV2.CloudWorkloadSecurityAgentPolicyCreateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentPolicyCreateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentPolicyCreateAttributes{
				Description: datadog.PtrString("My agent policy"),
				Enabled:     datadog.PtrBool(true),
				HostTagsLists: [][]string{
					{
						"env:test",
					},
				},
				Name: "my_agent_policy",
			},
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTPOLICYTYPE_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.CreateCSMThreatsAgentPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.CreateCSMThreatsAgentPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.CreateCSMThreatsAgentPolicy`:\n%s\n", responseContent)
}
