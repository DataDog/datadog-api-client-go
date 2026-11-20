// Get the details of an incident includes responders with modification timestamps

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
	// there is a valid "incident_with_commander" in the system
	IncidentWithCommanderDataID := os.Getenv("INCIDENT_WITH_COMMANDER_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncident", true)
	configuration.SetUnstableOperationEnabled("v2.GetIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.GetIncident(ctx, IncidentWithCommanderDataID, *datadogV2.NewGetIncidentOptionalParameters().WithInclude([]datadogV2.IncidentRelatedObject{
		datadogV2.INCIDENTRELATEDOBJECT_RESPONDERS,
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncident`:\n%s\n", responseContent)
}
