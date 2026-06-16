// Get agentless host facet info returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetCSMAgentlessHostFacetInfo", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCSMSettingsApi(apiClient)
	resp, r, err := api.GetCSMAgentlessHostFacetInfo(ctx, "cloud_provider", *datadogV2.NewGetCSMAgentlessHostFacetInfoOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CSMSettingsApi.GetCSMAgentlessHostFacetInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CSMSettingsApi.GetCSMAgentlessHostFacetInfo`:\n%s\n", responseContent)
}
