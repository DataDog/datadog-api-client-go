// Delete an existing incident service returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "service" in the system
	ServiceDataID := os.Getenv("SERVICE_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentService", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.IncidentServicesApi(apiClient)
	r, err := api.DeleteIncidentService(ctx, ServiceDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.DeleteIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
