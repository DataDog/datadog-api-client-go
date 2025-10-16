// Delete incident notification template returns "No Content" response

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


	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteIncidentNotificationTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	r, err := api.DeleteIncidentNotificationTemplate(ctx, NotificationTemplateDataID, *datadogV2.NewDeleteIncidentNotificationTemplateOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncidentNotificationTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}