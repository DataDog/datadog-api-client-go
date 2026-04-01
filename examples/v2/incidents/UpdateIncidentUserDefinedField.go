// Update an incident user-defined field returns "OK" response

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
	body := datadogV2.IncidentUserDefinedFieldUpdateRequest{
		Data: datadogV2.IncidentUserDefinedFieldUpdateData{
			Attributes: datadogV2.IncidentUserDefinedFieldAttributesUpdateRequest{
				Category:     *datadogV2.NewNullableIncidentUserDefinedFieldCategory(datadogV2.INCIDENTUSERDEFINEDFIELDCATEGORY_WHAT_HAPPENED.Ptr()),
				Collected:    *datadogV2.NewNullableIncidentUserDefinedFieldCollected(datadogV2.INCIDENTUSERDEFINEDFIELDCOLLECTED_ACTIVE.Ptr()),
				DefaultValue: *datadog.NewNullableString(datadog.PtrString("critical")),
				DisplayName:  datadog.PtrString("Root Cause"),
				Ordinal:      *datadog.NewNullableString(datadog.PtrString("1.5")),
				Required:     *datadog.NewNullableBool(datadog.PtrBool(false)),
				ValidValues: []datadogV2.IncidentUserDefinedFieldValidValue{
					{
						Description:      datadog.PtrString("A critical severity incident."),
						DisplayName:      "Critical",
						ShortDescription: datadog.PtrString("Critical"),
						Value:            "critical",
					},
				},
			},
			Id:   "00000000-0000-0000-0000-000000000000",
			Type: datadogV2.INCIDENTUSERDEFINEDFIELDTYPE_USER_DEFINED_FIELD,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentUserDefinedField", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentUserDefinedField(ctx, "00000000-0000-0000-0000-000000000000", body, *datadogV2.NewUpdateIncidentUserDefinedFieldOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentUserDefinedField`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentUserDefinedField`:\n%s\n", responseContent)
}
