// Update a CSM Threats Agent policy returns "OK" response

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

	body := datadogV2.CloudWorkloadSecurityAgentPolicyUpdateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentPolicyUpdateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentPolicyUpdateAttributes{
				Description: datadog.PtrString("Updated agent policy"),
				Enabled:     datadog.PtrBool(true),
				HostTagsLists: [][]string{
					{
						"env:test",
					},
				},
				Name: datadog.PtrString("updated_agent_policy"),
			},
			Id:   datadog.PtrString(PolicyDataID),
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTPOLICYTYPE_POLICY,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.UpdateCSMThreatsAgentPolicy(ctx, PolicyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.UpdateCSMThreatsAgentPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.UpdateCSMThreatsAgentPolicy`:\n%s\n", responseContent)
}
