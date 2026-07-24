// Create postmortem template returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.PostmortemTemplateRequest{
		Data: datadogV2.PostmortemTemplateDataRequest{
			Attributes: datadogV2.PostmortemTemplateAttributesRequest{
				ConfluencePostmortemSettings: &datadogV2.ConfluencePostmortemSettings{
					AccountId: "123456",
					ParentId:  *datadog.NewNullableString(datadog.PtrString("345678")),
					SpaceId:   "789012",
				},
				Content: datadog.PtrString(`# Overview

# What Happened

# Timeline

# Action Items`),
				GoogleDocsPostmortemSettings: &datadogV2.GoogleDocsPostmortemSettings{
					AccountId:      "123456",
					ParentFolderId: "789012",
				},
				IsDefault: *datadog.NewNullableTime(datadog.PtrTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))),
				Location:  datadogV2.POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS.Ptr(),
				Name:      "Standard Postmortem Template",
			},
			Id: datadog.PtrString("00000000-0000-0000-0000-000000000000"),
			Relationships: &datadogV2.PostmortemTemplateCreateRelationships{
				IncidentType: &datadogV2.PostmortemTemplateIncidentTypeRelationship{
					Data: datadogV2.PostmortemTemplateIncidentTypeRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000009"),
						Type: "incident_types",
					},
				},
			},
			Type: datadogV2.POSTMORTEMTEMPLATETYPE_POSTMORTEM_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentPostmortemTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentPostmortemTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentPostmortemTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentPostmortemTemplate`:\n%s\n", responseContent)
}
