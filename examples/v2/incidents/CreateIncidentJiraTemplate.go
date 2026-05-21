// Create an incident Jira template returns "Created" response

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
	body := datadogV2.IncidentJiraTemplateRequest{
		Data: datadogV2.IncidentJiraTemplateDataRequest{
			Attributes: datadogV2.IncidentJiraTemplateDataAttributesRequest{
				AccountId: "123456",
				FieldConfigurations: []datadogV2.IncidentJiraTemplateFieldConfiguration{
					{
						IncidentField: *datadog.NewNullableString(datadog.PtrString("title")),
						JiraFieldKey:  "summary",
						JiraFieldType: *datadog.NewNullableString(datadog.PtrString("string")),
						SyncDirection: "bidirectional",
					},
				},
				IsDefault:   datadog.PtrBool(false),
				IssueId:     "10001",
				Name:        datadog.PtrString("Default Jira Template"),
				ProjectId:   "10000",
				ProjectKey:  "INC",
				SyncEnabled: datadog.PtrBool(true),
				Type:        datadog.PtrString("jira"),
			},
			Relationships: &datadogV2.IncidentJiraTemplateRelationships{
				IncidentType: &datadogV2.IncidentJiraTemplateIncidentTypeRelationship{
					Data: datadogV2.IncidentJiraTemplateIncidentTypeRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
						Type: "incident_types",
					},
				},
			},
			Type: datadogV2.INCIDENTJIRATEMPLATETYPE_INCIDENTS_JIRA_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentJiraTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentJiraTemplate(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentJiraTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentJiraTemplate`:\n%s\n", responseContent)
}
