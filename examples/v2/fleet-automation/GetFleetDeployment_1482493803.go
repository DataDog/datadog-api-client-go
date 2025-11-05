// Get a deployment by ID returns "OK" response

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
	// there is a valid "deployment" in the system
	DeploymentID := os.Getenv("DEPLOYMENT_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetFleetDeployment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.GetFleetDeployment(ctx, DeploymentID, *datadogV2.NewGetFleetDeploymentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.GetFleetDeployment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.GetFleetDeployment`:\n%s\n", responseContent)
}
