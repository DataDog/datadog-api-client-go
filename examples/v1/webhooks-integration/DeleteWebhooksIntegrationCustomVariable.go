// Delete a custom variable returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "webhook_custom_variable" in the system
	WebhookCustomVariableName := os.Getenv("WEBHOOK_CUSTOM_VARIABLE_NAME")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable(ctx, WebhookCustomVariableName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
