// Update incident notification rule returns "OK" response

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
	// there is a valid "notification_rule" in the system
	NotificationRuleDataID := uuid.MustParse(os.Getenv("NOTIFICATION_RULE_DATA_ID"))

	// there is a valid "incident_type" in the system
	IncidentTypeDataID := os.Getenv("INCIDENT_TYPE_DATA_ID")

	body := datadogV2.PutIncidentNotificationRuleRequest{
		Data: datadogV2.IncidentNotificationRuleUpdateData{
			Attributes: datadogV2.IncidentNotificationRuleCreateAttributes{
				Enabled: datadog.PtrBool(false),
				Conditions: []datadogV2.IncidentNotificationRuleConditionsItems{
					{
						Field: "severity",
						Values: []string{
							"SEV-1",
						},
					},
				},
				Handles: []string{
					"@updated-team-email@company.com",
				},
				Visibility: datadogV2.INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_PRIVATE.Ptr(),
				Trigger:    "incident_modified_trigger",
			},
			Relationships: &datadogV2.IncidentNotificationRuleCreateDataRelationships{
				IncidentType: &datadogV2.RelationshipToIncidentType{
					Data: datadogV2.RelationshipToIncidentTypeData{
						Id:   IncidentTypeDataID,
						Type: datadogV2.INCIDENTTYPETYPE_INCIDENT_TYPES,
					},
				},
			},
			Id:   NotificationRuleDataID,
			Type: datadogV2.INCIDENTNOTIFICATIONRULETYPE_INCIDENT_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentNotificationRule(ctx, NotificationRuleDataID, body, *datadogV2.NewUpdateIncidentNotificationRuleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentNotificationRule`:\n%s\n", responseContent)
}
