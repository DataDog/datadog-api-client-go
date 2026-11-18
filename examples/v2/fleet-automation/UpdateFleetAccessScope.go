// Replace a fleet access scope returns "OK" response

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
	// there is a valid "fleet_access_scope" in the system
	FleetAccessScopeID := os.Getenv("FLEET_ACCESS_SCOPE_ID")

	body := datadogV2.FleetAccessScopeWriteRequest{
		Data: datadogV2.FleetAccessScopeWrite{
			Type: datadogV2.FLEETACCESSSCOPERESOURCETYPE_ACCESS_SCOPE,
			Attributes: &datadogV2.FleetAccessScopeAttributes{
				Description: datadog.PtrString("Example-Fleet-Automation updated"),
				Query:       datadog.PtrString("env:staging"),
				Features: &datadogV2.FleetAccessScopeFeatures{
					Install:   datadog.PtrBool(false),
					Upgrade:   datadog.PtrBool(false),
					Configure: datadog.PtrBool(true),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFleetAccessScope", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.UpdateFleetAccessScope(ctx, FleetAccessScopeID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.UpdateFleetAccessScope`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.UpdateFleetAccessScope`:\n%s\n", responseContent)
}
