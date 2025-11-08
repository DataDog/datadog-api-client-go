// Update deployment rule returns "OK" response

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
	// there is a valid "deployment_gate" in the system
	DeploymentGateDataID := os.Getenv("DEPLOYMENT_GATE_DATA_ID")

	// there is a valid "deployment_rule" in the system
	DeploymentRuleDataID := os.Getenv("DEPLOYMENT_RULE_DATA_ID")

	body := datadogV2.UpdateDeploymentRuleParams{
		Data: datadogV2.UpdateDeploymentRuleParamsData{
			Attributes: datadogV2.UpdateDeploymentRuleParamsDataAttributes{
				DryRun: false,
				Name:   "Updated deployment rule",
				Options: datadogV2.DeploymentRulesOptions{
					DeploymentRuleOptionsFaultyDeploymentDetection: &datadogV2.DeploymentRuleOptionsFaultyDeploymentDetection{
						ExcludedResources: []string{},
					}},
			},
			Type: datadogV2.DEPLOYMENTRULEDATATYPE_DEPLOYMENT_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDeploymentRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.UpdateDeploymentRule(ctx, DeploymentGateDataID, DeploymentRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.UpdateDeploymentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.UpdateDeploymentRule`:\n%s\n", responseContent)
}
