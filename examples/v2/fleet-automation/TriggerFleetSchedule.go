// Trigger a schedule deployment returns "CREATED - Deployment successfully created and started." response

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
	configuration.SetUnstableOperationEnabled("v2.TriggerFleetSchedule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	resp, r, err := api.TriggerFleetSchedule(ctx, "id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.TriggerFleetSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `FleetAutomationApi.TriggerFleetSchedule`:\n%s\n", responseContent)
}
