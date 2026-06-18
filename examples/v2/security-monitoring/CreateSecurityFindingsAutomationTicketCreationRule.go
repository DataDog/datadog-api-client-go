// Create a ticket creation rule returns "Successfully created the ticket creation rule" response

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
	body := datadogV2.TicketCreationRuleCreateRequest{
		Data: datadogV2.TicketCreationRuleDataCreate{
			Attributes: datadogV2.TicketCreationRuleAttributesCreate{
				Action: datadogV2.TicketCreationRuleAction{
					MaxTicketsPerDay: 10,
					ProjectId:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Target:           datadogV2.TICKETCREATIONTARGET_JIRA,
				},
				Enabled: datadog.PtrBool(true),
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
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityFindingsAutomationTicketCreationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFindingsAutomationTicketCreationRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFindingsAutomationTicketCreationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFindingsAutomationTicketCreationRule`:\n%s\n", responseContent)
}
