// Create an incident user-defined field returns "CREATED" response

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
	body := datadogV2.IncidentUserDefinedFieldCreateRequest{
		Data: datadogV2.IncidentUserDefinedFieldCreateData{
			Attributes: datadogV2.IncidentUserDefinedFieldAttributesCreateRequest{
				Category:     *datadogV2.NewNullableIncidentUserDefinedFieldCategory(datadogV2.INCIDENTUSERDEFINEDFIELDCATEGORY_WHAT_HAPPENED.Ptr()),
				Collected:    *datadogV2.NewNullableIncidentUserDefinedFieldCollected(datadogV2.INCIDENTUSERDEFINEDFIELDCOLLECTED_ACTIVE.Ptr()),
				DefaultValue: *datadog.NewNullableString(datadog.PtrString("critical")),
				DisplayName:  datadog.PtrString("Root Cause"),
				Name:         "root_cause",
				Ordinal:      *datadog.NewNullableString(datadog.PtrString("1.5")),
				Required:     datadog.PtrBool(false),
				TagKey:       *datadog.NewNullableString(datadog.PtrString("datacenter")),
				Type:         datadogV2.INCIDENTUSERDEFINEDFIELDFIELDTYPE_TEXTBOX,
				ValidValues: []datadogV2.IncidentUserDefinedFieldValidValue{
					{
						Description:      datadog.PtrString("A critical severity incident."),
						DisplayName:      "Critical",
						ShortDescription: datadog.PtrString("Critical"),
						Value:            "critical",
					},
				},
			},
			Relationships: datadogV2.IncidentUserDefinedFieldCreateRelationships{
				IncidentType: datadogV2.RelationshipToIncidentType{
					Data: datadogV2.RelationshipToIncidentTypeData{
						Id:   "00000000-0000-0000-0000-000000000000",
						Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
					},
				},
			},
			Type: datadogV2.INCIDENTUSERDEFINEDFIELDTYPE_USER_DEFINED_FIELD,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentUserDefinedField", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentUserDefinedField(ctx, body, *datadogV2.NewCreateIncidentUserDefinedFieldOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentUserDefinedField`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentUserDefinedField`:\n%s\n", responseContent)
}
