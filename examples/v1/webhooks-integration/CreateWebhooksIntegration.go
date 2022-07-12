// Create a webhooks integration returns "OK" response

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
	body := datadog.WebhooksIntegration{
		Name: "Example-Create_a_webhooks_integration_returns_OK_response",
		Url:  "https://example.com/webhook",
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.CreateWebhooksIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateWebhooksIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.CreateWebhooksIntegration`:\n%s\n", responseContent)
}
