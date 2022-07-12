// Get a webhook integration returns "OK" response

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
	// there is a valid "webhook" in the system
	WebhookName := os.Getenv("WEBHOOK_NAME")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.GetWebhooksIntegration(ctx, WebhookName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.GetWebhooksIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.GetWebhooksIntegration`:\n%s\n", responseContent)
}
