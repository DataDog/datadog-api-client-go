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
	body := datadogV2.FleetDeploymentPackageUpgradeV2CreateRequest{
		Data: datadogV2.FleetDeploymentPackageUpgradeV2Create{
			Attributes: datadogV2.FleetDeploymentPackageUpgradeV2Attributes{
				FilterQuery: "env:prod AND service:example-fleet-automation",
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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.CreateFleetDeploymentUpgradeV2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CreateFleetDeploymentUpgradeV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.CreateFleetDeploymentUpgradeV2`:\n%s\n", responseContent)
}
