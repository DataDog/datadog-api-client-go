// Update an incident user-defined role returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.IncidentUserDefinedRolePatchRequest{
		Data: datadogV2.IncidentUserDefinedRolePatchDataRequest{
			Attributes: &datadogV2.IncidentUserDefinedRolePatchDataAttributesRequest{
				Description: *datadog.NewNullableString(datadog.PtrString("The technical lead for the incident.")),
				Name:        datadog.PtrString("Tech Lead"),
				Policy: &datadogV2.IncidentUserDefinedRolePolicy{
					IsSingle: true,
				},
			},
			Id:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Type: datadogV2.INCIDENTUSERDEFINEDROLETYPE_INCIDENT_USER_DEFINED_ROLES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentUserDefinedRole", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentUserDefinedRole(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000002"), body, *datadogV2.NewUpdateIncidentUserDefinedRoleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentUserDefinedRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentUserDefinedRole`:\n%s\n", responseContent)
}
