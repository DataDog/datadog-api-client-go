// Delete a schedule returns "Schedule successfully deleted." response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteFleetSchedule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	r, err := api.DeleteFleetSchedule(ctx, "id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.DeleteFleetSchedule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
