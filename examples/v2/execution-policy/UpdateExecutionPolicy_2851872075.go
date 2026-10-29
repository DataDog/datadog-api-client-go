// Update an execution policy to Slack with scope and targets returns "OK" response

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
	// there is a valid "execution_policy" in the system
	ExecutionPolicyDataID := os.Getenv("EXECUTION_POLICY_DATA_ID")

	body := datadogV2.ExecutionPolicyUpdateRequest{
		Data: datadogV2.ExecutionPolicyUpdateRequestData{
			Id:   ExecutionPolicyDataID,
			Type: datadogV2.EXECUTIONPOLICYTYPE_EXECUTION_POLICY,
			Attributes: datadogV2.ExecutionPolicyWriteAttributes{
				Name:   "Cassette Execution Policy Updated",
				Effect: datadogV2.EXECUTIONPOLICYEFFECT_DENY,
				ActionPattern: datadogV2.ExecutionPolicyActionPattern{
					Integration: datadogV2.EXECUTIONPOLICYINTEGRATION_INTEGRATION_SLACK,
					ActionFqns: []string{
						"*",
					},
				},
				Scope: &datadogV2.ExecutionPolicyScope{
					Slack: &datadogV2.ExecutionPolicySlackScope{
						Rules: []datadogV2.ExecutionPolicySlackScopeRule{
							{
								ChannelNameSelectors: []string{
									"incident-denied",
								},
							},
						},
					},
				},
				Targets: []datadogV2.ExecutionPolicyTarget{
					{
						Name: *datadog.NewNullableString(datadog.PtrString("Incident workspace")),
						AgentTags: []string{
							"slack_workspace_id:T01234567",
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateExecutionPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExecutionPolicyApi(apiClient)
	resp, r, err := api.UpdateExecutionPolicy(ctx, ExecutionPolicyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionPolicyApi.UpdateExecutionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExecutionPolicyApi.UpdateExecutionPolicy`:\n%s\n", responseContent)
}
