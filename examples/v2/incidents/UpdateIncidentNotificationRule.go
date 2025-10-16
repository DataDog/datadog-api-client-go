// Update an incident notification rule returns "OK" response

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
	body := datadogV2.PutIncidentNotificationRuleRequest{
Data: datadogV2.IncidentNotificationRuleUpdateData{
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
Id: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentNotificationRule(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000001"), body, *datadogV2.NewUpdateIncidentNotificationRuleOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentNotificationRule`:\n%s\n", responseContent)
}