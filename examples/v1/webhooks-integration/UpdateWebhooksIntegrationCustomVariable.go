// Update a custom variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "webhook_custom_variable" in the system
	WebhookCustomVariableName := os.Getenv("WEBHOOK_CUSTOM_VARIABLE_NAME")

	body := datadog.WebhooksIntegrationCustomVariableUpdateRequest{
		Value: datadog.PtrString("variable-updated"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable(ctx, WebhookCustomVariableName, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable`:\n%s\n", responseContent)
}
