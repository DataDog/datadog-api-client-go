// Create a webhooks integration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.WebhooksIntegration{
		Name: "Example-Webhooks-Integration",
		Url:  "https://example.com/webhook",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.CreateWebhooksIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateWebhooksIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.CreateWebhooksIntegration`:\n%s\n", responseContent)
}
