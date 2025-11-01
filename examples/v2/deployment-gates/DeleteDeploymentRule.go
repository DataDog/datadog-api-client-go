// Delete deployment rule returns "No Content" response

package main

import (
	"context"
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
	configuration.SetUnstableOperationEnabled("v2.DeleteDeploymentRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	r, err := api.DeleteDeploymentRule(ctx, DeploymentGateDataID, DeploymentRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.DeleteDeploymentRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
