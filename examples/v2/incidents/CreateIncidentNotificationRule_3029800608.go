// Create incident notification rule returns "Created" response

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
Handles: []string{
"@test-email@company.com",
},
Visibility: datadogV2.INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ORGANIZATION.Ptr(),
Trigger: "incident_created_trigger",
Enabled: datadog.PtrBool(true),
},
Relationships: &datadogV2.IncidentNotificationRuleCreateDataRelationships{
IncidentType: &datadogV2.RelationshipToIncidentType{
Data: datadogV2.RelationshipToIncidentTypeData{
Id: IncidentTypeDataID,
Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
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