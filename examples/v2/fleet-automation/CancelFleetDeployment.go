// Cancel a deployment returns "Deployment successfully canceled." response

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
	configuration.SetUnstableOperationEnabled("v2.CancelFleetDeployment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewFleetAutomationApi(apiClient)
	r, err := api.CancelFleetDeployment(ctx, "deployment_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAutomationApi.CancelFleetDeployment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
