// Get an incident notification rule returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetIncidentNotificationRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.GetIncidentNotificationRule(ctx, uuid.MustParse("00000000-0000-0000-0000-000000000001"), *datadogV2.NewGetIncidentNotificationRuleOptionalParameters(), )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncidentNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncidentNotificationRule`:\n%s\n", responseContent)
}