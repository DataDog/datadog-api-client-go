// Create a status page email subscription returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.IncidentStatuspageSubscriptionRequest{
		Data: datadogV2.IncidentStatuspageSubscriptionDataRequest{
			Attributes: datadogV2.IncidentStatuspageSubscriptionDataAttributesRequest{
				Email: "user@example.com",
			},
			Type: datadogV2.INCIDENTSTATUSPAGESUBSCRIPTIONTYPE_STATUSPAGE_EMAIL_SUBSCRIPTION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateStatuspageEmailSubscription", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateStatuspageEmailSubscription(ctx, "page_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateStatuspageEmailSubscription`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateStatuspageEmailSubscription`:\n%s\n", responseContent)
}
