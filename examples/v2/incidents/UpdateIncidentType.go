// Update an incident type returns "OK" response

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
	body := datadogV2.IncidentTypePatchRequest{
		Data: datadogV2.IncidentTypePatchData{
			Attributes: datadogV2.IncidentTypeUpdateAttributes{
				Description: datadog.PtrString("Any incidents that harm (or have the potential to) the confidentiality, integrity, or availability of our data. Note: This will notify the security team."),
				IsDefault:   datadog.PtrBool(true),
				Name:        datadog.PtrString("Security Incident"),
			},
			Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentType", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentType(ctx, "incident_type_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentType`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentType`:\n%s\n", responseContent)
}
