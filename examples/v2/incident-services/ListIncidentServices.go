// Get a list of all incident services returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "service" in the system
	ServiceDataAttributesName := os.Getenv("SERVICE_DATA_ATTRIBUTES_NAME")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListIncidentServices", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentServicesApi(apiClient)
	resp, r, err := api.ListIncidentServices(ctx, *datadog.NewListIncidentServicesOptionalParameters().WithFilter(ServiceDataAttributesName))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.ListIncidentServices`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.ListIncidentServices`:\n%s\n", responseContent)
}
