// Create incident notification template returns "Created" response

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


	body := datadogV2.CreateIncidentNotificationTemplateRequest{
Data: datadogV2.IncidentNotificationTemplateCreateData{
Attributes: datadogV2.IncidentNotificationTemplateCreateAttributes{
Category: "alert",
Content: `An incident has been declared.

Title: Sample Incident Title
Severity: SEV-2
Affected Services: web-service, database-service
Status: active

Please join the incident channel for updates.`,
Name: "Example-Incident",
Subject: "SEV-2 Incident: Sample Incident Title",
},
Relationships: &datadogV2.IncidentNotificationTemplateCreateDataRelationships{
IncidentType: &datadogV2.RelationshipToIncidentType{
Data: datadogV2.RelationshipToIncidentTypeData{
Id: IncidentTypeDataID,
Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
},
},
},
Type: datadogV2.INCIDENTNOTIFICATIONTEMPLATETYPE_NOTIFICATION_TEMPLATES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentNotificationTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentNotificationTemplate(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentNotificationTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentNotificationTemplate`:\n%s\n", responseContent)
}