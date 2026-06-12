// Trigger a deployment gate evaluation returns "Accepted" response

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
	body := datadogV2.DeploymentGatesEvaluationRequest{
		Data: datadogV2.DeploymentGatesEvaluationRequestData{
			Attributes: datadogV2.DeploymentGatesEvaluationRequestAttributes{
				Configuration: &datadogV2.DeploymentGatesEvaluationConfiguration{
					DryRun: datadog.PtrBool(false),
					Rules: []datadogV2.DeploymentGatesEvaluationRule{
						datadogV2.DeploymentGatesEvaluationRule{
							DeploymentGatesMonitorRule: &datadogV2.DeploymentGatesMonitorRule{
								DryRun: datadog.PtrBool(false),
								Name:   "error rate monitors",
								Options: &datadogV2.DeploymentGatesMonitorRuleOptions{
									Duration: datadog.PtrInt64(300),
									Query:    "service:transaction-backend env:production",
								},
								Type: datadogV2.DEPLOYMENTGATESMONITORRULETYPE_MONITOR,
							}},
					},
				},
				Env:        "staging",
				Identifier: datadog.PtrString("pre-deploy"),
				PrimaryTag: datadog.PtrString("region:us-east-1"),
				Service:    "transaction-backend",
				Version:    datadog.PtrString("v1.2.3"),
			},
			Type: datadogV2.DEPLOYMENTGATESEVALUATIONREQUESTDATATYPE_DEPLOYMENT_GATES_EVALUATION_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.TriggerDeploymentGatesEvaluation", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.TriggerDeploymentGatesEvaluation(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.TriggerDeploymentGatesEvaluation`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.TriggerDeploymentGatesEvaluation`:\n%s\n", responseContent)
}
