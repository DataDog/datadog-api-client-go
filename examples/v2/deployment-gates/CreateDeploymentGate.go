// Create deployment gate returns "OK" response

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
	body := datadogV2.CreateDeploymentGateParams{
		Data: datadogV2.CreateDeploymentGateParamsData{
			Attributes: datadogV2.CreateDeploymentGateParamsDataAttributes{
				DryRun:     datadog.PtrBool(false),
				Env:        "production",
				Identifier: datadog.PtrString("my-gate-1"),
				Service:    "my-service",
			},
			Type: datadogV2.DEPLOYMENTGATEDATATYPE_DEPLOYMENT_GATE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateDeploymentGate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDeploymentGatesApi(apiClient)
	resp, r, err := api.CreateDeploymentGate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentGatesApi.CreateDeploymentGate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DeploymentGatesApi.CreateDeploymentGate`:\n%s\n", responseContent)
}
