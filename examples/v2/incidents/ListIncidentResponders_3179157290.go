// List incident responders includes users and role assignments

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
	configuration.SetUnstableOperationEnabled("v2.ListIncidentResponders", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.ListIncidentResponders(ctx, IncidentWithCommanderDataID, *datadogV2.NewListIncidentRespondersOptionalParameters().WithInclude([]datadogV2.IncidentResponderRelatedObject{
		datadogV2.INCIDENTRESPONDERRELATEDOBJECT_USERS,
		datadogV2.INCIDENTRESPONDERRELATEDOBJECT_ROLE_ASSIGNMENTS,
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidentResponders`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListIncidentResponders`:\n%s\n", responseContent)
}
