// Create a fleet access scope returns "Created" response

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
	body := datadogV2.FleetAccessScopeWriteRequest{
		Data: datadogV2.FleetAccessScopeWrite{
			Type: datadogV2.FLEETACCESSSCOPERESOURCETYPE_ACCESS_SCOPE,
			Attributes: &datadogV2.FleetAccessScopeAttributes{
				Description: datadog.PtrString("Example-Fleet-Automation"),
				Query:       datadog.PtrString("env:prod"),
				Features: &datadogV2.FleetAccessScopeFeatures{
					Install:   datadog.PtrBool(false),
					Upgrade:   datadog.PtrBool(true),
					Configure: datadog.PtrBool(false),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateFleetAccessScope", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.CreateFleetAccessScope(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CreateFleetAccessScope`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.CreateFleetAccessScope`:\n%s\n", responseContent)
}
