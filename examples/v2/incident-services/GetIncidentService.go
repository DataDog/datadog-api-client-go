// Get details of an incident service returns "OK" response

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
	ServiceDataID := os.Getenv("SERVICE_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetIncidentService", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.IncidentServicesApi(apiClient)
	resp, r, err := api.GetIncidentService(ctx, ServiceDataID, *datadog.NewGetIncidentServiceOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.GetIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.GetIncidentService`:\n%s\n", responseContent)
}
