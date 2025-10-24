// Update deployment gate returns "OK" response

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

	body := datadogV2.UpdateDeploymentGateParams{
		Data: datadogV2.UpdateDeploymentGateParamsData{
			Attributes: datadogV2.UpdateDeploymentGateParamsDataAttributes{
				DryRun: false,
			},
			Id:   "12345678-1234-1234-1234-123456789012",
			Type: datadogV2.DEPLOYMENTGATEDATATYPE_DEPLOYMENT_GATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateDeploymentGate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.UpdateDeploymentGate(ctx, DeploymentGateDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.UpdateDeploymentGate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.UpdateDeploymentGate`:\n%s\n", responseContent)
}
