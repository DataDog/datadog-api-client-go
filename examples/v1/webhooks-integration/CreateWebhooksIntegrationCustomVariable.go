// Create a custom variable returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.WebhooksIntegrationCustomVariable{
		IsSecret: true,
		Name:     "EXAMPLECREATEACUSTOMVARIABLERETURNSOKRESPONSE",
		Value:    "CUSTOM_VARIABLE_VALUE",
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewWebhooksIntegrationApi(apiClient)
	resp, r, err := api.CreateWebhooksIntegrationCustomVariable(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable`:\n%s\n", responseContent)
}
