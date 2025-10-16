// Get incident notification template returns "OK" response

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
	// there is a valid "notification_template" in the system
	NotificationTemplateDataID := uuid.MustParse(os.Getenv("NOTIFICATION_TEMPLATE_DATA_ID"))


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetIncidentNotificationTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.GetIncidentNotificationTemplate(ctx, NotificationTemplateDataID, *datadogV2.NewGetIncidentNotificationTemplateOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncidentNotificationTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncidentNotificationTemplate`:\n%s\n", responseContent)
}