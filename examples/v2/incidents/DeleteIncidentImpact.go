// Delete an incident impact returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// the "incident" has an "incident_impact"
	IncidentImpactDataID := os.Getenv("INCIDENT_IMPACT_DATA_ID")
	IncidentImpactDataRelationshipsIncidentDataID := os.Getenv("INCIDENT_IMPACT_DATA_RELATIONSHIPS_INCIDENT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentImpact", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	r, err := api.DeleteIncidentImpact(ctx, IncidentImpactDataRelationshipsIncidentDataID, IncidentImpactDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncidentImpact`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
