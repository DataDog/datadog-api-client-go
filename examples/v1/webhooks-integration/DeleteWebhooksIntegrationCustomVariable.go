// Delete a custom variable returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "webhook_custom_variable" in the system
	WebhookCustomVariableName := os.Getenv("WEBHOOK_CUSTOM_VARIABLE_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewWebhooksIntegrationApi(apiClient)
	r, err := api.DeleteWebhooksIntegrationCustomVariable(ctx, WebhookCustomVariableName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
