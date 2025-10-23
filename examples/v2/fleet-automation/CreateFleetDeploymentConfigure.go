// Create a deployment returns "CREATED" response

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
	body := datadogV2.FleetDeploymentConfigureCreateRequest{
		Data: datadogV2.FleetDeploymentConfigureCreate{
			Attributes: datadogV2.FleetDeploymentConfigureAttributes{
				ConfigOperations: []datadogV2.FleetDeploymentOperation{
					{
						FileOp:   datadogV2.FLEETDEPLOYMENTFILEOP_MERGE_PATCH,
						FilePath: "/datadog.yaml",
						Patch: map[string]interface{}{
							"apm_config":   "{'enabled': True}",
							"log_level":    "debug",
							"logs_enabled": "True",
						},
					},
				},
				FilterQuery: datadog.PtrString("env:prod AND service:web"),
			},
			Type: datadogV2.FLEETDEPLOYMENTRESOURCETYPE_DEPLOYMENT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateFleetDeploymentConfigure", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.CreateFleetDeploymentConfigure(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CreateFleetDeploymentConfigure`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.CreateFleetDeploymentConfigure`:\n%s\n", responseContent)
}
