// Update a ticket creation rule returns "Successfully updated the ticket creation rule" response

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
	// there is a valid "valid_ticket_creation_rule" in the system
	ValidTicketCreationRuleDataID := uuid.MustParse(os.Getenv("VALID_TICKET_CREATION_RULE_DATA_ID"))

	body := datadogV2.TicketCreationRuleUpdateRequest{
		Data: datadogV2.TicketCreationRuleDataCreate{
			Attributes: datadogV2.TicketCreationRuleAttributesCreate{
				Action: datadogV2.TicketCreationRuleAction{
					MaxTicketsPerDay: 5,
					ProjectId:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Target:           datadogV2.TICKETCREATIONTARGET_JIRA,
				},
				Enabled: datadog.PtrBool(false),
				Name:    "Example-Security-Monitoring",
				Rule: datadogV2.AutomationRuleScope{
					FindingTypes: []datadogV2.SecurityFindingType{
						datadogV2.SECURITYFINDINGTYPE_MISCONFIGURATION,
					},
					Query: datadog.PtrString("env:staging"),
				},
			},
			Type: datadogV2.TICKETCREATIONRULETYPE_TICKET_CREATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityFindingsAutomationTicketCreationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityFindingsAutomationTicketCreationRule(ctx, ValidTicketCreationRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityFindingsAutomationTicketCreationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityFindingsAutomationTicketCreationRule`:\n%s\n", responseContent)
}
