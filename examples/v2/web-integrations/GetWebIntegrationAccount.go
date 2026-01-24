// Get integration account returns "OK" response

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
	// there is a valid "web_integration_account" in the system
	WebIntegrationAccountDataID := os.Getenv("WEB_INTEGRATION_ACCOUNT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebIntegrationsApi(apiClient)
	resp, r, err := api.GetWebIntegrationAccount(ctx, "twilio", WebIntegrationAccountDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebIntegrationsApi.GetWebIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebIntegrationsApi.GetWebIntegrationAccount`:\n%s\n", responseContent)
}
