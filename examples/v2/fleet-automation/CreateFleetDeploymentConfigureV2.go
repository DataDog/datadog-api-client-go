// Create a configuration deployment returns "OK" response

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
	body := datadogV2.FleetDeploymentConfigureV2CreateRequest{
		Data: datadogV2.FleetDeploymentConfigureV2Create{
			Attributes: datadogV2.FleetDeploymentConfigureV2Attributes{
				ConfigOperations: []datadogV2.FleetDeploymentOperation{
					{
						FileOp:   datadogV2.FLEETDEPLOYMENTFILEOP_MERGE_PATCH,
						FilePath: "/datadog.yaml",
						Patch: map[string]interface{}{
							"log_level": "info",
						},
					},
				},
				DryRun:      datadog.PtrBool(true),
				FilterQuery: "env:prod AND service:example-fleet-automation",
			},
			Type: datadogV2.FLEETDEPLOYMENTRESOURCETYPE_DEPLOYMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.CreateFleetDeploymentConfigureV2(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CreateFleetDeploymentConfigureV2`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.CreateFleetDeploymentConfigureV2`:\n%s\n", responseContent)
}
