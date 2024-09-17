// Get details of an incident service returns "OK" response

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
	// there is a valid "service" in the system
	ServiceDataID := os.Getenv("SERVICE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetIncidentService", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentServicesApi(apiClient)
	resp, r, err := api.GetIncidentService(ctx, ServiceDataID, *datadogV2.NewGetIncidentServiceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.GetIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.GetIncidentService`:\n%s\n", responseContent)
}
