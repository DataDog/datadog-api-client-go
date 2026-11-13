// Create an execution policy with transitional remote shell command omission returns "Created" response

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
	body := datadogV2.ExecutionPolicyCreateRequest{
		Data: datadogV2.ExecutionPolicyCreateRequestData{
			Type: datadogV2.EXECUTIONPOLICYTYPE_EXECUTION_POLICY,
			Attributes: datadogV2.ExecutionPolicyWriteAttributes{
				Name:   "Legacy remote shell policy exampleexecutionpolicy",
				Effect: datadogV2.EXECUTIONPOLICYEFFECT_ALLOW,
				ActionPattern: datadogV2.ExecutionPolicyActionPattern{
					Integration: datadogV2.EXECUTIONPOLICYINTEGRATION_INTEGRATION_REMOTE_ACTION,
					ActionFqns: []string{
						"com.datadoghq.remoteaction.rshell.runCommand",
					},
				},
				Scope: &datadogV2.ExecutionPolicyScope{
					RemoteActionRshell: &datadogV2.ExecutionPolicyRemoteActionRshellScope{
						Rules: []datadogV2.ExecutionPolicyRemoteActionRshellScopeRule{
							{
								TargetPaths: []string{
									"/var/log",
								},
								Access: datadogV2.EXECUTIONPOLICYREMOTEACTIONRSHELLACCESS_READ_ONLY,
							},
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateExecutionPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExecutionPolicyApi(apiClient)
	resp, r, err := api.CreateExecutionPolicy(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionPolicyApi.CreateExecutionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExecutionPolicyApi.CreateExecutionPolicy`:\n%s\n", responseContent)
}
