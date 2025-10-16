// Update an incident type returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "incident_type" in the system
	IncidentTypeDataID := os.Getenv("INCIDENT_TYPE_DATA_ID")


	body := datadogV2.IncidentTypePatchRequest{
Data: datadogV2.IncidentTypePatchData{
Id: IncidentTypeDataID,
Attributes: datadogV2.IncidentTypeUpdateAttributes{
Name: datadog.PtrString("Security Incident-updated"),
},
Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentType", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentType(ctx, IncidentTypeDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentType`:\n%s\n", responseContent)
}