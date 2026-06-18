// Reorder ticket creation rules returns "Successfully reordered the ticket creation rules" response

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

	body := datadogV2.TicketCreationRuleReorderRequest{
		Data: []datadogV2.TicketCreationRuleReorderItem{
			{
				Id:   ValidTicketCreationRuleDataID,
				Type: datadogV2.TICKETCREATIONRULETYPE_TICKET_CREATION_RULES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ReorderSecurityFindingsAutomationTicketCreationRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ReorderSecurityFindingsAutomationTicketCreationRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ReorderSecurityFindingsAutomationTicketCreationRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ReorderSecurityFindingsAutomationTicketCreationRules`:\n%s\n", responseContent)
}
