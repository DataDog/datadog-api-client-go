// Update a custom variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "webhook_custom_variable" in the system
	WebhookCustomVariableName := os.Getenv("WEBHOOK_CUSTOM_VARIABLE_NAME")

	body := datadog.WebhooksIntegrationCustomVariableUpdateRequest{
		Value: common.PtrString("variable-updated"),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.UpdateWebhooksIntegrationCustomVariable(ctx, WebhookCustomVariableName, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable`:\n%s\n", responseContent)
}
