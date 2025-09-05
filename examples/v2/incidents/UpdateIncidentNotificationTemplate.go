// Update incident notification template returns "OK" response

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
	// there is a valid "notification_template" in the system
	NotificationTemplateDataID := uuid.MustParse(os.Getenv("NOTIFICATION_TEMPLATE_DATA_ID"))

	body := datadogV2.PatchIncidentNotificationTemplateRequest{
		Data: datadogV2.IncidentNotificationTemplateUpdateData{
			Attributes: &datadogV2.IncidentNotificationTemplateUpdateAttributes{
				Category: datadog.PtrString("update"),
				Content: datadog.PtrString(`Incident Status Update:

Title: Sample Incident Title
New Status: resolved
Severity: SEV-2
Services: web-service, database-service
Commander: John Doe

For more details, visit the incident page.`),
				Name:    datadog.PtrString("Example-Incident"),
				Subject: datadog.PtrString("Incident Update: Sample Incident Title - resolved"),
			},
			Id:   NotificationTemplateDataID,
			Type: datadogV2.INCIDENTNOTIFICATIONTEMPLATETYPE_NOTIFICATION_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentNotificationTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentNotificationTemplate(ctx, NotificationTemplateDataID, body, *datadogV2.NewUpdateIncidentNotificationTemplateOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentNotificationTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentNotificationTemplate`:\n%s\n", responseContent)
}
