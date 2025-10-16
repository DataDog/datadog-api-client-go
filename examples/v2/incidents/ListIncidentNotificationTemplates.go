// List incident notification templates returns "OK" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListIncidentNotificationTemplates", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.ListIncidentNotificationTemplates(ctx, *datadogV2.NewListIncidentNotificationTemplatesOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidentNotificationTemplates`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListIncidentNotificationTemplates`:\n%s\n", responseContent)
}