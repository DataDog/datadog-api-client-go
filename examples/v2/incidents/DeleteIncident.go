// Delete an existing incident returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "incident" in the system
	INCIDENT_DATA_ID := os.Getenv("INCIDENT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("DeleteIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.IncidentsApi.DeleteIncident(ctx, INCIDENT_DATA_ID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
