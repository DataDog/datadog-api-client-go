// Get incident notification rule returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetIncidentNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.GetIncidentNotificationRule(ctx, NotificationRuleDataID, *datadogV2.NewGetIncidentNotificationRuleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncidentNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncidentNotificationRule`:\n%s\n", responseContent)
}
