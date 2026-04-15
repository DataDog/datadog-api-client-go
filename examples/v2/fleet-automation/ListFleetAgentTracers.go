// List tracers for a specific agent returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListFleetAgentTracers", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.ListFleetAgentTracers(ctx, "agent_key", *datadogV2.NewListFleetAgentTracersOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.ListFleetAgentTracers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.ListFleetAgentTracers`:\n%s\n", responseContent)
}
