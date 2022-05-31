// Get a list of all incident services returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "service" in the system
	ServiceDataAttributesName := os.Getenv("SERVICE_DATA_ATTRIBUTES_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("ListIncidentServices", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentServicesApi.ListIncidentServices(ctx, *datadog.NewListIncidentServicesOptionalParameters().WithFilter(ServiceDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.ListIncidentServices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.ListIncidentServices`:\n%s\n", responseContent)
}
