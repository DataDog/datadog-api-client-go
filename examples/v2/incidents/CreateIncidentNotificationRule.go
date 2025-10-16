// Create an incident notification rule returns "Created" response

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
	body := datadogV2.CreateIncidentNotificationRuleRequest{
Data: datadogV2.IncidentNotificationRuleCreateData{
Attributes: datadogV2.IncidentNotificationRuleCreateAttributes{
Conditions: []datadogV2.IncidentNotificationRuleConditionsItems{
{
Field: "severity",
Values: []string{
"SEV-1",
"SEV-2",
},
},
},
Enabled: datadog.PtrBool(true),
Handles: []string{
"@team-email@company.com",
"@slack-channel",
},
RenotifyOn: []string{
"status",
"severity",
},
Trigger: "incident_created_trigger",
Visibility: datadogV2.INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ORGANIZATION.Ptr(),
},
Relationships: &datadogV2.IncidentNotificationRuleCreateDataRelationships{
IncidentType: &datadogV2.RelationshipToIncidentType{
Data: datadogV2.RelationshipToIncidentTypeData{
Id: "00000000-0000-0000-0000-000000000000",
Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
},
},
NotificationTemplate: &datadogV2.RelationshipToIncidentNotificationTemplate{
Data: datadogV2.RelationshipToIncidentNotificationTemplateData{
Id: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
Type: datadogV2.INCIDENTNOTIFICATIONTEMPLATETYPE_NOTIFICATION_TEMPLATES,
},
},
},
Type: datadogV2.INCIDENTNOTIFICATIONRULETYPE_INCIDENT_NOTIFICATION_RULES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentNotificationRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentNotificationRule`:\n%s\n", responseContent)
}