// Create an incident type returns "CREATED" response

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
	body := datadogV2.IncidentTypeCreateRequest{
Data: datadogV2.IncidentTypeCreateData{
Attributes: datadogV2.IncidentTypeAttributes{
Description: datadog.PtrString("Any incidents that harm (or have the potential to) the confidentiality, integrity, or availability of our data."),
IsDefault: datadog.PtrBool(false),
Name: "Security Incident",
},
Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentType", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentType(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentType`:\n%s\n", responseContent)
}