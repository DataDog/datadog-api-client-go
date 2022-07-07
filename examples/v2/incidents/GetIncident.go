// Get the details of an incident returns "OK" response

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
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetIncident", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.IncidentsApi(apiClient)
	resp, r, err := api.GetIncident(ctx, IncidentDataID, *datadog.NewGetIncidentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncident`:\n%s\n", responseContent)
}
