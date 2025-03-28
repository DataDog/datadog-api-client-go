// Batch update CSM Threats Agent policies returns "OK" response

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

	body := datadogV2.CloudWorkloadSecurityAgentPolicyBatchUpdateRequest{
		Data: datadogV2.CloudWorkloadSecurityAgentPolicyBatchUpdateData{
			Attributes: datadogV2.CloudWorkloadSecurityAgentPolicyBatchUpdateAttributes{
				Policies: []datadogV2.CloudWorkloadSecurityAgentPolicyBatchUpdateAttributesPoliciesItems{
					{
						Delete:      datadog.PtrBool(false),
						Description: datadog.PtrString("Updated agent policy"),
						Enabled:     datadog.PtrBool(true),
						HostTags: []string{
							"env:test",
						},
						Id:       datadog.PtrString(PolicyDataID),
						Name:     datadog.PtrString("updated_agent_policy"),
						Priority: datadog.PtrInt64(20),
					},
				},
			},
			Id:   "batch_update_req",
			Type: datadogV2.CLOUDWORKLOADSECURITYAGENTPOLICYBATCHUPDATEDATATYPE_POLICIES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMThreatsApi(apiClient)
	resp, r, err := api.BatchUpdateCSMThreatsAgentPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMThreatsApi.BatchUpdateCSMThreatsAgentPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMThreatsApi.BatchUpdateCSMThreatsAgentPolicy`:\n%s\n", responseContent)
}
