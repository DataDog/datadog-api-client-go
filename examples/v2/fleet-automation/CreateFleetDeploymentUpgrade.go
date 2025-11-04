// Upgrade hosts returns "CREATED" response

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
	body := datadogV2.FleetDeploymentPackageUpgradeCreateRequest{
		Data: datadogV2.FleetDeploymentPackageUpgradeCreate{
			Attributes: datadogV2.FleetDeploymentPackageUpgradeAttributes{
				FilterQuery: datadog.PtrString("env:prod AND service:web"),
				TargetPackages: []datadogV2.FleetDeploymentPackage{
					{
						Name:    "datadog-agent",
						Version: "7.52.0",
					},
				},
			},
			Type: datadogV2.FLEETDEPLOYMENTRESOURCETYPE_DEPLOYMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateFleetDeploymentUpgrade", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.CreateFleetDeploymentUpgrade(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CreateFleetDeploymentUpgrade`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.CreateFleetDeploymentUpgrade`:\n%s\n", responseContent)
}
