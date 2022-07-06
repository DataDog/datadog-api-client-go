// Delete a webhook returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "webhook" in the system
	WebhookName := os.Getenv("WEBHOOK_NAME")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.WebhooksIntegrationApi(apiClient)
	r, err := api.DeleteWebhooksIntegration(ctx, WebhookName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.DeleteWebhooksIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
