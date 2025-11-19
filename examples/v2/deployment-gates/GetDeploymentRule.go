// Get deployment rule returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetDeploymentRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.GetDeploymentRule(ctx, DeploymentGateDataID, DeploymentRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.GetDeploymentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.GetDeploymentRule`:\n%s\n", responseContent)
}
