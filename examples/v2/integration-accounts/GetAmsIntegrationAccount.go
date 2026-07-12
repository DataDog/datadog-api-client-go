// Get integration account returns "OK: The account details for the specified integration." response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	resp, r, err := api.GetAmsIntegrationAccount(ctx, "integration_name", "interface_id", "account_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.GetAmsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAccountsApi.GetAmsIntegrationAccount`:\n%s\n", responseContent)
}
