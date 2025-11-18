// Create deployment rule returns "OK" response

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

	body := datadogV2.CreateDeploymentRuleParams{
		Data: &datadogV2.CreateDeploymentRuleParamsData{
			Attributes: datadogV2.CreateDeploymentRuleParamsDataAttributes{
				DryRun: datadog.PtrBool(false),
				Name:   "My deployment rule",
				Options: datadogV2.DeploymentRulesOptions{
					DeploymentRuleOptionsFaultyDeploymentDetection: &datadogV2.DeploymentRuleOptionsFaultyDeploymentDetection{
						ExcludedResources: []string{},
					}},
				Type: "faulty_deployment_detection",
			},
			Type: datadogV2.DEPLOYMENTRULEDATATYPE_DEPLOYMENT_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDeploymentRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.CreateDeploymentRule(ctx, DeploymentGateDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.CreateDeploymentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.CreateDeploymentRule`:\n%s\n", responseContent)
}
